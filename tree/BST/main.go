package main

import (
	"algos/tree/utils"
	"fmt"
)

func main() {
	root := utils.GetNumbericRootNode()

	two := utils.NewNode(2)
	insert(root, two)
	utils.PrintTree(root)

	fmt.Println(lookUp(root, 12))
}

func insert(root *utils.Node, two *utils.Node) *utils.Node {
	if root == nil {
		return two
	}

	if two.GetValue().(int) <= root.GetValue().(int) {
		if root.LeftNode == nil {
			root.SetLeft(two)
		} else {
			insert(root.GetLeftNode(), two)
		}
	} else {
		if root.RightNode == nil {
			root.SetRight(two)
		} else {
			insert(root.GetRightNode(), two)
		}
	}

	return two
}

func lookUp(root *utils.Node, value int) *utils.Node {
	if root == nil || root.Value == value {
		return root
	}
	if value <= root.Value.(int) {
		return lookUp(root.GetLeftNode(), value)
	}

	return lookUp(root.GetRightNode(), value)
}
