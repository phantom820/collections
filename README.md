# collections
[![Build Status](https://app.travis-ci.com/phantom820/collections.svg?branch=main)](https://app.travis-ci.com/phantom820/collections) [![codecov](https://codecov.io/gh/phantom820/collections/branch/main/graph/badge.svg?token=TY4FD26RP0)](https://codecov.io/gh/phantom820/collections)

collections is a library aiming to bring common data structures into Go. These collections can be used with user define types that satisfy an interface required by that collection i.e collections such as `List`, `Queue` and `Stack` require types to satisfy `Equitable` interface while a `Map` requires a type that satisfies the `Hashable` interface and so forth. See [types](https://github.com/phantom820/collections/blob/main/types/types.go), in which wrappers around primitives `string` and `int` have been implemented. 

### Data structures
- List
	- `ForwardList` : singly linked list.
	- `List` : doubly linked list.
-Vector
	- `Vector` : a vector (wrapper around Go slice)
- Queue
	- `ListQueue` : `ForwardList` based implementation of a queue.
	- `SliceQueue` : slice based implementation of a queue.
- Stack 
	- `ListStack` : `ForwardList` based implementation of a stack.
	- `SliceStack` : slice based implementation of a stack.
- Tree
	- `Red Black Tree` : a red black tree implementation witho nodes that store a key and an associated value.
- Map
	- `HashMap` : a map that uses a hash table (slice) and red black tree for individual containers in buckets.
- Set
	- `HashSet` : a set implementation based on a `HashMap`.

### Interfaces
-  `List, Vector, Queue , Stack ,  HashSet`
```go
// Satisfied by List, Vector, Queue , Stack ,  HashSet.
type Collection[T types.Equitable[T]] interface {
	iterator.Iterable[T]              // Returns an iterator for iterating through the collection.
	Add(elements ...T) bool           // Adds elements to the collection.
	AddAll(c iterator.Iterable[T])    // Adds all elements from another collection into the collection.
	AddSlice(s []T)                   // Adds all elements from a slice into the collection.
	Len() int                         // Returns the size (number of items) stored in the collection.
	Contains(e T) bool                // Checks if the element e is a member of the collection.
	Remove(e T) bool                  // Tries to remove a specified element in the collection. It removes the first occurence of the element.
	RemoveAll(c iterator.Iterable[T]) // Removes all elements from another collections that appear in the collection.
	Empty() bool                      // Checks if the collection contains any elements.
	Clear()                           // Removes all elements in the collection.
}
```

- Satisfied by `HashMap` .
```go
type Map[K types.Hashable[K], V any] interface {
	MapIterable[K, V]
	Put(key K, value V) V             // Inserts the key and value into the map. Returns the previous value associated with the key if it was present otherwise zero value.
	PutIfAbsent(key K, value V) bool  // Insert the key and value in the map if the key does not already exist.
	PutAll(m Map[K, V])               // Inserts all entries from another map into the map.
	Get(key K) (V, bool)              // Retrieves the valuee associated with the key. Returns zero value if the key does not exist.
	Len() int                         // Returns the size of the map.
	Keys() []K                        // Returns the keys in the map as a slice.
	Contains(key K) bool              // Checks if the map contains the specified key.
	Remove(key K) bool                // Removes the map entry with the specified key.
	RemoveAll(c iterator.Iterable[K]) // Removes all keys in the map that appear in an iterable.
	Clear()                           // Removes all entries in the map.
	Empty() bool                      // Checks if the map is empty.
}
```

### install 
` go get github.com/phantom820/collections@v0.3.0-alpha`

