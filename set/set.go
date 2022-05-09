// Package set provides an interface that an implementation of a set must satisfy.
package set

import (
	"collections/interfaces"
	"collections/iterator"
	"collections/types"
)

// Set methods to be supported by an implementation of set i.e HashSet,TreeSet ...
type Set[T types.Equitable[T], E any] interface {
	iterator.Iterable[T]
	interfaces.Collection[T]
	interfaces.Functional[T, E]
}
