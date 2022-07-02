// Package iterator provides an interface for specifying an iterable. Any type that satisfies this interface is an Iterable.
package iterator

import "errors"

var (
	NoNextElementError = errors.New("iterator has no next element.")
)

// Iterator interface for anything that is iterable must satisfy.
type Iterator[T any] interface {
	HasNext() bool // Checks if the iterator has a next element to produce.
	Next() T       // Returns the next element.
	Cycle()        // Resets the iterator back to its initial position.
}

type PartitionedIterator[T any] struct {
	iterators []Iterator[T]
}

// Iterable interface for anything that is iterable and can be collected into a slice..
type Iterable[T any] interface {
	Collect() []T
	Iterator() Iterator[T]
}
