package main

import "fmt"

/*
 For the key values 1...numKeys, how many structurally unique
 binary search trees are possible that store those keys.
 Strategy: consider that each value could be the root.
 Recursively find the size of the left and right subtrees.
*/

func main() {
	fmt.Println(countTrees(4))
}

func countTrees(number int) int {
	if number <= 1 {
		return 1
	}

	treeCounter, left, right := 0, 0, 0

	for root := 1; root <= number; root++ {
		left = countTrees(root - 1)
		right = countTrees(number - root)
		treeCounter += left * right
	}

	return treeCounter
}
