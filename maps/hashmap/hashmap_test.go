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

// TestPut covers tests for Put, Contains.
func TestPut(t *testing.T) {

	m := NewHashMap[types.String, string]()

	assert.Equal(t, false, m.Contains("1"))
	m.Put("1", "20")
	assert.Equal(t, true, m.Contains("1"))
	assert.Equal(t, 1, m.Len())
	m.Put("0ab", "21")

	assert.Equal(t, true, m.Contains("0ab"))
	m.Put("1c", "21")
	assert.Equal(t, true, m.Contains("1c"))
	m.Put("1c", "90")
	v, _ := m.Get("1c")
	assert.Equal(t, "90", v)
	assert.Equal(t, 3, m.Len())

}

// TestPutIdAbsent covers tests for PutIfAbsent.
func TestPutIfAbsent(t *testing.T) {

	m := NewHashMap[types.String, string]()

	// Case 1 : A key that;s not there.
	assert.Equal(t, true, m.PutIfAbsent("222", "1"))
	m.Put("1", "20")

	// Case 2 : A key that is already present.
	p := m.PutIfAbsent("1", "21")
	assert.Equal(t, false, p)

	assert.Equal(t, true, m.PutIfAbsent("22", "23"))
	v, _ := m.Get("22")
	assert.Equal(t, "23", v)

}

// TestRemove covers tests for Remove.
func TestRemove(t *testing.T) {

	m := NewHashMap[types.String, string]()

	// Case 1 : Remove an absent key.
	assert.Equal(t, false, m.Remove("1"))
	m.Put("1", "20")
	assert.Equal(t, 1, m.Len())

	// Case 2 : Remove a present key.
	m.Remove("1")
	assert.Equal(t, true, m.Empty())
	_, b := m.Get("1")
	assert.Equal(t, false, b)

}

// TestRemoverAll covers tests for RemoveAll
func TestRemoveAll(t *testing.T) {

	m := NewHashMap[types.String, string]()

	m.Put("1", "20")
	m.Put("2", "20")
	m.Put("3", "20")

	l := list.New[types.String]()
	l.Add("2")
	l.Add("3")

	m.RemoveAll(l)
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, false, m.Contains("2"))
}

// TestResize covers tests for resize.
func TestResize(t *testing.T) {

	m := NewHashMap[types.Integer, int]()
	assert.Equal(t, 16, m.Capacity())
	assert.Equal(t, float32(0), m.LoadFactor())
	for i := 1; i <= 16; i++ {
		m.Put(types.Integer(i), i)
	}
	assert.Equal(t, 32, m.Capacity())             // should have doubled in caoacity after crossing threshold.
	assert.Equal(t, float32(0.5), m.LoadFactor()) // load factor should be half.

	for i := 17; i <= 34; i++ {
		m.Put(types.Integer(i), i)
	}

	assert.Equal(t, 64, m.Capacity()) // should have doubled in size once more.

}

// TestKeys covers tests for Keys
func TestKeys(t *testing.T) {

	m := NewHashMap[types.Integer, int]()

	// Keys should be collected correctly.
	for i := 1; i <= 6; i++ {
		m.Put(types.Integer(i), i)
	}
	keys := []types.Integer{1, 2, 3, 4, 5, 6}
	assert.ElementsMatch(t, keys, m.Keys())

}

// TestValues covers tests for values.
func TestValues(t *testing.T) {

	m := NewHashMap[types.Integer, int]()

	// Values should be collected .
	for i := 1; i <= 6; i++ {
		m.Put(types.Integer(i), i)
	}
	values := []int{1, 2, 3, 4, 5, 6}
	assert.ElementsMatch(t, values, m.Values())

}

// TestIterator covers tests for Iterator()
func TestIterator(t *testing.T) {

	m := NewHashMap[types.String, int]()

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

	m := NewHashMap[types.Integer, int]()
	for i := 0; i <= 5; i++ {
		m.Put(types.Integer(i), i)
	}

	// Case 1 : map equals its self.
	assert.Equal(t, true, m.Equals(m, func(a, b int) bool { return a == b }))

	other := NewHashMap[types.Integer, int]()

	// Case 2 : maps with different keys should not be equal.
	assert.Equal(t, false, m.Equals(other, func(a, b int) bool { return a == b }))

	// Case 3 : maps with same keys and values should be equal.
	other.PutAll(m)
	assert.Equal(t, true, m.Equals(other, func(a, b int) bool { return a == b }))

	m.Clear()
	other.Clear()

	// Case 4 : maps with same keys but different values should not be equal.
	m.Put(1, 2)
	other.Put(1, 4)

	assert.Equal(t, false, m.Equals(other, func(a, b int) bool { return a == b }))

}

// TestMap covers tests for Map
func TestMap(t *testing.T) {
	m := NewHashMap[types.Integer, int]()
	for i := 1; i <= 6; i++ {
		m.Put(types.Integer(i), i)
	}

	other := m.Map(func(e maps.MapEntry[types.Integer, int]) maps.MapEntry[types.Integer, int] {
		k := e.Key + 2
		v := e.Value + 10
		return maps.MapEntry[types.Integer, int]{Key: k, Value: v}
	})

	keys := []types.Integer{3, 4, 5, 6, 7, 8}
	values := []int{11, 12, 13, 14, 15, 16}
	assert.ElementsMatch(t, keys, other.Keys())
	assert.ElementsMatch(t, values, other.Values())

}

// TestFilter covers tests for Filter.
func TestFilter(t *testing.T) {
	m := NewHashMap[types.Integer, int]()
	for i := 1; i <= 6; i++ {
		m.Put(types.Integer(i), i)
	}

	other := m.Filter(func(e maps.MapEntry[types.Integer, int]) bool {
		return e.Key%2 == 0
	})

	keys := []types.Integer{2, 4, 6}
	assert.ElementsMatch(t, keys, other.Keys())

}
