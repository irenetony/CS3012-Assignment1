package lca

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func CreateAGraph() *Graph {
	return &Graph{
		NumNodes: 5,
		Edges:    make([][]Edge, 5),
	}
}
func TestEmptyTree(t *testing.T) {
	x := EmptyBinaryTree()
	expected := Node{}
	if *x != expected {
		t.Errorf("Your program does not create a binary tree")
	}
}

func TestInsert(t *testing.T) {

	x := CreateAGraph()
	x.AddEdge(1, 2, 1)
	expected := Graph{
		NumNodes: 1,
		Edges: &Edge{
			From:   1,
			To:     2,
			Weight: 1,
		}
	}
	if x.NumNodes != expected.NumNodes {
		t.Errorf("Your program does not insert a node correctly")
	}

	x.Insert(1, 1)
	x.Insert(3, 3)
	expected = Node{
		Left: &Node{
			Right: &Node{
				Key:   3,
				Value: 3,
			},
			Key:   1,
			Value: 1,
		},
		Key:   4,
		Value: 4,
	}
	if x.Value != expected.Value || x.Left.Value != expected.Left.Value || x.Left.Right.Value != expected.Left.Right.Value {
		t.Errorf("Your program does not insert many nodes correctly")
	}

}
func TestGet(t *testing.T) {
	x := Node{
		Left: &Node{
			Right: &Node{
				Key:   3,
				Value: 3,
			},
			Key:   1,
			Value: 1,
		},
		Key:   4,
		Value: 4,
	}
	expected := 1
	value := x.Get(1)

	if value != expected {
		t.Errorf("Your program does not return the searched node correctly. Expected: %v, Got: %v", expected, value)
	}
}

func TestFindPath(t *testing.T) {
	x := CreateATree()
	expected := []int{4, 6, 5}

	ans := x.FindPath(5)
	if !cmp.Equal(expected, ans) {
		t.Errorf("Expected: %v, Got: %v", expected, ans)
	}
}
func TestLca(t *testing.T) {
	x := CreateATree()
	expected := 6
	ans := x.LCA(5, 7)

	if ans != expected {
		t.Errorf("Your program does not return the lowest common ancestor node correctly. Expected: %v, Got: %v", expected, ans)
	}

}
