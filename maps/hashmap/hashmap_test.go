package hashmap

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/types"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	m := New[types.String, string]()

	assert.Equal(t, true, m.Empty())
	assert.Equal(t, 16, m.Capacity())
	assert.Equal(t, float32(0), m.LoadFactor())
	assert.Equal(t, float32(0.75), m.loadFactorLimit)

}

func TestPut(t *testing.T) {

	m := New[types.String, string]()

	// Case 1 : Put with a new key.
	assert.Equal(t, 0, m.Len())
	assert.Equal(t, false, m.ContainsKey("A"))
	m.Put("A", "A")
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, true, m.ContainsKey("A"))
	assert.Equal(t, true, m.ContainsValue("A", func(a, b string) bool { return a == b }))

	// Case 2 : Put with an existing key.
	value := m.Put("A", "B")
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, "A", value)
	assert.Equal(t, true, m.ContainsKey("A"))
	assert.Equal(t, false, m.ContainsValue("A", func(a, b string) bool { return a == b }))
	assert.Equal(t, true, m.ContainsValue("B", func(a, b string) bool { return a == b }))

	// Case 3 : Put with a key that will map to an empty bucket.
	m.Put("B", "B")
	assert.Equal(t, 2, m.Len())
	assert.Equal(t, true, m.ContainsKey("B"))

	// Case 4 : PutAll for keys comming from another map.
	otherMap := New[types.String, string]()
	otherMap.Put("A", "D")
	otherMap.Put("C", "C")
	otherMap.Put("D", "D")
	m.PutAll(otherMap)

	assert.Equal(t, 4, m.Len())
	value, _ = m.Get("A") // Value should have been updated with one in other map.
	assert.Equal(t, "D", value)

}

func TestPutIfAbsent(t *testing.T) {

	m := New[types.String, string]()

	// Case 1 : PutIfAbsent with a new key.
	assert.Equal(t, true, m.PutIfAbsent("A", "A"))

	// Case 2 : PutIfAbsent with an already mapped key.
	assert.Equal(t, false, m.PutIfAbsent("A", "B"))

	// Case 3 : PutIfAbsent with a key mapping to non empty bucket.
	assert.Equal(t, true, m.PutIfAbsent("\x10AAAA", "\x10AAAA"))
	assert.Equal(t, true, m.PutIfAbsent("\x00AAAA", "\x00AAAA")) // maps to non empty bucket.

}

func TestRemove(t *testing.T) {

	m := New[types.Int, string]()

	// Case 1 : Remove an absent key.
	_, ok := m.Remove(1)
	assert.Equal(t, false, ok)

	// Case 2 : Remove a mapped key
	m.Put(1, "A")
	value, ok := m.Remove(1)
	assert.Equal(t, true, ok)
	assert.Equal(t, "A", value)

	// Case 3 : Remove a number of keys.
	m.Put(1, "A")
	m.Put(2, "B")
	m.Put(3, "C")
	m.Put(4, "D")
	m.Put(5, "E")

	l := list.New[types.Int](1, 3, 4, 5)
	m.RemoveAll(l)

	assert.Equal(t, 1, m.Len())
	assert.Equal(t, false, m.ContainsKey(1))
	assert.Equal(t, false, m.ContainsKey(3))
	assert.Equal(t, false, m.ContainsKey(4))
	assert.Equal(t, false, m.ContainsKey(5))

}

func TestResize(t *testing.T) {

	m := New[types.Int, int]()

	// Case default: Properties should be initial values.
	assert.Equal(t, 16, m.Capacity())
	assert.Equal(t, float32(0), m.LoadFactor())

	// Case 1 : Put keys till we reach resize.
	for i := 1; i <= 16; i++ {
		m.Put(types.Int(i), i)
	}

	assert.Equal(t, 32, m.Capacity())             // should have doubled in capacity after crossing threshold.
	assert.Equal(t, float32(0.5), m.LoadFactor()) // load factor should be half.

	for i := 17; i <= 34; i++ {
		m.Put(types.Int(i), i)
	}

	assert.Equal(t, 64, m.Capacity()) // should have doubled in size once more.

}

func TestIterator(t *testing.T) {

	m := New[types.String, int]()

	// Case 1 : Iterator on map with elements.
	for i := 1; i <= 20; i++ {
		m.Put(types.String(strconv.Itoa(i)), i)
	}

	keys := make([]types.String, 0)
	values := make([]int, 0)

	it := m.Iterator()
	for it.HasNext() {
		entry := it.Next()
		keys = append(keys, entry.Key)
		values = append(values, entry.Value)
		fmt.Println(entry.Key)
	}

	assert.ElementsMatch(t, m.Keys(), keys)
	assert.ElementsMatch(t, m.Values(), values)

	// Case 2 : Next on exhausted iterator should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoNextElement, r.(errors.Error).Code())
			}
		}()
		it.Next()
	})

}

func TestIteratorConcurrentModification(t *testing.T) {

	m := New[types.String, int]()
	for i := 1; i <= 20; i++ {
		m.Put(types.String(strconv.Itoa(i)), i)
	}

	// Recovery for concurrent modifications.
	recovery := func() {
		if r := recover(); r != nil {
			assert.Equal(t, errors.ConcurrentModification, r.(*errors.Error).Code())
		}
	}
	// Case 1 : Put.
	it := m.Iterator()
	t.Run("Put while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			m.Put(types.String("D"), 22)
			it.Next()
		}
	})
	// Case 2 : PutIfAbsent.
	it = m.Iterator()
	t.Run("PutIfAbsent while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			m.PutIfAbsent(types.String("D"), 22)
			it.Next()
		}
	})
	// Case 3 : Remove.
	it = m.Iterator()
	t.Run("Remove while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			m.Remove(types.String("D"))
			it.Next()
		}
	})
	// Case 4 : Clear.
	it = m.Iterator()
	t.Run("Clear while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			m.Clear()
			it.Next()
		}
	})

}

func TestEquals(t *testing.T) {

	m := New[types.Int, string]()
	other := New[types.Int, string]()

	// Case 1 : Self equivalence and empty maps are equal.
	assert.Equal(t, true, m.Equals(other, func(a, b string) bool { return a == b }))
	assert.Equal(t, true, m.Equals(m, func(a, b string) bool { return a == b }))

	// Case 2 : Equals with different sizes should fail.
	m.Put(1, "A")
	m.Put(2, "B")

	assert.Equal(t, false, m.Equals(other, func(a, b string) bool { return a == b }))

	// Case 3 : Same sizes with different keys should fail.
	other.Put(1, "A")
	other.Put(3, "B")
	assert.Equal(t, false, m.Equals(other, func(a, b string) bool { return a == b }))

	// Case 4 : Same sizes with different values should fail.
	other.Remove(3)
	other.Put(2, "C")
	assert.Equal(t, false, m.Equals(other, func(a, b string) bool { return a == b }))

	// Case 5 : Same sizes and entries should pass.
	other.Put(2, "B")
	assert.Equal(t, true, m.Equals(other, func(a, b string) bool { return a == b }))

}

func TestClear(t *testing.T) {

	m := New[types.Int, int]()

	for i := 0; i < 20; i++ {
		m.Put(types.Int(i), i)
	}

	assert.Equal(t, 20, m.Len())
	m.Clear()
	assert.Equal(t, true, m.Empty())
	assert.Equal(t, 16, m.Capacity())
	assert.Equal(t, float32(0), m.LoadFactor())

}

func TestMap(t *testing.T) {

	m := New[types.Int, string]()

	// Case 1 : Mapping an empty map should give an empty map.
	other := m.Map(func(entry maps.MapEntry[types.Int, string]) maps.MapEntry[types.Int, string] {
		return maps.MapEntry[types.Int, string]{Key: entry.Key, Value: entry.Value}
	})
	assert.Equal(t, true, other.Empty())

	// Case 2 : Mapping a map with entries.
	m.Put(1, "A")
	m.Put(2, "B")
	m.Put(3, "C")

	other = m.Map(func(entry maps.MapEntry[types.Int, string]) maps.MapEntry[types.Int, string] {
		return maps.MapEntry[types.Int, string]{Key: entry.Key + 1, Value: entry.Value + "$"}
	})

	assert.Equal(t, 3, other.Len())
	value, _ := other.Get(2)
	assert.Equal(t, "A$", value)
	value, _ = other.Get(3)
	assert.Equal(t, "B$", value)
	value, _ = other.Get(4)
	assert.Equal(t, "C$", value)

}

func TestFilter(t *testing.T) {

	m := New[types.Int, string]()

	// Case 1 : Filtering an empty map should give an empty map.
	other := m.Filter(func(entry maps.MapEntry[types.Int, string]) bool { return entry.Key%2 == 0 })
	assert.Equal(t, true, other.Empty())

	// Case 2 : Filtering a map with entries.
	m.Put(1, "A")
	m.Put(2, "B")
	m.Put(3, "C")

	other = m.Filter(func(entry maps.MapEntry[types.Int, string]) bool { return entry.Key%2 == 0 })

	assert.Equal(t, 1, other.Len())
	assert.Equal(t, false, other.ContainsKey(1))
	assert.Equal(t, false, other.ContainsKey(3))
	value, _ := other.Get(2)
	assert.Equal(t, "B", value)

}
