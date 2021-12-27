package utils

import "fmt"

func GetAlphaRootNode() *Node {
	aNode := NewNode("A")
	bNode := NewNode("B")
	cNode := NewNode("C")
	dNode := NewNode("D")
	eNode := NewNode("E")
	fNode := NewNode("F")
	gNode := NewNode("G")
	hNode := NewNode("H")

	aNode.SetLeft(bNode)
	aNode.SetRight(cNode)
	bNode.SetLeft(dNode)
	bNode.SetRight(eNode)
	dNode.SetLeft(fNode)
	cNode.SetLeft(gNode)
	cNode.SetRight(hNode)
	return aNode
}

func GetNumbericRootNode() *Node {
	eight := NewNode(8)
	six := NewNode(6)
	fourteen := NewNode(14)
	four := NewNode(4)
	seven := NewNode(7)
	sixteen := NewNode(16)
	eighteen := NewNode(18)

	eight.SetLeft(six)
	eight.SetRight(fourteen)
	six.SetLeft(four)
	six.SetRight(seven)
	fourteen.SetRight(sixteen)
	sixteen.SetRight(eighteen)

	return eight
}

func PrintTree(root *Node) {
	queue := NewQueue()
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
