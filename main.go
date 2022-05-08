package main

import (
	"collections/queue"
	"collections/wrapper"
	"fmt"
)

func main() {
	// // Haha my own collections.
	q := queue.NewSliceQueue[wrapper.Integer]()
	q.Add(1)
	q.Add(2)
	q.Add(3)

	fmt.Println(q)
	c := q.Collect()
	fmt.Println(c)
	c = c[1:]
	fmt.Println(q)
	q.Remove(2)
	fmt.Println(q)

}
