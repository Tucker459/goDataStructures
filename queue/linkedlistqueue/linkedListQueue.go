package linkedlistqueue

import (
	"fmt"
	"sync"
)

// Node is the representation of a single item in the double
// linked list
type Node struct {
	previous *Node
	content  string
	next     *Node
}

// StringLinkedListQueue is string based
// double linked list queue
type StringLinkedListQueue struct {
	head *Node
	tail *Node
	size int
	lock sync.RWMutex
}

// newNode is constructor for creating a Node
func newNode(str string) *Node {
	return &Node{previous: nil, content: str, next: nil}
}

// PrintNode returns the content of that node
func (n *Node) PrintNode() {
	fmt.Println(n.content)
}

// Enqueue adds an Item to the beginning of the queue
func (s *StringLinkedListQueue) Enqueue(str string) {
	s.lock.Lock()
	node := newNode(str)
	if s.head == nil {
		s.head = node
		s.tail = node
		s.size++
	} else {
		node.next = s.head
		s.head.previous = node
		s.head = node
		s.size++
	}
	s.lock.Unlock()
}

// Dequeue removes the node from the end of the queue
func (s *StringLinkedListQueue) Dequeue() *Node {
	s.lock.Lock()
	defer s.lock.Unlock()
	var removedNode *Node
	if s.tail == nil {
		return nil
	} else if s.tail == s.head {
		removedNode = s.head
		s.head = nil
		s.tail = nil
		s.size--
		return removedNode
	}
	removedNode = s.tail
	newTailNode := removedNode.previous
	newTailNode.next = nil
	s.tail = newTailNode
	removedNode.previous = nil
	s.size--

	return removedNode
}

// Front returns the node that will we popped off next
func (s *StringLinkedListQueue) Front() *Node {
	return s.tail
}

// End returns the node at the end of the queue
func (s *StringLinkedListQueue) End() *Node {
	return s.head
}

// IsEmpty returns false if the node is not empty
func (s *StringLinkedListQueue) IsEmpty() bool {
	return s.head == nil
}

// Size is the current size of the queue
func (s *StringLinkedListQueue) Size() int {
	return s.size
}

// String prints out a string representation of the queue
func (s *StringLinkedListQueue) String() {
	s.lock.RLock()
	if s.size == 0 {
		fmt.Println("Queue is empty!")
	}
	j := 0
	node := s.head
	for j < s.size {
		fmt.Print(node.content)
		if j != s.size-1 {
			fmt.Print("-->")
		} else {
			fmt.Println("")
		}
		node = node.next
		j++
	}
	s.lock.RUnlock()
}
