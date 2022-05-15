// Package maps provides an innterface that a map like data structure should satisfy and some utils for implementing a map such a map entry and iterator.
package maps

import (
	"github.com/phantom820/collections/iterator"
)

// MapIterator an interface that an iterator implementation for a map must satisfy.
type MapIterator[K any, V any] interface {
	HasNext() bool
	Next() MapEntry[K, V]
	Cycle()
}

// MapIterable provides an interface for specifying a map iterable.
type MapIterable[K any, V any] interface {
	Keys() []K
	Values() []V
	Iterator() MapIterator[K, V]
}

// MapEntry wraps around the key and value of a map.
type MapEntry[K any, V any] struct {
	Key   K
	Value V
}

// Map interface specifying functions an implementation of a map should provide.
type Map[K any, V any] interface {
	MapIterable[K, V]
	Put(key K, value V) V                                 // Inserts the key and value into the map. Returns the previous value associated with the key if it was present otherwise zero value.
	PutIfAbsent(key K, value V) bool                      // Insert the key and value in the map if the key does not already exist.
	PutAll(m Map[K, V])                                   // Inserts all entries from another map into the map.
	Get(key K) (V, bool)                                  // Retrieves the valuee associated with the key. Returns zero value if the key does not exist.
	Len() int                                             // Returns the size of the map.
	Keys() []K                                            // Returns the keys in the map as a slice.
	ContainsKey(key K) bool                               // Checks if the map contains the specified key.
	ContainsValue(value V, equals func(a, b V) bool) bool // Checks if the map has an entry whose value is the specified value. func equals is used to compare value for equality.
	Remove(key K) (V, bool)                               // Removes the map entry with the specified key.Will return the value that was associated with the key.
	RemoveAll(c iterator.Iterable[K])                     // Removes all keys in the map that appear in an iterable.
	Clear()                                               // Removes all entries in the map.
	Empty() bool                                          // Checks if the map is empty.
}
