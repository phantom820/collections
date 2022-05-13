package stack

import (
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/stacks"
	"github.com/phantom820/collections/types"
)

// SliceStack a slice based implementation of a stack.
type SliceStack[T types.Equitable[T]] []T

// New creates a slice based stack with the specified elements, if there are none an empty stack is created.
func New[T types.Equitable[T]](elements ...T) SliceStack[T] {
	var s SliceStack[T] = make([]T, 0)
	s.AddSlice(elements)
	return s
}

// Peek returns the top element of stack without removing it. Will panic if s has no top element.
func (s *SliceStack[T]) Peek() T {
	if s.Empty() {
		panic(stacks.ErrNoTopElement)
	}

	return (*s)[s.Len()-1]
}

// Pop removes and returns the top element of stack. Will panic if s has no top element.
func (s *SliceStack[T]) Pop() T {
	if s.Empty() {
		panic(stacks.ErrNoTopElement)
	}
	t := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return t
}

// Add pushes the elements to the stack.
func (s *SliceStack[T]) Add(elements ...T) bool {
	for _, e := range elements {
		*s = append(*s, e)
	}
	return true
}

// AddAll pushes elements from iterable elements onto the stack.
func (s *SliceStack[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		s.Add(it.Next())
	}
}

// AddSlice adds element from a slice s into the stack q.
func (s *SliceStack[T]) AddSlice(slice []T) {
	for _, e := range slice {
		s.Add(e)
	}
}

// Clear removes all elements in the stack q.
func (q *SliceStack[T]) Clear() {
	*q = nil
	*q = make([]T, 0)
}

// Collect converts stack into a slice.
func (s *SliceStack[T]) Collect() []T {
	return *s
}

// Contains checks if the element e is in the stack.
func (s *SliceStack[T]) Contains(e T) bool {
	for i, _ := range *s {
		if (*s)[i].Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the stack is empty.
func (s *SliceStack[T]) Empty() bool {
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
func (s *SliceStack[T]) Iterator() iterator.Iterator[T] {
	return &sliceStackIterator[T]{slice: *s, i: len(*s) - 1}
}

// Len returns the size of the stack.
func (s *SliceStack[T]) Len() int {
	return len(*s)
}

// indexOf finds the index of an element e in the stack q. Gives -1 if the element is not present.
func (s *SliceStack[T]) indexOf(e T) int {
	for i, _ := range *s {
		if (*s)[i].Equals(e) {
			return i
		}
	}
	return -1
}

// Removes the first occurence of element e from the stack.
func (s *SliceStack[T]) Remove(e T) bool {
	i := s.indexOf(e)
	if i != -1 {
		*s = append((*s)[0:i], (*s)[i+1:]...)
		return true
	}
	return false
}

// RemoveAll removes all the elements from some iterable elements that are in the stack.
func (s *SliceStack[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		s.Remove(it.Next())
	}
}
