// Package vector provides an implementation of a vector which is b a wrapper around a slice.
package vector

import (
	"errors"
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types"
)

var (
	ErrOutOfBounds = errors.New("index out of bounds")
)

// Vector a wrapper around a slice.
type Vector[T types.Equitable[T]] struct {
	data []T
}

// New creates a vector with the specified elements, if there are none an empty vector is created.
func New[T types.Equitable[T]](elements ...T) *Vector[T] {
	v := Vector[T]{data: make([]T, 0)}
	v.AddSlice(elements)
	return &v
}

// Set replaces the element at index i in the vector with the new element. Returns the old element that was at index i.
func (v *Vector[T]) Set(i int, e T) T {
	if i < 0 || i >= v.Len() {
		panic(ErrOutOfBounds)
	}
	temp := v.data[i]
	v.data[i] = e
	return temp
}

// Add adds elements to the back of the vector.
func (v *Vector[T]) Add(elements ...T) bool {
	v.data = append(v.data, elements...)
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
	v.data = append(v.data, s...)
}

// Clear removes all elements in the vector.
func (v *Vector[T]) Clear() {
	v.data = nil
	v.data = make([]T, 0)
}

// At retrieves the element at the specified index.  Will panic if index is out of bounds.
func (v *Vector[T]) At(i int) T {
	if i < 0 || i >= v.Len() {
		panic(ErrOutOfBounds)
	}
	return v.data[i]
}

//AddAt adds an element at the specified index in the vector. Will panic if index is out of bounds.
func (v *Vector[T]) AddAt(i int, e T) {
	if i < 0 || i >= len(v.data) {
		panic(ErrOutOfBounds)
	}
	for j, _ := range v.data {
		if i == j {
			a := v.data[0:i]
			a = append(a, e)
			b := v.data[i:]
			v.data = append(a, b...)
		}
	}
}

// Collect converts vector into a slice.
func (v *Vector[T]) Collect() []T {
	return v.data
}

// Contains checks if the elemen e is in the vector.
func (v *Vector[T]) Contains(e T) bool {
	for i, _ := range v.data {
		if v.data[i].Equals(e) {
			return true
		}
	}
	return false
}

// Empty checks if the vector is empty.
func (v *Vector[T]) Empty() bool {
	return len(v.data) == 0
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
	} else if len(v.data) != len(other.data) {
		return false
	} else {
		for i, _ := range v.data {
			if !(v.data)[i].Equals((other.data)[i]) {
				return false
			}
		}
		return true
	}
}

// Iterator returns an iterator for iterating through a vector.
func (v *Vector[T]) Iterator() iterator.Iterator[T] {
	return &vectorIterator[T]{slice: v.data, i: 0}
}

// Len return the size of the vector.
func (v Vector[T]) Len() int {
	return len(v.data)
}

// indexOf finds the index of an element e in the vector. Gives -1 if the element is not present.
func (v *Vector[T]) indexOf(e T) int {
	for i, _ := range v.data {
		if (v.data)[i].Equals(e) {
			return i
		}
	}
	return -1
}

// Remove deletes the element e from the vector.
func (v *Vector[T]) Remove(e T) bool {
	i := v.indexOf(e)
	if i == -1 {
		return false
	}
	v.data = append((v.data)[0:i], (v.data)[i+1:]...)
	return true
}

// RemoveAt deletes the element at the specified index in the vector. Will panic if index is out of bounds.
func (v *Vector[T]) RemoveAt(i int) T {
	if i < 0 || i >= len(v.data) {
		panic(ErrOutOfBounds)
	}
	temp := (v.data)[i]
	v.data = append(v.data[0:i], v.data[i+1:]...)
	return temp
}

// RemoveAll removes all the elements from some iterable elements that are in the vector.
func (v *Vector[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		v.Remove(it.Next())
	}
}

// Map transforms each element in the vector using a specified function and returns a new vector with transformed elements.
func (v *Vector[T]) Map(f func(e T) T) *Vector[T] {
	newV := New[T]()
	for _, e := range v.data {
		newV.Add(f(e))
	}
	return newV
}

// Filter filters elements of the vector using a specified predicate function and returns a new vector with elements satisfying predicate.
func (v *Vector[T]) Filter(f func(e T) bool) *Vector[T] {
	newV := New[T]()
	for _, e := range v.data {
		if f(e) {
			newV.Add(e)
		}
	}
	return newV
}

// String for pretty printing the vector.
func (s *Vector[T]) String() string {
	return fmt.Sprint(s.data)
}
