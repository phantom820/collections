package hashmap

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections/maps"
)

const (
	DEFAULT_CAPACITY = 16 // Initial capacity of the hash table.
)

// HashMap wrapper around a map[K]V.
type HashMap[K comparable, V any] map[K]V

// New creates an empty HashMap.
func New[K comparable, V any]() HashMap[K, V] {
	return make(map[K]V, DEFAULT_CAPACITY)
}

// Put associates the specified value with the specified key in the map.
func (hashMap HashMap[K, V]) Put(key K, value V) V {
	storedValue := hashMap[key]
	hashMap[key] = value
	return storedValue
}

// PutIfAbsent associates the specified key with the given value if the key is not already mapped. Will return the
// current value if present otherwise the zero value.
func (hashMap HashMap[K, V]) PutIfAbsent(key K, value V) bool {
	if _, ok := hashMap[key]; ok {
		return false
	}
	hashMap[key] = value
	return true
}

// Get returns the value to which the specified key is mapped, or the zero value if the key is not present.
func (hashMap HashMap[K, V]) Get(key K) V {
	return hashMap[key]
}

// GetIf returns the values mapped by keys that match the given predicate.
func (hashMap HashMap[K, V]) GetIf(f func(K) bool) []V {
	values := make([]V, 0)
	for key, value := range hashMap {
		if f(key) {
			values = append(values, value)
		}
	}
	return values
}

// Remove removes the mapping for the specified key from the map if present
func (hashMap HashMap[K, V]) Remove(key K) V {
	storedValue := hashMap[key]
	delete(hashMap, key)
	return storedValue
}

// RemoveIf removes all the key, value mapping in which the key matches the given predicate.
func (hashMap HashMap[K, V]) RemoveIf(f func(K) bool) {
	keysToRemove := make([]K, 0)
	for key, _ := range hashMap {
		if f(key) {
			keysToRemove = append(keysToRemove, key)
		}
	}

	for _, key := range keysToRemove {
		delete(hashMap, key)
	}
}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (hashMap HashMap[K, V]) ContainsKey(key K) bool {
	_, ok := hashMap[key]
	return ok
}

// ContainsValue returns true if this map maps one or more keys to the specified value.
func (hashMap HashMap[K, V]) ContainsValue(value V, equals func(v1, v2 V) bool) bool {
	for _, storedValue := range hashMap {
		if equals(storedValue, value) {
			return true
		}
	}
	return false
}

// Clear removes all of the mappings from this map.
func (hashMap HashMap[K, V]) Clear() {
	for key := range hashMap {
		delete(hashMap, key)
	}
}

// Keys returns a slice containing the keys in the map.
func (hashMap HashMap[K, V]) Keys() []K {
	keys := make([]K, len(hashMap))
	i := 0
	for key, _ := range hashMap {
		keys[i] = key
		i++
	}
	return keys
}

// Values returns a slice containing the values in the map.
func (hashMap HashMap[K, V]) Values() []V {
	values := make([]V, len(hashMap))
	i := 0
	for _, value := range hashMap {
		values[i] = value
		i++
	}
	return values
}

// Len returns the number of key, value mappings in the map.
func (hashMap HashMap[K, V]) Len() int {
	return len(hashMap)
}

// Empty returns true if the map has no elements.
func (hashMap HashMap[K, V]) Empty() bool {
	return len(hashMap) == 0
}

// ForEach performs the given action for each key, value mapping in the map.
func (hashMap HashMap[K, V]) ForEach(f func(K, V)) {
	for key, value := range hashMap {
		f(key, value)
	}
}

// Iterator returns an iterator over the map.
func (hashMap HashMap[K, V]) Iterator() maps.MapIterator[K, V] {
	return &iterator[K, V]{initialized: false, index: 0, hashMap: hashMap, entries: make([]maps.Entry[K, V], 0)}
}

// iterator implementation of an iterator for [HashMap].
type iterator[K comparable, V any] struct {
	initialized bool
	index       int
	hashMap     map[K]V
	entries     []maps.Entry[K, V]
}

// HasNext returns true if the iterator has more elements.
func (it *iterator[K, V]) HasNext() bool {
	if it.hashMap == nil {
		return false
	} else if !it.initialized {
		it.initialized = true
		for key, value := range it.hashMap {
			it.entries = append(it.entries, maps.NewEntry(key, value))
		}
	}
	return it.index < len(it.entries)

}

// Next returns the next element in the iterator.
func (it *iterator[K, V]) Next() maps.Entry[K, V] {
	if !it.HasNext() {
		panic("iterator things shoould panic here")
	}
	index := it.index
	it.index++
	return it.entries[index]
}

// String returns the string representation of the map.
func (hashMap HashMap[K, V]) String() string {
	var sb strings.Builder
	if hashMap.Empty() {
		return "{}"
	}
	i := 0
	sb.WriteString("{")
	for key, value := range hashMap {
		if i == 0 {
			sb.WriteString(fmt.Sprintf("%v=%v", key, value))
			i++
		} else {
			sb.WriteString(fmt.Sprintf(", %v=%v", key, value))
		}
	}
	sb.WriteString("}")
	return sb.String()
}
