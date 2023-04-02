// package pair defines a pair type that can be used to represeny key, value mapping.
package pair

// Pair  represent key, value pair.
type Pair[K any, V any] struct {
	key   K
	value V
}

// Key returns the key in the pair.
func (pair Pair[K, V]) Key() K {
	return pair.key
}

// Value returns the value in the pair.
func (pair Pair[K, V]) Value() V {
	return pair.value
}

// Of creates a pair with the given key and value.
func Of[K any, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{key: key, value: value}
}
