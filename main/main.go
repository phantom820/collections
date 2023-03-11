package main

import (
	"time"

	"github.com/phantom820/collections"
)

type list[T comparable] interface {
	Add(e T) bool
	Iterator() collections.Iterator[T]
}

func benchmarkAdd(list list[int], m, n int) float64 {

	accumulator := 0
	for i := 0; i < m; i++ {
		start := time.Now()
		for j := 0; j < n; j++ {
			list.Add(j)
		}
		end := time.Now()
		accumulator = accumulator + int(end.Sub(start).Milliseconds())
	}
	return float64(accumulator) / float64(m)
}

func main() {

	// a := forwardlist.New[int]()
	// b := other.New[int]()

	// m := 10
	// n := 1000000
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("Default implementation add %v elements duration %v.\n", n, benchmarkAdd(a, m, n))
	// 	fmt.Printf("Type definition implementation add %v elements duration %v.\n", n, benchmarkAdd(b, m, n))
	// }

	// s := linkedhashmap.New[int, int]()
}
