package dequeue

import (
	"fmt"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
	"github.com/phantom820/collections/types/optional"
)

// Dequeue a double ended queue.
type Dequeue[T comparable] struct {
	list collections.List[T]
}

// ListDequeueOf creates a list based dequeue of the given elements.
func ListDequeueOf[T comparable](elements ...T) Dequeue[T] {
	list := linkedlist.Of(elements...)
	return Dequeue[T]{list: &list}
}

// VectorDequeueOf creates a vector based dequeue of the given elements.
func VectorDequeueOf[T comparable](elements ...T) Dequeue[T] {
	list := vector.Of(elements...)
	return Dequeue[T]{list: &list}
}

// NewVectorDequeue creates a [Vector] based dequeue.
func NewVectorDequeue[T comparable]() *Dequeue[T] {
	return &Dequeue[T]{list: vector.New[T]()}
}

// NewListDequeue creates a [LinkedList] based dequeue.
func NewListDequeue[T comparable]() *Dequeue[T] {
	return &Dequeue[T]{list: linkedlist.New[T]()}
}

// Add appends the specified element to the end of the dequeue.
func (dequeue *Dequeue[T]) Add(e T) bool {
	return dequeue.list.Add(e)
}

// AddFirst inserts the given element at the front of the dequeue, returns the previuos front element.
func (dequeue *Dequeue[T]) AddFirst(e T) optional.Optional[T] {
	front := dequeue.PeekFirst()
	dequeue.list.AddAt(0, e)
	return front
}

// AddLast inserts the given element at the back the dequeue, returns the previous last element.
func (dequeue *Dequeue[T]) AddLast(e T) optional.Optional[T] {
	last := dequeue.PeekLast()
	dequeue.list.Add(e)
	return last
}

// Len returns the number of elements in the dequeue.
func (dequeue *Dequeue[T]) Len() int {
	return dequeue.list.Len()
}

// AddAll adds all of the elements in the specified iterable to the dequeue.
func (dequeue *Dequeue[T]) AddAll(iterable collections.Iterable[T]) bool {
	return dequeue.list.AddAll(iterable)
}

// AddSlice adds all the elements in the slice to the dequeue.
func (dequeue *Dequeue[T]) AddSlice(s []T) bool {
	return dequeue.list.AddSlice(s)
}

// PeekFirst retrieves, but does not remove, the first element of this dequeue.
func (dequeue *Dequeue[T]) PeekFirst() optional.Optional[T] {
	if dequeue.Empty() {
		return optional.Empty[T]()
	}
	return optional.Of(dequeue.list.At(0))
}

// PeekLast retrieves, but does not remove, the last element of this dequeue.
func (dequeue *Dequeue[T]) PeekLast() optional.Optional[T] {
	if dequeue.Empty() {
		return optional.Empty[T]()
	}
	return optional.Of(dequeue.list.At(dequeue.Len() - 1))
}

// Contains returns true if the list contains the specified element.
func (dequeue *Dequeue[T]) Contains(e T) bool {
	return dequeue.list.Contains(e)
}

// RemoveFirst retrieves and removes the first element of the dequeue.
func (dequeue *Dequeue[T]) RemoveFirst() optional.Optional[T] {
	if dequeue.Empty() {
		return optional.Empty[T]()
	}
	return optional.Of(dequeue.list.RemoveAt(0))
}

// RemoveLast retrieves and removes the last element of the dequeue.
func (dequeue *Dequeue[T]) RemoveLast() optional.Optional[T] {
	if dequeue.Empty() {
		return optional.Empty[T]()
	}
	return optional.Of(dequeue.list.RemoveAt(dequeue.Len() - 1))
}

// Clear removes all of the elements from the dequeue.
func (dequeue *Dequeue[T]) Clear() {
	dequeue.list.Clear()
}

// Empty returns true if the dequeue contains no elements.
func (dequeue *Dequeue[T]) Empty() bool {
	return dequeue.list.Len() == 0
}

// Remove removes the first occurrence of the specified element from the dequeue.
func (dequeue *Dequeue[T]) Remove(e T) bool {
	return dequeue.list.Remove(e)
}

// RemoveIf removes all of the elements of the dequeue that satisfy the given predicate.
func (dequeue *Dequeue[T]) RemoveIf(f func(T) bool) bool {
	return dequeue.list.RemoveIf(f)
}

// RemoveAll removes from the dequeue all of its elements that are contained in the specified collection.
func (dequeue *Dequeue[T]) RemoveAll(iterable collections.Iterable[T]) bool {
	return dequeue.list.RemoveAll(iterable)
}

// RemoveSlice removes all of the dequeue elements that are also contained in the specified slice.
func (dequeue *Dequeue[T]) RemoveSlice(s []T) bool {
	return dequeue.list.RemoveSlice(s)
}

// RetainAll retains only the elements in the dequeue that are contained in the specified collection.
func (dequeue *Dequeue[T]) RetainAll(c collections.Collection[T]) bool {
	return dequeue.list.RetainAll(c)
}

// ForEach performs the given action for each element of the dequeue.
func (dequeue *Dequeue[T]) ForEach(f func(T)) {
	dequeue.list.ForEach(f)
}

// ToSlice returns a slice containing the elements of the dequeue.
func (dequeue *Dequeue[T]) ToSlice() []T {
	return dequeue.list.ToSlice()
}

// Equals returns true if the list is equivalent to the given list. Two dequeues are equal if they have the same size
// and contain the same elements in the same order.
func (dequeue *Dequeue[T]) Equals(other *Dequeue[T]) bool {
	if dequeue == other {
		return true
	}
	return dequeue.list.Equals(other.list)
}

// Iterator returns an iterator over the elements in the dequeue.
func (dequeue *Dequeue[T]) Iterator() collections.Iterator[T] {
	return dequeue.list.Iterator()
}

// String returns the string representation of the dequeue.
func (dequeue Dequeue[T]) String() string {
	return fmt.Sprint(dequeue.list)
}
