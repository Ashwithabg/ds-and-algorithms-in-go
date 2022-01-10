package main

type distanceTable map[int]distanceInfo

func (dt distanceTable) update(key int, distance int, previousVector int) {
	distanceInfo := dt[key]
	distanceInfo.setDistance(distance)
	distanceInfo.setPreviousVertex(previousVector)
	dt[key] = distanceInfo
}
