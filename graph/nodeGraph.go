package graph

import (
	"fmt"
	"github.com/davidul/go-vic/linkedlist"
)

type NodeGraph struct {
	m     map[*Node]*linkedlist.LinkedList
	count int
}

type Node struct {
	Value any
}

func NewGraph() *NodeGraph {
	i := new(NodeGraph)
	i.Init()
	return i
}

func (G *NodeGraph) Init() {
	G.m = make(map[*Node]*linkedlist.LinkedList)
}

func (G *NodeGraph) Add(value any) *Node {
	G.count++
	list := &linkedlist.LinkedList{}
	node := &Node{Value: value}
	G.m[node] = list
	return node
}

func (G *NodeGraph) AddEdge(start *Node, end *Node) {
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
		list := linkedlist.LinkedList{}
		list.Add(start)
		G.m[end] = &list
		return
	}

	if !e1 && e2 {
		l2.Add(start)
		list := linkedlist.LinkedList{}
		list.Add(end)
		G.m[start] = &list
		return
	}

	if !e1 && !e2 {
		list1 := linkedlist.LinkedList{}
		list2 := linkedlist.LinkedList{}
		list1.Add(end)
		list2.Add(start)
		G.m[start] = &list1
		G.m[end] = &list2
	}
}

func (G *NodeGraph) AddEdgeValues(start any, end any) (n1 *Node, n2 *Node) {
	node1 := new(Node)
	node2 := new(Node)
	node1.Value = start
	node2.Value = end
	G.AddEdge(node1, node2)
	return node1, node2
}

func (G *NodeGraph) AddEdgeNodeValue(start *Node, end any) *Node {
	n := new(Node)
	n.Value = end
	list := G.m[start]
	list.Add(n)
	G.m[n] = new(linkedlist.LinkedList)
	return n
}

func (G *NodeGraph) Bsf(root *Node, goal any) any {
	visited := linkedlist.LinkedList{}
	queue := linkedlist.LinkedList{}
	queue.Add(root)

	for !queue.IsEmpty() {
		v := queue.Poll()
		visited.Add(v)
		vv := v.(*Node).Value
		if vv == goal {
			return vv
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

func (G *NodeGraph) ShortestPath(root *Node, goal any) any {
	type Distance struct {
		i int
		n *Node
	}
	visited := linkedlist.LinkedList{}
	queue := linkedlist.LinkedList{}
	queue.Add(root)

	distanceMap := make(map[*Node]Distance)
	distanceMap[root] = Distance{
		i: 0,
		n: root,
	}

	for !queue.IsEmpty() {
		v := queue.Poll()
		visited.Add(v)
		vv := v.(*Node).Value
		if vv == goal {
			return vv
		}
		node := v.(*Node)
		list := G.m[node]
		for e := list.Head(); e != nil; e = e.Next() {
			distance := distanceMap[v.(*Node)]
			if !visited.Contains(e.Data()) {
				d := Distance{
					i: distance.i + 1,
					n: e.Data().(*Node),
				}
				distanceMap[e.Data().(*Node)] = d
				queue.Add(e.Data())
			}
		}

		for n := range distanceMap {
			fmt.Printf("Distance to root %d from %s \n", distanceMap[n].i, distanceMap[n].n.Value)
		}
	}
	return nil
}

func (G *NodeGraph) dsf() {

}
