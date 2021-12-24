package main

import "fmt"

func main() {
	numbers := []int{2, 16, 3, 15, 6, 10, -1}
	quickSort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
}

func quickSort(numbers []int, low int, high int) {
	if low < high {
		pivotIndex := partition(numbers, low, high)
		quickSort(numbers, low, pivotIndex-1)
		quickSort(numbers, pivotIndex+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivotIndex := high
	pivot := arr[high]
	smallerElementIndex := low

	for iteration := low; iteration <= high; iteration++ {
		if arr[iteration] < pivot {
			swap(arr, smallerElementIndex, iteration)
			smallerElementIndex++
		}
	}

	swap(arr, smallerElementIndex, pivotIndex)
	return smallerElementIndex
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
