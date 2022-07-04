// Package hashset provides an implementation of a set that is backed by a HashMap.
package hashset

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/types"
)

// HashSet an implementation of a hashset based on a HashMap.
type HashSet[T types.Hashable[T]] struct {
	data          *hashmap.HashMap[T, bool]
	modifications int
}

// New creates a HashSet with the specified elements.
func New[T types.Hashable[T]](elements ...T) *HashSet[T] {
	data := hashmap.New[T, bool]()
	set := HashSet[T]{data: data}
	set.Add(elements...)
	return &set
}

// modify increments the modification value
func (set *HashSet[T]) modify() {
	set.modifications++
}

// hashSetIterator type to implement an iterator for a HashSet.
type hashSetIterator[T types.Hashable[T]] struct {
	initialized      bool
	mapIterator      maps.MapIterator[T, bool]
	getMapIterator   func() maps.MapIterator[T, bool]
	modifications    int
	getModifications func() int
}

// HasNext checks if the iterator has a next element to yield.
func (it *hashSetIterator[T]) HasNext() bool {
	if !it.initialized {
		it.initialized = true
		it.modifications = it.getModifications()
		it.mapIterator = it.getMapIterator()
	}
	return it.mapIterator.HasNext()
}

// Next returns the next element in the iterator it. Will panic if iterator has no next element.
func (it *hashSetIterator[T]) Next() T {
	if !it.HasNext() {
		panic(errors.ErrNoNextElement())
	} else if it.modifications != it.getModifications() {
		panic(errors.ErrConcurrenModification())
	}
	entry := it.mapIterator.Next()
	return entry.Key
}

// Iterator returns an iterator for the set.
func (set *HashSet[T]) Iterator() iterator.Iterator[T] {
	return &hashSetIterator[T]{getMapIterator: set.data.Iterator, getModifications: func() int { return set.modifications }}
}

// String formats the set for pretty printing.
func (set *HashSet[T]) String() string {
	sb := make([]string, 0, set.data.Len())
	for _, k := range set.data.Keys() {
		sb = append(sb, fmt.Sprint(k))

	}
	return "{" + strings.Join(sb, ", ") + "}"
}

// Len returns the size of the set.
func (set *HashSet[T]) Len() int {
	return set.data.Len()
}

// Contains checks if an element is in the set.
func (set *HashSet[T]) Contains(element T) bool {
	_, ok := set.data.Get(element)
	return ok
}

// Add adds elements to the set. Only elements that are not in the set are added.
func (set *HashSet[T]) Add(elements ...T) bool {
	set.modify()
	n := set.Len()
	for _, element := range elements {
		set.data.PutIfAbsent(element, true)
	}
	return (n != set.Len())
}

// AddAll adds all elements from an iterable to the set. Only elements that are not in the set are added.
func (set *HashSet[T]) AddAll(iterable iterator.Iterable[T]) {
	it := iterable.Iterator()
	for it.HasNext() {
		set.Add(it.Next())
	}
}

// Remove removes elements from the set.
func (set *HashSet[T]) Remove(elements ...T) bool {
	set.modify()
	n := set.Len()
	for _, element := range elements {
		set.data.Remove(element)
		if set.Empty() {
			break
		}
	}
	return (n != set.Len())
}

// RemoveIf removes all elements from the set that satisfy the predicate function f.
func (set *HashSet[T]) RemoveIf(f func(element T) bool) bool {
	set.modify()
	n := set.Len()
	elements := set.Collect()
	for _, element := range elements {
		if f(element) {
			set.Remove(element)
		}
	}
	return n != set.Len()
}

// RemoveAll removes all the elements in the set that appear in the iterable.
func (set *HashSet[T]) RemoveAll(iterable iterator.Iterable[T]) {
	set.data.RemoveAll(iterable)
}

// RetainAll removes all entries from the set that do not appear in the other collection.
func (set *HashSet[T]) RetainAll(collection collections.Collection[T]) bool {
	set.modify()
	elements := set.Collect()
	n := set.Len()
	for _, element := range elements {
		if collection.Contains(element) {
			continue
		} else {
			set.Remove(element)
		}
	}
	return n != set.Len()
}

// Clear removes all elements from the set.
func (set *HashSet[T]) Clear() {
	set.modify()
	set.data.Clear()
}

// Empty checks if the set is empty.
func (set *HashSet[T]) Empty() bool {
	return set.data.Empty()
}

// Collect returns a slice containing all the elements in the set.
func (set *HashSet[T]) Collect() []T {
	data := make([]T, set.data.Len())
	i := 0
	for _, e := range set.data.Keys() {
		data[i] = e
		i += 1
	}
	return data
}

// Map applies a transformation on an elements of the set , using the function f and returns a new set with the
// transformed elements.
func (set *HashSet[T]) Map(f func(element T) T) *HashSet[T] {
	newSet := New[T]()
	for _, element := range set.data.Keys() { // Should we use the iterator here ??
		newSet.Add(f(element))
	}
	return newSet
}

// Filter filters the set using the predicate function  f and returns a new set containing only elements that satisfy the predicate.
func (set *HashSet[T]) Filter(f func(element T) bool) *HashSet[T] {
	newSet := New[T]()
	for _, element := range set.data.Keys() {
		if f(element) {
			newSet.Add(element)
		}
	}
	return newSet
}

// Union union operation on sets a and b. Will return a new set.
func (a *HashSet[T]) Union(b *HashSet[T]) *HashSet[T] {
	c := New[T]()
	c.AddAll(a)
	c.AddAll(b)
	return c
}

// intersection helper function to perform set intersection the idea is iterate over bigger set and lookup in smaller.
func intersection[T types.Hashable[T]](a *HashSet[T], b *HashSet[T]) *HashSet[T] {
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
func (a *HashSet[T]) Intersection(b *HashSet[T]) *HashSet[T] {
	c := New[T]()
	if a.Empty() || b.Empty() {
		return c
	}
	return intersection(a, b)
}

// Equals checks if the set is equal to another set. Two sets are equal if they are the same reference or have the same size and contain the same elements.
func (set *HashSet[T]) Equals(other *HashSet[T]) bool {
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
