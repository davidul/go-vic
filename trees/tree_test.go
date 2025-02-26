package trees

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hash/fnv"
	"strconv"
	"testing"
)

/*
		1
		\ 	\
		2  	 3
	    \
		4
*/
func TestTree_NewTree(t *testing.T) {
	tree := NewTree("1")
	add := tree.root.Add("2")
	tree.root.Add("3")
	add.Add("4")
	hasher := fnv.New32a()
	hasher.Write([]byte(tree.String()))
	sum32 := hasher.Sum32()
	fmt.Println(sum32)
	order := tree.PostOrder()
	assert.Len(t, order, 4)
}

func TestTree_Large(t *testing.T) {
	tree := NewTree("/")
	for i := 0; i < 1000000; i++ {
		tree.Add(strconv.Itoa(i))
	}
	order := tree.PostOrder()
	assert.Len(t, order, 1000001)
}

/*
	  1
	2    3

4 5 6
-> 4 5 6 2 3 1
*/
func TestTree_PreOrder(t *testing.T) {
	tree := NewTree("1")
	two := tree.root.Add("2")
	tree.root.Add("3")
	two.Add("4")
	two.Add("5")
	two.Add("6")
	tree.PreOrder()
}

/*
	   1
	2     3
	    4 5 6

-> 2 1 3 4 5 6
*/
func TestTree_PreOrder2(t *testing.T) {
	tree := NewTree("1")
	tree.root.Add("2")
	three := tree.root.Add("3")
	three.Add("4")
	three.Add("5")
	three.Add("6")
	tree.PreOrder()
}

func TestTree_PostOrder(t *testing.T) {
	tree := NewTree("1")
	tree.root.Add("2")
	three := tree.root.Add("3")
	three.Add("4")
	three.Add("5")
	three.Add("6")
	order := tree.PostOrder()
	assert.Len(t, order, 6)
}

/*
	1

2       3

	4 5 6
*/
func TestTree_PreOrderDepth(t *testing.T) {
	tree := NewTree("1")
	tree.root.Add("2")
	three := tree.root.Add("3")
	three.Add("4")
	three.Add("5")
	three.Add("6")
	tree.PreOrderDepth(2)
}

func TestTree_Bsf(t *testing.T) {
	tree := NewTree("1")
	tree.root.Add("2")
	three := tree.root.Add("3")
	three.Add("4")
	three.Add("5")
	three.Add("6")

	bsf := tree.Bsf("6")
	fmt.Println(bsf.Value)
}
