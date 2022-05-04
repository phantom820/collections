package _map

import "collections/interfaces"

type MapIterator[K interfaces.Hashable[K], V any] interface {
	HasNext() bool
	Next() MapEntry[K, V]
	Cycle()
}

// Iterable iterable for a map.
type MapIterable[K interfaces.Hashable[K], V any] interface {
	Keys() []K
	Values() []V
	Iterator() MapIterator[K, V]
}

// MapEntry wraps around the key and value of a map. For uniformity in terms of functional
type MapEntry[K interfaces.Hashable[K], V any] struct {
	key   K
	value V
}

// NewMapEntry creates a new MapEntry with specified key and value.
func NewMapEntry[K interfaces.Hashable[K], V any](k K, v V) MapEntry[K, V] {
	return MapEntry[K, V]{key: k, value: v}
}

// Key returns the Key value of the entry.
func (entry *MapEntry[K, V]) Key() K {
	return entry.key
}

// Value returns the value of thte entry.
func (entry *MapEntry[K, V]) Value() V {
	return entry.value
}

// Map interface for what an implementation of a map should provide.
type Map[K interfaces.Hashable[K], V any] interface {
	MapIterable[K, V]
	Put(k K, v V) V
	PutIfAbsent(k K, v V) bool
	PutAll(m Map[K, V])
	Get(k K) (V, bool)
	Len() int
	Capacity() int
	LoadFactor() float32
	Keys() []K
	Contains(k K) bool
	Remove(k K) bool
	RemoveAll(c interfaces.Iterable[K])
	Clear()
	Empty() bool
}
