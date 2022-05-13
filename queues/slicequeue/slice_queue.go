package slicequeue

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/queues"
	"github.com/phantom820/collections/types"
)

// SliceQueue slice based implementation of a queue.
type SliceQueue[T types.Equitable[T]] struct {
	data []T
}

// New creates a slice based queue with the specified elements. If no specified elements an empty queue is returned.
func New[T types.Equitable[T]](elements ...T) *SliceQueue[T] {
	q := SliceQueue[T]{data: make([]T, len(elements))}
	q.AddSlice(elements)
	return &q
}

// Add adds elements to the back of the queue.
func (q *SliceQueue[T]) Add(elements ...T) bool {
	q.data = append(q.data, elements...)
	return true
}

// AddAll adds the elements from some iterable elements to the queue.
func (q *SliceQueue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		q.Add(it.Next())
	}
}

// AddSlice adds element from a slice s into the queue.
func (q *SliceQueue[T]) AddSlice(s []T) {
	for _, e := range s {
		q.Add(e)
	}
}

// Clear removes all elements in the queue.
func (q *SliceQueue[T]) Clear() {
	q.data = make([]T, 0)
}

// Collect converts queue into a slice.
func (q *SliceQueue[T]) Collect() []T {
	return q.data
}

// Contains checks if the elemen e is in the queue.
func (q *SliceQueue[T]) Contains(e T) bool {
	for i, _ := range q.data {
		if q.data[i].Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the queue is empty.
func (q *SliceQueue[T]) Empty() bool {
	return len(q.data) == 0
}

// Front returns the front element of the queue without removing it.
func (q *SliceQueue[T]) Front() T {
	if q.Empty() {
		panic(queues.ErrNoFrontElement)
	}
	return q.data[0]
}

// sliceQueueIterator model for implementing an iterator on a slice based queue.
type sliceQueueIterator[T types.Equitable[T]] struct {
	slice []T
	i     int
}

// HasNext check if the iterator has next element to produce.
func (it *sliceQueueIterator[T]) HasNext() bool {
	if it.slice == nil || it.i >= len(it.slice) {
		return false
	}
	return true
}

// Next yields the next element from the iterator.
func (it *sliceQueueIterator[T]) Next() T {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	e := it.slice[it.i]
	it.i++
	return e
}

// Cycle resets the iterator.
func (it *sliceQueueIterator[T]) Cycle() {
	it.i = 0
}

// Iterator returns an iterator for iterating through queue.
func (q *SliceQueue[T]) Iterator() iterator.Iterator[T] {
	return &sliceQueueIterator[T]{slice: q.data, i: 0}
}

func (q *SliceQueue[T]) Len() int {
	return len(q.data)
}

// indexOf finds the index of an element e in the queue. Gives -1 if the element is not present.
func (q *SliceQueue[T]) indexOf(e T) int {
	for i, _ := range q.data {
		if q.data[i].Equals(e) {
			return i
		}
	}
	return -1
}

func (q *SliceQueue[T]) Remove(e T) bool {
	i := q.indexOf(e)
	if i != -1 {
		q.data = append(q.data[0:i], q.data[i+1:]...)
		return true
	}
	return false
}

// RemoveAll removes all the elements from some iterable elements that are in the queue.
func (q *SliceQueue[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		q.Remove(it.Next())
	}
}

// RemoveFront removes and returns the front element of the queue. Wil panic if no such element.
func (q *SliceQueue[T]) RemoveFront() T {
	if q.Empty() {
		panic(queues.ErrNoFrontElement)
	}
	f := q.data[0]
	q.data = q.data[1:]
	return f
}

// String for pretty printing the queue.
func (q *SliceQueue[T]) String() string {
	return fmt.Sprint(q.data)
}
