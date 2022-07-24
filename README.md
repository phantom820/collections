# collections
[![Build Status](https://app.travis-ci.com/phantom820/collections.svg?branch=main)](https://app.travis-ci.com/phantom820/collections) [![codecov](https://codecov.io/gh/phantom820/collections/branch/main/graph/badge.svg?token=TY4FD26RP0)](https://codecov.io/gh/phantom820/collections)

collections is a library tht aims to bring common data structures into Go. These collections can be used with user defined types that satisfy an interface required by that collection i.e collections such as `List`, `Queue` and `Stack` require types to satisfy `Equitable` interface while a `Map` requires a type that satisfies the `Hashable` interface and so forth. See [types](https://github.com/phantom820/collections/blob/main/types/types.go), in which wrappers around primitives `string` and `int` have been implemented. For some performance insights see [collections-benchmarks](https://github.com/phantom820/collections-benchmarks)

### Install 
` go get github.com/phantom820/collections@v0.3.0-alpha.2.7`

### Collections
"Linear" data structures.
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
- #### Iterating
```go
it := collection.Iterator()
for it.HasNext() {
	fmt.Println(it.Next())
}
```
- #### Sorting
Only supported by Lists.
```go
collections.Sort[T](collection[T]) // Sorting using natural ordering of elements in the collection.
// with custom comparator.	
collections.SortBy[T](l,func(a, b T) bool { ..... }) // Sorting with custom comparator function. Needs to define ordering for 2 elemetns.

```

- #### Lists
 Linked lists and  vector implementations.
```go	
type List[T types.Equitable[T]] interface {
	collections.Collection[T]
	Front() T         // Returns the front element in the list. Will panic if there is no front element.
	RemoveFront() T   // Returns and removes the front element in the list.
	Back() T          // Returns the element at the back of the list. Will panic if no back element.
	RemoveBack() T    // Returns and removes the element at the back of the list. Will panic if no back element.
	Set(i int, e T) T // Replaces the element at the specified index with the new element and returns old element. Will panic if index out of bounds.
	Swap(i, j int)    // Swaps the element at index i with the element at index j. Will panic if one or both indices out of bounds.
	At(i int) T       // Retrieves the element at the specified index. Will panic if index is out of bounds.
	RemoveAt(i int) T // Removes the element ath the specified index andreturns it. Will panic if index out of bounds.
	AddAt(i int, e T) // Adds the element at the specified index. Will panic if index out of bounds.
}

```
- `ForwardList` : singly linked list.
- `List` : doubly linked list.
- `Vector` : a vector by wrapping around a slice.


- #### Queues
	Various queue implementations.
```go
type Queue[T types.Equitable[T]] interface {
	collections.Collection[T]
	Front() T       //  Returns the front element of the queue. Will panic if no front element.
	RemoveFront() T // Returns and removes the front element of the queue. Will panic if no front element.
}

type DeqQueue[T types.Equitable[T]] interface {
	collections.Collection[T]
	AddFront(elements ...T) // Adds elements to the front of the queue.
	Front() T               //  Returns the front element of the queue. Will panic if no front element.
	RemoveFront() T         // Returns and removes the front element of the queue. Will panic if no front element.
	Back() T                //  Returns the back element of the queue. Will panic if no back element.
	RemoveBack() T          //  Returns and removes the item at the back of the queue
}
```
- `ListQueue` : `ForwardList` based implementation of a queue.
- `SliceQueue` : slice based implementation of a queue.
- `ListDequeue` : slice based implementation of a dequeue.
- `PriorityQueue` : binary heap based implementation of a priority queue (min and max queue).
- `SliceDequeue` : slice based implementation of a dequeue. 



- Stack
Various stack implementaions. 
```go
type Stack[T types.Equitable[T]] interface {
	collections.Collection[T]
	Peek() T // Returns the top element in the stack. Will panic if no top element.
	Pop() T  // Returns and  removes the top element in the stack. Will panic if no top element.
}
```
- `ListStack` : a singly linked list based implementation of a stack.
- `SliceStack` : slice based implementation of a stack.

### Trees
Tree based data structures.
```go
type Tree[K any, V any] interface {
	Insert(key K, value V) bool      // Inserts a node with the specified key and value.
	Delete(key K) bool               // Deletes the node with specified key. Returns true if such a node was found and deleted otherwise false.
	Clear()                          // Deleted all the nodes in the tree.
	Search(key K) bool               // Searches for a node with the specified key.
	Update(key K, value V) (V, bool) // Updates the node with specified key with the new value. Returns the old value if there was such a node.
	Get(key K) (V, bool)             // Retrieves the value of the node with the specified key.
	InOrderTraversal() []K           // Performs an in order traversal and returns results in a slice.
	Values() []V                     // Retrieves all the values sin the tree.
	Keys() []K                       // Retrieves all the keys in the tree.
	Empty() bool                     // Chekcs if the tree is empty.
	Len() int                        // Returns the size of the tree.
}
```

- `Red Black Tree` : a red black tree implementation witho nodes that store a key and an associated value.

### Maps
Map based data structures.
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

- `HashMap` : a map implementation that uses a hash table (slice) and red black tree for individual containers in buckets.
- `LinkedHashMap` : a map implementation similar to HashMap but preserves insertion order of entries.
- `TreeMap` : a map implementation in which entries are stored according to the natural sorting order of their keys. 

### Sets
Set data structures.
```go
type Set[T types.Equitable[T]] interface {
	collections.Collection[T]
	RetainAll(collection collections.Collection[T]) bool // Removes all entries from the set that do not appear in the other collection. Return true if the set was modified.
	RemoveIf(f func(element T) bool) bool                // Removes elements from the set that satisfy the predicate function f
}
```
- `HashSet` : a set implementation based on a HashMap.
- `LinkedHashSet` : a set implementation in which insertion order of items is preserved, uses LinkedHashMap.
- `TreeSet` : as set implementation in which elements are stored according to their defined ordering, uses TreeMap.


