// Package trees provides an interface that tree based data structures will satisfy and other utils.
package trees

// Tree interface specifying methods a tree based data structure should provide.
type Tree[K any, V any] interface {
	Insert(key K, value V) bool      // Inserts a node with the specified key and value.
	Delete(key K) bool               // Deletes the node with specified key. Returns true if such a node was found and deleted otherwise false.
	Clear()                          // Deleted all the nodes in the tree.
	Search(key K) bool               // Searches for a node with the specified key.
	Update(key K, value V) (V, bool) // Updates the node with specified key with the new value. Returns the old value if there was such a node.
	Get(key K) (V, bool)             // Retrieves the value of the node with the specified key.
	InOrderTraversal() []K           // Performs an in order traversal and returns results in a slice.
	Values() []V                     // Retrieves all the values sin the tree.
	Keys() []K                       // Retrieves all the keys in the tree.
	Empty() bool                     // Chekcs if the tree is empty.
	Len() int                        // Returns the size of the tree.
}

// Node a type for representing node of a tree.
type Node[K any, V any] struct {
	Key   K
	Value V
}
