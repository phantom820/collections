package main

import (
	"fmt"

	"github.com/phantom820/collections/lists"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
	"github.com/phantom820/collections/queues/dequeue"
)

type Matrix[T any] [][]T

func new[T any](m, n int) Matrix[T] {
	data := make([][]T, m)
	for i := 0; i < m; i++ {
		data[i] = make([]T, n)
	}
	return data
}

type Node[T any] struct {
	next  *Node[T]
	value T
}

// type LinkedList[T any] *Node[T]

// func (list LinkedList[T]) Add(e T) bool {
// 	list.next = *Node[T]{e}
// 	return true
// }

func main() {

	// m := hashmap.New[string, int]()
	// m.Put("A", 0)
	// m["ADD"] = 22

	// fmt.Println(m)

	a := vector.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b := linkedlist.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	sum, _ := lists.Reduce[int](&a, func(x, y int) int { return x + y })
	fmt.Println(sum)
	sum, _ = lists.Reduce[int](&b, func(x, y int) int { return x + y })
	fmt.Println(sum)

	partitions := lists.Partition[int](&a, 3)
	fmt.Println(partitions)

	f := func() {}
	f()

	f()

	c := new[int](2, 2)
	c[0][0] = 1
	fmt.Println(c)

	d := dequeue.NewListDequeue[int]()
	fmt.Println(d.PeekFirst())
}
