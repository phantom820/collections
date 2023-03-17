package collections

import (
	"reflect"

	"github.com/phantom820/collections/iterable"
)

type Collection[T comparable] interface {
	iterable.Iterable[T]
	Add(e T) bool
	AddAll(iterable iterable.Iterable[T]) bool
	AddSlice(s []T) bool
	Clear()
	Contains(e T) bool
	Empty() bool
	Remove(e T) bool
	RemoveIf(func(T) bool) bool
	RemoveAll(iterable iterable.Iterable[T]) bool
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

type Set[T comparable] interface {
	Collection[T]
	ContainsAll(iterable iterable.Iterable[T]) bool
}

func IsNil[T comparable](c Collection[T]) bool {
	return c == nil || reflect.ValueOf(c).IsNil()
}
