package graph

type Type int

const (
	DIRECTED Type = 1
	UNDIRECTED
)

type Graph interface {
	AddEdge(v1, v2 int) error
	GetAdjacentVertices(vertice int) ([]int, error)
}
