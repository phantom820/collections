package vector

import (
	"github.com/phantom820/collections"
	"github.com/phantom820/collections/sets/hashset"
)

type Vector[T comparable] struct {
	data []T
}

func New[T comparable]() *Vector[T] {
	return &Vector[T]{data: make([]T, 0, 16)}
}

// Of creates a list with the given elements.
func Of[T comparable](elements ...T) Vector[T] {
	list := Vector[T]{data: make([]T, 0, len(elements))}
	list.data = append(list.data, elements...)
	return list
}

// Add appends the specified element to the end of this list.
func (list *Vector[T]) Add(e T) bool {
	list.data = append(list.data, e)
	return true
}

// Len the number of elements in this list.
func (list *Vector[T]) Len() int {
	return len(list.data)
}

// AddAll adds all of the elements in the specified iterable to the set.
func (list *Vector[T]) AddAll(iterable collections.Iterable[T]) bool {
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

// At returns the element at the specified position in the list.
func (list *Vector[T]) At(i int) T {
	if i < 0 || i >= list.Len() {
		panic("")
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

// Remove removes the first occurrence of the specified element from this list, if it is present.
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
func (list *Vector[T]) RemoveAll(iterable collections.Iterable[T]) bool {
	if list.Empty() {
		return false
	}
	// introduce a set so we can make the lookups fast, also passing a collection here introduces
	// uncertainty about performance of contains so we just need an iterable and enforce the set.
	set := hashset.New[T]()
	it := iterable.Iterator()
	for it.HasNext() {
		set.Add(it.Next())
	}
	return list.RemoveIf(func(t T) bool { return set.Contains(t) })
}

// RemoveSlice removes all of the list elements that are also contained in the specified slice.
func (list *Vector[T]) RemoveSlice(s []T) bool {
	if list.Empty() {
		return false
	}
	// introduce a set so we can make the lookups fast, also passing a collection here introduces
	// uncertainty about performance of contains so we just need an iterable and enforce the set.
	set := hashset.New[T]()
	set.AddSlice(s)
	return list.RemoveIf(func(t T) bool { return set.Contains(t) })
}

// RetainAll retains only the elements in the list that are contained in the specified collection.
func (list *Vector[T]) RetainAll(c collections.Collection[T]) bool {
	if list.Empty() {
		return false
	}
	// create a predicate that removes elements that are not in the passed collection.
	// performance here is mainly affected by how the given collection performs with contains.
	return list.RemoveIf(func(t T) bool { return !c.Contains(t) })
}

// ForEach performs the given action for each element of the list.
func (list *Vector[T]) ForEach(f func(T)) {
	for _, e := range list.data {
		f(e)
	}
}

// ToSlice returns the underlying slice.
func (list *Vector[T]) ToSlice() []T {
	return list.data
}

// Iterator returns an iterator over the elements in the list.
func (list *Vector[T]) Iterator() collections.Iterator[T] {
	return &iterator[T]{initialized: false, initialize: func() []T { return list.data }, index: 0, data: nil}
}

// iterator implememantation for [Vector].
type iterator[T comparable] struct {
	initialized bool
	initialize  func() []T
	index       int
	data        []T
}

// HasNext returns true if the iterator has more elements.
func (it *iterator[T]) HasNext() bool {
	if !it.initialized {
		it.data = it.initialize()
		it.initialized = true
	} else if it.data == nil {
		return false
	}
	return it.index < len(it.data)
}

// Next returns the next element in the iterator.
func (it *iterator[T]) Next() T {
	if !it.HasNext() {
		panic("iterator things shoould panic here")
	}
	index := it.index
	it.index++
	return it.data[index]
}
