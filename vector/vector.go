// Package vector provides an implementation of a vector which is a wrapper around a slice.
package vector

import (
	"errors"
	"fmt"
	"sort"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types"
)

var (
	ErrOutOfBounds = errors.New("index out of bounds")
)

// Vector an implementation of a vector by wrapping around a slice.
type Vector[T types.Equitable[T]] struct {
	data []T
}

// New creates a vector with the specified elements.
func New[T types.Equitable[T]](elements ...T) *Vector[T] {
	vector := Vector[T]{data: make([]T, 0)}
	vector.Add(elements...)
	return &vector
}

// Set replaces the element at index i in the vector with the new element. Returns the old element that was at index i.
func (vector *Vector[T]) Set(i int, e T) T {
	if i < 0 || i >= vector.Len() {
		panic(ErrOutOfBounds)
	}
	temp := vector.data[i]
	vector.data[i] = e
	return temp
}

// Add adds elements to the vector.
func (vector *Vector[T]) Add(elements ...T) bool {
	n := vector.Len()
	vector.data = append(vector.data, elements...)
	return (n != vector.Len())
}

// AddAll adds the elements from some iterable elements to the vector.
func (vector *Vector[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		vector.Add(it.Next())
	}
}

// Clear removes all elements in the vector.
func (vector *Vector[T]) Clear() {
	vector.data = nil
	vector.data = make([]T, 0)
}

// At retrieves the element at the specified index.  Will panic if index is out of bounds.
func (vector *Vector[T]) At(i int) T {
	if i < 0 || i >= vector.Len() {
		panic(ErrOutOfBounds)
	}
	return vector.data[i]
}

//AddAt adds an element at the specified index in the vector. Will panic if index is out of bounds.
func (vector *Vector[T]) AddAt(i int, e T) {
	if i < 0 || i >= len(vector.data) {
		panic(ErrOutOfBounds)
	}
	if i == 0 {
		vector.data = append([]T{e}, vector.data...)
		return
	} else if i == vector.Len()-1 {
		temp := vector.data[vector.Len()-1]
		vector.data = append(vector.data[0:vector.Len()-1], e, temp)
		return
	}
	for j, _ := range vector.data {
		if i == j {
			a := vector.data[0:i]
			a = append(a, e)
			b := vector.data[i:]
			vector.data = append(a, b...)
			return
		}
	}
}

// Collect returns a slice containing all the elements in the vector.
func (vector *Vector[T]) Collect() []T {
	return vector.data
}

// Contains checks if the element is in the vector.
func (vector *Vector[T]) Contains(element T) bool {
	for i, _ := range vector.data {
		if vector.data[i].Equals(element) {
			return true
		}
	}
	return false
}

// Empty checks if the vector is empty.
func (vector *Vector[T]) Empty() bool {
	return len(vector.data) == 0
}

// vectorIterator type for implementing an iterator on a vector.
type vectorIterator[T types.Equitable[T]] struct {
	slice []T
	i     int
}

// HasNext checks if the iterator has a next element to yield.
func (it *vectorIterator[T]) HasNext() bool {
	if it.slice == nil || it.i >= len(it.slice) {
		return false
	}
	return true
}

// Next returns the next element in the iterator it. Will panic if iterator has no next element.
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
func (vector *Vector[T]) Equals(other *Vector[T]) bool {
	if vector == other {
		return true
	} else if len(vector.data) != len(other.data) {
		return false
	} else {
		for i, _ := range vector.data {
			if !(vector.data)[i].Equals((other.data)[i]) {
				return false
			}
		}
		return true
	}
}

// Iterator returns an iterator for iterating through a vector.
func (vector *Vector[T]) Iterator() iterator.Iterator[T] {
	return &vectorIterator[T]{slice: vector.data, i: 0}
}

// Len return the size of the vector.
func (vector *Vector[T]) Len() int {
	return len(vector.data)
}

// indexOf finds the index of an element e in the vector. Gives -1 if the element is not present.
func (vector *Vector[T]) indexOf(e T) int {
	for i, _ := range vector.data {
		if (vector.data)[i].Equals(e) {
			return i
		}
	}
	return -1
}

// Remove removes elements from the list. Only the first occurence of each element is removed.
func (vector *Vector[T]) Remove(elements ...T) bool {
	n := vector.Len()
	for _, element := range elements {
		vector.remove(element)
		if vector.Empty() {
			break
		}
	}
	return (n != vector.Len())
}

// remove deletes the element from the vector. For internal use to support Remove.
func (vector *Vector[T]) remove(element T) bool {
	i := vector.indexOf(element)
	if i == -1 {
		return false
	}
	vector.data = append((vector.data)[0:i], (vector.data)[i+1:]...)
	return true
}

// RemoveAt deletes the element at the specified index in the vector. Will panic if index is out of bounds.
func (vector *Vector[T]) RemoveAt(i int) T {
	if i < 0 || i >= len(vector.data) {
		panic(ErrOutOfBounds)
	}
	temp := (vector.data)[i]
	vector.data = append(vector.data[0:i], vector.data[i+1:]...)
	return temp
}

// RemoveAll removes all the elements from iterable elements that are in the vector.
func (vector *Vector[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		vector.Remove(it.Next())
	}
}

// Map transforms each element in the vector using a specified function and returns a new vector with transformed elements.
func (vector *Vector[T]) Map(f func(e T) T) *Vector[T] {
	newVector := New[T]()
	for _, e := range vector.data {
		newVector.Add(f(e))
	}
	return newVector
}

// Filter filters elements of the vector using the predicate function f and returns a new vector with elements satisfying predicate.
func (vector *Vector[T]) Filter(f func(e T) bool) *Vector[T] {
	newVector := New[T]()
	for _, e := range vector.data {
		if f(e) {
			newVector.Add(e)
		}
	}
	return newVector
}

// String for pretty printing the vector.
func (s *Vector[T]) String() string {
	return fmt.Sprint(s.data)
}

// Sort the vector using the natural ordering of its elements.
func Sort[T types.Comparable[T]](vector *Vector[T]) {
	sort.Slice(vector.data, func(i, j int) bool {
		return vector.data[i].Less(vector.data[j])
	})
}

// SortBy sorts the vector using the function less for comparison of two element . if less(a,b) = true then it means a comes before b in the sorted list.
func SortBy[T types.Equitable[T]](vector *Vector[T], less func(a, b T) bool) {
	sort.Slice(vector.data, func(i, j int) bool {
		return less(vector.data[i], vector.data[j])
	})
}
