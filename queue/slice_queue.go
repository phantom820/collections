package queue

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types"
)

// SliceQueue slice based implementation of a queue.
type SliceQueue[T types.Equitable[T]] []T

// NewSliceQueue creates a slice based queue with the specified elements. If no specified elements an empty queue is returned.
func NewSliceQueue[T types.Equitable[T]](elements ...T) *SliceQueue[T] {
	var q SliceQueue[T]
	q.AddSlice(elements)
	return &q
}

// Add adds an element to the back of the queue q.
func (q *SliceQueue[T]) Add(e T) bool {
	*q = append(*q, e)
	return true
}

// AddAll adds the elements from some iterable elements to the queue q.
func (q *SliceQueue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		q.Add(it.Next())
	}
}

// AddSlice adds element from a slice s into the queue q.
func (q *SliceQueue[T]) AddSlice(s []T) {
	for _, e := range s {
		q.Add(e)
	}
}

// Clear removes all elements in the queue q.
func (q *SliceQueue[T]) Clear() {
	*q = nil
}

// Collect converts queue q into a slice.
func (q *SliceQueue[T]) Collect() []T {
	return *q
}

// Contains checks if the elemen e is in the queue q.
func (q *SliceQueue[T]) Contains(e T) bool {
	for i, _ := range *q {
		if (*q)[i].Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the queue q is empty.
func (q *SliceQueue[T]) Empty() bool {
	return len(*q) == 0
}

// Front returns the front element of the queue without removing it.
func (q *SliceQueue[T]) Front() T {
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
func (q *SliceQueue[T]) Iterator() iterator.Iterator[T] {
	return &sliceQueueIterator[T]{slice: *q, i: 0}
}

func (q *SliceQueue[T]) Len() int {
	return len(*q)
}

// indexOf finds the index of an element e in the queue q. Gives -1 if the element is not present.
func (q *SliceQueue[T]) indexOf(e T) int {
	for i, _ := range *q {
		if (*q)[i].Equals(e) {
			return i
		}
	}
	return -1
}

func (q *SliceQueue[T]) Remove(e T) bool {
	i := q.indexOf(e)
	if i != -1 {
		*q = append((*q)[0:i], (*q)[i+1:]...)
		return true
	}
	return false
}

// RemoveAll removes all the elements from some iterable elements that are in the queue q.
func (q *SliceQueue[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		q.Remove(it.Next())
	}
}

// RemoveFront removes and returns the front element of the queue q. Wil panic if no such element.
func (q *SliceQueue[T]) RemoveFront() T {
	if q.Empty() {
		panic(NoFrontElementError)
	}
	f := (*q)[0]
	*q = (*q)[1:]
	return f
}

// String for pretty printing a slice based queue.
func (q *SliceQueue[T]) String() string {
	return fmt.Sprint(*q)
}
