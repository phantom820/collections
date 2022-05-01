package hashmap

import (
	"collections/interfaces"
	"collections/tree"
)

const (
	LoadFactorLimit = 0.75
	Capacity        = 16
)

// HashMap a hashmap associated some keys k witth values v.
type HashMap[K interfaces.Hashable[K], V any] interface {
	interfaces.Map[K, V]
}

// hashMap underlying concrete implementation for a Hashmap.
// capacity -> initial number of buckets, loadFactorLimit -> dictates when should expansion take place.
// buckets -> slice with actual underlying containers (Red Black Trees) in this case. len -> no keys stored.
type hashMap[K interfaces.Hashable[K], V any] struct {
	capacity        int
	loadFactorLimit float32
	buckets         []tree.RedBlackTree[K, V]
	len             int
}

// NewHashMap creates a new empty HashMap with default initial capacity and load factor limit.
func NewHashMap[K interfaces.Hashable[K], V any]() HashMap[K, V] {
	buckets := make([]tree.RedBlackTree[K, V], Capacity)
	m := hashMap[K, V]{capacity: Capacity, loadFactorLimit: LoadFactorLimit, buckets: buckets, len: 0}
	return &m
}

// NewHashMapWith creates a new empty HashMap with the specified capacity and load factor limit.
func NewHashMapWith[K interfaces.Hashable[K], V any](capacity int, loadFactorLimit float32) HashMap[K, V] {
	buckets := make([]tree.RedBlackTree[K, V], capacity)
	m := hashMap[K, V]{capacity: capacity, loadFactorLimit: loadFactorLimit, buckets: buckets, len: 0}
	return &m
}

type mapIterator[K interfaces.Comparable[K], V any] struct {
	index      int
	maxIndex   int
	exhausted  bool // check if the iterator has been used up (parsed all values)
	bucket     []K
	keys       int
	maxkeys    int
	nextBucket func(i int) tree.RedBlackTree[K, V]
}

// Cycle refreshes the iterator if it has been exhausted.
func (it *mapIterator[K, V]) Cycle() {
	if it.exhausted {
		it.index = 0
		it.bucket = nil
		it.keys = 0
		it.exhausted = false
	}
}

// HasNext checks if there is next value to be produced by iterator.
func (it *mapIterator[K, V]) HasNext() bool {
	if it.index < it.maxIndex && it.keys < it.maxkeys {
		return true
	}
	it.exhausted = true
	return false
}

// Next gives out the next element for a map iterator.  Start from bucket 0 and grab values from underlying container "emitting" them
// until the container is exhausted, then move on to next container. Now what happens when someone calls next on an exhausted iterator ??
// they should get the zero value up until the iterator is cycled back.
func (it *mapIterator[K, V]) Next() K {
	if it.exhausted {
		var k K
		return k
	}
	next := func() K {
		if it.bucket != nil && len(it.bucket) > 0 {
			k := it.bucket[0]
			it.keys++
			it.bucket = it.bucket[1:len(it.bucket)]
			return k
		} else {
			var newBucket []K
			if it.index == 0 && it.bucket == nil { // Special case for starting. Can we be more concise here ??
				newBucketPtr := it.nextBucket(0)
				if newBucketPtr != nil {
					newBucket = newBucketPtr.Keys()
				}
				k := newBucket[0]
				it.keys++
				it.bucket = newBucket[1:]
				return k
			}
			// otherwise search for a new bucket that has values.
			for j := it.index + 1; j < it.maxIndex; j++ {
				newBucketPtr := it.nextBucket(j)
				it.index++
				if newBucketPtr != nil {
					it.index = j
					newBucket = newBucketPtr.Keys()
					break
				}
			}
			k := newBucket[0]
			it.keys++
			it.bucket = newBucket[1:]
			return k
		}
	}
	return next()
}

// Iterator returns an iterator on the keys of the map m.
func (m *hashMap[K, V]) Iterator() interfaces.Iterator[K] {
	nextBucket := func(i int) tree.RedBlackTree[K, V] {
		if i < len(m.buckets) {
			return m.buckets[i]
		}
		return nil
	}
	it := mapIterator[K, V]{index: 0, nextBucket: nextBucket,
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

// Put associates the specified value with the specified key in m. If the key alreaady exists then its value will be updated.
func (m *hashMap[K, V]) Put(k K, v V) bool {
	if m.loadFactor() >= m.loadFactorLimit { // if we have crossed the load factor limit resize.
		m.resize()
	}
	c := k.HashCode()
	i := c % m.capacity
	if m.buckets[i] == nil {
		m.buckets[i] = tree.NewRedBlackTree[K, V]()
		m.buckets[i].Insert(k, v)
		m.len++
		return true
	} else {
		if m.buckets[i].Update(k, v) { // try updating otherwise make a new insertion.
			return true
		}
		m.len++
		return m.buckets[i].Insert(k, v)
	}
}

// PutIfAbsent adds the value v associate with key k to map m only if k has not been mapped already.
func (m *hashMap[K, V]) PutIfAbsent(k K, v V) bool {
	c := k.HashCode()
	i := c % m.capacity
	if m.buckets[i] == nil {
		return m.Put(k, v)
	} else if !m.buckets[i].Search(k) {
		m.buckets[i].Insert(k, v)
		m.len++
		return true
	}
	return false
}

// PutAll adds all the values from other map into the map m. Note this has the side effect that if a key
// is present in m and other then the associated value  in m will be replaced by the associated value  in other.
func (m *hashMap[K, V]) PutAll(other interfaces.Map[K, V]) {
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
	i := k.HashCode() % m.capacity
	if m.buckets[i] == nil {
		var e V
		return e, false
	}
	return m.buckets[i].Get(k)
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
	i := k.HashCode() % m.capacity
	if m.buckets[i] == nil {
		return false
	}
	r := m.buckets[i].Delete(k)
	if r {
		m.len--
		if m.buckets[i].Empty() {
			m.buckets[i] = nil
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
func (m *hashMap[K, V]) loadFactor() float32 {
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
		for _, v := range bucket.Collect() { // This should not be costly since the buckets should be small rbt trees O(n*k) n -> capacity
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
		for _, k := range bucket.Keys() { // This should not be costly since the buckets should be small rbt trees O(n*k) n -> capacity
			data[i] = k // k -> average size of a bucket.
			i = i + 1
		}
	}
	return data
}

// Empty checks if the map m is empty.
func (m *hashMap[K, V]) Empty() bool {
	return m.len == 0
}
