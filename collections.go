// package collections defines the interfaces for common container data structures. Each container type is an iterable and an iterator can be obtained to access
// its elements. See below for some brief descriptions.
//
//  1. Maps[K comparable, V any] : A Map is an Iterable consisting of pairs of keys and values (also named mappings or associations).
//     1.1 HashMap[K, V] : This is a wrapper around a standard map[K]V i.e has map[K]V as its base type and can be ranged over.
//     1.2 LinkedHashMap[K, V] : This is similar to a HashMap[K, V] however elements are iterated over following their insertion order.
//     1.2 TreeMap[K, V] : A sorted map that stored elements in a sorted order, this backed by a Red Black Tree.
//
// 2.Collection[T comparable] : This is an interface satisfied by
//
//			2.1 List[T] : Linear ordered data structure that supports index based operations, this interface is satisfied by the following concrete types.
//			   a. Vector[T] : This is backed by a standard slice.
//			   b. ImmutableVector[T] : An immutable version of a [Vector[T]].
//			   c. ForwardList[T] : A singly linked list with a tail pointer.
//			   d. ImmutableForwardList[T] : An immutable version of a [ForwardList[T]].
//			   e. LinkedList[T] : A doubly linkedList.
//
//		    2.2 Queue[T] : Linear data structure that is used in for First in First Out Operations (FIFO).
//				 - Dequeue[T] : A double ended queue to support additions and removals on both ends.
//	          a. ListDequeue[T] : A linked list based implementation of a dequeue.
//			  b. VectorDequeue[T] : A slice based implementastion of a dequeue.
//
// Dequeue[T] : This is a double ended queue and can either be backed by a Vector[T] or a LinkedList[T].
//
//	 2.3 Set[T] : Non-linear data structure that stores unique elements and has quick lookups, this interface is satisfied by the following concrete types.
//		   a. HashSet[T] : A set implementation backed by a [HashMap] with no particular ordering for element iteration.
//		   b. LinkedHashSet[T] : A set implementation backed by a [LinkedHashMap] in which elements are iterated on following their insertion order.
//		   c.TreeSet[T] : A set implementation backed by a [TreeMap] in which elements are iterated on following particular ordering.
package collections

import (
	"reflect"

	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/types/optional"
	"github.com/phantom820/collections/types/pair"
)

// Map a key, value container that supports efficient lookups, insertions and deletions.
type Map[K comparable, V any] interface {
	iterable.Iterable[pair.Pair[K, V]]
	ContainsKey(k K) bool                            // Returns true if the map contains a mapping for the given key.
	ContainsValue(v V, f func(V, V) bool) bool       // Returns true if any key in the map is mapped to the value.
	ForEach(f func(K, V))                            // Perform the action f for each key, value pair in the map.
	Clear()                                          // Clears the contents of the map.
	Get(k K) optional.Optional[V]                    // Optionally returns the value associated with a key.
	GetIf(f func(K) bool) []V                        // Returns all values with keys that satisfy the given predicate.
	Put(k K, v V) optional.Optional[V]               // Adds a new key/value pair to this map and optionally returns previously bound value.
	PutIfAbsent(k K, v V) optional.Optional[V]       // Adds a new key/value pair to the map if the key is not already bounded and optionally returns bound value.
	Len() int                                        // Returns the size of the map.
	Remove(k K) optional.Optional[V]                 // Removes a key from the map, returning the value associated previously with that key as an option.
	RemoveIf(f func(K) bool) bool                    // Removes all the key, value mapping in which the key satisfies the given predicate.
	Keys() []K                                       // Returns a slice containing all the keys in the map.
	Values() []V                                     // Returns a slice containing all the values in the map.
	Empty() bool                                     // Returns true if the map has no elements.
	Equals(m Map[K, V], equals func(V, V) bool) bool // Returns true if the 2 maps are equal. Two maps are equal if thay have the same size and have the same key, value mappings.
}

// Collection a container for a grouping of elements.
type Collection[T comparable] interface {
	iterable.Iterable[T]
	Add(e T) bool                                 // Adds the given element to the collection and returns true if the element was added.
	AddAll(iterable iterable.Iterable[T]) bool    // Adds all of the elements in the specified iterable to the collection and returns true if the collection changed as a result of the operation.
	AddSlice(s []T) bool                          // Adds all of the elements in the specified slice to the collection and returns true if the collection changed as a result of the operation.
	Clear()                                       // Removes all of the elements from the collection.
	Contains(e T) bool                            // Returns true if this collection contains the specified element.
	Empty() bool                                  // Returns true if the collection contains no elements.
	Remove(e T) bool                              // Returns the first occurence of the given element and returns true if the collection changed as a result of the operation.
	RemoveIf(func(T) bool) bool                   // Removes all of the elements of the collection that satisfy the given predicate and returns true if the collection changed as a result of the operation.
	RemoveAll(iterable iterable.Iterable[T]) bool // Removes all of this collection's elements that are also contained in the specified iterable.
	RemoveSlice(s []T) bool                       // Removes all of this collection's elements that are also contained in the specified slice.
	RetainAll(c Collection[T]) bool               // Retains only the elements in this collection that are contained in the specified collection.
	ForEach(func(T))                              // Performs the given action for each element of the collection.
	Len() int                                     // Returns the number of elements in the collection.
	ToSlice() []T                                 // Returns a slice containing all of the elements of the collection.
}

// List a linear ordered data structure that supports index based operations.
type List[T comparable] interface {
	Collection[T]
	AddAt(i int, e T)         // Inserts the specified element at the specified index in the list.
	At(i int) T               // Returns the element at the specified index in the list
	Set(i int, e T) T         // Replaces the element at the specified index in the list with the specified element.
	RemoveAt(i int) T         // Removes the element at the specified index in the list.
	Equals(list List[T]) bool // Returns true if the list is equals to the given list. Two list are equal if they have the same size and have
	// the same elements in the same order.
	Sort(less func(a, b T) bool) // Sorts the list according to the ordering defined by the given less function for elements.
}

// Queue a linear data structure for processing elements in a First In First Out fashion.
type Queue[T comparable] interface {
	Collection[T]
	AddLast(e T) optional.Optional[T]  // Adds an element to the back of the queue and returns the previous back element as an option.
	PeekFirst() optional.Optional[T]   // Returns the front element of the queue as an option.
	RemoveFirst() optional.Optional[T] // Returns and removes the front element of the queue as an option.
}

// Dequeue a double ended [Queue] that also supports processing elements in a Last In First Out fashion.
type Dequeue[T comparable] interface {
	Queue[T]
	AddFirst(e T) optional.Optional[T] // Adds an element to the front of the dequeue and returns the previous front element as an option.
	PeekLast() optional.Optional[T]    // Returns the back element of the dequeue as an option.
	RemoveLast() optional.Optional[T]  // Returns and removes the back element of the dequeue as an option.
}

// Set a non-linear data structure that stores unique elements and supports quick lookups, insertions and deletions.
type Set[T comparable] interface {
	Collection[T]
	ContainsAll(iterable iterable.Iterable[T]) bool // Returns true if the set contains all of the elements in the specified iterable.
}

// IsNil returns true if the collection is nil.
func IsNil[T comparable](c Collection[T]) bool {
	return c == nil || reflect.ValueOf(c).IsNil()
}
