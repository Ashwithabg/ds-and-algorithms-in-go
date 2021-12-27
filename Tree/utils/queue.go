package utils

import (
	"container/list"
)

type Queue struct {
	values *list.List
}

func NewQueue() Queue {
	return Queue{values: list.New()}
}

func (q Queue) Enqueue(node *Node) {

	q.values.PushBack(node)
}

func (q Queue) Dequeue() *Node {
	front := q.values.Front()
	q.values.Remove(front)

	return front.Value.(*Node)
}
func (q Queue) IsEmpty() bool {
	return q.values.Len() == 0
}
