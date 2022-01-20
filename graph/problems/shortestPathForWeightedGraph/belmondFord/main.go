package main

import (
	"fmt"
)

func main() {

}

func buildDistanceTable(g graph, source int) {
	distanceTable := map[int]distanceInformation{}
	for i := 0; i < g.getNodeCount(); i++ {
		distanceTable[i] = newDistanceInformation()
	}

	information := distanceTable[source]
	information.setDistance(0)
	information.setLastVertex(source)

	q := queue{}
	q.enqueue(source)

	for numInteration := 0; numInteration < g.getNodeCount()-1; numInteration++ {
		for v := 0; v < g.getNodeCount(); v++ {
			q.enqueue(v)
		}

		visitedEdges := []string{}

		for !q.isEmpty() {
			currentVertex := q.dequeue()
			for _, neighbour := range g.getAdjacentVertices(currentVertex) {
				edge := fmt.Sprintf("%s-%s", currentVertex, neighbour)
				if contains(visitedEdges, edge) {
					continue
				}
			}
		}

	}

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
