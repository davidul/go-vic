package bst

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_Add(t *testing.T) {
	bst := NewBST()
	bst.Add(5, 5)
	bst.Add(4, 4)
	bst.Add(3, 3)
	bst.Add(6, 6)
	bst.Add(7, 7)
	bst.Add(5, 5)
	h := Height{height: 0}
	balanced, height := bst.isHeightBalanced(bst.root, h)
	fmt.Println(balanced)
	fmt.Println(height)
}

func TestBinarySearchTree_Balance(t *testing.T) {
	bst := NewBST()
	bst.Add(1, 1)
	bst.Add(2, 2)
	bst.Add(3, 3)
	h := Height{height: 0}
	balanced, height := bst.isHeightBalanced(bst.root, h)
	fmt.Println(balanced)
	fmt.Println(height)
}
