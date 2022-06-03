// Package vector provides an implementation of a list which is a wrapper around a slice.
package vector

import (
	"errors"
	"fmt"
	"sort"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists"
	"github.com/phantom820/collections/types"
)

var (
	ErrOutOfBounds = errors.New("index out of bounds")
)

// Vector an implementation of a list by wrapping around a slice.
type Vector[T types.Equitable[T]] struct {
	data []T
}

// New creates a list with the specified elements. If no elements are specified an empty list is created.
func New[T types.Equitable[T]](elements ...T) *Vector[T] {
	list := Vector[T]{data: make([]T, 0)}
	list.Add(elements...)
	return &list
}

// AddFront adds elements to the front of the list.
func (list *Vector[T]) AddFront(elements ...T) {
	list.data = append(elements, list.data...)
}

// AddBack adds elements to the back of the list.
func (list *Vector[T]) AddBack(elements ...T) {
	list.data = append(list.data, elements...)
}

// Set replaces the element at index i in the list with the new element. Returns the old element that was at index i.
func (list *Vector[T]) Set(i int, e T) T {
	if i < 0 || i >= list.Len() {
		panic(ErrOutOfBounds)
	}
	temp := list.data[i]
	list.data[i] = e
	return temp
}

// Add adds elements to the list.
func (list *Vector[T]) Add(elements ...T) bool {
	n := list.Len()
	list.data = append(list.data, elements...)
	return (n != list.Len())
}

// AddAll adds the elements from iterable to the list.
func (list *Vector[T]) AddAll(iterable iterator.Iterable[T]) {
	it := iterable.Iterator()
	for it.HasNext() {
		list.Add(it.Next())
	}
}

// Clear removes all elements in the list.
func (list *Vector[T]) Clear() {
	list.data = nil
	list.data = make([]T, 0)
}

// At retrieves the element at the specified index.  Will panic if index is out of bounds.
func (list *Vector[T]) At(i int) T {
	if i < 0 || i >= list.Len() {
		panic(ErrOutOfBounds)
	}
	return list.data[i]
}

// AddAt adds an element to the list at specified index, all subsequent elements will be shifted right. Will panic if index is out of bounds.
func (list *Vector[T]) AddAt(i int, e T) {
	if i < 0 || i >= len(list.data) {
		panic(ErrOutOfBounds)
	}
	if i == 0 {
		list.data = append([]T{e}, list.data...)
		return
	}
	for j, _ := range list.data {
		if j == i-1 {
			a := list.data[0:i]
			a = append(a, e)
			b := list.data[i:]
			list.data = append(a, b...)
			return
		}
	}
}

// Collect returns a slice containing all the elements in the list.
func (list *Vector[T]) Collect() []T {
	return list.data
}

// Contains checks if the element is in the list.
func (list *Vector[T]) Contains(element T) bool {
	for i, _ := range list.data {
		if list.data[i].Equals(element) {
			return true
		}
	}
	return false
}

// Empty checks if the list is empty.
func (list *Vector[T]) Empty() bool {
	return len(list.data) == 0
}

// listIterator type for implementing an iterator on a list.
type listIterator[T types.Equitable[T]] struct {
	slice []T
	i     int
}

// HasNext checks if the iterator has a next element to yield.
func (it *listIterator[T]) HasNext() bool {
	if it.slice == nil || it.i >= len(it.slice) {
		return false
	}
	return true
}

// Next returns the next element in the iterator it. Will panic if iterator has no next element.
func (it *listIterator[T]) Next() T {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	e := it.slice[it.i]
	it.i++
	return e
}

// Cycle resets the iterator.
func (it *listIterator[T]) Cycle() {
	it.i = 0
}

// Equals checks if the list is equals to other. This only true if they are the same reference or have the same size and their elements match.
func (list *Vector[T]) Equals(other *Vector[T]) bool {
	if list == other {
		return true
	} else if len(list.data) != len(other.data) {
		return false
	} else {
		for i, _ := range list.data {
			if !(list.data)[i].Equals((other.data)[i]) {
				return false
			}
		}
		return true
	}
}

// Iterator returns an iterator for the list.
func (list *Vector[T]) Iterator() iterator.Iterator[T] {
	return &listIterator[T]{slice: list.data, i: 0}
}

// Len return the size of the list.
func (list *Vector[T]) Len() int {
	return len(list.data)
}

// indexOf finds the index of an element e in the list. Gives -1 if the element is not present.
func (list *Vector[T]) indexOf(e T) int {
	for i, _ := range list.data {
		if (list.data)[i].Equals(e) {
			return i
		}
	}
	return -1
}

// Remove removes elements from the list. Only the first occurence of an element is removed.
func (list *Vector[T]) Remove(elements ...T) bool {
	n := list.Len()
	for _, element := range elements {
		list.remove(element)
		if list.Empty() {
			break
		}
	}
	return (n != list.Len())
}

// RemoveFront removes and returns the front element of the list. Will panic if list has no front element.
func (list *Vector[T]) RemoveFront() T {
	if list.Empty() {
		panic(lists.ErrEmptyList)
	}
	front := list.data[0]
	list.data = list.data[1:]
	return front
}

// RemoveBack removes and returns the back element of the list. Will panic if the list has no back element.
func (list *Vector[T]) RemoveBack() T {
	if list.Empty() {
		panic(lists.ErrEmptyList)
	}
	back := list.data[list.Len()-1]
	list.data = list.data[:list.Len()-1]
	return back
}

// remove deletes the element from the list. For internal use to support Remove.
func (list *Vector[T]) remove(element T) bool {
	i := list.indexOf(element)
	if i == -1 {
		return false
	}
	list.data = append((list.data)[0:i], (list.data)[i+1:]...)
	return true
}

// RemoveAt deletes the element at the specified index in the list. Will panic if index is out of bounds.
func (list *Vector[T]) RemoveAt(i int) T {
	if i < 0 || i >= len(list.data) {
		panic(ErrOutOfBounds)
	}
	temp := (list.data)[i]
	list.data = append(list.data[0:i], list.data[i+1:]...)
	return temp
}

// RemoveAll removes all the elements in the list that appear in the iterable.
func (list *Vector[T]) RemoveAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		list.Remove(it.Next())
	}
}

// Front returns the front of the list. Will panic if list has no front element.
func (list *Vector[T]) Front() T {
	if list.Empty() {
		panic(lists.ErrEmptyList)
	}
	return list.data[0]
}

// Back returns the back element of the list.  Will panic if list has no back element.
func (list *Vector[T]) Back() T {
	if list.Empty() {
		panic(lists.ErrEmptyList)
	}
	return list.data[list.Len()-1]
}

// Map transforms each element in the list the function f and returns a new list with transformed elements.
func (list *Vector[T]) Map(f func(element T) T) *Vector[T] {
	newVector := New[T]()
	for _, e := range list.data {
		newVector.Add(f(e))
	}
	return newVector
}

// Filter filters elements of the list using the predicate function f and returns a new list with elements satisfying predicate.
func (list *Vector[T]) Filter(f func(element T) bool) *Vector[T] {
	newVector := New[T]()
	for _, e := range list.data {
		if f(e) {
			newVector.Add(e)
		}
	}
	return newVector
}

// Swap swaps the element at index i and the element at index j. This is done using links. Will panic if one/both of the specified indices is
//  out of bounds.
func (list *Vector[T]) Swap(i, j int) {
	if i < 0 || i >= list.Len() || j < 0 || j >= list.Len() {
		panic(lists.ErrOutOfBounds)
	}
	temp := list.data[i]
	list.data[i] = list.data[j]
	list.data[j] = temp
}

// Reverse reverses the list in place.
func (list *Vector[T]) Reverse() {
	n := list.Len()
	for i := 0; i < n/2; i++ {
		list.Swap(i, n-i-1)
	}
}

// String for pretty printing the list.
func (s *Vector[T]) String() string {
	return fmt.Sprint(s.data)
}

// Sort the list using the natural ordering of its elements.
func Sort[T types.Comparable[T]](list *Vector[T]) {
	sort.Slice(list.data, func(i, j int) bool {
		return list.data[i].Less(list.data[j])
	})
}

// SortBy sorts the list using the function less for comparison of two element . if less(a,b) = true then it means a comes before b in the sorted list.
func SortBy[T types.Equitable[T]](list *Vector[T], less func(a, b T) bool) {
	sort.Slice(list.data, func(i, j int) bool {
		return less(list.data[i], list.data[j])
	})
}
