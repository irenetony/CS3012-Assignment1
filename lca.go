package lca

import "fmt"

//Graph Construct a Graph
type Graph struct {
	NumNodes int
	Edges    [][]int
	Visited  []int
	//Min      int
}

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

//DFS takes in the root node and the node to be found. It returns the path to the node.
func (g *Graph) DFS(root int, find int, path int) int {
	pathT := 0
	if root == find {
		return path
	}
	for j := 0; j < g.NumNodes; j++ {
		if g.Edges[root][j] == 1 && g.Visited[j] == 0 {
			g.Visited[j] = 1
			pathT = g.DFS(j, find, path)
			fmt.Printf("%v -> %v %v\n", root, j, pathT)
		}
	}
	return pathT
}

//LCA takes in two nodes and returns their lowest common ancestor
func (g *Graph) LCA(x int, y int) int {
	lca := 0
	return lca
}
