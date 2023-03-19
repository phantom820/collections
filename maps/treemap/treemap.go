// package treemap defines a map implementation backed by a red black tree to keep entries in a sorted order.
package treemap

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/maps/linkedhashmap"
	"github.com/phantom820/collections/trees/rbt"
	"github.com/phantom820/collections/types/optional"
	"github.com/phantom820/collections/types/pair"
)

// TreeMap implementation of a map in which entries are stored in a sorted order.
type TreeMap[K comparable, V any] struct {
	tree *rbt.RedBlackTree[K, V]
}

// New creates a map with the given key, value pairs. Keys are compared using the lessThan function which should satisfy.
// k1 < k2 => lessThan(k1, k2) = true and lessThan(k2,k1) = false.
// k1 = k2 => lessThan(k1,k2) = false and lessThan(k2,k1) = false.
// k1 > k2 -> lessThan(k1,k2) = false and lessThan(k2,k1) = true.
func New[K comparable, V any](lessThan func(k1, k2 K) bool, pairs ...pair.Pair[K, V]) *TreeMap[K, V] {
	treeMap := TreeMap[K, V]{rbt.New[K, V](lessThan)}
	for _, pair := range pairs {
		treeMap.Put(pair.Key(), pair.Value())
	}
	return &treeMap
}

// Put adds a new key/value pair to the map and optionally returns previously bound value.
func (treeMap *TreeMap[K, V]) Put(key K, value V) optional.Optional[V] {
	return treeMap.tree.Insert(key, value)
}

// PutIfAbsent adds a new key/value pair to the map if the key is not already bounded and optionally returns bound value.
func (treeMap *TreeMap[K, V]) PutIfAbsent(key K, value V) optional.Optional[V] {
	if storedValue := treeMap.tree.Get(key); !storedValue.Empty() {
		return storedValue
	}
	treeMap.tree.Insert(key, value)
	return optional.Empty[V]()
}

// Get optionally returns the value associated with a key.
func (treeMap *TreeMap[K, V]) Get(key K) optional.Optional[V] {
	return treeMap.tree.Get(key)
}

// GetIf returns the values mapped by keys that match the given predicate.
func (treeMap *TreeMap[K, V]) GetIf(f func(K) bool) []V {
	return treeMap.tree.GetIf(f)
}

// Remove removes a key from the map, returning the value associated previously with that key as an option.
func (treeMap TreeMap[K, V]) Remove(key K) optional.Optional[V] {
	return treeMap.tree.Delete(key)
}

// RemoveIf removes all the key, value mapping in which the key matches the given predicate.
func (treeMap *TreeMap[K, V]) RemoveIf(f func(K) bool) bool {
	n := treeMap.Len()
	keysToRemove := make([]K, 0)
	for _, key := range treeMap.tree.Keys() {
		if f(key) {
			keysToRemove = append(keysToRemove, key)
		}
	}

	for _, key := range keysToRemove {
		treeMap.tree.Delete(key)
	}
	return n != treeMap.Len()
}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (treeMap *TreeMap[K, V]) ContainsKey(key K) bool {
	return treeMap.tree.Search(key)
}

// ContainsValue returns true if this map maps one or more keys to the specified value.
func (treeMap *TreeMap[K, V]) ContainsValue(value V, equals func(v1, v2 V) bool) bool {
	it := treeMap.Iterator()
	for it.HasNext() {
		pair := it.Next()
		if equals(pair.Value(), value) {
			return true
		}
	}
	return false
}

// Clear removes all of the mappings from this map.
func (treeMap *TreeMap[K, V]) Clear() {
	treeMap.tree.Clear()
}

// Keys returns a slice containing the keys in the map.
func (treeMap *TreeMap[K, V]) Keys() []K {
	return treeMap.tree.Keys()
}

// Values returns a slice containing the values in the map.
func (treeMap *TreeMap[K, V]) Values() []V {
	return treeMap.tree.Values()
}

// Len returns the number of key, value mappings in the map.
func (treeMap *TreeMap[K, V]) Len() int {
	return treeMap.tree.Len()
}

// Empty returns true if the map has no elements.
func (treeMap *TreeMap[K, V]) Empty() bool {
	return treeMap.Len() == 0
}

// ForEach performs the given action for each key, value mapping in the map.
func (treeMap *TreeMap[K, V]) ForEach(f func(K, V)) {
	for _, pair := range treeMap.tree.Nodes() {
		f(pair.Key(), pair.Value())
	}
}

// Iterator returns an iterator over the map.
func (treeMap *TreeMap[K, V]) Iterator() iterator.Iterator[pair.Pair[K, V]] {
	return &mapIterator[K, V]{initialized: false, index: 0, entries: make([]pair.Pair[K, V], 0), initialize: treeMap.tree.Nodes}
}

// mapIterator implementation of an mapIterator for [HashMap].
type mapIterator[K comparable, V any] struct {
	initialized bool
	initialize  func() []pair.Pair[K, V]
	index       int
	entries     []pair.Pair[K, V]
}

// HasNext returns true if the iterator has more elements.
func (it *mapIterator[K, V]) HasNext() bool {
	if !it.initialized {
		it.initialized = true
		it.entries = it.initialize()
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
func (treeMap *TreeMap[K, V]) String() string {
	var sb strings.Builder
	if treeMap.Empty() {
		return "{}"
	}
	sb.WriteString("{")
	for i, node := range treeMap.tree.Nodes() {
		if i == 0 {
			sb.WriteString(fmt.Sprintf("%v=%v", node.Key(), node.Value()))
		} else {
			sb.WriteString(fmt.Sprintf(", %v=%v", node.Key(), node.Value()))
		}
	}
	sb.WriteString("}")
	return sb.String()
}

// Equals return true if the map is is equal to the given map. Two maps are equal if they contain the same
// key, value pairs.
func (treeMap TreeMap[K, V]) Equals(other collections.Map[K, V], equals func(V, V) bool) bool {
	if treeMap.Len() != other.Len() {
		return false
	}

	switch other.(type) {
	case hashmap.HashMap[K, V], *hashmap.HashMap[K, V], *linkedhashmap.LinkedHashMap[K, V]:
		{
			it := treeMap.Iterator()
			for it.HasNext() {
				pair := it.Next()
				result := other.Get(pair.Key())
				if result.Empty() {
					return false
				} else if !equals(pair.Value(), result.Value()) {
					return false
				}
			}
			return true
		}
	default:
		{
			it := other.Iterator()
			for it.HasNext() {
				pair := it.Next()
				result := treeMap.Get(pair.Key())
				if result.Empty() {
					return false
				} else if !equals(pair.Value(), result.Value()) {
					return false
				}
			}
			return true
		}
	}

}
