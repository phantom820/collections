package interfaces

// Equitable how to check for element equality for our collections.
type Equitable[T any] interface {
	Equals(other T) bool
}

// Hashable for collections that rely on hashing capabilities.
type Hashable[T any] interface {
	HashCode() int
	Equitable[T]
}

// Comparable interface to allow comparison operators.
type Comparable[T any] interface {
	Equitable[T]
	Less(other T) bool
}

// Collection some underlying container which supports these operations.
type Collection[T Equitable[T]] interface {
	Iterable[T]
	Add(e T) bool
	AddAll(c Iterable[T])
	Len() int
	Contains(e T) bool
	Remove(e T) bool
	RemoveAll(c Iterable[T])
	Empty() bool
	Clear()
}
