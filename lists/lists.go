package lists

import (
	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
)

// newMutable return a new list that is mutable that has an underlying type derived from the given list.
func newMutable[T comparable, U comparable](l collections.List[T]) collections.List[U] {
	switch l.(type) {
	case *vector.Vector[T]:
		return vector.New[U]()
	case *vector.ImmutableVector[T]:
		return vector.New[U]()
	case *forwardlist.ForwardList[T]:
		return forwardlist.New[U]()
	case *forwardlist.ImmutableForwadList[T]:
		return forwardlist.New[U]()
	case *linkedlist.LinkedList[T]:
		return linkedlist.New[U]()
	default:
		panic("")
	}
}

// Partition returns consecutive sublists of the list, each of the same size, the last list may be smaller.
func Partition[T comparable](list collections.List[T], size int) []collections.List[T] {
	if list.Empty() {
		return []collections.List[T]{}
	}

	it := list.Iterator()
	subList := newMutable[T, T](list)
	subLists := make([]collections.List[T], 0)
	for it.HasNext() {
		if subList.Len() < size {
			subList.Add(it.Next())
		} else {
			subLists = append(subLists, subList)
			subList = newMutable[T, T](subList)
			subList.Add(it.Next())
		}
	}

	if !subList.Empty() {
		subLists = append(subLists, subList)
	}

	return subLists
}

// Filter filters the given list and returns a new list containing only elements that satisfy the given predicate.
func Filter[T comparable](list collections.List[T], f func(T) bool) collections.List[T] {
	it := list.Iterator()
	newList := newMutable[T, T](list)
	for it.HasNext() {
		e := it.Next()
		if f(e) {
			newList.Add(e)
		}
	}
	return newList
}

// Map returns a new list obtained by applying the given mapping on members of the given list.
func Map[T comparable, U comparable](list collections.List[T], f func(T) U) collections.List[U] {
	it := list.Iterator()
	newList := newMutable[T, U](list)
	for it.HasNext() {
		e := it.Next()
		newList.Add(f(e))
	}
	return newList
}

// Reduce returns the result of applying binary operator on members of the list. The operator should be associative.
func Reduce[T comparable](list collections.List[T], f func(x, y T) T) (T, bool) {
	if list.Empty() {
		var zero T
		return zero, false
	}
	it := list.Iterator()
	x := it.Next()
	for it.HasNext() {
		y := it.Next()
		x = f(x, y)
	}

	return x, true
}
