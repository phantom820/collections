package set

import "collections/interfaces"

const (
	HashSetRank     = 3
	LoadFactorLimit = 0.75
	Capacity        = 16
)

// Set methods to be supported by an implementation of set i.e HashSet,TreeSet ...
type Set[T interfaces.Equitable[T], E any] interface {
	interfaces.Iterable[T]
	interfaces.Collection[T]
	interfaces.Functional[T, E]
}
