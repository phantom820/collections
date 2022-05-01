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
	Rank() int8
	LoadFactor() float32
	// Map(func(e T) T) Set[T]
	// Filter(func(e T) bool) Set[T]
}

// newSet makes a new set from another set.
// func newSet[T interfaces.Hashable,E any](a Set[T,E]) Set[T,E] {
// 	rank := a.Rank()
// 	switch rank {
// 	case HashSetRank:
// 		s := NewHashSet[T]()
// 		s.AddAll(a)
// 		return s
// 	default:
// 		s := NewHashSet[T]()
// 		s.AddAll(a)
// 		return s
// 	}
// }
