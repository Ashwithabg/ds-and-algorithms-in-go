package main

import "fmt"

func main() {
	num := []int{50, 4, 6, 17, 8}
	fmt.Println(sort(num))
}

func sort(numbers []int) []int {
	for j := 0; j < len(numbers); j++ {
		smallest := numbers[j]
		smallestIndex := j

		for i := j; i < len(numbers); i++ {
			if numbers[i] < smallest {
				smallest = numbers[i]
				smallestIndex = i
			}
		}

		swap(numbers, j, smallestIndex)
	}

	return numbers
}

func swap(numbers []int, i int, j int) {
	temp := numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp
}
