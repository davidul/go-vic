package graph

import "github.com/davidul/go-vic/linkedlist"

type WNode[T comparable] struct {
	value  *Node[T]
	weight int64
}

type WNodeGraph[T comparable] struct {
	m map[*Node[T]]linkedlist.DoublyLinkedList[*WNode[T]]
}

func NewWGraph[T comparable]() *WNodeGraph[T] {
	w := new(WNodeGraph[T])
	w.m = make(map[*Node[T]]linkedlist.DoublyLinkedList[*WNode[T]])
	return w
}

func (G *WNodeGraph[T]) Add(value T) *Node[T] {
	n := new(Node[T])
	n.Value = value
	G.m[n] = linkedlist.DoublyLinkedList[*WNode[T]]{}
	return n
}

func (G *WNodeGraph[T]) AddEdge(start *Node[T], end *Node[T], weight int64) {
	l := G.m[start]
	l2 := G.m[end]

	w := new(WNode[T])
	w.weight = weight
	w.value = end
	l.Add(w)

	w2 := new(WNode[T])
	w2.weight = weight
	w2.value = start
	l2.Add(w2)
}
