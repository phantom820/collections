// Package slicedequeue provides a circular array based implementation of a double ended queue.
package slicedequeue

import (
	"fmt"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types"
)

const (
	capaity = 8
)

// SliceDequeue slice based implementation of a queue.
type SliceDequeue[T types.Equitable[T]] struct {
	front         int
	back          int
	data          []T
	len           int
	modifications int
}

// New creates a slice based queue with the specified elements. If no specified elements an empty queue is returned.
func New[T types.Equitable[T]](elements ...T) *SliceDequeue[T] {
	queue := SliceDequeue[T]{data: make([]T, capaity), front: -1, back: -1, len: 0}
	queue.Add(elements...)
	return &queue
}

// AddFront adds elements to the front of the queue.
func (queue *SliceDequeue[T]) AddFront(elements ...T) bool {
	queue.modify()
	if len(elements) == 0 {
		return false
	}
	for _, element := range elements {
		queue.addFront(element)
	}
	return true
}

// Add adds elements to the back of the queue.
func (queue *SliceDequeue[T]) Add(elements ...T) bool {
	queue.modify()
	if len(elements) == 0 {
		return false
	}
	for _, element := range elements {
		queue.addRear(element)
	}
	return true
}

// AddAll adds the elements from some iterable elements to the queue.
func (queue *SliceDequeue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		queue.Add(it.Next())
	}
}

// Clear removes all elements in the queue.
func (queue *SliceDequeue[T]) Clear() {
	queue.front = -1
	queue.back = -1
	queue.len = 0
	queue.data = make([]T, capaity)
}

// Collect returns a slice containing all the elements in the queue. The ordering of elements is not gueranteed.
func (queue *SliceDequeue[T]) Collect() []T {
	return queue.data[queue.front:]
}

// Contains checks if the element  is in the queue.
func (queue *SliceDequeue[T]) Contains(e T) bool {
	for i, _ := range queue.data {
		if queue.data[i].Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the queue is empty.
func (queue *SliceDequeue[T]) Empty() bool {
	return queue.len == 0
}

// isFull checks if the queue is full or not.
func (queue *SliceDequeue[T]) isFull() bool {
	return queue.front == 0 || queue.len == len(queue.data)
}

// expand doubles the size of the underlying slice.
func (queue *SliceDequeue[T]) expand() {
	// double the memory and copy across.
	n := len(queue.data) * 2
	data := make([]T, n)
	queue.front = n - len(queue.data)
	queue.back = len(data) - 1
	copy(data[queue.front:], queue.data)
	queue.data = data
}

// addRear adds an element to the back of the queue. For internal use to support AddFront.
func (queue *SliceDequeue[T]) addFront(element T) {
	if queue.isFull() {
		queue.expand()
	}
	// If queue is initially empty
	if queue.Empty() {
		queue.front = len(queue.data) - 1
		queue.back = len(queue.data) - 1
	} else {
		queue.front--
	}
	queue.data[queue.front] = element
	queue.len++
}

// addRear adds an element to the back of the queue. For internal use to support Add.
func (queue *SliceDequeue[T]) addRear(element T) {
	// If queue is initially empty
	if queue.Empty() {
		queue.addFront(element)
		return
	}
	queue.data = append(queue.data, element)
	queue.back++
	queue.len++
}

// Front returns the front element of the queue without removing it. Will panic if there is no such element.
func (queue *SliceDequeue[T]) Front() T {
	if queue.Empty() {
		panic(errors.ErrNoSuchElement(queue.len))
	}
	return queue.data[queue.front]
}

// Back returns the back element of the queue without removing it. Will panic if there is no such element.
func (queue *SliceDequeue[T]) Back() T {
	if queue.Empty() {
		panic(errors.ErrNoSuchElement(queue.len))
	}
	return queue.data[queue.back]
}

// modify increments the modification value
func (queue *SliceDequeue[T]) modify() {
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

// HasNext checks if the iterator has a next element to yield.
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

// Iterator returns an iterator for the queue.
func (queue *SliceDequeue[T]) Iterator() iterator.Iterator[T] {
	if queue.Empty() {
		return &sliceQueueIterator[T]{getSlice: func() []T { return []T{} }, index: 0, getModifications: func() int { return queue.modifications }}
	}
	return &sliceQueueIterator[T]{getSlice: func() []T { return queue.data[queue.front:] }, index: 0, getModifications: func() int { return queue.modifications }}
}

func (queue *SliceDequeue[T]) Len() int {
	return queue.len
}

// indexOf finds the index of an element in the queue. Gives -1 if the element is not present.
func (queue *SliceDequeue[T]) indexOf(element T) int {
	if queue.Empty() {
		return -1
	}
	for i, e := range queue.data[queue.front:] {
		if e.Equals(element) {
			return i + queue.front
		}
	}
	return -1
}

// Remove removes elements from the list. Only the first occurence of each element is removed.
func (queue *SliceDequeue[T]) Remove(elements ...T) bool {
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
func (queue *SliceDequeue[T]) remove(element T) bool {
	i := queue.indexOf(element)
	if i == -1 || queue.Empty() {
		return false
	} else if i == queue.front {
		queue.RemoveFront()
		return true
	} else if i == queue.back {
		queue.RemoveBack()
		return true
	}
	// front<i<back
	queue.data = append(queue.data[:i], queue.data[i+1:]...)
	queue.len--
	queue.back = len(queue.data) - 1
	return true
}

// shrinkFront reduces the memory used by the queue if less than a quarter is used. This is from the front perspective.
func (queue *SliceDequeue[T]) shrinkFront() {
	loadFactor := float32(queue.len) / float32(len(queue.data))
	if loadFactor >= 0.25 || queue.Empty() {
		return
	}
	data := make([]T, len(queue.data)/2)
	offset := len(data) - queue.len
	copy(data[offset:], queue.data[queue.front:])
	queue.data = nil
	queue.data = data
	queue.front = offset
	queue.back = len(queue.data)
}

// RemoveAll removes all the elements in queue set that appear in the iterable.
func (queue *SliceDequeue[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		queue.Remove(it.Next())
	}
}

// RemoveFront removes and returns the front element of the queue. Wil panic if no such element.
func (queue *SliceDequeue[T]) RemoveFront() T {
	queue.modify()
	if queue.Empty() {
		panic(errors.ErrNoSuchElement(queue.len))
	}
	front := queue.data[queue.front]
	queue.len--
	queue.front++
	queue.shrinkFront()
	if queue.Empty() {
		queue.front = -1
		queue.back = -1
	}
	return front
}

// shrinkBack shrinks the underlying slice if its capacity becomes too large compared to its len.
func (queue *SliceDequeue[T]) shrinkBack() {
	if cap(queue.data) > 0 {
		loadFactor := float32(queue.len) / float32(cap(queue.data))
		if loadFactor >= 0.25 {
			return
		}
	}
	data := make([]T, cap(queue.data)/2)
	offset := len(data) - queue.len
	copy(data[offset:], queue.data[queue.front:])
	queue.data = nil
	queue.data = data
	queue.front = offset
	queue.back = len(queue.data) - 1
}

// RemoveBack removes and returns the element at the back of the queue. Will panic if there is no such element.
func (queue *SliceDequeue[T]) RemoveBack() T {
	queue.modify()
	if queue.Empty() {
		panic(errors.ErrNoSuchElement(queue.len))
	}
	back := queue.data[queue.back]
	queue.data = queue.data[:queue.back]
	queue.back--
	queue.len--
	queue.shrinkBack()
	if queue.Empty() {
		queue.front = -1
		queue.back = -1
	}
	return back
}

// String for pretty printing the queue.
func (queue *SliceDequeue[T]) String() string {
	return fmt.Sprint(queue.data[queue.front:])
}
