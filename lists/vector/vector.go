// package vector defines a wrapper over a standard slice.
package vector

import (
	"fmt"
	"sort"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/sets"
	"github.com/phantom820/collections/types/optional"
)

// Vector a wrapper around a slice []T.
type Vector[T comparable] struct {
	data []T
}

// New creates a mutable list with the given elements.
func New[T comparable](elements ...T) *Vector[T] {
	list := Vector[T]{data: make([]T, 0)}
	for _, e := range elements {
		list.Add(e)
	}
	return &list
}

// Of creates an immutable list with the given elements.
func Of[T comparable](elements ...T) ImmutableVector[T] {
	list := Vector[T]{data: make([]T, 0, len(elements))}
	list.data = append(list.data, elements...)
	return ImmutableVector[T]{vector: list}
}

// Add appends the specified element to the end of the list.
func (list *Vector[T]) Add(e T) bool {
	list.data = append(list.data, e)
	return true
}

// Len returns the number of elements in the list.
func (list *Vector[T]) Len() int {
	return len(list.data)
}

// AddAll adds all of the elements in the specified iterable to the list.
func (list *Vector[T]) AddAll(iterable iterable.Iterable[T]) bool {
	it := iterable.Iterator()
	for it.HasNext() {
		list.data = append(list.data, it.Next())
	}
	return true
}

// AddSlice adds all the elements in the slice to the list.
func (list *Vector[T]) AddSlice(s []T) bool {
	list.data = append(list.data, s...)
	return true
}

// AddAt inserts the specified element at the specified index in the list.
func (list *Vector[T]) AddAt(i int, e T) {
	if i < 0 || i >= list.Len() && !list.Empty() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	} else if i == 0 {
		data := make([]T, 0, list.Len()+1)
		data = append(data, e)
		list.data = append(data, list.data...)
		return
	} else if i == list.Len()-1 {
		list.Add(e)
		return
	}
	left := list.data[:i]
	right := list.data[i:]
	data := make([]T, 0, list.Len()+1)
	data = append(data, left...)
	data = append(data, e)
	data = append(data, right...)
	list.data = data
}

// At returns the element at the specified index in the list.
func (list *Vector[T]) At(i int) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	}
	return list.data[i]
}

// Contains returns true if the list contains the specified element.
func (list *Vector[T]) Contains(e T) bool {
	for i := range list.data {
		if list.data[i] == e {
			return true
		}
	}
	return false
}

// Clear removes all of the elements from the list.
func (list *Vector[T]) Clear() {
	list.data = nil
	list.data = make([]T, 0)
}

// Empty returns true if the list contains no elements.
func (list *Vector[T]) Empty() bool {
	return list.Len() == 0
}

// indexOf returns the index of the given element if it is present in the given slice.
func indexOf[T comparable](data []T, e T) int {
	index := -1
	for i := range data {
		if data[i] == e {
			index = i
			break
		}
	}
	return index
}

// IndexOf returns the index of the first occurrence of the specified element in the list.
func (list *Vector[T]) IndexOf(e T) optional.Optional[int] {
	index := indexOf(list.data, e)
	if index == -1 {
		return optional.Empty[int]()
	}
	return optional.Of(index)
}

// removeAt removes the element at the specified index.
func removeAt[T comparable](data []T, index int) []T {
	if index == 0 { // remove first element.
		if len(data) == 1 {
			return []T{}
		}
		return data[1:]
	} else if index == len(data)-1 { // remove the last element.
		return data[:index]
	}
	// shift every element after the index one back.
	for i := index; i < len(data)-1; i++ {
		data[i] = data[i+1]
	}
	// chop of the end to reduce the length.
	return data[:len(data)-1]
}

// Remove removes the first occurrence of the specified element from the list, if it is present.
func (list *Vector[T]) Remove(e T) bool {
	index := indexOf(list.data, e)
	if index == -1 {
		return false
	}
	list.data = removeAt(list.data, index)
	return true
}

// shift shifts the elements between the stat and end indices. The shift here moves all elements between [start, end] to
// [end, end + (end - start)] .
func shift[T comparable](data []T, start int, end int) {
	for i := 0; i < end-start; i++ {
		if end+i < len(data) {
			temp := data[start+i]
			data[start+i] = data[end+i]
			data[end+i] = temp
		}
	}
}

// RemoveIf removes all of the elements of the list that satisfy the given predicate.
func (list *Vector[T]) RemoveIf(f func(T) bool) bool {
	tail := len(list.data)
	// move all elements that do not satisfy the predicate to the front and then chop of the tail with elements to remove.
	// we first identify an element that is to be removed , then find next element to be kept and shift elements around to
	// bring the element to be kept to index of element to be removed.
	for i, _ := range list.data {
		if f(list.data[i]) {
			start := i
			end := i
			for j := i; j < len(list.data); j++ {
				if !f(list.data[j]) {
					end = j
					break
				}
			}

			if start == end && end != len(list.data)-1 {
				temp := list.data[start]
				list.data[start] = list.data[start+1]
				list.data[start+1] = temp
				tail = end
				break
			}
			tail = end
			shift(list.data, start, end)
		}
		i++
	}

	n := len(list.data)
	list.data = list.data[:tail]
	return n != len(list.data)
}

// RemoveAll removes from the list all of its elements that are contained in the specified collection.
func (list *Vector[T]) RemoveAll(iterable iterable.Iterable[T]) bool {
	if list.Empty() {
		return false
	} else if sets.IsSet(iterable) {
		set := iterable.(collections.Set[T])
		return list.RemoveIf(func(e T) bool {
			return set.Contains(e)
		})
	}

	set := make(map[T]struct{})
	it := iterable.Iterator()
	for it.HasNext() {
		set[it.Next()] = struct{}{}
	}
	return list.RemoveIf(func(t T) bool {
		_, ok := set[t]
		return ok
	})
}

// RemoveSlice removes all of the list elements that are also contained in the specified slice.
func (list *Vector[T]) RemoveSlice(s []T) bool {
	if list.Empty() {
		return false
	}
	set := make(map[T]struct{})
	for _, e := range s {
		set[e] = struct{}{}
	}
	return list.RemoveIf(func(e T) bool {
		_, ok := set[e]
		return ok
	})
}

// ImmutableCopy returns an immutable copy of the list.
func (list *Vector[T]) ImmutableCopy() ImmutableVector[T] {
	return Of(list.data...)
}

// Copy returns a copy of the list.
func (list *Vector[T]) Copy() *Vector[T] {
	copy := New(list.data...)
	return copy
}

// SubList returns a copy of the portion of the list between the specified start and end indices (exclusive).
func (list *Vector[T]) SubList(start int, end int) *Vector[T] {
	if start < 0 || start >= list.Len() {
		panic(errors.IndexOutOfBounds(start, list.Len()))
	} else if end < 0 || end > list.Len() {
		panic(errors.IndexOutOfBounds(end, list.Len()))
	} else if start > end {
		panic(errors.IndexBoundsOutOfRange(start, end))
	} else if start == end {
		return &Vector[T]{data: make([]T, 0)}
	}
	subData := make([]T, end-start)
	copy(subData, list.data[start:end])
	return &Vector[T]{data: subData}
}

// RetainAll retains only the elements in the list that are contained in the specified collection.
func (list *Vector[T]) RetainAll(c collections.Collection[T]) bool {
	if list.Empty() {
		return false
	} else if sets.IsSet[T](c) {
		return list.RemoveIf(func(t T) bool { return !c.Contains(t) })
	}
	set := make(map[T]struct{})
	it := c.Iterator()
	for it.HasNext() {
		set[it.Next()] = struct{}{}
	}
	return list.RemoveIf(func(e T) bool {
		_, ok := set[e]
		return !ok
	})
}

// ForEach performs the given action for each element of the list.
func (list *Vector[T]) ForEach(f func(T)) {
	for _, e := range list.data {
		f(e)
	}
}

// Set replaces the element at the specified index in the list with the specified element.
func (list *Vector[T]) Set(i int, e T) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	}
	temp := list.data[i]
	list.data[i] = e
	return temp
}

// ToSlice returns a slice containing the elements of the list.
func (list *Vector[T]) ToSlice() []T {
	return list.data
}

// Equals returns true if the list is equivalent to the given list. Two lists are equal if they have the same size
// and contain the same elements in the same order.
func (list *Vector[T]) Equals(other collections.List[T]) bool {
	if list == other {
		return true
	} else if list.Len() != other.Len() {
		return false
	}
	it := other.Iterator()
	for i := range list.data {
		if list.data[i] != it.Next() {
			return false
		}
	}
	return true

}

// RemoveAt removes the element at the specified position in the list.
func (list *Vector[T]) RemoveAt(i int) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	} else if i == 0 {
		temp := list.data[i]
		list.data = list.data[1:]
		return temp
	} else if i == list.Len()-1 {
		temp := list.data[i]
		list.data = list.data[:list.Len()-1]
		return temp
	}
	temp := list.data[i]
	list.data = removeAt(list.data, i)
	return temp
}

// Iterator returns an iterator over the elements in the list.
func (list *Vector[T]) Iterator() iterator.Iterator[T] {
	return &listIterator[T]{initialized: false, initialize: func() []T { return list.data }, index: 0, data: nil}
}

// iterator implememantation for [Vector].
type listIterator[T comparable] struct {
	initialized bool
	initialize  func() []T
	index       int
	data        []T
}

// HasNext returns true if the iterator has more elements.
func (it *listIterator[T]) HasNext() bool {
	if !it.initialized {
		it.data = it.initialize()
		it.initialized = true
	} else if it.data == nil {
		return false
	}
	return it.index < len(it.data)
}

// Next returns the next element in the iterator.
func (it *listIterator[T]) Next() T {
	if !it.HasNext() {
		panic(errors.NoSuchElement())
	}
	index := it.index
	it.index++
	return it.data[index]
}

// String returns the string representation of the list.
func (list Vector[T]) String() string {
	return fmt.Sprint(list.data)
}

// Sort sorts the list using the given less function. if less(a,b) = true then a would be before b in a sorted list.
func (list *Vector[T]) Sort(less func(a, b T) bool) {
	sort.Slice(list.data, func(i, j int) bool {
		return less(list.data[i], list.data[j])
	})
}
