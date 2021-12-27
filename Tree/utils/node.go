package utils

type Node struct {
	Value     interface{}
	LeftNode  *Node
	RightNode *Node
}

func NewNode(value interface{}) *Node{
	return &Node{Value: value}
}

func (n Node) GetLeftNode() *Node {
	return n.LeftNode
}

func (n Node) GetRightNode() *Node {
	return n.RightNode
}

func (n *Node) SetLeft(node *Node) {
	n.LeftNode = node
}

func (n *Node) SetRight(node *Node) {
	n.RightNode = node
}

func (n *Node) GetValue() interface{} {
	return n.Value
}
