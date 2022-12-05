package trees

import (
	"fmt"
	"github.com/davidul/go-vic/linkedlist"
	"math"
)

type TreeNode[T interface{}] struct {
	Value    any
	Children *linkedlist.LinkedList[T]
}

type Tree[T interface{}] struct {
	root TreeNode[T]
}

func NewTree[T interface{}](value any) *Tree[T] {
	return &Tree[T]{
		root: TreeNode[T]{
			Value:    value,
			Children: &linkedlist.LinkedList[T]{},
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
func (TN *TreeNode[T]) Add(value any) *TreeNode[T] {
	tn := TreeNode[T]{
		Value:    value,
		Children: &linkedlist.LinkedList[T]{},
	}
	TN.Children.Add(tn)
	return &tn
}

func (TN *TreeNode[T]) Delete() {

}

// Bsf - Breadth First Search on the tree
func (R *Tree[T]) Bsf(goal any) *TreeNode[T] {
	visited := linkedlist.LinkedList[T]{}
	queue := linkedlist.LinkedList[T]{}
	queue.Add(R.root)

	for !queue.IsEmpty() {
		v := queue.Poll()
		visited.Add(v)
		vv := v.(TreeNode[T]).Value
		if vv == goal {
			treeNode := v.(TreeNode[T])
			return &treeNode
		}
		node := v.(TreeNode[T])
		list := node.Children
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e.Data()) {
				queue.Add(e.Data())
			}
		}
	}
	return nil
}

func (R *Tree[T]) PostOrder() {
	R.postOrder(R.root.Children)
	fmt.Println(R.root.Value)
}

func (R *Tree[T]) postOrder(nodes *linkedlist.LinkedList[T]) {
	for x := nodes.Head(); x != nil; x = x.Next() {
		R.postOrder(x.Data().(TreeNode[T]).Children)
		fmt.Println(x.Data().(TreeNode[T]).Value)
	}
}

func (R *Tree[T]) PreOrder() {
	fmt.Println(R.root.Value)
	R.preOrder(R.root.Children, 1, math.MaxInt)
}

func (R *Tree[T]) preOrder(nodes *linkedlist.LinkedList[T], lvl int, stopLvl int) {
	lvl += 1
	if stopLvl <= lvl {
		return
	}
	for x := nodes.Head(); x != nil; x = x.Next() {
		fmt.Println(x.Data().(TreeNode[T]).Value)
		R.preOrder(x.Data().(TreeNode[T]).Children, lvl, stopLvl)
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

func (R *Tree[T]) preOrderFunc(nodes *linkedlist.LinkedList[T], lvl int, stopLvl int, f func(any)) {
	lvl += 1
	if stopLvl <= lvl {
		return
	}
	for x := nodes.Head(); x != nil; x = x.Next() {
		f(x.Data().(TreeNode[T]).Value)
		R.preOrder(x.Data().(TreeNode[T]).Children, lvl, stopLvl)
	}
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
