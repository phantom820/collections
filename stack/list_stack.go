// ForwardList based implementation of a stack.
package stack

import (
	"collections/iterator"
	"collections/list"
	"collections/types"
)

// ListStack an interface providing singly linked list based implementation of a stack. This interface serves as an  abstraction for
// operating on an underlying list.
type ListStack[T types.Equitable[T]] interface {
	Stack[T]
}

// listStack concrete type for a singly linked list based stack.
type listStack[T types.Equitable[T]] struct {
	list list.ForwardList[T]
}

// NewListStack creates an empty list based stack.
func NewListStack[T types.Equitable[T]]() ListStack[T] {
	var s listStack[T] = listStack[T]{list: list.NewForwardList[T]()}
	return &s
}

// Peek returns the element at the top of the stack s without removing it. Will panic if stack has no top element to peek.
func (s *listStack[T]) Peek() T {
	defer func() {
		if r := recover(); r != nil {
			panic(NoTopElementError)
		}
	}()
	return s.list.Front()
}

// Pop returns and removes the element at the top of the stack s. Will panic if stack has no top element to pop.
func (s *listStack[T]) Pop() T {
	defer func() {
		if r := recover(); r != nil {
			panic(NoTopElementError)
		}
	}()
	return s.list.RemoveFront()
}

// Add pushes element e to the stack s.
func (s *listStack[T]) Add(e T) bool {
	s.list.AddFront(e)
	return true
}

// AddAll pushes all the element from an iterable elements to the stack s.
func (s *listStack[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		s.list.AddFront(it.Next())
	}
}

// AddSlices pushes all the elements from a slice to the stack s.
func (s *listStack[T]) AddSlice(slice []T) {
	for _, e := range slice {
		s.list.AddFront(e)
	}
}

// Clear removes all items on the stack.
func (s *listStack[T]) Clear() {
	s.list.Clear()
}

// Collect returns the stack as a slice.
func (s *listStack[T]) Collect() []T {
	return s.list.Collect()
}

// Contains checks if element e is part of the stack s.
func (s *listStack[T]) Contains(e T) bool {
	return s.list.Contains(e)
}

// Empty checks if the stack s is empty.
func (s *listStack[T]) Empty() bool {
	return s.list.Empty()
}

// Iterator returns in iterator for iterating through the stack s.
func (s *listStack[T]) Iterator() iterator.Iterator[T] {
	return s.list.Iterator()
}

// Len returns the size of the stack s.
func (s *listStack[T]) Len() int {
	return s.list.Len()
}

// Remove removes the first occurence of element e from the stack s.
func (s *listStack[T]) Remove(e T) bool {
	return s.list.Remove(e)
}

// RemoveAll removes all elements from the stack s that occur in iterable elements.
func (s *listStack[T]) RemoveAll(elements iterator.Iterable[T]) {
	s.list.RemoveAll(elements)
}
