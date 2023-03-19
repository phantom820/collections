// package iterator defines a way to access the elements of a collection one by one. The two basic operations on an iterator
// are Next and HasNext. A call to Next() will return the next element of the iterator and advance the state of the iterator.
// A call to Next() should always be preceded by a call to HasNext(), otherwise a NoSuchElement panic may occur if the iterator has no next elements.
package iterator

import (
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/types/optional"
)

// Iterator a sequential accessor to elements of an underlying container.
type Iterator[T any] interface {
	HasNext() bool // Returns true if the iterator has more elements.
	Next() T       // Returns the next element in the iterator.
}

// Of returns an iterator over the given elements.
func Of[T any](elements ...T) Iterator[T] {
	return &iterator[T]{elements: elements, index: 0}
}

// iterator slice based iterator implementation.
type iterator[T any] struct {
	elements []T
	index    int
}

// HasNext returns true if the iterator has more elements.
func (it *iterator[T]) HasNext() bool {
	if it.elements == nil {
		return false
	}
	return it.index < len(it.elements)
}

// Next returns the next element in the iterator.
func (it *iterator[T]) Next() T {
	if !it.HasNext() {
		panic(errors.NoSuchElement())
	}
	element := it.elements[it.index]
	it.index++
	return element
}

// Map return the iterator obtained from applying the transformation function to every element on the given iterator.
func Map[T, U comparable](it Iterator[T], f func(T) U) Iterator[U] {
	elements := make([]U, 0, 8)
	for it.HasNext() {
		elements = append(elements, f(it.Next()))
	}
	return Of(elements...)
}

// Filter returns an iterator with all elements that satisfy the given predicate.
func Filter[T comparable](it Iterator[T], f func(T) bool) Iterator[T] {
	elements := make([]T, 0, 8)
	for it.HasNext() {
		element := it.Next()
		if f(element) {
			elements = append(elements, element)
		}
	}
	return Of(elements...)
}

// Reduce reduces the elements of the iterator using the associative binary function and returns result as an option.
func Reduce[T comparable](it Iterator[T], f func(x, y T) T) optional.Optional[T] {
	if !it.HasNext() {
		return optional.Empty[T]()
	}
	x := it.Next()
	for it.HasNext() {
		y := it.Next()
		x = f(x, y)
	}
	return optional.Of(x)
}

func ToSlice[T any](it Iterator[T]) []T {
	data := make([]T, 0)
	for it.HasNext() {
		data = append(data, it.Next())
	}
	return data
}

// func Partition[T any](it Iterator[T], n int) [][]T {

// 	partitions := make([][]T, 0, n)

// 	for it.HasNext() {
// 		complete := false
// 		for i := 0; i < n; i++ {
// 			if it.HasNext() {
// 				element := it.Next()
// 				if i >= len(partitions) {
// 					partitions = append(partitions, make([]T, 0, 8))
// 				}
// 				partitions[i] = append(partitions[i], element)
// 			} else {
// 				complete = true
// 				break
// 			}
// 		}
// 		if complete {
// 			break
// 		}
// 	}

// 	return partitions
// }
