package orderStatistics

import (
	"container/heap"
	"sort"
)

//K’th Smallest/Largest Element in Unsorted Array
//Given an array and a number k where k is smaller than the size of the array, we need to find the k’th
//smallest element in the given array. It is given that all array elements are distinct.

func sortAndGetElement(arr []int, k int) int {
	sort.Ints(arr)
	return arr[k-1]
}

func getMinElementUsingMaxheap(arr []int, k int) int {
	maxHeap := &maxHeap{}
	for i := 0; i <= k; i++ {
		maxHeap.Push(arr[i])
	}
	heap.Init(maxHeap)

	counter := k + 1
	for maxHeap.Len() == k+1 {
		maxHeap.Pop()
		if counter != len(arr) {
			maxHeap.Push(arr[counter])
			counter++
		}

		heap.Init(maxHeap)
	}

	pop := maxHeap.Pop()
	return pop.(int)
}
