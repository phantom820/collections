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
