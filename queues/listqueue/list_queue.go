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
	queue := ListQueue[T]{list: forwardlist.New[T]()}
	queue.AddSlice(elements)
	return &queue
}

// Add adds elements to the back of the queue.
func (queue *ListQueue[T]) Add(elements ...T) bool {
	for _, e := range elements {
		queue.list.Add(e)
	}
	return true
}

// AddAll adds the elements from some iterable elements to the queue.
func (queue *ListQueue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		queue.Add(it.Next())
	}
}

// AddSlice adds element from some slice s into the queue.
func (queue *ListQueue[T]) AddSlice(s []T) {
	queue.list.AddSlice(s)
}

// Remove removes the element e from the queue .
func (queue *ListQueue[T]) Remove(e T) bool {
	return queue.list.Remove(e)
}

// Clear removes all elements in the queue.
func (queue *ListQueue[T]) Clear() {
	queue.list.Clear()
}

// Collect converts queue into a slice.
func (queue *ListQueue[T]) Collect() []T {
	return queue.list.Collect()
}

// Contains checks if the elemen e is in the queue.
func (queue *ListQueue[T]) Contains(e T) bool {
	return queue.list.Contains(e)
}

// Empty checks if the queue is empty.
func (queue *ListQueue[T]) Empty() bool {
	return queue.list.Empty()
}

// Front returns the front element of the queue without removing it.
func (queue *ListQueue[T]) Front() T {
	defer func() {
		if r := recover(); r != nil {
			panic(queues.ErrNoFrontElement)
		}
	}()
	return queue.list.Front()
}

// Iterator returns an iterator for iterating through queue.
func (queue *ListQueue[T]) Iterator() iterator.Iterator[T] {
	return queue.list.Iterator()
}

// Len returns the size of the queue.
func (queue *ListQueue[T]) Len() int {
	return queue.list.Len()
}

// RemoveAll removes all the elements from some iterable elements that are in the queue.
func (queue *ListQueue[T]) RemoveAll(elements iterator.Iterable[T]) {
	queue.list.RemoveAll(elements)
}

// RemoveFront removes and returns the front element of the queue. Wil panic if no such element.
func (queue *ListQueue[T]) RemoveFront() T {
	defer func() {
		if r := recover(); r != nil {
			panic(queues.ErrNoFrontElement)
		}
	}()
	return queue.list.RemoveFront()
}

// String for printing the queue.
func (queue *ListQueue[T]) String() string {
	return fmt.Sprint(queue.list)
}
