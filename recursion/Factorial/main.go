package main

import "fmt"

func main() {
	fmt.Println(factorial(5, 1))
}

func factorial(n int, acc int) int {
	if n == 1 {
		return acc
	}

	return factorial(n-1, n*acc)
}
