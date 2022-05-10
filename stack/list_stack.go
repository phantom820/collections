package stack

import (
	"collections/iterator"
	"collections/list"
	"collections/types"
)

// ListStack singly linked list based implementation of a stack.
type ListStack[T types.Equitable[T]] struct {
	list *list.ForwardList[T]
}

// NewListStack creates a stack with the specified elements, if there are none an empty stack is created.
func NewListStack[T types.Equitable[T]](elements ...T) *ListStack[T] {
	var s ListStack[T] = ListStack[T]{list: list.NewForwardList[T]()}
	s.AddSlice(elements)
	return &s
}

// Peek returns the element at the top of the stack s without removing it. Will panic if stack has no top element to peek.
func (s *ListStack[T]) Peek() T {
	defer func() {
		if r := recover(); r != nil {
			panic(NoTopElementError)
		}
	}()
	return s.list.Front()
}

// Pop returns and removes the element at the top of the stack s. Will panic if stack has no top element to pop.
func (s *ListStack[T]) Pop() T {
	defer func() {
		if r := recover(); r != nil {
			panic(NoTopElementError)
		}
	}()
	return s.list.RemoveFront()
}

// Add pushes element e to the stack s.
func (s *ListStack[T]) Add(e T) bool {
	s.list.AddFront(e)
	return true
}

// AddAll pushes all the element from an iterable elements to the stack s.
func (s *ListStack[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		s.list.AddFront(it.Next())
	}
}

// AddSlices pushes all the elements from a slice to the stack s.
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

// Contains checks if element e is part of the stack s.
func (s *ListStack[T]) Contains(e T) bool {
	return s.list.Contains(e)
}

// Empty checks if the stack s is empty.
func (s *ListStack[T]) Empty() bool {
	return s.list.Empty()
}

// Iterator returns in iterator for iterating through the stack s.
func (s *ListStack[T]) Iterator() iterator.Iterator[T] {
	return s.list.Iterator()
}

// Len returns the size of the stack s.
func (s *ListStack[T]) Len() int {
	return s.list.Len()
}

// Remove removes the first occurence of element e from the stack s.
func (s *ListStack[T]) Remove(e T) bool {
	return s.list.Remove(e)
}

// RemoveAll removes all elements from the stack s that occur in iterable elements.
func (s *ListStack[T]) RemoveAll(elements iterator.Iterable[T]) {
	s.list.RemoveAll(elements)
}
