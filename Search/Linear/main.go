package main

import "fmt"

func main() {
	numbers := []int{1, 6, 2, 5, 7, 8, 9}
	fmt.Println(search(numbers, 6))
	fmt.Println(search(numbers, 9))
	fmt.Println(search(numbers, 19))
}

func search(numbers []int, number int) int {
	for index, currentNumber := range numbers {
		if currentNumber == number {
			return index
		}
	}

	return -1
}
