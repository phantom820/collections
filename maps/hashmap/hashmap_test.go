package hashmap

import (
	"strconv"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/types"
	"github.com/stretchr/testify/assert"
)

// TestPut also covers tests for PutAll, ContainsKey, ContainsValue,  Empty, Len, Keys, Clear.
func TestPut(t *testing.T) {

	m := New[types.String, string]()

	// Case 1 : An empty map.
	assert.Equal(t, true, m.Empty())
	assert.ElementsMatch(t, []types.Int{}, m.Keys())
	v := m.Put("1", "A")
	assert.Equal(t, false, m.Empty())
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, "", v)
	assert.Equal(t, []types.String{"1"}, m.Keys())
	assert.ElementsMatch(t, []string{"A"}, m.Values())

	// Case 2 : An already mapped key.
	v = m.Put("1", "B")
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, "A", v)

	// Case 3 : A non empty map with a new key.
	v = m.Put("2", "B")
	assert.Equal(t, 2, m.Len())
	assert.Equal(t, "", v)

	// Case 4 : Adding all values from another map.
	other := New[types.String, string]()
	other.Put("15", "O")
	other.Put("16", "P")
	other.Put("17", "Q")

	assert.Equal(t, false, m.ContainsKey("15"))
	assert.Equal(t, false, m.ContainsKey("O"))

	m.PutAll(other)
	assert.Equal(t, true, m.ContainsKey("15"))
	assert.Equal(t, true, m.ContainsValue("O", func(a, b string) bool { return a == b }))
	assert.Equal(t, true, m.ContainsKey("16"))
	assert.Equal(t, true, m.ContainsValue("P", func(a, b string) bool { return a == b }))
	assert.Equal(t, true, m.ContainsKey("17"))
	assert.Equal(t, false, m.ContainsValue("W", func(a, b string) bool { return a == b }))

	m.Clear()
	assert.Equal(t, true, m.Empty())

}

// TestPutIfAbsent
func TestPutIfAbsent(t *testing.T) {

	m := New[types.String, string]()

	// Case 1 : An empty map.
	assert.Equal(t, true, m.Empty())
	assert.Equal(t, false, m.ContainsKey("1"))
	assert.Equal(t, false, m.ContainsValue("A", func(a, b string) bool { return a == b }))
	b := m.PutIfAbsent("1", "A")
	assert.Equal(t, false, m.Empty())
	assert.Equal(t, true, m.ContainsKey("1"))
	assert.Equal(t, true, m.ContainsValue("A", func(a, b string) bool { return a == b }))
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, true, b)

	// Case 2 : An already mapped key.
	b = m.PutIfAbsent("1", "B")
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, false, b)

	// Case 3 : A non empty map with a new key.
	b = m.PutIfAbsent("2", "C")
	assert.Equal(t, 2, m.Len())
	assert.Equal(t, true, b)

	// Case 5 : Bucket non empty but does not have key.
	m.PutIfAbsent("\x10AAAA", "AA")

	assert.Equal(t, true, m.ContainsKey("\x10AAAA"))
	assert.Equal(t, false, m.ContainsKey("\x00AAAA"))

	m.PutIfAbsent("\x00AAAA", "B")
	assert.Equal(t, true, m.ContainsKey("\x00AAAA"))

}

// TestRemove also covers tests for Remove, RemoveAll
func TestRemove(t *testing.T) {

	m := New[types.Int, string]()

	// Case 1 : Key to remove absent.
	_, b := m.Remove(1)
	assert.Equal(t, false, b)

	// Case 2 : Keys to remove is mapped.
	m.Put(1, "A")
	v, _ := m.Remove(1)
	assert.Equal(t, "A", v)

	m.Put(1, "A")
	m.Put(2, "B")
	m.Put(3, "C")
	m.Put(4, "D")
	m.Put(5, "E")

	v, _ = m.Remove(2)
	assert.Equal(t, "B", v)

	// Case 3 : Remove a number of keys.
	l := list.New[types.Int](1, 3, 4, 5)
	m.RemoveAll(l)

	assert.Equal(t, true, m.Empty())

}

func TestResize(t *testing.T) {

	m := New[types.Int, int]()
	assert.Equal(t, 16, m.Capacity())
	assert.Equal(t, float32(0), m.LoadFactor())
	for i := 1; i <= 16; i++ {
		m.Put(types.Int(i), i)
	}
	assert.Equal(t, 32, m.Capacity())             // should have doubled in caoacity after crossing threshold.
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
	}
	assert.ElementsMatch(t, m.Keys(), keys)
	assert.ElementsMatch(t, m.Values(), values)

	// Case 2 : Next on exhausted iterator should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, iterator.NoNextElementError, r.(error))
			}
		}()

		it.Next()
	})

	// Case 3 : Cycling should reset iterator.
	it.Cycle()
	entry := it.Next()
	assert.Equal(t, keys[0], entry.Key)
	assert.Equal(t, values[0], entry.Value)

}

// TestEquals covers tests for Equals.
func TestEquals(t *testing.T) {

	m := New[types.Int, string]()
	other := New[types.Int, string]()

	// Case 1 : Self equivalence and empty maps are equal.
	assert.Equal(t, true, m.Equals(other, func(a, b string) bool { return a == b }))
	assert.Equal(t, true, m.Equals(m, func(a, b string) bool { return a == b }))

	// Case 2 : Different sizes.
	m.Put(1, "A")
	m.Put(2, "B")

	assert.Equal(t, false, m.Equals(other, func(a, b string) bool { return a == b }))

	// Case 3 : Same sizes different keys.
	other.Put(1, "A")
	other.Put(3, "B")
	assert.Equal(t, false, m.Equals(other, func(a, b string) bool { return a == b }))

	// Case 4 : Same sizes different values.
	other.Remove(3)
	other.Put(2, "C")
	assert.Equal(t, false, m.Equals(other, func(a, b string) bool { return a == b }))

	// Case 5 : Equal maps.
	other.Put(2, "B")
	assert.Equal(t, true, m.Equals(other, func(a, b string) bool { return a == b }))
}

func TestMap(t *testing.T) {

	m := New[types.Int, string]()
	m.Put(1, "A")
	m.Put(2, "B")
	m.Put(3, "C")

	other := m.Map(func(e maps.MapEntry[types.Int, string]) maps.MapEntry[types.Int, string] {
		return maps.MapEntry[types.Int, string]{Key: e.Key, Value: e.Value + "$"}
	})

	assert.Equal(t, true, m.Equals(other, func(a, b string) bool { return a+"$" == b }))

}

func TestFilter(t *testing.T) {

	m := New[types.Int, string]()
	m.Put(1, "A")
	m.Put(2, "B")
	m.Put(3, "C")

	other := m.Filter(func(e maps.MapEntry[types.Int, string]) bool { return e.Key%2 == 0 })

	n := New[types.Int, string]()
	n.Put(2, "B")

	assert.Equal(t, true, n.Equals(other, func(a, b string) bool { return a == b }))

}
