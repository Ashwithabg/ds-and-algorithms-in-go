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
	minHeap := heap.NewHeap(make([]int, 0, 5), 5)

	for _, value := range data {
		if minHeap.GetCount() == 0 || minHeap.GetCount() < count {
			minHeap.Insert(value)
		} else {
			if minHeap.GetElement(0) < value {
				minHeap.Remove()
				minHeap.Insert(value)
			}
		}
	}

	fmt.Println(minHeap.Data())
}
