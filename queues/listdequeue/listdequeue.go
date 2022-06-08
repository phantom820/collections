// Package listdequeue provides a singly linked list (with tail pointer) based implementation of a double ended queue.
package listdequeue

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/queues"

	"github.com/phantom820/collections/types"
)

// ListDequeue a list based implementation of a queue.
type ListDequeue[T types.Equitable[T]] struct {
	list *list.List[T]
}

// New creates a list based queue with the specified elements. If no specified elements an empty queue is returned.
func New[T types.Equitable[T]](elements ...T) *ListDequeue[T] {
	queue := ListDequeue[T]{list: list.New[T]()}
	queue.Add(elements...)
	return &queue
}

// Add adds elements to the back of the queue.
func (queue *ListDequeue[T]) Add(elements ...T) bool {
	if len(elements) == 0 {
		return false
	}
	for _, e := range elements {
		queue.list.Add(e)
	}
	return true
}

// Add adds elements to the front of the queue.
func (queue *ListDequeue[T]) AddFront(elements ...T) bool {
	n := queue.Len()
	for _, e := range elements {
		queue.list.AddFront(e)
	}
	return (n != queue.Len())
}

// AddAll adds the elements from some iterable elements to the back of the queue.
func (queue *ListDequeue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		queue.Add(it.Next())
	}
}

// Remove removes the elements from the queue.
func (queue *ListDequeue[T]) Remove(elements ...T) bool {
	return queue.list.Remove(elements...)
}

// Clear removes all elements in the queue.
func (queue *ListDequeue[T]) Clear() {
	queue.list.Clear()
}

// Collect converts queue into a slice.
func (queue *ListDequeue[T]) Collect() []T {
	return queue.list.Collect()
}

// Contains checks if the element is in the queue.
func (queue *ListDequeue[T]) Contains(element T) bool {
	return queue.list.Contains(element)
}

// Empty checks if the queue is empty.
func (queue *ListDequeue[T]) Empty() bool {
	return queue.list.Empty()
}

// Front returns the front element of the queue without removing it.
func (queue *ListDequeue[T]) Front() T {
	defer func() {
		if r := recover(); r != nil {
			panic(queues.ErrNoFrontElement)
		}
	}()
	return queue.list.Front()
}

// Back returns the back element of the queue without removing it.
func (queue *ListDequeue[T]) Back() T {
	defer func() {
		if r := recover(); r != nil {
			panic(queues.ErrNoBackElement)
		}
	}()
	return queue.list.Back()
}

// Iterator returns an iterator for iterating through queue.
func (queue *ListDequeue[T]) Iterator() iterator.Iterator[T] {
	return queue.list.Iterator()
}

// Len returns the size of the queue.
func (queue *ListDequeue[T]) Len() int {
	return queue.list.Len()
}

// RemoveAll removes all the elements from some iterable elements that are in the queue.
func (queue *ListDequeue[T]) RemoveAll(elements iterator.Iterable[T]) {
	queue.list.RemoveAll(elements)
}

// RemoveFront removes and returns the front element of the queue. Wil panic if no such element.
func (queue *ListDequeue[T]) RemoveFront() T {
	defer func() {
		if r := recover(); r != nil {
			panic(queues.ErrNoFrontElement)
		}
	}()
	return queue.list.RemoveFront()
}

// RemoveBack removes and returns the back element of the queue. Wil panic if no such element.
func (queue *ListDequeue[T]) RemoveBack() T {
	defer func() {
		if r := recover(); r != nil {
			panic(queues.ErrNoBackElement)
		}
	}()
	return queue.list.RemoveBack()
}

// String for printing the queue.
func (queue *ListDequeue[T]) String() string {
	return fmt.Sprint(queue.list)
}
