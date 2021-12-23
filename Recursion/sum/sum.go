package main

import (
	"fmt"
	"time"
)

func main() {
	n := 5000000
	start := time.Now()
	fmt.Println(iterativeSum(n))
	fmt.Println("iterative time ", time.Now().Sub(start).Seconds())

	start = time.Now()
	fmt.Println(recursiveSum(n))
	fmt.Println("head recursion time ", time.Now().Sub(start).Seconds())

	start = time.Now()
	fmt.Println(tailRecursionSum(n, 0))
	fmt.Println("tail recursion time ", time.Now().Sub(start).Seconds())
}

func iterativeSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func recursiveSum(n int) int {
	if n == 1 {
		return 1
	}

	return n + recursiveSum(n-1)
}

func tailRecursionSum(n int, result int) int {
	if n == 1 {
		return result + 1
	}

	return tailRecursionSum(n-1, result+n)
}

