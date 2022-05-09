// ForwardList based implementation of a queue.
package queue

import (
	"collections/iterator"
	"collections/list"
	"collections/types"
	"fmt"
)

// ListQueue an interface providing singly linked list based implementation of a queue. This interface serves as an  abstraction for
// operating on an underlying list.
type ListQueue[T types.Equitable[T]] interface {
	Queue[T]
}

// listQueue concrete type for a singly linked list based queue.
type listQueue[T types.Equitable[T]] struct {
	list list.ForwardList[T]
}

// NewListQueue creates an empty list based queue.
func NewListQueue[T types.Equitable[T]]() ListQueue[T] {
	q := listQueue[T]{list: list.NewForwardList[T]()}
	return &q
}

// Add adds an element to the back of the queue.
func (q *listQueue[T]) Add(e T) bool {
	return q.list.Add(e)
}

// AddAll adds the elements from some iterable elements to the queue q.
func (q *listQueue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		q.Add(it.Next())
	}
}

// AddSlice adds element from some slice s into the queue q.
func (q *listQueue[T]) AddSlice(s []T) {
	q.list.AddSlice(s)
}

// Remove removes the element e from the queue q .
func (q *listQueue[T]) Remove(e T) bool {
	return q.list.Remove(e)
}

// Clear removes all elements in the queue q.
func (q *listQueue[T]) Clear() {
	q.list.Clear()
}

// Collect converts queue q into a slice.
func (q listQueue[T]) Collect() []T {
	return q.list.Collect()
}

// Contains checks if the elemen e is in the queue q.
func (q *listQueue[T]) Contains(e T) bool {
	return q.list.Contains(e)
}

// Empty checks if the queue q is empty.
func (q *listQueue[T]) Empty() bool {
	return q.list.Empty()
}

// Front returns the front element of the queue without removing it.
func (q *listQueue[T]) Front() T {
	defer func() {
		if r := recover(); r != nil {
			panic(NoFrontElementError)
		}
	}()
	return q.list.Front()
}

// Iterator returns an iterator for iterating through queue q.
func (q *listQueue[T]) Iterator() iterator.Iterator[T] {
	return q.list.Iterator()
}

// Len returns the size of the queue q.
func (q *listQueue[T]) Len() int {
	return q.list.Len()
}

// RemoveAll removes all the elements from some iterable elements that are in the queue q.
func (q *listQueue[T]) RemoveAll(elements iterator.Iterable[T]) {
	q.list.RemoveAll(elements)
}

// RemoveFront removes and returns the front element of the queue q. Wil panic if no such element.
func (q *listQueue[T]) RemoveFront() T {
	defer func() {
		if r := recover(); r != nil {
			panic(NoFrontElementError)
		}
	}()
	return q.list.RemoveFront()
}

// String for printing the queue q.
func (q listQueue[T]) String() string {
	return fmt.Sprint(q.list)
}
