package utils

import "algos/heap"

func GetMinHeap() *heap.Heap {
	data := []int{5, 8, 6, 9, 12, 11, 7, 15, 10}
	return heap.NewHeap(data)
}
