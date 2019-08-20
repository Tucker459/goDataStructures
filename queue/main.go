package main

import (
	"fmt"

	"github.com/Tucker459/goDataStructures/queue/arrayqueue"

	"github.com/Tucker459/goDataStructures/queue/linkedlistqueue"
)

func main() {
	// String Linked List Queue
	var myLinkedListQueue linkedlistqueue.StringLinkedListQueue

	myLinkedListQueue.Enqueue("Hi")
	headNode := myLinkedListQueue.Front()
	headNode.PrintNode() // "Hi"
	fmt.Print("End Node: ")
	endNode := myLinkedListQueue.End()
	endNode.PrintNode() // "End Node: Hi"
	myLinkedListQueue.Enqueue("Bye")
	myLinkedListQueue.String() // Bye --> Hi
	myLinkedListQueue.Enqueue("Hello")
	headNode = myLinkedListQueue.Front()
	headNode.PrintNode()                     // "Hi"
	fmt.Println(myLinkedListQueue.Size())    // 3
	myLinkedListQueue.String()               // Hello --> Bye --> Hi
	fmt.Println(myLinkedListQueue.IsEmpty()) // False
	removedNode := myLinkedListQueue.Dequeue()
	removedNode.PrintNode() // Hi
	headNode = myLinkedListQueue.Front()
	fmt.Print("New Front Node: ")
	headNode.PrintNode() // "Bye"
	endNode = myLinkedListQueue.End()
	fmt.Print("End Node: ")
	endNode.PrintNode()                   // New End Node: Hello
	myLinkedListQueue.String()            // Hello --> Bye
	fmt.Println(myLinkedListQueue.Size()) // 2
	myLinkedListQueue.Dequeue()
	myLinkedListQueue.String() // Hello
	myLinkedListQueue.Dequeue()
	myLinkedListQueue.String() // Queue is empty!

	fmt.Println("------------------------ Array Queue ------------------------------------")
	// String Array Queue
	var myArrayQueue arrayqueue.StringQueue

	myArrayQueue.Enqueue("Hi - Array Queue")
	myArrayQueue.Enqueue("Hello")
	fmt.Println(myArrayQueue.Size())    // 2
	fmt.Println(myArrayQueue.IsEmpty()) // False
	frontNode := myArrayQueue.Front()
	fmt.Println(*frontNode) // Hi - Array Queue
	removedArrayNode := myArrayQueue.Dequeue()
	fmt.Println(*removedArrayNode)   // Hi - Array Queue
	fmt.Println(myArrayQueue.Size()) // 1

}
