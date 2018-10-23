package lca

import "testing"

func CreateAGraph() *Graph {
	testGraph := CreateGraph(8)
	testGraph.AddEdge(0, 1)
	testGraph.AddEdge(0, 2)
	testGraph.AddEdge(2, 3)
	testGraph.AddEdge(1, 2)
	testGraph.AddEdge(2, 4)
	testGraph.AddEdge(1, 5)
	testGraph.AddEdge(4, 7)
	testGraph.AddEdge(3, 6)

	return testGraph
}
func TestAddEdge(t *testing.T) {
	testGraph := CreateGraph(3)
	testGraph.AddEdge(1, 2)
	testGraph.AddEdge(0, 1)
	testGraph.AddEdge(0, 2)

	expected := true
	ans := testGraph.ValidEdge(0, 1)
	if ans != expected {
		t.Errorf("Failed to add an Edge. Expected: %v, Got: %v", expected, ans)
	}
}

//Case 1: Tests for the LCA of two nodes where one is the ancestor of the other.
func TestLca(t *testing.T) {
	testGraph := CreateAGraph()
	ans := testGraph.LCA(1, 2)
	expected := 1
	if ans != expected {
		t.Errorf("Failed Case 1. Expected: %v, Got: %v", expected, ans)
	}
}
func TestLca2(t *testing.T) {
	testGraph := CreateAGraph()
	ans := testGraph.LCA(5, 3)
	expected := 1
	if ans != expected {
		t.Errorf("Failed Case 2. Expected: %v, Got: %v", expected, ans)
	}

}
func TestLca3(t *testing.T) {
	testGraph := CreateAGraph()
	ans := testGraph.LCA(6, 7)
	expected := 2
	if ans != expected {
		t.Errorf("Failed Case 3. Expected: %v, Got: %v", expected, ans)
	}

}

//Case 4: Test for the LCA of two nodes that do not have and ancestor.
func TestLca4(t *testing.T) {
	testGraph := CreateAGraph()
	testGraph.RemoveEdge(4, 7)
	ans := testGraph.LCA(7, 3)
	expected := -1
	if ans != expected {
		t.Errorf("Failed Case 3. Expected: %v, Got: %v", expected, ans)
	}
}

// func TestEmptyGraph(t *testing.T) {
// 	x := EmptyGraph()
// 	expected := Graph{}
// 	if *x != expected {
// 		t.Errorf("Your program does not create a binary tree")
//	}
//}

// func TestInsert(t *testing.T) {

// 	x := CreateAGraph()
// 	x.AddEdge(1, 2, 1)
// 	expected := Graph{
// 		NumNodes: 1,
// 		Edges:    &Edge{From: 1, To: 2, Weight: 1},
// 	}
// 	if x.NumNodes != expected.NumNodes {
// 		t.Errorf("Your program does not insert a node correctly")
// 	}

// 	x.AddEdge(1, 3, 1)
// 	x.AddEdge(3, 4, 1)
// 	expected = Graph{
// 		NumNodes: 4,
// 		Edges:    &Edge{From: 1, To: 2, Weight: 1},
// 		&Edge{From: 1, To: 3, Weight: 1},
// 		&Edge{From: 3, To: 4, Weight: 1},
// 	}
// 	if x.NumNodes != expected.NumNodes {
// 		t.Errorf("Your program does not insert many nodes correctly")
// 	}

// }
// func TestGet(t *testing.T) {
// 	x := Node{
// 		Left: &Node{
// 			Right: &Node{
// 				Key:   3,
// 				Value: 3,
// 			},
// 			Key:   1,
// 			Value: 1,
// 		},
// 		Key:   4,
// 		Value: 4,
// 	}
// 	expected := 1
// 	value := x.Get(1)

// 	if value != expected {
// 		t.Errorf("Your program does not return the searched node correctly. Expected: %v, Got: %v", expected, value)
// 	}
// }

// func TestFindPath(t *testing.T) {
// 	x := CreateATree()
// 	expected := []int{4, 6, 5}

// 	ans := x.FindPath(5)
// 	if !cmp.Equal(expected, ans) {
// 		t.Errorf("Expected: %v, Got: %v", expected, ans)
// 	}
// }
// func TestLca(t *testing.T) {
// 	x := CreateATree()
// 	expected := 6
// 	ans := x.LCA(5, 7)

// 	if ans != expected {
// 		t.Errorf("Your program does not return the lowest common ancestor node correctly. Expected: %v, Got: %v", expected, ans)
// 	}

// }
