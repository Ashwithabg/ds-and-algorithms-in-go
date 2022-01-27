package orderStatistics

import "container/heap"

func getLargestElement(arr []int, k int) int {
	minHeap := &minHeap{}
	for i := 0; i <= k; i++ {
		minHeap.Push(arr[i])
	}
	heap.Init(minHeap)

	counter := k + 1
	for minHeap.Len() == k+1 {
		minHeap.Pop()
		if counter < len(arr) {
			minHeap.Push(arr[counter])
			counter++
		}

		heap.Init(minHeap)
	}

	return minHeap.Pop().(int)
}
