package main

import (
	"algos/tree/utils"
	"fmt"
)

func main() {
	rootNode := utils.GetAlphaRootNode()
	fmt.Println("preorder")
	preorder(rootNode)
	fmt.Println()

	fmt.Println("inorder")
	inorder(rootNode)
	fmt.Println()

	fmt.Println("postOrder")
	postOrder(rootNode)
	fmt.Println()
}

func preorder(root *utils.Node) {
	if root == nil {
		return
	}

	fmt.Print(root.Value)

	preorder(root.GetLeftNode())
	preorder(root.GetRightNode())
}

func inorder(root *utils.Node) {
	if root == nil {
		return
	}

	inorder(root.GetLeftNode())
	fmt.Print(root.Value)
	inorder(root.GetRightNode())
}

func postOrder(root *utils.Node) {
	if root == nil {
		return
	}

	postOrder(root.GetLeftNode())
	postOrder(root.GetRightNode())
	fmt.Print(root.Value)
}
