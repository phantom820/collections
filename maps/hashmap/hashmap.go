// Package hashmap provides an implementation of a HashMap that uses a red black tree as an underlying container in a bucket.
package hashmap

import (
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/trees"
	"github.com/phantom820/collections/trees/rbt"
	"github.com/phantom820/collections/types"
)

const (
	loadFactorLimit = 0.75 // threshold for doubling the capacity of the hash table
	capacity        = 16   // initial size of the hash table.
)

// HashMap an implementation of a map that uses a hash table with red black tree as a container for its buckets.
type HashMap[K types.Hashable[K], V any] struct {
	defaultCapacity int
	capacity        int
	loadFactorLimit float32
	buckets         []*rbt.RedBlackTree[mapKey[K], V]
	len             int
	modifications   int // keeps track of modifications made by mutating operations, so we can have panics with concurrent modifications.
}

// New creates and returns an empty HashMap with capacity of 16 and load factor limit of 0.75 .
func New[K types.Hashable[K], V any]() *HashMap[K, V] {
	buckets := make([]*rbt.RedBlackTree[mapKey[K], V], capacity)
	hashMap := HashMap[K, V]{defaultCapacity: capacity, capacity: capacity, loadFactorLimit: loadFactorLimit, buckets: buckets, len: 0,
		modifications: 0}
	return &hashMap
}

// NewWith creates and returns an empty HashMap with the specified capacity and load factor limit.
func NewWith[K types.Hashable[K], V any](capacity int, loadFactorLimit float32) *HashMap[K, V] {
	buckets := make([]*rbt.RedBlackTree[mapKey[K], V], capacity)
	hashMap := HashMap[K, V]{defaultCapacity: capacity, capacity: capacity, loadFactorLimit: loadFactorLimit, buckets: buckets, len: 0,
		modifications: 0}
	return &hashMap
}

// modify increments the modification value.
func (hashMap *HashMap[K, V]) modify() {
	hashMap.modifications++
}

// mapKey a type to be used to represent an underlying key for a HashMap. The key and its hash value are stored to avoid recomputing the
// hash value when using this as the key for the underlying red black tree in a bucket. For internal use only.
type mapKey[K types.Hashable[K]] struct {
	key  K
	hash int
}

// Less compares two mapKeys based using their hash values.
func (mapKey mapKey[K]) Less(other mapKey[K]) bool {
	return mapKey.hash < other.hash
}

// Equals checks if 2 mapKeys are equal by using their underlying keys.
func (mapKey mapKey[K]) Equals(other mapKey[K]) bool {
	return mapKey.key.Equals(other.key)
}

// hashMapIterator a type to implement an iterator for the map.
type hashMapIterator[K types.Hashable[K], V any] struct {
	modifications    int
	getModifications func() int                                                                    // keep track of any modifications to the source if we out of parity then concurrent modification occured.
	initialized      bool                                                                          // keeps track if we have initialized the iterator or not.
	index            int                                                                           // index of the bucket we are currently in.
	maxIndex         int                                                                           // max number of buckets.
	entries          []trees.Node[mapKey[K], V]                                                    // entries in the current bucket's red black tree.
	keys             int                                                                           // a count of how many keys the iterator has seen.
	maxKeys          int                                                                           // maximum number of keys (size of the map).
	getMaxKeys       func() int                                                                    // returns the maximum number of keys we need to set.
	bucket           func(index int, maxKeys int) (*rbt.RedBlackTree[mapKey[K], V], *errors.Error) // returns a bucket at the specified index and checks for concurrent modification.
}

// HasNext checks if the iterator has a next element to yield.
func (it *hashMapIterator[K, V]) HasNext() bool {
	if !it.initialized {
		it.initialized = true
		it.modifications = it.getModifications()
		it.maxKeys = it.getMaxKeys()
	}
	if it.index < it.maxIndex && it.keys < it.maxKeys {
		return true
	}
	return false
}

// Next yields the next element in the iterator. Will panic if the iterator has no next element.
func (it *hashMapIterator[K, V]) Next() maps.MapEntry[K, V] {
	if !it.HasNext() {
		panic(errors.ErrNoNextElement())
	}
	next := func() maps.MapEntry[K, V] {
		if it.entries != nil && len(it.entries) > 0 {
			node := it.entries[0]
			key := node.Key.key
			value := node.Value
			it.keys++
			it.entries = it.entries[1:len(it.entries)]
			entry := maps.MapEntry[K, V]{Key: key, Value: value}
			return entry
		} else {
			// Find a non empty bucket and take values from it using iterator. This should be somewhat quick, O(m) collecting keys, O(m) collecting values
			// where m is the average number of items in the underlying red black tree.
			var key K
			var value V
			for i := it.index; i < it.maxIndex; i++ {
				tree, err := it.bucket(i, it.modifications)
				if err != nil {
					panic(err)
				} else if tree != nil {
					it.entries = tree.Nodes()
					it.index = i + 1
					node := it.entries[0]
					key = node.Key.key
					value = node.Value
					it.keys++
					it.entries = it.entries[1:len(it.entries)]
					break
				}
			}
			entry := maps.MapEntry[K, V]{Key: key, Value: value}
			return entry
		}
	}
	return next()
}

// Iterator returns an iterator for the map.
func (hashMap *HashMap[K, V]) Iterator() maps.MapIterator[K, V] {
	bucket := func(i int, modifications int) (*rbt.RedBlackTree[mapKey[K], V], *errors.Error) {
		if modifications != hashMap.modifications {
			return nil, errors.ErrConcurrenModification()
		} else if i < len(hashMap.buckets) {
			return hashMap.buckets[i], nil
		}
		return nil, nil
	}
	it := hashMapIterator[K, V]{index: 0, bucket: bucket,
		maxIndex: len(hashMap.buckets), keys: 0, getMaxKeys: func() int { return hashMap.len }, getModifications: func() int { return hashMap.modifications }}
	return &it
}

// resize expands the capacity of the map when we exceed the load factor limit. This creates a new map with twice the capacity of the
// old map but with the same load factor limit. Can we be more clever here ??
func (hashMap *HashMap[K, V]) resize() {
	newMap := NewWith[K, V](hashMap.capacity*2, hashMap.loadFactorLimit)
	newMap.defaultCapacity = hashMap.defaultCapacity
	newMap.PutAll(hashMap)
	*hashMap = *newMap
	newMap = nil
}

// Capacity returns the capacity (size of the hash table) of the map.
func (hashMap *HashMap[K, V]) Capacity() int {
	return hashMap.capacity
}

// Put inserts the entry <key,value> into the map. If an entry with the given key already exists then its value is updated. Returns the previous value
// associated with the key or zero value if there is no previous value.
func (hashMap *HashMap[K, V]) Put(key K, value V) V {
	hashMap.modify()
	if hashMap.LoadFactor() >= hashMap.loadFactorLimit { // If we have crossed the load factor limit resize.
		hashMap.resize()
	}
	_key := mapKey[K]{key: key, hash: key.HashCode()} // Internal key for use by underlying red black tree.
	index := _key.hash % hashMap.capacity
	if hashMap.buckets[index] == nil {
		hashMap.buckets[index] = rbt.New[mapKey[K], V]()
		hashMap.buckets[index].Insert(_key, value)
		hashMap.len++
		var zero V
		return zero
	} else {
		oldValue, ok := hashMap.buckets[index].Update(_key, value)
		if ok { // try updating otherwise make a new insertion.
			return oldValue
		}
		hashMap.len++
		hashMap.buckets[index].Insert(_key, value)
		return oldValue
	}
}

// PutIfAbsent inserts the entry <key,value> into the map if the key does not already exist in the map. Returns true if the new entry was made.
func (hashMap *HashMap[K, V]) PutIfAbsent(key K, value V) bool {
	hashMap.modify()
	_key := mapKey[K]{key: key, hash: key.HashCode()}
	index := _key.hash % hashMap.capacity
	if hashMap.buckets[index] == nil {
		hashMap.Put(key, value)
		return true
	} else if !hashMap.buckets[index].Search(_key) {
		hashMap.buckets[index].Insert(_key, value)
		hashMap.len++
		return true
	}
	return false
}

// PutAll adds all the values from another map into the map. Note this has the side effect that if a key
// is present in the map and in the passed map then the associated value in the map will be replaced by the associated value from the passed map.
func (hashMap *HashMap[K, V]) PutAll(other maps.Map[K, V]) {
	for _, key := range other.Keys() {
		value, _ := other.Get(key)
		hashMap.Put(key, value)
	}
}

// Len returns the size of the map.
func (hashMap *HashMap[K, V]) Len() int {
	return hashMap.len
}

// Get retrieves the value associated with the key in the map. Returns a value and a boolean indicating if the value is valid or invalid.
// An invalid value results when there is no entry for the given key and the zero value is returned.
func (hashMap *HashMap[K, V]) Get(key K) (V, bool) {
	_key := mapKey[K]{key: key, hash: key.HashCode()}
	index := _key.hash % hashMap.capacity
	if hashMap.buckets[index] == nil {
		var value V
		return value, false
	}
	return hashMap.buckets[index].Get(_key)
}

// ContainsKey checks if the map contains an entry with the given key.
func (hashMap *HashMap[K, V]) ContainsKey(key K) bool {
	hash := key.HashCode()
	index := hash % hashMap.capacity
	if hashMap.buckets[index] == nil {
		return false
	}
	_key := mapKey[K]{key: key, hash: hash}
	return hashMap.buckets[index].Search(_key)
}

// ContainsValue checks if the map has an entry whose value is the specified value. The function equals is used to check values for equality.
func (hashMap *HashMap[K, V]) ContainsValue(v V, equals func(a, b V) bool) bool {
	it := hashMap.Iterator()
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
func (hashMap *HashMap[K, V]) Remove(key K) (V, bool) {
	hashMap.modify()
	_key := mapKey[K]{key: key, hash: key.HashCode()}
	index := _key.hash % hashMap.capacity
	if hashMap.buckets[index] == nil {
		var value V
		return value, false
	}
	value, ok := hashMap.buckets[index].Delete(_key)
	if ok {
		hashMap.len--
		if hashMap.buckets[index].Empty() {
			hashMap.buckets[index] = nil
		}
	}
	return value, ok
}

// RemoveAll removes all key entries from the map that appear in the iterable keys.
func (hashMap *HashMap[K, V]) RemoveAll(keys iterator.Iterable[K]) {
	it := keys.Iterator()
	for it.HasNext() {
		hashMap.Remove(it.Next())
	}
}

// LoadFactor returns the load factor of the map.
func (hashMap *HashMap[K, V]) LoadFactor() float32 {
	return float32(hashMap.len) / float32(hashMap.capacity)
}

// Values returns a slice containing all the values in the map.
func (hashMap *HashMap[K, V]) Values() []V {
	data := make([]V, 0)
	i := 0
	for _, bucket := range hashMap.buckets {
		if bucket == nil {
			continue
		}
		for _, value := range bucket.Values() { // This should not be costly since the buckets should be small in size O(n*k) n -> capacity
			data = append(data, value) // k -> average size of a bucket
			i = i + 1
		}
	}
	return data
}

// Keys returns a slice containing all the keys in the map.
func (hashMap *HashMap[K, V]) Keys() []K {
	data := make([]K, hashMap.len)
	i := 0
	for _, bucket := range hashMap.buckets {
		if bucket == nil {
			continue
		}
		for _, key := range bucket.Keys() { // This should not be costly since the buckets should be small rbt tree O(n*k) n -> capacity
			data[i] = key.key // k -> average size of a bucket.
			i = i + 1
		}
	}
	return data
}

// Empty checks if the map is empty.
func (hashMap *HashMap[K, V]) Empty() bool {
	return hashMap.len == 0
}

// Clear removes all entries from the map.
func (hashMap *HashMap[K, V]) Clear() {
	hashMap.modify()
	hashMap.len = 0
	hashMap.buckets = nil
	hashMap.capacity = hashMap.defaultCapacity
	hashMap.buckets = make([]*rbt.RedBlackTree[mapKey[K], V], hashMap.defaultCapacity)
}

// Equals checks if the map is equal to map other. This checks that the two maps have the same entries (k,v), the values are compared
// using the specified equals function. Keys are compared using their corresponding Equals method.
// Only returns true if the 2 maps are the same reference or have the same size and entries.
func (hashMap *HashMap[K, V]) Equals(other *HashMap[K, V], equals func(a V, b V) bool) bool {
	if hashMap == other {
		return true
	} else if hashMap.len != other.Len() {
		return false
	} else {
		if hashMap.Empty() && other.Empty() {
			return true
		}
		iterator := hashMap.Iterator()
		for iterator.HasNext() {
			entry := iterator.Next()
			value, ok := other.Get(entry.Key)
			if ok && equals(entry.Value, value) {
				continue
			} else {
				return false
			}
		}
		return true
	}
}

// Map applies a transformation on an entry of the map i.e f((k,v)) -> (k*,v*) , using the function f and returns a new map with the
// transformed entries.
func (hashMap *HashMap[K, V]) Map(f func(entry maps.MapEntry[K, V]) maps.MapEntry[K, V]) *HashMap[K, V] {
	newMap := New[K, V]()
	it := hashMap.Iterator()
	for it.HasNext() {
		oldEntry := it.Next()
		newEntry := f(oldEntry)
		newMap.Put(newEntry.Key, newEntry.Value)
	}
	return newMap
}

// Filter filters the map using the predicate function  f and returns a new map containing only entries that satisfy the predicate.
func (hashMap *HashMap[K, V]) Filter(f func(entry maps.MapEntry[K, V]) bool) *HashMap[K, V] {
	newMap := New[K, V]()
	iterator := hashMap.Iterator()
	for iterator.HasNext() {
		entry := iterator.Next()
		if f(entry) {
			newMap.Put(entry.Key, entry.Value)
		}
	}
	return newMap
}
