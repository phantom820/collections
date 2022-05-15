package sort

import (
	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/types"
)

func Sort[T types.Comparable[T]](collection collections.Collection[T]) {

	switch t := collection.(type) {
	case *forwardlist.ForwardList[T]:
		// fmt.Println(t)
		forwardlist.Sort(t)
		// fmt.Println(t)
	default:
		return
	}
	// var slice slice[T] = collection.Collect() // linear time to collect all members into a slice O(n).
	// sort.Sort(slice)                          // log linear time to sort O(nlogn)
	// collection.Clear()                        // constant time O(1)
	// collection.AddSlice(slice)                // linear time O(n) resulting in overall time complexity O(nlogn)
}
