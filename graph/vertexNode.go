package graph

import (
	"algos/utils"
	"sort"
)

type VertexNode struct {
	Id            int
	adjacentNodes utils.Set
}

func NewNode(vertex int) *VertexNode {
	return &VertexNode{Id: vertex}
}

func (n *VertexNode) AddEdge(vertex int) {
	n.adjacentNodes.Add(vertex)
}

func (n *VertexNode) GetAdjacentVertices() []int {
	nodes := n.adjacentNodes.Values()
	sort.Ints(nodes)
	return nodes
}
