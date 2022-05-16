package sort

import (
	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/types"
)

// Sort sorts the given collection using the natural define ordering of its elements.
func Sort[T types.Comparable[T]](collection collections.Collection[T]) {

	switch t := collection.(type) {
	case *forwardlist.ForwardList[T]:
		forwardlist.Sort(t)
		return
	case *list.List[T]:
		list.Sort(t)
	default:
		return
	}
	// var slice slice[T] = collection.Collect() // linear time to collect all members into a slice O(n).
	// sort.Sort(slice)                          // log linear time to sort O(nlogn)
	// collection.Clear()                        // constant time O(1)
	// collection.AddSlice(slice)                // linear time O(n) resulting in overall time complexity O(nlogn)
}

// Sort sorts the given collection using the function less for ordering if less(a,b) = true, then a comes before b in the sorted collection
func SortBy[T types.Equitable[T]](collection collections.Collection[T], less func(a, b T) bool) {

	switch t := collection.(type) {
	case *forwardlist.ForwardList[T]:
		forwardlist.SortBy(t, less)
		return
	default:
		return
	}
	// var slice slice[T] = collection.Collect() // linear time to collect all members into a slice O(n).
	// sort.Sort(slice)                          // log linear time to sort O(nlogn)
	// collection.Clear()                        // constant time O(1)
	// collection.AddSlice(slice)                // linear time O(n) resulting in overall time complexity O(nlogn)
}
