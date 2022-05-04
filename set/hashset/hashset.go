package hashset

import (
	"collections/interfaces"
	_map "collections/map"
	"collections/map/hashmap"
	"collections/set"
	"fmt"

	"strings"
)

// set the actual underlying set.
type hashSet[T interfaces.Hashable[T]] struct {
	data hashmap.HashMap[T, bool]
	len  int
}

// HashSet implements Set interface with some type T and all operations that result in a set
// will return a HashSet.
type HashSet[T interfaces.Hashable[T]] interface {
	set.Set[T, HashSet[T]]
}

// NewHashSet creates a new empty set with default initial capacity and load factor limit.
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

// String formats the set for printing.
func (s hashSet[T]) String() string {
	sb := make([]string, 0, s.len)
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
	p := s.data.Put(e, true)
	return p
}

// AddAll adds all elements from an iterable to the set.
func (a *hashSet[T]) AddAll(b interfaces.Iterable[T]) {
	for _, e := range b.Collect() {
		a.Add(e)
	}
}

// Remove remove the element from the set if it is present.
func (s *hashSet[T]) Remove(e T) bool {
	return s.data.Remove(e)
}

// RemoveAll deletes all entries of set b from set a.
func (s *hashSet[T]) RemoveAll(b interfaces.Iterable[T]) {
	s.data.RemoveAll(b)
}

// Clear clears the set and returns to its initial state.
func (s *hashSet[T]) Clear() {
	s.data.Clear()
}

// Empty checks if the set is empty.
func (s *hashSet[T]) Empty() bool {
	return s.len == 0
}

// Collect collects all elements of the set into a slice.
func (s *hashSet[T]) Collect() []T {
	data := make([]T, s.len)
	i := 0
	for _, e := range s.data.Keys() {
		data[i] = e
		i += 1
	}
	return data
}

// Map applies some transformation on elements of the set to produce a new set.
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
