package sets_benchmarks

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/vector"
	"github.com/phantom820/collections/sets/hashset"
	"github.com/phantom820/collections/sets/linkedhashset"
	"github.com/phantom820/collections/sets/treeset"
)

// go test -bench=./... -benchmem -benchtime=5x github.com/phantom820/collections/benchmarks/sets
const (
	size = 1000000
	m    = 500000
)

var (
	data       = generateSetData(size)
	collection = vector.Of(generateData(m)...)
	lessThan   = func(i1, i2 int) bool { return i1 < i2 }
)

type constructor struct {
	new  func(element ...int) collections.Set[int]
	name string
}

func generateSetData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	return data
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
			name: "HashSet",
			new:  func(elements ...int) collections.Set[int] { return hashset.New(elements...) },
		},
		{
			name: "LinkedHashSet",
			new:  func(elements ...int) collections.Set[int] { return linkedhashset.New(elements...) },
		},
		{
			name: "TreeSet",
			new:  func(elements ...int) collections.Set[int] { return treeset.New(lessThan, elements...) },
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

func BenchmarkRemove(b *testing.B) {

	constructors := []constructor{
		{
			name: "HashSet",
			new:  func(elements ...int) collections.Set[int] { return hashset.New(elements...) },
		},
		{
			name: "LinkedHashSet",
			new:  func(elements ...int) collections.Set[int] { return linkedhashset.New(elements...) },
		},
		{
			name: "TreeSet",
			new:  func(elements ...int) collections.Set[int] { return treeset.New(lessThan, elements...) },
		},
	}

	for _, constructor := range constructors {
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				set := constructor.new(data...)
				b.StartTimer()
				set.Remove(rand.Intn(set.Len()))
			}
		})
	}

}

func BenchmarkRemoveIf(b *testing.B) {

	constructors := []constructor{
		{
			name: "HashSet",
			new:  func(elements ...int) collections.Set[int] { return hashset.New(elements...) },
		},
		{
			name: "LinkedHashSet",
			new:  func(elements ...int) collections.Set[int] { return linkedhashset.New(elements...) },
		},
		{
			name: "TreeSet",
			new:  func(elements ...int) collections.Set[int] { return treeset.New(lessThan, elements...) },
		},
	}

	predicate := func(i int) bool { return i%2 == 0 }
	for _, constructor := range constructors {
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				set := constructor.new(data...)
				b.StartTimer()
				set.RemoveIf(predicate)
			}
		})
	}

}

func BenchmarkRemoveAll(b *testing.B) {

	constructors := []constructor{
		{
			name: "HashSet",
			new:  func(elements ...int) collections.Set[int] { return hashset.New(elements...) },
		},
		{
			name: "LinkedHashSet",
			new:  func(elements ...int) collections.Set[int] { return linkedhashset.New(elements...) },
		},
		{
			name: "TreeSet",
			new:  func(elements ...int) collections.Set[int] { return treeset.New(lessThan, elements...) },
		},
	}

	for _, constructor := range constructors {
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				set := constructor.new(data...)
				b.StartTimer()
				set.RemoveAll(collection)
			}
		})
	}

}

func BenchmarkRetainAll(b *testing.B) {

	constructors := []constructor{
		{
			name: "HashSet",
			new:  func(elements ...int) collections.Set[int] { return hashset.New(elements...) },
		},
		{
			name: "LinkedHashSet",
			new:  func(elements ...int) collections.Set[int] { return linkedhashset.New(elements...) },
		},
		{
			name: "TreeSet",
			new:  func(elements ...int) collections.Set[int] { return treeset.New(lessThan, elements...) },
		},
	}

	for _, constructor := range constructors {
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				set := constructor.new(data...)
				b.StartTimer()
				set.RetainAll(collection)
			}
		})
	}

}

func BenchmarkContains(b *testing.B) {

	constructors := []constructor{
		{
			name: "HashSet",
			new:  func(elements ...int) collections.Set[int] { return hashset.New(elements...) },
		},
		{
			name: "LinkedHashSet",
			new:  func(elements ...int) collections.Set[int] { return linkedhashset.New(elements...) },
		},
		{
			name: "TreeSet",
			new:  func(elements ...int) collections.Set[int] { return treeset.New(lessThan, elements...) },
		},
	}

	for _, constructor := range constructors {
		set := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				set.Contains(rand.Intn(set.Len() + 10))
			}
		})
	}

}

func BenchmarkContainsAll(b *testing.B) {

	constructors := []constructor{
		{
			name: "HashSet",
			new:  func(elements ...int) collections.Set[int] { return hashset.New(elements...) },
		},
		{
			name: "LinkedHashSet",
			new:  func(elements ...int) collections.Set[int] { return linkedhashset.New(elements...) },
		},
		{
			name: "TreeSet",
			new:  func(elements ...int) collections.Set[int] { return treeset.New(lessThan, elements...) },
		},
	}

	for _, constructor := range constructors {
		set := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				set.ContainsAll(collection)
			}
		})
	}

}

func BenchmarkIterator(b *testing.B) {

	constructors := []constructor{
		{
			name: "HashSet",
			new:  func(elements ...int) collections.Set[int] { return hashset.New(elements...) },
		},
		{
			name: "LinkedHashSet",
			new:  func(elements ...int) collections.Set[int] { return linkedhashset.New(elements...) },
		},
		{
			name: "TreeSet",
			new:  func(elements ...int) collections.Set[int] { return treeset.New(lessThan, elements...) },
		},
	}

	for _, constructor := range constructors {
		set := constructor.new(data...)
		b.Run(fmt.Sprintf("%v-input-count-%d", constructor.name, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				it := set.Iterator()
				for it.HasNext() {
					it.Next()
				}
			}
		})
	}

}
