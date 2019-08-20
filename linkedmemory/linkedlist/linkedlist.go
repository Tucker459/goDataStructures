package linkedlist

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of LinkedList
type Item generic.Type

// Node is a single node that composes the List
type Node struct {
	Content string
	Next    *Node
}

// ItemLinkedList the linked list of Items
type ItemLinkedList struct {
	Head *Node
	Size int
	lock sync.RWMutex
}

// Append adds an Item to the end of the linkedlist
func (ll *ItemLinkedList) Append(t string) {
	ll.lock.Lock()
	node := Node{t, nil}
	if ll.Head == nil {
		ll.Head = &node
	} else {
		last := ll.Head
		for {
			if last.Next == nil {
				break
			}
			last = last.Next
		}
		last.Next = &node
	}
	ll.Size++
	ll.lock.Unlock()
}

// Insert adds an Item at position i
func (ll *ItemLinkedList) Insert(i int, t string) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if i < 0 || i > ll.Size {
		return fmt.Errorf("Index out of bounds")
	}
	addNode := Node{t, nil}

	if i == 0 {
		addNode.Next = ll.Head
		ll.Head = &addNode
		return nil
	}

	node := ll.Head
	j := 0
	for j < i-2 {
		j++
		node = node.Next
	}
	addNode.Next = node.Next
	node.Next = &addNode
	ll.Size++
	return nil
}

// RemoveAt removes a node at position i
func (ll *ItemLinkedList) RemoveAt(i int) (*string, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if i < 0 || i > ll.Size {
		return nil, fmt.Errorf("Index out of bounds")
	}
	node := ll.Head
	j := 0
	for j < i-1 {
		j++
		node = node.Next
	}
	remove := node.Next
	node.Next = remove.Next
	ll.Size--
	return &remove.Content, nil
}

// IndexOf returns the position of the Item i
func (ll *ItemLinkedList) IndexOf(t string) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.Head
	j := 0
	for {
		if node.Content == t {
			return j
		}
		if node.Next == nil {
			return -1
		}
		node = node.Next
		j++
	}
}

// IsEmpty returns true if linked list is empty
func (ll *ItemLinkedList) IsEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	if ll.Head == nil {
		return true
	}
	return false
}

// SizeOf returns the size of the linked list
func (ll *ItemLinkedList) SizeOf() int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	size := 1
	last := ll.Head
	for {
		if last == nil || last.Next == nil {
			break
		}
		last = last.Next
		size++
	}
	return size
}

// String returns a string representation of the linked list
func (ll *ItemLinkedList) String() {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.Head
	for {
		if node == nil {
			break
		}
		fmt.Print(node.Content)
		fmt.Print(" ")
		node = node.Next
	}
	fmt.Println()
}

// HeadOf returns the first node of the linked list
func (ll *ItemLinkedList) HeadOf() *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	return ll.Head
}
