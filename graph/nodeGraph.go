package graph

import (
	"fmt"
	"github.com/davidul/go-vic/linkedlist"
)

// NodeGraph Graph represented as adjacency list
// *Node is a key
// You can have duplicate in value
type NodeGraph[T interface{}] struct {
	m     map[*Node]*linkedlist.LinkedList[T]
	count int
}

// Node is just a wrapper for any
type Node struct {
	Value any
}

type Distance struct {
	i int
	n *Node
}

// NewGraph initialize the empty graph
func NewGraph[T interface{}]() *NodeGraph[T] {
	i := new(NodeGraph[T])
	i.init()
	return i
}

func (G *NodeGraph[T]) init() {
	G.m = make(map[*Node]*linkedlist.LinkedList[T])
}

// Add value to a graph and return it as *Node
// If value already exists inside the graph, this value
// is returned and no new node is added
func (G *NodeGraph[T]) Add(value any) *Node {
	if len(G.m) > 0 {
		for node, _ := range G.m {
			bsf := G.Bsf(node, value)
			if bsf != nil {
				return bsf
			}
		}
	}
	G.count++
	list := &linkedlist.LinkedList[T]{}
	node := &Node{Value: value}
	G.m[node] = list
	return node
}

// AddEdge creates new edge between two nodes start and end.
func (G *NodeGraph[T]) AddEdge(start *Node, end *Node) {
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
		list := linkedlist.LinkedList[T]{}
		list.Add(start)
		G.m[end] = &list
		return
	}

	if !e1 && e2 {
		l2.Add(start)
		list := linkedlist.LinkedList[T]{}
		list.Add(end)
		G.m[start] = &list
		return
	}

	if !e1 && !e2 {
		list1 := linkedlist.LinkedList[T]{}
		list2 := linkedlist.LinkedList[T]{}
		list1.Add(end)
		list2.Add(start)
		G.m[start] = &list1
		G.m[end] = &list2
	}
}

// AddEdgeValues Create new nodes for start and end value.
// Adds an edge between these two.
func (G *NodeGraph[T]) AddEdgeValues(start any, end any) (n1 *Node, n2 *Node) {
	node1 := new(Node)
	node2 := new(Node)
	node1.Value = start
	node2.Value = end
	G.AddEdge(node1, node2)
	return node1, node2
}

func (G *NodeGraph[T]) AddEdgeNodeValue(start *Node, end any) *Node {
	n := new(Node)
	n.Value = end
	list := G.m[start]
	list.Add(n)
	G.m[n] = new(linkedlist.LinkedList[T])
	return n
}

// Bsf Breadth First Search
func (G *NodeGraph[T]) Bsf(root *Node, goal any) *Node {
	visited := linkedlist.LinkedList[T]{}
	queue := linkedlist.LinkedList[T]{}
	queue.Add(root)

	for !queue.IsEmpty() {
		v := queue.Poll()
		visited.Add(v)
		vv := v.(*Node).Value
		if vv == goal {
			return v.(*Node)
		}
		node := v.(*Node)
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
func (G *NodeGraph[T]) ShortestPath(root *Node, goal any) map[*Node]Distance {

	visited := linkedlist.LinkedList[T]{}
	queue := linkedlist.LinkedList[T]{}
	queue.Add(root)

	distanceMap := make(map[*Node]Distance)
	distanceMap[root] = Distance{
		i: 0,
		n: root,
	}

	for !queue.IsEmpty() {
		v := queue.Poll()
		//fmt.Printf("Processing node %s \n", v.(*Node).Value)
		visited.Add(v)
		vv := v.(*Node).Value
		if vv == goal {
			return distanceMap
		}
		node := v.(*Node)
		list := G.m[node]
		for e := list.Head(); e != nil; e = e.Next() {
			//fmt.Printf("Children %s \n", e.Data().(*Node).Value)
			distance := distanceMap[v.(*Node)]
			//fmt.Printf("Distance %d \n", distance.i)
			if !visited.Contains(e.Data()) {
				//fmt.Println("Not visited yet")
				d := Distance{
					i: distance.i + 1,
					n: e.Data().(*Node),
				}
				//fmt.Printf("New Distance %d \n", distance.i+1)
				if nn, ok := distanceMap[e.Data().(*Node)]; ok {
					if nn.i >= d.i {
						distanceMap[e.Data().(*Node)] = d
					}
				} else {
					distanceMap[e.Data().(*Node)] = d
				}
				queue.Add(e.Data())
			}
		}
	}
	return nil
}

// Dsf Depth first search
func (G *NodeGraph[T]) Dsf(root *Node, goal any) any {
	stack := linkedlist.LinkedList[T]{}
	visited := linkedlist.LinkedList[T]{}

	stack.AddFirst(root)

	for !stack.IsEmpty() {
		v := stack.Poll()
		visited.Add(v)
		value := v.(*Node).Value
		if value == goal {
			return value
		}
		list := G.m[v.(*Node)]
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e.Data()) {
				stack.AddFirst(e.Data())
			}
		}
	}
	return nil
}

func (G *NodeGraph[T]) SimpleDelete(node *Node) {
	if list, ok := G.m[node]; ok {
		delete(G.m, node)
		for e := list.Head(); e != nil; e = e.Next() {
			list.RemoveNode(e)
		}
	}
}

func (G *NodeGraph[T]) Print(node *Node) {
	visited := linkedlist.LinkedList[T]{}
	queue := linkedlist.LinkedList[T]{}
	queue.Add(node)

	for !queue.IsEmpty() {
		v := queue.Poll()
		visited.Add(v)
		vv := v.(*Node).Value
		fmt.Println(vv)
		node := v.(*Node)
		list := G.m[node]
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e.Data()) {
				queue.Add(e.Data())
			}
		}
	}
}
