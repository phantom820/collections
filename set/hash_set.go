package set

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	_map "github.com/phantom820/collections/map"
	"github.com/phantom820/collections/types"

	"strings"
)

// HashSet an implementation of a hashset based on a HashMap.
type HashSet[T types.Hashable[T]] struct {
	data *_map.HashMap[T, bool]
}

// NewHashSet creates a HashSet with the specified elements, if there none an empty set is returned.
func NewHashSet[T types.Hashable[T]](elements ...T) *HashSet[T] {
	data := _map.NewHashMap[T, bool]()
	s := HashSet[T]{data: data}
	s.AddSlice(elements)
	return &s
}

// hashSetIterator concrete type to implement an iterator for a hashSet.
type hashSetIterator[T types.Hashable[T]] struct {
	_it _map.MapIterator[T, bool]
}

// HasNext check if the iterator it has a next element to yield.
func (it *hashSetIterator[T]) HasNext() bool {
	return it._it.HasNext()
}

// Next yields the next element in the iterator.
func (it *hashSetIterator[T]) Next() T {
	entry := it._it.Next()
	return entry.Key()
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

// Add adds the element  e if its not already in the set s.
func (s *HashSet[T]) Add(e T) bool {
	p := s.data.PutIfAbsent(e, true)
	return p
}

// AddAll adds all elements from an iterable it to the set s.
func (s *HashSet[T]) AddAll(it iterator.Iterable[T]) {
	iter := it.Iterator()
	for iter.HasNext() {
		s.Add(iter.Next())
	}
}

// AddSlice adds element from some slice sl into the set s.
func (s *HashSet[T]) AddSlice(sl []T) {
	for _, e := range sl {
		s.Add(e)
	}
}

// Remove removes the element from the set s if it is present.
func (s *HashSet[T]) Remove(e T) bool {
	return s.data.Remove(e)
}

// RemoveAll removes all entries from some iterable it from set s.
func (s *HashSet[T]) RemoveAll(it iterator.Iterable[T]) {
	s.data.RemoveAll(it)
}

// Clear clears the set s by removing all elements.
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
	other := NewHashSet[T]()
	for _, e := range s.data.Keys() {
		other.Add(f(e))
	}
	return other
}

// Filter filters set s using a predicate function f and returns a new set with elements satisfying the predicate.
func (s *HashSet[T]) Filter(f func(e T) bool) *HashSet[T] {
	other := NewHashSet[T]()
	for _, e := range s.data.Keys() {
		if f(e) {
			other.Add(e)
		}
	}
	return other
}

// Equals check if the set s is equals the other set. This is true only if they are the same reference or they are of the same size with the
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
