package rbt

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/phantom820/collections/types/optional"
	"github.com/phantom820/collections/types/pair"
)

const (
	BLACK bool = true
	RED   bool = false
)

var (
	COLOR_MAP = map[bool]string{true: "B", false: "R"}
)

// redBlackNode represents the node for a red black tree.
type redBlackNode[K comparable, V any] struct {
	parent *redBlackNode[K, V] // Parent of the node.
	left   *redBlackNode[K, V] // Left child of the node.
	right  *redBlackNode[K, V] // Right child of the node.
	color  bool                // Color of the node.
	key    K                   // Key of the node.
	value  V                   // Value of the node.
}

// Color return the color of the node, false -> black and true -> red.
func (node *redBlackNode[K, V]) Color() bool {
	return node.color
}

// newRedBlackNode creates and returns a red black node with the specified key and value.
func newRedBlackNode[K comparable, V any](key K, value V, sentinel *redBlackNode[K, V]) *redBlackNode[K, V] {
	return &redBlackNode[K, V]{parent: sentinel, left: sentinel, right: sentinel, key: key, value: value}
}

// String returns a string of the form (key, value , color) representing the node.
func (node redBlackNode[K, V]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", node.key, node.value, COLOR_MAP[node.color])
}

// RedBlackTree implementation of a red black tree in which each node has a key and associate value.
type RedBlackTree[K comparable, V any] struct {
	root     *redBlackNode[K, V] // The root of the tree.
	sentinel *redBlackNode[K, V] // The sentinel node.
	len      int                 // Number of nodes in the tree.
	lessThan func(K, K) bool     // The comparison for ordering keys.
}

// New creates a RedBlackTree. Keys are compared using the lessThan function which should satisfy.
// k1 < k2 => lessThan(k1, k2) = true and lessThan(k2,k1) = false.
// k1 = k2 => lessThan(k1,k2) = false and lessThan(k2,k1) = false.
// k1 > k2 -> lessThan(k1,k2) = false and lessThan(k2,k1) = true.
func New[K comparable, V any](lessThan func(K, K) bool) *RedBlackTree[K, V] {
	sentinel := redBlackNode[K, V]{parent: nil, left: nil, right: nil, color: BLACK}
	return &RedBlackTree[K, V]{
		root:     &sentinel,
		lessThan: lessThan,
		sentinel: &sentinel}
}

// Insert inserts a node of the form (key,value) into the tree. If the key already exist its value will be updated,
// the currently stored value is returned.
func (tree *RedBlackTree[K, V]) Insert(key K, value V) optional.Optional[V] {
	node := newRedBlackNode(key, value, tree.sentinel)
	stored, ok := tree.insert(node)
	if ok {
		tree.insertFix(node)
		tree.len++
		return optional.Empty[V]()
	}
	return optional.Of(stored)
}

// Update replaces the value stored in the node with given key and returns the previous value that was stored.
func (tree *RedBlackTree[K, V]) Update(key K, value V) (V, bool) {
	node := tree.search(key)
	if node == tree.sentinel {
		return node.value, false
	}
	temp := node.value
	node.value = value
	return temp, true
}

// insert inserts a node into the tree. For internal use to support Insert function.
func (tree *RedBlackTree[K, V]) insert(z *redBlackNode[K, V]) (V, bool) {
	var y *redBlackNode[K, V] = tree.sentinel
	x := tree.root
	for x != tree.sentinel {
		y = x
		if z.key == x.key {
			stored := x.value
			x.value = z.value
			return stored, false
		} else if tree.lessThan(z.key, x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == tree.sentinel {
		tree.root = z
	} else if tree.lessThan(z.key, y.key) {
		y.left = z
	} else {
		y.right = z
	}
	z.color = RED
	return tree.sentinel.value, true
}

// insertFix fixes the tree after an insertion. For internal use to support Insert function.
func (tree *RedBlackTree[K, V]) insertFix(z *redBlackNode[K, V]) {
	var y *redBlackNode[K, V]
	for z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					tree.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				tree.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					tree.rightRotate(z)
				}

				z.parent.color = BLACK
				z.parent.parent.color = RED
				tree.leftRotate(z.parent.parent)

			}
		}
	}
	tree.root.color = BLACK
}

// leftRotate performs a left rotation around node x of the tree. For internal use to support deleteFix and insertFix functions.
func (tree *RedBlackTree[K, V]) leftRotate(x *redBlackNode[K, V]) {
	y := x.right
	x.right = y.left

	if y.left != tree.sentinel {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == tree.sentinel {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

// rightRotate performs a right rotation around the node x of the tree. For internal use to support deleteFix and insertFix functions.
func (tree *RedBlackTree[K, V]) rightRotate(x *redBlackNode[K, V]) {
	y := x.left
	x.left = y.right

	if y.right != tree.sentinel {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == tree.sentinel {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.right = x
	x.parent = y
}

// transplant performs transplant operation on the tree. For internal use to support deleteFix and insertFix functions.
func (t *RedBlackTree[K, V]) transplant(u *redBlackNode[K, V], v *redBlackNode[K, V]) {
	if u.parent == t.sentinel {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

// minimum returns the node with smallest key value in the tree. For internal use to support Minimum and Delete functions.
func (tree *RedBlackTree[K, V]) minimum(node *redBlackNode[K, V]) *redBlackNode[K, V] {
	if node.left == tree.sentinel {
		return node
	} else {
		return tree.minimum(node.left)
	}
}

// search finds the node with the given key in the tree. For internal use to support Search function.
func (tree *RedBlackTree[K, V]) search(key K) *redBlackNode[K, V] {
	x := tree.root
	for x != tree.sentinel {
		if x.key == key {
			return x
		} else if tree.lessThan(x.key, key) {
			x = x.right
		} else {
			x = x.left
		}
	}
	return x
}

// SubTree returns a new tree that consists of nodes with keys that are in the specified key range [fromKey,toKey]. If fromInclusive is
// true then range includes fromKey otherwise it is left out and if toInclusive is true toKey is included in the range.
func (tree *RedBlackTree[K, V]) SubTree(fromKey K, fromInclusive bool, toKey K, toInclusive bool) *RedBlackTree[K, V] {
	if tree.lessThan(toKey, fromKey) && !(toKey == fromKey) {
		panic(errors.New("undefined range lower key cannot be greater than upper key bound"))
	}
	subTree := New[K, V](tree.lessThan)
	var traverse func(node *redBlackNode[K, V])
	traverse = func(node *redBlackNode[K, V]) {
		if node == tree.sentinel {
			return
		}

		if node.left != tree.sentinel {
			traverse(node.left)
		}

		if node.key == fromKey && fromInclusive {
			subTree.Insert(node.key, node.value)
		} else if node.key == toKey && toInclusive {
			subTree.Insert(node.key, node.value)
		} else if tree.lessThan(fromKey, node.key) && tree.lessThan(node.key, toKey) {
			subTree.Insert(node.key, node.value)
		}

		if node.right != tree.sentinel {
			traverse(node.right)
		}

	}
	traverse(tree.root)
	return subTree
}

// LeftSubTree returns a new tree that consists of nodes with keys that are less than or equals the specified key. If inclusive is
// true then the node with an equal key is included otherwise its left out.
func (tree *RedBlackTree[K, V]) LeftSubTree(key K, inclusive bool) *RedBlackTree[K, V] {
	subTree := New[K, V](tree.lessThan)
	var traverse func(node *redBlackNode[K, V])
	traverse = func(node *redBlackNode[K, V]) {
		if node == tree.sentinel {
			return
		}

		if node.left != tree.sentinel {
			traverse(node.left)
		}
		if node.key == key && inclusive {
			subTree.Insert(node.key, node.value)
		} else if tree.lessThan(node.key, key) {
			subTree.Insert(node.key, node.value)
		}

		if node.right != tree.sentinel {
			traverse(node.right)
		}

	}
	traverse(tree.root)
	return subTree
}

// RightSubTree returns a new tree that consists of nodes with keys that are greater than or equals than the specified key. If inclusive is
// true then the node with an equal key is included otherwise its left out.
func (tree *RedBlackTree[K, V]) RightSubTree(key K, inclusive bool) *RedBlackTree[K, V] {
	subTree := New[K, V](tree.lessThan)
	var traverse func(node *redBlackNode[K, V])
	traverse = func(node *redBlackNode[K, V]) {
		if node == tree.sentinel {
			return
		}

		if node.left != tree.sentinel {
			traverse(node.left)
		}
		if node.key == key && inclusive {
			subTree.Insert(node.key, node.value)
		} else if tree.lessThan(key, node.key) {
			subTree.Insert(node.key, node.value)
		}

		if node.right != tree.sentinel {
			traverse(node.right)
		}

	}
	traverse(tree.root)
	return subTree
}

// Search checks if the tree contains a node with the specified key.
func (tree *RedBlackTree[K, V]) Search(key K) bool {
	return (tree.search(key) != tree.sentinel)
}

// Get returns the value of the node with the given key.
func (tree *RedBlackTree[K, V]) Get(key K) optional.Optional[V] {
	node := tree.search(key)
	if node == tree.sentinel {
		return optional.Empty[V]()
	}
	return optional.Of(node.value)
}

// GetIf returns the values of the nodes with keys that satisfy the given predicate.
func (tree *RedBlackTree[K, V]) GetIf(f func(K) bool) []V {
	values := make([]V, 0)
	var traverse func(node *redBlackNode[K, V])
	traverse = func(node *redBlackNode[K, V]) {
		if node == tree.sentinel {
			return
		}

		if node.left != tree.sentinel {
			traverse(node.left)
		}

		if f(node.key) {
			values = append(values, node.value)
		}

		if node.right != tree.sentinel {
			traverse(node.right)
		}

	}
	traverse(tree.root)
	return values
}

// Delete deletes the node with the specified key from the tree and returns the value that was stored.
func (tree *RedBlackTree[K, V]) Delete(key K) optional.Optional[V] {
	node := tree.search(key)
	if node == tree.sentinel {
		return optional.Empty[V]()
	}
	tree.delete(node)
	tree.len = int(math.Max(0, float64(tree.len-1)))
	node.left = nil
	node.right = nil
	temp := node.value
	node = nil
	return optional.Of(temp)
}

// delete deletes the node z from the tree. For internal use to support Delete function.
func (tree *RedBlackTree[K, V]) delete(z *redBlackNode[K, V]) {
	var x, y *redBlackNode[K, V]
	y = z
	yOriginalColor := y.color
	if z.left == tree.sentinel {
		x = z.right
		tree.transplant(z, z.right)
	} else if z.right == tree.sentinel {
		x = z.left
		tree.transplant(z, z.left)
	} else {
		y = tree.minimum(z.right)
		yOriginalColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			tree.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		tree.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if yOriginalColor == BLACK {
		tree.deleteFix(x)
	}
}

// deleteFix fixes the tree after a delete operation. For internal use to support Delete function.
func (tree *RedBlackTree[K, V]) deleteFix(x *redBlackNode[K, V]) {
	var s *redBlackNode[K, V]
	for x != tree.root && x.color == BLACK {
		if x == x.parent.left {
			s = x.parent.right
			if s.color == RED {
				s.color = BLACK
				x.parent.color = RED
				tree.leftRotate(x.parent)
				s = x.parent.right
			}

			if s.left.color == BLACK && s.right.color == BLACK {
				s.color = RED
				x = x.parent
			} else {
				if s.right.color == BLACK {
					s.left.color = BLACK
					s.color = RED
					tree.rightRotate(s)
					s = x.parent.right
				}

				s.color = x.parent.color
				x.parent.color = BLACK
				s.right.color = BLACK
				tree.leftRotate(x.parent)
				x = tree.root
			}
		} else {
			s = x.parent.left
			if s.color == RED {
				s.color = BLACK
				x.parent.color = RED
				tree.rightRotate(x.parent)
				s = x.parent.left
			}

			if s.left.color == BLACK && s.right.color == BLACK {
				s.color = RED
				x = x.parent
			} else {
				if s.left.color == BLACK {
					s.right.color = BLACK
					s.color = RED
					tree.leftRotate(s)
					s = x.parent.left
				}

				s.color = x.parent.color
				x.parent.color = BLACK
				s.left.color = BLACK
				tree.rightRotate(x.parent)
				x = tree.root
			}
		}
	}
	x.color = BLACK
}

// values collects all the values in the tree into a slice using an in order traversal. For internal use to support Values function.
func (tree *RedBlackTree[K, V]) values(node *redBlackNode[K, V], data []V, index *int) {
	if node == tree.sentinel {
		return
	}
	if node.left != tree.sentinel {
		tree.values(node.left, data, index)
	}
	data[*index] = node.value
	*index++
	if node.right != tree.sentinel {
		tree.values(node.right, data, index)
	}
}

// Values returns a slice of values stored in the nodes of the tree using an in order traversal.
func (tree *RedBlackTree[K, V]) Values() []V {
	data := make([]V, tree.len)
	index := 0
	tree.values(tree.root, data, &index)
	return data
}

// nodes collects the nodes of the tree using an in order traversal. For internal use to support Nodes function.
func (tree *RedBlackTree[K, V]) nodes(node *redBlackNode[K, V], nodes []pair.Pair[K, V], index *int) {
	if node == tree.sentinel {
		return
	}
	if node.left != tree.sentinel {
		tree.nodes(node.left, nodes, index)
	}
	nodes[*index] = pair.Of(node.key, node.value)
	*index++
	if node.right != tree.sentinel {
		tree.nodes(node.right, nodes, index)
	}
}

// Nodes returns the nodes of the tree using an in order traversal.
func (tree *RedBlackTree[K, V]) Nodes() []pair.Pair[K, V] {
	nodes := make([]pair.Pair[K, V], tree.len)
	index := 0
	tree.nodes(tree.root, nodes, &index)
	return nodes
}

// keys collects the keys in the tree into a slice using an in order traversal. For internal use to support Keys function.
func (tree *RedBlackTree[K, V]) keys(node *redBlackNode[K, V], data []K, index *int) {
	if node == tree.sentinel {
		return
	}
	if node.left != tree.sentinel {
		tree.keys(node.left, data, index)
	}
	data[*index] = node.key
	*index++
	if node.right != tree.sentinel {
		tree.keys(node.right, data, index)
	}
}

// Keys returns a slice of the keys in the tree using an in order traversal.
func (tree *RedBlackTree[K, V]) Keys() []K {
	data := make([]K, tree.len)
	index := 0
	tree.keys(tree.root, data, &index)
	return data
}

// Len returns the size of the tree.
func (tree *RedBlackTree[K, V]) Len() int {
	return tree.len
}

// Clear deletes all the nodes in the tree.
func (tree *RedBlackTree[K, V]) Clear() {
	tree.root = nil
	tree.sentinel = nil
	tree.len = 0
	sentinel := &redBlackNode[K, V]{parent: nil, left: nil, right: nil, color: BLACK}
	tree.root = sentinel
	tree.sentinel = sentinel
}

// Empty checks if the tree is empty.
func (tree *RedBlackTree[K, V]) Empty() bool {
	return tree.len == 0
}

// printInOrder a helper for string formatting the tree for pretty printing. For internal use to support String function.
func (tree *RedBlackTree[K, V]) printInOrder(node *redBlackNode[K, V], sb *strings.Builder) {
	if node == tree.sentinel {
		return
	}
	if node.left != tree.sentinel {
		tree.printInOrder(node.left, sb)
	}
	sb.WriteString(fmt.Sprint(node) + " ")
	if node.right != tree.sentinel {
		tree.printInOrder(node.right, sb)
	}
}

// String for pretty printing the tree.
func (tree *RedBlackTree[K, V]) String() string {
	var sb strings.Builder

	tree.printInOrder(tree.root, &sb)
	return "{" + strings.TrimSpace(sb.String()) + "}"
}
