// package hashmap defines a wrapper around a standard map to provide extended functionality.
package hashmap

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types/optional"
	"github.com/phantom820/collections/types/pair"
)

// HashMap wrapper around a map[K]V.
type HashMap[K comparable, V any] map[K]V

// New creates a map with the given key, value pairs.
func New[K comparable, V any](pairs ...pair.Pair[K, V]) HashMap[K, V] {
	var hashMap HashMap[K, V] = make(map[K]V)
	for _, pair := range pairs {
		hashMap.Put(pair.Key(), pair.Value())
	}
	return hashMap
}

// Put adds a new key/value pair to the map and optionally returns previously bound value.
func (hashMap HashMap[K, V]) Put(key K, value V) optional.Optional[V] {
	if storedValue, ok := hashMap[key]; ok {
		hashMap[key] = value
		return optional.Of(storedValue)
	}
	hashMap[key] = value
	return optional.Empty[V]()
}

// PutIfAbsent adds a new key/value pair to the map if the key is not already bounded and optionally returns bound value.
func (hashMap HashMap[K, V]) PutIfAbsent(key K, value V) optional.Optional[V] {
	if storedValue, ok := hashMap[key]; ok {
		return optional.Of(storedValue)
	}
	hashMap[key] = value
	return optional.Empty[V]()
}

// Get optionally returns the value associated with a key.
func (hashMap HashMap[K, V]) Get(key K) optional.Optional[V] {
	if storedValue, ok := hashMap[key]; ok {
		return optional.Of(storedValue)
	}
	return optional.Empty[V]()
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

// Remove removes a key from the map, returning the value associated previously with that key as an option.
func (hashMap HashMap[K, V]) Remove(key K) optional.Optional[V] {
	if storedValue, ok := hashMap[key]; ok {
		delete(hashMap, key)
		return optional.Of(storedValue)
	}
	return optional.Empty[V]()
}

// RemoveIf removes all the key, value mapping in which the key matches the given predicate.
func (hashMap HashMap[K, V]) RemoveIf(f func(K) bool) bool {
	n := hashMap.Len()
	keysToRemove := make([]K, 0)
	for key, _ := range hashMap {
		if f(key) {
			keysToRemove = append(keysToRemove, key)
		}
	}

	for _, key := range keysToRemove {
		delete(hashMap, key)
	}
	return n != hashMap.Len()

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

// Len returns the size of the map.
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
func (hashMap HashMap[K, V]) Iterator() iterator.Iterator[pair.Pair[K, V]] {
	return &mapIterator[K, V]{initialized: false, index: 0, hashMap: hashMap, entries: make([]pair.Pair[K, V], 0)}
}

// mapIterator implementation of an mapIterator for [HashMap].
type mapIterator[K comparable, V any] struct {
	initialized bool
	index       int
	hashMap     map[K]V
	entries     []pair.Pair[K, V]
}

// HasNext returns true if the iterator has more elements.
func (it *mapIterator[K, V]) HasNext() bool {
	if it.hashMap == nil {
		return false
	} else if !it.initialized {
		it.initialized = true
		for key, value := range it.hashMap {
			it.entries = append(it.entries, pair.Of(key, value))
		}
	}
	return it.index < len(it.entries)

}

// Next returns the next element in the iterator.
func (it *mapIterator[K, V]) Next() pair.Pair[K, V] {
	if !it.HasNext() {
		panic(errors.NoSuchElement())
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

// Equals return true if the map is is equal to the given map. Two maps are equal if they contain the same
// key, value pairs.
func (hashMap HashMap[K, V]) Equals(other collections.Map[K, V], equals func(V, V) bool) bool {
	if hashMap.Len() != other.Len() {
		return false
	}
	it := other.Iterator()
	for it.HasNext() {
		pair := it.Next()
		result := hashMap.Get(pair.Key())
		if result.Empty() {
			return false
		} else if !equals(pair.Value(), result.Value()) {
			return false
		}
	}
	return true
}
