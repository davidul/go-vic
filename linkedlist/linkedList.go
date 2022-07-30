package linkedlist

import "fmt"

type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

type Node struct {
	next *Node
	data interface{}
}

/*type LinkedListIterator struct {
	index int
	list  *LinkedList
}*/

// a -> b -> c ->
//
func (L *LinkedList) Add(data interface{}) {
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

func (L *LinkedList) AddLast(data interface{}) {
	L.Add(data)
}

func (L *LinkedList) AddFirst(data interface{}) {
	oldHead := L.head
	list := &Node{
		next: oldHead,
		data: data,
	}
	L.head = list
	L.count++
}

// Peek returns but does not remove the head element
func (L *LinkedList) Peek() interface{} {
	if L.head != nil {
		return L.head.data
	} else {
		return nil
	}
}

// PeekLast returns but does not remove the tail element
func (L *LinkedList) PeekLast() interface{} {
	if L.tail != nil {
		return L.tail.data
	} else {
		return nil
	}
}

// Remove retrieves and removes the head
// memory leak or gc will collect it?
func (L *LinkedList) Remove() interface{} {
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

func (L *LinkedList) ToArray() []interface{} {
	i := make([]interface{}, L.count)
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
