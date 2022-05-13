// Package rbt provides an implementation of a Red Black Tree with nodes that store a key and an associated value.
package rbt

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections/types"
)

const (
	black bool = true
	red   bool = false
)

// RedBlackTree an implementation of a red black tree that has nodes identified by key K and stores values V.
type RedBlackTree[K types.Comparable[K], V any] struct {
	root *redBlackNode[K, V]
	Nil  *redBlackNode[K, V]
	len  int
}

// New creates an empty red black tree.
func New[K types.Comparable[K], V any]() *RedBlackTree[K, V] {
	Nil := redBlackNode[K, V]{parent: nil, left: nil, right: nil, color: black}
	return &RedBlackTree[K, V]{root: &Nil, Nil: &Nil}
}

// redBlackNode represent a node for an red black tree. Stores a key k and and associated data.
type redBlackNode[K types.Comparable[K], V any] struct {
	parent *redBlackNode[K, V]
	left   *redBlackNode[K, V]
	right  *redBlackNode[K, V]
	color  bool
	key    K
	value  V
}

// color gets the color of a node as a string. for pretty printing.
func (n *redBlackNode[K, V]) Color() string {
	if n.color {
		return "(B)"
	}
	return "(R)"
}

// newRedBlackNode constructs a new red black tree node.
func newRedBlackNode[K types.Comparable[K], V any](k K, v V, Nil *redBlackNode[K, V]) *redBlackNode[K, V] {
	return &redBlackNode[K, V]{parent: Nil, left: Nil, right: Nil, key: k, value: v}
}

// Update replaces the value at node with specified key  with the new given value and returns the old value. Will return zero value and false
// if there was no such node in the tree.
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

// insert adds the node z to the tree.
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

// insertFix fixes the tree t after an insertion.
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

// Insert insert a node with the specified key and value to the tree.
func (tree *RedBlackTree[K, V]) Insert(key K, value V) bool {
	node := newRedBlackNode(key, value, tree.Nil)
	tree.insert(node)
	tree.insertFix(node)
	tree.len++
	return true
}

// leftRotate performs a left rotation aound node x of the tree.
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

// rightRotate performs a right rotation around the node y of the tree.
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

// transplant performs transplant operation on the tree. For internal use only.
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

// minimum retrieves node with smallest key value in the tree. For internal use to support operation such as Minimum,Delete.
func (t *RedBlackTree[K, V]) minimum(r *redBlackNode[K, V]) *redBlackNode[K, V] {
	if r.left == t.Nil {
		return r
	} else {
		return t.minimum(r.left)
	}
}

// search finds the node with the given key in the tree.
func (t *RedBlackTree[K, V]) search(key K) *redBlackNode[K, V] {
	x := t.root
	for x != t.Nil {
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

// Search checks if the tree contains a node with the specified key.
func (tree *RedBlackTree[K, V]) Search(key K) bool {
	if tree.search(key) != tree.Nil {
		return true
	}
	return false
}

// Get returns the value in the node with the specified key. Will return the zero value and false if there is no such node.
func (tree *RedBlackTree[K, V]) Get(key K) (V, bool) {
	node := tree.search(key)
	if node == tree.Nil {
		var e V
		return e, false
	}
	return node.value, true
}

// Delete deletes the node with the specified key from the tree.
func (tree *RedBlackTree[K, V]) Delete(key K) bool {
	node := tree.search(key)
	if node == tree.Nil {
		return false
	}
	tree.delete(node)
	tree.len--
	node.left = nil
	node.right = nil
	node = nil
	return true
}

// Delete deletes the node z from the tree. for internal use to support Delete operation.
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

// deleteFix fixes the tree after a delete operation.
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

// inOrdeTraversal performs an in order traversal of the tree appending results to a slice.
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

// InOrderTraversal performs an in order traversal of the tree returning results as a slice.
func (tree *RedBlackTree[K, V]) InOrderTraversal() []K {
	data := []K{}
	tree.inOrderTraversal(tree.root, &data)
	return data
}

// values collects all the values in the tree into a slice using an in order traversal.
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

// Values collects values in the tree into a slice using an in order traversal.
func (tree *RedBlackTree[K, V]) Values() []V {
	data := []V{}
	tree.values(tree.root, &data)
	return data
}

// keys collects the keys in the ttrr into a slice using an in order traversal.
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

// Keys collects all the keys in the tree int a slice using an in order traversal.
func (t *RedBlackTree[K, V]) Keys() []K {
	data := []K{}
	t.keys(t.root, &data)
	return data
}

// Len returns the size of the tree.
func (tree *RedBlackTree[K, V]) Len() int {
	return tree.len
}

// Clear deletes all the nodes in the tree.
func (t *RedBlackTree[K, V]) Clear() {
	keys := t.Keys()
	for _, k := range keys {
		t.Delete(k)
	}
}

// Empty checks if the tree is empty.
func (tree *RedBlackTree[K, V]) Empty() bool {
	return tree.root == tree.Nil
}

// printInOrder a helper for string formatting the tree for pretty printing.
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
