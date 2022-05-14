package main

import (
	"fmt"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/types"
)

func main() {

	l1 := forwardlist.New[types.Int](5, 3, 6, 7, 20)
	l2 := forwardlist.New[types.Int](5, 3, 6, 7, 20)

	// it := l1.Iterator()
	// for it.HasNext() {
	// 	fmt.Println(it.Next())
	// }

	collections.Sort[types.Int](l1)                                                  // sorting using type defined ordering.
	collections.SortBy[types.Int](l2, func(a, b types.Int) bool { return !(a < b) }) // sorting with custom comparator for 2 elements.
	fmt.Println(l1)
	fmt.Println(l2)

	m := hashmap.New[types.Int, string]()
	m.Put(1, "A")
	m.Put(2, "B")

	it := m.Iterator()
	for it.HasNext() {
		entry := it.Next()
		fmt.Printf("Key:%d Value:%s\n", entry.Key, entry.Value)
	}

}
