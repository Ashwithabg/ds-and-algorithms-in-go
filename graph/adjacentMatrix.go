package graph

import (
	"fmt"
	"sort"
)

type adjacentMatrixGraph struct {
	matrix        [][]int
	graphType     Type
	verticesCount int
}

func (a adjacentMatrixGraph) GetNumberOfVertices() int {
	return a.verticesCount
}

func (a adjacentMatrixGraph) AddEdge(v1, v2 int) error {
	if v1 >= a.verticesCount || v1 < 0 || v2 >= a.verticesCount || v2 < 0 {
		return fmt.Errorf("not valid vertices")
	}

	a.matrix[v1][v2] = 1
	if a.graphType == UNDIRECTED {
		a.matrix[v2][v1] = 1
	}

	return nil
}

func (a adjacentMatrixGraph) GetAdjacentVertices(vertex int) ([]int, error) {
	vertices := []int{}
	if vertex >= a.verticesCount || vertex < 0 {
		return vertices, fmt.Errorf("invalid Vertex")
	}

	for i := 0; i < a.verticesCount; i++ {
		if a.matrix[vertex][i] == 1 {
			vertices = append(vertices, i)
		}
	}
	sort.Ints(vertices)
	return vertices, nil
}

func NewAdjacentMatrixGraph(graphType Type, verticesCount int) Graph {
	matrix := make([][]int, verticesCount, verticesCount)
	for i := 0; i < verticesCount; i++ {
		for j := 0; j < verticesCount; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}

	return &adjacentMatrixGraph{
		matrix:        matrix,
		graphType:     graphType,
		verticesCount: verticesCount,
	}
}

func (a adjacentMatrixGraph) GetIndegree(v int) (int, error) {
	indegree := 0
	if v < 0 || v >= len(a.matrix) {
		return indegree, fmt.Errorf("invalid vertex")
	}

	for i := 0; i < a.verticesCount; i++ {
		if a.matrix[i][v] == 1 {
			indegree++
		}
	}

	return indegree, nil
}
