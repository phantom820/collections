package lists_benchmarks

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
)

// go test -bench=./... -benchmem -benchtime=5x github.com/phantom820/collections/benchmarks/lists
const (
	size = 1000000
	m    = 500000
)

var (
	data       = generateData(size)
	collection = vector.Of(generateData(m)...)
)

type constructor struct {
	new  func(element ...int) collections.List[int]
	name string
}

func generateData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(size)
	}
	return data
}

func BenchmarkAdd(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	for _, constructor := range constructors {

		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				constructor.new(data...)
			}
		})
	}

}

func BenchmarkAddAt(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	for _, constructor := range constructors {
		list := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				list.AddAt(rand.Intn(list.Len()), -1)
			}
		})
	}

}

func BenchmarkRemove(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	for _, constructor := range constructors {

		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				list := constructor.new(data...)
				b.StartTimer()
				list.Remove(rand.Intn(list.Len()))
			}
		})
	}

}

func BenchmarkRemoveAll(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	for _, constructor := range constructors {

		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				list := constructor.new(data...)
				b.StartTimer()
				list.RemoveAll(collection)
			}
		})
	}

}

func BenchmarkRetainAll(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	for _, constructor := range constructors {
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				list := constructor.new(data...)
				b.StartTimer()
				list.RetainAll(collection)
			}
		})
	}

}

func BenchmarkRemoveIf(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	predicate := func(i int) bool { return i%2 == 0 }
	for _, constructor := range constructors {
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				list := constructor.new(data...)
				b.StartTimer()
				list.RemoveIf(predicate)
			}
		})
	}
}

func BenchmarkRemoveAt(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	for _, constructor := range constructors {

		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				list := constructor.new(data...)
				b.StartTimer()
				list.RemoveAt(rand.Intn(list.Len()))
			}
		})
	}
}

func BenchmarkContains(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	for _, constructor := range constructors {
		list := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				list.Contains(rand.Intn(list.Len() + 1000))
			}
		})
	}

}

func BenchmarkAt(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	for _, constructor := range constructors {
		list := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				list.At(rand.Intn(list.Len()))
			}
		})
	}

}

func BenchmarkSet(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	for _, constructor := range constructors {
		list := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				list.Set(rand.Intn(list.Len()), -1)
			}
		})
	}

}

func BenchmarkSort(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	comparator := func(a, b int) bool { return a < b }
	for _, constructor := range constructors {
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				list := constructor.new(data...)
				b.StartTimer()
				list.Sort(comparator)
			}
		})
	}
}

func BenchmarkIterator(b *testing.B) {

	constructors := []constructor{
		{
			name: "Vector",
			new:  func(elements ...int) collections.List[int] { return vector.New(elements...) },
		},
		{
			name: "ForwardList",
			new:  func(elements ...int) collections.List[int] { return forwardlist.New(elements...) },
		},
		{
			name: "LinkedList",
			new:  func(elements ...int) collections.List[int] { return linkedlist.New(elements...) },
		},
	}

	for _, constructor := range constructors {
		list := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-size-%d", constructor.name, size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				it := list.Iterator()
				for it.HasNext() {
					it.Next()
				}
			}
		})
	}

}
