package hashset

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/iterator"
)

// ImmutableHashSet an immutable version of [HashSet].
type ImmutableHashSet[T comparable] struct {
	hashSet HashSet[T]
}

// Contains returns true if this set contains the specified element.
func (set ImmutableHashSet[T]) Contains(e T) bool {
	return set.hashSet.Contains(e)
}

// ContainsAll returns true if the set contains all of the elements of the specified iterable.
func (set ImmutableHashSet[T]) ContainsAll(iterable iterable.Iterable[T]) bool {
	return set.hashSet.ContainsAll(iterable)
}

// Len returns the number of elements in the set.
func (set ImmutableHashSet[T]) Len() int {
	return set.hashSet.Len()
}

// Empty returns true if the set contains no elements.
func (set ImmutableHashSet[T]) Empty() bool {
	return set.hashSet.Empty()
}

// ForEach performs the given action for each element of the set.
func (set ImmutableHashSet[T]) ForEach(f func(T)) {
	set.hashSet.ForEach(f)
}

// Iterator returns an iterator over the elements in the set.
func (set ImmutableHashSet[T]) Iterator() iterator.Iterator[T] {
	return set.hashSet.Iterator()
}

// Equals returns true if the set is equivalent to the given set. Two sets are equal if they are the same reference or have the same size and contain
// the same elements.
func (set ImmutableHashSet[T]) Equals(otherSet collections.Set[T]) bool {
	return set.hashSet.Equals(otherSet)
}

// ToSlice returns a slice containing all the elements in the set.
func (set ImmutableHashSet[T]) ToSlice() []T {
	return set.hashSet.ToSlice()
}

// String returns the string representation of the set.
func (set ImmutableHashSet[T]) String() string {
	var sb strings.Builder
	if set.Empty() {
		return "{}"
	}
	sb.WriteString("{")
	it := set.Iterator()
	sb.WriteString(fmt.Sprint(it.Next()))
	for it.HasNext() {
		sb.WriteString(fmt.Sprintf(" ,%v", it.Next()))
	}
	sb.WriteString("}")
	return sb.String()
}

// Remove unsupported operation.
func (set ImmutableHashSet[T]) Remove(e T) bool {
	panic(errors.UnsupportedOperation("Remove", "ImmutableHashSet"))
}

// RemoveIf unsupported operation.
func (set ImmutableHashSet[T]) RemoveIf(func(T) bool) bool {
	panic(errors.UnsupportedOperation("RemoveIf", "ImmutableHashSet"))
}

// RemoveAll unsupported operation.
func (set ImmutableHashSet[T]) RemoveAll(iterable iterable.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("RemoveAll", "ImmutableHashSet"))
}

// RemoveSlice unsupported operation.
func (set ImmutableHashSet[T]) RemoveSlice(s []T) bool {
	panic(errors.UnsupportedOperation("RemoveSlice", "ImmutableHashSet"))
}

// RetainAll unsupported operation.
func (set ImmutableHashSet[T]) RetainAll(c collections.Collection[T]) bool {
	panic(errors.UnsupportedOperation("RetainAll", "ImmutableHashSet"))
}

// Add unsupported operation.
func (set ImmutableHashSet[T]) Add(e T) bool {
	panic(errors.UnsupportedOperation("Add", "ImmutableHashSet"))
}

// AddAll unsupported operation.
func (set ImmutableHashSet[T]) AddAll(iterable iterable.Iterable[T]) bool {
	panic(errors.UnsupportedOperation("AddAll", "ImmutableHashSet"))
}

// Clear unsupported operation.
func (set ImmutableHashSet[T]) Clear() {
	panic(errors.UnsupportedOperation("Clear", "ImmutableHashSet"))
}

// AddSlice unsupported operation.
func (set ImmutableHashSet[T]) AddSlice(s []T) bool {
	panic(errors.UnsupportedOperation("AddSlice", "ImmutableHashSet"))
}
