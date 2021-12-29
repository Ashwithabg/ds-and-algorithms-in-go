package main

import (
	"algos/Tree/utils"
	"fmt"
)

/*
 Compute the number of nodes in a tree.
*/

func main() {
	root := utils.GetNumbericRootNode()
	counter := 0
	traverse(root, &counter)
	fmt.Println(counter)
	fmt.Println(traverseBFS(root))
}

func traverse(root *utils.Node, counter *int) {
	if root == nil {
		return
	}
	*counter++

	traverse(root.GetRightNode(), counter)
	traverse(root.GetLeftNode(), counter)
}

func traverseBFS(root *utils.Node) int {
	counter := 0
	queue := utils.NewQueue()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		node := queue.Dequeue()
		if node != nil {
			counter++
			queue.Enqueue(node.GetLeftNode())
			queue.Enqueue(node.GetRightNode())
		}
	}

	return counter
}
