package graph

import (
	"fmt"
)

type AdjacentSetGraph struct {
	nodes     []*VertexNode
	graphType Type
}

func (a *AdjacentSetGraph) AddEdge(v1, v2 int) error {
	if v1 >= len(a.nodes) || v1 < 0 || v2 >= len(a.nodes) || v2 < 0 {
		return fmt.Errorf("not valid vertices")
	}

	a.nodes[v1].AddEdge(v2)
	return nil
}

func (a *AdjacentSetGraph) GetAdjacentVertices(vertice int) ([]int, error) {
	panic("implement me")
}

func NewAdjacentSetGraph(vertexCount int, grphType Type) Graph {
	nodes := []*VertexNode{}
	for i := 0; i < vertexCount; i++ {
		nodes = append(nodes, NewNode(i))
	}

	return &AdjacentSetGraph{
		nodes:     nodes,
		graphType: grphType,
	}
}
