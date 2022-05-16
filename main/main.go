package main

import (
	"fmt"

	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/maps/treemap"
	"github.com/phantom820/collections/sort"
	"github.com/phantom820/collections/types"
)

func main() {

	m := hashmap.New[types.Int, string]()
	m.Put(1, "A")

	n := treemap.New[types.Int, string]()
	n.Put(2, "B")

	n.PutAll(m)

	// var s1 sets.Set[types.Int] = hashset.New[types.Int]()
	// v := reflect.ValueOf(*m)
	// y := v.FieldByName("capacity")
	// fmt.Println(y)
	l := list.New[types.Int](5, 4, 3, 2, 1, 13, 15, 6778, 90)
	fmt.Println(l)
	sort.SortBy[types.Int](l, func(a, b types.Int) bool { return a >= b })
	// fmt.Println(l.Len())
	fmt.Println(l)
}
