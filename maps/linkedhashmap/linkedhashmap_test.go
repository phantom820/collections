package linkedhashmap

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/types"
	"github.com/stretchr/testify/assert"
)

// TestPut also covers tests for PutAll, Empty, Len, Keys, Values, Clear.
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

	// Case 2 : An already mapped key.
	v = m.Put("1", "B")
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, "A", v)

	// Case 3 : A non empty map with a new key.
	v = m.Put("2", "B")
	assert.Equal(t, 2, m.Len())
	assert.Equal(t, "", v)
	assert.Equal(t, []types.String{"1", "2"}, m.Keys())

	// Case 4 : Ordering of keys and values.
	for i := 3; i < 15; i++ {
		m.Put(types.String(fmt.Sprint(i)), string(rune(64+i)))
	}

	keys := []types.String{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14"}
	values := []string{"B", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}

	assert.Equal(t, keys, m.Keys())
	assert.Equal(t, values, m.Values())

	// Case 5 : Adding all values from another map.
	other := New[types.String, string]()
	other.Put("15", "O")
	other.Put("16", "P")
	other.Put("17", "Q")

	m.PutAll(other)
	assert.ElementsMatch(t, append(keys, "15", "16", "17"), m.Keys())
	assert.ElementsMatch(t, append(values, "O", "P", "Q"), m.Values())

	m.Clear()
	assert.Equal(t, true, m.Empty())

}

// TestPutIfAbsent also covers tests for ContainsKey, ContainsValue.
func TestPutIfAbsent(t *testing.T) {

	m := New[types.Int, string]()

	// Case 1 : An empty map.
	assert.Equal(t, true, m.Empty())
	assert.Equal(t, false, m.ContainsKey(1))
	assert.Equal(t, false, m.ContainsValue("A", func(a, b string) bool { return a == b }))
	b := m.PutIfAbsent(1, "A")
	assert.Equal(t, false, m.Empty())
	assert.Equal(t, true, m.ContainsKey(1))
	assert.Equal(t, true, m.ContainsValue("A", func(a, b string) bool { return a == b }))
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, true, b)

	// Case 2 : An already mapped key.
	b = m.PutIfAbsent(1, "B")
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, false, b)

	// Case 3 : A non empty map with a new key.
	b = m.PutIfAbsent(2, "C")
	assert.Equal(t, 2, m.Len())
	assert.Equal(t, true, b)

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
	assert.ElementsMatch(t, []types.Int{1, 3, 4, 5}, m.Keys())
	assert.ElementsMatch(t, []string{"A", "C", "D", "E"}, m.Values())

	// Case 3 : Remove a number of keys.
	l := list.New[types.Int](1, 3, 4, 5)
	m.RemoveAll(l)

	assert.Equal(t, true, m.Empty())

}

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

func TestIterator(t *testing.T) {

	m := New[types.Int, string]()

	// Case 1 : Legal iterator.
	m.Put(1, "A")
	m.Put(2, "B")
	m.Put(3, "C")

	it := m.Iterator()
	keys := []types.Int{}
	for it.HasNext() {
		keys = append(keys, it.Next().Key)
	}

	assert.ElementsMatch(t, []types.Int{1, 2, 3}, keys)

	// Case 2 : Exhausted iterator.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, iterator.NoNextElementError, r.(error))
			}
		}()
		it.Next()
	})

	// Case 3 : Recylce iterator.
	it.Cycle()

	v := it.Next().Value
	assert.Equal(t, "A", v)

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
