package main

type distanceInfo struct {
	distance   int
	lastVertex int
}

func (d distanceInfo) getDistance() int {
	return d.distance
}

func (d distanceInfo) getLastVertex() int {
	return d.lastVertex
}

func (d *distanceInfo) setDistance(distance int) {
	d.distance = distance
}

func (d *distanceInfo) setLastVertex(lastVertex int) {
	d.lastVertex = lastVertex
}
