package main

import "fmt"

const MaxSize = 40
const rootIndex = 0

type Heap struct {
	data []interface{}
}

func (h *Heap) getElement(index int) interface{} {
	if index > len(h.data) || index < 0 {
		return -1
	}

	return h.data[index]
}

func (h *Heap) getLeftChildIndex(index int) int {
	if index > len(h.data) || index < 0 {
		return -1
	}

	childIndex := 2*index + 1
	if childIndex > len(h.data) || childIndex < 0 {
		return -1
	}

	return childIndex
}

func (h *Heap) getRightChildIndex(index int) int {
	if index > len(h.data) || index < 0 {
		return -1
	}

	childIndex := 2*index + 2
	if childIndex > len(h.data) || childIndex < 0 {
		return -1
	}
	return childIndex
}

func (h *Heap) getParentIndex(index int) int {
	if index > len(h.data) || index < 0 {
		return -1
	}
	return (index - 1) / 2
}

func (h *Heap) swap(index1, index2 int) {
	temp := h.data[index1]
	h.data[index1] = h.data[index2]
	h.data[index2] = temp
}

func (h *Heap) getCount() int {
	return len(h.data)
}

func (h *Heap) isEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) isFull() bool {
	return len(h.data) == cap(h.data)
}

//min Heap
func (h *Heap) siftDown(index int) {
	leftIndex := h.getLeftChildIndex(index)
	rightIndex := h.getRightChildIndex(index)

	smallerIndex := -1

	if leftIndex != -1 && rightIndex != -1 {
		if h.getElement(leftIndex).(int) <= h.getElement(rightIndex).(int) {
			smallerIndex = leftIndex
		} else {
			smallerIndex = rightIndex
		}
	} else if leftIndex != -1 {
		smallerIndex = leftIndex
	} else if rightIndex != -1 {
		smallerIndex = rightIndex
	}
	if smallerIndex == -1 {
		return
	}

	if h.getElement(smallerIndex).(int) < h.getElement(index).(int) {
		h.swap(smallerIndex, index)
		h.siftDown(smallerIndex)
	}
}

//min Heap
func (h *Heap) siftUp(index int) {
	parentIndex := h.getParentIndex(index)
	if parentIndex != -1 && h.getElement(index).(int) < h.getElement(parentIndex).(int) {
		h.swap(index, parentIndex)
		h.siftUp(parentIndex)
	}
}

func (h *Heap) insert(value int) error {
	if h.isFull() {
		return fmt.Errorf("heap is full")
	}
	count := len(h.data)
	h.data = append(h.data, value)
	h.siftUp(count)

	return nil
}

func (h *Heap) remove() (int, error) {
	if h.isEmpty() {
		return rootIndex, fmt.Errorf("heap is empty")
	}
	highPriorityElement := h.data[rootIndex]
	lastIndex := len(h.data) - 1

	h.swap(rootIndex, lastIndex)
	h.data = h.data[:lastIndex]
	h.siftDown(rootIndex)

	return highPriorityElement.(int), nil
}

//maxHeap
func (h *Heap) findIndex(index, endIndex int, fn func(int) int) int {
	if index > endIndex || index == -1 {
		return -1
	}

	childIndex := fn(index)
	if childIndex > endIndex || childIndex < 0 {
		return -1
	}

	return childIndex
}

func (h *Heap) findLeftChildIndex(index, endIndex int) int {
	indexCalculator := func(index int) int {
		return 2*index + 1
	}
	return h.findIndex(index, endIndex, indexCalculator)
}

func (h *Heap) findRightChildIndex(index, endIndex int) int {
	indexCalculator := func(index int) int {
		return 2*index + 2
	}
	return h.findIndex(index, endIndex, indexCalculator)
}

func (h *Heap) findParentIndex(index, endIndex int) int {
	indexCalculator := func(index int) int {
		return (index - 1) / 2
	}
	return h.findIndex(index, endIndex, indexCalculator)
}

func (h *Heap) percolateDown(index, endIndex int) {
	leftChildIndex := h.findLeftChildIndex(index, endIndex)
	rightChildIndex := h.findRightChildIndex(index, endIndex)

	if leftChildIndex != -1 && h.getElement(leftChildIndex).(int) > h.getElement(index).(int) {
		h.swap(leftChildIndex, index)
		h.percolateDown(leftChildIndex, endIndex)
	}
	if rightChildIndex != -1 && h.getElement(rightChildIndex).(int) > h.getElement(index).(int) {
		h.swap(rightChildIndex, index)
		h.percolateDown(rightChildIndex, endIndex)
	}
}

func (h *Heap) Heapify(endIndex int) {
	parentIndex := h.findParentIndex(endIndex, endIndex)
	for parentIndex >= 0 {
		h.percolateDown(parentIndex, endIndex)
		parentIndex--
	}
}
