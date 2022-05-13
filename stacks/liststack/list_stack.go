package liststack

import (
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/stacks"
	"github.com/phantom820/collections/types"
)

// ListStack singly linked list based implementation of a stack.
type ListStack[T types.Equitable[T]] struct {
	list *forwardlist.ForwardList[T]
}

// New creates a stack with the specified elements, if there are none an empty stack is created.
func New[T types.Equitable[T]](elements ...T) *ListStack[T] {
	var s ListStack[T] = ListStack[T]{list: forwardlist.New[T]()}
	s.AddSlice(elements)
	return &s
}

// Peek returns the element at the top of the stack without removing it. Will panic if stack has no top element to peek.
func (s *ListStack[T]) Peek() T {
	defer func() {
		if r := recover(); r != nil {
			panic(stacks.ErrNoTopElement)
		}
	}()
	return s.list.Front()
}

// Pop returns and removes the element at the top of the stack. Will panic if stack has no top element to pop.
func (s *ListStack[T]) Pop() T {
	defer func() {
		if r := recover(); r != nil {
			panic(stacks.ErrNoTopElement)
		}
	}()
	return s.list.RemoveFront()
}

// Add pushes elements  to the stack.
func (s *ListStack[T]) Add(elements ...T) bool {
	for _, e := range elements {
		s.list.AddFront(e)
	}
	return true
}

// AddAll pushes all the element from an iterable elements to the stack.
func (s *ListStack[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		s.list.AddFront(it.Next())
	}
}

// AddSlices pushes all the elements from a slice to the stack.
func (s *ListStack[T]) AddSlice(slice []T) {
	for _, e := range slice {
		s.list.AddFront(e)
	}
}

// Clear removes all items on the stack.
func (s *ListStack[T]) Clear() {
	s.list.Clear()
}

// Collect returns the stack as a slice.
func (s *ListStack[T]) Collect() []T {
	return s.list.Collect()
}

// Contains checks if element e is part of the stack.
func (s *ListStack[T]) Contains(e T) bool {
	return s.list.Contains(e)
}

// Empty checks if the stack is empty.
func (s *ListStack[T]) Empty() bool {
	return s.list.Empty()
}

// Iterator returns in iterator for iterating through the stack.
func (s *ListStack[T]) Iterator() iterator.Iterator[T] {
	return s.list.Iterator()
}

// Len returns the size of the stack.
func (s *ListStack[T]) Len() int {
	return s.list.Len()
}

// Remove removes the first occurence of element e from the stack.
func (s *ListStack[T]) Remove(e T) bool {
	return s.list.Remove(e)
}

// RemoveAll removes all elements from the stack that occur in iterable elements.
func (s *ListStack[T]) RemoveAll(elements iterator.Iterable[T]) {
	s.list.RemoveAll(elements)
}
