// Package tree provides an interface that a tree based data structure must satisfy.
package tree

// Tree interface for tree based data structures
type Tree[K any, V any] interface {
	Insert(k K, v V) bool
	Delete(k K) bool
	Clear()
	Search(k K) bool
	Update(k K, v V) (V, bool)
	Get(k K) (V, bool)
	InOrderTraversal() []K
	Values() []V
	Keys() []K
	Empty() bool
	Len() int
}
