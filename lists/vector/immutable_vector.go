package vector

import (
	"fmt"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
)

// ImmutableVector an immutable version of [Vector].
type ImmutableVector[T comparable] struct {
	vector Vector[T]
}

// ImmutableOf creates an ImmutableVector with the specified elements.
func ImmutableOf[T comparable](elements ...T) ImmutableVector[T] {
	return ImmutableVector[T]{Of(elements...)}
}

// At returns the element at the specified position in the list.
func (list ImmutableVector[T]) At(i int) T {
	return list.vector.At(i)
}

// Contains returns true if this list contains the specified element.
func (list ImmutableVector[T]) Contains(e T) bool {
	return list.vector.Contains(e)
}

// Len returns the number of elements in the list.
func (list ImmutableVector[T]) Len() int {
	return list.vector.Len()
}

// Empty returns true if the list contains no elements.
func (list ImmutableVector[T]) Empty() bool {
	return list.vector.Empty()
}

// ForEach performs the given action for each element of the list.
func (list ImmutableVector[T]) ForEach(f func(T)) {
	list.vector.ForEach(f)
}

// Iterator returns an iterator over the elements in the list.
func (list ImmutableVector[T]) Iterator() collections.Iterator[T] {
	return list.vector.Iterator()
}

// Equals returns true if the list is equivalent to the given list. Two lists are equal if they are the same reference or have the same size and contain
// the same elements.
func (list ImmutableVector[T]) Equals(otherSet ImmutableVector[T]) bool {
	return list.vector.Equals(&otherSet.vector)
}

// ToSlice returns a slice containing all the elements in the list.
func (list ImmutableVector[T]) ToSlice() []T {
	return list.vector.ToSlice()
}

// String returns the string representation of the list.
func (list ImmutableVector[T]) String() string {
	return fmt.Sprint(list.vector.data)
}

func (list ImmutableVector[T]) Remove(e T) bool {
	panic(errors.UnsupportedOperation("Remove", "ImmutableVector"))
}
func (list ImmutableVector[T]) RemoveIf(func(T) bool) bool {
	panic(errors.UnsupportedOperation("RemoveIf", "ImmutableVector"))
}

func (list ImmutableVector[T]) RemoveAll(iterable collections.Iterable[T]) bool {
	panic("RemoveAll")
}

// RetainAll retains only the elements in the list that are contained in the specified collection.
func (list ImmutableVector[T]) RetainAll(c collections.Collection[T]) bool {
	panic(errors.UnsupportedOperation("RetainAll", "ImmutableVector"))
}

// Add adds the specified element to this list if it is not already present.
func (list ImmutableVector[T]) Add(e T) bool {
	panic(errors.UnsupportedOperation("Add", "ImmutableVector"))
}

// AddAll adds all of the elements in the specified iterable to the list.
func (list ImmutableVector[T]) AddAll(iterable collections.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("AddAll", "ImmutableVector"))
}

// Clear removes all of the elements from the list.
func (list ImmutableVector[T]) Clear() {
	panic(errors.UnsupportedOperation("Clear", "ImmutableVector"))
}

// AddSlice adds all the elements in the slice to the list.
func (list ImmutableVector[T]) AddSlice(s []T) bool {
	panic(errors.UnsupportedOperation("AddSlice", "ImmutableVector"))
}
