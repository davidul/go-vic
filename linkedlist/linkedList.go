package linkedlist

// DoublyLinkedList Simple linked list structure.
// Pointer to head and tail of the list.
// Count is the number of node.
type DoublyLinkedList[T comparable] struct {
	head  *Node[T]
	tail  *Node[T]
	count int
}

// Node of the linked list
// Next and previous node.
type Node[T comparable] struct {
	next *Node[T]
	prev *Node[T]
	//list *DoublyLinkedList
	data T
}

func EmptyDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head:  nil,
		tail:  nil,
		count: 0,
	}
}

// NewDoublyLinkedList create a new linked list with data
// head.next -> tail
// head <- tail.prev
func NewDoublyLinkedList[T comparable](data T) *DoublyLinkedList[T] {
	l := DoublyLinkedList[T]{}
	n := Node[T]{}
	l.head = &n
	l.tail = &n
	l.head.data = data
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
func (L *DoublyLinkedList[T]) Head() *Node[T] {
	return L.head
}

// Add append at the end of the list
// H == T
// head -> tail
func (L *DoublyLinkedList[T]) Add(data T) *Node[T] {
	newNode := &Node[T]{
		next: nil,
		prev: nil,
		data: data}

	if L.head == nil {
		L.head = newNode
		L.tail = newNode
	} else {

		// A N -> T
		newNode.prev = L.tail
		L.tail.next = newNode
		L.tail = newNode
	}
	L.count++
	return newNode
}

// AddLast append to tail
func (L *DoublyLinkedList[T]) AddLast(data T) {
	L.Add(data)
}

// AddFirst prepend to head. New head is being created.
func (L *DoublyLinkedList[T]) AddFirst(data T) {
	newNode := &Node[T]{
		data: data,
	}

	if L.head == nil {
		L.head = newNode
		L.tail = newNode
	} else {
		newNode.next = L.head
		L.head.prev = newNode
		L.head = newNode
	}
	L.count++
}

// Peek returns but does not remove the head element
func (L *DoublyLinkedList[T]) Peek() any {
	if L.head != nil {
		return L.head.data
	} else {
		return nil
	}
}

// PeekLast returns but does not remove the tail element
func (L *DoublyLinkedList[T]) PeekLast() T {
	if L.tail != nil {
		return L.tail.data
	} else {
		var result T
		return result
	}
}

// Remove retrieves and removes the head
func (L *DoublyLinkedList[T]) Remove() T {
	//head -> tail

	if L.head == nil {
		var result T
		return result
	}

	data := L.head.data
	L.head = L.head.next
	if L.head == nil {
		L.tail = nil
	} else {
		L.head.prev = nil
	}
	L.count--
	return data
}

// RemoveLast remove and return tail of the list
func (L *DoublyLinkedList[T]) RemoveLast() T {
	if L.tail == nil {
		var zero T
		return zero
	}

	data := L.tail.data
	L.tail = L.tail.prev
	if L.tail == nil {
		L.head = nil
	} else {
		L.tail.next = nil
	}
	L.count--

	return data
}

// Remove node by reference to the node
func (L *DoublyLinkedList[T]) RemoveNode(node *Node[T]) {
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
func (L *DoublyLinkedList[T]) Poll() T {
	return L.Remove()
}

// Converts linked list to array
func (L *DoublyLinkedList[T]) ToArray() []T {
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

func (L *DoublyLinkedList[T]) Contains(v T) bool {
	for e := L.head; e != nil; e = e.Next() {
		if e.Data() == v {
			return true
		}
	}
	return false
}

func (L *DoublyLinkedList[T]) IsEmpty() bool {
	if L.head == nil {
		return true
	}
	return false
}

//func (L DoublyLinkedList) Reverse() DoublyLinkedList {
//
//}

func (L *DoublyLinkedList[T]) Size() int {
	return L.count
}

func (L *DoublyLinkedList[T]) Print() {
	list := L.head
	for list != nil {
		//fmt.Println(list.data)
		list = list.next
	}
}

func (L *DoublyLinkedList[T]) Compare(other *DoublyLinkedList[T]) bool {
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

func (L *DoublyLinkedList[T]) AddAll(other *DoublyLinkedList[T]) {
	if other == nil || other.head == nil {
		return
	}

	if L.head == nil {
		L.head = other.head
		L.tail = other.tail
		L.count = other.count
		return
	}

	L.tail.next = other.head
	other.head.prev = L.tail
	L.tail = other.tail
	L.count += other.count

}

func (L *DoublyLinkedList[T]) Append(other *DoublyLinkedList[T]) {
	L.tail.next = other.head
}
