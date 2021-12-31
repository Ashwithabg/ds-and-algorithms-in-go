package main

import (
	"algos/heap"
	"fmt"
)

func main() {
	data := []int{5, 8, 6, 9, 12, 11, 7, 15, 10}
	elementCount := 5
	findLargestElement(data, elementCount)
}

func findLargestElement(data []int, count int) {
	minHeap := heap.NewHeap([]int{}, 5)

	for _, value := range data {
		if minHeap.GetCount() == 0 || minHeap.GetCount() < count {
			minHeap.Insert(value)
		} else {
			if minHeap.GetElement(0) < value {
				minHeap.Update(value, 0)
				minHeap.SiftDown(0)
			}
		}
	}

	fmt.Println(minHeap.Data())
}
