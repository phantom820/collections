package main

import (
	"fmt"

	"github.com/phantom820/collections/lists"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
)

func main() {

	// m := hashmap.New[string, int]()
	// m.Put("A", 0)
	// m["ADD"] = 22

	// fmt.Println(m)

	a := vector.Of[int]()
	b := linkedlist.New[int]()
	b.Add(1)
	fmt.Println(b.ImmutableCopy())
	// b := forwadlist.Of(1, 2, 3, 4)
	// a.SubList(1, 3)

	// fmt.Println(a.SubList(1, 2))
	// fmt.Println(b.SubList(1, 2))

	c := lists.Map[int](&a, func(s int) int { return s * 2 })
	fmt.Println(c)
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
