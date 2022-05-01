package interfaces

// Iterable iterable for a map.
type MapIterable[K Comparable[K], V any] interface {
	Keys() []K
	Values() []V
	Iterator() Iterator[K]
}

// Map interface for what an implementation of a map should provide.
type Map[K Hashable[K], V any] interface {
	MapIterable[K, V]
	Put(k K, v V) bool
	PutIfAbsent(k K, v V) bool
	PutAll(m Map[K, V])
	Get(k K) (V, bool)
	Len() int
	Keys() []K
	Contains(k K) bool
	Remove(k K) bool
	RemoveAll(c Iterable[K])
	Empty() bool
}
