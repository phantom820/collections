package slicequeue

import (
	"fmt"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types"
)

// SliceQueue slice based implementation of a queue.
type SliceQueue[T types.Equitable[T]] struct {
	data          []T
	modifications int
}

// New creates a slice based queue with the specified elements. If no specified elements an empty queue is returned.
func New[T types.Equitable[T]](elements ...T) *SliceQueue[T] {
	queue := SliceQueue[T]{data: make([]T, 0)}
	queue.Add(elements...)
	return &queue
}

// Add adds elements to the back of the queue.
func (queue *SliceQueue[T]) Add(elements ...T) bool {
	queue.modify()
	if len(elements) == 0 {
		return false
	}
	queue.data = append(queue.data, elements...)
	return true
}

// AddAll adds the elements from some iterable elements to the queue.
func (queue *SliceQueue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		queue.Add(it.Next())
	}
}

// Clear removes all elements in the queue.
func (queue *SliceQueue[T]) Clear() {
	queue.modify()
	queue.data = make([]T, 0)
}

// Collect converts queue into a slice.
func (queue *SliceQueue[T]) Collect() []T {
	return queue.data
}

// Contains checks if the elemen e is in the queue.
func (queue *SliceQueue[T]) Contains(e T) bool {
	for i, _ := range queue.data {
		if queue.data[i].Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the queue is empty.
func (queue *SliceQueue[T]) Empty() bool {
	return len(queue.data) == 0
}

// Front returns the front element of the queue without removing it.
func (queue *SliceQueue[T]) Front() T {
	if queue.Empty() {
		panic(errors.ErrNoSuchElement(queue.Len()))
	}
	return queue.data[0]
}

// modify increments the modification value
func (queue *SliceQueue[T]) modify() {
	queue.modifications++
}

// sliceQueueIterator model for implementing an iterator on a slice based queue.
type sliceQueueIterator[T types.Equitable[T]] struct {
	initialized      bool
	slice            []T
	getSlice         func() []T
	index            int
	modifications    int
	getModifications func() int
}

// HasNext check if the iterator has next element to produce.
func (it *sliceQueueIterator[T]) HasNext() bool {
	if !it.initialized {
		it.initialized = true
		it.modifications = it.getModifications()
		it.slice = it.getSlice()
	}
	if it.slice == nil || it.index >= len(it.slice) {
		return false
	}
	return true
}

// Next yields the next element from the iterator.
func (it *sliceQueueIterator[T]) Next() T {
	if !it.HasNext() {
		panic(errors.ErrNoNextElement())
	} else if it.modifications != it.getModifications() {
		panic(errors.ErrConcurrenModification())
	}
	e := it.slice[it.index]
	it.index++
	return e
}

// Cycle resets the iterator.
func (it *sliceQueueIterator[T]) Cycle() {
	it.index = 0
}

// Iterator returns an iterator for the queue.
func (queue *SliceQueue[T]) Iterator() iterator.Iterator[T] {
	return &sliceQueueIterator[T]{slice: queue.data, index: 0, getSlice: func() []T { return queue.data }, getModifications: func() int { return queue.modifications }}
}

func (queue *SliceQueue[T]) Len() int {
	return len(queue.data)
}

// indexOf finds the index of an element e in the queue. Gives -1 if the element is not present.
func (queue *SliceQueue[T]) indexOf(e T) int {
	for i, _ := range queue.data {
		if queue.data[i].Equals(e) {
			return i
		}
	}
	return -1
}

// Remove removes elements from the list. Only the first occurence of each element is removed.
func (queue *SliceQueue[T]) Remove(elements ...T) bool {
	queue.modify()
	n := queue.Len()
	for _, element := range elements {
		queue.remove(element)
		if queue.Empty() {
			break
		}
	}
	return n != queue.Len()
}

// remove removes the element from the queue. For internal use to support Remove.
func (queue *SliceQueue[T]) remove(element T) bool {
	i := queue.indexOf(element)
	if i == -1 {
		return false
	}
	queue.data = append(queue.data[0:i], queue.data[i+1:]...)
	return true
}

// RemoveAll removes all the elements in the queue that appear in the iterable.
func (queue *SliceQueue[T]) RemoveAll(iterable iterator.Iterable[T]) {
	it := iterable.Iterator()
	for it.HasNext() {
		queue.Remove(it.Next())
	}
}

// RemoveFront removes and returns the front element of the queue. Wil panic if no such element.
func (queue *SliceQueue[T]) RemoveFront() T {
	queue.modify()
	if queue.Empty() {
		panic(errors.ErrNoSuchElement(queue.Len()))
	}
	f := queue.data[0]
	queue.data = queue.data[1:]
	return f
}

// String for pretty printing the queue.
func (queue *SliceQueue[T]) String() string {
	return fmt.Sprint(queue.data)
}
