package main

import "fmt"

var dp = map[int]int{}

func main() {
	fmt.Println(solve(6))
	fmt.Println(solve(4))
}

func solve(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	val, ok := dp[n]
	if ok {
		return val
	}

	dp[n] = solve(n-1) + solve(n-3) + solve(n-5)
	return dp[n]
}
