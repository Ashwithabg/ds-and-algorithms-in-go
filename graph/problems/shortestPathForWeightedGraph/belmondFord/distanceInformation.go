package main

const (
	MaxDistance = 999999999
	lastvertex  = -1
)

type distanceInformation struct {
	distance   int
	lastVertex int
}

func newDistanceInformation() distanceInformation {
	return distanceInformation{
		distance: MaxDistance,
		lastVertex: lastvertex,
	}
}

func (d distanceInformation) getDistance() int {
	return d.distance
}

func (d distanceInformation) getLastVertex() int {
	return d.lastVertex
}

func (d *distanceInformation) setDistance(distance int) {
	d.distance = distance
}

func (d *distanceInformation) setLastVertex(lastVertex int) {
	d.lastVertex = lastVertex
}
