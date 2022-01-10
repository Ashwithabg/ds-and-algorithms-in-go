package main

type distanceInfo struct {
	distance       int
	previousVertex int
}

func (d distanceInfo) getDistance() int {
	return d.distance
}

func (d distanceInfo) getPreviousVertex() int {
	return d.previousVertex
}

func (d *distanceInfo) setDistance(distance int) {
	d.distance = distance
}

func (d *distanceInfo) setPreviousVertex(v int) {
	d.previousVertex = v
}
