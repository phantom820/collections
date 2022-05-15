// Package linkedhashset provides an implementation of a set that is backed by a LinkedHashMap.
package linkedhashset

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/maps/linkedhashmap"
	"github.com/phantom820/collections/types"
)

// LinkedHashSet an implementation of a set based on a LinkedHashMap.
type LinkedHashSet[T types.Hashable[T]] struct {
	data *linkedhashmap.LinkedHashMap[T, bool]
}

// New creates a LinkedHashSet with the specified elements, if there none an empty set is returned.
func New[T types.Hashable[T]](elements ...T) *LinkedHashSet[T] {
	data := linkedhashmap.New[T, bool]()
	set := LinkedHashSet[T]{data: data}
	set.AddSlice(elements)
	return &set
}

// linkedHashSetIterator type to implement an iterator for a LinkedHashSet.
type linkedHashSetIterator[T types.Hashable[T]] struct {
	mapIterator maps.MapIterator[T, bool]
}

// HasNext checks if the iterator has a next element to yield.
func (iterator *linkedHashSetIterator[T]) HasNext() bool {
	return iterator.mapIterator.HasNext()
}

// Next returns the next element in the iterator it. Will panic if iterator has no next element.
func (iter *linkedHashSetIterator[T]) Next() T {
	if !iter.HasNext() {
		panic(iterator.NoNextElementError)
	}
	entry := iter.mapIterator.Next()
	return entry.Key
}

// Cycle resets the iterator.
func (iterator *linkedHashSetIterator[T]) Cycle() {
	iterator.mapIterator.Cycle()
}

// Iterator returns an iterator for the set.
func (set *LinkedHashSet[T]) Iterator() iterator.Iterator[T] {
	return &linkedHashSetIterator[T]{set.data.Iterator()}
}

// String formats the set for pretty printing.
func (set *LinkedHashSet[T]) String() string {
	sb := make([]string, 0, set.data.Len())
	for _, k := range set.data.Keys() {
		sb = append(sb, fmt.Sprint(k))

	}
	return "{" + strings.Join(sb, ", ") + "}"
}

// Len returns the size of the set.
func (set *LinkedHashSet[T]) Len() int {
	return set.data.Len()
}

// Contains checks if an element is in the set.
func (set *LinkedHashSet[T]) Contains(element T) bool {
	_, ok := set.data.Get(element)
	return ok
}

// Add adds elements  if not already in the set. Returns true if the set changed as a result of this call false otherwise.
func (set *LinkedHashSet[T]) Add(elements ...T) bool {
	ok := false
	for _, element := range elements {
		ok = set.data.PutIfAbsent(element, true)
	}
	return ok
}

// AddAll adds all elements from an iterable to the set.
func (set *LinkedHashSet[T]) AddAll(iterable iterator.Iterable[T]) {
	iterator := iterable.Iterator()
	for iterator.HasNext() {
		set.Add(iterator.Next())
	}
}

// AddSlice adds element from a slice to the set.
func (set *LinkedHashSet[T]) AddSlice(slice []T) {
	for _, element := range slice {
		set.Add(element)
	}
}

// Remove removes the element from the set if it is present.
func (set *LinkedHashSet[T]) Remove(e T) bool {
	_, ok := set.data.Remove(e)
	return ok
}

// RemoveAll removes all entries from an iterable from the set.
func (set *LinkedHashSet[T]) RemoveAll(iterable iterator.Iterable[T]) {
	set.data.RemoveAll(iterable)
}

// RetainAll removes all entries from the set that do not appear in the other collection. Returns true if the set was modified.
func (set *LinkedHashSet[T]) RetainAll(collection collections.Collection[T]) bool {
	iterator := set.Iterator()
	changed := false
	for iterator.HasNext() {
		element := iterator.Next()
		if collection.Contains(element) {
			continue
		} else {
			set.Remove(element)
			changed = true
		}
	}
	return changed
}

// Clear removes all elements in the set.
func (set *LinkedHashSet[T]) Clear() {
	set.data.Clear()
}

// Empty checks if the set is empty.
func (set *LinkedHashSet[T]) Empty() bool {
	return set.data.Empty()
}

// Collect collects all elements of the set into a slice.
func (set *LinkedHashSet[T]) Collect() []T {
	data := make([]T, set.data.Len())
	i := 0
	for _, e := range set.data.Keys() {
		data[i] = e
		i += 1
	}
	return data
}

// Map applies a transformation on elements of the set using the function f and returns a new set with transformed element.
func (set *LinkedHashSet[T]) Map(f func(e T) T) *LinkedHashSet[T] {
	newSet := New[T]()
	for _, element := range set.data.Keys() { // Should we use the iterator here ??
		newSet.Add(f(element))
	}
	return newSet
}

// Filter filters the set using the predicate function f and returns a new set with elements satisfying the predicate.
func (set *LinkedHashSet[T]) Filter(f func(e T) bool) *LinkedHashSet[T] {
	newSet := New[T]()
	for _, element := range set.data.Keys() {
		if f(element) {
			newSet.Add(element)
		}
	}
	return newSet
}

// Union union operation on sets a and b. Will return a new set.
func (a *LinkedHashSet[T]) Union(b *LinkedHashSet[T]) *LinkedHashSet[T] {
	c := New[T]()
	c.AddAll(a)
	c.AddAll(b)
	return c
}

// intersection helper function to perform set intersection the idea is iterate over bigger set and lookup in smaller.
func intersection[T types.Hashable[T]](a *LinkedHashSet[T], b *LinkedHashSet[T]) *LinkedHashSet[T] {
	c := New[T]()
	if a.Len() > b.Len() {
		it := a.Iterator()
		for it.HasNext() {
			e := it.Next()
			if b.Contains(e) {
				c.Add(e)
			}
		}
		return c
	}
	it := b.Iterator()
	for it.HasNext() {
		e := it.Next()
		if a.Contains(e) {
			c.Add(e)
		}
	}
	return c
}

// Intersection intersection operation on sets a and b. Will return a new set.
func (a *LinkedHashSet[T]) Intersection(b *LinkedHashSet[T]) *LinkedHashSet[T] {
	c := New[T]()
	if a.Empty() || b.Empty() {
		return c
	}
	return intersection(a, b)
}

// Equals check if the set is equals the other set. This is true only if they are the same reference or they are of the same size with the
// same elements.
func (set *LinkedHashSet[T]) Equals(other *LinkedHashSet[T]) bool {
	if set == other {
		return true
	} else if set.Len() != other.Len() {
		return false
	} else {
		it := set.Iterator()
		for it.HasNext() {
			if !other.Contains(it.Next()) {
				return false
			}
		}
		return true
	}
}
