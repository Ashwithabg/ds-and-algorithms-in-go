package main

import "algos/Tree/utils"

/*
 For each node in a binary search tree,
 create a new duplicate node, and insert
 the duplicate as the left child of the original node.
 The resulting tree should still be a binary search tree.
 So the tree...
 2
 / \
 1 3
 Is changed to...
 2
 / \
 2 3
 / /
 1 3
 /
 1
*/
func main() {
	root := utils.GetNumbericRootNode()
	printTree(root)
	doubleTree(root)
	println()
	printTree(root)

}

func doubleTree(node *utils.Node) {
	if node == nil {
		return
	}

	doubleTree(node.GetLeftNode())
	doubleTree(node.GetRightNode())

	var oldLeft *utils.Node
	oldLeft = node.GetLeftNode()
	node.SetLeft(utils.NewNode(node.Value))
	node.GetLeftNode().SetLeft(oldLeft)
}

func printTree(node *utils.Node) {
	if node == nil {
		return
	}

	print(node.Value.(int), ",")
	printTree(node.GetLeftNode())
	printTree(node.GetRightNode())
}
