package graph

import "data-structures-go/linkedlist"

type IntGraph struct {
	m     map[int]linkedlist.LinkedList
	count int
}

func NewGraph() *IntGraph {
	i := new(IntGraph)
	i.Init()
	return i
}

func (G *IntGraph) Init() {
	G.m = make(map[int]linkedlist.LinkedList)
}

func (G *IntGraph) Add(value any) int {
	G.count++
	list := linkedlist.LinkedList{}
	node := Node{Value: value,
		Id: G.count}
	list.Add(node)
	G.m[G.count] = list
	return G.count
}

func (G *IntGraph) AddEdge(start int, value any) int {
	list := G.m[start]
	G.count++
	node := Node{Id: G.count, Value: value}
	list.Add(node)
	G.m[G.count] = linkedlist.LinkedList{}
	return G.count
}

func (G *IntGraph) Bsf(root int, goal any) any {
	visited := linkedlist.LinkedList{}
	queue := linkedlist.LinkedList{}
	queue.Add(root)
	for h := queue.Head(); h != nil; h = h.Next() {
		v := queue.Poll()
		if v == goal {
			return v
		}
		list := G.m[v.(int)]
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e) {
				queue.Add(e.Data())
			}
		}
	}
	return nil
}

func (G *IntGraph) dsf() {

}
