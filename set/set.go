// Package set provides an interface that an implementation of a set must satisfy.
package set

import (
	"collections/interfaces"
	"collections/iterator"
	"collections/types"
)

// Set interface specifying a list of methods a set implementation is expected to provide.
type Set[T types.Equitable[T], E any] interface {
	iterator.Iterable[T]
	interfaces.Collection[T]
	interfaces.Functional[T, E]
}
