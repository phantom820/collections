package linkedhashmap

import (
	"fmt"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/types"
)

// LinkedHashMap a HashMap that maintains insertion order.
type LinkedHashMap[K types.Hashable[K], V any] struct {
	data *hashmap.HashMap[K, *linkedMapEntry[K, V]]
	head *linkedMapEntry[K, V]
	tail *linkedMapEntry[K, V]
}

type linkedMapEntry[K types.Hashable[K], V any] struct {
	key   K
	value V
	prev  *linkedMapEntry[K, V]
	next  *linkedMapEntry[K, V]
}

// New creates an empty LinkedHashMap.
func New[K types.Hashable[K], V any]() *LinkedHashMap[K, V] {
	data := hashmap.New[K, *linkedMapEntry[K, V]]()
	m := LinkedHashMap[K, V]{data: data}
	return &m
}

// linkedHashMapIterator an iterator for moving through the keys and value of a LinkedHashMap.
type linkedHashMapIterator[K types.Hashable[K], V any] struct {
	head *linkedMapEntry[K, V]
	n    *linkedMapEntry[K, V]
}

// Cycle resets the iterator.
func (it *linkedHashMapIterator[K, V]) Cycle() {
	it.n = it.head
}

// HasNext checks if the iterator has a next value to yield.
func (it *linkedHashMapIterator[K, V]) HasNext() bool {
	return it.n != nil
}

// Next returns the next element in the iterator it. Will panic if iterator has been exhausted.
func (it *linkedHashMapIterator[K, V]) Next() maps.MapEntry[K, V] {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	entry := maps.MapEntry[K, V]{Key: it.n.key, Value: it.n.value}
	it.n = it.n.next
	return entry
}

// Iterator returns an iterator for the map.
func (m *LinkedHashMap[K, V]) Iterator() maps.MapIterator[K, V] {
	it := linkedHashMapIterator[K, V]{head: m.head, n: m.head}
	return &it
}

// Put associates the specified value with the specified key in the map. If the key already exists then its value will be updated. It
// returns the old value associated with the key or zero value if no previous association.
func (m *LinkedHashMap[K, V]) Put(k K, v V) V {
	if m.Empty() {
		entry := linkedMapEntry[K, V]{key: k, value: v, prev: nil, next: nil}
		m.head = &entry
		m.tail = &entry
		var e V
		m.data.Put(k, &entry)
		return e
	}

	if entry, ok := m.data.Get(k); ok { // Key already mapped no need to change prev and next values.
		temp := entry.value
		entry.value = v
		return temp
	}
	entry := linkedMapEntry[K, V]{key: k, value: v, prev: m.tail, next: nil}
	entry.prev.next = &entry
	entry.prev = m.tail
	entry.next = nil
	m.tail = &entry
	m.data.Put(k, &entry)
	var e V
	return e
}

// PutIfAbsent adds the value with the specified key to the map only if the key has not been mapped already.
func (m *LinkedHashMap[K, V]) PutIfAbsent(k K, v V) bool {
	if m.data.ContainsKey(k) {
		return false
	}
	m.Put(k, v)
	return true
}

// PutAll adds all the values from other map into the map. Note this has the side effect that if a key
// is present in the map and in other map then the associated value  in m will be replaced by the associated value  in other.
func (m *LinkedHashMap[K, V]) PutAll(other maps.Map[K, V]) {
	for _, k := range other.Keys() {
		v, _ := other.Get(k)
		m.Put(k, v)
	}
}

// Len returns the size of the map.
func (m *LinkedHashMap[K, V]) Len() int {
	return m.data.Len()
}

// Get retrieves the value associated with key in the map m. If there is no such value the zero value is returned along with false.
func (m *LinkedHashMap[K, V]) Get(k K) (V, bool) {
	entry, b := m.data.Get(k)
	if b {
		return entry.value, b
	}
	var e V
	return e, b
}

// ContainsKey checks if the map contains a mapping for the key.
func (m *LinkedHashMap[K, V]) ContainsKey(k K) bool {
	return m.data.ContainsKey(k)
}

// ContainsValue checks if the map has an entry whose value is the specified value. func equals is used to compare value for equality.
func (m *LinkedHashMap[K, V]) ContainsValue(v V, equals func(a, b V) bool) bool {
	it := m.data.Iterator()
	for it.HasNext() {
		entry := it.Next()
		if equals(entry.Value.value, v) {
			return true
		}
	}
	return false
}

// Remove removes the map entry <k,V> from map m if it exists.
func (m *LinkedHashMap[K, V]) Remove(k K) (V, bool) {
	entry, ok := m.data.Get(k)
	if !ok {
		var e V
		return e, false
	}
	if entry == m.head { // if was a head
		m.head = entry.next
		v, b := m.data.Remove(k)
		return v.value, b
	}
	if entry == m.tail { // if was a tail
		m.tail = entry.prev
	}
	entry.prev.next = entry.next
	entry.prev = nil
	entry.next = nil
	entry = nil
	v, b := m.data.Remove(k)
	return v.value, b
}

// RemoveAll removes all keys that are in the specified iterable from m.
func (m *LinkedHashMap[K, V]) RemoveAll(keys iterator.Iterable[K]) {
	it := keys.Iterator()
	for it.HasNext() {
		m.Remove(it.Next())
	}
}

// LoadFactor computes the load factor of the map m.
func (m *LinkedHashMap[K, V]) LoadFactor() float32 {
	return m.data.LoadFactor()
}

// Values collects all the values of the map into a slice.
func (m *LinkedHashMap[K, V]) Values() []V {
	data := make([]V, m.Len())
	i := 0
	for curr := m.head; curr != nil; curr = curr.next {
		data[i] = curr.value
		i++
	}
	return data
}

// Keys collects the keys of the map into a slice.
func (m *LinkedHashMap[K, V]) Keys() []K {
	data := make([]K, m.Len())
	i := 0
	for curr := m.head; curr != nil; curr = curr.next {
		data[i] = curr.key
		i++
	}
	return data
}

// Empty checks if the map is empty.
func (m *LinkedHashMap[K, V]) Empty() bool {
	return m.data.Empty()
}

// Clear removes all entries in the map.
func (m *LinkedHashMap[K, V]) Clear() {
	m.head = nil
	m.tail = nil
	m.data.Clear()
}

// Equals check if map m is equal to map other. This checks that the two maps have the same entries (k,v), the values are compared
// using the specified equals function (a is the mapped value in map m and b is the mapped value in the other map). Keys are compared using their corresponding Equals method.
// Only returns true if the 2 maps are the same reference or have the same size and entries.
func (m *LinkedHashMap[K, V]) Equals(other *LinkedHashMap[K, V], equals func(a V, b V) bool) bool {
	if m == other {
		return true
	} else if m.Len() != other.Len() {
		return false
	} else {
		if m.Empty() && other.Empty() {
			return true
		}
		it := m.Iterator()
		for it.HasNext() {
			entry := it.Next()
			v, b := other.Get(entry.Key)
			fmt.Println(v)
			fmt.Println(b)
			if b && equals(entry.Value, v) {
			} else {
				return false
			}
		}
		return true
	}
}

// Map applies a transformation on an entry of m i.e f((k,v)) -> (k*,v*) , using a function f and returns a new HashMap of which its keys
// and values have been transformed.
func (m LinkedHashMap[K, V]) Map(f func(e maps.MapEntry[K, V]) maps.MapEntry[K, V]) *LinkedHashMap[K, V] {
	newMap := New[K, V]()
	it := m.Iterator()
	for it.HasNext() {
		oldEntry := it.Next()
		newEntry := f(oldEntry)
		newMap.Put(newEntry.Key, newEntry.Value)
	}
	return newMap
}

// Filter filters the HashMap m using a predicate function that indicates whether an entry should be kept or not in a
// HashMap to be returned.
func (m LinkedHashMap[K, V]) Filter(f func(e maps.MapEntry[K, V]) bool) *LinkedHashMap[K, V] {
	newMap := New[K, V]()
	it := m.Iterator()
	for it.HasNext() {
		entry := it.Next()
		if f(entry) {
			newMap.Put(entry.Key, entry.Value)
		}
	}
	return newMap
}
