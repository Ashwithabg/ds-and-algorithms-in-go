package main

import "sync"

type graph struct {
	edges map[int][]vertex
	mutex sync.RWMutex
}

func (g *graph) addEdge(vertex vertex, vertex2 vertex) {
	g.mutex.Lock()
	g.edges[vertex.vertexId] = append(g.edges[vertex.vertexId], vertex2)
	g.mutex.Unlock()
}

func (g *graph) getAdjacentWeightedvectors(vertexId int) []vertex {
	return g.edges[vertexId]
}

func (g *graph) getNodeCount() int {
	return len(g.edges)
}

func (g *graph) hasAdjacentVectors(vertexId int) bool {
	return len(g.edges[vertexId]) != 0
}
