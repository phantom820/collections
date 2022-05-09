// Slice based implementation of a queue.
package queue

import (
	"collections/iterator"
	"collections/types"
	"fmt"
)

// SliceQueue an interface providing a slice based implementation of a queue. This interface serves as an  abstraction for
// operating on an underlying slice.
type SliceQueue[T types.Equitable[T]] interface {
	Queue[T]
}

type sliceQueue[T types.Equitable[T]] []T

// NewSliceQueue creates an empty slice based queue.
func NewSliceQueue[T types.Equitable[T]]() SliceQueue[T] {
	var q sliceQueue[T]
	return &q
}

// Add adds an element to the back of the queue q.
func (q *sliceQueue[T]) Add(e T) bool {
	*q = append(*q, e)
	return true
}

// AddAll adds the elements from some iterable elements to the queue q.
func (q *sliceQueue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		q.Add(it.Next())
	}
}

// AddSlice adds element from a slice s into the queue q.
func (q *sliceQueue[T]) AddSlice(s []T) {
	for _, e := range s {
		q.Add(e)
	}
}

// Clear removes all elements in the queue q.
func (q *sliceQueue[T]) Clear() {
	*q = nil
}

// Collect converts queue q into a slice.
func (q *sliceQueue[T]) Collect() []T {
	return *q
}

// Contains checks if the elemen e is in the queue q.
func (q *sliceQueue[T]) Contains(e T) bool {
	for i, _ := range *q {
		if (*q)[i].Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the queue q is empty.
func (q *sliceQueue[T]) Empty() bool {
	return len(*q) == 0
}

// Front returns the front element of the queue without removing it.
func (q *sliceQueue[T]) Front() T {
	if q.Empty() {
		panic(NoFrontElementError)
	}
	return (*q)[0]
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

// Iterator returns an iterator for iterating through queue q.
func (q *sliceQueue[T]) Iterator() iterator.Iterator[T] {
	return &sliceQueueIterator[T]{slice: *q, i: 0}
}

func (q *sliceQueue[T]) Len() int {
	return len(*q)
}

// indexOf finds the index of an element e in the queue q. Gives -1 if the element is not present.
func (q *sliceQueue[T]) indexOf(e T) int {
	for i, _ := range *q {
		if (*q)[i].Equals(e) {
			return i
		}
	}
	return -1
}

func (q *sliceQueue[T]) Remove(e T) bool {
	i := q.indexOf(e)
	if i != -1 {
		*q = append((*q)[0:i], (*q)[i+1:]...)
		return true
	}
	return false
}

// RemoveAll removes all the elements from some iterable elements that are in the queue q.
func (q *sliceQueue[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		q.Remove(it.Next())
	}
}

// RemoveFront removes and returns the front element of the queue q. Wil panic if no such element.
func (q *sliceQueue[T]) RemoveFront() T {
	if q.Empty() {
		panic(NoFrontElementError)
	}
	f := (*q)[0]
	*q = (*q)[1:]
	return f
}

// String for pretty printing a slice based queue.
func (q *sliceQueue[T]) String() string {
	return fmt.Sprint(*q)
}
