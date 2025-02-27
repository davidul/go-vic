package graph

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWGraph(t *testing.T) {
	graph := NewWGraph[string]()
	assert.NotNil(t, graph)
}

func TestWNodeGraph_Add(t *testing.T) {
	graph := NewWGraph[string]()
	n1 := graph.Add("A")
	n2 := graph.Add("B")
	graph.AddEdge(n1, n2, 100)
	fmt.Println("")
}
