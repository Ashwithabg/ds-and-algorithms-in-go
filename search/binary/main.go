package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(search(numbers, 6, 0, len(numbers)-1))
	fmt.Println(searchIterative(numbers, 6))
}

func search(numbers []int, number int, start int, end int) int {
	if start <= end {
		mid := (start + end) / 2
		if numbers[mid] == number {
			return mid
		}

		if number < numbers[mid] {
			return search(numbers, number, start, mid-1)
		} else {
			return search(numbers, number, mid+1, end)
		}
	}

	return -1
}

func searchIterative(numbers []int, number int) int {
	start := 0
	end := len(numbers)

	for start <= end {
		mid := (start + end) / 2
		if numbers[mid] == number {
			return mid
		}

		if number < numbers[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
