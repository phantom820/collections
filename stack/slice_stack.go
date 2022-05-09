package stack

import (
	"collections/iterator"
	"collections/types"
)

// SliceStack an interface providing slice based implementation of a stack. This interface serves as an  abstraction for
// operating on an underlying slice.
type SliceStack[T types.Equitable[T]] interface {
	Stack[T]
}

type sliceStack[T types.Equitable[T]] []T

// NewSliceStack creates an empty slice based stack.
func NewSliceStack[T types.Equitable[T]]() ListStack[T] {
	var s sliceStack[T]
	return &s
}

// Peek returns the top element of stack s without removing it. Will panic if s has no top element.
func (s *sliceStack[T]) Peek() T {
	if s.Empty() {
		panic(NoTopElementError)
	}

	return (*s)[s.Len()-1]
}

// Pop removes and returns the top element of stack s. Will panic if s has no top element.
func (s *sliceStack[T]) Pop() T {
	if s.Empty() {
		panic(NoTopElementError)
	}
	t := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return t
}

// Add pushes the element e to the stack s.
func (s *sliceStack[T]) Add(e T) bool {
	*s = append(*s, e)
	return true
}

// AddAll pushes elements from iterable elements onto the stack s.
func (s *sliceStack[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		s.Add(it.Next())
	}
}

// AddSlice adds element from a slice s into the stack q.
func (s *sliceStack[T]) AddSlice(slice []T) {
	for _, e := range slice {
		s.Add(e)
	}
}

// Clear removes all elements in the stack q.
func (q *sliceStack[T]) Clear() {
	*q = nil
}

// Collect converts stack s into a slice.
func (s *sliceStack[T]) Collect() []T {
	return *s
}

// Contains checks if the element e is in the stack s.
func (s *sliceStack[T]) Contains(e T) bool {
	for i, _ := range *s {
		if (*s)[i].Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the stack s is empty.
func (s *sliceStack[T]) Empty() bool {
	return len(*s) == 0
}

// sliceStackIterator model for implementing an iterator on a slice based stack.
type sliceStackIterator[T types.Equitable[T]] struct {
	slice []T
	i     int
}

// HasNext check if the iterator has next element to produce.
func (it *sliceStackIterator[T]) HasNext() bool {
	if it.slice == nil || it.i < 0 {
		return false
	}
	return true
}

// Next yields the next element from the iterator it.
func (it *sliceStackIterator[T]) Next() T {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	e := it.slice[it.i]
	it.i--
	return e
}

// Cycle resets the iterator it.
func (it *sliceStackIterator[T]) Cycle() {
	it.i = len(it.slice) - 1
}

// Iterator returns an iterator for iterating through stack q.
func (s *sliceStack[T]) Iterator() iterator.Iterator[T] {
	return &sliceStackIterator[T]{slice: *s, i: len(*s) - 1}
}

// Len returns the size of the stack s.
func (s *sliceStack[T]) Len() int {
	return len(*s)
}

// indexOf finds the index of an element e in the stack q. Gives -1 if the element is not present.
func (s *sliceStack[T]) indexOf(e T) int {
	for i, _ := range *s {
		if (*s)[i].Equals(e) {
			return i
		}
	}
	return -1
}

// Removes the first occurence of element e from the stack s.
func (s *sliceStack[T]) Remove(e T) bool {
	i := s.indexOf(e)
	if i != -1 {
		*s = append((*s)[0:i], (*s)[i+1:]...)
		return true
	}
	return false
}

// RemoveAll removes all the elements from some iterable elements that are in the stack s.
func (s *sliceStack[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		s.Remove(it.Next())
	}
}
