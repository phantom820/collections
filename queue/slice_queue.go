package queue

import "collections/interfaces"

type SliceQueue[T interfaces.Equitable[T]] interface {
	Queue[T]
}

type sliceQueue[T interfaces.Equitable[T]] []T

// NewSliceQueue creates a new slice based queue.
func NewSliceQueue[T interfaces.Equitable[T]]() SliceQueue[T] {
	var q sliceQueue[T]
	return &q
}

// Add adds an element to the back of the queue q.
func (q *sliceQueue[T]) Add(e T) bool {
	*q = append(*q, e)
	return true
}

// AddAll adds the elements from some iterable elements to the queue q.
func (q *sliceQueue[T]) AddAll(elements interfaces.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		q.Add(it.Next())
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

type sliceQueueIterator[T interfaces.Equitable[T]] struct {
	slice []T
	i     int
}

func (it *sliceQueueIterator[T]) HasNext() bool {
	if it.slice == nil || it.i >= len(it.slice) {
		return false
	}
	return true
}

func (it *sliceQueueIterator[T]) Next() T {
	if !it.HasNext() {
		panic(NoNextElementError)
	}
	e := it.slice[it.i]
	it.i++
	return e
}

func (it *sliceQueueIterator[T]) Cycle() {
	it.i = 0
}

// Iterator returns an iterator for iterating through queue q.
func (q *sliceQueue[T]) Iterator() interfaces.Iterator[T] {
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
func (q *sliceQueue[T]) RemoveAll(elements interfaces.Iterable[T]) {
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
