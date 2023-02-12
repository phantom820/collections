package main

import (
	"fmt"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/sets/hashset"
)

func main() {

	m := hashmap.New[string, int]()
	m.Put("A", 0)
	m["ADD"] = 22

	fmt.Println(m)

	s := hashset.Of[string]("1", "2", "1")
	fmt.Println(s)

	var a collections.ImmutableCollection[int] = hashset.ImmutableOf[int](1, 2, 345)
	fmt.Println(a)
}
