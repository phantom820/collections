package slicedequeue

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/queues"
	"github.com/phantom820/collections/types"
)

// SliceDequeue slice based implementation of a queue.
type SliceDequeue[T types.Equitable[T]] struct {
	front int
	back  int
	data  []T
	len   int
}

// New creates a slice based queue with the specified elements. If no specified elements an empty queue is returned.
func New[T types.Equitable[T]](elements ...T) *SliceDequeue[T] {
	queue := SliceDequeue[T]{data: make([]T, len(elements)), front: -1, back: 0, len: 0}
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
	queue.back = 0
	queue.len = 0
	queue.data = make([]T, 0)
}

// Collect converts queue into a slice.
func (queue *SliceDequeue[T]) Collect() []T {
	return queue.data[0:queue.len]
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
	if len(queue.data) == 0 {
		return true
	}
	return ((queue.front == 0 && queue.back == len(queue.data)-1) ||
		queue.front == queue.back+1)
}

// expand doubles the size of the underlying slice.
func (queue *SliceDequeue[T]) expand() {
	// double the memory and copy across.
	n := queue.len
	if n == 0 {
		n = 1
	} else {
		n = n * 2
	}
	data := make([]T, n)
	for i, _ := range queue.data {
		data[i] = queue.data[i]
	}
	queue.data = data
}

// addRear adds an element to the back of the queue. For internal use to support AddFront.
func (queue *SliceDequeue[T]) addFront(element T) {
	if queue.isFull() {
		queue.expand()
	}
	// If queue is initially empty
	if queue.front == -1 {
		queue.front = 0
		queue.back = 0
	} else if queue.front == 0 { // front is at first position of queue
		queue.front = len(queue.data) - 1
	} else {
		queue.front = queue.front - 1
	}
	// insert current element into Deque
	queue.data[queue.front] = element
	queue.len++
}

// addRear adds an element to the back of the queue. For internal use to support Add.
func (queue *SliceDequeue[T]) addRear(element T) {
	if queue.isFull() {
		queue.expand()
	}

	// If queue is initially empty
	if queue.front == -1 {
		queue.front = 0
		queue.back = 0
	} else if queue.back == len(queue.data)-1 { // rear is at last position of queue
		queue.back = 0
	} else { // increment rear end by '1'
		queue.back = queue.back + 1
	}
	// insert current element into Deque
	// fmt.Println(len(queue.data))
	queue.data[queue.back] = element
	queue.len++
}

// Front returns the front element of the queue without removing it. Will panic if no such element.
func (queue *SliceDequeue[T]) Front() T {
	if queue.Empty() {
		panic(queues.ErrNoFrontElement)
	}
	return queue.data[queue.front]
}

// Back returns the back element of the queue without removing it. Will panic if no such element.
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
func (queue *SliceDequeue[T]) Iterator() iterator.Iterator[T] {
	return &sliceQueueIterator[T]{slice: queue.data[0:queue.len], i: 0}
}

func (queue *SliceDequeue[T]) Len() int {
	return queue.len
}

// indexOf finds the index of an element e in the queue. Gives -1 if the element is not present.
func (queue *SliceDequeue[T]) indexOf(e T) int {
	for i, _ := range queue.data {
		if queue.data[i].Equals(e) {
			return i
		}
	}
	return -1
}

// Remove removes elements from the list. Only the first occurence of each element is removed.
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

// remove removes the element from the queue. For internal use to support Remove.
func (queue *SliceDequeue[T]) remove(element T) bool {
	i := queue.indexOf(element)
	if i == -1 {
		return false
	}
	queue.data = append(queue.data[0:i], queue.data[i+1:]...)
	if queue.front != 0 && queue.front != len(queue.data)-1 {
		queue.front = queue.front - 1
	}
	if queue.back != 0 && queue.back != len(queue.data)-1 {
		queue.back = queue.back - 1
	}
	queue.len--
	return true
}

// RemoveAll removes all the elements from an iterable elements that are in the queue.
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

	f := queue.data[queue.front]
	// Deque has only one element
	if queue.front == queue.back {
		queue.front = -1
		queue.back = -1
	} else
	// back to initial position
	if queue.front == len(queue.data) {
		queue.front = 0
	} else {
		queue.front = queue.front + 1 // increment front by '1' to remove current
		// front value from Deque
	}
	queue.len--
	return f
}

func (queue *SliceDequeue[T]) RemoveBack() T {
	if queue.Empty() {
		panic(queues.ErrNoBackElement)
	}

	back := queue.data[queue.back]
	// Deque has only one element
	if queue.front == queue.back {
		queue.front = -1
		queue.back = -1
	} else if queue.back == 0 {
		queue.back = len(queue.data) - 1
	} else {
		queue.back = queue.back - 1
	}
	return back
}

// String for pretty printing the queue.
func (queue *SliceDequeue[T]) String() string {
	return fmt.Sprint(queue.data[0:queue.len])
}
