package linkedlist

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	next *Node
	data interface{}
}

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
}

func (L *LinkedList) Peek() interface{} {
	if L.head != nil {
		return L.head.data
	} else {
		return nil
	}
}

func (L *LinkedList) PeekLast() interface{} {
	if L.tail != nil {
		return L.tail.data
	} else {
		return nil
	}
}

func (L *LinkedList) Print() {
	list := L.head
	for list != nil {
		fmt.Println(list.data)
		list = list.next
	}
}
