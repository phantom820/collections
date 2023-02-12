package collections

type Iterator[T any] interface {
	Next() T
	HasNext() bool
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type Collection[T comparable] interface {
	Iterable[T]
	Add(e T) bool
	AddAll(iterable Iterable[T]) bool
	AddSlice(s []T) bool
	Clear()
	Contains(e T) bool
	Empty() bool
	Remove(e T) bool
	RemoveIf(func(T) bool) bool
	RemoveAll(iterable Iterable[T]) bool
	RetainAll(c Collection[T]) bool
	ForEach(func(T))
	Len() int
}

type ImmutableCollection[T comparable] interface {
	Iterable[T]
	Contains(e T) bool
	Empty() bool
	ForEach(func(T))
	Len() int
}
