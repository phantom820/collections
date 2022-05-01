package tree

// Tree interface for tree data structures
type Tree[K any, V any] interface {
	Insert(k K, v V) bool
	Delete(k K) bool
	Search(k K) bool
	Update(k K, v V) bool
	Get(k K) (V, bool)
	InOrderTraversal() string
	Collect() []V
	Keys() []K
	Empty() bool
}
