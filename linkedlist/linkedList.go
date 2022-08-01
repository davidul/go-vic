package linkedlist

import "fmt"

type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

type Node struct {
	next *Node
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

// a -> b -> c ->
//
func (L *LinkedList) Add(data any) {
	list := &Node{
		next: nil,
		data: data}
	if L.head == nil {
		L.head = list
		L.tail = list
	} else {
		L.tail.next = list
		L.tail = list
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

/*func (L *LinkedList) createIterator() iterator {
	return &LinkedListIterator{
		list: L,
	}
}*/

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

func (L *LinkedList) Print() {
	list := L.head
	for list != nil {
		fmt.Println(list.data)
		list = list.next
	}
}
