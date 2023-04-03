# collections
[![Build Status](https://app.travis-ci.com/phantom820/collections.svg?branch=main)](https://app.travis-ci.com/phantom820/collections) [![codecov](https://codecov.io/gh/phantom820/collections/branch/main/graph/badge.svg?token=TY4FD26RP0)](https://codecov.io/gh/phantom820/collections)

collections is a library tht aims to bring common data structures into Go.

### Install 
` go get github.com/phantom820/collections@v0.3.0-alpha.2.8`

#### Interfaces 
```go

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

```

#### Examples
```go
list := linkedlist.New[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
it := list.Iterator()
for it.HasNext() {
	fmt.Printf("%v ", it.Next())
}
// 1 2 3 4 5 6 7 8 9 10


set := hashset.Of(1,2,3,4,5,6,7,8,9,10)
fmt.Println(set.Contains(2))
// true
``` 


#### Benchmarks
The performance of different collections can be examined by running the `run_benchmarks.sh` script. This will generate becnhmarks in the standard go format and save results to csv files. 

##### Lists
```
goos: linux
goarch: amd64
pkg: github.com/phantom820/collections/benchmarks/lists
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkAdd/Vector-input-count-1000000-12         	       5	   8957225 ns/op	41678534 B/op	      40 allocs/op
BenchmarkAdd/ForwardList-input-count-1000000-12    	       5	  38264903 ns/op	16000024 B/op	 1000001 allocs/op
BenchmarkAdd/LinkedList-input-count-1000000-12     	       5	  45452442 ns/op	24000024 B/op	 1000001 allocs/op
BenchmarkAddAt/Vector-input-count-1-12             	       5	   2204156 ns/op	 8003584 B/op	       1 allocs/op
BenchmarkAddAt/ForwardList-input-count-1-12        	       5	    795382 ns/op	      16 B/op	       1 allocs/op
BenchmarkAddAt/LinkedList-input-count-1-12         	       5	   1269604 ns/op	      24 B/op	       1 allocs/op
BenchmarkRemove/Vector-input-count-1-12            	       5	    740261 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemove/ForwardList-input-count-1-12       	       5	   1029362 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemove/LinkedList-input-count-1-12        	       5	   1673496 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveAll/Vector-input-count-1-12         	       5	 273469922 ns/op	12407144 B/op	    9527 allocs/op
BenchmarkRemoveAll/ForwardList-input-count-1-12    	       5	  82623416 ns/op	12430732 B/op	    9517 allocs/op
BenchmarkRemoveAll/LinkedList-input-count-1-12     	       5	  88656244 ns/op	12411545 B/op	    9532 allocs/op
BenchmarkRetainAll/Vector-input-count-1-12         	       5	 282168293 ns/op	12425340 B/op	    9500 allocs/op
BenchmarkRetainAll/ForwardList-input-count-1-12    	       5	  85509110 ns/op	12399974 B/op	    9464 allocs/op
BenchmarkRetainAll/LinkedList-input-count-1-12     	       5	  81559467 ns/op	12428032 B/op	    9585 allocs/op
BenchmarkRemoveIf/Vector-input-count-1-12          	       5	  30316786 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveIf/ForwardList-input-count-1-12     	       5	  13718785 ns/op	      38 B/op	       0 allocs/op
BenchmarkRemoveIf/LinkedList-input-count-1-12      	       5	  11166855 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveAt/Vector-input-count-1-12          	       5	    502123 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveAt/ForwardList-input-count-1-12     	       5	    848262 ns/op	      19 B/op	       0 allocs/op
BenchmarkRemoveAt/LinkedList-input-count-1-12      	       5	   1075833 ns/op	       0 B/op	       0 allocs/op
BenchmarkContains/Vector-input-count-1-12          	       5	    347923 ns/op	       0 B/op	       0 allocs/op
BenchmarkContains/ForwardList-input-count-1-12     	       5	   1132103 ns/op	       0 B/op	       0 allocs/op
BenchmarkContains/LinkedList-input-count-1-12      	       5	   1261600 ns/op	       0 B/op	       0 allocs/op
BenchmarkAt/Vector-input-count-1-12                	       5	       191.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAt/ForwardList-input-count-1-12           	       5	    788561 ns/op	       0 B/op	       0 allocs/op
BenchmarkAt/LinkedList-input-count-1-12            	       5	   1413185 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet/Vector-input-count-1-12               	       5	       206.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet/ForwardList-input-count-1-12          	       5	   1132151 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet/LinkedList-input-count-1-12           	       5	   1573239 ns/op	       0 B/op	       0 allocs/op
BenchmarkSort/Vector-input-count-1-12              	       5	 167904304 ns/op	      56 B/op	       2 allocs/op
BenchmarkSort/ForwardList-input-count-1-12         	       5	 283105762 ns/op	16000022 B/op	  999999 allocs/op
BenchmarkSort/LinkedList-input-count-1-12          	       5	 352588642 ns/op	23999976 B/op	  999999 allocs/op
BenchmarkIterator/Vector-input-size-1000000-12     	       5	   8236392 ns/op	      72 B/op	       2 allocs/op
BenchmarkIterator/ForwardList-input-size-1000000-12         	       5	   8861225 ns/op	      72 B/op	       2 allocs/op
BenchmarkIterator/LinkedList-input-size-1000000-12          	       5	   8830911 ns/op	      72 B/op	       2 allocs/op
PASS
ok  	github.com/phantom820/collections/benchmarks/lists	16.089s
```


##### Queues
```
goos: linux
goarch: amd64
pkg: github.com/phantom820/collections/benchmarks/queues
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkAddFirst/ListDequeue-input-count-1000000-12         	       5	  78188654 ns/op	32000385 B/op	 1999999 allocs/op
BenchmarkAddFirst/VectorDequeue-input-count-1000000-12       	       5	  21490675 ns/op	24777131 B/op	 1000017 allocs/op
BenchmarkPeekFirst/ListDequeue-input-count-1-12              	       5	       560.2 ns/op	       9 B/op	       1 allocs/op
BenchmarkPeekFirst/VectorDequeue-input-count-1-12            	       5	       279.2 ns/op	       9 B/op	       1 allocs/op
BenchmarkRemoveFirst/ListDequeue-input-count-1-12            	       5	       693.6 ns/op	       9 B/op	       1 allocs/op
BenchmarkRemoveFirst/VectorDequeue-input-count-1-12          	       5	       281.8 ns/op	       9 B/op	       1 allocs/op
BenchmarkAddLast/ListDequeue-input-count-1000000-12          	       5	  82677500 ns/op	32000011 B/op	 1999999 allocs/op
BenchmarkAddLast/VectorDequeue-input-count-1000000-12        	       5	  30036007 ns/op	49678396 B/op	 1000034 allocs/op
BenchmarkPeekLast/ListDequeue-input-count-1-12               	       5	       572.0 ns/op	       9 B/op	       1 allocs/op
BenchmarkPeekLast/VectorDequeue-input-count-1-12             	       5	       229.0 ns/op	       9 B/op	       1 allocs/op
BenchmarkRemoveLast/ListDequeue-input-count-1-12             	       5	       570.2 ns/op	       9 B/op	       1 allocs/op
BenchmarkRemoveLast/VectorDequeue-input-count-1-12           	       5	       255.0 ns/op	       9 B/op	       1 allocs/op
PASS
ok  	github.com/phantom820/collections/benchmarks/queues	1.884s
```

#### Sets 
```
goos: linux
goarch: amd64
pkg: github.com/phantom820/collections/benchmarks/sets
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkAdd/HashSet-input-count-1000000-12         	       5	 135528384 ns/op	49513742 B/op	   38376 allocs/op
BenchmarkAdd/LinkedHashSet-input-count-1000000-12   	       5	 261277944 ns/op	109745966 B/op	 1038254 allocs/op
BenchmarkAdd/TreeSet-input-count-1000000-12         	       5	 663497244 ns/op	48000124 B/op	 1000004 allocs/op
BenchmarkRemove/HashSet-input-count-1-12            	       5	      2158 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemove/LinkedHashSet-input-count-1-12      	       5	      2075 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemove/TreeSet-input-count-1-12            	       5	      3109 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveIf/HashSet-input-count-1-12          	       5	  34354400 ns/op	21083392 B/op	      35 allocs/op
BenchmarkRemoveIf/LinkedHashSet-input-count-1-12    	       5	  69920193 ns/op	21083392 B/op	      35 allocs/op
BenchmarkRemoveIf/TreeSet-input-count-1-12          	       5	 181857277 ns/op	29086976 B/op	      36 allocs/op
BenchmarkRemoveAll/HashSet-input-count-1-12         	       5	  55787289 ns/op	     120 B/op	       4 allocs/op
BenchmarkRemoveAll/LinkedHashSet-input-count-1-12   	       5	  99068615 ns/op	     120 B/op	       4 allocs/op
BenchmarkRemoveAll/TreeSet-input-count-1-12         	       5	 347484894 ns/op	     120 B/op	       4 allocs/op
BenchmarkRetainAll/HashSet-input-count-1-12         	       5	 115644669 ns/op	45646355 B/op	    9550 allocs/op
BenchmarkRetainAll/LinkedHashSet-input-count-1-12   	       5	 169196601 ns/op	45651430 B/op	    9539 allocs/op
BenchmarkRetainAll/TreeSet-input-count-1-12         	       5	 246793469 ns/op	53673148 B/op	    9613 allocs/op
BenchmarkContains/HashSet-input-count-1-12          	       5	       261.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkContains/LinkedHashSet-input-count-1-12    	       5	       880.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkContains/TreeSet-input-count-1-12          	       5	      1572 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsAll/HashSet-input-count-1-12       	       5	  42178541 ns/op	     120 B/op	       4 allocs/op
BenchmarkContainsAll/LinkedHashSet-input-count-1-12 	       5	  46578857 ns/op	     120 B/op	       4 allocs/op
BenchmarkContainsAll/TreeSet-input-count-1-12       	       5	 296351125 ns/op	     120 B/op	       4 allocs/op
BenchmarkIterator/HashSet-input-count-1-12          	       5	  25824776 ns/op	16007232 B/op	       3 allocs/op
BenchmarkIterator/LinkedHashSet-input-count-1-12    	       5	   8345736 ns/op	      88 B/op	       3 allocs/op
BenchmarkIterator/TreeSet-input-count-1-12          	       5	  59680328 ns/op	16007256 B/op	       4 allocs/op
PASS
ok  	github.com/phantom820/collections/benchmarks/sets	46.970s
```

