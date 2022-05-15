package main

import (
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/maps/treemap"
	"github.com/phantom820/collections/types"
)

func main() {

	m := hashmap.New[types.Int, string]()
	m.Put(1, "A")

	n := treemap.New[types.Int, string]()
	n.Put(2, "B")

	n.PutAll(m)

	// var s1 sets.Set[types.Int] = hashset.New[types.Int]()
}
