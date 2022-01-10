package main

import "container/heap"

type priorityQueue []*weightedVector

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].weight > p[j].weight
}

func (p priorityQueue) Swap(i, j int) {
	temp := p[i]
	p[i] = p[j]
	p[j] = temp
	p[i].index = j
	p[j].index = i
}

func (p *priorityQueue) Push(x interface{}) {
	n := len(*p)
	item := x.(*weightedVector)
	item.index = n
	*p = append(*p, item)
	heap.Fix(p, n)
}

func (p *priorityQueue) Pop() interface{} {
	n := len(*p)
	item := (*p)[n-1]
	*p = (*p)[:n-1]
	return item
}

func newPriorityQueue() heap.Interface {
	return &priorityQueue{}
}
