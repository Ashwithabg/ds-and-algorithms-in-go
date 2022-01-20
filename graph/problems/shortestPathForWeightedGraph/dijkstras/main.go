package main

import "fmt"

const verylargeNumber = 999999

func main() {
	a := 0
	b := 1
	c := 2
	d := 3
	e := 4
	g := &graph{
		edges: map[int][]vertex{
			a: {newVertex(b, 2), newVertex(c, 3)},
			b: {newVertex(d, 2)},
			c: {newVertex(e, 6)},
			d: {},
			e: {newVertex(d, 4)},
		},
	}
	fmt.Printf("%+v\n", createDistanceTable(g, a))
	findShortestDistance(g, a, d)
}

func findShortestDistance(g *graph, source int, dest int) {
	distanceTable := createDistanceTable(g, source)
	stack := newStack()
	stack.push(dest)

	distanceInfo := distanceTable[dest]
	for distanceInfo.getDistance() != verylargeNumber && distanceInfo.previousVertex != source {
		stack.push(distanceInfo.previousVertex)
		distanceInfo = distanceTable[distanceInfo.getPreviousVertex()]
	}

	if distanceInfo.previousVertex == -1 {
		fmt.Errorf("there is no path")
	} else {
		fmt.Printf("shortest distance is %v", source)
		for !stack.isEmpty() {
			fmt.Printf(" ==> %+v", stack.pop())
		}
	}

}

func createDistanceTable(g *graph, source int) distanceTable {
	distanceTable := initDistanceTable(g, source)
	queue := newPriorityQueue()
	queue.Push(newWeightedVector(source, 0))

	for queue.Len() != 0 {
		currNode := queue.Pop().(*weightedVector)
		currDistance := currNode.weight
		currVectorId := currNode.vertexId
		adjacentNodes := g.getAdjacentWeightedvectors(currVectorId)

		for _, adjacentNode := range adjacentNodes {
			adjacentVectorId := adjacentNode.vertexId
			currAdjacentNode := distanceTable[adjacentVectorId]
			newDistance := adjacentNode.distance + currDistance

			if currAdjacentNode.getDistance() > newDistance {
				distanceTable.update(adjacentVectorId, newDistance, currVectorId)
			}

			if g.hasAdjacentVectors(adjacentVectorId) {
				queue.Push(newWeightedVector(adjacentVectorId, newDistance))
			}
		}

	}

	return distanceTable
}

func initDistanceTable(g *graph, source int) distanceTable {
	table := distanceTable{}
	for i := 0; i < g.getNodeCount(); i++ {
		table[i] = distanceInfo{
			distance:       verylargeNumber,
			previousVertex: -1,
		}
	}

	sourceDistance := table[source]
	sourceDistance.setDistance(0)
	sourceDistance.setPreviousVertex(source)
	table[source] = sourceDistance

	return table
}
