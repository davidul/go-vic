package graph

import "github.com/davidul/go-vic/linkedlist"

type DiNodeGraph[T comparable] struct {
	m     map[*Node[T]]*linkedlist.DoublyLinkedList[*Node[T]]
	count int
}

func NewDiGraph[T comparable]() *DiNodeGraph[T] {
	i := new(DiNodeGraph[T])
	i.init()
	return i
}

func (G *DiNodeGraph[T]) init() {
	G.m = make(map[*Node[T]]*linkedlist.DoublyLinkedList[*Node[T]])
}

func (G *DiNodeGraph[T]) Add(value T) *Node[T] {
	if len(G.m) > 0 {
		//for node, _ := range G.m {
		//bsf := G.Bsf(node, value)
		//if bsf != nil {
		//	return bsf
		//}
		//}
	}
	G.count++
	list := &linkedlist.DoublyLinkedList[*Node[T]]{}
	node := &Node[T]{Value: value}
	G.m[node] = list
	return node
}
