package collections

import (
	"reflect"
)

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
	RemoveSlice(s []T) bool
	RetainAll(c Collection[T]) bool
	ForEach(func(T))
	Len() int
	ToSlice() []T
}

type List[T comparable] interface {
	Collection[T]
	AddAt(i int, e T)
	At(i int) T
	Set(i int, e T) T
	RemoveAt(i int) T
	Equals(list List[T]) bool
	Sort(less func(a, b T) bool)
}

func IsNil[T comparable](c Collection[T]) bool {
	return c == nil || reflect.ValueOf(c).IsNil()
}
