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
	var stack ListStack[T] = ListStack[T]{list: forwardlist.New[T]()}
	stack.Add(elements...)
	return &stack
}

// Peek returns the element at the top of the stack without removing it. Will panic if stack has no top element to peek.
func (stack *ListStack[T]) Peek() T {
	defer func() {
		if r := recover(); r != nil {
			panic(stacks.ErrNoTopElement)
		}
	}()
	return stack.list.Front()
}

// Pop returns and removes the element at the top of the stack. Will panic if stack has no top element to pop.
func (stack *ListStack[T]) Pop() T {
	defer func() {
		if r := recover(); r != nil {
			panic(stacks.ErrNoTopElement)
		}
	}()
	return stack.list.RemoveFront()
}

// Add pushes elements  to the stack.
func (stack *ListStack[T]) Add(elements ...T) bool {
	if len(elements) == 0 {
		return false
	}
	for _, e := range elements {
		stack.list.AddFront(e)
	}
	return true
}

// AddAll pushes all the element from an iterable elements to the stack.
func (stack *ListStack[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		stack.list.AddFront(it.Next())
	}
}

// Clear removes all items on the stack.
func (stack *ListStack[T]) Clear() {
	stack.list.Clear()
}

// Collect returns the stack as a slice.
func (stack *ListStack[T]) Collect() []T {
	return stack.list.Collect()
}

// Contains checks if element e is part of the stack.
func (stack *ListStack[T]) Contains(e T) bool {
	return stack.list.Contains(e)
}

// Empty checks if the stack is empty.
func (stack *ListStack[T]) Empty() bool {
	return stack.list.Empty()
}

// Iterator returns in iterator for iterating through the stack.
func (stack *ListStack[T]) Iterator() iterator.Iterator[T] {
	return stack.list.Iterator()
}

// Len returns the size of the stack.
func (stack *ListStack[T]) Len() int {
	return stack.list.Len()
}

// Remove removes elements from the stack.
func (stack *ListStack[T]) Remove(elements ...T) bool {
	return stack.list.Remove(elements...)
}

// RemoveAll removes all elements from the stack that occur in iterable elementstack.
func (stack *ListStack[T]) RemoveAll(elements iterator.Iterable[T]) {
	stack.list.RemoveAll(elements)
}
