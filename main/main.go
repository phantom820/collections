package main

import (
	"fmt"

	"github.com/phantom820/collections/heaps/maxheap"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/types"
)

func main() {

	heap := maxheap.New[types.Int]()

	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(9)
	heap.Insert(5)
	heap.Insert(2)

	fmt.Println(heap)
	heap.DeleteTop()
	fmt.Println(heap)
	heap.DeleteTop()
	fmt.Println(heap)

	l := forwardlist.New[types.Int](1, 2, 3)
	fmt.Println(l)
	l.Reverse()
	l.Add(23)
	fmt.Println(l)
	l.AddFront(34)
	fmt.Println(l)
	l.Reverse()
	fmt.Println(l)
	l.AddAt(2, 45)
	fmt.Println(l)

}
