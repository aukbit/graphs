// Package graph based on http://algs4.cs.princeton.edu/41graph/
package graph

import (
	"bytes"
	"fmt"

	"github.com/aukbit/cache/bag"
)

type Graph struct {
	// number of vertices
	v uint
	// number of edges
	e uint
	// adj slice of bags, each contains vertices adjancents to a given vertex
	adj []*bag.Bag
}

// New create a V-vertex graph with no edges
func New(vertex uint) *Graph {
	return &Graph{
		v: vertex,
		adj: func() (out []*bag.Bag) {
			var i uint
			for i = 0; i < vertex; i++ {
				out[i] = bag.New()
			}
			return out
		}(),
	}
}

// V number of vertices
func (g *Graph) V() uint {
	return g.v
}

// E number of edges
func (g *Graph) E() uint {
	return g.e
}

func (g *Graph) validateVertex(v uint) error {
	if v >= g.v {
		return fmt.Errorf("invalid vertex %v >= %v", v, g.v)
	}
	return nil
}

// addEdge add the undirected edge v-w to this graph
func (g *Graph) addEdge(v, w uint) error {
	if err := g.validateVertex(v); err != nil {
		return err
	}
	if err := g.validateVertex(w); err != nil {
		return err
	}
	g.e++
	g.adj[v].Add(w)
	g.adj[w].Add(v)
	return nil
}

// Adj vertices adjancent to v
func (g *Graph) Adj(v uint) (*bag.Bag, error) {
	if err := g.validateVertex(v); err != nil {
		return nil, err
	}
	return g.adj[v], nil
}

// Degree returns the degree of vertex
func (g *Graph) Degree(v uint) (int, error) {
	if err := g.validateVertex(v); err != nil {
		return 0, err
	}
	return g.adj[v].Size(), nil
}

// String representation
func (g *Graph) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("vertices: %v, edges: %v\n", g.v, g.e))
	var i uint
	for i = 0; i < g.v; i++ {
		buffer.WriteString(fmt.Sprintf("%v: ", i))
		iter := g.adj[i].Iterator()
		for iter.HasNext() {
			n, _ := iter.Next()
			buffer.WriteString(fmt.Sprintf("%v ", n))
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
