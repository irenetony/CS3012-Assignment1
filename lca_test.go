package lca

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func CreateATree() *Node {
	return &Node{
		Left: &Node{
			Right: &Node{
				Key:   3,
				Value: 3,
			},
			Key:   1,
			Value: 1,
		},
		Right: &Node{
			Left: &Node{
				Key:   5,
				Value: 5,
			},
			Right: &Node{
				Key:   7,
				Value: 7,
			},
			Key:   6,
			Value: 6,
		},
		Key:   4,
		Value: 4,
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
	// var nilNode *Node

	// nilNode, err := nilNode.Insert(1, 1)
	// if err != nil {
	// 	t.Error(err)
	// }

	x := CreateATree()
	x.Insert(4, 4)
	expected := Node{Key: 4, Value: 4}
	if x.Value != expected.Value {
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
