// Package interfaces specifies an interface for a collection. Any type that satisfies this interface is a collection.
package interfaces

import (
	"collections/iterator"
	"collections/types"
)

// Collection a blue print for which methods a collection must implement. A collection is a linear
// data structures such a linked list, queue, stack and a set.
type Collection[T types.Equitable[T]] interface {
	iterator.Iterable[T]              // Returns an iterator for iterating through the collection.
	Add(e T) bool                     // Adds element e to the collection.
	AddAll(c iterator.Iterable[T])    // Adds all elements from another collection into the collection.
	AddSlice(s []T)                   // Adds all elements from a slice into the collection.
	Len() int                         // Returns the size (number of items) stored in the collection.
	Contains(e T) bool                // Checks if the element e is a member of the collection.
	Remove(e T) bool                  // Tries to remove a specified element in the collection. It removes the first occurence of the element.
	RemoveAll(c iterator.Iterable[T]) // Removes all elements from another collections that appear in the collection.
	Empty() bool                      // Checks if the collection contains any elements.
	Clear()                           // Removes all elements in the collection.
}
