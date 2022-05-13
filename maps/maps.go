// Package maps provides an innterface that a map like data structure should satisfy and some utils for implementing a map such a map entry and iterator.
package maps

import (
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types"
)

// MapIterator an interface that an iterator implementation for a map must satisfy.
type MapIterator[K types.Hashable[K], V any] interface {
	HasNext() bool
	Next() MapEntry[K, V]
	Cycle()
}

// MapIterable provides an interface for specifying a map iterable.
type MapIterable[K types.Hashable[K], V any] interface {
	Keys() []K
	Values() []V
	Iterator() MapIterator[K, V]
}

// MapEntry wraps around the key and value of a map.
type MapEntry[K types.Hashable[K], V any] struct {
	Key   K
	Value V
}

// Map interface specifying functions an implementation of a map should provide.
type Map[K types.Hashable[K], V any] interface {
	MapIterable[K, V]
	Put(k K, v V) V                   // Inserts the key and value into the map. Returns the previous value associated with the key if it was present otherwise zero value.
	PutIfAbsent(k K, v V) bool        // Insert the key and value in the map if the key does not already exist.
	PutAll(m Map[K, V])               // Inserts all entries from another map into the map.
	Get(k K) (V, bool)                // Retrieves the valuee associated with the key. Returns zero value if the key does not exist.
	Len() int                         // Returns the size of the map.
	Capacity() int                    // Returns the capacity of the map (number pf buckets).
	LoadFactor() float32              // Returns the load factor of the map.
	Keys() []K                        // Returns the keys in the map as a slice.
	Contains(k K) bool                // Checks if the map contains the specified key.
	Remove(k K) bool                  // Removes the map entry with the specified key.
	RemoveAll(c iterator.Iterable[K]) // Removes all keys in the map that appear in an iterable.
	Clear()                           // Removes all entries in the map.
	Empty() bool                      // Checks if the map is empty.
}