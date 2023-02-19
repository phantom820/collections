package main

import (
	"fmt"

	"github.com/phantom820/collections/sets"
	"github.com/phantom820/collections/sets/hashset"
)

func main() {

	// m := hashmap.New[string, int]()
	// m.Put("A", 0)
	// m["ADD"] = 22

	// fmt.Println(m)

	a := hashset.Of[string]("1", "2")
	b := hashset.Of[string]("3", "1", "4", "2")

	fmt.Println(a)
	fmt.Println(b)
	c := sets.Intersection[string](&a, &b)

	fmt.Println(c.Contains("1"))
	fmt.Println(c.Len())
	c.ForEach(func(s string) { fmt.Println(s) })
	// var a collections.ImmutableCollection[int] = hashset.ImmutableOf[int](1, 2, 345)
	// fmt.Println(a)

	// tree := rbt.New[int, int](func(i1, i2 int) bool { return i1 < i2 })

	// tree.Insert(1, 1)
	// tree.Insert(2, 2)
	// root := tree.Root()

	// trees.InOrderTraversal(root)
	// fmt.Println(tree.Root())
	// fmt.Println(root.Left().Left())
	// // fmt.Println(root.eft())
}
