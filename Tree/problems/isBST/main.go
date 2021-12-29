package main

import (
	"algos/Tree/utils"
	"fmt"
)

func main() {
	invalidBST := utils.NewNode(3)
	invalidBST.SetLeft(utils.NewNode(10))
	invalidBST.SetLeft(utils.NewNode(4))

	root := utils.GetNumbericRootNode()
	fmt.Println(isBST(root))
	fmt.Println(isBST(invalidBST))
}

func isBST(node *utils.Node) bool {
	if node == nil {
		return true
	}

	if node.GetLeftNode() != nil && node.GetLeftNode().Value.(int) > node.Value.(int) {
		return false
	}

	if node.GetRightNode() != nil && node.GetRightNode().Value.(int) < node.Value.(int) {
		return false
	}

	return isBST(node.GetLeftNode()) && isBST(node.GetRightNode())
}
