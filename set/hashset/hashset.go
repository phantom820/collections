package hashset

import (
	"collections/interfaces"
	_map "collections/map"
	"collections/map/hashmap"
	"collections/set"
	"errors"
	"fmt"

	"strings"
)

var (
	NoNextElementError = errors.New("Iterator has no next element.")
)

// set the actual underlying set.
type hashSet[T interfaces.Hashable[T]] struct {
	data hashmap.HashMap[T, bool]
}

// HashSet implements Set interface with some type T and all operations that result in a set
// will return a HashSet.
type HashSet[T interfaces.Hashable[T]] interface {
	set.Set[T, HashSet[T]]
	Equals(other HashSet[T]) bool
}

// NewHashSet creates a new empty set with default initial capacity(16) and load factor limit(0.75).
func NewHashSet[T interfaces.Hashable[T]]() HashSet[T] {
	data := hashmap.NewHashMap[T, bool]()
	s := hashSet[T]{data: data}
	return &s
}

type hashSetIterator[T interfaces.Hashable[T]] struct {
	_it _map.MapIterator[T, bool]
}

func (it *hashSetIterator[T]) HasNext() bool {
	if it._it.HasNext() {
		return true
	}
	return false
}

func (it *hashSetIterator[T]) Next() T {
	entry := it._it.Next()
	return entry.Key()
}

func (it *hashSetIterator[T]) Cycle() {
	it._it.Cycle()
}

func (s hashSet[T]) Iterator() interfaces.Iterator[T] {
	return &hashSetIterator[T]{s.data.Iterator()}
}

// String formats the set for pretty printing.
func (s hashSet[T]) String() string {
	sb := make([]string, 0, s.data.Len())
	for _, k := range s.data.Keys() {
		sb = append(sb, fmt.Sprint(k))

	}
	return "{" + strings.Join(sb, ", ") + "}"
}

// Len returns the size of the set.
func (s *hashSet[T]) Len() int {
	return s.data.Len()
}

// Contains checks if an element is in the set.
func (s *hashSet[T]) Contains(e T) bool {
	_, p := s.data.Get(e)
	return p
}

// Add adds the element  e if its not already in the set s.
func (s *hashSet[T]) Add(e T) bool {
	p := s.data.PutIfAbsent(e, true)
	return p
}

// AddAll adds all elements from an iterable it to the set s.
func (s *hashSet[T]) AddAll(it interfaces.Iterable[T]) {
	iter := it.Iterator()
	for iter.HasNext() {
		s.Add(iter.Next())
	}
}

// Remove removes the element from the set s if it is present.
func (s *hashSet[T]) Remove(e T) bool {
	return s.data.Remove(e)
}

// RemoveAll removes all entries from some iterable it from set s.
func (s *hashSet[T]) RemoveAll(it interfaces.Iterable[T]) {
	s.data.RemoveAll(it)
}

// Clear clears the set s by removing all elements.
func (s *hashSet[T]) Clear() {
	s.data.Clear()
}

// Empty checks if the set is empty.
func (s *hashSet[T]) Empty() bool {
	return s.data.Empty()
}

// Collect collects all elements of the set into a slice.
func (s *hashSet[T]) Collect() []T {
	data := make([]T, s.data.Len())
	i := 0
	for _, e := range s.data.Keys() {
		data[i] = e
		i += 1
	}
	return data
}

// Map applies some transformation on elements of the set s to produce a new set.
func (s *hashSet[T]) Map(f func(e T) T) HashSet[T] {
	other := NewHashSet[T]()
	for _, e := range s.data.Keys() {
		other.Add(f(e))
	}
	return other
}

// Filter filters this set and produces a new set.
func (s *hashSet[T]) Filter(f func(e T) bool) HashSet[T] {
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
func (s *hashSet[T]) Equals(other HashSet[T]) bool {
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
