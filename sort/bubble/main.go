package main

import "fmt"

func main() {
	num := []int{50, 4, 6, 17, 8}
	fmt.Println(sort(num))
}

func sort(numbers []int) []int {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				swap(numbers, j, j+1)
			}
		}
	}

	return numbers
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
