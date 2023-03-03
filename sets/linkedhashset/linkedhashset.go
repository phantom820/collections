package linkedhashset

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/maps/linkedhashmap"
)

// LinkedHashSet implementation of a set backed by a LinkedHashMap.
type LinkedHashSet[T comparable] struct {
	linkedHashMap *linkedhashmap.LinkedHashMap[T, struct{}]
}

// New creates an empty set.
func New[T comparable]() *LinkedHashSet[T] {
	return &LinkedHashSet[T]{linkedhashmap.New[T, struct{}]()}
}

// Of creates a set with the given elements.
func Of[T comparable](elements ...T) LinkedHashSet[T] {
	set := LinkedHashSet[T]{linkedhashmap.New[T, struct{}]()}
	for i := range elements {
		set.Add(elements[i])
	}
	return set
}

// Add adds the specified element to the set if it is not already present.
func (set LinkedHashSet[T]) Add(e T) bool {
	return set.linkedHashMap.PutIfAbsent(e, struct{}{})
}

// AddAll adds all of the elements in the specified iterable to the set.
func (set LinkedHashSet[T]) AddAll(iterable collections.Iterable[T]) bool {
	n := set.linkedHashMap.Len()
	it := iterable.Iterator()
	for it.HasNext() {
		set.Add(it.Next())
	}
	return n != set.linkedHashMap.Len()
}

// AddSlice adds all the elements in the slice to the set.
func (set LinkedHashSet[T]) AddSlice(s []T) bool {
	n := set.linkedHashMap.Len()
	for _, value := range s {
		set.Add(value)
	}
	return n != set.linkedHashMap.Len()
}

// Remove removes the specified element from this set if it is present.
func (set LinkedHashSet[T]) Remove(e T) bool {
	n := set.Len()
	_ = set.linkedHashMap.Remove(e)
	return n != set.Len()
}

// RemoveIf removes all of the elements of this collection that satisfy the given predicate.
func (set LinkedHashSet[T]) RemoveIf(f func(T) bool) bool {
	n := set.linkedHashMap.Len()
	set.linkedHashMap.RemoveIf(f)
	return n != set.linkedHashMap.Len()
}

// RetainAll retains only the elements in the set that are contained in the specified collection.
func (set LinkedHashSet[T]) RetainAll(c collections.Collection[T]) bool {
	return set.RemoveIf(func(e T) bool { return !c.Contains(e) })
}

// RemoveAll removes all of the set's elements that are also contained in the specified iterable.
func (set LinkedHashSet[T]) RemoveAll(iterable collections.Iterable[T]) bool {
	n := set.linkedHashMap.Len()
	it := iterable.Iterator()
	for it.HasNext() {
		set.Remove(it.Next())
	}
	return n != set.linkedHashMap.Len()
}

// RemoveSlice removes all of the set's elements that are also contained in the specified slice.
func (set LinkedHashSet[T]) RemoveSlice(s []T) bool {
	n := set.linkedHashMap.Len()
	for i := range s {
		set.Remove(s[i])
	}
	return n != set.linkedHashMap.Len()
}

// Clear removes all of the elements from the set.
func (set LinkedHashSet[T]) Clear() {
	set.linkedHashMap.Clear()
}

// Contains returns true if this set contains the specified element.
func (set LinkedHashSet[T]) Contains(e T) bool {
	return set.linkedHashMap.ContainsKey(e)
}

// Len returns the number of elements in the set.
func (set LinkedHashSet[T]) Len() int {
	return set.linkedHashMap.Len()
}

// Empty returns true if the set contains no elements.
func (set LinkedHashSet[T]) Empty() bool {
	return set.linkedHashMap.Len() == 0
}

// Equals returns true if the set is equivalent to the given set. Two sets are equal if they are the same reference or have the same size and contain
// the same elements.
func (set *LinkedHashSet[T]) Equals(otherSet *LinkedHashSet[T]) bool {
	if set == otherSet {
		return true
	} else if set.Len() != otherSet.Len() {
		return false
	}
	it := set.Iterator()
	for it.HasNext() {
		if !otherSet.Contains(it.Next()) {
			return false
		}
	}
	return true
}

// ForEach performs the given action for each element of the set.
func (set LinkedHashSet[T]) ForEach(f func(T)) {
	it := set.linkedHashMap.Iterator()
	for it.HasNext() {
		f(it.Next().Key())
	}
}

// Iterator returns an iterator over the elements in the set.
func (set *LinkedHashSet[T]) Iterator() collections.Iterator[T] {
	return &iterator[T]{mapIterator: set.linkedHashMap.Iterator()}
}

// iterator implememantation for [LinkedHashSet].
type iterator[T comparable] struct {
	mapIterator maps.Iterator[T, struct{}]
}

// HasNext returns true if the iterator has more elements.
func (it *iterator[T]) HasNext() bool {
	return it.mapIterator.HasNext()
}

// Next returns the next element in the iterator.
func (it iterator[T]) Next() T {
	return it.mapIterator.Next().Key()
}

// ToSlice returns a slice containing all the elements in the set.
func (set *LinkedHashSet[T]) ToSlice() []T {
	slice := make([]T, set.Len())
	i := 0
	it := set.Iterator()
	for it.HasNext() {
		slice[i] = it.Next()
		i++
	}
	return slice
}

// ImmutableCopy returns an immutable copy of the set.
func (set *LinkedHashSet[T]) ImmutableCopy() ImmutableLinkedHashSet[T] {
	return ImmutableOf(set.ToSlice()...)
}

// String returns the string representation of a set.
func (set LinkedHashSet[T]) String() string {
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
