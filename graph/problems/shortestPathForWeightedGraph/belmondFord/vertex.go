package main

type vertex struct {
	vertexId int
	distance int
}

func newVertex(vertexId int, distance int) vertex {
	return vertex{
		vertexId: vertexId,
		distance: distance,
	}
}
