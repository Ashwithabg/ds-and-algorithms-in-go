package main

import (
	"algos/Tree/utils"
	"fmt"
)

func main() {
	rootNode := constuctTree()

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

func constuctTree() *utils.Node {
	aNode := utils.NewNode("A")
	bNode := utils.NewNode("B")
	cNode := utils.NewNode("C")
	dNode := utils.NewNode("D")
	eNode := utils.NewNode("E")
	rootNode := utils.NewNode("root")

	rootNode.SetLeft(aNode)
	rootNode.SetRight(bNode)
	bNode.SetLeft(cNode)
	bNode.SetRight(dNode)
	dNode.SetLeft(eNode)
	return rootNode
}
