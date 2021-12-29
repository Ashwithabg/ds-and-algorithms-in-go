package main

import (
	"algos/Tree/utils"
	"fmt"
)

/*
 Given a tree and a sum, return true if there is a path from the root
 down to a leaf, such that adding up all the values along the path
 equals the given sum.

 Strategy: subtract the node value from the sum when recurring down,
 and check to see if the sum is 0 when you run out of tree.
 */

func main() {
	root := utils.GetNumbericRootNode()
	sum := 21
	fmt.Println(hasSumPath(root, sum))
}

func hasSumPath(node *utils.Node, sum int) bool {
	if node == nil {
		return sum==0
	}

	subSum := sum - node.Value.(int)
	return hasSumPath(node.GetLeftNode(), subSum) ||
		hasSumPath(node.GetRightNode(), subSum)
}