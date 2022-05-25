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
	queue := SliceQueue[T]{data: make([]T, len(elements))}
	queue.Add(elements...)
	return &queue
}

// Add adds elements to the back of the queue.
func (queue *SliceQueue[T]) Add(elements ...T) bool {
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
		panic(queues.ErrNoFrontElement)
	}
	return queue.data[0]
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
func (iter *sliceQueueIterator[T]) Next() T {
	if !iter.HasNext() {
		panic(iterator.NoNextElementError)
	}
	e := iter.slice[iter.i]
	iter.i++
	return e
}

// Cycle resets the iterator.
func (it *sliceQueueIterator[T]) Cycle() {
	it.i = 0
}

// Iterator returns an iterator for the queue.
func (queue *SliceQueue[T]) Iterator() iterator.Iterator[T] {
	return &sliceQueueIterator[T]{slice: queue.data, i: 0}
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

// RemoveAll removes all the elements from an iterable elements that are in the queue.
func (queue *SliceQueue[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		queue.Remove(it.Next())
	}
}

// RemoveFront removes and returns the front element of the queue. Wil panic if no such element.
func (queue *SliceQueue[T]) RemoveFront() T {
	if queue.Empty() {
		panic(queues.ErrNoFrontElement)
	}
	f := queue.data[0]
	queue.data = queue.data[1:]
	return f
}

// String for pretty printing the queue.
func (queue *SliceQueue[T]) String() string {
	return fmt.Sprint(queue.data)
}
