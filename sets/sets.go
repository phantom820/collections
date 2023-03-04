package sets

import (
	"reflect"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/sets/hashset"
	"github.com/phantom820/collections/sets/linkedhashset"
	"github.com/phantom820/collections/sets/treeset"
)

const (
	UNION        = 0
	INTERSECTION = 1
	DIFFERENCE   = 2
)

// SetView an unmodifiable view of a set which is backed by other sets, this view will change as the backing sets change.
type SetView[T comparable] struct {
	a    collections.Collection[T]
	b    collections.Collection[T]
	view int
}

// Type return the type of set view (union, intersection , ...)
func (setView SetView[T]) Type() int {
	return setView.view
}

// Len returns the number of elements in the set. This needs to be calculated based on the backing sets of the set view.
func (setView SetView[T]) Len() int {
	switch setView.view {
	case UNION:
		return unionLen(setView.a, setView.b)
	case INTERSECTION:
		return intersectionLen(setView.a, setView.b)
	case DIFFERENCE:
		return differenceLen(setView.a, setView.b)
	default:
		panic("")
	}
}

// unionLen computes the size of the set union of sets a and b.
func unionLen[T comparable](a collections.Collection[T], b collections.Collection[T]) int {
	// iterate over small set and skip counted elements.
	if a.Len() <= b.Len() {
		len := b.Len()
		it := a.Iterator()
		for it.HasNext() {
			if !b.Contains(it.Next()) {
				len++
			}
		}
		return len
	}
	// b is the smaller set so swap around.
	return unionLen(b, a)
}

// intersectionLen computes the size of the set intersection of sets a and b.
func intersectionLen[T comparable](a collections.Collection[T], b collections.Collection[T]) int {
	//iterate over smaller set and count common elements.
	if a.Len() <= b.Len() {
		len := 0
		it := a.Iterator()
		for it.HasNext() {
			if b.Contains(it.Next()) {
				len++
			}
		}
		return len
	}
	// b is the smaller set so swap around.
	return intersectionLen(b, a)
}

// differenceLen computes the size of the set difference of sets a and b (a-b).
func differenceLen[T comparable](a collections.Collection[T], b collections.Collection[T]) int {
	it := a.Iterator()
	len := 0
	for it.HasNext() {
		if !b.Contains(it.Next()) {
			len++
		}
	}
	return len
}

// Empty returns true if the set contains no elements. This needs to be calculated based on the backing sets.
func (setView SetView[T]) Empty() bool {
	return setView.Len() == 0
}

// unionForEach forEach operation on a union view.
func unionForEach[T comparable](a collections.Collection[T], b collections.Collection[T], f func(T)) {
	// iterate over both sets and avoid applying twice.
	it1 := a.Iterator()
	for it1.HasNext() {
		f(it1.Next())
	}
	it2 := b.Iterator()
	for it2.HasNext() {
		e := it2.Next()
		if !a.Contains(e) {
			f(e)
		}
	}
}

// intersectionForEach forEach operation on an intersection view.
func intersectionForEach[T comparable](a collections.Collection[T], b collections.Collection[T], f func(T)) {
	// iterate over smaller set and apply on common elements.
	if a.Len() <= b.Len() {
		it := a.Iterator()
		for it.HasNext() {
			e := it.Next()
			if b.Contains(e) {
				f(e)
			}
		}
		return
	}
	// otherwise swap around since b is the smaller set.
	intersectionForEach(b, a, f)
}

// differenceForEach forEach operation on an difference set view.
func differenceForEach[T comparable](a collections.Collection[T], b collections.Collection[T], f func(T)) {
	it := a.Iterator()
	for it.HasNext() {
		e := it.Next()
		if !b.Contains(e) {
			f(e)
		}
	}
}

// ForEach performs the given action for each element of the set view.
func (setView SetView[T]) ForEach(f func(T)) {

	switch setView.view {
	case UNION:
		unionForEach(setView.a, setView.b, f)
		return
	case INTERSECTION:
		intersectionForEach(setView.a, setView.b, f)
		return
	case DIFFERENCE:
		differenceForEach(setView.a, setView.b, f)
	default:
		panic("")
	}

}

// ToSlice returns a slice containing all the elements in the set view.
func (setView SetView[T]) ToSlice() []T {
	slice := make([]T, 0)
	setView.ForEach(func(t T) { slice = append(slice, t) })
	return slice
}

// ToHashSet returns a [HashSet] with all the elements from the set view.
func (setView SetView[T]) ToHashSet() *hashset.HashSet[T] {
	set := hashset.New[T]()
	setView.ForEach(func(t T) {
		set.Add(t)
	})
	return set
}

// ToLinkedHashSet returns a [LinkedHashSet] with all the elements from the set view.
func (setView SetView[T]) ToLinkedHashSet() *linkedhashset.LinkedHashSet[T] {
	set := linkedhashset.New[T]()
	setView.ForEach(func(t T) {
		set.Add(t)
	})
	return set
}

// ToTreeSet returns a [TreeSet] with all the elements from the set view.
func (setView SetView[T]) ToTreeSet(lessThan func(e1, e2 T) bool) *treeset.TreeSet[T] {
	set := treeset.New(lessThan)
	setView.ForEach(func(t T) {
		set.Add(t)
	})
	return set
}

// Contains returns true if the set view contains the specified element.
func (setView SetView[T]) Contains(e T) bool {
	switch setView.view {
	case UNION:
		return setView.a.Contains(e) || setView.b.Contains(e)

	case INTERSECTION:
		return setView.a.Contains(e) && setView.b.Contains(e)

	case DIFFERENCE:
		return setView.a.Contains(e) && !setView.b.Contains(e)

	default:
		panic("")
	}
}

// Iterator returns an iterator over the elements in the set view.
func (setView SetView[T]) Iterator() collections.Iterator[T] {
	return nil
}

// IsSet returns true if the given collection is a set.
func IsSet[T comparable](c collections.Collection[T]) bool {

	if c == nil || reflect.ValueOf(c).IsNil() {
		return false
	}

	switch c.(type) {
	case *hashset.HashSet[T]:
		return true
	case *hashset.ImmutableHashSet[T]:
		return true
	case *linkedhashset.LinkedHashSet[T]:
		return true
	case *linkedhashset.ImmutableLinkedHashSet[T]:
		return true
	case *treeset.TreeSet[T]:
		return true
	case *treeset.ImmutableTreeSet[T]:
		return true
	default:
		return false
	}
}

// Union returns an unmodifiable view of the union of two sets. Will panic if any of the given arguments is not a set.
func Union[T comparable](a collections.Collection[T], b collections.Collection[T]) SetView[T] {
	if !IsSet(a) || !IsSet(b) {
		panic("")
	}
	return SetView[T]{a: a, b: b, view: UNION}

}

// Intersection returns an unmodifiable view of the intersection of two sets. Will panic if any of the given arguments is not a set.
func Intersection[T comparable](a collections.Collection[T], b collections.Collection[T]) SetView[T] {
	if !IsSet(a) || !IsSet(b) {
		panic("")
	}
	return SetView[T]{a: a, b: b, view: INTERSECTION}
}

// Difference returns an unmodifiable view of the difference of two sets. Will panic if any of the given arguments is not a set.
func Difference[T comparable](a collections.Collection[T], b collections.Collection[T]) SetView[T] {
	if !IsSet(a) || !IsSet(b) {
		panic("")
	}
	return SetView[T]{a: a, b: b, view: DIFFERENCE}
}
