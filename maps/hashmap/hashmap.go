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

// HashMap an implementation of a hashmap that uses a red black tree as a container in its individual buckets.
type HashMap[K types.Hashable[K], V any] struct {
	capacity        int
	loadFactorLimit float32
	buckets         []*rbt.RedBlackTree[key[K], V]
	len             int
}

// New creates an empty HashMap with default initial capacity of 16 and load factor limit of 0.75.
func New[K types.Hashable[K], V any]() *HashMap[K, V] {
	buckets := make([]*rbt.RedBlackTree[key[K], V], Capacity)
	m := HashMap[K, V]{capacity: Capacity, loadFactorLimit: loadFactorLimit, buckets: buckets, len: 0}
	return &m
}

// NewHashMapWith creates an empty HashMap with the specified capacity and load factor limit.
func NewHashMapWith[K types.Hashable[K], V any](capacity int, loadFactorLimit float32) *HashMap[K, V] {
	buckets := make([]*rbt.RedBlackTree[key[K], V], capacity)
	m := HashMap[K, V]{capacity: capacity, loadFactorLimit: loadFactorLimit, buckets: buckets, len: 0}
	return &m
}

// key a struct used to represent an underlying key for a HashMap. The actual supplied key and its hash value so that it can
// be used in a red black tree that needs values to compare keys for operations.
type key[K types.Hashable[K]] struct {
	key  K
	hash int
}

// Less how Less operates for a key. It simply compares underlying hash value of the keys.
func (k key[K]) Less(other key[K]) bool {
	return k.hash < other.hash
}

// Equals checks if 2 keys are equal by using their supllied equals method.
func (k key[K]) Equals(other key[K]) bool {
	return k.key.Equals(other.key)
}

// hashMapIterator an iterator for moving through the keys and value of a HashMap.
type hashMapIterator[K types.Hashable[K], V any] struct {
	index      int
	maxIndex   int
	exhausted  bool // check if the iterator has been used up (parsed all values)
	bucket     []key[K]
	values     []V
	keys       int
	maxkeys    int
	nextBucket func(i int) *rbt.RedBlackTree[key[K], V]
}

// Cycle resets the iterator.
func (it *hashMapIterator[K, V]) Cycle() {
	it.index = 0
	it.bucket = nil
	it.keys = 0
}

// HasNext checks if the iterator has a next value to yield.
func (it *hashMapIterator[K, V]) HasNext() bool {
	if it.index < it.maxIndex && it.keys < it.maxkeys {
		return true
	}
	it.exhausted = true
	return false
}

// Next returns the next element in the iterator it. Will panic if iterator has been exhausted.
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
func (m *HashMap[K, V]) Iterator() maps.MapIterator[K, V] {
	nextBucket := func(i int) *rbt.RedBlackTree[key[K], V] {
		if i < len(m.buckets) {
			return m.buckets[i]
		}
		return nil
	}
	it := hashMapIterator[K, V]{index: 0, nextBucket: nextBucket,
		maxIndex: len(m.buckets), maxkeys: m.len, keys: 0}
	return &it
}

// resize expands the capacity of the map when we exceed the load factor limit. This creates a new map with twice the capacity of the
// old map but with the same load factor limit. Can we be more clever here ??
func (m *HashMap[K, V]) resize() {
	newMap := NewHashMapWith[K, V](m.capacity*2, m.loadFactorLimit)
	newMap.PutAll(m)
	*m = *newMap
	newMap = nil
}

// Capacity retrieves the capacity of the map i.e number of buckets.
func (m *HashMap[K, V]) Capacity() int {
	return m.capacity
}

// Put associates the specified value with the specified key in the map. If the key already exists then its value will be updated. It
// returns the old value associated with the key or zero value if no previous association.
func (m *HashMap[K, V]) Put(k K, v V) V {
	if m.LoadFactor() >= m.loadFactorLimit { // if we have crossed the load factor limit resize.
		m.resize()
	}
	_key := key[K]{key: k, hash: k.HashCode()} // internal key for use by undelring container.
	index := _key.hash % m.capacity
	if m.buckets[index] == nil {
		m.buckets[index] = rbt.New[key[K], V]()
		m.buckets[index].Insert(_key, v)
		m.len++
		var e V
		return e
	} else {
		old, p := m.buckets[index].Update(_key, v)
		if p { // try updating otherwise make a new insertion.
			return old
		}
		m.len++
		m.buckets[index].Insert(_key, v)
		return old
	}
}

// PutIfAbsent adds the value with the specified key to the map only if the key has not been mapped already.
func (m *HashMap[K, V]) PutIfAbsent(k K, v V) bool {
	_key := key[K]{key: k, hash: k.HashCode()}
	index := _key.hash % m.capacity
	if m.buckets[index] == nil {
		m.Put(k, v)
		return true
	} else if !m.buckets[index].Search(_key) {
		m.buckets[index].Insert(_key, v)
		m.len++
		return true
	}
	return false
}

// PutAll adds all the values from other map into the map. Note this has the side effect that if a key
// is present in the map and in other map then the associated value  in m will be replaced by the associated value  in other.
func (m *HashMap[K, V]) PutAll(other maps.Map[K, V]) {
	for _, k := range other.Keys() {
		v, _ := other.Get(k)
		m.Put(k, v)
	}
}

// Len returns the size of the map.
func (m *HashMap[K, V]) Len() int {
	return m.len
}

// Get retrieves the value associated with key in the map m. If there is no such value the zero value is returned along with false.
func (m *HashMap[K, V]) Get(k K) (V, bool) {
	_key := key[K]{key: k, hash: k.HashCode()}
	index := _key.hash % m.capacity
	if m.buckets[index] == nil {
		var e V
		return e, false
	}
	return m.buckets[index].Get(_key)
}

// ContainsKey checks if the map contains a mapping for the key.
func (m *HashMap[K, V]) ContainsKey(k K) bool {
	hash := k.HashCode()
	i := hash % m.capacity
	if m.buckets[i] == nil {
		return false
	}
	_key := key[K]{key: k, hash: hash}
	return m.buckets[i].Search(_key)

}

// ContainsValue checks if the map has an entry whose value is the specified value. func equals is used to compare value for equality.
func (m *HashMap[K, V]) ContainsValue(v V, equals func(a, b V) bool) bool {
	it := m.Iterator()
	for it.HasNext() {
		entry := it.Next()
		if equals(entry.Value, v) {
			return true
		}
	}
	return false
}

// Remove removes the map entry <k,V> from map m if it exists.
func (m *HashMap[K, V]) Remove(k K) (V, bool) {
	_key := key[K]{key: k, hash: k.HashCode()}
	index := _key.hash % m.capacity
	if m.buckets[index] == nil {
		var e V
		return e, false
	}
	v, r := m.buckets[index].Delete(_key)
	if r {
		m.len--
		if m.buckets[index].Empty() {
			m.buckets[index] = nil
		}
	}
	return v, r
}

// RemoveAll removes all keys that are in the specified iterable from m.
func (m *HashMap[K, V]) RemoveAll(keys iterator.Iterable[K]) {
	it := keys.Iterator()
	for it.HasNext() {
		m.Remove(it.Next())
	}
}

// LoadFactor computes the load factor of the map m.
func (m *HashMap[K, V]) LoadFactor() float32 {
	return float32(m.len) / float32(m.capacity)
}

// Values collects all the values of the map into a slice.
func (m *HashMap[K, V]) Values() []V {
	data := make([]V, 0)
	i := 0
	for _, bucket := range m.buckets {
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
func (m *HashMap[K, V]) Keys() []K {
	data := make([]K, m.len)
	i := 0
	for _, bucket := range m.buckets {
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
func (m *HashMap[K, V]) Empty() bool {
	return m.len == 0
}

// Clear removes all entries in the map.
func (m *HashMap[K, V]) Clear() {
	m.len = 0
	m.buckets = nil
	m.buckets = make([]*rbt.RedBlackTree[key[K], V], m.capacity)
}

// Equals check if map m is equal to map other. This checks that the two maps have the same entries (k,v), the values are compared
// using the specified equals function for two values. Keys are compared using their corresponding Equals method.
// Only returns true if the 2 maps are the same reference or have the same size and entries.
func (m *HashMap[K, V]) Equals(other *HashMap[K, V], equals func(a V, b V) bool) bool {
	if m == other {
		return true
	} else if m.len != other.Len() {
		return false
	} else {
		if m.Empty() && other.Empty() {
			return true
		}
		it := m.Iterator()
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

// Map applies a transformation on an entry of m i.e f((k,v)) -> (k*,v*) , using a function f and returns a new HashMap of which its keys
// and values have been transformed.
func (m HashMap[K, V]) Map(f func(e maps.MapEntry[K, V]) maps.MapEntry[K, V]) *HashMap[K, V] {
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
func (m HashMap[K, V]) Filter(f func(e maps.MapEntry[K, V]) bool) *HashMap[K, V] {
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
