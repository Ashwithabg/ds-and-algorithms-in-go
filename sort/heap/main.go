package main

import (
	"algos/heap"
	"fmt"
)

func main() {
	unsortedNumbers := []int{4, 6, 9, 2, 10, 56, 12, 5, 1, 17, 14}
	endIndex := len(unsortedNumbers) - 1

	newHeap := heap.NewHeap(unsortedNumbers, 11)
	newHeap.Heapify(endIndex)
	for endIndex > 0 {
		newHeap.Swap(0, endIndex)
		endIndex--
		newHeap.PercolateDown(0, endIndex)
	}

	fmt.Printf("%+v\n", newHeap.Data())
}
