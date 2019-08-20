package main

import (
	"fmt"

	"github.com/Tucker459/goDataStructures/queue/linkedlistqueue"
)

func main() {
	var myLinkedListQueue linkedlistqueue.StringLinkedListQueue

	myLinkedListQueue.Enqueue("Hi")
	headNode := myLinkedListQueue.Front()
	headNode.PrintNode()
	fmt.Print("End Node: ")
	endNode := myLinkedListQueue.End()
	endNode.PrintNode()
	myLinkedListQueue.Enqueue("Bye")
	myLinkedListQueue.String()
	myLinkedListQueue.Enqueue("Hello")
	headNode = myLinkedListQueue.Front()
	headNode.PrintNode()
	fmt.Println(myLinkedListQueue.Size())
	myLinkedListQueue.String()
	removedNode := myLinkedListQueue.Dequeue()
	myLinkedListQueue.String()
	removedNode = myLinkedListQueue.Dequeue()
	removedNode.PrintNode()
	myLinkedListQueue.String()
	removedNode = myLinkedListQueue.Dequeue()
	removedNode.PrintNode()
	myLinkedListQueue.String()

}
