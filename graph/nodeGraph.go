package graph

import (
	"fmt"
	"github.com/davidul/go-vic/linkedlist"
)

// NodeGraph Graph represented as adjacency list
// *Node is a key
// You can have duplicate in value
type NodeGraph[T comparable] struct {
	m     map[*Node[T]]*linkedlist.LinkedList[*Node[T]]
	count int
}

// Node is just a wrapper for any
type Node[T comparable] struct {
	Value T
}

type Distance[T comparable] struct {
	i int
	n *Node[T]
}

// NewGraph initialize the empty graph
func NewGraph[T comparable]() *NodeGraph[T] {
	i := new(NodeGraph[T])
	i.init()
	return i
}

func (G *NodeGraph[T]) init() {
	G.m = make(map[*Node[T]]*linkedlist.LinkedList[*Node[T]])
}

// Add value to a graph and return it as *Node
// If value already exists inside the graph, this value
// is returned and no new node is added
func (G *NodeGraph[T]) Add(value T) *Node[T] {
	if len(G.m) > 0 {
		for node, _ := range G.m {
			bsf := G.Bsf(node, value)
			if bsf != nil {
				return bsf
			}
		}
	}
	G.count++
	list := &linkedlist.LinkedList[*Node[T]]{}
	node := &Node[T]{Value: value}
	G.m[node] = list
	return node
}

// AddEdge creates new edge between two nodes start and end.
func (G *NodeGraph[T]) AddEdge(start *Node[T], end *Node[T]) {
	l1, e1 := G.m[start]
	l2, e2 := G.m[end]

	if e1 && e2 {
		if !l1.Contains(end) {
			l1.Add(end)
		}
		if !l2.Contains(start) {
			l2.Add(start)
		}
		return
	}

	if e1 && !e2 {
		l1.Add(end)
		list := linkedlist.LinkedList[*Node[T]]{}
		list.Add(start)
		G.m[end] = &list
		return
	}

	if !e1 && e2 {
		l2.Add(start)
		list := linkedlist.LinkedList[*Node[T]]{}
		list.Add(end)
		G.m[start] = &list
		return
	}

	if !e1 && !e2 {
		list1 := linkedlist.LinkedList[*Node[T]]{}
		list2 := linkedlist.LinkedList[*Node[T]]{}
		list1.Add(end)
		list2.Add(start)
		G.m[start] = &list1
		G.m[end] = &list2
	}
}

// AddEdgeValues Create new nodes for start and end value.
// Adds an edge between these two.
func (G *NodeGraph[T]) AddEdgeValues(start T, end T) (n1 *Node[T], n2 *Node[T]) {
	node1 := new(Node[T])
	node2 := new(Node[T])
	node1.Value = start
	node2.Value = end
	G.AddEdge(node1, node2)
	return node1, node2
}

func (G *NodeGraph[T]) AddEdgeNodeValue(start *Node[T], end T) *Node[T] {
	n := new(Node[T])
	n.Value = end
	list := G.m[start]
	list.Add(n)
	G.m[n] = new(linkedlist.LinkedList[*Node[T]])
	return n
}

// Bsf Breadth First Search
func (G *NodeGraph[T]) Bsf(root *Node[T], goal T) *Node[T] {
	visited := linkedlist.LinkedList[*Node[T]]{}
	queue := linkedlist.LinkedList[*Node[T]]{}
	queue.Add(root)

	for !queue.IsEmpty() {
		v := queue.Poll()
		visited.Add(v)
		vv := v.Value
		if vv == goal {
			return v
		}
		node := v
		list := G.m[node]
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e.Data()) {
				queue.Add(e.Data())
			}
		}
	}
	return nil
}

// ShortestPath with implemented as BSF
// Returned map contains the shortest distance from root to
// any node on the path
func (G *NodeGraph[T]) ShortestPath(root *Node[T], goal T) map[*Node[T]]Distance[T] {

	visited := linkedlist.LinkedList[*Node[T]]{}
	queue := linkedlist.LinkedList[*Node[T]]{}
	queue.Add(root)

	distanceMap := make(map[*Node[T]]Distance[T])
	distanceMap[root] = Distance[T]{
		i: 0,
		n: root,
	}

	for !queue.IsEmpty() {
		v := queue.Poll()
		//fmt.Printf("Processing node %s \n", v.(*Node).Value)
		visited.Add(v)
		vv := v.Value
		if vv == goal {
			return distanceMap
		}
		node := v
		list := G.m[node]
		for e := list.Head(); e != nil; e = e.Next() {
			//fmt.Printf("Children %s \n", e.Data().(*Node).Value)
			distance := distanceMap[v]
			//fmt.Printf("Distance %d \n", distance.i)
			if !visited.Contains(e.Data()) {
				//fmt.Println("Not visited yet")
				d := Distance[T]{
					i: distance.i + 1,
					n: e.Data(),
				}
				//fmt.Printf("New Distance %d \n", distance.i+1)
				if nn, ok := distanceMap[e.Data()]; ok {
					if nn.i >= d.i {
						distanceMap[e.Data()] = d
					}
				} else {
					distanceMap[e.Data()] = d
				}
				queue.Add(e.Data())
			}
		}
	}
	return nil
}

// Dsf Depth first search
func (G *NodeGraph[T]) Dsf(root *Node[T], goal T) T {
	stack := linkedlist.LinkedList[*Node[T]]{}
	visited := linkedlist.LinkedList[*Node[T]]{}

	stack.AddFirst(root)

	for !stack.IsEmpty() {
		v := stack.Poll()
		visited.Add(v)
		if v.Value == goal {
			return v.Value
		}
		list := G.m[v]
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e.Data()) {
				stack.AddFirst(e.Data())
			}
		}
	}
	var result T
	return result
}

func (G *NodeGraph[T]) SimpleDelete(node *Node[T]) {
	if list, ok := G.m[node]; ok {
		delete(G.m, node)
		for e := list.Head(); e != nil; e = e.Next() {
			list.RemoveNode(e)
		}
	}
}

func (G *NodeGraph[T]) Print(node *Node[T]) {
	visited := linkedlist.LinkedList[*Node[T]]{}
	queue := linkedlist.LinkedList[*Node[T]]{}
	queue.Add(node)

	for !queue.IsEmpty() {
		v := queue.Poll()
		visited.Add(v)
		vv := v.Value
		fmt.Println(vv)
		node := v
		list := G.m[node]
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e.Data()) {
				queue.Add(e.Data())
			}
		}
	}
}
