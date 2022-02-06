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
	preOrderWithoutRecursion(rootNode)
	fmt.Println()

	fmt.Println("inorder")
	inorder(rootNode)
	fmt.Println()
	inorderWithoutRecursion(rootNode)
	fmt.Println()

	fmt.Println("postOrder")
	postOrder(rootNode)
	fmt.Println()
	postOrderWithoutRecursion(rootNode)
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

func inorderWithoutRecursion(root *utils.Node) {
	stack := []*utils.Node{}
	if root == nil {
		return
	}

	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.LeftNode
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Print(curr.Value)
		curr = curr.GetRightNode()
	}
}

func preOrderWithoutRecursion(root *utils.Node) {
	if root == nil {
		return
	}

	curr := root
	stack := []*utils.Node{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			fmt.Print(curr.Value)
			stack = append(stack, curr)
			curr = curr.LeftNode
		}
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr = popped.RightNode
	}
}

func postOrderWithoutRecursion(root *utils.Node) {
	if root == nil {
		return
	}

	prev := root
	stack := []*utils.Node{root}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		if curr.LeftNode != nil && curr.LeftNode != prev && curr.RightNode != prev {
			stack = append(stack, curr.LeftNode)
		} else if curr.RightNode != nil && curr.RightNode != prev {
			stack = append(stack, curr.RightNode)
		} else {
			prev = curr
			fmt.Print(curr.Value)
			stack = stack[:len(stack)-1]
		}
	}
}
