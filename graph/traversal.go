package graph

import "algos/Tree/utils"

func bfs(graph Graph, visited []int, currentVertex int) {
	queue := utils.NewIntegerQueue()
	queue.Enqueue(currentVertex)

	for !queue.IsEmpty() {
		value := queue.Dequeue()

		if visited[value] == 1 {
			return
		}

		visited = append(visited, currentVertex)
		vertices, _ := graph.GetAdjacentVertices(value)

		for vertex := range vertices {
			queue.Enqueue(vertex)
		}
	}
}

func dfs(graph Graph, visited []int, currentVertex int) {
	if visited[currentVertex] == 1 {
		return
	}

	vertices, _ := graph.GetAdjacentVertices(currentVertex)
	visited = append(visited, currentVertex)
	for vertex := range vertices {
		dfs(graph, visited, vertex)
	}
}
