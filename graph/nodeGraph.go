package graph

import (
	"github.com/davidul/go-vic/linkedlist"
)

type NodeGraph struct {
	m     map[Node]linkedlist.LinkedList
	count int
}

type Node struct {
	Value any
	Id    int
}

func NewGraph() *NodeGraph {
	i := new(NodeGraph)
	i.Init()
	return i
}

func (G *NodeGraph) Init() {
	G.m = make(map[Node]linkedlist.LinkedList)
}

func (G *NodeGraph) Add(value any) int {
	G.count++
	list := linkedlist.LinkedList{}
	node := Node{Value: value,
		Id: G.count}
	list.Add(node)
	G.m[node] = list
	return G.count
}

func (G *NodeGraph) AddEdge(start Node, value any) int {
	list := G.m[start]
	G.count++
	node := Node{Id: G.count, Value: value}
	list.Add(node)
	G.m[node] = linkedlist.LinkedList{}
	return G.count
}

func (G *NodeGraph) Bsf(root Node, goal any) any {
	visited := linkedlist.LinkedList{}
	queue := linkedlist.LinkedList{}
	queue.Add(root)
	for h := queue.Head(); h != nil; h = h.Next() {
		v := queue.Poll()
		if v == goal {
			return v
		}
		list := G.m[v.(Node)]
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e) {
				queue.Add(e.Data())
			}
		}
	}
	return nil
}

func (G *NodeGraph) dsf() {

}
