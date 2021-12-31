package main


import (
	"algos/tree/utils"
	"fmt"
	"math"
)

func main() {
	invalidBST := utils.NewNode(3)
	invalidBST.SetLeft(utils.NewNode(10))
	invalidBST.SetLeft(utils.NewNode(4))

	root := utils.GetNumbericRootNode()
	fmt.Println(isBST(root))
	fmt.Println(isBST(invalidBST))

	fmt.Println(isBSTEfficient(invalidBST, math.MinInt, math.MaxInt))
	fmt.Println(isBSTEfficient(root, math.MinInt, math.MaxInt))
}

func isBST(node *utils.Node) bool {
	if node == nil {
		return true
	}

	if node.GetLeftNode() != nil && node.GetLeftNode().Value.(int) > node.Value.(int) {
		return false
	}

	if node.GetRightNode() != nil && node.GetRightNode().Value.(int) < node.Value.(int) {
		return false
	}

	return isBST(node.GetLeftNode()) && isBST(node.GetRightNode())
}

func isBSTEfficient(node *utils.Node, min int, max int) bool {
	if node == nil {
		return true
	}

	data := node.Value.(int)
	if data < min || data > max {
		return false
	}

	return isBSTEfficient(node.GetLeftNode(), min, data) &&
		isBSTEfficient(node.GetRightNode(), data, max)
}
