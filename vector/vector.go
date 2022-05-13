// Package vector provides an implementation of a vector which is basicall a wrapper around a slice.
package vector

import (
	"errors"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types"
)

var (
	ErrOutOfBounds = errors.New("index out of bounds")
)

// Vector a wrapper around a slice.
type Vector[T types.Equitable[T]] []T

// New creates a vector with the specified elements, if there are none an empty vector is created.
func New[T types.Equitable[T]](elements ...T) Vector[T] {
	var v Vector[T] = make([]T, len(elements))
	v.AddSlice(elements)
	return v
}

// Add adds elements to the back of the vector.
func (v *Vector[T]) Add(elements ...T) bool {
	for _, e := range elements {
		*v = append(*v, e)
	}
	return true
}

// AddAll adds the elements from some iterable elements to the vector.
func (v *Vector[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		v.Add(it.Next())
	}
}

// AddSlice adds element from a slice s into the vector.
func (v *Vector[T]) AddSlice(s []T) {
	*v = append(*v, s...)
}

// Clear removes all elements in the vector.
func (v *Vector[T]) Clear() {
	*v = nil
	*v = make([]T, 0)
}

//AddAt adds an element at the specified index in the vector. Will panic if index is out of bounds.
func (v *Vector[T]) AddAt(i int, e T) {
	if i < 0 || i >= len(*v) {
		panic(ErrOutOfBounds)
	}
	for j, _ := range *v {
		if i == j {
			a := (*v)[0:i]
			a = append(a, e)
			b := (*v)[i:]
			(*v) = append(a, b...)
		}
	}
}

// Collect converts vector into a slice.
func (v *Vector[T]) Collect() []T {
	return *v
}

// Contains checks if the elemen e is in the vector.
func (v *Vector[T]) Contains(e T) bool {
	for i, _ := range *v {
		if (*v)[i].Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the vector is empty.
func (v *Vector[T]) Empty() bool {
	return len(*v) == 0
}

// vectorIterator model for implementing an iterator on a vector.
type vectorIterator[T types.Equitable[T]] struct {
	slice []T
	i     int
}

// HasNext check if the iterator has next element to produce.
func (it *vectorIterator[T]) HasNext() bool {
	if it.slice == nil || it.i >= len(it.slice) {
		return false
	}
	return true
}

// Next yields the next element from the iterator.
func (it *vectorIterator[T]) Next() T {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	e := it.slice[it.i]
	it.i++
	return e
}

// Cycle resets the iterator.
func (it *vectorIterator[T]) Cycle() {
	it.i = 0
}

// Equals checks if the vector is equals to other. This only true if they are the same reference or have the same size and element wise comparison passes
// otherwise false.
func (v *Vector[T]) Equals(other *Vector[T]) bool {
	if v == other {
		return true
	} else if len(*v) != len(*other) {
		return false
	} else {
		for i, _ := range *v {
			if !(*v)[i].Equals((*other)[i]) {
				return false
			}
		}
		return true
	}
}

// Iterator returns an iterator for iterating through a vector.
func (v *Vector[T]) Iterator() iterator.Iterator[T] {
	return &vectorIterator[T]{slice: *v, i: 0}
}

// Len return the size of the vector.
func (v Vector[T]) Len() int {
	return len(v)
}

// indexOf finds the index of an element e in the vector. Gives -1 if the element is not present.
func (v *Vector[T]) indexOf(e T) int {
	for i, _ := range *v {
		if (*v)[i].Equals(e) {
			return i
		}
	}
	return -1
}

// Remove deletes the element e from the vector.
func (v *Vector[T]) Remove(e T) bool {
	i := v.indexOf(e)
	if i != -1 {
		*v = append((*v)[0:i], (*v)[i+1:]...)
		return true
	}
	return false
}

// RemoveAt deletes the element at the specified index in the vector. Will panic if index is out of bounds.
func (v *Vector[T]) RemoveAt(i int) T {
	if i < 0 || i >= len(*v) {
		panic(ErrOutOfBounds)
	}
	temp := (*v)[i]
	*v = append((*v)[0:i], (*v)[i+1:]...)
	return temp
}

// RemoveAll removes all the elements from some iterable elements that are in the vector.
func (v *Vector[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		v.Remove(it.Next())
	}
}
