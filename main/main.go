package main

import (
	"fmt"

	"github.com/phantom820/collections/queues/dequeue"
)

func main() {

	q := dequeue.NewListDequeue[int](1, 2, 3)
	fmt.Println(q.PeekLast())
}
