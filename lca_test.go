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
func TestCase1(t *testing.T) {
	testGraph := CreateAGraph()
	ans := testGraph.LCA(1, 2)
	expected := 1
	if ans != expected {
		t.Errorf("Failed Case 1. Expected: %v, Got: %v", expected, ans)
	}
}
func TestCase3(t *testing.T) {
	testGraph := CreateAGraph()
	ans := testGraph.LCA(6, 7)
	expected := 2
	if ans != expected {
		t.Errorf("Failed Case 3. Expected: %v, Got: %v", expected, ans)
	}

}

func TestCase2(t *testing.T) {
	testGraph := CreateAGraph()
	ans := testGraph.LCA(5, 3)
	expected := 1
	if ans != expected {
		t.Errorf("Failed Case 2. Expected: %v, Got: %v", expected, ans)
	}

}

//Case 4: Test for the LCA of two nodes that do not have and ancestor.
func TestCase4(t *testing.T) {
	testGraph := CreateAGraph()
	testGraph.RemoveEdge(4, 7)
	ans := testGraph.LCA(7, 3)
	expected := -1
	if ans != expected {
		t.Errorf("Failed Case 4. Expected: %v, Got: %v", expected, ans)
	}
}
func TestCase5(t *testing.T) {
	testGraph := CreateAGraph()
	testGraph.RemoveEdge(4, 7)
	ans := testGraph.LCA(3, 7)
	expected := -1
	if ans != expected {
		t.Errorf("Failed Case 5. Expected: %v, Got: %v", expected, ans)
	}
}

//Case 6: Test for the lca of the same node, which should be the same node
func TestCase6(t *testing.T) {
	testGraph := CreateAGraph()
	ans := testGraph.LCA(3, 3)
	expected := 3
	if ans != expected {
		t.Errorf("Failed Case 6. Expected: %v, Got: %v", expected, ans)
	}
}
