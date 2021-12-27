package utils

func GetRootNode() *Node {
	aNode := NewNode("A")
	bNode := NewNode("B")
	cNode := NewNode("C")
	dNode := NewNode("D")
	eNode := NewNode("E")
	fNode := NewNode("F")
	gNode := NewNode("G")
	rootNode := NewNode("root")

	rootNode.SetLeft(aNode)
	rootNode.SetRight(bNode)
	bNode.SetLeft(cNode)
	bNode.SetRight(dNode)
	dNode.SetLeft(eNode)
	cNode.SetLeft(fNode)
	cNode.SetRight(gNode)
	return rootNode
}
