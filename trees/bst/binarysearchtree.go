package bst

type BinarySearchTree struct {
	root *BSTNode
}

type BSTNode struct {
	value any
	key   int32
	left  *BSTNode
	right *BSTNode
}

type Height struct {
	height int
}

func NewBST() *BinarySearchTree {
	b := new(BinarySearchTree)
	return b
}

func (B *BinarySearchTree) Add(key int32, value any) {
	if B.root == nil {
		b := new(BSTNode)
		b.value = value
		b.key = key
		B.root = b
	} else {
		B.insert(key, value, B.root)
	}
}

func (B *BinarySearchTree) isHeightBalanced(node *BSTNode, height Height) (bool, Height) {
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

func (B *BinarySearchTree) Traverse() {

}

func (B *BinarySearchTree) find(key int32, node *BSTNode) *BSTNode {
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

func (B *BinarySearchTree) insert(key int32, value any, node *BSTNode) {
	find := B.find(key, node)
	if find != nil {
		if find.key > key {
			find.left = new(BSTNode)
			find.left.key = key
			find.left.value = value
		} else {
			find.right = new(BSTNode)
			find.right.key = key
			find.right.value = value
		}

	}
}
