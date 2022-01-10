package main

import "container/list"

type queue struct {
	*list.List
}

func newQueue() *queue {
	return &queue{list.New()}
}

func (q *queue) enqueue(value int) {
	q.PushBack(value)
}

func (q *queue) dequeue() int {
	frontNode := q.Front()
	q.Remove(frontNode)
	return frontNode.Value.(int)
}

func (q queue) isEmpty() bool {
	return q.Len() == 0
}
