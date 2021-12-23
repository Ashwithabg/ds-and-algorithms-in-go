package main

import "fmt"

func main() {
	numbers := []int{1, 6, 2, 5, 7, 8, 9}
	fmt.Println(search(numbers, 6))
	fmt.Println(search(numbers, 9))
	fmt.Println(search(numbers, 19))

	fmt.Println(searchRecursion(numbers, 9, 0))
}

func search(numbers []int, number int) int {
	for index, currentNumber := range numbers {
		if currentNumber == number {
			return index
		}
	}

	return -1
}

func searchRecursion(numbers []int, number int, currentIndex int) int {
	if numbers[currentIndex] == number {
		return currentIndex
	}

	if currentIndex == len(numbers)-1 {
		return -1
	}

	return searchRecursion(numbers, number, currentIndex+1)
}
