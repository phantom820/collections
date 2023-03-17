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
