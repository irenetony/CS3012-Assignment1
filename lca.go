package lca

var path1 []int
var path2 []int

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
	// path1 = make([]int, x)
	// path2 = make([]int, x)

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
func (g *Graph) DFS(root int, find int) int {
	pathT := -1
	//if the root is the node to be found, return 0
	if root == find {
		path1 = append(path1, root) //add root to path
		return 0
	}
	// //if the root has an edge with the node to be found, then search it and expect a 0 in return
	// if g.Edges[root][find] == 1 {
	// 	g.Visited[root] = 1
	// 	path1 = append(path1, root) //add root to path
	// 	return g.DFS(find, find)
	// }
	//if the root doesnt have an edge with the node to be found, then search the adj nodes.
	//If one search retuns a 0, then add that node to the path
	for j := 0; j < g.NumNodes; j++ {
		if g.Edges[root][j] == 1 && g.Visited[j] == 0 {
			g.Visited[root] = 1
			pathT = g.DFS(j, find)
			if pathT == 0 {
				path1 = append(path1, root) //add j to the path
				return pathT
			}
		}
	}
	return pathT
}

//LCA takes in two nodes and returns their lowest common ancestor
func (g *Graph) LCA(x int, y int) int {
	lca := -1
	if x == y {
		return x
	}
	//find path to x
	g.DFS(0, x)
	pathA := make([]int, len(path1))
	copy(pathA, path1)

	//clear the global array
	path1 = path1[:0]
	g.Visited = make([]int, g.NumNodes)
	//find path to y
	g.DFS(0, y)
	pathB := make([]int, len(path1))
	copy(pathB, path1)

	for i := 0; i < len(pathA); i++ {
		for j := 0; j < len(pathB); j++ {
			if pathA[i] == pathB[j] {
				return pathA[i]
			}
		}

	}
	path1 = path1[:0]
	return lca
}
