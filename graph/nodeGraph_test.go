package graph

import (
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
	graph.Add(TestType{
		n: 0,
		s: "n1",
	})

	assert.Equal(t, graph.count, 1)

	graph.Add(TestType{
		n: 1,
		s: "n2",
	})

	assert.Equal(t, graph.count, 2)
}
