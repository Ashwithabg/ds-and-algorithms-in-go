package main

import "container/list"

type queue struct {
	*list.List
}

func (q *queue) enqueue(elem int) {
	q.PushBack(elem)
}

func (q *queue) dequeue() int {
	elem := q.Front()
	q.Remove(elem)
	return elem.Value.(int)
}

func (q *queue) isEmpty() bool {
	return q.Len() == 0
}


