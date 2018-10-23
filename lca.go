package lca

//Garph: Construct a Graph
type Graph struct {
	NumNodes int
	Edges    [][]int
	Visited  []int
	Min      int
}

//Construct an edge for a graph
type Edge struct {
	From   int
	To     int
	Weight int
}

//EmptyGraph: returns an empty graph
func EmptyGraph() *Graph {
	return &Graph{}
}
func CreateGraph(x int) *Graph {
	edges := make([][]int, x)
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			edges[i][j] = -1
		}
	}

	graph := &Graph{
		NumNodes: x,
		Edges:    edges,
		Visited:  make([]int, x),
	}
	return graph
}
func (g *Graph) AddEdge(x int, y int) {

}
func (g *Graph) RemoveEdge(x int, y int) {

}
