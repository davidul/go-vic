package linkedlist

// LinkedList Simple linked list structure.
// Pointer to head and tail of the list.
// Count is the number of node.
type LinkedList[T comparable] struct {
	head  *Node[T]
	tail  *Node[T]
	count int
}

type Node[T comparable] struct {
	next *Node[T]
	prev *Node[T]
	//list *LinkedList
	data T
}

func NewLinkedList[T comparable](data T) *LinkedList[T] {
	l := LinkedList[T]{}
	l.head = &Node[T]{}
	l.head.data = data
	l.tail = &Node[T]{}
	l.tail.prev = l.head
	l.head.next = l.tail
	l.count = 1
	return &l
}

func (n *Node[T]) Next() *Node[T] {
	if n.next != nil {
		return n.next
	}
	return nil
}

// Data Node data
func (n *Node[T]) Data() T {
	return n.data
}

// Head return head of the list
func (L *LinkedList[T]) Head() *Node[T] {
	return L.head
}

// Add append at the end of the list
// H == T
func (L *LinkedList[T]) Add(data T) *Node[T] {
	node := &Node[T]{
		next: nil,
		prev: nil,
		data: data}
	if L.head == nil {
		L.head = node
		L.tail = &Node[T]{}
		L.tail.prev = L.head
		L.head.next = L.tail
	} else {

		// A N -> T
		a := L.tail.prev
		node.next = L.tail
		node.prev = L.tail.prev
		L.tail.prev = node
		a.next = node
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
	//head -> tail
	if L.count == 0 {
		var result T
		return result
	}
	if L.head.next.next == nil {
		data := L.head.data
		L.count--
		return data
	} else {
		data := L.head.data
		head := L.head
		newHead := head.next
		L.head = newHead
		L.head.prev = nil
		L.count--
		head = nil
		return data
	}
}

// RemoveLast remove and return tail of the list
func (L *LinkedList[T]) RemoveLast() T {
	if L.tail != nil {
		tail := L.tail
		L.tail.prev.next = nil
		L.tail = nil
		return tail.data
	} else if L.head != nil {
		data := L.head.data
		L.head = nil
		return data
	}
	var result T
	return result
}

// Remove node by reference to the node
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

func (L *LinkedList[T]) IsEmpty() bool {
	if L.head == nil {
		return true
	}
	return false
}

//func (L LinkedList) Reverse() LinkedList {
//
//}

func (L *LinkedList[T]) Size() int {
	return L.count
}

func (L *LinkedList[T]) Print() {
	list := L.head
	for list != nil {
		//fmt.Println(list.data)
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
	if other == nil {
		return
	}

	if other.head == nil {
		return
	}
	head := other.Head()
	for head.next != nil {
		L.Add(head.data)
		head = head.next
	}

}

func (L *LinkedList[T]) Append(other *LinkedList[T]) {
	L.tail.next = other.head
}
