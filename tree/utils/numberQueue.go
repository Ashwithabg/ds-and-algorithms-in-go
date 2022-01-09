package utils

import (
	"container/list"
)

type IntQueue struct {
	values *list.List
}

func NewIntegerQueue() IntQueue {
	return IntQueue{values: list.New()}
}

func (i IntQueue) Enqueue(value int) {
	i.values.PushBack(value)
}

func (i IntQueue) Dequeue() int {
	front := i.values.Front()
	i.values.Remove(front)
	return front.Value.(int)
}

func (i IntQueue) IsEmpty() bool {
	return i.values.Len() == 0
}
