package main

import (
	"fmt"

	"github.com/Tucker459/goDataStructures/linkedmemory/linkedlist"
)

func main() {
	firstNode := linkedlist.Node{Content: "Hi", Next: nil}

	myList := linkedlist.ItemLinkedList{Head: &firstNode, Size: 1}

	myList.String()
	empty := myList.IsEmpty()
	fmt.Println(empty)
	size := myList.SizeOf()
	fmt.Println(size)
	myList.Append("Bye")
	size = myList.SizeOf()
	fmt.Println(size)
	indexNum := myList.IndexOf("Bye")
	fmt.Println(indexNum)
	headNode := myList.HeadOf()
	fmt.Println(headNode.Content)
	removedNode, err := myList.RemoveAt(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*removedNode)

}
