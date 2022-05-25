package main

import (
	"fmt"

	"github.com/phantom820/collections/heaps/maxheap"
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

}
