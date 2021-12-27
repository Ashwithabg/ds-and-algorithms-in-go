package utils

func GetRootNode() *Node {
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
