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
