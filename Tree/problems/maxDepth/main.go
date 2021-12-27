package main

import (
	"algos/Tree/utils"
	"fmt"
)

func main() {
	root := utils.GetNumbericRootNode()
	fmt.Println(maxDepth(root))
}

func maxDepth(root *utils.Node) int {
	if root == nil {
		return 0
	}

	lDepth := maxDepth(root.GetLeftNode())
	rDepth := maxDepth(root.GetRightNode())

	if lDepth > rDepth {
		return lDepth + 1
	}

	return rDepth + 1
}
