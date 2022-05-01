package interfaces

// Iterator interface for how we can get elements of a data structure.
type Iterator[T any] interface {
	HasNext() bool
	Next() T
	Cycle()
}

// Iterable anything that can be collected to a slice and iterated on.
type Iterable[T any] interface {
	Collect() []T
	Iterator() Iterator[T]
}
