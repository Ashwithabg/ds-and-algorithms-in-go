package heap

import "fmt"

const MaxCapacity = 40
const rootIndex = 0

type Heap struct {
	capacity int
	data     []int
}

func NewHeap(arr []int, capacity int) *Heap {
	if capacity == 0 {
		capacity = MaxCapacity
	}
	return &Heap{
		capacity: capacity,
		data:     arr,
	}
}

func (h *Heap) GetElement(index int) int {
	if index >= len(h.data) || index < 0 {
		return -1
	}

	return h.data[index]
}

func (h *Heap) getLeftChildIndex(index int) int {
	if index >= len(h.data) || index < 0 {
		return -1
	}

	childIndex := 2*index + 1
	if childIndex >= len(h.data) || childIndex < 0 {
		return -1
	}

	return childIndex
}

func (h *Heap) getRightChildIndex(index int) int {
	if index >= len(h.data) || index < 0 {
		return -1
	}

	childIndex := 2*index + 2
	if childIndex >= len(h.data) || childIndex < 0 {
		return -1
	}
	return childIndex
}

func (h *Heap) GetParentIndex(index int) int {
	if index > len(h.data) || index < 0 {
		return -1
	}
	return (index - 1) / 2
}

func (h *Heap) Swap(index1, index2 int) {
	temp := h.data[index1]
	h.data[index1] = h.data[index2]
	h.data[index2] = temp
}

func (h *Heap) GetCount() int {
	return len(h.data)
}

func (h *Heap) isEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) isFull() bool {
	return len(h.data) == h.capacity
}

//min Heap
func (h *Heap) SiftDown(index int) {
	leftIndex := h.getLeftChildIndex(index)
	rightIndex := h.getRightChildIndex(index)

	smallerIndex := -1

	if leftIndex != -1 && rightIndex != -1 {
		if h.GetElement(leftIndex) <= h.GetElement(rightIndex) {
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

	if h.GetElement(smallerIndex) < h.GetElement(index) {
		h.Swap(smallerIndex, index)
		h.SiftDown(smallerIndex)
	}
}

//min Heap
func (h *Heap) siftUp(index int) {
	parentIndex := h.GetParentIndex(index)
	if parentIndex != -1 && h.GetElement(index) < h.GetElement(parentIndex) {
		h.Swap(index, parentIndex)
		h.siftUp(parentIndex)
	}
}

func (h *Heap) Insert(value int) error {
	if h.isFull() {
		return fmt.Errorf("heap is full")
	}
	count := len(h.data)
	h.data = append(h.data, value)
	h.siftUp(count)

	return nil
}

func (h *Heap) Remove() (int, error) {
	if h.isEmpty() {
		return rootIndex, fmt.Errorf("heap is empty")
	}
	highPriorityElement := h.data[rootIndex]
	lastIndex := len(h.data) - 1

	h.Swap(rootIndex, lastIndex)
	h.data = h.data[:lastIndex]
	h.SiftDown(rootIndex)

	return highPriorityElement, nil
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

func (h *Heap) PercolateDown(index, endIndex int) {
	leftChildIndex := h.findLeftChildIndex(index, endIndex)
	rightChildIndex := h.findRightChildIndex(index, endIndex)

	if leftChildIndex != -1 && h.GetElement(leftChildIndex) > h.GetElement(index) {
		h.Swap(leftChildIndex, index)
		h.PercolateDown(leftChildIndex, endIndex)
	}
	if rightChildIndex != -1 && h.GetElement(rightChildIndex) > h.GetElement(index) {
		h.Swap(rightChildIndex, index)
		h.PercolateDown(rightChildIndex, endIndex)
	}
}

func (h *Heap) Heapify(endIndex int) {
	parentIndex := h.findParentIndex(endIndex, endIndex)
	for parentIndex >= 0 {
		h.PercolateDown(parentIndex, endIndex)
		parentIndex--
	}
}

func (h *Heap) Data() []int {
	return h.data
}
