# collections
[![Build Status](https://app.travis-ci.com/phantom820/collections.svg?branch=main)](https://app.travis-ci.com/phantom820/collections) [![codecov](https://codecov.io/gh/phantom820/collections/branch/main/graph/badge.svg?token=TY4FD26RP0)](https://codecov.io/gh/phantom820/collections)

collections is a library aiming to bring collections (common data structures) into Go. 

- `Tree`
  - `Red Black Tree`
  An implementation  of a red black tree that stores a key and an associated value in its node. See the usage example below
  ```go
    import (
      "github.com/phantom820/collections/tree"
      "github.com/phantom820/collections/types"
    )

    t := tree.NewRedBlackTree[types.Integer, string]() // empty rbt that uses Integer as a key and string for associated value.
    t.Insert(1, "2")                                   // creates a node  (1,"2").
    t.Search(1)                                        // Searches for a node with the key 1.
    t.Get(1)                                           // retrieves a node with the key 1.
    t.Delete(1)                                        // removes node with key 1 from the tree
    ...
  ```

- `Map`
  - `HashMap`
  An implementation of a hashmap that uses a red black tree as the underlting container in its individual buckets. See usage example below.
  ```go

  import (
    _map "github.com/phantom820/collections/map"
    "github.com/phantom820/collections/types"
  )

	m := _map.NewHashMap[types.Integer, string]()
	m.Put(1, "A") // makes map insertion.
	m.Put(2, "B")
	m.Get(1) // retrieves key from map.

	// creates a new map from the map m by adding 22 to the key and appending "$" to value.
	tMap := m.Map(func(e _map.MapEntry[types.Integer, string]) _map.MapEntry[types.Integer, string] {
		n := _map.NewMapEntry(e.Key()+22, e.Value()+"$")
		return n
	})

	// creates a new map by filtering the map based on even number key values.
	fMap := tMap.Filter(func(e _map.MapEntry[types.Integer, string]) bool {
		return e.Key()%2 == 0
	})
  ```

- `Linked List`
  - `List`
  A doubly linked list that only stores elements of the same type (homogeeous list). See usage examples below.
  ```go
  
	l := list.NewList[types.Integer](1, 2, 3) // creates a list with element 1,2,3 .
	l.AddFront(23)                            // adds an element to the front of the list.
	l.Front()                                 // retrieves front element of list.
	l.AddBack(34)                             // adds an element to the back of the list an alias for Add.
	l.Back()                                  // retrives back element of list.
	l.Contains(23)                            // checks if list contains 23.

	// create a new list that has transformed elements using specified function
	tList = l.Map(func(e types.Integer) types.Integer {
		return e + 1
	})

	// creates a new filtered listered using specified function.
	fList = l.Filter(func(e types.Integer) bool {
		return e > 3
	})
  ```

  - `ForwardList`
  A singly linked list that only stores elements of the same type. See usage examples below.

- Stack 
  - `ListStack`
  A `ForwardList` based implementation of a stack. See usage examples below.
  - `SliceStack` 
  A slice based implementation of a stack. See usage examples belo.w
