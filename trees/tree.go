package trees

import (
	"fmt"
	"github.com/davidul/go-vic/linkedlist"
	"math"
)

type TreeNode struct {
	Value    any
	Children *linkedlist.LinkedList
}

type Tree struct {
	root TreeNode
}

func NewTree(value any) *Tree {
	return &Tree{
		root: TreeNode{
			Value:    value,
			Children: &linkedlist.LinkedList{},
		},
	}
}

func (T *Tree) Add(value any) *TreeNode {
	return T.root.Add(value)
}

func (T *Tree) Root() *TreeNode {
	return &T.root
}

// Add new TreeNode to the tree.
func (TN *TreeNode) Add(value any) *TreeNode {
	tn := TreeNode{
		Value:    value,
		Children: &linkedlist.LinkedList{},
	}
	TN.Children.Add(tn)
	return &tn
}

func (TN *TreeNode) Delete() {

}

// Bsf - Breadth First Search on the tree
func (T *Tree) Bsf(goal any) *TreeNode {
	visited := linkedlist.LinkedList{}
	queue := linkedlist.LinkedList{}
	queue.Add(T.root)

	for !queue.IsEmpty() {
		v := queue.Poll()
		visited.Add(v)
		vv := v.(TreeNode).Value
		if vv == goal {
			treeNode := v.(TreeNode)
			return &treeNode
		}
		node := v.(TreeNode)
		list := node.Children
		for e := list.Head(); e != nil; e = e.Next() {
			if !visited.Contains(e.Data()) {
				queue.Add(e.Data())
			}
		}
	}
	return nil
}

func (T *Tree) PostOrder() {
	T.postOrder(T.root.Children)
	fmt.Println(T.root.Value)
}

func (T *Tree) postOrder(nodes *linkedlist.LinkedList) {
	for x := nodes.Head(); x != nil; x = x.Next() {
		T.postOrder(x.Data().(TreeNode).Children)
		fmt.Println(x.Data().(TreeNode).Value)
	}
}

func (T *Tree) PreOrder() {
	fmt.Println(T.root.Value)
	T.preOrder(T.root.Children, 1, math.MaxInt)
}

func (T *Tree) preOrder(nodes *linkedlist.LinkedList, lvl int, stopLvl int) {
	lvl += 1
	if stopLvl <= lvl {
		return
	}
	for x := nodes.Head(); x != nil; x = x.Next() {
		fmt.Println(x.Data().(TreeNode).Value)
		T.preOrder(x.Data().(TreeNode).Children, lvl, stopLvl)
	}
}

func (T *Tree) PreOrderDepth(depth int) {
	fmt.Println(T.root.Value)
	T.preOrder(T.root.Children, 0, depth)
}

func (T *Tree) PreOrderDepthFunc(depth int, f func(any)) {
	f(T.root.Value)
	T.preOrderFunc(T.root.Children, 0, depth, f)
}

func (T *Tree) preOrderFunc(nodes *linkedlist.LinkedList, lvl int, stopLvl int, f func(any)) {
	lvl += 1
	if stopLvl <= lvl {
		return
	}
	for x := nodes.Head(); x != nil; x = x.Next() {
		f(x.Data().(TreeNode).Value)
		T.preOrder(x.Data().(TreeNode).Children, lvl, stopLvl)
	}
}

func Compare(t1 *Tree, t2 Tree) {
	/*tr1 := t1.root.value
	tr2 := t2.root.value
	if tr1 == tr2 {
		t1.root.Children
		t2.root.Children
	}*/
}

func (T *Tree) Walk() {

}
