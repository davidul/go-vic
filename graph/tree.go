package graph

import (
	"fmt"
	"github.com/davidul/go-vic/linkedlist"
)

type TreeNode struct {
	Value    any
	children *linkedlist.LinkedList
}

type Tree struct {
	root TreeNode
}

func NewTree(value any) *Tree {
	return &Tree{
		root: TreeNode{
			Value:    value,
			children: &linkedlist.LinkedList{},
		},
	}
}

func (T *Tree) Add(value any) *TreeNode {
	return T.root.Add(value)
}

func (T *Tree) Root() *TreeNode {
	return &T.root
}

func (TN *TreeNode) Add(value any) *TreeNode {
	tn := TreeNode{
		Value:    value,
		children: &linkedlist.LinkedList{},
	}
	TN.children.Add(tn)
	return &tn
}

func (T *Tree) PostOrder() {
	T.postOrder(T.root.children)
	fmt.Println(T.root.Value)
}

func (T *Tree) postOrder(nodes *linkedlist.LinkedList) {
	for x := nodes.Head(); x != nil; x = x.Next() {
		T.postOrder(x.Data().(TreeNode).children)
		fmt.Println(x.Data().(TreeNode).Value)
	}
}

func (T *Tree) PreOrder() {
	fmt.Println(T.root.Value)
	T.preOrder(T.root.children, 1)
}

func (T *Tree) preOrder(nodes *linkedlist.LinkedList, lvl int) {
	for x := nodes.Head(); x != nil; x = x.Next() {
		fmt.Println(x.Data().(TreeNode).Value)
		lvl += 1
		T.preOrder(x.Data().(TreeNode).children, lvl)
	}
}

func Compare(t1 *Tree, t2 Tree) {
	/*tr1 := t1.root.value
	tr2 := t2.root.value
	if tr1 == tr2 {
		t1.root.children
		t2.root.children
	}*/
}

func (T *Tree) Walk() {

}
