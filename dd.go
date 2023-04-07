package main

import (
	"fmt"
	"github.com/davidul/go-vic/graph/intgraph"
	linkedlist "github.com/davidul/go-vic/linkedlist"
)

func main() {
	list := linkedlist.LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.RemoveLast()
	list.Print()
	fmt.Println("=======Peek first/last========")
	fmt.Println(list.Peek())
	fmt.Println(list.PeekLast())
	fmt.Println("=======AddFirst -4- ========")
	list.AddFirst(4)
	list.Print()
	fmt.Println("=====AddFirst -5- ==========")
	list.AddFirst(5)
	list.Print()
	array := list.ToArray()
	for i := range array {
		fmt.Print(i)
		fmt.Println(array[i])
	}

	fmt.Println("================")
	fmt.Printf("Removed %d \n", list.Remove())
	fmt.Println("Print list")
	list.Print()

	fmt.Println("================")
	for h := list.Head(); h != nil; h = h.Next() {
		fmt.Printf("Value -> %d \n", h.Data())
	}

	intGraph := intgraph.NewIntGraph()
	one := intGraph.Add(1)
	intGraph.AddEdge(one, 2)
	bsf := intGraph.Bsf(1, 2)
	fmt.Printf("Bsf search for 2, found %d \n", bsf)
	bsf = intGraph.Bsf(1, 3)
	fmt.Printf("Bsf search for 3, found %d \n", bsf)

	fmt.Println("\033[C \033[C Cursor right")
}
