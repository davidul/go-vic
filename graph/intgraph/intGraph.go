package intgraph

import (
	"github.com/davidul/go-vic/linkedlist"
)

type IntGraph struct {
	m     map[int]linkedlist.DoublyLinkedList[int]
	count int
}

func NewIntGraph() *IntGraph {
	i := new(IntGraph)
	i.initGraph()
	return i
}

func (G *IntGraph) initGraph() {
	G.m = make(map[int]linkedlist.DoublyLinkedList[int])
}

func (G *IntGraph) Add(value int) int {
	G.count++
	list := linkedlist.DoublyLinkedList[int]{}
	list.Add(value)
	G.m[G.count] = list
	return G.count
}

func (G *IntGraph) AddEdge(start int, value int) int {
	list := G.m[start]
	G.count++
	list.Add(value)
	G.m[value] = linkedlist.DoublyLinkedList[int]{}
	return value
}

func (G *IntGraph) Bsf(root int, goal int) any {
	visited := linkedlist.DoublyLinkedList[int]{}
	queue := linkedlist.DoublyLinkedList[int]{}
	queue.Add(root)
	visited.Add(root)
	for !queue.IsEmpty() {
		v := queue.Poll()
		if v == goal {
			return v
		}
		list := G.m[v]
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e.Data()) {
				queue.Add(e.Data())
				visited.Add(e.Data())
			}
		}
	}
	return nil
}

func (G *IntGraph) dsf() {

}
