package utils

import (
	"container/list"
)

type IntQueue struct {
	*list.List
}

func NewIntegerQueue() IntQueue {
	return IntQueue{
		list.New(),
	}
}

func (i *IntQueue) Enqueue(value int) {
	i.PushBack(value)
}

func (i *IntQueue) Dequeue() int {
	front := i.Front()
	i.Remove(front)
	return front.Value.(int)
}

func (i *IntQueue) IsEmpty() bool {
	return i.Len() == 0
}
