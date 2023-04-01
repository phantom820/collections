package dequeue_benchmarks

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/queues/listdequeue"
	"github.com/phantom820/collections/queues/vectordequeue"
)

// go test -bench=./... -benchmem -benchtime=5x github.com/phantom820/collections/benchmarks/queues
const (
	size = 1000000
)

var (
	data = generateData(size)
)

type constructor struct {
	new  func(element ...int) collections.Dequeue[int]
	name string
}

func generateData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(size)
	}
	return data
}

func BenchmarkAddFirst(b *testing.B) {

	constructors := []constructor{
		{
			name: "ListDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return listdequeue.New(elements...) },
		},
		{
			name: "VectorDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return vectordequeue.New(elements...) },
		},
	}

	for _, constructor := range constructors {

		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				queue := constructor.new()
				b.StartTimer()
				for j := 0; j < size; j++ {
					queue.AddFirst(data[j])
				}
			}
		})
	}

}

func BenchmarkPeekFirst(b *testing.B) {

	constructors := []constructor{
		{
			name: "ListDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return listdequeue.New(elements...) },
		},
		{
			name: "VectorDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return vectordequeue.New(elements...) },
		},
	}

	for _, constructor := range constructors {
		queue := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				queue.PeekFirst()
			}
		})
	}

}

func BenchmarkRemoveFirst(b *testing.B) {

	constructors := []constructor{
		{
			name: "ListDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return listdequeue.New(elements...) },
		},
		{
			name: "VectorDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return vectordequeue.New(elements...) },
		},
	}

	for _, constructor := range constructors {
		queue := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				queue.RemoveFirst()
			}
		})
	}

}

func BenchmarkAddLast(b *testing.B) {

	constructors := []constructor{
		{
			name: "ListDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return listdequeue.New(elements...) },
		},
		{
			name: "VectorDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return vectordequeue.New(elements...) },
		},
	}

	for _, constructor := range constructors {

		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				queue := constructor.new()
				b.StartTimer()
				for j := 0; j < size; j++ {
					queue.AddLast(data[j])
				}
			}
		})
	}

}

func BenchmarkPeekLast(b *testing.B) {

	constructors := []constructor{
		{
			name: "ListDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return listdequeue.New(elements...) },
		},
		{
			name: "VectorDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return vectordequeue.New(elements...) },
		},
	}

	for _, constructor := range constructors {
		queue := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				queue.PeekLast()
			}
		})
	}

}

func BenchmarkRemoveLast(b *testing.B) {

	constructors := []constructor{
		{
			name: "ListDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return listdequeue.New(elements...) },
		},
		{
			name: "VectorDequeue",
			new:  func(elements ...int) collections.Dequeue[int] { return vectordequeue.New(elements...) },
		},
	}

	for _, constructor := range constructors {
		queue := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				queue.RemoveLast()
			}
		})
	}

}
