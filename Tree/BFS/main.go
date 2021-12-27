package main

import (
	"algos/Tree/utils"
	"fmt"
)

func main() {
	rootNode := utils.GetRootNode()

	findBFS(rootNode)
}

func findBFS(root *utils.Node) {
	queue := utils.NewQueue()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		n := queue.Dequeue()
		fmt.Println(n.GetValue())
		if n.GetLeftNode() != nil {
			queue.Enqueue(n.GetLeftNode())
		}

		if n.GetRightNode() != nil {
			queue.Enqueue(n.GetRightNode())
		}
	}
}
