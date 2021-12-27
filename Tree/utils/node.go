package utils

type Node struct {
	Value     string
	LeftNode  *Node
	RightNode *Node
}

func NewNode(value string) *Node{
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

func (n *Node) GetValue() string {
	return n.Value
}
