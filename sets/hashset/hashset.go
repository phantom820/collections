// Package hashset provides an implementation of a HashSet that is backed by a HashMap.
package hashset

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/maps/hashmap"

	"github.com/phantom820/collections/types"

	"strings"
)

// HashSet an implementation of a hashset based on a HashMap.
type HashSet[T types.Hashable[T]] struct {
	data *hashmap.HashMap[T, bool]
}

// New creates a HashSet with the specified elements, if there none an empty set is returned.
func New[T types.Hashable[T]](elements ...T) *HashSet[T] {
	data := hashmap.New[T, bool]()
	s := HashSet[T]{data: data}
	s.AddSlice(elements)
	return &s
}

// hashSetIterator concrete type to implement an iterator for a hashSet.
type hashSetIterator[T types.Hashable[T]] struct {
	_it maps.MapIterator[T, bool]
}

// HasNext check if the iterator it has a next element to yield.
func (it *hashSetIterator[T]) HasNext() bool {
	return it._it.HasNext()
}

// Next yields the next element in the iterator.
func (it *hashSetIterator[T]) Next() T {
	entry := it._it.Next()
	return entry.Key
}

// Cycle resets the iterator.
func (it *hashSetIterator[T]) Cycle() {
	it._it.Cycle()
}

func (s HashSet[T]) Iterator() iterator.Iterator[T] {
	return &hashSetIterator[T]{s.data.Iterator()}
}

// String formats the set for pretty printing.
func (s HashSet[T]) String() string {
	sb := make([]string, 0, s.data.Len())
	for _, k := range s.data.Keys() {
		sb = append(sb, fmt.Sprint(k))

	}
	return "{" + strings.Join(sb, ", ") + "}"
}

// Len returns the size of the set.
func (s *HashSet[T]) Len() int {
	return s.data.Len()
}

// Contains checks if an element is in the set.
func (s *HashSet[T]) Contains(e T) bool {
	_, p := s.data.Get(e)
	return p
}

// Add adds elements  if not already in the set. Returns true if the set changed as a result of this call false otherwise
func (s *HashSet[T]) Add(elements ...T) bool {
	p := false
	for _, e := range elements {
		p = s.data.PutIfAbsent(e, true)
	}
	return p
}

// AddAll adds all elements from an iterable it to the set.
func (s *HashSet[T]) AddAll(it iterator.Iterable[T]) {
	iter := it.Iterator()
	for iter.HasNext() {
		s.Add(iter.Next())
	}
}

// AddSlice adds element from some slice sl into the set.
func (s *HashSet[T]) AddSlice(sl []T) {
	for _, e := range sl {
		s.Add(e)
	}
}

// Remove removes the element from the set if it is present.
func (s *HashSet[T]) Remove(e T) bool {
	_, ok := s.data.Remove(e)
	return ok
}

// RemoveAll removes all entries from some iterable it from set.
func (s *HashSet[T]) RemoveAll(it iterator.Iterable[T]) {
	s.data.RemoveAll(it)
}

// Clear clears the set by removing all elements.
func (s *HashSet[T]) Clear() {
	s.data.Clear()
}

// Empty checks if the set is empty.
func (s *HashSet[T]) Empty() bool {
	return s.data.Empty()
}

// Collect collects all elements of the set into a slice.
func (s *HashSet[T]) Collect() []T {
	data := make([]T, s.data.Len())
	i := 0
	for _, e := range s.data.Keys() {
		data[i] = e
		i += 1
	}
	return data
}

// Map applies a transformation on elements of using a function f and returns a new set with transformed elements.
func (s *HashSet[T]) Map(f func(e T) T) *HashSet[T] {
	other := New[T]()
	for _, e := range s.data.Keys() {
		other.Add(f(e))
	}
	return other
}

// Filter filters set using a predicate function f and returns a new set with elements satisfying the predicate.
func (s *HashSet[T]) Filter(f func(e T) bool) *HashSet[T] {
	other := New[T]()
	for _, e := range s.data.Keys() {
		if f(e) {
			other.Add(e)
		}
	}
	return other
}

// Union performs union operation using sets a and b. This will return a new set.
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

// Intersection performs intersection operation using sets a and b. This will return a new set.
func (a *HashSet[T]) Intersection(b *HashSet[T]) *HashSet[T] {
	c := New[T]()
	if a.Empty() || b.Empty() {
		return c
	}
	return intersection(a, b)
}

// Equals check if the set is equals the other set. This is true only if they are the same reference or they are of the same size with the
// same elements.
func (s *HashSet[T]) Equals(other *HashSet[T]) bool {
	if s == other {
		return true
	} else if s.Len() != other.Len() {
		return false
	} else {
		it := s.Iterator()
		for it.HasNext() {
			if !other.Contains(it.Next()) {
				return false
			}
		}
		return true
	}
}
