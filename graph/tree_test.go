package graph

import "testing"

func TestTree_NewTree(t *testing.T) {
	tree := NewTree("1")
	add := tree.root.Add("2")
	tree.root.Add("3")
	add.Add("4")
	tree.InOrder()
}

/*
    1
  2    3
4 5 6
-> 4 5 6 2 3 1
*/
func TestTree_InOrder(t *testing.T) {
	tree := NewTree("1")
	two := tree.root.Add("2")
	tree.root.Add("3")
	two.Add("4")
	two.Add("5")
	two.Add("6")
	tree.InOrder()
}

/*
     1
  2     3
      4 5 6
-> 2 1 3 4 5 6
*/
func TestTree_InOrder_2(t *testing.T) {
	tree := NewTree("1")
	tree.root.Add("2")
	three := tree.root.Add("3")
	three.Add("4")
	three.Add("5")
	three.Add("6")
	tree.InOrder()
}