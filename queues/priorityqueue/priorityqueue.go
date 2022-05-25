// Package priorityqueue provides an implementation of a priority queue
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

// New creates an empty priority queue. If min is true the its a min priority queue if min is false then its a max priority queue with the specified elements.
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
	ok := false
	for _, e := range elements {
		queue.heap.Insert(e)
		ok = true
	}
	return ok
}

// AddAll adds the elements from some iterable elements to the queue.
func (queue *PriorityQueue[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		queue.Add(it.Next())
	}
}

// Remove removes the elements from the queue .
func (queue *PriorityQueue[T]) Remove(elements ...T) bool {
	ok := false
	for _, element := range elements {
		ok = queue.heap.Delete(element)
		if queue.Empty() {
			return ok
		}
	}
	return ok
}

// Clear removes all elements in the queue.
func (queue *PriorityQueue[T]) Clear() {
	queue.heap.Clear()
}

// Collect converts queue into a slice. The elements in the slice are not ordered.
func (queue *PriorityQueue[T]) Collect() []T {
	return queue.heap.Collect()
}

// Contains checks if the element is in the queue.
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
func (iter *priorityQueueIterator[T]) Next() T {
	if !iter.HasNext() {
		panic(iterator.NoNextElementError)
	}
	n := iter.data[iter.i]
	iter.i++
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

// RemoveAll removes all the elements from some iterable elements that are in the queue.
func (queue *PriorityQueue[T]) RemoveAll(elements iterator.Iterable[T]) {
	iter := elements.Iterator()
	for iter.HasNext() {
		queue.heap.Delete(iter.Next())
		if queue.Empty() {
			break
		}
	}
}

// RemoveFront removes and returns the front element of the queue. Wil panic if no such element.
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
