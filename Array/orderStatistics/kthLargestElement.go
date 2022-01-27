package orderStatistics

import "container/heap"

func getLargestElement(arr []int, k int) int {
	minHeap := &minHeap{}
	n := len(arr)
	for i := 0; i < n; i++ {
		minHeap.Push(arr[i])
		heap.Init(minHeap)

		if minHeap.Len() == k+1 {
			minHeap.Pop()
		}
	}

	return minHeap.Pop().(int)
}
