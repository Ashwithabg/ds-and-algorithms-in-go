package graph

import "container/list"

type adjacentListNode struct {
	vertexId int
	nodes    *list.List
}

type adjacentListGraph struct {
	vertices *list.List
}


