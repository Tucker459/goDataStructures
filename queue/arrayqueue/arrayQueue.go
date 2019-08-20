package arrayqueue

import "sync"

// StringQueue is an array-based queue of strings
type StringQueue struct {
	Items []string
	lock  sync.RWMutex
}

// New is the constructor for creating a queue of strings
func (s *StringQueue) New() *StringQueue {
	s.Items = []string{}
	return s
}

// Enqueue adds an String to the end of the queue
func (s *StringQueue) Enqueue(str string) {
	s.lock.Lock()
	s.Items = append(s.Items, str)
	s.lock.Unlock()
}

// Dequeue removes an item from the start of the queue
func (s *StringQueue) Dequeue(str string) *string {
	s.lock.Lock()
	removedStr := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	s.lock.Unlock()
	return &removedStr
}

// Front returns the item next in the queue without removing it
func (s *StringQueue) Front() *string {
	s.lock.RLock()
	firstItem := s.Items[0]
	s.lock.RUnlock()
	return &firstItem
}

// IsEmpty returns false if the queue is not empty
func (s *StringQueue) IsEmpty() bool {
	return len(s.Items) == 0
}

// Size is the current size of the queue
func (s *StringQueue) Size() int {
	return len(s.Items)
}
