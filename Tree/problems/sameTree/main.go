package main

import (
	"algos/Tree/utils"
	"fmt"
)

func main() {
	root := utils.GetNumbericRootNode()
	root2 := utils.NewNode(1)
	root2.SetLeft(utils.NewNode(2))
	root2.SetRight(utils.NewNode(3))
	fmt.Println(sameTree(root, root2))
}

func sameTree(node1 *utils.Node, node2 *utils.Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1.Value.(int) == node2.Value.(int) {
		return sameTree(node1.GetLeftNode(), node2.GetLeftNode()) &&
			sameTree(node1.GetRightNode(), node2.GetRightNode())
	}

	return false
}
