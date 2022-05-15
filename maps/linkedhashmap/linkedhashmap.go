package linkedhashmap

import (
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

// linkedMapEntry a type for entries of a LinkedHashMap.
type linkedMapEntry[K types.Hashable[K], V any] struct {
	key   K
	value V
	prev  *linkedMapEntry[K, V]
	next  *linkedMapEntry[K, V]
}

// New creates an empty LinkedHashMap.
func New[K types.Hashable[K], V any]() *LinkedHashMap[K, V] {
	data := hashmap.New[K, *linkedMapEntry[K, V]]()
	linkedMap := LinkedHashMap[K, V]{data: data}
	return &linkedMap
}

// linkedHashMapIterator an iterator for moving through the keys and value of a LinkedHashMap.
type linkedHashMapIterator[K types.Hashable[K], V any] struct {
	head *linkedMapEntry[K, V]
	n    *linkedMapEntry[K, V]
}

// Cycle resets the iterator.
func (iterator *linkedHashMapIterator[K, V]) Cycle() {
	iterator.n = iterator.head
}

// HasNext checks if the iterator has a next element to yield.
func (iterator *linkedHashMapIterator[K, V]) HasNext() bool {
	return iterator.n != nil
}

// Next returns the next element in the iterator it. Will panic if iterator has been exhausted.
func (iter *linkedHashMapIterator[K, V]) Next() maps.MapEntry[K, V] {
	if !iter.HasNext() {
		panic(iterator.NoNextElementError)
	}
	entry := maps.MapEntry[K, V]{Key: iter.n.key, Value: iter.n.value}
	iter.n = iter.n.next
	return entry
}

// Iterator returns an iterator for the map.
func (m *LinkedHashMap[K, V]) Iterator() maps.MapIterator[K, V] {
	it := linkedHashMapIterator[K, V]{head: m.head, n: m.head}
	return &it
}

// Put associates the specified value with the specified key in the map. If the key already exists then its value will be updated. It
// returns the old value associated with the key or zero value if no previous association.
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
	var e V
	return e
}

// PutIfAbsent adds the value with the specified key to the map only if the key has not been mapped already.
func (linkedHashMap *LinkedHashMap[K, V]) PutIfAbsent(key K, value V) bool {
	if linkedHashMap.data.ContainsKey(key) {
		return false
	}
	linkedHashMap.Put(key, value)
	return true
}

// PutAll adds all the values from other map into the map. Note this has the side effect that if a key
// is present in the map and in other map then the associated value  in m will be replaced by the associated value  in other.
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

// Get retrieves the value associated with key in the map m. If there is no such value the zero value is returned along with false.
func (linkedHashMap *LinkedHashMap[K, V]) Get(key K) (V, bool) {
	entry, ok := linkedHashMap.data.Get(key)
	if ok {
		return entry.value, ok
	}
	var e V
	return e, ok
}

// ContainsKey checks if the map contains a mapping for the key.
func (linkedHashMap *LinkedHashMap[K, V]) ContainsKey(k K) bool {
	return linkedHashMap.data.ContainsKey(k)
}

// ContainsValue checks if the map has an entry whose value is the specified value. func equals is used to compare value for equality.
func (linkedHashMap *LinkedHashMap[K, V]) ContainsValue(value V, equals func(a, b V) bool) bool {
	iterator := linkedHashMap.data.Iterator()
	for iterator.HasNext() {
		entry := iterator.Next()
		if equals(entry.Value.value, value) {
			return true
		}
	}
	return false
}

// Remove removes the map entry <k,V> from map m if it exists.
func (linkedHashMap *LinkedHashMap[K, V]) Remove(key K) (V, bool) {
	entry, ok := linkedHashMap.data.Get(key)
	if !ok {
		var e V
		return e, false
	}
	if entry == linkedHashMap.head { // if was a head
		linkedHashMap.head = entry.next
		prevEntry, ok := linkedHashMap.data.Remove(key)
		return prevEntry.value, ok
	}
	if entry == linkedHashMap.tail { // if was a tail
		linkedHashMap.tail = entry.prev
	}
	entry.prev.next = entry.next
	entry.prev = nil
	entry.next = nil
	entry = nil
	prevEntry, ok := linkedHashMap.data.Remove(key)
	return prevEntry.value, ok
}

// RemoveAll removes all keys that are in the specified iterable from the map.
func (linkedHashMap *LinkedHashMap[K, V]) RemoveAll(keys iterator.Iterable[K]) {
	iterator := keys.Iterator()
	for iterator.HasNext() {
		linkedHashMap.Remove(iterator.Next())
	}
}

// LoadFactor computes the load factor of the map m.
func (linkedHashMap *LinkedHashMap[K, V]) LoadFactor() float32 {
	return linkedHashMap.data.LoadFactor()
}

// Values collects all the values of the map into a slice.
func (linkedHashMap *LinkedHashMap[K, V]) Values() []V {
	values := make([]V, linkedHashMap.Len())
	i := 0
	for curr := linkedHashMap.head; curr != nil; curr = curr.next {
		values[i] = curr.value
		i++
	}
	return values
}

// Keys collects the keys of the map into a slice.
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

// Clear removes all entries in the map.
func (linkedHashMap *LinkedHashMap[K, V]) Clear() {
	linkedHashMap.head = nil
	linkedHashMap.tail = nil
	linkedHashMap.data.Clear()
}

// Equals check if map m is equal to map other. This checks that the two maps have the same entries (k,v), the values are compared
// using the specified equals function (a is the mapped value in map m and b is the mapped value in the other map). Keys are compared using their corresponding Equals method.
// Only returns true if the 2 maps are the same reference or have the same size and entries.
func (linkedHashMap *LinkedHashMap[K, V]) Equals(other *LinkedHashMap[K, V], equals func(a V, b V) bool) bool {
	if linkedHashMap == other {
		return true
	} else if linkedHashMap.Len() != other.Len() {
		return false
	} else {
		if linkedHashMap.Empty() && other.Empty() {
			return true
		}
		iterator := linkedHashMap.Iterator()
		for iterator.HasNext() {
			entry := iterator.Next()
			value, ok := other.Get(entry.Key)
			if ok && equals(entry.Value, value) {
			} else {
				return false
			}
		}
		return true
	}
}

// Map applies a transformation on an entry of m i.e f((k,v)) -> (k*,v*) , using a function f and returns a new HashMap of which its keys
// and values have been transformed.
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

// Filter filters the HashMap m using a predicate function that indicates whether an entry should be kept or not in a
// HashMap to be returned.
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
