package linkedhashset

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
)

// ImmutableLinkedHashSet an immutable version of [LinkedHashSet].
type ImmutableLinkedHashSet[T comparable] struct {
	linkedHashSet LinkedHashSet[T]
}

// ImmutableOf creates an immutable set with the specified elements.
func ImmutableOf[T comparable](elements ...T) ImmutableLinkedHashSet[T] {
	return ImmutableLinkedHashSet[T]{Of(elements...)}
}

// Contains returns true if this set contains the specified element.
func (set ImmutableLinkedHashSet[T]) Contains(e T) bool {
	return set.linkedHashSet.Contains(e)
}

// Len returns the number of elements in the set.
func (set ImmutableLinkedHashSet[T]) Len() int {
	return set.linkedHashSet.Len()
}

// Empty returns true if the set contains no elements.
func (set ImmutableLinkedHashSet[T]) Empty() bool {
	return set.linkedHashSet.Empty()
}

// ForEach performs the given action for each element of the set.
func (set ImmutableLinkedHashSet[T]) ForEach(f func(T)) {
	set.linkedHashSet.ForEach(f)
}

// Iterator returns an iterator over the elements in the set.
func (set ImmutableLinkedHashSet[T]) Iterator() collections.Iterator[T] {
	return set.linkedHashSet.Iterator()
}

// ToSlice returns a slice containing all the elements in the set.
func (set ImmutableLinkedHashSet[T]) ToSlice() []T {
	return set.linkedHashSet.ToSlice()
}

// Equals returns true if the set is equivalent to the given set. Two sets are equal if they are the same reference or have the same size and contain
// the same elements.
func (set ImmutableLinkedHashSet[T]) Equals(otherSet ImmutableLinkedHashSet[T]) bool {
	return set.linkedHashSet.Equals(&otherSet.linkedHashSet)
}

// String returns the string representation of the set.
func (set ImmutableLinkedHashSet[T]) String() string {
	var sb strings.Builder
	if set.Empty() {
		return "{}"
	}
	sb.WriteString("{")
	it := set.Iterator()
	sb.WriteString(fmt.Sprint(it.Next()))
	for it.HasNext() {
		sb.WriteString(fmt.Sprintf(", %v", it.Next()))
	}
	sb.WriteString("}")
	return sb.String()
}

// Remove unsupported operation.
func (set ImmutableLinkedHashSet[T]) Remove(e T) bool {
	panic(errors.UnsupportedOperation("Remove", "ImmutableLinkedHashSet"))
}

// RemoveIf unsupported operation.
func (set ImmutableLinkedHashSet[T]) RemoveIf(func(T) bool) bool {
	panic(errors.UnsupportedOperation("RemoveIf", "ImmutableLinkedHashSet"))
}

// RemoveAll unsupported operation.
func (set ImmutableLinkedHashSet[T]) RemoveAll(iterable collections.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("RemoveAll", "ImmutableLinkedHashSet"))
}

// RemoveSlice unsupported operation.
func (set ImmutableLinkedHashSet[T]) RemoveSlice(s []T) bool {
	panic(errors.UnsupportedOperation("RemoveSlice", "ImmutableLinkedHashSet"))
}

// RetainAll unsupported operation.
func (set ImmutableLinkedHashSet[T]) RetainAll(c collections.Collection[T]) bool {
	panic(errors.UnsupportedOperation("RetainAll", "ImmutableLinkedHashSet"))
}

// Add unsupported operation.
func (set ImmutableLinkedHashSet[T]) Add(e T) bool {
	panic(errors.UnsupportedOperation("Add", "ImmutableLinkedHashSet"))
}

// AddAll unsupported operation
func (set ImmutableLinkedHashSet[T]) AddAll(iterable collections.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("AddAll", "ImmutableLinkedHashSet"))
}

// Clear unsupported operation.
func (set ImmutableLinkedHashSet[T]) Clear() {
	panic(errors.UnsupportedOperation("Clear", "ImmutableLinkedHashSet"))
}

// AddSlice unsupported operation
func (set ImmutableLinkedHashSet[T]) AddSlice(s []T) bool {
	panic(errors.UnsupportedOperation("AddSlice", "ImmutableLinkedHashSet"))
}
