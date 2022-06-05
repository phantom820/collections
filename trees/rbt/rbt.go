// Package rbt provides an implementation of a red black tree in which nodes have a key and a value.
package rbt

import (
	"errors"
	"fmt"
	"strings"

	"github.com/phantom820/collections/trees"
	"github.com/phantom820/collections/types"
)

// colors
const (
	black bool = true
	red   bool = false
)

var (
	errKeyRange = errors.New("undefined range lower key cannot be greater than upper key bound")
)

// RedBlackTree an implementation of a red black tree in which nodes have a key and a value.
type RedBlackTree[K types.Comparable[K], V any] struct {
	root *redBlackNode[K, V]
	Nil  *redBlackNode[K, V]
	len  int
}

// New creates and returns an empty RedBlackTree.
func New[K types.Comparable[K], V any]() *RedBlackTree[K, V] {
	Nil := redBlackNode[K, V]{parent: nil, left: nil, right: nil, color: black}
	return &RedBlackTree[K, V]{root: &Nil, Nil: &Nil}
}

// redBlackNode a node in a red black tree.
type redBlackNode[K types.Comparable[K], V any] struct {
	parent *redBlackNode[K, V]
	left   *redBlackNode[K, V]
	right  *redBlackNode[K, V]
	color  bool
	key    K
	value  V
}

// color returns the color of a red black node as a string for pretty printing.
func (n *redBlackNode[K, V]) Color() string {
	if n.color {
		return "(B)"
	}
	return "(R)"
}

// newRedBlackNode creates and returns a red black node with the specified key and value.
func newRedBlackNode[K types.Comparable[K], V any](key K, value V, Nil *redBlackNode[K, V]) *redBlackNode[K, V] {
	return &redBlackNode[K, V]{parent: Nil, left: Nil, right: Nil, key: key, value: value}
}

// Insert inserts a node of the form (key,value) into the tree.
func (tree *RedBlackTree[K, V]) Insert(key K, value V) bool {
	node := newRedBlackNode(key, value, tree.Nil)
	tree.insert(node)
	tree.insertFix(node)
	tree.len++
	return true
}

// Update replaces the value stored in the node identified by the key with the new value. Returns the previous value that was stored in the node and a
// boolean indicating if the previous value is valid or invalid. An invalid value is when there is no node with the specified key in the tree.
func (tree *RedBlackTree[K, V]) Update(key K, value V) (V, bool) {
	node := tree.search(key)
	if node == tree.Nil {
		var z V
		return z, false
	}
	temp := node.value
	node.value = value
	return temp, true
}

// insert inserts a node into the tree. For internal use to support Insert function.
func (tree *RedBlackTree[K, V]) insert(z *redBlackNode[K, V]) {
	var y *redBlackNode[K, V] = tree.Nil
	x := tree.root
	for x != tree.Nil {
		y = x
		if z.key.Less(x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == tree.Nil {
		tree.root = z
	} else if z.key.Less(y.key) {
		y.left = z
	} else {
		y.right = z
	}
	z.color = red
}

// insertFix fixes the tree after an insertion. For internal use to support insert function.
func (tree *RedBlackTree[K, V]) insertFix(z *redBlackNode[K, V]) {
	var y *redBlackNode[K, V]
	for z.parent.color == red {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					tree.leftRotate(z)
				}

				z.parent.color = black
				z.parent.parent.color = red
				tree.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					tree.rightRotate(z)
				}

				z.parent.color = black
				z.parent.parent.color = red
				tree.leftRotate(z.parent.parent)

			}
		}
	}
	tree.root.color = black
}

// leftRotate performs a left rotation around node x of the tree. For internal use to support deleteFix and insertFix functions.
func (tree *RedBlackTree[K, V]) leftRotate(x *redBlackNode[K, V]) {
	y := x.right
	x.right = y.left

	if y.left != tree.Nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == tree.Nil {
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

	if y.right != tree.Nil {
		y.right.parent = x
	}

	y.parent = x.parent
	if x.parent == tree.Nil {
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
	if u.parent == t.Nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

// minimum returns the node with smallest key value in the tree. For internal use to support Minimum and delete functions.
func (tree *RedBlackTree[K, V]) minimum(node *redBlackNode[K, V]) *redBlackNode[K, V] {
	if node.left == tree.Nil {
		return node
	} else {
		return tree.minimum(node.left)
	}
}

// search finds the node with the given key in the tree. For internal use to support Search function.
func (tree *RedBlackTree[K, V]) search(key K) *redBlackNode[K, V] {
	x := tree.root
	for x != tree.Nil {
		if x.key.Equals(key) {
			return x
		} else if x.key.Less(key) {
			x = x.right
		} else {
			x = x.left
		}
	}
	return x
}

// inRange checks if a given key lies withing the range [fromKey,toKey]. For internal use to support Subtree.
func (tree *RedBlackTree[K, V]) inRange(node *redBlackNode[K, V], fromKey K, toKey K) bool {
	if fromKey.Less(node.key) && node.key.Less(toKey) {
		return true
	}
	return false
}

// SubTree returns a new tree that consists of nodes with keys that are in the specified key range [fromKey,toKey]. If fromInclusive is
// true then range includes fromKey otherwise it is left out and if toInclusive is true toKey is included in the range.
func (tree *RedBlackTree[K, V]) SubTree(fromKey K, fromInclusive bool, toKey K, toInclusive bool) *RedBlackTree[K, V] {
	subTree := New[K, V]()

	if toKey.Less(fromKey) && !toKey.Equals(fromKey) {
		panic(errKeyRange)
	}

	var traverse func(node *redBlackNode[K, V])
	traverse = func(node *redBlackNode[K, V]) {
		if node == tree.Nil {
			return
		}

		if node.left != tree.Nil {
			traverse(node.left)
		}

		if node.key.Equals(fromKey) && fromInclusive {
			subTree.Insert(node.key, node.value)
		} else if node.key.Equals(toKey) && toInclusive {
			subTree.Insert(node.key, node.value)
		} else if tree.inRange(node, fromKey, toKey) {
			subTree.Insert(node.key, node.value)
		}

		if node.right != tree.Nil {
			traverse(node.right)
		}

	}
	traverse(tree.root)
	return subTree
}

// LeftSubTree returns a new tree that consists of nodes with keys that are less than or equals than the specified key. If inclusive is
// true then the node with an equal key is included otherwise its left out.
func (tree *RedBlackTree[K, V]) LeftSubTree(key K, inclusive bool) *RedBlackTree[K, V] {
	subTree := New[K, V]()
	var traverse func(node *redBlackNode[K, V])
	traverse = func(node *redBlackNode[K, V]) {
		if node == tree.Nil {
			return
		}

		if node.left != tree.Nil {
			traverse(node.left)
		}
		if node.key.Equals(key) && inclusive {
			subTree.Insert(node.key, node.value)
		} else if node.key.Less(key) {
			subTree.Insert(node.key, node.value)
		}

		if node.right != tree.Nil {
			traverse(node.right)
		}

	}
	traverse(tree.root)
	return subTree
}

// RightSubTree returns a new tree that consists of nodes with keys that are greater than or equals than the specified key. If inclusive is
// true then the node with an equal key is included otherwise its left out.
func (tree *RedBlackTree[K, V]) RightSubTree(key K, inclusive bool) *RedBlackTree[K, V] {
	subTree := New[K, V]()
	var traverse func(node *redBlackNode[K, V])
	traverse = func(node *redBlackNode[K, V]) {
		if node == tree.Nil {
			return
		}

		if node.left != tree.Nil {
			traverse(node.left)
		}
		if node.key.Equals(key) && inclusive {
			subTree.Insert(node.key, node.value)
		} else if key.Less(node.key) {
			subTree.Insert(node.key, node.value)
		}

		if node.right != tree.Nil {
			traverse(node.right)
		}

	}
	traverse(tree.root)
	return subTree
}

// Search checks if the tree contains a node with the specified key.
func (tree *RedBlackTree[K, V]) Search(key K) bool {
	return (tree.search(key) != tree.Nil)
}

// Get retrieves the value of the node with the given key. Returns a value and a boolean indicating whether the value is valid or invalid.
// An invalid value results when there is no node with the specified key in the tree.
func (tree *RedBlackTree[K, V]) Get(key K) (V, bool) {
	node := tree.search(key)
	if node == tree.Nil {
		var value V
		return value, false
	}
	return node.value, true
}

// Delete deletes the node with the specified key from the tree. Returns the value that was stored at the node and a boolean indicating if the returned value
// is valid or invalid. An invalid value results when there is no node with the specified key to be deleted from the tree.
func (tree *RedBlackTree[K, V]) Delete(key K) (V, bool) {
	node := tree.search(key)
	if node == tree.Nil {
		return node.value, false
	}
	tree.delete(node)
	tree.len--
	node.left = nil
	node.right = nil
	temp := node.value
	node = nil
	return temp, true
}

// delete deletes the node z from the tree. For internal use to support Delete function.
func (tree *RedBlackTree[K, V]) delete(z *redBlackNode[K, V]) {
	var x, y *redBlackNode[K, V]
	y = z
	yOriginalColor := y.color
	if z.left == tree.Nil {
		x = z.right
		tree.transplant(z, z.right)
	} else if z.right == tree.Nil {
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
	if yOriginalColor == black {
		tree.deleteFix(x)
	}
}

// deleteFix fixes the tree after a delete operation. For internal use to support delete function.
func (tree *RedBlackTree[K, V]) deleteFix(x *redBlackNode[K, V]) {
	var s *redBlackNode[K, V]
	for x != tree.root && x.color == black {
		if x == x.parent.left {
			s = x.parent.right
			if s.color == red {
				s.color = black
				x.parent.color = red
				tree.leftRotate(x.parent)
				s = x.parent.right
			}

			if s.left.color == black && s.right.color == black {
				s.color = red
				x = x.parent
			} else {
				if s.right.color == black {
					s.left.color = black
					s.color = red
					tree.rightRotate(s)
					s = x.parent.right
				}

				s.color = x.parent.color
				x.parent.color = black
				s.right.color = black
				tree.leftRotate(x.parent)
				x = tree.root
			}
		} else {
			s = x.parent.left
			if s.color == red {
				s.color = black
				x.parent.color = red
				tree.rightRotate(x.parent)
				s = x.parent.left
			}

			if s.left.color == black && s.right.color == black {
				s.color = red
				x = x.parent
			} else {
				if s.left.color == black {
					s.right.color = black
					s.color = red
					tree.leftRotate(s)
					s = x.parent.left
				}

				s.color = x.parent.color
				x.parent.color = black
				s.left.color = black
				tree.rightRotate(x.parent)
				x = tree.root
			}
		}
	}
	x.color = black
}

// inOrdeTraversal performs an in order traversal of the tree appending results to a slice. For internal use to support InOrderTraversal function.
func (tree *RedBlackTree[K, V]) inOrderTraversal(node *redBlackNode[K, V], data *[]K) {
	if node == tree.Nil {
		return
	}
	if node.left != tree.Nil {
		tree.inOrderTraversal(node.left, data)
	}
	*data = append(*data, node.key)
	if node.right != tree.Nil {
		tree.inOrderTraversal(node.right, data)
	}
}

// InOrderTraversal performs an in order traversal of the tree. Returns a slice containing the values collected from the traversal.
func (tree *RedBlackTree[K, V]) InOrderTraversal() []K {
	data := []K{}
	tree.inOrderTraversal(tree.root, &data)
	return data
}

// nodes collects the nodes of the tree using an in order traversal. For internal use to support Nodes function.
func (tree *RedBlackTree[K, V]) nodes(node *redBlackNode[K, V], nodes *[]trees.Node[K, V]) {
	if node == tree.Nil {
		return
	}
	if node.left != tree.Nil {
		tree.nodes(node.left, nodes)
	}
	*nodes = append(*nodes, trees.Node[K, V]{Key: node.key, Value: node.value})
	if node.right != tree.Nil {
		tree.nodes(node.right, nodes)
	}
}

// Nodes returns the nodes of the tree using an in order traversal.
func (tree *RedBlackTree[K, V]) Nodes() []trees.Node[K, V] {
	nodes := []trees.Node[K, V]{}
	tree.nodes(tree.root, &nodes)
	return nodes
}

// values collects all the values in the tree into a slice using an in order traversal. For internal use to support Values function.
func (tree *RedBlackTree[K, V]) values(node *redBlackNode[K, V], data *[]V) {
	if node == tree.Nil {
		return
	}
	if node.left != tree.Nil {
		tree.values(node.left, data)
	}
	*data = append(*data, node.value)
	if node.right != tree.Nil {
		tree.values(node.right, data)
	}
}

// Values returns a slice of values stored in the nodes of the tree using an in order traversal.
func (tree *RedBlackTree[K, V]) Values() []V {
	data := []V{}
	tree.values(tree.root, &data)
	return data
}

// keys collects the keys in the tree into a slice using an in order traversal. For internal use to support Keys function.
func (tree *RedBlackTree[K, V]) keys(node *redBlackNode[K, V], data *[]K) {
	if node == tree.Nil {
		return
	}
	if node.left != tree.Nil {
		tree.keys(node.left, data)
	}
	*data = append(*data, node.key)
	if node.right != tree.Nil {
		tree.keys(node.right, data)
	}
}

// Keys returns a slice of the keys in the tree using an in order traversal.
func (tree *RedBlackTree[K, V]) Keys() []K {
	data := []K{}
	tree.keys(tree.root, &data)
	return data
}

// Len returns the size of the tree.
func (tree *RedBlackTree[K, V]) Len() int {
	return tree.len
}

// Clear deletes all the nodes in the tree.
func (t *RedBlackTree[K, V]) Clear() {
	t.root = nil
	t.Nil = nil
	t.len = 0
	Nil := &redBlackNode[K, V]{parent: nil, left: nil, right: nil, color: black}
	t.root = Nil
	t.Nil = Nil
}

// Empty checks if the tree is empty.
func (tree *RedBlackTree[K, V]) Empty() bool {
	return tree.root == tree.Nil
}

// printInOrder a helper for string formatting the tree for pretty printing. For internal use to support String function.
func (tree *RedBlackTree[K, V]) printInOrder(node *redBlackNode[K, V], sb *strings.Builder) {
	if node == tree.Nil {
		return
	}
	if node.left != tree.Nil {
		tree.printInOrder(node.left, sb)
	}
	sb.WriteString(fmt.Sprint(node.key) + node.Color() + " ")
	if node.right != tree.Nil {
		tree.printInOrder(node.right, sb)
	}
}

// String for pretty printing the tree.
func (tree *RedBlackTree[K, V]) String() string {
	var sb strings.Builder
	tree.printInOrder(tree.root, &sb)
	return strings.TrimSpace(sb.String())
}
