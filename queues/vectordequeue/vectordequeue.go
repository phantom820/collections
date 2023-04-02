// package vectordequeue defines an implementation of a double-ended queue that is backed by vector.

package vectordequeue

import (
	"fmt"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types/optional"
)

// VectorDequeue
type VectorDequeue[T comparable] struct {
	data   []T
	head   int
	offset int
	len    int
}

const (
	default_capacity = 4
	default_offset   = 2
)

// New creates a mutable list with the given elements.
func New[T comparable](elements ...T) *VectorDequeue[T] {
	dequeue := VectorDequeue[T]{
		data:   make([]T, 0, default_capacity),
		head:   -1,
		offset: default_offset,
		len:    0,
	}
	for _, e := range elements {
		dequeue.Add(e)
	}
	return &dequeue
}

// Add appends the specified element to the end of the list.
func (dequeue *VectorDequeue[T]) Add(e T) bool {
	if dequeue.Empty() {
		dequeue.data = append(dequeue.data, e)
		dequeue.head = 0
		dequeue.len++
		return true

	}
	dequeue.data = append(dequeue.data, e)
	dequeue.len++
	return true
}

func (dequeue *VectorDequeue[T]) expandFront() {
	newOffset := dequeue.offset * 2
	newData := make([]T, newOffset+len(dequeue.data))
	copy(newData[newOffset:], dequeue.data)
	dequeue.head = newOffset
	dequeue.data = newData
	dequeue.offset = newOffset
}

func (dequeue *VectorDequeue[T]) shrink() {
	if float64(dequeue.len) >= 0.25*float64(cap(dequeue.data)) {
		return
	} else if dequeue.Empty() && cap(dequeue.data) > default_capacity {
		newData := make([]T, 0, default_capacity)
		dequeue.data = newData
		return
	} else if dequeue.Empty() {
		return
	}
	newData := make([]T, cap(dequeue.data)/2)
	view := dequeue.data[dequeue.head:]
	j := len(newData) - 1
	for i := len(view) - 1; i >= 0; i-- {
		newData[j] = view[i]
		j--
	}
	dequeue.data = newData
	dequeue.head = j + 1
}

// AddFirst inserts the given element at the front of the dequeue, returns the previous front element.
func (dequeue *VectorDequeue[T]) AddFirst(e T) optional.Optional[T] {
	if dequeue.Empty() {
		dequeue.Add(e)
		return optional.Empty[T]()
	} else if dequeue.head == 0 {
		dequeue.expandFront()
		first := dequeue.data[dequeue.head]
		dequeue.head--
		dequeue.data[dequeue.head] = e
		dequeue.len++
		return optional.Of(first)
	}
	first := dequeue.data[dequeue.head]
	dequeue.head--
	dequeue.data[dequeue.head] = e
	dequeue.len++
	return optional.Of(first)
}

// RemoveFirst retrieves and removes the first element of the dequeue.
func (dequeue *VectorDequeue[T]) RemoveFirst() optional.Optional[T] {
	defer dequeue.shrink()
	if dequeue.Empty() {
		return optional.Empty[T]()
	} else if dequeue.len == 1 {
		temp := dequeue.data[dequeue.head]
		dequeue.head = -1
		dequeue.len = 0
		return optional.Of(temp)
	}
	temp := dequeue.data[dequeue.head]
	dequeue.len--
	dequeue.head++
	return optional.Of(temp)
}

// PeekFirst retrieves, but does not remove, the first element of this dequeue.
func (dequeue *VectorDequeue[T]) PeekFirst() optional.Optional[T] {
	if dequeue.Empty() {
		return optional.Empty[T]()
	}
	return optional.Of(dequeue.data[dequeue.head])
}

// AddLast inserts the given element at the back the dequeue, returns the previous last element.
func (dequeue *VectorDequeue[T]) AddLast(e T) optional.Optional[T] {
	last := dequeue.PeekLast()
	dequeue.Add(e)
	return last
}

// RemoveLast retrieves and removes the last element of the dequeue.
func (dequeue *VectorDequeue[T]) RemoveLast() optional.Optional[T] {
	defer dequeue.shrink()
	if dequeue.Empty() {
		return optional.Empty[T]()
	} else if dequeue.len == 1 {
		temp := dequeue.data[len(dequeue.data)-1]
		dequeue.len = 0
		dequeue.head = -1
		return optional.Of(temp)
	}
	temp := dequeue.data[len(dequeue.data)-1]
	dequeue.data = dequeue.data[:len(dequeue.data)-1]
	dequeue.len--
	return optional.Of(temp)
}

// PeekLast retrieves, but does not remove, the last element of this dequeue.
func (dequeue *VectorDequeue[T]) PeekLast() optional.Optional[T] {
	if dequeue.Empty() {
		return optional.Empty[T]()
	}
	return optional.Of(dequeue.data[len(dequeue.data)-1])
}

// Len returns the number of elements in the list.
func (dequeue *VectorDequeue[T]) Len() int {
	return dequeue.len
}

// AddAll adds all of the elements in the specified iterable to the dequeue.
func (dequeue *VectorDequeue[T]) AddAll(iterable iterable.Iterable[T]) bool {
	it := iterable.Iterator()
	for it.HasNext() {
		dequeue.Add(it.Next())
	}
	return true
}

// AddSlice adds all the elements in the slice to the dequeue.
func (dequeue *VectorDequeue[T]) AddSlice(s []T) bool {
	for _, e := range s {
		dequeue.Add(e)
	}
	return true
}

// Contains returns true if the dequeue contains the specified element.
func (dequeue *VectorDequeue[T]) Contains(e T) bool {
	if dequeue.Empty() {
		return false
	}
	for i := range dequeue.data[dequeue.head:] {
		if dequeue.data[i] == e {
			return true
		}
	}
	return false
}

// Clear removes all of the elements from the dequeue.
func (dequeue *VectorDequeue[T]) Clear() {
	dequeue.data = nil
	dequeue.data = make([]T, 0, default_capacity)
	dequeue.offset = default_offset
	dequeue.len = 0
}

// Empty returns true if the dequeue contains no elements.
func (dequeue *VectorDequeue[T]) Empty() bool {
	return dequeue.Len() == 0
}

// Remove unsupported operation.
func (dequeue *VectorDequeue[T]) Remove(e T) bool {
	panic(errors.UnsupportedOperation("Remove", "VectorDequeue"))
}

// RemoveIf removes all of the elements of the list that satisfy the given predicate.
func (dequeue *VectorDequeue[T]) RemoveIf(f func(T) bool) bool {
	panic(errors.UnsupportedOperation("RemoveIf", "VectorDequeue"))

}

// RemoveAll unsupported operation.
func (dequeue *VectorDequeue[T]) RemoveAll(iterable iterable.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("RemoveAll", "VectorDequeue"))
}

// RemoveSlice unsupported operation.
func (list *VectorDequeue[T]) RemoveSlice(s []T) bool {
	panic(errors.UnsupportedOperation("RemoveSlice", "VectorDequeue"))
}

// Copy returns a copy of the queue.
func (dequeue *VectorDequeue[T]) Copy() *VectorDequeue[T] {
	return nil
}

// RetainAll unsupported operation
func (dequeue *VectorDequeue[T]) RetainAll(c collections.Collection[T]) bool {
	panic(errors.UnsupportedOperation("RetainAll", "Dequeue"))
}

// ForEach performs the given action for each element of the dequeue.
func (dequeue *VectorDequeue[T]) ForEach(f func(T)) {

	if dequeue.Empty() {
		return
	}

	for _, e := range dequeue.data[dequeue.head:] {
		f(e)
	}
}

// ToSlice returns a slice containing the elements of the dequue.
func (dequeue *VectorDequeue[T]) ToSlice() []T {
	if dequeue.Empty() {
		return []T{}
	}
	slice := make([]T, 0, dequeue.len)
	slice = append(slice, dequeue.data[dequeue.head:]...)
	return slice
}

// // Equals returns true if the dequeue is equivalent to the given queue. Two lists are equal if they have the same size
// and contain the same elements in the same order.
func (dequeue *VectorDequeue[T]) Equals(other collections.Queue[T]) bool {
	if dequeue == other {
		return true
	} else if dequeue.Len() != other.Len() {
		return false
	}

	it1 := dequeue.Iterator()
	it2 := other.Iterator()

	for it1.HasNext() {
		if it1.Next() != it2.Next() {
			return false
		}
	}

	return true

}

// Iterator returns an iterator over the elements in the dequeue.
func (dequeue *VectorDequeue[T]) Iterator() iterator.Iterator[T] {
	return &dequeueIterator[T]{initialized: false, initialize: func() []T { return dequeue.ToSlice() }, index: 0, data: nil}
}

// iterator implememantation for [Vector].
type dequeueIterator[T comparable] struct {
	initialized bool
	initialize  func() []T
	index       int
	data        []T
}

// HasNext returns true if the iterator has more elements.
func (it *dequeueIterator[T]) HasNext() bool {
	if !it.initialized {
		it.data = it.initialize()
		it.initialized = true
	} else if it.data == nil {
		return false
	}
	return it.index < len(it.data)
}

// Next returns the next element in the iterator.
func (it *dequeueIterator[T]) Next() T {
	if !it.HasNext() {
		panic(errors.NoSuchElement())
	}
	index := it.index
	it.index++
	return it.data[index]
}

// String returns the string representation of the list.
func (dequeue VectorDequeue[T]) String() string {
	if dequeue.Empty() {
		return "[]"
	}
	return fmt.Sprint(dequeue.data[dequeue.head:])
}
