// Package priorityqueue provides an implementation of a priority queue.
package priorityqueue

import (
	"fmt"

	"github.com/phantom820/collections/heaps"
	"github.com/phantom820/collections/heaps/maxheap"
	"github.com/phantom820/collections/heaps/minheap"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/queues"
	"github.com/phantom820/collections/types"
)

// PriorityQueue an implementation of a priority queue.
type PriorityQueue[T types.Comparable[T]] struct {
	heap heaps.Heap[T]
}

// New creates an empty priority queue. If min is true the its a min priority queue otherwise a max priority queue with the specified elements.
func New[T types.Comparable[T]](min bool, elements ...T) *PriorityQueue[T] {
	var queue PriorityQueue[T]
	if min {
		queue.heap = minheap.New[T]()
	} else {
		queue.heap = maxheap.New[T]()
	}
	queue.Add(elements...)
	return &queue
}

// Add adds elements to the queue.
func (queue *PriorityQueue[T]) Add(elements ...T) bool {
	if len(elements) == 0 {
		return false
	}
	for _, element := range elements {
		queue.heap.Insert(element)
	}
	return true
}

// AddAll adds the elements from iterable elements to the queue.
func (queue *PriorityQueue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		queue.Add(it.Next())
	}
}

// Remove removes the elements from the queue.
func (queue *PriorityQueue[T]) Remove(elements ...T) bool {
	n := queue.Len()
	for _, element := range elements {
		queue.heap.Delete(element)
		if queue.Empty() {
			return n != queue.Len()
		}
	}
	return n != queue.Len()
}

// Clear removes all elements in the queue.
func (queue *PriorityQueue[T]) Clear() {
	queue.heap.Clear()
}

// Collect returns a slice containing all the elements in the priority. The elements of the slice are not ordered.
func (queue *PriorityQueue[T]) Collect() []T {
	return queue.heap.Collect()
}

//  Contains checks if an element is in the queue.
func (queue *PriorityQueue[T]) Contains(element T) bool {
	return queue.heap.Search(element)
}

// Empty checks if the queue is empty.
func (queue *PriorityQueue[T]) Empty() bool {
	return queue.heap.Empty()
}

// Front returns the front element of the queue without removing it.
func (queue *PriorityQueue[T]) Front() T {
	defer func() {
		if r := recover(); r != nil {
			panic(queues.ErrNoFrontElement)
		}
	}()
	return queue.heap.Top()
}

// priorityQueue iterator for iterating through a priority queue.
type priorityQueueIterator[T types.Comparable[T]] struct {
	data []T
	i    int
}

// HasNext checks if the iterator has a next element to yield.
func (iterator *priorityQueueIterator[T]) HasNext() bool {
	return iterator.i < len(iterator.data)
}

// Next returns the next element in the iterator it. Will panic if iterator is exhausted.
func (it *priorityQueueIterator[T]) Next() T {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	n := it.data[it.i]
	it.i++
	return n
}

// Cycle resets the iterator.
func (iterator *priorityQueueIterator[T]) Cycle() {
	iterator.i = 0
}

// Iterator returns an iterator for iterating through queue. The elements from the iterator are not ordered.
func (queue *PriorityQueue[T]) Iterator() iterator.Iterator[T] {
	iterator := priorityQueueIterator[T]{data: queue.Collect(), i: 0}
	return &iterator
}

// Len returns the size of the queue.
func (queue *PriorityQueue[T]) Len() int {
	return queue.heap.Len()
}

// RemoveAll removes all the elements in the queue that appear in the iterable.
func (queue *PriorityQueue[T]) RemoveAll(iterable iterator.Iterable[T]) {
	it := iterable.Iterator()
	for it.HasNext() {
		queue.heap.Delete(it.Next())
		if queue.Empty() {
			break
		}
	}
}

// RemoveFront removes and returns the front element of the queue. Will panic if there is no such element.
func (queue *PriorityQueue[T]) RemoveFront() T {
	defer func() {
		if r := recover(); r != nil {
			panic(queues.ErrNoFrontElement)
		}
	}()
	return queue.heap.DeleteTop()
}

// String for printing the queue.
func (queue *PriorityQueue[T]) String() string {
	return fmt.Sprint(queue.heap)
}
