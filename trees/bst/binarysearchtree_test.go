package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//	   5
//	4     6
//
// 3         7
func TestBinarySearchTree_Add(t *testing.T) {
	bst := NewBST[int]()
	bst.Add(5, 5)
	bst.Add(4, 4)
	bst.Add(3, 3)
	bst.Add(6, 6)
	bst.Add(7, 7)
	bst.Add(5, 5)
	h := Height{height: 0}
	balanced, height := bst.isHeightBalanced(bst.root, h)
	assert.Equal(t, balanced, true)
	assert.Equal(t, height.height, 3)
}

// 1
//
//	2
//	  3
func TestBinarySearchTree_Balance(t *testing.T) {
	bst := NewBST[int]()
	bst.Add(1, 1)
	bst.Add(2, 2)
	bst.Add(3, 3)
	h := Height{height: 0}
	balanced, height := bst.isHeightBalanced(bst.root, h)
	assert.Equal(t, balanced, false)
	assert.Equal(t, height.height, 3)
}

func TestBinarySearchTree_Traverse(t *testing.T) {
	bst := NewBST[int]()
	bst.Add(5, 5)
	bst.Add(4, 4)
	bst.Add(3, 3)
	bst.Add(6, 6)
	bst.Add(7, 7)
	bst.Add(5, 5)
	bst.InOrderTraverse()
}
