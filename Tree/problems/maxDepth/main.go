package main

import (
	"algos/Tree/utils"
	"fmt"
)

/*
 Compute the "maxDepth" of a tree -- the number of nodes along
 the longest path from the root node down to the farthest leaf node.
*/

func main() {
	root := utils.GetNumbericRootNode()
	fmt.Println(maxDepth(root))
	fmt.Println(maxDepthBFS(root))
}

func maxDepthBFS(root *utils.Node) int {
	depth := 0
	queue := utils.NewQueue()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		depth++

		size := queue.Size()
		for ;size > 0; size-- {
			node := queue.Dequeue()
			if node.GetLeftNode() != nil {
				queue.Enqueue(node.GetLeftNode())
			}

			if node.GetRightNode() != nil {
				queue.Enqueue(node.GetRightNode())
			}
		}

	}

	return depth
}

func maxDepth(root *utils.Node) int {
	if root == nil {
		return 0
	}

	lDepth := maxDepth(root.GetLeftNode())
	rDepth := maxDepth(root.GetRightNode())

	if lDepth > rDepth {
		return lDepth + 1
	}

	return rDepth + 1
}
