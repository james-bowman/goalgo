package graph_test

import (
	"testing"

	"github.com/nfisher/goalgo/graph"
)

func Test_vertice(t *testing.T) {
	td := []struct {
		name  string
		list  *graph.AdjacencySet
		edges []int
		id    int
		err   error
	}{
		{"adds vertice to empty list", newList(), nil, 0, nil},
		{"adds vertice to populated list", newList(WithEdges()), nil, 1, nil},
		{"adds vertice with valid edge", newList(WithEdges()), []int{0}, 1, nil},
		{"rejects vertice with invalid edge", newList(), []int{1}, -1, graph.ErrCannotAddVertices},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.list.Vertice(tc.edges...)
			if id != tc.id {
				t.Errorf("list.Add(%v) id = %v, want %v", tc.edges, id, tc.id)
			}
			if err != tc.err {
				t.Errorf("list.Add(%v) id = %v, want %v", tc.edges, err, tc.err)
			}
		})
	}
}

func Test_counters(t *testing.T) {
	td := []struct {
		name     string
		actual   int
		expected int
	}{
		{"Vertices() for empty list", newList().Vertices(), 0},
		{"Vertices() for populated list", newList(WithEdges(), WithEdges(0)).Vertices(), 2},
		{"Edges() for empty list", newList().Edges(), 0},
		{"Edges() for populated list", newList(WithEdges()).Edges(), 0},
		{"Edges() for populated list with edge", newList(WithEdges(), WithEdges(0), WithEdges(0, 1)).Edges(), 3},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			if tc.actual != tc.expected {
				t.Errorf("len = %v, want %v", tc.actual, tc.expected)
			}
		})
	}
}

type Modifier func(*graph.AdjacencySet)

func WithEdges(i ...int) Modifier {
	return func(as *graph.AdjacencySet) {
		as.Vertice(i...)
	}
}

func newList(mm ...Modifier) *graph.AdjacencySet {
	s := &graph.AdjacencySet{}
	for _, m := range mm {
		m(s)
	}
	return s
}