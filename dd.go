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
	fmt.Println("===============")
	fmt.Println(list.Peek())
	fmt.Println(list.PeekLast())
	fmt.Println("===============")
	list.AddFirst(4)
	list.Print()
	fmt.Println("===============")
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

}
