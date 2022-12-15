package bst

type BinarySearchTree[T comparable] struct {
	root *BSTNode[T]
}

type BSTNode[T comparable] struct {
	value T
	key   int32
	left  *BSTNode[T]
	right *BSTNode[T]
}

type Height struct {
	height int
}

func NewBST[T comparable]() *BinarySearchTree[T] {
	b := new(BinarySearchTree[T])
	return b
}

func (B *BinarySearchTree[T]) Add(key int32, value T) {
	if B.root == nil {
		b := new(BSTNode[T])
		b.value = value
		b.key = key
		B.root = b
	} else {
		B.insert(key, value, B.root)
	}
}

func (B *BinarySearchTree[T]) isHeightBalanced(node *BSTNode[T], height Height) (bool, Height) {
	if node == nil {
		height.height = 0
		return true, height
	}

	leftHeight := Height{height: 0}
	rightHeight := Height{height: 0}

	l, left := B.isHeightBalanced(node.left, leftHeight)
	r, right := B.isHeightBalanced(node.right, rightHeight)

	if left.height > right.height {
		height.height = left.height + 1
	} else {
		height.height = right.height + 1
	}

	if left.height-right.height >= 2 || right.height-left.height >= 2 {
		return false, height
	} else {
		return l && r, height
	}

}

func (B *BinarySearchTree[T]) Traverse() {

}

func (B *BinarySearchTree[T]) find(key int32, node *BSTNode[T]) *BSTNode[T] {
	if node.key > key {
		if node.left != nil {
			return B.find(key, node.left)
		} else {
			return node
		}
	} else if node.key < key {
		if node.right != nil {
			return B.find(key, node.right)
		} else {
			return node
		}
	}

	//equal
	return nil
}

func (B *BinarySearchTree[T]) insert(key int32, value T, node *BSTNode[T]) {
	find := B.find(key, node)
	if find != nil {
		if find.key > key {
			find.left = new(BSTNode[T])
			find.left.key = key
			find.left.value = value
		} else {
			find.right = new(BSTNode[T])
			find.right.key = key
			find.right.value = value
		}

	}
}
