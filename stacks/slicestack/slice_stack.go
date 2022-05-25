package slicestack

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/stacks"
	"github.com/phantom820/collections/types"
)

// SliceStack a slice based implementation of a stack.
type SliceStack[T types.Equitable[T]] struct {
	data []T
}

// New creates a slice based stack with the specified elements, if there are none an empty stack is created.
func New[T types.Equitable[T]](elements ...T) *SliceStack[T] {
	stack := SliceStack[T]{data: make([]T, 0)}
	stack.Add(elements...)
	return &stack
}

// Peek returns the top element of stack without removing it. Will panic if s has no top element.
func (stack *SliceStack[T]) Peek() T {
	if stack.Empty() {
		panic(stacks.ErrNoTopElement)
	}
	return stack.data[stack.Len()-1]
}

// Pop removes and returns the top element of stack. Will panic if s has no top element.
func (stack *SliceStack[T]) Pop() T {
	if stack.Empty() {
		panic(stacks.ErrNoTopElement)
	}
	t := stack.data[stack.Len()-1]
	stack.data = stack.data[:stack.Len()-1]
	return t
}

// Add pushes the elements to the stack.
func (stack *SliceStack[T]) Add(elements ...T) bool {
	if len(elements) == 0 {
		return false
	}
	stack.data = append(stack.data, elements...)
	return true
}

// AddAll pushes elements from iterable elements onto the stack.
func (stack *SliceStack[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		stack.Add(it.Next())
	}
}

// Clear removes all elements in the stack.
func (stack *SliceStack[T]) Clear() {
	stack.data = make([]T, 0)
}

// Collect converts stack into a slice by returning underlying slice.
func (stack *SliceStack[T]) Collect() []T {
	return stack.data
}

// Contains checks if the element e is in the stack.
func (stack *SliceStack[T]) Contains(e T) bool {
	for i, _ := range stack.data {
		if stack.data[i].Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the stack is empty.
func (stack *SliceStack[T]) Empty() bool {
	return len(stack.data) == 0
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
func (stack *SliceStack[T]) Iterator() iterator.Iterator[T] {
	return &sliceStackIterator[T]{slice: stack.data, i: len(stack.data) - 1}
}

// Len returns the size of the stack.
func (stack *SliceStack[T]) Len() int {
	return len(stack.data)
}

// indexOf finds the index of an element e in the stack q. Gives -1 if the element is not present.
func (stack *SliceStack[T]) indexOf(e T) int {
	for i, _ := range stack.data {
		if stack.data[i].Equals(e) {
			return i
		}
	}
	return -1
}

// Remove remove the elements from the stack.
func (stack *SliceStack[T]) Remove(elements ...T) bool {
	n := stack.Len()
	for _, element := range elements {
		stack.remove(element)
		if stack.Empty() {
			break
		}
	}
	return (n != stack.Len())
}

// removes the first occurence of element  from the stack.
func (stack *SliceStack[T]) remove(e T) bool {
	i := stack.indexOf(e)
	if i == -1 {
		return false
	}
	stack.data = append(stack.data[0:i], stack.data[i+1:]...)
	return true
}

// RemoveAll removes all the elements from some iterable elements that are in the stack.
func (stack *SliceStack[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		stack.Remove(it.Next())
	}
}

// String for pretty printing the stack.
func (stack *SliceStack[T]) String() string {
	return fmt.Sprint(stack.data)
}
