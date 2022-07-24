package main

import (
	"fmt"

	"github.com/phantom820/collections/queues/slicequeue"
	"github.com/phantom820/collections/types"
)

func main() {
	l := slicequeue.New[types.Int](1, 2, 3, 4, 5, 6)
	it := l.Iterator()
	l = nil
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
