package bst

type BinarySearchTree[T comparable] struct {
	root *BSTNode[T]
	size int
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
	return new(BinarySearchTree[T])
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

func (B *BinarySearchTree[T]) InOrderTraverse() []T {
	var result []T
	B.inorder(B.root, &result)
	return result
}

func (b *BinarySearchTree[T]) inorder(node *BSTNode[T], result *[]T) {
	if node != nil {
		b.inorder(node.left, result)
		*result = append(*result, node.value)
		b.inorder(node.right, result)
	}
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

func (b *BinarySearchTree[T]) Delete(key int32) bool {
	if b.root == nil {
		return false
	}

	var found bool
	b.root, found = b.deleteRecursive(b.root, key)
	if found {
		b.size--
	}
	return found
}

func (b *BinarySearchTree[T]) deleteRecursive(node *BSTNode[T], key int32) (*BSTNode[T], bool) {
	if node == nil {
		return nil, false
	}

	if key < node.key {
		var found bool
		node.left, found = b.deleteRecursive(node.left, key)
		return node, found
	} else if key > node.key {
		var found bool
		node.right, found = b.deleteRecursive(node.right, key)
		return node, found
	}

	// Node to delete found
	if node.left == nil {
		return node.right, true
	} else if node.right == nil {
		return node.left, true
	}

	// Node has two children
	minNode := b.findMin(node.right)
	node.key = minNode.key
	node.value = minNode.value
	node.right, _ = b.deleteRecursive(node.right, minNode.key)
	return node, true
}

func (b *BinarySearchTree[T]) findMin(node *BSTNode[T]) *BSTNode[T] {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}
