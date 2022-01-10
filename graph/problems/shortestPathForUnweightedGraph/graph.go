package main

import "sync"

type graph struct {
	nodes []int
	edges map[int][]int
	lock  sync.RWMutex
}

func (g *graph) addEdge(node1 int, node2 int) {
	g.lock.Lock()
	g.edges[node1] = append(g.edges[node1], node2)
	g.lock.Unlock()
}

func (g *graph) getAdjacentNodes(node int) []int {
	g.lock.Lock()
	strings := g.edges[node]
	g.lock.Unlock()
	return strings
}

func (g *graph) getNumberOfNodes() int {
	return len(g.nodes)
}
