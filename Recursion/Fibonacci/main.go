package main

import "fmt"

func main() {
	fmt.Println(fibonacci(5, 0, 1))
}

func fibonacci(n int, num1 int, num2 int) int {
	if n == 0 {
		return num1
	}

	if n == 1 {
		return num2
	}

	return fibonacci(n-1, num2, num1 + num2)
}
