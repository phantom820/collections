package main

import (
	"github.com/phantom820/collections/queues"
	"github.com/phantom820/collections/types"
)

const (
	size = 10000
)

func addFront(dequeue queues.Dequeue[types.Int]) {
	for i := 0; i < size; i++ {
		dequeue.AddFront(types.Int(i))
	}
}

func addBack(dequeue queues.Dequeue[types.Int]) {
	for i := 0; i < size; i++ {
		dequeue.Add(types.Int(i))
	}
}

func main() {
	// err := errors.New(4, 5, 2)
	// fmt.Println(err.Error())
	// a := listdequeue.New[types.Int](1, 2, 3, 4, 5, 6)
	// b := slicedequeue.New[types.Int](1, 2, 3, 4, 5, 6)

	// start := time.Now()
	// addFront(a)
	// end := time.Now()
	// duration := float32(end.Sub(start).Nanoseconds()) / size
	// fmt.Printf("list dequeue addFront : %v\n", duration)

	// start = time.Now()
	// addFront(b)
	// end = time.Now()
	// duration = float32(end.Sub(start).Nanoseconds()) / size
	// fmt.Printf("Slice dequeue addFront : %v\n", duration)

	// c := listdequeue.New[types.Int](1, 2, 3, 4, 5, 6)
	// d := slicedequeue.New[types.Int](1, 2, 3, 4, 5, 6)

	// start = time.Now()
	// addBack(c)
	// end = time.Now()
	// duration = float32(end.Sub(start).Nanoseconds()) / size
	// fmt.Printf("list dequeue addBack : %v\n", duration)

	// start = time.Now()
	// addBack(d)
	// end = time.Now()
	// duration = float32(end.Sub(start).Nanoseconds()) / size
	// fmt.Printf("Slice dequeue addBack : %v\n", duration)

}
