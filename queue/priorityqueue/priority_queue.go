package priorityqueue

import "container/heap"

type priorityQueue []int

func (h priorityQueue) Len() int           { return len(h) }
func (h priorityQueue) Less(i, j int) bool { return h[i] < h[j] }
func (h priorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *priorityQueue) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *priorityQueue) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func NewPriorityQueue(arr []int) *priorityQueue {
	p := priorityQueue(arr)
	heap.Init(&p)
	return &p
}
