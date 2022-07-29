package main

import (
	"data-structures-go/linkedlist"
	"fmt"
)

func main() {
	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Print()

	fmt.Println(list.Peek())
	fmt.Println(list.PeekLast())
}
