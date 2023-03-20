// package treeset defines a set implementation that is backed by a [TreeMap].
package treeset

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps/treemap"
	"github.com/phantom820/collections/types/pair"
)

// TreeSet implementation of a set backed by a [TreeMap].
type TreeSet[T comparable] struct {
	treeMap  *treemap.TreeMap[T, struct{}]
	lessThan func(e1, e2 T) bool
}

// New creates a mutable set with the given elements. Elements are compared using the lessThan function which should satisfy.
// e1 < e2 => lessThan(e1, e2) = true and lessThan(e2,e1) = false.
// e1 = e2 => lessThan(e1,e2) = false and lessThan(e2,e1) = false.
// e1 > e2 -> lessThan(e1,e2) = false and lessThan(e2,e1) = true.
func New[T comparable](lessThan func(e1, e2 T) bool, elements ...T) *TreeSet[T] {
	set := TreeSet[T]{lessThan: lessThan, treeMap: treemap.New[T, struct{}](lessThan)}
	for _, e := range elements {
		set.Add(e)
	}
	return &set
}

// Of creates an immutable set with the given elements. Elements are compared using the lessThan function which should satisfy.
// e1 < e2 => lessThan(e1, e2) = true and lessThan(e2,e1) = false.
// e1 = e2 => lessThan(e1,e2) = false and lessThan(e2,e1) = false.
// e1 > e2 -> lessThan(e1,e2) = false and lessThan(e2,e1) = true.
func Of[T comparable](lessThan func(e1, e2 T) bool, elements ...T) ImmutableTreeSet[T] {
	set := TreeSet[T]{lessThan: lessThan, treeMap: treemap.New[T, struct{}](lessThan)}
	for i := range elements {
		set.Add(elements[i])
	}
	return ImmutableTreeSet[T]{treeSet: set}
}

// Add adds the specified element to this set if it is not already present.
func (set *TreeSet[T]) Add(e T) bool {
	value := set.treeMap.Put(e, struct{}{})
	return value.Empty()
}

// AddAll adds all of the elements in the specified iterable to the set.
func (set *TreeSet[T]) AddAll(iterable iterable.Iterable[T]) bool {
	n := set.Len()
	it := iterable.Iterator()
	for it.HasNext() {
		set.Add(it.Next())
	}
	return n != set.Len()
}

// AddSlice adds all the elements in the slice to the set.
func (set *TreeSet[T]) AddSlice(s []T) bool {
	n := set.Len()
	for _, value := range s {
		set.Add(value)
	}
	return n != set.Len()
}

// Remove removes the specified element from this set if it is present.
func (set *TreeSet[T]) Remove(e T) bool {
	n := set.Len()
	set.treeMap.Remove(e)
	return n != set.Len()
}

// RemoveIf removes all of the elements of this collection that satisfy the given predicate.
func (set TreeSet[T]) RemoveIf(f func(T) bool) bool {
	n := set.Len()
	set.treeMap.RemoveIf(f)
	return n != set.Len()
}

// RetainAll retains only the elements in the set that are contained in the specified collection.
func (set *TreeSet[T]) RetainAll(c collections.Collection[T]) bool {
	switch c.(type) {
	case collections.Set[T]:
		return set.RemoveIf(func(e T) bool { return !c.Contains(e) })
	default:
		{
			otherSet := make(map[T]struct{})
			it := c.Iterator()
			for it.HasNext() {
				e := it.Next()
				if _, ok := otherSet[e]; !ok {
					otherSet[e] = struct{}{}
				}
			}
			return set.RemoveIf(func(e T) bool {
				_, ok := otherSet[e]
				return !ok
			})
		}
	}

}

// RemoveAll removes all of the set's elements that are also contained in the specified iterable.
func (set TreeSet[T]) RemoveAll(iterable iterable.Iterable[T]) bool {
	n := set.Len()
	it := iterable.Iterator()
	for it.HasNext() {
		set.Remove(it.Next())
	}
	return n != set.Len()
}

// RemoveSlice removes all of the set's elements that are also contained in the specified slice.
func (set TreeSet[T]) RemoveSlice(s []T) bool {
	n := set.Len()
	for i := range s {
		set.Remove(s[i])
	}
	return n != set.Len()
}

// ToSlice returns a slice containing all the elements in the set.
func (set *TreeSet[T]) ToSlice() []T {
	return set.treeMap.Keys()
}

// Clear removes all of the elements from the set.
func (set TreeSet[T]) Clear() {
	set.treeMap.Clear()
}

// ImmutableCopy returns an immutable copy of the set.
func (set *TreeSet[T]) ImmutableCopy() ImmutableTreeSet[T] {
	return Of(set.lessThan, set.ToSlice()...)
}

// Contains returns true if this set contains the specified element.
func (set TreeSet[T]) Contains(e T) bool {
	return set.treeMap.ContainsKey(e)
}

// ContainsAll returns true if the set contains all of the elements of the specified iterable.
func (set *TreeSet[T]) ContainsAll(iterable iterable.Iterable[T]) bool {
	it := iterable.Iterator()
	for it.HasNext() {
		if !set.Contains(it.Next()) {
			return false
		}
	}
	return true
}

// Len returns the number of elements in the set.
func (set TreeSet[T]) Len() int {
	return set.treeMap.Len()
}

// Empty returns true if the set contains no elements.
func (set TreeSet[T]) Empty() bool {
	return set.treeMap.Len() == 0
}

// Equals returns true if the set is equivalent to the given set. Two sets are equal if they are the same reference or have the same size and contain
// the same elements.
func (set *TreeSet[T]) Equals(otherSet collections.Set[T]) bool {
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
func (set TreeSet[T]) ForEach(f func(T)) {
	it := set.treeMap.Iterator()
	for it.HasNext() {
		f(it.Next().Key())
	}
}

// Iterator returns an iterator over the elements in the set.
func (set *TreeSet[T]) Iterator() iterator.Iterator[T] {
	return &setIterator[T]{mapIterator: set.treeMap.Iterator()}
}

// setIterator implememantation for [HashSet].
type setIterator[T comparable] struct {
	mapIterator iterator.Iterator[pair.Pair[T, struct{}]]
}

// HasNext returns true if the iterator has more elements.
func (it *setIterator[T]) HasNext() bool {
	return it.mapIterator.HasNext()
}

// Next returns the next element in the iterator.
func (it setIterator[T]) Next() T {
	return it.mapIterator.Next().Key()
}

// String returns the string representation of a set.
func (set TreeSet[T]) String() string {
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
