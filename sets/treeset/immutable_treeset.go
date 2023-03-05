package treeset

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
)

// ImmutableTreeSet an immutable version of [TreeSet].
type ImmutableTreeSet[T comparable] struct {
	treeSet TreeSet[T]
}

// ImmutableOf creates an immutable set with the specified elements. Elements are compared using the lessThan function which should satisfy.
// e1 < e2 => lessThan(e1, e2) = true and lessThan(e2,e1) = false.
// e1 = e2 => lessThan(e1,e2) = false and lessThan(e2,e1) = false.
// e1 > e2 -> lessThan(e1,e2) = false and lessThan(e2,e1) = true.
func ImmutableOf[T comparable](lessThan func(e1, e2 T) bool, elements ...T) ImmutableTreeSet[T] {
	return ImmutableTreeSet[T]{Of(lessThan, elements...)}
}

// Contains returns true if this set contains the specified element.
func (set ImmutableTreeSet[T]) Contains(e T) bool {
	return set.treeSet.Contains(e)
}

// Len returns the number of elements in the set.
func (set ImmutableTreeSet[T]) Len() int {
	return set.treeSet.Len()
}

// Empty returns true if the set contains no elements.
func (set ImmutableTreeSet[T]) Empty() bool {
	return set.treeSet.Empty()
}

// ForEach performs the given action for each element of the set.
func (set ImmutableTreeSet[T]) ForEach(f func(T)) {
	set.treeSet.ForEach(f)
}

// Iterator returns an iterator over the elements in the set.
func (set ImmutableTreeSet[T]) Iterator() collections.Iterator[T] {
	return set.treeSet.Iterator()
}

// ToSlice returns a slice containing all the elements in the set.
func (set ImmutableTreeSet[T]) ToSlice() []T {
	return set.treeSet.ToSlice()
}

// Equals returns true if the set is equivalent to the given set. Two sets are equal if they are the same reference or have the same size and contain
// the same elements.
func (set ImmutableTreeSet[T]) Equals(otherSet ImmutableTreeSet[T]) bool {
	return set.treeSet.Equals(&otherSet.treeSet)
}

// String returns the string representation of the set.
func (set ImmutableTreeSet[T]) String() string {
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
func (set ImmutableTreeSet[T]) Remove(e T) bool {
	panic(errors.UnsupportedOperation("Remove", "ImmutableTreeSet"))
}

// RemoveIf unsupported operation.
func (set ImmutableTreeSet[T]) RemoveIf(func(T) bool) bool {
	panic(errors.UnsupportedOperation("RemoveIf", "ImmutableTreeSet"))
}

// RemoveAll unsupported operation.
func (set ImmutableTreeSet[T]) RemoveAll(iterable collections.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("RemoveAll", "ImmutableTreeSet"))
}

// RemoveSlice unsupported operation.
func (set ImmutableTreeSet[T]) RemoveSlice(s []T) bool {
	panic(errors.UnsupportedOperation("RemoveSlice", "ImmutableTreeSet"))
}

// RetainAll unsupported operation.
func (set ImmutableTreeSet[T]) RetainAll(c collections.Collection[T]) bool {
	panic(errors.UnsupportedOperation("RetainAll", "ImmutableTreeSet"))
}

// Add unsupported operation.
func (set ImmutableTreeSet[T]) Add(e T) bool {
	panic(errors.UnsupportedOperation("Add", "ImmutableTreeSet"))
}

// AddAll unsupported operation
func (set ImmutableTreeSet[T]) AddAll(iterable collections.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("AddAll", "ImmutableTreeSet"))
}

// Clear unsupported operation.
func (set ImmutableTreeSet[T]) Clear() {
	panic(errors.UnsupportedOperation("Clear", "ImmutableTreeSet"))
}

// AddSlice unsupported operation
func (set ImmutableTreeSet[T]) AddSlice(s []T) bool {
	panic(errors.UnsupportedOperation("AddSlice", "ImmutableTreeSet"))
}
