package jobqueue

import (
	"context"
	"log"
	"sync"
)

type (
	JobQueue struct {
		jobChan  chan Job
		workers  []*worker
		stopChan chan struct{}
		wg       *sync.WaitGroup
		ctx      context.Context

		workerNums int
		jobNums    int
		once       sync.Once
	}
	worker struct {
		id  int
		ctx context.Context
	}
)

type (
	Job         func()
	QueueOption func(*JobQueue)
)

const (
	defaultWorkerNums = 5
	defaultJobNums    = 10
)

func NewJobQueue(ctx context.Context, options ...QueueOption) *JobQueue {
	j := &JobQueue{
		stopChan:   make(chan struct{}, 1),
		wg:         &sync.WaitGroup{},
		ctx:        ctx,
		workerNums: defaultWorkerNums,
		jobNums:    defaultJobNums,
	}
	for _, o := range options {
		o(j)
	}
	for i := 0; i < j.workerNums; i++ {
		j.workers = append(j.workers, j.newWorker(i))
	}
	j.jobChan = make(chan Job, j.jobNums)
	j.start()
	return j
}

func SetWorkerNum(i int) QueueOption {
	return func(j *JobQueue) {
		j.workerNums = i
	}
}

func SetJobNumber(i int) QueueOption {
	return func(j *JobQueue) {
		j.jobNums = i
	}
}

// 执行一批任务，自动关闭队列
func (j *JobQueue) Run(funcs ...func()) {
	defer j.Close()
	for _, f := range funcs {
		j.Push(f)
	}
}

// 提交任务
func (j *JobQueue) Push(job Job) {
	defer rescue()
	select {
	case j.jobChan <- job:
	case <-j.ctx.Done():
	case <-j.stopChan:
		return
	}
}

// 关闭队列
func (j *JobQueue) Close() {
	j.once.Do(func() {
		close(j.jobChan)
	})
}

// 获取结束信号，用于堵塞主线程
func (j *JobQueue) StopChannel() <-chan struct{} {
	return j.stopChan
}

func (j *JobQueue) Wait() {
	select {
	case <-j.stopChan:
	case <-j.ctx.Done():
	}
}

func (j *JobQueue) newWorker(i int) *worker {
	return &worker{id: i, ctx: j.ctx}
}

func (j *JobQueue) start() {
	go func() {
		defer j.stop()
		for _, w := range j.workers {
			j.wg.Add(1)
			go w.start(j.jobChan, j.wg)
		}
		j.wg.Wait()
	}()
}

func (j *JobQueue) stop() {
	defer close(j.stopChan)
	j.stopChan <- struct{}{}
}

func (w *worker) start(jobChan chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobChan:
			if ok {
				run(job)
			} else {
				return
			}
		case <-w.ctx.Done():
			return
		}
	}
}

func run(f func()) {
	defer rescue()
	f()
}

func rescue()  {
	if err := recover(); err != nil {
		log.Println(err)
	}
}