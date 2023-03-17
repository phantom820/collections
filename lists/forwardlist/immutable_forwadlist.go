package forwardlist

import (
	"fmt"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types/optional"
)

// ImmutableForwardList an immutable version of [ForwadList].
type ImmutableForwadList[T comparable] struct {
	list ForwardList[T]
}

// Of creates a list with the given elements.
func ImmutableOf[T comparable](elements ...T) ImmutableForwadList[T] {
	list := ImmutableForwadList[T]{Of(elements...)}
	return list
}

// AddSlice unsupported operation.
func (list ImmutableForwadList[T]) AddSlice(s []T) bool {
	panic(errors.UnsupportedOperation("AddSlice", "ImmutableForwardList"))
}

// AddAll unsupported operation.
func (list ImmutableForwadList[T]) AddAll(iterable iterable.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("AddAll", "ImmutableForwardList"))
}

// Empty returns true if the list contains no elements.
func (list ImmutableForwadList[T]) Empty() bool {
	return list.list.Empty()
}

// Add unsupported operation.
func (list ImmutableForwadList[T]) Add(e T) bool {
	panic(errors.UnsupportedOperation("Add", "ImmutableForwardList"))
}

// AddAt unsupported operation.
func (list ImmutableForwadList[T]) AddAt(i int, e T) {
	panic(errors.UnsupportedOperation("AddAt", "ImmutableForwardList"))
}

// At returns the element at the specified index in the list.
func (list ImmutableForwadList[T]) At(i int) T {
	return list.list.At(i)
}

// IndexOf returns the index of the first occurrence of the specified element in the list.
func (list ImmutableForwadList[T]) IndexOf(e T) optional.Optional[int] {
	return list.list.IndexOf(e)
}

// Set unsupported operation.
func (list ImmutableForwadList[T]) Set(i int, e T) T {
	panic(errors.UnsupportedOperation("Set", "ImmutableForwardList"))
}

// Len returns the number of elements in the list.
func (list ImmutableForwadList[T]) Len() int {
	return list.list.len
}

// Clear unsupported operation.
func (list ImmutableForwadList[T]) Clear() {
	panic(errors.UnsupportedOperation("Clear", "ImmutableForwardList"))
}

// Contains returns true if the list contains the specified element.
func (list ImmutableForwadList[T]) Contains(e T) bool {
	return list.list.Contains(e)
}

// Remove unsupported operation.
func (list ImmutableForwadList[T]) Remove(e T) bool {
	panic(errors.UnsupportedOperation("Remove", "ImmutableForwardList"))
}

// RemoveAt unsupported operation.
func (list ImmutableForwadList[T]) RemoveAt(i int) T {
	panic(errors.UnsupportedOperation("RemoveAt", "ImmutableForwardList"))
}

// RemoveIf unsupported operation.
func (list ImmutableForwadList[T]) RemoveIf(f func(T) bool) bool {
	panic(errors.UnsupportedOperation("RemoveIf", "ImmutableForwardList"))
}

// RemoveAll unsupported operation
func (list ImmutableForwadList[T]) RemoveAll(iterable iterable.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("RemoveAll", "ImmutableForwardList"))
}

// RetainAll unsupported operation.
func (list ImmutableForwadList[T]) RetainAll(c collections.Collection[T]) bool {
	panic(errors.UnsupportedOperation("RetainAll", "ImmutableForwardList"))
}

// RemoveSlice unsupported operation.
func (list ImmutableForwadList[T]) RemoveSlice(s []T) bool {
	panic(errors.UnsupportedOperation("RemoveSlice", "ImmutableForwardList"))
}

// ToSlice returns a slice containing the elements of the list.
func (list ImmutableForwadList[T]) ToSlice() []T {
	return list.list.ToSlice()
}

// ForEach performs the given action for each element of the list.
func (list ImmutableForwadList[T]) ForEach(f func(T)) {
	list.list.ForEach(f)
}

// SubList returns a view of the portion of the list between the specified start and end indices (exclusive).
func (list ImmutableForwadList[T]) SubList(start int, end int) ImmutableForwadList[T] {
	if start < 0 || start >= list.Len() {
		panic(errors.IndexOutOfBounds(start, list.Len()))
	} else if end < 0 || end > list.Len() {
		panic(errors.IndexOutOfBounds(end, list.Len()))
	} else if start > end {
		panic(errors.IndexBoundsOutOfRange(start, end))
	} else if start == end {
		return ImmutableOf[T]()
	}
	_, startNode := chaseIndex(list.list.head, start)
	return ImmutableForwadList[T]{list: ForwardList[T]{head: startNode, tail: list.list.tail, len: end - start}}
}

// Equals returns true if the list is equivalent to the given list. Two lists are equal if they have the same size
// and contain the same elements in the same order.
func (list ImmutableForwadList[T]) Equals(other collections.List[T]) bool {
	return list.list.Equals(other)
}

// Iterator returns an iterator over the elements in the list.
func (list ImmutableForwadList[T]) Iterator() iterator.Iterator[T] {
	return list.list.Iterator()
}

// String returns the string representation of the list.
func (list ImmutableForwadList[T]) String() string {
	return fmt.Sprint(list.ToSlice())
}

// Sort unsupported operation.
func (list ImmutableForwadList[T]) Sort(less func(a, b T) bool) {
	panic(errors.UnsupportedOperation("Sort", "ImmutableForwardList"))
}
