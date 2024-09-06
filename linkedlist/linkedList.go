package linkedlist

import (
	"fmt"
)

// LinkedList Double linked list structure.
// Pointer to head and tail of the list.
// Count is the number of nodes.
type LinkedList[T comparable] struct {
	head  *Node[T]
	tail  *Node[T]
	count int
}

// Node has pointers to next and previous nodes.
// Can iterate in forward or backward direction.
type Node[T comparable] struct {
	next *Node[T]
	prev *Node[T]
	data T
}

type LinkedListIterator[T comparable] interface {
	Next() T
	Done() bool
}

// NewLinkedList creates new linked list with data T.
// Head and tail are the same.
func NewLinkedList[T comparable](data T) *LinkedList[T] {
	head := &Node[T]{
		data: data,
		next: nil,
		prev: nil,
	}
	tail := head

	return &LinkedList[T]{
		head:  head,
		tail:  tail,
		count: 1,
	}
}

// Next returns next node. If the next
// pointer is nil returns nil
func (n *Node[T]) Next() *Node[T] {
	if n.next != nil {
		return n.next
	}
	return nil
}

// Data returns Node data
func (n *Node[T]) Data() T {
	return n.data
}

// Head returns head of the list
func (L *LinkedList[T]) Head() *Node[T] {
	return L.head
}

// Add append at the end of the list
func (L *LinkedList[T]) Add(data T) *Node[T] {
	node := &Node[T]{
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
	return node
}

// AddLast append to tail
func (L *LinkedList[T]) AddLast(data T) {
	L.Add(data)
}

// AddFirst prepend to head. New head is being created.
func (L *LinkedList[T]) AddFirst(data T) {
	oldHead := L.head
	list := &Node[T]{
		next: oldHead,
		data: data,
	}
	L.head = list
	L.count++
}

// Peek returns but does not remove the head element
func (L *LinkedList[T]) Peek() any {
	if L.head != nil {
		return L.head.data
	} else {
		return nil
	}
}

// PeekLast returns but does not remove the tail element
func (L *LinkedList[T]) PeekLast() T {
	if L.tail != nil {
		return L.tail.data
	} else {
		var result T
		return result
	}
}

// Remove retrieves and removes the head
func (L *LinkedList[T]) Remove() T {
	if L.count == 0 {
		var result T
		return result
	}
	head := L.head
	newHead := head.next
	if newHead == nil {
		L.count--
		return head.data
	}
	L.head = newHead
	L.head.prev = nil
	L.count--
	return head.data
}

// RemoveLast remove and return tail of the list
func (L *LinkedList[T]) RemoveLast() T {
	if L.count == 0 {
		var result T
		return result
	}
	if L.head == L.tail {
		L.count--
		return L.head.data
	}
	if L.tail != nil {
		tail := L.tail
		L.tail.prev.next = nil
		L.tail = tail.prev
		L.count--
		return tail.data
	} else if L.head != nil {
		data := L.head.data
		L.head = nil
		L.count--
		return data
	}

	var result T
	return result
}

// RemoveNode by reference to the node
func (L *LinkedList[T]) RemoveNode(node *Node[T]) {
	for e := L.head; e != nil; e = e.Next() {
		if e == node {
			prev := e.prev
			next := e.next
			prev.next = next
			next.prev = prev
			e = nil
			break
		}
	}
}

// Poll remove and return
func (L *LinkedList[T]) Poll() T {
	return L.Remove()
}

// Converts linked list to array
func (L *LinkedList[T]) ToArray() []T {
	i := make([]T, L.count)
	e := L.head
	c := 0
	for e != nil {
		i[c] = e.data
		e = e.next
		c++
	}
	return i
}

func (L *LinkedList[T]) Contains(v T) bool {
	for e := L.head; e != nil; e = e.Next() {
		if e.Data() == v {
			return true
		}
	}
	return false
}

// IsEmpty returns true if the list is empty
func (L *LinkedList[T]) IsEmpty() bool {
	if L.count == 0 {
		return true
	}
	return false
}

//func (L LinkedList) Reverse() LinkedList {
//
//}

// Size return size of the linkedlist
func (L *LinkedList[T]) Size() int {
	return L.count
}

func (L *LinkedList[T]) Print() {
	list := L.head
	for list != nil {
		fmt.Println(list.data)
		list = list.next
	}
}

func (L *LinkedList[T]) Compare(other *LinkedList[T]) bool {
	if L.Size() == other.Size() {
		oe := other.Head()
		for e := L.head; e != nil; e = e.next {
			if e.Data() != oe.Data() {
				return false
			} else {
				oe = oe.Next()
			}
		}
	} else {
		return false
	}
	return true
}

func (L *LinkedList[T]) AddAll(other *LinkedList[T]) {
	head := other.Head()
	for head != nil {
		L.Add(head.data)
		head = head.next
	}
}

func (L *LinkedList[T]) Append(other *LinkedList[T]) {
	if L.tail == L.head {
		L.head.next = other.head
		other.head.prev = L.head
		L.tail = other.tail
		L.count += other.count
	} else {
		L.tail.next = other.head
	}
}

type LLIterator[T comparable] struct {
	l    *LinkedList[T]
	next *Node[T]
}

func (L *LinkedList[T]) CreateIterator() *LLIterator[T] {
	return &LLIterator[T]{
		l:    L,
		next: L.head,
	}
}
func (L *LLIterator[T]) Next() T {
	data := L.next.data
	L.next = L.next.next
	return data
}

// Done returns true if there
// are no more elements
func (L *LLIterator[T]) Done() bool {
	if L.next == nil {
		return true
	}

	return false
}
