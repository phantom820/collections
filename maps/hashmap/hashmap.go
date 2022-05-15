// package hashmap provides an implementation of a HashMap that uses a Red Black Tree as an underlying container in a bucket.
package hashmap

import (
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/trees/rbt"
	"github.com/phantom820/collections/types"
)

const (
	loadFactorLimit = 0.75
	Capacity        = 16
)

// HashMap an implementation of a hashmap that uses a red black tree as a container for its individual buckets.
type HashMap[K types.Hashable[K], V any] struct {
	capacity        int
	loadFactorLimit float32
	buckets         []*rbt.RedBlackTree[mapKey[K], V]
	len             int
}

// New creates an empty HashMap with default initial capacity of 16 and load factor limit of 0.75.
func New[K types.Hashable[K], V any]() *HashMap[K, V] {
	buckets := make([]*rbt.RedBlackTree[mapKey[K], V], Capacity)
	m := HashMap[K, V]{capacity: Capacity, loadFactorLimit: loadFactorLimit, buckets: buckets, len: 0}
	return &m
}

// NewWith creates an empty HashMap with the specified capacity and load factor limit.
func NewWith[K types.Hashable[K], V any](capacity int, loadFactorLimit float32) *HashMap[K, V] {
	buckets := make([]*rbt.RedBlackTree[mapKey[K], V], capacity)
	hashMap := HashMap[K, V]{capacity: capacity, loadFactorLimit: loadFactorLimit, buckets: buckets, len: 0}
	return &hashMap
}

// mapKey a struct used to represent an underlying key for a HashMap. The actual supplied key and its hash value so that it can
// be used in a red black tree that needs values to compare keys for operations.
type mapKey[K types.Hashable[K]] struct {
	key  K
	hash int
}

// Less how Less operates for a key. It simply compares underlying hash value of the keys.
func (mapKey mapKey[K]) Less(other mapKey[K]) bool {
	return mapKey.hash < other.hash
}

// Equals checks if 2 keys are equal by using their supllied equals method.
func (mapKey mapKey[K]) Equals(other mapKey[K]) bool {
	return mapKey.key.Equals(other.key)
}

// hashMapIterator an iterator for moving through the keys and value of a HashMap.
type hashMapIterator[K types.Hashable[K], V any] struct {
	index      int
	maxIndex   int
	exhausted  bool // check if the iterator has been used up (parsed all values)
	bucket     []mapKey[K]
	values     []V
	keys       int
	maxkeys    int
	nextBucket func(i int) *rbt.RedBlackTree[mapKey[K], V]
}

// Cycle resets the iterator.
func (iterator *hashMapIterator[K, V]) Cycle() {
	iterator.index = 0
	iterator.bucket = nil
	iterator.keys = 0
}

// HasNext checks if the iterator has a next element to yield.
func (iterator *hashMapIterator[K, V]) HasNext() bool {
	if iterator.index < iterator.maxIndex && iterator.keys < iterator.maxkeys {
		return true
	}
	iterator.exhausted = true
	return false
}

// Next returns the next element in the iterator it. Will panic if iterator has no next element.
func (it *hashMapIterator[K, V]) Next() maps.MapEntry[K, V] {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	next := func() maps.MapEntry[K, V] {

		if it.bucket != nil && len(it.bucket) > 0 {
			k := it.bucket[0].key
			v := it.values[0]
			it.keys++
			it.bucket = it.bucket[1:len(it.bucket)]
			it.values = it.values[1:len(it.values)]
			entry := maps.MapEntry[K, V]{Key: k, Value: v}
			return entry
		} else {
			// find a non empty bucket and take values from it. This should be somewhat quick, O(m) collecting keys, O(m) collecting values
			// Where m is the number of items in the rbt.
			var k K
			var v V
			for i := it.index; i < it.maxIndex; i++ {
				bucket := it.nextBucket(i)
				if bucket != nil {
					it.bucket = bucket.Keys()
					it.values = bucket.Values()
					it.index = i + 1
					k = it.bucket[0].key
					v = it.values[0]
					it.keys++
					it.bucket = it.bucket[1:len(it.bucket)]
					it.values = it.values[1:len(it.values)]
					break
				}
			}

			entry := maps.MapEntry[K, V]{Key: k, Value: v}
			return entry
		}
	}
	return next()
}

// Iterator returns an iterator for the map.
func (hashMap *HashMap[K, V]) Iterator() maps.MapIterator[K, V] {
	nextBucket := func(i int) *rbt.RedBlackTree[mapKey[K], V] {
		if i < len(hashMap.buckets) {
			return hashMap.buckets[i]
		}
		return nil
	}
	iterator := hashMapIterator[K, V]{index: 0, nextBucket: nextBucket,
		maxIndex: len(hashMap.buckets), maxkeys: hashMap.len, keys: 0}
	return &iterator
}

// resize expands the capacity of the map when we exceed the load factor limit. This creates a new map with twice the capacity of the
// old map but with the same load factor limit. Can we be more clever here ??
func (hashMap *HashMap[K, V]) resize() {
	newMap := NewWith[K, V](hashMap.capacity*2, hashMap.loadFactorLimit)
	newMap.PutAll(hashMap)
	*hashMap = *newMap
	newMap = nil
}

// Capacity retrieves the capacity of the map i.e number of buckets.
func (hashMap *HashMap[K, V]) Capacity() int {
	return hashMap.capacity
}

// Put associates the specified value with the specified key in the map. If the key already exists then its value will be updated. It
// returns the old value associated with the key or zero value if no previous association.
func (hashMap *HashMap[K, V]) Put(key K, value V) V {
	if hashMap.LoadFactor() >= hashMap.loadFactorLimit { // if we have crossed the load factor limit resize.
		hashMap.resize()
	}
	_key := mapKey[K]{key: key, hash: key.HashCode()} // internal key for use by undelring container.
	index := _key.hash % hashMap.capacity
	if hashMap.buckets[index] == nil {
		hashMap.buckets[index] = rbt.New[mapKey[K], V]()
		hashMap.buckets[index].Insert(_key, value)
		hashMap.len++
		var e V
		return e
	} else {
		old, ok := hashMap.buckets[index].Update(_key, value)
		if ok { // try updating otherwise make a new insertion.
			return old
		}
		hashMap.len++
		hashMap.buckets[index].Insert(_key, value)
		return old
	}
}

// PutIfAbsent adds the value with the specified key to the map only if the key has not been mapped already.
func (hashMap *HashMap[K, V]) PutIfAbsent(key K, value V) bool {
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

// PutAll adds all the values from other map into the map. Note this has the side effect that if a key
// is present in the map and in other map then the associated value  in m will be replaced by the associated value  in other.
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

// Get retrieves the value associated with the key in the map hashMap. If there is no such value the zero value is returned along with false.
func (hashMap *HashMap[K, V]) Get(key K) (V, bool) {
	_key := mapKey[K]{key: key, hash: key.HashCode()}
	index := _key.hash % hashMap.capacity
	if hashMap.buckets[index] == nil {
		var value V
		return value, false
	}
	return hashMap.buckets[index].Get(_key)
}

// ContainsKey checks if the map contains a mapping for the key.
func (hashMap *HashMap[K, V]) ContainsKey(key K) bool {
	hash := key.HashCode()
	i := hash % hashMap.capacity
	if hashMap.buckets[i] == nil {
		return false
	}
	_key := mapKey[K]{key: key, hash: hash}
	return hashMap.buckets[i].Search(_key)

}

// ContainsValue checks if the map has an entry whose value is the specified value. func equals is used to compare value for equality.
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

// Remove removes the map entry <k,V> from map if it exists.
func (hashMap *HashMap[K, V]) Remove(key K) (V, bool) {
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

// RemoveAll removes all keys that are in the specified iterable from the map.
func (hashMap *HashMap[K, V]) RemoveAll(keys iterator.Iterable[K]) {
	it := keys.Iterator()
	for it.HasNext() {
		hashMap.Remove(it.Next())
	}
}

// LoadFactor computes the load factor of the map.
func (hashMap *HashMap[K, V]) LoadFactor() float32 {
	return float32(hashMap.len) / float32(hashMap.capacity)
}

// Values collects all the values of the map into a slice.
func (hashMap *HashMap[K, V]) Values() []V {
	data := make([]V, 0)
	i := 0
	for _, bucket := range hashMap.buckets {
		if bucket == nil {
			continue
		}
		for _, v := range bucket.Values() { // This should not be costly since the buckets should be small rbt tree O(n*k) n -> capacity
			data = append(data, v) // k -> average size of a bucket
			i = i + 1
		}
	}
	return data
}

// Keys collects the keys of the map into a slice.
func (hashMap *HashMap[K, V]) Keys() []K {
	data := make([]K, hashMap.len)
	i := 0
	for _, bucket := range hashMap.buckets {
		if bucket == nil {
			continue
		}
		for _, k := range bucket.Keys() { // This should not be costly since the buckets should be small rbt tree O(n*k) n -> capacity
			data[i] = k.key // k -> average size of a bucket.
			i = i + 1
		}
	}
	return data
}

// Empty checks if the map is empty.
func (hashMap *HashMap[K, V]) Empty() bool {
	return hashMap.len == 0
}

// Clear removes all entries in the map.
func (hashMap *HashMap[K, V]) Clear() {
	hashMap.len = 0
	hashMap.buckets = nil
	hashMap.buckets = make([]*rbt.RedBlackTree[mapKey[K], V], hashMap.capacity)
}

// Equals checks if the map is equal to map other. This checks that the two maps have the same entries (k,v), the values are compared
// using the specified equals function for two values. Keys are compared using their corresponding Equals method.
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

// Map applies a transformation on an entry of m i.e f((k,v)) -> (k*,v*) , using a function f and returns a new HashMap of which its keys
// and values have been transformed.
func (hashMap *HashMap[K, V]) Map(f func(e maps.MapEntry[K, V]) maps.MapEntry[K, V]) *HashMap[K, V] {
	newMap := New[K, V]()
	iterator := hashMap.Iterator()
	for iterator.HasNext() {
		oldEntry := iterator.Next()
		newEntry := f(oldEntry)
		newMap.Put(newEntry.Key, newEntry.Value)
	}
	return newMap
}

// Filter filters the HashMap m using a predicate function that indicates whether an entry should be kept or not in a
// HashMap to be returned.
func (hashMap *HashMap[K, V]) Filter(f func(e maps.MapEntry[K, V]) bool) *HashMap[K, V] {
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
