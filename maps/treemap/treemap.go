package treemap

import (
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/trees/rbt"
	"github.com/phantom820/collections/types"
)

// TreeMap a map that stores entries in sorted order.
type TreeMap[K types.Comparable[K], V any] struct {
	tree *rbt.RedBlackTree[K, V]
}

// New creates an empty TreeMap.
func New[K types.Comparable[K], V any]() *TreeMap[K, V] {
	treeMap := TreeMap[K, V]{tree: rbt.New[K, V]()}
	return &treeMap
}

// Put associates the specified value with the specified key in the map. If the key already exists then its value will be updated. It
// returns the old value associated with the key or zero value if no previous association.
func (treeMap *TreeMap[K, V]) Put(k K, v V) V {
	if val, ok := treeMap.tree.Get(k); ok {
		treeMap.tree.Update(k, v)
		return val
	}
	treeMap.tree.Insert(k, v)
	var val V
	return val
}

// PutIfAbsent adds the value with the specified key to the map only if the key has not been mapped already.
func (treeMap *TreeMap[K, V]) PutIfAbsent(k K, v V) bool {
	if _, ok := treeMap.tree.Get(k); ok {
		return false
	}
	treeMap.tree.Insert(k, v)
	return true
}

// PutAll adds all the values from other map into the map. Note this has the side effect that if a key
// is present in the map and in other map then the associated value  in m will be replaced by the associated value  in other.
func (treeMap *TreeMap[K, V]) PutAll(other maps.Map[K, V]) {
	for _, k := range other.Keys() {
		v, _ := other.Get(k)
		treeMap.Put(k, v)
	}
}

// Len returns the size of the map.
func (treeMap *TreeMap[K, V]) Len() int {
	return treeMap.tree.Len()
}

// Get retrieves the value associated with key in the map treeMap. If there is no such value the zero value is returned along with false.
func (treeMap *TreeMap[K, V]) Get(k K) (V, bool) {
	return treeMap.tree.Get(k)
}

// ContainsKey checks if the map contains a mapping for the key.
func (treeMap *TreeMap[K, V]) ContainsKey(k K) bool {
	return treeMap.tree.Search(k)
}

// treeMapIterator an iterator for moving through the keys and value of a HashMap.
type treeMapIterator[K types.Comparable[K], V any] struct {
	index  int
	keys   []K
	values []V
}

// Cycle resets the iterator.
func (it *treeMapIterator[K, V]) Cycle() {
	it.index = 0
}

// HasNext checks if the iterator has a next value to yield.
func (it *treeMapIterator[K, V]) HasNext() bool {
	return it.index < len(it.keys)
}

// Next returns the next element in the iterator it. Will panic if iterator has been exhausted.
func (it *treeMapIterator[K, V]) Next() maps.MapEntry[K, V] {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	key := it.keys[it.index]
	value := it.values[it.index]
	it.index++
	return maps.MapEntry[K, V]{Key: key, Value: value}
}

// Iterator returns an iterator for the map.
func (treeMap *TreeMap[K, V]) Iterator() maps.MapIterator[K, V] {
	keys, values := treeMap.tree.Pairs()
	it := treeMapIterator[K, V]{keys: keys, values: values, index: 0}
	return &it
}

// ContainsValue checks if the map has an entry whose value is the specified value. func equals is used to compare value for equality.
func (treeMap *TreeMap[K, V]) ContainsValue(v V, equals func(a, b V) bool) bool {
	it := treeMap.Iterator()
	for it.HasNext() {
		entry := it.Next()
		if equals(entry.Value, v) {
			return true
		}
	}
	return false
}

// Remove removes the map entry <k,V> from map m if it exists.
func (treeMap *TreeMap[K, V]) Remove(k K) (V, bool) {
	return treeMap.tree.Delete(k)
}

// RemoveAll removes all keys that are in the specified iterable from treeMap.
func (treeMap *TreeMap[K, V]) RemoveAll(keys iterator.Iterable[K]) {
	it := keys.Iterator()
	for it.HasNext() {
		treeMap.Remove(it.Next())
	}
}

// Values collects all the values of the map into a slice.
func (treeMap *TreeMap[K, V]) Values() []V {
	return treeMap.tree.Values()
}

// Keys collects the keys of the map into a slice.
func (treeMap *TreeMap[K, V]) Keys() []K {
	return treeMap.tree.Keys()
}

// Empty checks if the map is empty.
func (treeMap *TreeMap[K, V]) Empty() bool {
	return treeMap.tree.Empty()
}

// Clear removes all entries in the map.
func (treeMap *TreeMap[K, V]) Clear() {
	treeMap.tree.Clear()
}

// Equals check if map m is equal to map other. This checks that the two maps have the same entries (k,v), the values are compared
// using the specified equals function for two values. Keys are compared using their corresponding Equals method.
// Only returns true if the 2 maps are the same reference or have the same size and entries.
func (treeMap *TreeMap[K, V]) Equals(other *TreeMap[K, V], equals func(a V, b V) bool) bool {
	if treeMap == other {
		return true
	} else if treeMap.Len() != other.Len() {
		return false
	} else {
		if treeMap.Empty() && other.Empty() {
			return true
		}
		it := treeMap.Iterator()
		for it.HasNext() {
			entry := it.Next()
			v, b := other.Get(entry.Key)
			if b && equals(entry.Value, v) {
				continue
			} else {
				return false
			}
		}
		return true
	}
}

// Map applies a transformation on an entry of m i.e f((k,v)) -> (k*,v*) , using a function f and returns a new TreeMap of which its keys
// and values have been transformed.
func (treeMap *TreeMap[K, V]) Map(f func(e maps.MapEntry[K, V]) maps.MapEntry[K, V]) *TreeMap[K, V] {
	newMap := New[K, V]()
	it := treeMap.Iterator()
	for it.HasNext() {
		oldEntry := it.Next()
		newEntry := f(oldEntry)
		newMap.Put(newEntry.Key, newEntry.Value)
	}
	return newMap
}

// Filter filters the TreeMap m using a predicate function that indicates whether an entry should be kept or not in a
// TreeMap to be returned.
func (treeMap *TreeMap[K, V]) Filter(f func(e maps.MapEntry[K, V]) bool) *TreeMap[K, V] {
	newMap := New[K, V]()
	it := treeMap.Iterator()
	for it.HasNext() {
		entry := it.Next()
		if f(entry) {
			newMap.Put(entry.Key, entry.Value)
		}
	}
	return newMap
}
