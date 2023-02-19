package hashset

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
)

// ImmutableHashSet an immutable version of [HashSet].
type ImmutableHashSet[T comparable] struct {
	hashSet HashSet[T]
}

// ImmutableOf creates an ImmutableHashSet with the specified elements.
func ImmutableOf[T comparable](elements ...T) ImmutableHashSet[T] {
	return ImmutableHashSet[T]{Of(elements...)}
}

// Contains returns true if this set contains the specified element.
func (set ImmutableHashSet[T]) Contains(e T) bool {
	return set.hashSet.Contains(e)
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
func (set ImmutableHashSet[T]) Iterator() collections.Iterator[T] {
	return set.hashSet.Iterator()
}

// Equals returns true if the set is equivalent to the given set. Two sets are equal if they are the same reference or have the same size and contain
// the same elements.
func (set ImmutableHashSet[T]) Equals(otherSet ImmutableHashSet[T]) bool {
	return set.hashSet.Equals(&otherSet.hashSet)
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

func (set ImmutableHashSet[T]) Remove(e T) bool {
	panic("")
}
func (set ImmutableHashSet[T]) RemoveIf(func(T) bool) bool {
	panic("")
}

func (set ImmutableHashSet[T]) RemoveAll(iterable collections.Iterable[T]) bool {
	panic("")
}

// RetainAll retains only the elements in the set that are contained in the specified collection.
func (set ImmutableHashSet[T]) RetainAll(c collections.Collection[T]) bool {
	panic("")
}

// Add adds the specified element to this set if it is not already present.
func (set ImmutableHashSet[T]) Add(e T) bool {
	panic("")
}

// AddAll adds all of the elements in the specified iterable to the set.
func (set ImmutableHashSet[T]) AddAll(iterable collections.Iterable[T]) bool {
	panic("")
}

// Clear removes all of the elements from the set.
func (set ImmutableHashSet[T]) Clear() {
	panic("")
}

// AddSlice adds all the elements in the slice to the set.
func (set ImmutableHashSet[T]) AddSlice(s []T) bool {
	panic("")
}
