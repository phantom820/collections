package listqueue

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/queues"

	"github.com/phantom820/collections/types"
)

// ListQueue a list based implementation of a queue.
type ListQueue[T types.Equitable[T]] struct {
	list *forwardlist.ForwardList[T]
}

// New creates a list based queue with the specified elements. If no specified elements an empty queue is returned.
func New[T types.Equitable[T]](elements ...T) *ListQueue[T] {
	q := ListQueue[T]{list: forwardlist.New[T]()}
	q.AddSlice(elements)
	return &q
}

// Add adds elements to the back of the queue.
func (q *ListQueue[T]) Add(elements ...T) bool {
	for _, e := range elements {
		q.list.Add(e)
	}
	return true
}

// AddAll adds the elements from some iterable elements to the queue.
func (q *ListQueue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		q.Add(it.Next())
	}
}

// AddSlice adds element from some slice s into the queue.
func (q *ListQueue[T]) AddSlice(s []T) {
	q.list.AddSlice(s)
}

// Remove removes the element e from the queue .
func (q *ListQueue[T]) Remove(e T) bool {
	return q.list.Remove(e)
}

// Clear removes all elements in the queue.
func (q *ListQueue[T]) Clear() {
	q.list.Clear()
}

// Collect converts queue into a slice.
func (q ListQueue[T]) Collect() []T {
	return q.list.Collect()
}

// Contains checks if the elemen e is in the queue.
func (q *ListQueue[T]) Contains(e T) bool {
	return q.list.Contains(e)
}

// Empty checks if the queue is empty.
func (q *ListQueue[T]) Empty() bool {
	return q.list.Empty()
}

// Front returns the front element of the queue without removing it.
func (q *ListQueue[T]) Front() T {
	defer func() {
		if r := recover(); r != nil {
			panic(queues.ErrNoFrontElement)
		}
	}()
	return q.list.Front()
}

// Iterator returns an iterator for iterating through queue.
func (q *ListQueue[T]) Iterator() iterator.Iterator[T] {
	return q.list.Iterator()
}

// Len returns the size of the queue.
func (q *ListQueue[T]) Len() int {
	return q.list.Len()
}

// RemoveAll removes all the elements from some iterable elements that are in the queue.
func (q *ListQueue[T]) RemoveAll(elements iterator.Iterable[T]) {
	q.list.RemoveAll(elements)
}

// RemoveFront removes and returns the front element of the queue. Wil panic if no such element.
func (q *ListQueue[T]) RemoveFront() T {
	defer func() {
		if r := recover(); r != nil {
			panic(queues.ErrNoFrontElement)
		}
	}()
	return q.list.RemoveFront()
}

// String for printing the queue.
func (q ListQueue[T]) String() string {
	return fmt.Sprint(q.list)
}
