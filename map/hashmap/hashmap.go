package hashmap

import (
	"collections/interfaces"
	_map "collections/map"
	"collections/tree/rbt"
	"errors"
)

var (
	NoNextElementError = errors.New("Iterator has no next element.")
)

const (
	LoadFactorLimit = 0.75
	Capacity        = 16
)

// key a struct used to represent an underlying key for a HashMap. The actual supplied key and its hash value so that it can
// be used in a red black tree that needs values to compare for operations.
type key[K interfaces.Hashable[K]] struct {
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

// HashMap a hashmap with allowed keys of type K that are associated with values of type V.
type HashMap[K interfaces.Hashable[K], V any] interface {
	_map.Map[K, V] // See this for available methods.
	interfaces.Functional[_map.MapEntry[K, V], HashMap[K, V]]
	Equals(other HashMap[K, V], equals func(a V, b V) bool) bool
}

// hashMap underlying concrete implementation for a Hashmap.
// capacity -> initial number of buckets, loadFactorLimit -> dictates when should expansion take place.
// buckets -> slice with actual underlying containers (Red Black Trees) in this case. len -> no keys stored.
type hashMap[K interfaces.Hashable[K], V any] struct {
	capacity        int
	loadFactorLimit float32
	buckets         []rbt.RedBlackTree[key[K], V]
	len             int
}

// NewHashMap creates a new empty HashMap with default initial capacity and load factor limit.
func NewHashMap[K interfaces.Hashable[K], V any]() HashMap[K, V] {
	buckets := make([]rbt.RedBlackTree[key[K], V], Capacity)
	m := hashMap[K, V]{capacity: Capacity, loadFactorLimit: LoadFactorLimit, buckets: buckets, len: 0}
	return &m
}

// NewHashMapWith creates a new empty HashMap with the specified capacity and load factor limit.
func NewHashMapWith[K interfaces.Hashable[K], V any](capacity int, loadFactorLimit float32) HashMap[K, V] {
	buckets := make([]rbt.RedBlackTree[key[K], V], capacity)
	m := hashMap[K, V]{capacity: capacity, loadFactorLimit: loadFactorLimit, buckets: buckets, len: 0}
	return &m
}

// hashMapIterator an iterator for moving through the keys and value of a HashMap.
type hashMapIterator[K interfaces.Hashable[K], V any] struct {
	index      int
	maxIndex   int
	exhausted  bool // check if the iterator has been used up (parsed all values)
	bucket     []key[K]
	values     []V
	keys       int
	maxkeys    int
	nextBucket func(i int) rbt.RedBlackTree[key[K], V]
}

// Cycle resets the iterator.
func (it *hashMapIterator[K, V]) Cycle() {
	it.index = 0
	it.bucket = nil
	it.keys = 0
}

// HasNext checks if there is next value to be produced by iterator.
func (it *hashMapIterator[K, V]) HasNext() bool {
	if it.index < it.maxIndex && it.keys < it.maxkeys {
		return true
	}
	it.exhausted = true
	return false
}

// Next gives out the next element for a map iterator.  Start from bucket 0 and grab values from underlying container "emitting" them
// until the container is exhausted, then move on to next container. Now what happens when someone calls next on an exhausted iterator ??
// they should get the zero value up until the iterator is cycled back. Calling Next in the worst case scenarion is O(m) where m is
// the average number of items in a bucket.
func (it *hashMapIterator[K, V]) Next() _map.MapEntry[K, V] {
	if !it.HasNext() {
		panic(NoNextElementError)
	}
	next := func() _map.MapEntry[K, V] {

		if it.bucket != nil && len(it.bucket) > 0 {
			k := it.bucket[0].key
			v := it.values[0]
			it.keys++
			it.bucket = it.bucket[1:len(it.bucket)]
			it.values = it.values[1:len(it.values)]
			entry := _map.NewMapEntry(k, v)
			return entry
		} else {
			// find a non empty bucket and take values from it. This should be somewhat quick, O(m) collecting keys, O(m) collecting values
			// Where m is the number of items in the tree.
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

			entry := _map.NewMapEntry(k, v)
			return entry
		}
	}
	return next()
}

// Iterator returns an iterator on the keys of the map m.
func (m *hashMap[K, V]) Iterator() _map.MapIterator[K, V] {
	nextBucket := func(i int) rbt.RedBlackTree[key[K], V] {
		if i < len(m.buckets) {
			return m.buckets[i]
		}
		return nil
	}
	it := hashMapIterator[K, V]{index: 0, nextBucket: nextBucket,
		maxIndex: len(m.buckets), maxkeys: m.len, keys: 0}
	return &it
}

// resize expnds the capacity of the map m when we exceed the load factor limit. This creates a new map with twice the capacity of the
// old map but with the same load factor limit. Can we be more clever here ??
func (m *hashMap[K, V]) resize() {
	newMap := NewHashMapWith[K, V](m.capacity*2, m.loadFactorLimit)
	newMap.PutAll(m)
	newMapPtr, ok := newMap.(*hashMap[K, V])
	if ok {
		*m = *newMapPtr
	}
}

// Capacity retrieves the capacity of the map i.e number of buckets in m.
func (m *hashMap[K, V]) Capacity() int {
	return m.capacity
}

// Put associates the specified value with the specified key in m and. If the key alreaady exists then its value will be updated. It
// returns the old value associated with the key or zero value if no previous association.
func (m *hashMap[K, V]) Put(k K, v V) V {
	if m.LoadFactor() >= m.loadFactorLimit { // if we have crossed the load factor limit resize.
		m.resize()
	}

	_key := key[K]{key: k, hash: k.HashCode()} // internal key for use by undelring container.
	index := _key.hash % m.capacity
	if m.buckets[index] == nil {
		m.buckets[index] = rbt.NewRedBlackTree[key[K], V]()
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

// PutIfAbsent adds the value v associate with key k to map m only if k has not been mapped already.
func (m *hashMap[K, V]) PutIfAbsent(k K, v V) bool {
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

// PutAll adds all the values from other map into the map m. Note this has the side effect that if a key
// is present in m and other then the associated value  in m will be replaced by the associated value  in other.
func (m *hashMap[K, V]) PutAll(other _map.Map[K, V]) {
	for _, k := range other.Keys() {
		v, _ := other.Get(k)
		m.Put(k, v)
	}
}

// Len returns the size (number of keys currently mapped) of the map m.
func (m *hashMap[K, V]) Len() int {
	return m.len
}

// Get retrieves the value associated with key k in map m. When there is no such value the zero value is returned along with false
// indicating value was not found.
func (m *hashMap[K, V]) Get(k K) (V, bool) {
	_key := key[K]{key: k, hash: k.HashCode()}
	index := _key.hash % m.capacity
	if m.buckets[index] == nil {
		var e V
		return e, false
	}
	return m.buckets[index].Get(_key)
}

// Contains checks if the map m contains a mapping for the key k.
func (m *hashMap[K, V]) Contains(k K) bool {
	i := k.HashCode() % m.capacity
	if m.buckets[i] != nil {
		return true
	}
	return false
}

// Remove removes the map entry <k,V> from map m if there is a value associated with the key k.
func (m *hashMap[K, V]) Remove(k K) bool {
	_key := key[K]{key: k, hash: k.HashCode()}
	index := _key.hash % m.capacity
	if m.buckets[index] == nil {
		return false
	}
	r := m.buckets[index].Delete(_key)
	if r {
		m.len--
		if m.buckets[index].Empty() {
			m.buckets[index] = nil
		}
	}
	return r
}

// RemoveAll removes all keys that in the specified iterable from m. This could be potentially slow (iterator perfomance) ?? just remove things from outside ?
func (m *hashMap[K, V]) RemoveAll(keys interfaces.Iterable[K]) {
	it := keys.Iterator()
	for it.HasNext() {
		m.Remove(it.Next())
	}
}

// loadFactor computes the load factor of the map m.
func (m *hashMap[K, V]) LoadFactor() float32 {
	return float32(m.len) / float32(m.capacity)
}

// Values collects all the values of the map m into a slice.
func (m *hashMap[K, V]) Values() []V {
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

// Keys collects the keys of the map m into a slice.
func (m *hashMap[K, V]) Keys() []K {
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

// Empty checks if the map m is empty.
func (m *hashMap[K, V]) Empty() bool {
	return m.len == 0
}

// Clear removes all entries in the map m.
func (m *hashMap[K, V]) Clear() {
	keys := m.Keys()
	for _, key := range keys {
		m.Remove(key)
	}
}

// Equals check if map m is equal to map n. This checks that the two maps have the same entries (k,v), the values are compared
// using the specified equals function for tow values. Keys are compared using their corresponding Equals method.
func (m *hashMap[K, V]) Equals(other HashMap[K, V], equals func(a V, b V) bool) bool {
	if m == other {
		return true
	} else if m.len != other.Len() {
		return false
	} else {
		it := m.Iterator()
		for it.HasNext() {
			entry := it.Next()
			v, b := other.Get(entry.Key())
			if b && equals(entry.Value(), v) {
				continue
			} else {
				return false
			}
		}
		return true
	}
}

// Map applies a transformation on an entry of m i.e f((k,v)) -> (k*,v*) , using some function f and returns a new hashmap of which its keys
// and values have been transformed.
func (m hashMap[K, V]) Map(f func(e _map.MapEntry[K, V]) _map.MapEntry[K, V]) HashMap[K, V] {
	newMap := NewHashMap[K, V]()
	it := m.Iterator()
	for it.HasNext() {
		oldEntry := it.Next()
		newEntry := f(oldEntry)
		newMap.Put(newEntry.Key(), newEntry.Value())
	}
	return newMap
}

// Filter filters the hashmmap m using some predicate function that indicates whether an entry should be kept or not in a
// hashmap to be returned.
func (m hashMap[K, V]) Filter(f func(e _map.MapEntry[K, V]) bool) HashMap[K, V] {
	newMap := NewHashMap[K, V]()
	it := m.Iterator()
	for it.HasNext() {
		entry := it.Next()
		if f(entry) {
			newMap.Put(entry.Key(), entry.Value())

		}
	}
	return newMap
}
