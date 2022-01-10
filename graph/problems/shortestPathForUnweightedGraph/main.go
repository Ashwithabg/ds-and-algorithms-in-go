package main

import "fmt"

type DistanceTable map[int]distanceInfo

func main() {
	a := 0
	b := 1
	c := 2
	d := 3
	e := 4
	g := &graph{
		nodes: []int{a, b, c, d, e},
		edges: map[int][]int{
			a: {b, c},
			b: {d},
			c: {e},
			d: {},
			e: {d},
		},
	}

	findShortestDistance(g, a, d)
}

func findShortestDistance(graph *graph, source int, dest int) {
	distanceTable := buildDistanceTable(graph, source)
	stack := newStack()
	stack.push(dest)

	previousVertex := distanceTable[dest].getLastVertex()
	if previousVertex != -1 && previousVertex != source {
		stack.push(previousVertex)
		previousVertex = distanceTable[previousVertex].getLastVertex()
	}

	if previousVertex == -1 {
		fmt.Errorf("there is no path")
	} else {
		fmt.Printf("smallest path %+v ", source)
		for !stack.isEmpty() {
			fmt.Printf("-> %+v ", stack.pop())
		}
	}
}

func buildDistanceTable(graph *graph, source int) DistanceTable {
	distanceTable := initDistanceTable(graph, source)

	q := newQueue()
	q.PushBack(source)

	for !q.isEmpty() {
		vertex := q.dequeue()
		vertexDistance := distanceTable[vertex].getDistance()
		adjacentVertex := graph.getAdjacentNodes(vertex)
		for _, adjVertex := range adjacentVertex {
			distanceInfo := distanceTable[adjVertex]
			if distanceInfo.getDistance() == -1 {
				distanceInfo.setDistance(vertexDistance + 1)
				distanceInfo.setLastVertex(vertex)
				distanceTable[adjVertex] = distanceInfo

				if !(len(graph.getAdjacentNodes(adjVertex)) == 0) {
					q.enqueue(adjVertex)
				}
			}

		}
	}

	return distanceTable
}

func initDistanceTable(graph *graph, source int) DistanceTable {
	distanceTable := make(DistanceTable)
	for i := 0; i < graph.getNumberOfNodes(); i++ {
		distanceTable[i] = distanceInfo{
			distance:   -1,
			lastVertex: -1,
		}
	}

	distanceTable[source] = distanceInfo{
		distance:   0,
		lastVertex: source,
	}
	return distanceTable
}
