package main

import "fmt"

func main() {
	num := []int{50, 4, 6, 17, 8}
	fmt.Println(sort(num))
}

func sort(numbers []int) []int {
	for unsortedIndex := 1; unsortedIndex < len(numbers); unsortedIndex++ {
		sortedIndex := unsortedIndex-1
		currElement := numbers[unsortedIndex]

		for sortedIndex >= 0 && currElement <= numbers[sortedIndex] {
			numbers[sortedIndex+1] = numbers[sortedIndex]
			sortedIndex--
		}

		numbers[sortedIndex+1] = currElement
	}

	return numbers
}
