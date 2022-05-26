// Package linkedhashmap provides an implementation of a LinkedHashMap which is a map that preserves insertion order oof its entries.
package linkedhashmap

import (
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/types"
)

// LinkedHashMap an implementation of a map that keeps track of insertion order of entries.
type LinkedHashMap[K types.Hashable[K], V any] struct {
	data *hashmap.HashMap[K, *linkedMapEntry[K, V]]
	head *linkedMapEntry[K, V]
	tail *linkedMapEntry[K, V]
}

// linkedMapEntry a type for entries of a LinkedHashMap.
type linkedMapEntry[K types.Hashable[K], V any] struct {
	key   K
	value V
	prev  *linkedMapEntry[K, V]
	next  *linkedMapEntry[K, V]
}

// New creates and returns an empty LinkedHashMap.
func New[K types.Hashable[K], V any]() *LinkedHashMap[K, V] {
	data := hashmap.New[K, *linkedMapEntry[K, V]]()
	linkedMap := LinkedHashMap[K, V]{data: data}
	return &linkedMap
}

// linkedHashMapIterator  a type to implement an iterator for the map.
type linkedHashMapIterator[K types.Hashable[K], V any] struct {
	head *linkedMapEntry[K, V]
	node *linkedMapEntry[K, V]
}

// Cycle resets the iterator.
func (it *linkedHashMapIterator[K, V]) Cycle() {
	it.node = it.head
}

// HasNext checks if the iterator has a next element to yield.
func (it *linkedHashMapIterator[K, V]) HasNext() bool {
	return it.node != nil
}

// Next yields the next element in the iterator. Will panic if the iterator has no next element.
func (it *linkedHashMapIterator[K, V]) Next() maps.MapEntry[K, V] {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	entry := maps.MapEntry[K, V]{Key: it.node.key, Value: it.node.value}
	it.node = it.node.next
	return entry
}

// Iterator returns an iterator for the map.
func (m *LinkedHashMap[K, V]) Iterator() maps.MapIterator[K, V] {
	it := linkedHashMapIterator[K, V]{head: m.head, node: m.head}
	return &it
}

// Put inserts the entry <key,value> into the map. If an entry with the given key already exists then its value is updated. Returns the previous value
// associated with the key or zero value if there is no previous value.
func (linkedHashMap *LinkedHashMap[K, V]) Put(key K, value V) V {
	if linkedHashMap.Empty() {
		entry := linkedMapEntry[K, V]{key: key, value: value, prev: nil, next: nil}
		linkedHashMap.head = &entry
		linkedHashMap.tail = &entry
		var prevValue V
		linkedHashMap.data.Put(key, &entry)
		return prevValue
	}
	if entry, ok := linkedHashMap.data.Get(key); ok { // Key already mapped no need to change prev and next values.
		temp := entry.value
		entry.value = value
		return temp
	}
	entry := linkedMapEntry[K, V]{key: key, value: value, prev: linkedHashMap.tail, next: nil}
	entry.prev.next = &entry
	entry.prev = linkedHashMap.tail
	entry.next = nil
	linkedHashMap.tail = &entry
	linkedHashMap.data.Put(key, &entry)
	var zero V
	return zero
}

// PutIfAbsent inserts the entry <key,value> into the map if the key does not already exist in the map. Returns true if the new entry was made.
func (linkedHashMap *LinkedHashMap[K, V]) PutIfAbsent(key K, value V) bool {
	if linkedHashMap.data.ContainsKey(key) {
		return false
	}
	linkedHashMap.Put(key, value)
	return true
}

// PutAll adds all the values from another map into the map. Note this has the side effect that if a key
// is present in the map and in the passed map then the associated value in the map will be replaced by the associated value from the passed map.
func (linkedHashMap *LinkedHashMap[K, V]) PutAll(other maps.Map[K, V]) {
	for _, key := range other.Keys() {
		value, _ := other.Get(key)
		linkedHashMap.Put(key, value)
	}
}

// Len returns the size of the map.
func (linkedHashMap *LinkedHashMap[K, V]) Len() int {
	return linkedHashMap.data.Len()
}

// Get retrieves the value associated with the key in the map. Returns a value and a boolean indicating if the value is valid or invalid.
// An invalid value results when there is no entry for the given key and the zero value is returned.
func (linkedHashMap *LinkedHashMap[K, V]) Get(key K) (V, bool) {
	entry, ok := linkedHashMap.data.Get(key)
	if ok {
		return entry.value, ok
	}
	var e V
	return e, ok
}

// ContainsKey checks if the map contains an entry with the given key.
func (linkedHashMap *LinkedHashMap[K, V]) ContainsKey(k K) bool {
	return linkedHashMap.data.ContainsKey(k)
}

// ContainsValue checks if the map has an entry whose value is the specified value. The function equals is used to check values for equality.
func (linkedHashMap *LinkedHashMap[K, V]) ContainsValue(value V, equals func(a, b V) bool) bool {
	it := linkedHashMap.data.Iterator()
	for it.HasNext() {
		entry := it.Next()
		if equals(entry.Value.value, value) {
			return true
		}
	}
	return false
}

// Remove removes the map entry <key,value> from the map if it exists. Returns the previous value associated with the key and a boolean indicating if the returned
// values is valid or invalid. An invalid value results when there is no entry in the map associated with the given key.
func (linkedHashMap *LinkedHashMap[K, V]) Remove(key K) (V, bool) {
	entry, ok := linkedHashMap.data.Get(key)
	if !ok {
		var e V
		return e, false
	}
	if entry == linkedHashMap.head { // if was a head.
		linkedHashMap.head = entry.next
		prevEntry, ok := linkedHashMap.data.Remove(key)
		return prevEntry.value, ok
	}
	if entry == linkedHashMap.tail { // if was a tail.
		linkedHashMap.tail = entry.prev
	}
	entry.prev.next = entry.next
	entry.prev = nil
	entry.next = nil
	entry = nil
	prevEntry, ok := linkedHashMap.data.Remove(key)
	return prevEntry.value, ok
}

// RemoveAll removes all key entries from the map that appear in the iterable keys.
func (linkedHashMap *LinkedHashMap[K, V]) RemoveAll(keys iterator.Iterable[K]) {
	it := keys.Iterator()
	for it.HasNext() {
		linkedHashMap.Remove(it.Next())
	}
}

// LoadFactor returns the load factor of the map.
func (linkedHashMap *LinkedHashMap[K, V]) LoadFactor() float32 {
	return linkedHashMap.data.LoadFactor()
}

// Values returns a slice containing all the values in the map.
func (linkedHashMap *LinkedHashMap[K, V]) Values() []V {
	values := make([]V, linkedHashMap.Len())
	i := 0
	for curr := linkedHashMap.head; curr != nil; curr = curr.next {
		values[i] = curr.value
		i++
	}
	return values
}

// Keys returns a slice containing all the keys in the map.
func (linkedHashMap *LinkedHashMap[K, V]) Keys() []K {
	keys := make([]K, linkedHashMap.Len())
	i := 0
	for curr := linkedHashMap.head; curr != nil; curr = curr.next {
		keys[i] = curr.key
		i++
	}
	return keys
}

// Empty checks if the map is empty.
func (linkedHashMap *LinkedHashMap[K, V]) Empty() bool {
	return linkedHashMap.data.Empty()
}

// Clear removes all entries from the map.
func (linkedHashMap *LinkedHashMap[K, V]) Clear() {
	linkedHashMap.head = nil
	linkedHashMap.tail = nil
	linkedHashMap.data.Clear()
}

// Equals checks if the map is equal to map other. This checks that the two maps have the same entries (k,v), the values are compared
// using the specified equals function for two values. Keys are compared using their corresponding Equals method.
// Only returns true if the 2 maps are the same reference or have the same size and entries. The insertion order is not taken into account.
func (linkedHashMap *LinkedHashMap[K, V]) Equals(other *LinkedHashMap[K, V], equals func(a V, b V) bool) bool {
	if linkedHashMap == other {
		return true
	} else if linkedHashMap.Len() != other.Len() {
		return false
	} else {
		if linkedHashMap.Empty() && other.Empty() {
			return true
		}
		it := linkedHashMap.Iterator()
		for it.HasNext() {
			entry := it.Next()
			value, ok := other.Get(entry.Key)
			if ok && equals(entry.Value, value) {
			} else {
				return false
			}
		}
		return true
	}
}

// Map applies a transformation on an entry of the map i.e f((k,v)) -> (k*,v*) , using the function f and returns a new map with the
// transformed entries.
func (linkedHashMap LinkedHashMap[K, V]) Map(f func(e maps.MapEntry[K, V]) maps.MapEntry[K, V]) *LinkedHashMap[K, V] {
	newMap := New[K, V]()
	iterator := linkedHashMap.Iterator()
	for iterator.HasNext() {
		oldEntry := iterator.Next()
		newEntry := f(oldEntry)
		newMap.Put(newEntry.Key, newEntry.Value)
	}
	return newMap
}

// Filter filters the map using the predicate function  f and returns a new map containing only entries that satisfy the predicate.
func (linkedHashMap LinkedHashMap[K, V]) Filter(f func(e maps.MapEntry[K, V]) bool) *LinkedHashMap[K, V] {
	newMap := New[K, V]()
	iterator := linkedHashMap.Iterator()
	for iterator.HasNext() {
		entry := iterator.Next()
		if f(entry) {
			newMap.Put(entry.Key, entry.Value)
		}
	}
	return newMap
}
