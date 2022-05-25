// Package collections provides an interface that a data structure must implement to count as a collection/ a tree.
package collections

import (
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types"
)

// Collection a blue print for which methods a collection must implement. A collection is a linear
// data structures such a linked list, queue, stack and a set.
type Collection[T types.Equitable[T]] interface {
	iterator.Iterable[T]              // Returns an iterator for iterating through the collection.
	Add(elements ...T) bool           // Adds elements to the collection.
	AddAll(c iterator.Iterable[T])    // Adds all elements from another collection into the collection.
	Len() int                         // Returns the size (number of items) stored in the collection.
	Contains(e T) bool                // Checks if the element e is a member of the collection.
	Remove(elements ...T) bool        // Tries to remove the specified element(s) from the collection. Only first occurence of an element is removed.
	RemoveAll(c iterator.Iterable[T]) // Removes all elements from another collections that appear in the collection.
	Empty() bool                      // Checks if the collection contains any elements.
	Clear()                           // Removes all elements in the collection.
}

// slice type used for allowing sort using defined ordering.
type slice[T types.Comparable[T]] []T

// Len returns size of underlying slice.
func (slice slice[T]) Len() int {
	return len(slice)
}

// Swap swaps elements in the slice.
func (slice slice[T]) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Less defines the relation slice[i] < slice[j] by using the natural ordering of the stored types.
func (slice slice[T]) Less(i, j int) bool { return slice[i].Less(slice[j]) }

func Sort[T types.Comparable[T]](collection Collection[T]) {

	// sort.Sort[T](collection)
	// var slice slice[T] = collection.Collect() // linear time to collect all members into a slice O(n).
	// sort.Sort(slice)                          // log linear time to sort O(nlogn)
	// collection.Clear()                        // constant time O(1)
	// collection.AddSlice(slice)                // linear time O(n) resulting in overall time complexity O(nlogn)
}

// SortBy sorts a collection given the custom less function which defines the ralation a < b i.e if a < b then less should return true
// otherwise should return false.
func SortBy[T types.Equitable[T]](collection Collection[T], less func(a, b T) bool) {
	// var slice = collection.Collect()
	// sort.Slice(slice, func(i, j int) bool {
	// 	return less(slice[i], slice[j])
	// })
	// collection.Clear()
	// collection.AddSlice(slice)
}
