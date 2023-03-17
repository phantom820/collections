package hashset

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/types/pair"
)

// HashSet implementation of a set backed by a HashMap.
type HashSet[T comparable] struct {
	hashmap hashmap.HashMap[T, struct{}]
}

// New creates an empty set.
func New[T comparable]() *HashSet[T] {
	return &HashSet[T]{hashmap.New[T, struct{}]()}
}

// Of creates a set with the given elements.
func Of[T comparable](elements ...T) HashSet[T] {
	set := HashSet[T]{hashmap.New[T, struct{}]()}
	for i := range elements {
		set.Add(elements[i])
	}
	return set
}

// Add adds the specified element to this set if it is not already present.
func (set *HashSet[T]) Add(e T) bool {
	value := set.hashmap.PutIfAbsent(e, struct{}{})
	return value.Empty()
}

// AddAll adds all of the elements in the specified iterable to the set.
func (set *HashSet[T]) AddAll(iterable iterable.Iterable[T]) bool {
	n := set.hashmap.Len()
	it := iterable.Iterator()
	for it.HasNext() {
		set.Add(it.Next())
	}
	return n != set.hashmap.Len()
}

// AddSlice adds all the elements in the slice to the set.
func (set *HashSet[T]) AddSlice(s []T) bool {
	n := set.hashmap.Len()
	for _, value := range s {
		set.Add(value)
	}
	return n != set.hashmap.Len()
}

// Remove removes the specified element from this set if it is present.
func (set *HashSet[T]) Remove(e T) bool {
	n := set.Len()
	_ = set.hashmap.Remove(e)
	return n != set.Len()
}

// RemoveIf removes all of the elements of this collection that satisfy the given predicate.
func (set *HashSet[T]) RemoveIf(f func(T) bool) bool {
	n := set.hashmap.Len()
	set.hashmap.RemoveIf(f)
	return n != set.hashmap.Len()
}

// RetainAll retains only the elements in the set that are contained in the specified collection.
func (set *HashSet[T]) RetainAll(c collections.Collection[T]) bool {
	return set.RemoveIf(func(e T) bool { return !c.Contains(e) })
}

// RemoveAll removes all of the set's elements that are also contained in the specified iterable.
func (set *HashSet[T]) RemoveAll(iterable iterable.Iterable[T]) bool {
	n := set.hashmap.Len()
	it := iterable.Iterator()
	for it.HasNext() {
		set.Remove(it.Next())
	}
	return n != set.hashmap.Len()
}

// RemoveSlice removes all of the set's elements that are also contained in the specified slice.
func (set *HashSet[T]) RemoveSlice(s []T) bool {
	n := set.hashmap.Len()
	for i := range s {
		set.Remove(s[i])
	}
	return n != set.hashmap.Len()
}

// Clear removes all of the elements from the set.
func (set *HashSet[T]) Clear() {
	set.hashmap.Clear()
}

// Contains returns true if the set contains the specified element.
func (set *HashSet[T]) Contains(e T) bool {
	return set.hashmap.ContainsKey(e)
}

// ContainsAll returns true if the set contains all of the elements of the specified iterable.
func (set *HashSet[T]) ContainsAll(iterable iterable.Iterable[T]) bool {
	it := iterable.Iterator()
	for it.HasNext() {
		if !set.Contains(it.Next()) {
			return false
		}
	}
	return true
}

// Len returns the number of elements in the set.
func (set *HashSet[T]) Len() int {
	return set.hashmap.Len()
}

// Empty returns true if the set contains no elements.
func (set *HashSet[T]) Empty() bool {
	return set.hashmap.Len() == 0
}

// Equals returns true if the set is equivalent to the given set. Two sets are equal if they are the same reference or have the same size and contain
// the same elements.
func (set *HashSet[T]) Equals(otherSet collections.Set[T]) bool {
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
func (set *HashSet[T]) ForEach(f func(T)) {
	for key := range set.hashmap {
		f(key)
	}
}

// Iterator returns an iterator over the elements in the set.
func (set *HashSet[T]) Iterator() iterator.Iterator[T] {
	return &setIterator[T]{iterator: set.hashmap.Iterator()}
}

// setIterator implememantation for [HashSet].
type setIterator[T comparable] struct {
	iterator iterator.Iterator[pair.Pair[T, struct{}]]
}

// HasNext returns true if the iterator has more elements.
func (it *setIterator[T]) HasNext() bool {
	return it.iterator.HasNext()
}

// Next returns the next element in the iterator.
func (it setIterator[T]) Next() T {
	return it.iterator.Next().Key()
}

// ToSlice returns a slice containing all the elements in the set.
func (set *HashSet[T]) ToSlice() []T {
	slice := make([]T, set.Len())
	i := 0
	for e, _ := range set.hashmap {
		slice[i] = e
		i++
	}
	return slice
}

// ImmutableCopy returns an immutable copy of the set.
func (set *HashSet[T]) ImmutableCopy() ImmutableHashSet[T] {
	return ImmutableOf(set.ToSlice()...)
}

// String returns the string representation of a set.
func (set HashSet[T]) String() string {
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
