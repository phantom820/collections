package sort

import (
	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/lists/vector"
	"github.com/phantom820/collections/types"
)

// Sort sorts the given collection using the natural ordering of its elements. This should only be used with collections that can actually be sorted
// such as a list or a vector where sorting operation is well defined. In case of stacks,queues, sets and so on this will not do anything to the collection.
func Sort[T types.Comparable[T]](collection collections.Collection[T]) {

	switch t := collection.(type) {
	case *forwardlist.ForwardList[T]:
		forwardlist.Sort(t) // Uses merge sort on a singly linked list. See forwardlist package. O(nlogn)
		return
	case *list.List[T]:
		list.Sort(t) // Uses merge sort on a doubly linked list. See forwardlist package. O(nlogn)
		return
	case *vector.Vector[T]:
		vector.Sort(t)
		return
	default:
		return
	}
}

// Sort sorts the given collection using the function less for ordering if less(a,b) = true, then a comes before b in the sorted collection.
// This should only be used with collections that can actually be sorted such as a list or a vector where sorting operation is well defined.
// In case of stacks,queues, sets and so on this will not do anything to the collection.
func SortBy[T types.Equitable[T]](collection collections.Collection[T], less func(a, b T) bool) {

	switch t := collection.(type) {
	case *forwardlist.ForwardList[T]:
		forwardlist.SortBy(t, less) // Uses merge sort on a singly linked list. See forwardlist package. O(nlogn)
		return
	case *list.List[T]:
		list.SortBy(t, less) // Uses merge sort on a double linked lis. See list package.  O(nlogn)
	case *vector.Vector[T]:
		vector.SortBy(t, less) // Uses stdlib sort package since underlying container is a slice.
		return
	default:
		return
	}

}
