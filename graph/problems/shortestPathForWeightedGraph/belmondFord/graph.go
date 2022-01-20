package main

type graph struct {
	edges map[int][]vertex
}

func (g *graph) getNodeCount() int {
	return len(g.edges)
}

func (g *graph) getAdjacentVertices(vertex int) []vertex {
	return g.edges[vertex]
}
