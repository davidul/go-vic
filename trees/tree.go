package trees

import (
	"fmt"
	"github.com/davidul/go-vic/linkedlist"
	"math"
	"strings"
)

// TreeNode node in the tree holds value T
// and list to its children of type T
type TreeNode[T comparable] struct {
	Value    T
	Children *linkedlist.DoublyLinkedList[TreeNode[T]]
}

// Tree TreeNode holds reference to root
type Tree[T comparable] struct {
	root TreeNode[T]
}

func NewTree[T comparable](value T) *Tree[T] {
	return &Tree[T]{
		root: TreeNode[T]{
			Value:    value,
			Children: &linkedlist.DoublyLinkedList[TreeNode[T]]{},
		},
	}
}

func (R *Tree[T]) Add(value T) *TreeNode[T] {
	return R.root.Add(value)
}

func (R *Tree[T]) Root() *TreeNode[T] {
	return &R.root
}

// Add new TreeNode to the tree.
func (TN *TreeNode[T]) Add(value T) *TreeNode[T] {
	tn := TreeNode[T]{
		Value:    value,
		Children: &linkedlist.DoublyLinkedList[TreeNode[T]]{},
	}
	TN.Children.Add(tn)
	return &tn
}

func (TN *TreeNode[T]) Delete() {

}

// Bsf - Breadth First Search on the tree
func (R *Tree[T]) Bsf(goal T) *TreeNode[T] {
	visited := linkedlist.DoublyLinkedList[TreeNode[T]]{}
	queue := linkedlist.DoublyLinkedList[TreeNode[T]]{}
	queue.Add(R.root)

	for !queue.IsEmpty() {
		v := queue.Poll()
		visited.Add(v)
		vv := v.Value
		if vv == goal {
			treeNode := v
			return &treeNode
		}
		node := v
		list := node.Children
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e.Data()) {
				queue.Add(e.Data())
			}
		}
	}
	return nil
}

func (R *Tree[T]) PostOrder() []T {
	tt := make([]T, 0)
	tt = R.postOrder(R.root.Children, tt)
	tt = append(tt, R.root.Value)
	return tt
}

func (R *Tree[T]) postOrder(nodes *linkedlist.DoublyLinkedList[TreeNode[T]], collector []T) []T {
	if nodes == nil {
		return collector
	}

	if nodes.Head() == nil {
		return collector
	}
	for x := nodes.Head(); x != nil; x = x.Next() {
		collector = R.postOrder(x.Data().Children, collector)
		collector = append(collector, x.Data().Value)
	}
	return collector
}

func (R *Tree[T]) PreOrder() {
	fmt.Println(R.root.Value)
	R.preOrder(R.root.Children, 1, math.MaxInt)
}

func (R *Tree[T]) preOrder(nodes *linkedlist.DoublyLinkedList[TreeNode[T]], lvl int, stopLvl int) {
	if nodes == nil {
		return
	}

	lvl += 1
	if stopLvl <= lvl {
		return
	}
	for x := nodes.Head(); x != nil; x = x.Next() {
		fmt.Println(x.Data().Value)
		R.preOrder(x.Data().Children, lvl, stopLvl)
	}
}

func (R *Tree[T]) PreOrderDepth(depth int) {
	fmt.Println(R.root.Value)
	R.preOrder(R.root.Children, 0, depth)
}

func (R *Tree[T]) PreOrderDepthFunc(depth int, f func(any)) {
	f(R.root.Value)
	R.preOrderFunc(R.root.Children, 0, depth, f)
}

func (R *Tree[T]) preOrderFunc(nodes *linkedlist.DoublyLinkedList[TreeNode[T]], lvl int, stopLvl int, f func(any)) {
	lvl += 1
	if stopLvl <= lvl {
		return
	}
	for x := nodes.Head(); x != nil; x = x.Next() {
		f(x.Data().Value)
		R.preOrder(x.Data().Children, lvl, stopLvl)
	}
}

func (R *Tree[T]) String() string {
	order := R.PostOrder()
	s := make([]string, len(order))
	for _, i := range order {
		s = append(s, fmt.Sprintf("%v", i))
	}
	return strings.Join(s, "")
}

//func Compare(t1 *Tree, t2 Tree) {
/*tr1 := t1.root.value
tr2 := t2.root.value
if tr1 == tr2 {
	t1.root.Children
	t2.root.Children
}*/
//}

func (R *Tree[T]) Walk() {

}
