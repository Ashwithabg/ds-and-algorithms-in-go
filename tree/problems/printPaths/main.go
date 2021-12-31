package main

import (
	"algos/tree/utils"
	"fmt"
)

/*
 Given a binary tree, print out all of its root-to-leaf
 paths, one per line. Uses a recursive helper to do the work.
*/

func main() {
	root := utils.GetNumbericRootNode()
	paths := [][]int{}
	path := []int{}
	printPath(root, &paths, path)
	fmt.Println(paths)
}

func printPath(node *utils.Node, paths *[][]int, path []int) {
	if node == nil {
		return
	}

	path = append(path, node.Value.(int))

	if node.IsLeafNode() {
		*paths = append(*paths, path)
		return
	}

	printPath(node.GetLeftNode(), paths, path)
	printPath(node.GetRightNode(), paths, path)
}
