package main

import (
	"algos/Tree/utils"
	"fmt"
)

func main() {
	rootNode := utils.GetRootNode()
	preorder(rootNode)
}

func preorder(root *utils.Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)

	preorder(root.GetLeftNode())
	preorder(root.GetRightNode())
}
