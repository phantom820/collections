// package iterable defines a base interface for all collections that define an iterator method to step a collection's elements.
// Implementations of this interface need to provide a concrete Iterator() method that yields an iterator to a given collection.
package iterable

import (
	"github.com/phantom820/collections/iterator"
)

// Iterable a base interface for iterable collections.
type Iterable[T any] interface {
	Iterator() iterator.Iterator[T] // Returns a new iterator over all elements contained in the iterable.
}

// Of returns an iterable of the given elements.
func Of[T any](elements ...T) Iterable[T] {
	return &iterable[T]{iterator: iterator.Of(elements...)}
}

// iterable represents an iterable constructed from slice.
type iterable[T any] struct {
	iterator iterator.Iterator[T]
}

// Iterator returns iterator for the iterable.
func (iterable *iterable[T]) Iterator() iterator.Iterator[T] {
	return iterable.iterator
}
