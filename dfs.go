package graph

// DFS DepthFirstSearch is a classic recursive method for systematically
// examining each of the vertices and edges in a graph.
type DFS struct {
	// marked[v] = is there an s-v path?
	marked []bool
	// number of vertices connected to s
	n uint
}

func NewDFS(g *Graph, v uint) *DFS {
	if err := g.validateVertex(v); err != nil {
		panic(err)
	}
	return &DFS{
		marked: make([]bool, g.V()),
	}
}

// depth first search from v
// func (d *DFS) dfs(g *Graph, v uint) {
// 	d.n++
// 	d.marked[v] = true
//   i := g.Adj(v).Iterator()
//   for i.HasNext(){
//     if d.marked[]
//   }
// }
