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

// ImmutableOf creates an immutable list with the specified elements.
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

// IndexOf returns the index of the first occurrence of the specified element in the list or -1 if the list does not contain the element.
func (list ImmutableVector[T]) IndexOf(e T) int {
	return list.vector.IndexOf(e)
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
// the same elements in the same order.
func (list ImmutableVector[T]) Equals(otherList collections.List[T]) bool {
	return list.vector.Equals(otherList)
}

// ToSlice returns a slice containing all the elements in the list.
func (list ImmutableVector[T]) ToSlice() []T {
	return list.vector.ToSlice()
}

// SubList returns a view of the portion of the list between the specified start and end indices (exclusive).
func (list ImmutableVector[T]) SubList(start int, end int) ImmutableVector[T] {
	if start < 0 || start >= list.Len() {
		panic(errors.IndexOutOfBounds(start, list.Len()))
	} else if end < 0 || end > list.Len() {
		panic(errors.IndexOutOfBounds(end, list.Len()))
	} else if start > end {
		panic(errors.IndexBoundsOutOfRange(start, end))
	}
	return ImmutableVector[T]{Vector[T]{data: list.vector.data[start:end]}}
}

// String returns the string representation of the list.
func (list ImmutableVector[T]) String() string {
	return fmt.Sprint(list.vector.data)
}

// Remove unsupported operation.
func (list ImmutableVector[T]) Remove(e T) bool {
	panic(errors.UnsupportedOperation("Remove", "ImmutableVector"))
}

// RemoveAt unsupported operation.
func (list ImmutableVector[T]) RemoveAt(i int) T {
	panic(errors.UnsupportedOperation("RemoveAt", "ImmutableVector"))
}

// Set unsupported operation.
func (list ImmutableVector[T]) Set(i int, e T) T {
	panic(errors.UnsupportedOperation("Set", "ImmutableVector"))
}

// RemoveIf unsupported operation.
func (list ImmutableVector[T]) RemoveIf(func(T) bool) bool {
	panic(errors.UnsupportedOperation("RemoveIf", "ImmutableVector"))
}

// RemoveAll unsupported operation.
func (list ImmutableVector[T]) RemoveAll(iterable collections.Iterable[T]) bool {
	panic("RemoveAll")
}

// RemoveSlice unsupported operation.
func (list ImmutableVector[T]) RemoveSlice(s []T) bool {
	panic(errors.UnsupportedOperation("RemoveSlice", "ImmutableVector"))
}

// RetainAll unsupported operation.
func (list ImmutableVector[T]) RetainAll(c collections.Collection[T]) bool {
	panic(errors.UnsupportedOperation("RetainAll", "ImmutableVector"))
}

// Add unsupported operation.
func (list ImmutableVector[T]) Add(e T) bool {
	panic(errors.UnsupportedOperation("Add", "ImmutableVector"))
}

// AddAt (unsupported operation).
func (list ImmutableVector[T]) AddAt(i int, e T) {
	panic(errors.UnsupportedOperation("AddAt", "ImmutableVector"))
}

// AddAll unsupported operation.
func (list ImmutableVector[T]) AddAll(iterable collections.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("AddAll", "ImmutableVector"))
}

// Clear unsupported operation.
func (list ImmutableVector[T]) Clear() {
	panic(errors.UnsupportedOperation("Clear", "ImmutableVector"))
}

// AddSlice unsupported operation.
func (list ImmutableVector[T]) AddSlice(s []T) bool {
	panic(errors.UnsupportedOperation("AddSlice", "ImmutableVector"))
}

func (list ImmutableVector[T]) Sort(less func(a, b T) bool) {
	panic(errors.UnsupportedOperation("Sort", "ImmutableVector"))
}
