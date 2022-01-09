package main

import (
	"algos/Tree/utils"
	graph "algos/graph"
	"fmt"
)

func main() {
	graph := graph.NewAdjacentMatrixGraph(graph.DIRECTED, 5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(4, 3)
	graph.AddEdge(4, 1)
	fmt.Printf("%+v\n", graph)
	fmt.Printf("%+v\n", sort(graph))
}

func sort(g graph.Graph) []int {
	indegreeMap := make(map[int]int)
	queue := utils.NewIntegerQueue()
	var sortedVertices []int

	for i := 0; i < g.GetNumberOfVertices(); i++ {
		indegree, _ := g.GetIndegree(i)
		indegreeMap[i] = indegree
		if indegree == 0 {
			queue.Enqueue(i)
		}
	}

	for !queue.IsEmpty() {
		element := queue.Dequeue()
		sortedVertices = append(sortedVertices, element)
		adjVertices, _ := g.GetAdjacentVertices(element)
		for _, a := range adjVertices {
			indegreeMap[a]--
			if indegreeMap[a] == 0 {
				queue.Enqueue(a)
			}
		}
	}

	if len(sortedVertices) != g.GetNumberOfVertices() {
		fmt.Errorf("Has cycles")
		return []int{}
	}

	return sortedVertices
}
