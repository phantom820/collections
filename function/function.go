package function

import (
	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/sets/hashset"
	"github.com/phantom820/collections/sets/linkedhashset"
	"github.com/phantom820/collections/sets/treeset"
	"github.com/phantom820/collections/types/optional"
)

type View[T comparable] interface {
	iterable.Iterable[T]
	ForEach(f func(T))
	ToVector() *vector.Vector[T]
	ToSlice() []T
	ToHashSet() *hashset.HashSet[T]
	ToLinkedHashSet() *linkedhashset.LinkedHashSet[T]
	ToTreeSet(lessThan func(k1, k2 T) bool) *treeset.TreeSet[T]
	ToLinkedList() *linkedlist.LinkedList[T]
	ToForwardList() *forwardlist.ForwardList[T]
}

type view[T comparable] struct {
	iterator func() iterator.Iterator[T]
}

func (view *view[T]) Iterator() iterator.Iterator[T] {
	return view.iterator()
}

func (view *view[T]) ForEach(f func(T)) {
	it := view.iterator()
	for it.HasNext() {
		f(it.Next())
	}
}

func (view *view[T]) ToVector() *vector.Vector[T] {
	vector := vector.New[T]()
	it := view.iterator()
	for it.HasNext() {
		vector.Add(it.Next())
	}
	return vector
}

func (view *view[T]) ToSlice() []T {
	it := view.iterator()
	slice := make([]T, 0)
	for it.HasNext() {
		slice = append(slice, it.Next())
	}
	return slice
}

func (view *view[T]) ToLinkedList() *linkedlist.LinkedList[T] {
	it := view.iterator()
	list := linkedlist.New[T]()
	for it.HasNext() {
		list.Add(it.Next())
	}
	return list
}

func (view *view[T]) ToForwardList() *forwardlist.ForwardList[T] {
	it := view.iterator()
	list := forwardlist.New[T]()
	for it.HasNext() {
		list.Add(it.Next())
	}
	return list
}

func (view *view[T]) ToHashSet() *hashset.HashSet[T] {
	it := view.iterator()
	set := hashset.New[T]()
	for it.HasNext() {
		set.Add(it.Next())
	}
	return set
}

func (view *view[T]) ToLinkedHashSet() *linkedhashset.LinkedHashSet[T] {
	it := view.iterator()
	set := linkedhashset.New[T]()
	for it.HasNext() {
		set.Add(it.Next())
	}
	return set
}

func (view *view[T]) ToTreeSet(lessThan func(k1, k2 T) bool) *treeset.TreeSet[T] {
	it := view.iterator()
	set := treeset.New(lessThan)
	for it.HasNext() {
		set.Add(it.Next())
	}
	return set
}

func Filter[T comparable](iterable iterable.Iterable[T], f func(T) bool) View[T] {
	return &view[T]{
		iterator: func() iterator.Iterator[T] {
			return iterator.Filter(iterable.Iterator(), f)
		},
	}
}

func Identity[T comparable](iterable iterable.Iterable[T]) View[T] {
	return &view[T]{
		iterator: func() iterator.Iterator[T] {
			return iterable.Iterator()
		},
	}
}

func Map[T comparable, U comparable](iterable iterable.Iterable[T], f func(T) U) View[U] {
	return &view[U]{
		iterator: func() iterator.Iterator[U] {
			return iterator.Map(iterable.Iterator(), f)
		},
	}
}

func Reduce[T comparable](iterable iterable.Iterable[T], f func(x, y T) T) optional.Optional[T] {
	return iterator.Reduce(iterable.Iterator(), f)
}

func GroupBy[T comparable, U comparable](iterable iterable.Iterable[T], f func(T) U) hashmap.HashMap[U, []T] {
	it := iterable.Iterator()
	groups := hashmap.New[U, []T]()
	for it.HasNext() {
		element := it.Next()
		key := f(element)
		if group, ok := groups[key]; ok {
			group = append(group, element)
			groups[key] = group
		} else {
			group = []T{element}
			groups[key] = group
		}
	}
	return groups
}
