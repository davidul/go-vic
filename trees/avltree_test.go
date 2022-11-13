package trees

import (
	"fmt"
	"testing"
)

func TestNewEmptyAVLTree(t *testing.T) {
	tree := NewEmptyAVLTree()
	fmt.Println(tree.root)
}

func TestAVLTree_Add(t *testing.T) {
	tree := NewEmptyAVLTree()
	tree.Add(2)
	tree.Add(3)
	tree.Add(4)
	tree.Add(1)
	tree.Add(5)
	tree.PreOrder(tree.root)
	tree.PrintTree(tree.root, "", false)
}
