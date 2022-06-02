// Package sets provides some utils functons for sets and an interface set implementation must follow.
package sets

import (
	"errors"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/types"
)

var (
	errInvalidDestinationSet = errors.New("invalid destination set for operation, detsination must be empty set")
)

// Set an interface that a set implementation should satisfy.
type Set[T types.Equitable[T]] interface {
	collections.Collection[T]
	RetainAll(collection collections.Collection[T]) bool // Removes all entries from the set that do not appear in the other collection. Return true if the set was modified.
	RemoveIf(f func(element T) bool) bool                // Removes elements from the set that satisfy the predicate function f
}

// Union union of sets a and b, results of the operation are placed in the set c.
func Union[T types.Equitable[T]](a Set[T], b Set[T], c Set[T]) error {
	if !c.Empty() {
		return errInvalidDestinationSet
	}
	c.AddAll(a)
	c.AddAll(b)
	return nil
}

// intersection helper function to perform set intersection the idea is iterate over bigger set and lookup in smaller.
func intersection[T types.Equitable[T]](a Set[T], b Set[T], c Set[T]) error {
	if a.Len() > b.Len() {
		it := a.Iterator()
		for it.HasNext() {
			e := it.Next()
			if b.Contains(e) {
				c.Add(e)
			}
		}
		return nil
	}
	it := b.Iterator()
	for it.HasNext() {
		e := it.Next()
		if a.Contains(e) {
			c.Add(e)
		}
	}
	return nil
}

// Intersection performs the union of sets a and b, results of the operation are placed in the set c.
func Intersection[T types.Equitable[T]](a Set[T], b Set[T], c Set[T]) error {
	if !c.Empty() {
		return errInvalidDestinationSet
	}
	intersection(a, b, c)
	return nil
}

// Difference performs the difference of sets a and b, results of the operation are placed in the set c.
func Difference[T types.Equitable[T]](a Set[T], b Set[T], c Set[T]) error {
	if !c.Empty() {
		return errInvalidDestinationSet
	}
	c.AddAll(a)
	c.RemoveAll(b)
	return nil
}
