package graph

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

type TestType struct {
	n int
	s string
}

func TestNewGraph(t *testing.T) {
	graph := NewGraph[string]()
	assert.Equal(t, graph.count, 0)
}

func TestNodeGraph_Add(t *testing.T) {
	graph := NewGraph[TestType]()
	testType := TestType{
		n: 0,
		s: "n1",
	}
	node1 := graph.Add(testType)
	node2 := graph.Add(testType) // no duplicates

	assert.Equal(t, node1, node2)
	assert.Equal(t, graph.count, 1)

	graph.Add(TestType{
		n: 1,
		s: "n2",
	})

	assert.Equal(t, graph.count, 2)

	bsf := graph.Bsf(node1, testType)
	assert.Equal(t, bsf.Value, testType)
}

func TestNodeGraph_AddEdge(t *testing.T) {
	graph := NewGraph[string]()

	add := graph.Add("n1")
	n2 := graph.Add("n2")
	graph.AddEdge(add, n2)

	bsf := graph.Bsf(add, "n2")
	fmt.Println(bsf)
}

// A -- B -- C
// |         |
// |		 |
// D -- E -- F
func TestBox(t *testing.T) {
	graph := NewGraph[string]()

	A, B := graph.AddEdgeValues("A", "B")
	D, E := graph.AddEdgeValues("D", "E")
	C, F := graph.AddEdgeValues("C", "F")
	graph.AddEdge(B, C)
	graph.AddEdge(C, F)
	graph.AddEdge(E, F)
	graph.AddEdge(D, E)

	bsf := graph.Bsf(A, "F")
	assert.Equal(t, bsf.Value, "F")

	dsf := graph.Dsf(A, "F")
	assert.Equal(t, dsf, "F")
}

// A -- B -- C -- D -- I
//
//	|              |
//	E -- F -- H ---|
//	|
//	G
func TestBox1(t *testing.T) {
	graph := NewGraph[string]()
	A, B := graph.AddEdgeValues("A", "B")
	C := graph.Add("C")
	D := graph.Add("D")
	E := graph.Add("E")
	F := graph.Add("F")
	G := graph.Add("G")
	H := graph.Add("H")
	I := graph.Add("I")

	graph.AddEdge(B, C)
	graph.AddEdge(C, D)
	graph.AddEdge(D, I)
	graph.AddEdge(B, E)
	graph.AddEdge(E, G)
	graph.AddEdge(E, F)
	graph.AddEdge(F, H)
	graph.AddEdge(H, I)

	bsf := graph.Bsf(A, "H")
	assert.Equal(t, bsf.Value, "H")

	dsf := graph.Dsf(A, "H")
	assert.Equal(t, dsf, "H")

	path := graph.ShortestPath(A, "H")
	assert.Equal(t, path[H].i, 4)
}

func addChildren(G *NodeGraph[string], parentNode *Node[string], parentDir string, entry os.DirEntry) {
	dir, err := os.ReadDir(parentDir)
	if err != nil {
		fmt.Println(err)
	}

	for _, i := range dir {
		if i.IsDir() {
			G.Add(filepath.Join(parentDir, i.Name()))
			addChildren(G, parentNode, filepath.Join(parentDir, i.Name()), i)
		} else {
			G.AddEdgeNodeValue(parentNode, filepath.Join(parentDir, i.Name()))
		}
	}
}

//func TestX(t *testing.T) {
//	//fmt.Println("\033[31m Hello Red")
//	fmt.Println("\033[S")
//	fmt.Println("\033[C Move right")
//}

//	func TestFileSystem(t *testing.T) {
//		topDir := "/Users/david/godir-test"
//		dir, err := os.ReadDir(topDir)
//		if err != nil {
//			panic(err)
//		}
//
//		graph := NewGraph[string]()
//
//		parentNode := graph.Add(topDir)
//		for _, d := range dir {
//			if d.IsDir() {
//				next := graph.AddEdgeNodeValue(parentNode, filepath.Join(topDir, d.Name()))
//				addChildren(graph, next, filepath.Join(topDir, d.Name()), d)
//			} else {
//				graph.AddEdgeNodeValue(parentNode, filepath.Join(topDir, d.Name()))
//			}
//		}
//
//		graph.Print(parentNode)
//	}
func TestNodeGraph_Dsf(t *testing.T) {

}
