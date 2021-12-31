package main

const MaxSize = 40

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
