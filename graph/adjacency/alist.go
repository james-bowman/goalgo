package adjacency

import "github.com/nfisher/goalgo/graph/errors"

// List is an adjacency list using an array.
type List struct {
	list  [][]int
	edges int
}

// Vertex adds a new vertex, optionally with the specified edges.
func (as *List) Vertex(edges ...int) (id int, err error) {
	var edgeset []int
	l := len(as.list)
	for _, i := range edges {
		if i >= l {
			return -1, errors.ErrCannotAddVertices
		}
		edgeset = append(edgeset, i)
	}

	as.edges += len(edgeset)
	as.list = append(as.list, edgeset)

	return l, nil
}

// Edge adds an edge from v to w.
func (as *List) Edge(v, w int) error {
	l := len(as.list)
	if v >= l {
		return errors.ErrCannotAddEdge
	}

	if w >= l {
		return errors.ErrCannotAddEdge
	}

	as.list[v] = append(as.list[v], w)
	as.edges++

	return nil
}

// Adjacent returns all vertices adjacent to this vertex.
func (as *List) Adjacent(v int) ([]int, error) {
	if v >= len(as.list) {
		return nil, errors.ErrVertexNotFound
	}

	var a []int
	for k := range as.list[v] {
		a = append(a, k)
	}

	return a, nil
}

// Vertices returns the number of vertices in the list.
func (as *List) Vertices() int {
	return len(as.list)
}

// Edges returns the number edges in the list.
func (as *List) Edges() int {
	return as.edges
}
