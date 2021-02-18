package jobqueue

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

var (
	ctx = context.Background()
)

type (
	testJob struct {
	}
)

func (t *testJob) Run() {
	fmt.Println(time.Now())
	time.Sleep(time.Millisecond * 100)
}

func NewTestJob() *testJob {
	fmt.Println("add job")
	return &testJob{}
}

func TestJobQueue_Close(t *testing.T) {
	queue := NewJobQueue(ctx, SetWorkerNum(5), SetJobNumber(50))
	tc := time.After(time.Second * 1)
OUT:
	for {
		select {
		case <-tc:
			queue.Close()
			break OUT
		default:
			fmt.Println(len(queue.jobChan), len(queue.workers))
			time.Sleep(time.Millisecond * 1)
			queue.Push(NewTestJob().Run)
		}
	}
	<-queue.StopChannel()
	fmt.Println("done")
}

func TestJobQueue_less(t *testing.T) {
	queue := NewJobQueue(ctx, SetWorkerNum(5), SetJobNumber(500))
	tc := time.After(time.Second * 1)
OUT:
	for {
		select {
		case <-tc:
			queue.Close()
			break OUT
		default:
			fmt.Println(len(queue.jobChan), len(queue.workers))
			time.Sleep(time.Millisecond * 10)
			queue.Push(NewTestJob().Run)
		}
	}
	<-queue.StopChannel()
	fmt.Println("done")
}

func TestJobQueue_more(t *testing.T) {
	queue := NewJobQueue(ctx, SetWorkerNum(50), SetJobNumber(500))
	tc := time.After(time.Second * 1)
OUT:
	for {
		select {
		case <-tc:
			queue.Close()
			break OUT
		default:
			fmt.Println(len(queue.jobChan), len(queue.workers))
			time.Sleep(time.Millisecond * 10)
			queue.Push(NewTestJob().Run)
		}
	}
	<-queue.StopChannel()
	fmt.Println("done")
}

func TestJobQueue_cancel(t *testing.T) {
	timeoutCtx, _ := context.WithTimeout(ctx, time.Second)
	queue := NewJobQueue(timeoutCtx, SetWorkerNum(50), SetJobNumber(500))
	for {
		select {
		case <-queue.StopChannel():
			goto DONE
		default:
			fmt.Println(len(queue.jobChan), len(queue.workers))
			time.Sleep(time.Millisecond * 10)
			queue.Push(NewTestJob().Run)
		}
	}
DONE:
	fmt.Println("done")
}

func TestJobQueue_Run(t *testing.T) {
	var testCases = []struct {
		workerNum int
		jobNum    int
		loop      int
	}{
		{
			workerNum: 100,
			jobNum:    1,
			loop:      1,
		},
		{
			workerNum: 1,
			jobNum:    1,
			loop:      10,
		},
		{
			workerNum: 10,
			jobNum:    5,
			loop:      100,
		},
		{
			workerNum: 5,
			jobNum:    10,
			loop:      100,
		},
		{
			workerNum: 500,
			jobNum:    10,
			loop:      1000,
		},
		{
			workerNum: 10,
			jobNum:    500,
			loop:      100,
		},
		{
			workerNum: 500,
			jobNum:    500,
			loop:      5,
		},
	}
	for _, c := range testCases {
		var count int32
		var job = func() {
			atomic.AddInt32(&count, 1)
		}
		queue := NewJobQueue(ctx, SetWorkerNum(c.workerNum), SetJobNumber(c.jobNum))
		var funcs []func()
		for i := 0; i < c.loop; i++ {
			funcs = append(funcs, func() {
				time.Sleep(time.Millisecond * 1)
				job()
			})
		}
		queue.Run(funcs...)
		<-queue.StopChannel()
		assert.Equal(t, int32(c.loop), count)
	}
	fmt.Println("done")
}

// 测试 goroutine 泄漏
func TestJobQueue_goroutine(t *testing.T) {
	var loop int = 1e4
	for i := 0; i < loop; i++ {
		fmt.Println(runtime.NumGoroutine())
		queue := NewJobQueue(ctx, SetWorkerNum(100), SetJobNumber(500))
		queue.Close()
	}
	for i := 0; i < loop; i++ {
		fmt.Println(runtime.NumGoroutine())
		cancelCtx, cancel := context.WithCancel(context.Background())
		NewJobQueue(cancelCtx, SetWorkerNum(100), SetJobNumber(50))
		cancel()
	}
}
