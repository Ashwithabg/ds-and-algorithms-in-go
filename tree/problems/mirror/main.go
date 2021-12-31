package main

import (
	"algos/tree/utils"
	"fmt"
)

/*
 Change a tree so that the roles of the
 left and right pointers are swapped at every node.
 So the tree...
 4
 / \
 2 5
 / \
 1 3
 is changed to...
 4
 / \
 5 2
 / \
 3 1
*/

func main() {
	root := utils.GetNumbericRootNode()
	printTree(root)
	mirror(root)
	fmt.Println()
	printTree(root)
}

func mirror(node *utils.Node) {
	if node == nil {
		return
	}

	mirror(node.GetLeftNode())
	mirror(node.GetRightNode())

	var temp *utils.Node
	temp = node.GetLeftNode()
	node.SetLeft(node.GetRightNode())
	node.SetRight(temp)
}

func printTree(node *utils.Node) {
	if node == nil {
		return
	}

	fmt.Print(node.Value.(int), ",")
	printTree(node.GetLeftNode())
	printTree(node.GetRightNode())
}
