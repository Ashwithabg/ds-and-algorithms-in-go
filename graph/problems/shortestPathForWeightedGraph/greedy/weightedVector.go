package main

type weightedVector struct {
	vertexId int
	index    int
	weight   int
}

func newWeightedVector(vertexId int, weight int) *weightedVector {
	return &weightedVector{
		vertexId: vertexId,
		weight:   weight,
	}
}
