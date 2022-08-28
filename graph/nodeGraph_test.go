package graph

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestType struct {
	n int
	s string
}

func TestNewGraph(t *testing.T) {
	graph := NewGraph()
	assert.Equal(t, graph.count, 0)
}

func TestNodeGraph_Add(t *testing.T) {
	graph := NewGraph()
	testType := TestType{
		n: 0,
		s: "n1",
	}
	node1 := graph.Add(testType)

	assert.Equal(t, graph.count, 1)

	graph.Add(TestType{
		n: 1,
		s: "n2",
	})

	assert.Equal(t, graph.count, 2)

	bsf := graph.Bsf(node1, testType)
	assert.Equal(t, bsf, testType)
}

func TestNodeGraph_AddEdge(t *testing.T) {
	graph := NewGraph()
	node1 := Node{
		Value: "n1",
	}

	add := graph.Add(node1)

	node2 := Node{
		Value: "n2",
	}
	graph.AddEdge(add, &node2)

	bsf := graph.Bsf(add, "n2")
	fmt.Println(bsf)
}

// A -- B -- C
// |         |
// |		 |
// D -- E -- F
func TestBox(t *testing.T) {
	graph := NewGraph()

	A, B := graph.AddEdgeValues("A", "B")
	D, E := graph.AddEdgeValues("D", "E")
	C, F := graph.AddEdgeValues("C", "F")
	graph.AddEdge(B, C)
	graph.AddEdge(C, F)
	graph.AddEdge(E, F)
	graph.AddEdge(D, E)

	bsf := graph.Bsf(A, "F")
	assert.Equal(t, bsf, "F")

}
