package lists

import (
	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
)

type List[T comparable] interface {
	collections.Collection[T]
	AddAt(i int, e T)
	At(i int) T
	Set(i int, e T) T
	RemoveAt(i int) T
	// Equals(other List[T]) bool
	// SubList(start int, end int) List[T]
	// Copy() List[T]
}

// Equal returns true if the given lists are equal. Two list are equal if they are the same reference or have the same size an d
// the same elements in the same order.
func Equal[T comparable](l1, l2 List[T]) bool {
	if l1 == l2 {
		return true
	} else if l1.Len() != l2.Len() {
		return false
	}
	it1, it2 := l1.Iterator(), l2.Iterator()
	_, _ = it1.HasNext(), it2.HasNext() // initializes each iterator.
	for it1.HasNext() {
		if it1.Next() != it2.Next() {
			return false
		}
	}
	return true
}

func newMutable[T comparable, U comparable](l List[T]) List[U] {
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

func Partition[T comparable](n int) []List[T] {
	return nil
}

func Filter[T comparable](list List[T], f func(T) bool) List[T] {
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

func Map[T comparable, U comparable](list List[T], f func(T) U) List[U] {
	it := list.Iterator()
	newList := newMutable[T, U](list)
	for it.HasNext() {
		e := it.Next()
		newList.Add(f(e))
	}
	return newList
}

func Reduce[T comparable, U any](list List[T], f func(T, T) U) *U {
	return nil
}

// func ImmutableCopyOf[T comparable](list List[T]) List[T] {
// 	immutableList = vector.Of[T]()
// }
