// Package graph based on http://algs4.cs.princeton.edu/41graph/
package graph

import (
	"fmt"
	"os"
	"testing"
)

func TestGraph(t *testing.T) {
	g := New(5)
	if g.V() != 5 {
		t.Fatalf("graph should have 5 vertexes got %v", g.V())
	}
	if g.E() != 0 {
		t.Fatalf("graph should have 0 edges got %v", g.E())
	}
}

func TestAddEdge(t *testing.T) {
	g := New(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 4)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(2, 3)
	g.AddEdge(2, 1)
	if g.V() != 5 {
		t.Fatalf("graph should have 5 vertexes got %v", g.V())
	}
	if g.E() != 6 {
		t.Fatalf("graph should have 6 edges got %v", g.E())
	}
	b, _ := g.Adj(0)
	if b.Size() != 2 {
		t.Fatalf("bag should have 2 objects got %v", b.Size())
	}
	d, _ := g.Degree(0)
	if d != 2 {
		t.Fatalf("vertex degree should be 2 got %v", d)
	}
	if err := g.AddEdge(5, 1); err == nil {
		t.Fatalf("vertex 5 should be invalid")
	}
	if err := g.AddEdge(3, 6); err == nil {
		t.Fatalf("vertex 6 should be invalid")
	}
	fmt.Print(g.String())
}

func TestLoadTiny(t *testing.T) {
	f, err := os.Open("tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	g, err := Load(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(g.String())
}

func TestLoadMedium(t *testing.T) {
	f, err := os.Open("mediumG.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	g, err := Load(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(g.String())
}
