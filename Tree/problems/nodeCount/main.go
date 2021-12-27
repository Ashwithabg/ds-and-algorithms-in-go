package main

import (
	"algos/Tree/utils"
	"fmt"
)

func main() {
	root := utils.GetNumbericRootNode()
	counter := 0
	traverse(root, &counter)
	fmt.Println(counter)
}

func traverse(root *utils.Node, counter *int) {
	if root == nil {
		return
	}
	*counter++

	traverse(root.GetRightNode(), counter)
	traverse(root.GetLeftNode(), counter)
}
