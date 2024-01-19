package main

import (
	"GO_LinkedList/linkedlist"
	"fmt"
)

func main() {
	list := linkedlist.Make()
	list.Push(1)
	list.Push(2)
	list.Append(3)
	list.Put(3, 5)
	//fmt.Println(list.Pop())
	list.Remove(3)
	list.Print()
	fmt.Println(list.Get(0))
}
