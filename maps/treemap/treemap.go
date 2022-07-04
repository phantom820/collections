// Package treemap provides an implementation of a TreeMap which is a map that stores entries in a sorted order.
package treemap

import (
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterator"

	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/trees"
	"github.com/phantom820/collections/trees/rbt"
	"github.com/phantom820/collections/types"
)

// TreeMap an implementation of a map in which entries are stored according to their defined ordering.
type TreeMap[K types.Comparable[K], V any] struct {
	tree          *rbt.RedBlackTree[K, V]
	modifications int
}

// New creates and returns an empty TreeMap.
func New[K types.Comparable[K], V any]() *TreeMap[K, V] {
	treeMap := TreeMap[K, V]{tree: rbt.New[K, V]()}
	return &treeMap
}

// modify increments the modification value.
func (treeMap *TreeMap[K, V]) modify() {
	treeMap.modifications++
}

// Put inserts the entry <key,value> into the map. If an entry with the given key already exists then its value is updated. Returns the previous value
// associated with the key or zero value if there is no previous value.
func (treeMap *TreeMap[K, V]) Put(k K, v V) V {
	treeMap.modify()
	if val, ok := treeMap.tree.Get(k); ok {
		treeMap.tree.Update(k, v)
		return val
	}
	treeMap.tree.Insert(k, v)
	var val V
	return val
}

// PutIfAbsent inserts the entry <key,value> into the map if the key does not already exist in the map. Returns true if the new entry was made.
func (treeMap *TreeMap[K, V]) PutIfAbsent(k K, v V) bool {
	treeMap.modify()
	if _, ok := treeMap.tree.Get(k); ok {
		return false
	}
	treeMap.tree.Insert(k, v)
	return true
}

// PutAll adds all the values from another map into the map. Note this has the side effect that if a key
// is present in the map and in the passed map then the associated value in the map will be replaced by the associated value from the passed map.
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

// Get retrieves the value associated with the key in the map. Returns a value and a boolean indicating if the value is valid or invalid.
// An invalid value results when there is no entry for the given key and the zero value is returned.
func (treeMap *TreeMap[K, V]) Get(k K) (V, bool) {
	return treeMap.tree.Get(k)
}

// treeMapIterator a type to implement an iterator for the map.
type treeMapIterator[K types.Comparable[K], V any] struct {
	initialized      bool
	index            int
	nodes            []trees.Node[K, V]
	getNodes         func() []trees.Node[K, V]
	modifications    int
	getModifications func() int
}

// HasNext checks if the iterator has a next element to yield.
func (it *treeMapIterator[K, V]) HasNext() bool {
	if !it.initialized {
		it.initialized = true
		it.nodes = it.getNodes()
		it.modifications = it.getModifications()
	} else if it.modifications != it.getModifications() {
		panic(errors.ErrConcurrenModification())
	}
	return it.index < len(it.nodes)
}

// Next yields the next element in the iterator. Will panic if the iterator has no next element.
func (it *treeMapIterator[K, V]) Next() maps.MapEntry[K, V] {
	if !it.HasNext() {
		panic(errors.ErrNoNextElement())
	}
	node := it.nodes[it.index]
	it.index++
	return maps.MapEntry[K, V]{Key: node.Key, Value: node.Value}
}

// Iterator returns an iterator for the map.
func (treeMap *TreeMap[K, V]) Iterator() maps.MapIterator[K, V] {
	it := treeMapIterator[K, V]{nodes: []trees.Node[K, V]{}, index: 0, getNodes: func() []trees.Node[K, V] { return treeMap.tree.Nodes() },
		getModifications: func() int { return treeMap.modifications }}
	return &it
}

// ContainsKey checks if the map contains an entry with the given key.
func (treeMap *TreeMap[K, V]) ContainsKey(k K) bool {
	return treeMap.tree.Search(k)
}

// ContainsValue checks if the map has an entry whose value is the specified value. The function equals is used to check values for equality.
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

// Remove removes the map entry <key,value> from the map if it exists. Returns the previous value associated with the key and a boolean indicating if the returned
// values is valid or invalid. An invalid value results when there is no entry in the map associated with the given key.
func (treeMap *TreeMap[K, V]) Remove(k K) (V, bool) {
	treeMap.modify()
	return treeMap.tree.Delete(k)
}

// RemoveAll removes all key entries from the map that appear in the iterable keys.
func (treeMap *TreeMap[K, V]) RemoveAll(keys iterator.Iterable[K]) {
	it := keys.Iterator()
	for it.HasNext() {
		treeMap.Remove(it.Next())
	}
}

// Values returns a slice containing all the values in the map.
func (treeMap *TreeMap[K, V]) Values() []V {
	return treeMap.tree.Values()
}

// Keys returns a slice containing all the keys in the map.
func (treeMap *TreeMap[K, V]) Keys() []K {
	return treeMap.tree.Keys()
}

// Empty checks if the map is empty.
func (treeMap *TreeMap[K, V]) Empty() bool {
	return treeMap.tree.Empty()
}

// Clear removes all entries from the map.
func (treeMap *TreeMap[K, V]) Clear() {
	treeMap.modify()
	treeMap.tree.Clear()
}

// Equals checks if the map is equal to map other. This checks that the two maps have the same entries (k,v), the values are compared
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

// LeftSubMap returns a new map which is a subset of the original map containing keys less than or equals the specified key. If inclusive is true
// a key equals to the specified key will be included otherwise excluded.
func (treeMap *TreeMap[K, V]) LeftSubMap(key K, inclusive bool) *TreeMap[K, V] {
	leftSubTree := treeMap.tree.LeftSubTree(key, inclusive)
	leftSubMap := TreeMap[K, V]{tree: leftSubTree}
	return &leftSubMap
}

// RightSubMap returns a new map which is a subset of the original map containing keys greater than or equals the specified key. If inclusive is true
// a key equals to the specified key will be included otherwise excluded.
func (treeMap *TreeMap[K, V]) RightSubMap(key K, inclusive bool) *TreeMap[K, V] {
	rightSubTree := treeMap.tree.RightSubTree(key, inclusive)
	rightSubMap := TreeMap[K, V]{tree: rightSubTree}
	return &rightSubMap
}

func (treeMap *TreeMap[K, V]) SubMap(fromKey K, fromInclusive bool, toKey K, toInclusive bool) *TreeMap[K, V] {
	if toKey.Less(fromKey) && !toKey.Equals(fromKey) {
		panic(errors.ErrMapKeyRange(fromKey, toKey))
	}
	subTree := treeMap.tree.SubTree(fromKey, fromInclusive, toKey, toInclusive)
	subMap := TreeMap[K, V]{tree: subTree}
	return &subMap
}

// Map applies a transformation on an entry of the map i.e f((k,v)) -> (k*,v*) , using the function f and returns a new map with the
// transformed entries.
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

// Filter filters the map using the predicate function  f and returns a new map containing only entries that satisfy the predicate.
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
