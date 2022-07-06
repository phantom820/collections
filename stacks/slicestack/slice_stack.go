// Package slicestack provides a slice based implementation of a stack.
package slicestack

import (
	"fmt"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types"
)

// SliceStack a slice based implementation of a stack.
type SliceStack[T types.Equitable[T]] struct {
	data          []T
	modifications int
}

// New creates a slice based stack with the specified elements, if there are none an empty stack is created.
func New[T types.Equitable[T]](elements ...T) *SliceStack[T] {
	stack := SliceStack[T]{data: make([]T, 0)}
	stack.Add(elements...)
	return &stack
}

// Peek returns the top element of stack without removing it. Will panic if there is no top element.
func (stack *SliceStack[T]) Peek() T {
	if stack.Empty() {
		panic(errors.ErrNoSuchElement(stack.Len()))
	}
	return stack.data[stack.Len()-1]
}

// Pop removes and returns the top element of stack. Will panic if there is no top element.
func (stack *SliceStack[T]) Pop() T {
	stack.modify()
	if stack.Empty() {
		panic(errors.ErrNoSuchElement(stack.Len()))
	}
	t := stack.data[stack.Len()-1]
	stack.data = stack.data[:stack.Len()-1]
	return t
}

// Add pushes elements to the stack.
func (stack *SliceStack[T]) Add(elements ...T) bool {
	stack.modify()
	if len(elements) == 0 {
		return false
	}
	stack.data = append(stack.data, elements...)
	return true
}

// AddAll pushes elements from an iterable onto the stack.
func (stack *SliceStack[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		stack.Add(it.Next())
	}
}

// Clear removes all elements in the stack.
func (stack *SliceStack[T]) Clear() {
	stack.modify()
	stack.data = make([]T, 0)
}

// Collect returns a slice containing all the elements in the stack.
func (stack *SliceStack[T]) Collect() []T {
	return stack.data
}

// Contains checks if the element is in the stack.
func (stack *SliceStack[T]) Contains(element T) bool {
	for i, _ := range stack.data {
		if stack.data[i].Equals(element) {
			return true
		}
	}
	return false
}

// Empty checks if the stack is empty.
func (stack *SliceStack[T]) Empty() bool {
	return len(stack.data) == 0
}

// modify increments the modification value
func (stack *SliceStack[T]) modify() {
	stack.modifications++
}

// sliceStackIterator model for implementing an iterator on a slice based stack.
type sliceStackIterator[T types.Equitable[T]] struct {
	initialized      bool
	slice            []T
	getSlice         func() []T
	index            int
	modifications    int
	getModifications func() int
}

// HasNext checks if the iterator has a next element to yield.
func (it *sliceStackIterator[T]) HasNext() bool {
	if !it.initialized {
		it.initialized = true
		it.modifications = it.getModifications()
		it.slice = it.getSlice()
		it.index = len(it.slice) - 1
	}
	if it.slice == nil || it.index < 0 {
		return false
	}
	return true
}

// Next yields the next element in the iterator. Will panic if the iterator has no next element.
func (it *sliceStackIterator[T]) Next() T {
	if !it.HasNext() {
		panic(errors.ErrNoNextElement())
	} else if it.modifications != it.getModifications() {
		panic(errors.ErrConcurrenModification())
	}
	e := it.slice[it.index]
	it.index--
	return e
}

// Iterator returns an iterator for the stack.
func (stack *SliceStack[T]) Iterator() iterator.Iterator[T] {
	return &sliceStackIterator[T]{slice: []T{}, getSlice: func() []T { return stack.data }, getModifications: func() int { return stack.modifications }}
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

// Remove removes elements from the stack.
func (stack *SliceStack[T]) Remove(elements ...T) bool {
	stack.modify()
	n := stack.Len()
	for _, element := range elements {
		stack.remove(element)
		if stack.Empty() {
			break
		}
	}
	return (n != stack.Len())
}

// removes the first occurence of the element  from the stack.
func (stack *SliceStack[T]) remove(element T) bool {
	i := stack.indexOf(element)
	if i == -1 {
		return false
	}
	stack.data = append(stack.data[0:i], stack.data[i+1:]...)
	return true
}

// RemoveAll removes all the elements in the stack that appear in the iterable.
func (stack *SliceStack[T]) RemoveAll(iterable iterator.Iterable[T]) {
	it := iterable.Iterator()
	for it.HasNext() {
		stack.Remove(it.Next())
	}
}

// String for pretty printing the stack.
func (stack *SliceStack[T]) String() string {
	return fmt.Sprint(stack.data)
}
