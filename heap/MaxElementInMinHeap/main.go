package main

import (
	"algos/heap"
	"algos/heap/utils"
	"fmt"
)

//Find the first leaf node and traverse
//Largest element will be present in leaf node only

func main() {
	minHeap := utils.GetMinHeap()
	fmt.Println(findMaxElement(minHeap))
}

func findMaxElement(minHeap *heap.Heap) int {
	lastIndex := minHeap.GetCount() - 1
	lastParentIndex := minHeap.GetParentIndex(lastIndex)
	maxElement := minHeap.GetElement(lastIndex)

	for index := lastIndex; index != lastParentIndex; index-- {
		if maxElement < minHeap.GetElement(index) {
			maxElement = minHeap.GetElement(index)
		}

	}

	return maxElement
}
