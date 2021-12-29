package main

import (
	"algos/Tree/utils"
	"fmt"
)

func main() {
	root := utils.GetNumbericRootNode()
	printTree(root)
}

func printTree(node *utils.Node) {
	if node == nil {
		return
	}

	printTree(node.GetLeftNode())
	fmt.Printf("%d ", node.Value)
	printTree(node.GetRightNode())
}
