package main

import (
	"algos/Tree/utils"
	"fmt"
)

func main() {
	fmt.Println(minBFS(utils.GetNumbericRootNode()))

	minValue := 9999999
	min(utils.GetNumbericRootNode(), &minValue)
	fmt.Println(minValue)
}

func minBFS(root *utils.Node) int {
	min := 9999999
	queue := utils.NewQueue()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		currNode := queue.Dequeue()

		if currNode != nil {
			if currNode.GetValue().(int) < min {
				min = currNode.GetValue().(int)
			}
			queue.Enqueue(currNode.GetLeftNode())
			queue.Enqueue(currNode.GetRightNode())
		}

	}

	return min
}

func min(root *utils.Node, i *int) {
	if root == nil {
		return
	}

	if root.Value.(int) < *i {
		 *i = root.Value.(int)
	}

	min(root.GetLeftNode(), i)
	min(root.GetRightNode(), i)
}
