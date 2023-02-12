package linkedhashset

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
)

// ImmutableLinkedHashSet an immutable version of [LinkedHashSet].
type ImmutableLinkedHashSet[T comparable] struct {
	linkedHashSet LinkedHashSet[T]
}

// ImmutableOf creates an ImmutableLinkedHashSet with the specified elements.
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
