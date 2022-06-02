// Package slicedequeue provides a circular array based implementation of a double ended queue.
package slicedequeue

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/queues"
	"github.com/phantom820/collections/types"
)

const (
	buffer = 4
	scale  = 2
)

// SliceDequeue slice based implementation of a queue.
type SliceDequeue[T types.Equitable[T]] struct {
	front  int
	back   int
	buffer int // offset at the start of the slice
	data   []T
	len    int
	scale  int // how much we want to scale the buffer by in each expansion.
}

// New creates a slice based queue with the specified elements. If no specified elements an empty queue is returned.
func New[T types.Equitable[T]](elements ...T) *SliceDequeue[T] {
	queue := SliceDequeue[T]{data: make([]T, buffer), front: -1, back: -1, len: 0, buffer: buffer, scale: scale}
	queue.Add(elements...)
	return &queue
}

// AddFront adds elements to the front of the queue.
func (queue *SliceDequeue[T]) AddFront(elements ...T) bool {
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
	queue.data = make([]T, buffer)
	queue.buffer = buffer
	queue.scale = scale
}

// Collect returns a slice containing all the elements in the queue. The ordering of elements is not gueranteed.
func (queue *SliceDequeue[T]) Collect() []T {
	return queue.data[queue.front:]
}

// Contains checks if the element is in the queue.
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
	return queue.buffer == 0
}

// expand doubles the size of the front buffer of the underlying slice. For internal use to amortize cost of AddFront.
func (queue *SliceDequeue[T]) expand() {
	// double the memory and copy across.
	n := buffer * queue.scale
	data := make([]T, n)
	queue.data = append(data, queue.data...)
	queue.buffer = n
	queue.front = queue.front + n
	queue.scale++
}

// addRear adds an element to the back of the queue. For internal use to support AddBack.
func (queue *SliceDequeue[T]) addFront(element T) {
	if queue.isFull() {
		queue.expand()
	}
	// If queue is initially empty
	if queue.Empty() {
		queue.front = queue.buffer - 1
		queue.back = queue.buffer - 1
	} else { // front is at first position of queue
		queue.front--
	}
	queue.buffer--
	queue.data[queue.front] = element
	queue.len++
}

// addRear adds an element to the back of the queue. For internal use to support Add.
func (queue *SliceDequeue[T]) addRear(element T) {
	// If queue is initially empty
	if !queue.Empty() {
		queue.back++
		queue.data = append(queue.data, element)
		queue.len++
		return
	}
	queue.front = queue.buffer - 1
	queue.back = queue.buffer - 1
	queue.buffer--
	queue.data[queue.front] = element
	queue.len++
}

// Front returns the front element of the queue without removing it. Will panic if there is no such element.
func (queue *SliceDequeue[T]) Front() T {
	if queue.Empty() {
		panic(queues.ErrNoFrontElement)
	}
	return queue.data[queue.front]
}

// Back returns the back element of the queue without removing it. Will panic if there is no such element.
func (queue *SliceDequeue[T]) Back() T {
	if queue.Empty() {
		panic(queues.ErrNoBackElement)
	}
	return queue.data[queue.back]
}

// sliceQueueIterator model for implementing an iterator on a slice based queue.
type sliceQueueIterator[T types.Equitable[T]] struct {
	slice []T
	i     int
}

// HasNext checks if the iterator has a next element to yield.
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

// Iterator returns an iterator for the queue.
func (queue *SliceDequeue[T]) Iterator() iterator.Iterator[T] {
	if queue.Empty() {
		return &sliceQueueIterator[T]{slice: queue.data[0:0], i: 0}
	}
	return &sliceQueueIterator[T]{slice: queue.data[queue.front:], i: 0}
}

// Len returns the size of the queue.
func (queue *SliceDequeue[T]) Len() int {
	return queue.len
}

// indexOf finds the index of an element in the queue. Returns -1 if the element is not present.
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

// Remove removes elements from the queue. Only the first occurence of an element is removed.
func (queue *SliceDequeue[T]) Remove(elements ...T) bool {
	n := queue.Len()
	for _, element := range elements {
		queue.remove(element)
		if queue.Empty() {
			break
		}
	}
	return n != queue.Len()
}

// remove removes the element from the queue. For internal use to support Remove. Only the first occurence is removed.
func (queue *SliceDequeue[T]) remove(element T) bool {
	i := queue.indexOf(element)
	if i == -1 || queue.Empty() {
		return false
	}
	queue.data = append(queue.data[:i], queue.data[i+1:]...)
	queue.len--
	queue.back--
	if queue.Empty() {
		queue.front = -1
		queue.back = -1
		queue.buffer = len(queue.data)
	}
	return true
}

// shrinkFront for reducing the memory of the underlying slice for the queue. The parameter i is the index of an item that is to be skipped (for remove) ,
// to avoid skipping we should specify an index greater than the length of the slice.
func (queue *SliceDequeue[T]) shrinkFront() {
	loadFactor := float32(queue.len) / float32(len(queue.data))
	if loadFactor > 0.25 || queue.Empty() {
		return
	}
	data := make([]T, len(queue.data)/2)
	j := len(data) - queue.len
	for _, element := range queue.data[queue.front:] {
		data[j] = element
		j++
	}
	queue.front = len(data) - queue.len
	queue.buffer = queue.front - 1
	queue.scale = 2
	queue.back = j - 1
	queue.data = nil
	queue.data = data
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
	if queue.Empty() {
		panic(queues.ErrNoFrontElement)
	}
	front := queue.data[queue.front]
	queue.len--
	queue.front++
	queue.buffer++
	queue.shrinkFront()
	if queue.Empty() {
		queue.front = -1
		queue.back = -1
		queue.buffer = len(queue.data)
	}
	return front
}

// shrinkBack reduces the capacity underlying slice to avoid wasting memory.
func (queue *SliceDequeue[T]) shrinkBack() {
	if cap(queue.data) > 0 {
		loadFactor := float32(len(queue.data)) / float32(cap(queue.data))
		if loadFactor > 0.25 {
			return
		}
	}
	data := make([]T, len(queue.data))
	copy(data, queue.data)
	queue.data = nil
	queue.data = data
}

// RemoveBack removes and returns the element at the back of the queue. Will panic if there is no such element.
func (queue *SliceDequeue[T]) RemoveBack() T {
	if queue.Empty() {
		panic(queues.ErrNoBackElement)
	}
	back := queue.data[queue.back]
	queue.data = queue.data[:queue.back]
	queue.back--
	queue.len--
	queue.shrinkBack()
	if queue.Empty() {
		queue.front = -1
		queue.back = -1
		queue.buffer = len(queue.data)
	}

	return back
}

// String for pretty printing the queue.
func (queue *SliceDequeue[T]) String() string {
	return fmt.Sprint(queue.data[queue.front:])
}
