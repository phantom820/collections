package hashmap

// import (
// 	"collections/list"
// 	_map "collections/map"
// 	"collections/wrapper"
// 	"strconv"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestPut(t *testing.T) {
// 	m := NewHashMap[wrapper.String, string]()

// 	assert.Equal(t, false, m.Contains("1"))
// 	m.Put("1", "20")
// 	assert.Equal(t, true, m.Contains("1"))
// 	assert.Equal(t, 1, m.Len())
// 	m.Put("0ab", "21")
// 	assert.Equal(t, true, m.Contains("0ab"))
// 	m.Put("1c", "21")
// 	assert.Equal(t, true, m.Contains("1c"))
// 	m.Put("1c", "90")
// 	v, _ := m.Get("1c")
// 	assert.Equal(t, "90", v)
// 	assert.Equal(t, 3, m.Len())

// }

// func TestPutIfAbsent(t *testing.T) {
// 	m := NewHashMap[wrapper.String, string]()

// 	assert.Equal(t, true, m.PutIfAbsent("222", "1"))
// 	m.Put("1", "20")
// 	p := m.PutIfAbsent("1", "21")
// 	assert.Equal(t, false, p)
// 	assert.Equal(t, true, m.PutIfAbsent("22", "23"))
// 	v, _ := m.Get("22")
// 	assert.Equal(t, "23", v)

// }

// func TestRemove(t *testing.T) {
// 	m := NewHashMap[wrapper.String, string]()

// 	assert.Equal(t, false, m.Remove("1"))
// 	m.Put("1", "20")
// 	assert.Equal(t, 1, m.Len())
// 	m.Remove("1")
// 	assert.Equal(t, true, m.Empty())
// 	_, b := m.Get("1")
// 	assert.Equal(t, false, b)

// }

// func TestRemoveAll(t *testing.T) {
// 	m := NewHashMap[wrapper.String, string]()

// 	m.Put("1", "20")
// 	m.Put("2", "20")
// 	m.Put("3", "20")

// 	l := list.NewList[wrapper.String]()
// 	l.Add("2")
// 	l.Add("3")

// 	m.RemoveAll(l)
// 	assert.Equal(t, 1, m.Len())
// 	assert.Equal(t, false, m.Contains("2"))
// }

// func TestResize(t *testing.T) {

// 	m := NewHashMap[wrapper.Integer, int]()
// 	assert.Equal(t, 16, m.Capacity())
// 	assert.Equal(t, float32(0), m.LoadFactor())
// 	for i := 1; i <= 16; i++ {
// 		m.Put(wrapper.Integer(i), i)
// 	}
// 	assert.Equal(t, 32, m.Capacity())
// 	assert.Equal(t, float32(0.5), m.LoadFactor())
// 	for i := 17; i <= 34; i++ {
// 		m.Put(wrapper.Integer(i), i)
// 	}
// 	assert.Equal(t, 64, m.Capacity())

// }

// func TestKeys(t *testing.T) {

// 	m := NewHashMap[wrapper.Integer, int]()
// 	for i := 1; i <= 6; i++ {
// 		m.Put(wrapper.Integer(i), i)
// 	}
// 	keys := []wrapper.Integer{1, 2, 3, 4, 5, 6}
// 	assert.ElementsMatch(t, keys, m.Keys())

// }

// func TestValues(t *testing.T) {

// 	m := NewHashMap[wrapper.Integer, int]()
// 	for i := 1; i <= 6; i++ {
// 		m.Put(wrapper.Integer(i), i)
// 	}
// 	values := []int{1, 2, 3, 4, 5, 6}
// 	assert.ElementsMatch(t, values, m.Values())

// }

// func TestIterator(t *testing.T) {

// 	m := NewHashMap[wrapper.String, int]()
// 	for i := 1; i <= 20; i++ {
// 		m.Put(wrapper.String(strconv.Itoa(i)), i)
// 	}

// 	keys := make([]wrapper.String, 0)
// 	values := make([]int, 0)
// 	it := m.Iterator()
// 	for it.HasNext() {
// 		entry := it.Next()
// 		keys = append(keys, entry.Key())
// 		values = append(values, entry.Value())
// 	}
// 	assert.ElementsMatch(t, m.Keys(), keys)
// 	assert.ElementsMatch(t, m.Values(), values)

// 	t.Run("panics", func(t *testing.T) {
// 		// If the function panics, recover() will
// 		// return a non nil value.
// 		defer func() {
// 			if r := recover(); r != nil {
// 				assert.Equal(t, NoNextElementError, r.(error))
// 			}
// 		}()

// 		it.Next()
// 	})

// 	it.Cycle()
// 	entry := it.Next()
// 	assert.Equal(t, keys[0], entry.Key())
// 	assert.Equal(t, values[0], entry.Value())

// }

// func TestEquals(t *testing.T) {

// 	m := NewHashMap[wrapper.Integer, int]()
// 	for i := 0; i <= 5; i++ {
// 		m.Put(wrapper.Integer(i), i)
// 	}

// 	assert.Equal(t, true, m.Equals(m, func(a, b int) bool { return a == b }))
// 	other := NewHashMap[wrapper.Integer, int]()
// 	assert.Equal(t, false, m.Equals(other, func(a, b int) bool { return a == b }))
// 	other.PutAll(m)
// 	assert.Equal(t, true, m.Equals(other, func(a, b int) bool { return a == b }))

// 	m.Clear()
// 	other.Clear()

// 	m.Put(1, 2)
// 	other.Put(1, 4)

// 	assert.Equal(t, false, m.Equals(other, func(a, b int) bool { return a == b }))

// }

// func TestMap(t *testing.T) {
// 	m := NewHashMap[wrapper.Integer, int]()
// 	for i := 1; i <= 6; i++ {
// 		m.Put(wrapper.Integer(i), i)
// 	}

// 	other := m.Map(func(e _map.MapEntry[wrapper.Integer, int]) _map.MapEntry[wrapper.Integer, int] {
// 		k := e.Key() + 2
// 		v := e.Value() + 10
// 		return _map.NewMapEntry(k, v)
// 	})

// 	keys := []wrapper.Integer{3, 4, 5, 6, 7, 8}
// 	values := []int{11, 12, 13, 14, 15, 16}
// 	assert.ElementsMatch(t, keys, other.Keys())
// 	assert.ElementsMatch(t, values, other.Values())

// }

// func TestFilter(t *testing.T) {
// 	m := NewHashMap[wrapper.Integer, int]()
// 	for i := 1; i <= 6; i++ {
// 		m.Put(wrapper.Integer(i), i)
// 	}

// 	other := m.Filter(func(e _map.MapEntry[wrapper.Integer, int]) bool {
// 		return e.Key()%2 == 0
// 	})

// 	keys := []wrapper.Integer{2, 4, 6}
// 	assert.ElementsMatch(t, keys, other.Keys())

// }
