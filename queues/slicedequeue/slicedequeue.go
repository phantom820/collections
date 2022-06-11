// Package slicedequeue provides a circular array based implementation of a double ended queue.
package slicedequeue

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/queues"
	"github.com/phantom820/collections/types"
)

const (
	capacity        = 8
	loadFactorLimit = 0.25
)

type entry[T types.Equitable[T]] struct {
	element T
}

// SliceDequeue slice based implementation of a queue.
type SliceDequeue[T types.Equitable[T]] struct {
	len   int
	data  []*entry[T]
	front int
	back  int
}

// New creates a slice based queue with the specified elements. If no specified elements an empty queue is returned.
func New[T types.Equitable[T]](elements ...T) *SliceDequeue[T] {
	queue := SliceDequeue[T]{data: make([]*entry[T], capacity), front: -1, back: -1}
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
		queue.addBack(element)
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
	queue.data = make([]*entry[T], 0)
}

// Collect returns a slice containing all the elements in the queue. The ordering of elements is not gueranteed.
func (queue *SliceDequeue[T]) Collect() []T {
	data := make([]T, 0)
	for _, e := range queue.data {
		if e != nil {
			data = append(data, e.element)
		}
	}
	return data
}

// Contains checks if the element is in the queue.
func (queue *SliceDequeue[T]) Contains(e T) bool {
	for i, _ := range queue.data {
		if queue.data[i] != nil && queue.data[i].element.Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the queue is empty.
func (queue *SliceDequeue[T]) Empty() bool {
	return queue.len == 0
}

// full checks if the queue is full or not.
func (queue *SliceDequeue[T]) full() bool {
	return queue.len == len(queue.data)
}

// expandFront doubles the size of the front buffer of the underlying slice. For internal use to amortize cost of AddFront.
func (queue *SliceDequeue[T]) expand() {
	if len(queue.data) == 0 {
		queue.data = make([]*entry[T], capacity)
		return
	} else if queue.front == 0 && queue.back == len(queue.data)-1 {
		data := make([]*entry[T], len(queue.data)*2)
		copy(data, queue.data)
		queue.data = nil
		queue.data = data
		return
	}
	data := make([]*entry[T], len(queue.data)*2)
	size := len(queue.data) - queue.front
	j := 0
	for i, _ := range data {
		if i <= queue.back {
			data[i] = queue.data[i]
			j++
		} else if i >= len(data)-size && j < len(queue.data) {
			data[i] = queue.data[j]
			j++
		}
	}
	queue.front = len(data) - size
	queue.data = nil
	queue.data = data
}

// addRear adds an element to the back of the queue. For internal use to support AddBack.
func (queue *SliceDequeue[T]) addFront(element T) {
	if queue.full() {
		queue.expand()
	}
	switch queue.front {
	case -1:
		queue.front = 0
		queue.back = 0
	case 0:
		queue.front = len(queue.data) - 1
	default:
		queue.front--
	}
	queue.len++
	queue.data[queue.front] = &entry[T]{element: element}
}

// addBack adds an element to the back of the queue. For internal use to support Add.
func (queue *SliceDequeue[T]) addBack(element T) {
	if queue.full() {
		queue.expand()
	}
	switch queue.back {
	case -1:
		queue.front = 0
		queue.back = 0
	case len(queue.data) - 1:
		queue.back = 0
	default:
		queue.back++
	}
	queue.len++
	queue.data[queue.back] = &entry[T]{element: element}
}

// Front returns the front element of the queue without removing it. Will panic if there is no such element.
func (queue *SliceDequeue[T]) Front() T {
	if queue.Empty() {
		panic(queues.ErrNoFrontElement)
	}
	front := queue.data[queue.front].element
	return front
}

// loadFactor returns the load factor of the queue.
func (queue *SliceDequeue[T]) loadFactor() float32 {
	if len(queue.data) == 0 {
		return 0
	}
	return float32(queue.len) / float32(len(queue.data))
}

// Back returns the back element of the queue without removing it. Will panic if there is no such element.
func (queue *SliceDequeue[T]) Back() T {
	if queue.Empty() {
		panic(queues.ErrNoBackElement)
	}
	back := queue.data[queue.back].element
	return back
}

// sliceQueueIterator model for implementing an iterator on a slice based queue.
type sliceQueueIterator[T types.Equitable[T]] struct {
	slice []*entry[T]
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
	return e.element
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
	return nil
}

// Len returns the size of the queue.
func (queue *SliceDequeue[T]) Len() int {
	return queue.len
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
	i := -1
	if i == -1 || queue.Empty() {
		return false
	}
	return true
}

// RemoveAll removes all the elements in queue set that appear in the iterable.
func (queue *SliceDequeue[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		queue.Remove(it.Next())
	}
}

// shrink reduces the memory allocated for the queue there is few items compared to actual size.
func (queue *SliceDequeue[T]) shrink() {
	if queue.loadFactor() > loadFactorLimit {
		return
	}
	if queue.Empty() {
		queue.data = nil
		queue.data = make([]*entry[T], capacity)
		return
	}
	data := make([]*entry[T], len(queue.data)/2)
	start := -1
	end := -1
	for i, element := range queue.data {
		if element == nil {
			if start == -1 {
				start = i
				end = i
			} else {
				end++
			}
		}
	}
	copy(data[:start], queue.data[:start])
	offset := len(queue.data) - end - 1
	tail := len(queue.data) - offset
	newFront := len(data) - offset
	copy(data[newFront:], queue.data[tail:])
	queue.data = nil
	queue.data = data
	if queue.front != 0 {
		queue.front = newFront
	}
}

// RemoveFront removes and returns the front element of the queue. Wil panic if no such element.
func (queue *SliceDequeue[T]) RemoveFront() T {
	if queue.Empty() {
		panic(queues.ErrNoFrontElement)
	}
	front := queue.data[queue.front].element
	queue.data[queue.front] = nil
	switch queue.front {
	case queue.back:
		queue.front = -1
		queue.back = -1
	case len(queue.data) - 1:
		queue.front = 0
	default:
		queue.front++
	}
	queue.len--
	queue.shrink()
	return front
}

// RemoveBack removes and returns the element at the back of the queue. Will panic if there is no such element.
func (queue *SliceDequeue[T]) RemoveBack() T {
	if queue.Empty() {
		panic(queues.ErrNoBackElement)
	}
	back := queue.data[queue.back].element
	queue.data[queue.back] = nil
	switch queue.back {
	case queue.front:
		queue.front = -1
		queue.back = -1
	case 0:
		queue.back = len(queue.data) - 1
	default:
		queue.back--
	}
	queue.len--
	queue.shrink()
	return back
}

// String for pretty printing the queue.
func (queue *SliceDequeue[T]) String() string {
	return fmt.Sprint(queue.Collect())
}
