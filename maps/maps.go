package maps

// NewEntry creates a
func NewEntry[K any, V any](key K, value V) Entry[K, V] {
	return Entry[K, V]{key: key, value: value}
}

// Entry a key, value pair map entry.
type Entry[K any, V any] struct {
	key   K
	value V
}

// Key returns the key of the entry.
func (entry Entry[K, V]) Key() K {
	return entry.key
}

// Value returns the value of the entry.
func (entry Entry[K, V]) Value() V {
	return entry.value
}

// MapIterator an iterator for a map.
type MapIterator[K any, V any] interface {
	Next() Entry[K, V] // returns the next entry in the iterator.
	HasNext() bool     // returns true if the iterator has more emtries.
}
