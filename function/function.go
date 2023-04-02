// package function defines common transformer functions for collections such as Filter, Map and Reduce. These functions do not immediately yield a collection
// but produce a view which can then be materialized to the desired collection.
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

// View represents a proxy for a base collection that is being transformed to derive another collectio. A view
// is materialized when it is converted to a specific collection.
type View[T comparable] interface {
	iterable.Iterable[T]
	ForEach(f func(T))                                          // Performs the given action on each element in the view.
	Filter(f func(T) bool) View[T]                              // Returns a view with all elements that satisfy the given predicate.
	Reduce(f func(x, y T) T) optional.Optional[T]               // Reduces the elements of the view using the associative binary function and returns result as an option.
	Map(f func(T) T) View[T]                                    // Returns the view obtained from applying the transformation function to every element of the view.
	ToVector() *vector.Vector[T]                                // Materializes the view to a [Vector].
	ToSlice() []T                                               // Materializes the view to a slice.
	ToHashSet() *hashset.HashSet[T]                             // Materializes the view to a [HashSet].
	ToLinkedHashSet() *linkedhashset.LinkedHashSet[T]           // Materializes the view to a [LinkedHashSet].
	ToTreeSet(lessThan func(k1, k2 T) bool) *treeset.TreeSet[T] // Materializes the view to a [TreeSet].
	ToLinkedList() *linkedlist.LinkedList[T]                    // Materializes the view to a [LinkedList].
	ToForwardList() *forwardlist.ForwardList[T]                 // Materializes the view to a [ForwardList].
}

// view represents a proxy of a collection that is being transformed.
type view[T comparable] struct {
	iterator func() iterator.Iterator[T]
}

// Iterator returns an iterator over the view elements.
func (view *view[T]) Iterator() iterator.Iterator[T] {
	return view.iterator()
}

// ForEach performs the given action for each element of the view.
func (view *view[T]) ForEach(f func(T)) {
	it := view.iterator()
	for it.HasNext() {
		f(it.Next())
	}
}

// Map returns the view obtained from applying the transformation function to every element of the view.
func (_view *view[T]) Map(f func(T) T) View[T] {
	return &view[T]{
		iterator: func() iterator.Iterator[T] {
			return iterator.Map(_view.Iterator(), f)
		},
	}
}

// Filter returns a view with all elements that satisfy the given predicate.
func (_view *view[T]) Filter(f func(T) bool) View[T] {
	return &view[T]{
		iterator: func() iterator.Iterator[T] {
			return iterator.Filter(_view.Iterator(), f)
		},
	}
}

// Reduce reduces the elements of the view using the associative binary function and returns result as an option.
func (_view *view[T]) Reduce(f func(T, T) T) optional.Optional[T] {
	return iterator.Reduce(_view.iterator(), f)
}

// ToVector materializes the view to a [Vector].
func (view *view[T]) ToVector() *vector.Vector[T] {
	vector := vector.New[T]()
	it := view.iterator()
	for it.HasNext() {
		vector.Add(it.Next())
	}
	return vector
}

// ToSlice materializes the view to a slice.
func (view *view[T]) ToSlice() []T {
	it := view.iterator()
	slice := make([]T, 0)
	for it.HasNext() {
		slice = append(slice, it.Next())
	}
	return slice
}

// ToLinkedList materializes the view to a [LinkedList].
func (view *view[T]) ToLinkedList() *linkedlist.LinkedList[T] {
	it := view.iterator()
	list := linkedlist.New[T]()
	for it.HasNext() {
		list.Add(it.Next())
	}
	return list
}

// ToForwardList materializes the view to a [ForwardList].
func (view *view[T]) ToForwardList() *forwardlist.ForwardList[T] {
	it := view.iterator()
	list := forwardlist.New[T]()
	for it.HasNext() {
		list.Add(it.Next())
	}
	return list
}

// ToHashSet materializes the view to a [HashSet].
func (view *view[T]) ToHashSet() *hashset.HashSet[T] {
	it := view.iterator()
	set := hashset.New[T]()
	for it.HasNext() {
		set.Add(it.Next())
	}
	return set
}

// ToHashSet materializes the view to a [LinkedHashSet].
func (view *view[T]) ToLinkedHashSet() *linkedhashset.LinkedHashSet[T] {
	it := view.iterator()
	set := linkedhashset.New[T]()
	for it.HasNext() {
		set.Add(it.Next())
	}
	return set
}

// ToTreeSet materializes the view to a [TreeSet].
func (view *view[T]) ToTreeSet(lessThan func(k1, k2 T) bool) *treeset.TreeSet[T] {
	it := view.iterator()
	set := treeset.New(lessThan)
	for it.HasNext() {
		set.Add(it.Next())
	}
	return set
}

// Filter returns a view with all elements that satisfy the given predicate.
func Filter[T comparable](iterable iterable.Iterable[T], f func(T) bool) View[T] {
	return &view[T]{
		iterator: func() iterator.Iterator[T] {
			return iterator.Filter(iterable.Iterator(), f)
		},
	}
}

// Identity returns a view that is identical to the given iterable.
func Identity[T comparable](iterable iterable.Iterable[T]) View[T] {
	return &view[T]{
		iterator: func() iterator.Iterator[T] {
			return iterable.Iterator()
		},
	}
}

// Map returns a view obtained from applying the transformation function to every element on the given iterable.
func Map[T comparable, U comparable](iterable iterable.Iterable[T], f func(T) U) View[U] {
	return &view[U]{
		iterator: func() iterator.Iterator[U] {
			return iterator.Map(iterable.Iterator(), f)
		},
	}
}

// Reduce reduces the elements of the iterable using the associative binary function and returns result as an option.
func Reduce[T comparable](iterable iterable.Iterable[T], f func(x, y T) T) optional.Optional[T] {
	return iterator.Reduce(iterable.Iterator(), f)
}

// GroupBy returns a grouping of elements from the iterable using the given discriminator function.
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
