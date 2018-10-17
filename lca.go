package lca

import "errors"

//Contruct a Node
type Node struct {
	Left  *Node
	Right *Node
	Key   int
	Value int
}

//Garph: Construct a Graph
type Graph struct {
	NumNodes int
	Edges    [][]Edge
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

func (n *Node) Insert(key int, data int) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	//If the data is already in the tree, return.
	case key == n.Key:
		return nil
		//If the data value is less than the current node’s value, and if the left child node is nil, insert a new left child node. Else call Insert on the left subtree.
	case key < n.Key:
		if n.Left == nil {
			n.Left = &Node{Key: key, Value: data}
			return nil
		}
		return n.Left.Insert(key, data)
		//If the data value is greater than the current node’s value, do the same but for the right subtree.
	case key > n.Key:
		if n.Right == nil {
			n.Right = &Node{Key: key, Value: data}
			return nil
		}
		return n.Right.Insert(key, data)
	}
	return nil
}

func (n *Node) Get(x int) int {
	if n == nil {
		return -1
	}
	switch {
	//If the current node contains the value, return the node.
	case x == n.Key:
		return n.Value
		//If the data value is less than the current node’s value, call Find for the left child node,
	case x < n.Key:
		return n.Left.Get(x)
		//else call Find for the right child node.
	default:
		return n.Right.Get(x)
	}
}
func (n *Node) FindPath(x int) []int {
	var path []int
	if hasPath := n.findPath(&path, x); !hasPath {
		return nil
	}
	return path
}
func (n *Node) findPath(path *[]int, x int) bool {
	var contains bool

	if n == nil {
		contains = false
	}
	*path = append(*path, n.Key)

	if n.Key == x {
		contains = true
	}
	if x < n.Key && n.Left != nil && n.Left.findPath(path, x) {
		contains = true
	}
	if x > n.Key && n.Right != nil && n.Right.findPath(path, x) {
		contains = true
	}
	return contains
}

func (n *Node) LCA(x int, y int) int {

	node1Path := n.FindPath(x)
	node2Path := n.FindPath(y)

	var shortLen int

	if len(node1Path) < len(node2Path) {
		shortLen = len(node1Path)
	} else {
		shortLen = len(node2Path)
	}
	for i := shortLen - 1; i >= 0; i-- {
		if node1Path[i] == node2Path[i] {
			return node1Path[i]
		}
	}
	return -1
}
