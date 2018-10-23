package lca

//Graph Construct a Graph
type Graph struct {
	NumNodes int
	Edges    [][]int
	Visited  []int
	Min      int
}

//Edge Construct an edge for a graph
// type Edge struct {
// 	From   int
// 	To     int
// 	Weight int
// }

//CreateGraph returns an empty Digraph
func CreateGraph(x int) *Graph {
	edges := make([][]int, x)
	for i := range edges {
		edges[i] = make([]int, x)
	}
	graph := &Graph{
		NumNodes: x,
		Edges:    edges,
		Visited:  make([]int, x),
	}
	return graph
}

//AddEdge adds an edge to theDigraph
func (g *Graph) AddEdge(x int, y int) {
	g.Edges[x][y] = 1
}

//RemoveEdge adds an edge to theDigraph
func (g *Graph) RemoveEdge(x int, y int) {
	g.Edges[x][y] = 0
}

//ValidNode finds a node and returns true if found
func (g *Graph) ValidNode(x int, y int) bool {
	isfound := false
	if x < g.NumNodes || x > 0 {
		isfound = true
	}
	if y < g.NumNodes || y > 0 {
		isfound = true
	}
	return isfound
}

//ValidEdge finds a edge and returns true if found
func (g *Graph) ValidEdge(x int, y int) bool {
	isfound := false
	if g.ValidNode(x, y) {
		if g.Edges[x][y] == 1 {
			isfound = true
		}
	}
	return isfound
}
