package linkedlist

import (
	"fmt"
)

type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

type Node struct {
	next *Node
	prev *Node
	list *LinkedList
	data any
}

func (n *Node) Next() *Node {
	if n.next != nil {
		return n.next
	}
	return nil
}

func (n *Node) Data() any {
	return n.data
}

func (L *LinkedList) Head() *Node {
	return L.head
}

// Add append at the end of the list
func (L *LinkedList) Add(data any) {
	node := &Node{
		next: nil,
		prev: nil,
		data: data}
	if L.head == nil {
		L.head = node
		L.tail = node
	} else {
		node.prev = L.tail
		L.tail.next = node
		L.tail = L.tail.next
	}
	L.count++
}

func (L *LinkedList) AddLast(data any) {
	L.Add(data)
}

func (L *LinkedList) AddFirst(data any) {
	oldHead := L.head
	list := &Node{
		next: oldHead,
		data: data,
	}
	L.head = list
	L.count++
}

// Peek returns but does not remove the head element
func (L *LinkedList) Peek() any {
	if L.head != nil {
		return L.head.data
	} else {
		return nil
	}
}

// PeekLast returns but does not remove the tail element
func (L *LinkedList) PeekLast() any {
	if L.tail != nil {
		return L.tail.data
	} else {
		return nil
	}
}

// Remove retrieves and removes the head
// memory leak or gc will collect it?
func (L *LinkedList) Remove() any {
	head := L.head
	newHead := head.next
	L.head = newHead
	L.count--
	return head.data
}

func (L *LinkedList) RemoveLast() any {
	if L.tail != nil {
		tail := L.tail
		L.tail = nil
		return tail.data
	}
	return nil
}

func (L *LinkedList) Poll() any {
	return L.Remove()
}

func (L *LinkedList) ToArray() []any {
	i := make([]any, L.count)
	e := L.head
	c := 0
	for e != nil {
		i[c] = e.data
		e = e.next
		c++
	}
	return i
}

func (L *LinkedList) Contains(v any) bool {
	for e := L.head; e != nil; e = e.Next() {
		if e.Data() == v {
			return true
		}
	}
	return false
}

func (L *LinkedList) IsEmpty() bool {
	if L.head == nil {
		return true
	}
	return false
}

//func (L LinkedList) Reverse() LinkedList {
//
//}

func (L *LinkedList) Size() int {
	return L.count
}
func (L *LinkedList) Print() {
	list := L.head
	for list != nil {
		fmt.Println(list.data)
		list = list.next
	}
}
