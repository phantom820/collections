// Package iterator provides an interface which specifies methods to be specified for implementing in iterator
// over a collection.
package iterator

import "errors"

var (
	NoNextElementError = errors.New("iterator has no next element.")
)

// Iterator specifies methods a collection must implement to allow iterating through it.
type Iterator[T any] interface {
	HasNext() bool // Checks if the iterator has not been exhausted.
	Next() T       // Retrieves the next element from the iterator.
	Cycle()
}

// Iterable effectively anything that can be converted to a slice and can be iterated on.
type Iterable[T any] interface {
	Collect() []T
	Iterator() Iterator[T]
}
