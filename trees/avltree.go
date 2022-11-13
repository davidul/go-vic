package trees

import "fmt"

type AVLTree struct {
	root *AVLNode
}

type AVLNode struct {
	key    int32
	value  any
	height int
	left   *AVLNode
	right  *AVLNode
	parent *AVLNode
}

func NewEmptyAVLTree() *AVLTree {
	tree := new(AVLTree)
	return tree
}

func NewAVLTree(key int32, value any) *AVLTree {
	n := &AVLNode{
		key:    key,
		value:  value,
		parent: nil,
		left:   nil,
		right:  nil,
		height: 0,
	}
	tree := new(AVLTree)
	tree.root = n
	return tree
}

func (A *AVLTree) Add(key int32) {
	A.root = A.insert(key, A.root)
}

func (A *AVLTree) insert(key int32, node *AVLNode) *AVLNode {
	if node == nil {
		node := &AVLNode{
			key:    key,
			value:  nil,
			height: 0,
			left:   nil,
			right:  nil,
			parent: nil,
		}
		return node
	}
	if node.key < key {
		node.right = A.insert(key, node.right)
	} else if node.key > key {
		node.left = A.insert(key, node.left)
	} else {
		return node
	}

	node.height = 1 + max(height(node.left), height(node.right))
	factor := A.balanceFactor(node)
	if factor > 1 {
		if key < node.left.key {
			return A.rightRotate(node)
		} else if key > node.left.key {
			node.left = A.leftRotate(node.left)
			return A.rightRotate(node)
		}
	}

	if factor < -1 {
		if key > node.right.key {
			return A.leftRotate(node)
		} else if key < node.right.key {
			node.right = A.rightRotate(node.right)
			return A.leftRotate(node)
		}
	}
	return node
}

func height(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (A *AVLTree) balanceFactor(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}

func (A *AVLTree) rightRotate(y *AVLNode) *AVLNode {
	x := y.left
	T2 := x.right
	x.right = y
	y.left = T2
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1
	return x
}

func (A *AVLTree) leftRotate(x *AVLNode) *AVLNode {
	y := x.right
	T2 := y.left
	y.left = x
	x.right = T2
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y
}

func (A *AVLTree) height() {

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (A *AVLTree) PrintTree(node *AVLNode, indent string, last bool) {
	if node != nil {
		fmt.Printf(indent)
		if last {
			fmt.Printf("R---")
			indent += "   "
		} else {
			fmt.Printf("L---")
			indent += "|  "
		}
		fmt.Println(node.key)
		A.PrintTree(node.left, indent, false)
		A.PrintTree(node.right, indent, true)
	}
}

func (A *AVLTree) PreOrder(node *AVLNode) {
	if node != nil {
		fmt.Println(node.key)
		A.PreOrder(node.left)
		A.PreOrder(node.right)
	}
}
