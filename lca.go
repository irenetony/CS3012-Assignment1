package lca

import "errors"

//Graph Construct a Graph
type Graph struct {
	NumNodes int
	Edges    [][]int
	Visited  []int
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
func (g *Graph) DFS(root int, find int, path []int) ([]int, error) {
	//pathT := make([]int, g.NumNodes)

	//if the root is the node to be found, return 0
	if root == find {
		path = append(path, root) //add root to path
		return path, nil
	}
	//if the root doesnt have an edge with the node to be found, then search the adj nodes.
	//If one search retuns a 0, then add that node to the path
	for j := 0; j < g.NumNodes; j++ {
		if g.Edges[root][j] == 1 && g.Visited[j] == 0 {
			g.Visited[root] = 1
			pathT, err := g.DFS(j, find, path)
			if err == nil {
				if pathT[0] == find {
					path = pathT
					path = append(path, root) //add j to the path
					return path, nil
				}
			}
		}
	}
	return path, errors.New("path not found")
}

//LCA takes in two nodes and returns their lowest common ancestor
func (g *Graph) LCA(x int, y int) int {
	lca := 0
	var path1 []int
	var path2 []int

	if g.ValidNode(x, y) {
		if x == y {
			return x
		}
		//find path to x

		pathA, err := g.DFS(0, x, path1)
		if err != nil {
			lca = -1
		}
		//clear the visited array

		g.Visited = make([]int, g.NumNodes)
		//find path to y
		pathB, err := g.DFS(0, y, path2)
		if err != nil {
			lca = -1
		}
		if lca != -1 {
			for i := 0; i < len(pathA); i++ {
				for j := 0; j < len(pathB); j++ {
					if pathA[i] == pathB[j] {
						return pathA[i]
					}
				}

			}
		}

	}
	return lca
}
