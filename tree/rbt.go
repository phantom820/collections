// Red Black Tree implementation.
package tree

import (
	"collections/types"
	"fmt"
	"strings"
)

const (
	Black bool = true
	Red   bool = false
)

// RedBlackTree interface to abstract away concrete data. Provides various method for interaccting with
// underlying concrete data.
type RedBlackTree[K types.Comparable[K], V any] interface {
	Tree[K, V]
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

// newRedBlackNode constructs a new rbt node.
func newRedBlackNode[K types.Comparable[K], V any](k K, v V, Nil *redBlackNode[K, V]) *redBlackNode[K, V] {
	return &redBlackNode[K, V]{parent: Nil, left: Nil, right: Nil, key: k, value: v}
}

// redBlackTree actual red black tree type.
type redBlackTree[K types.Comparable[K], V any] struct {
	root *redBlackNode[K, V]
	Nil  *redBlackNode[K, V]
	len  int
}

// NewRedBlackTree creates an empty rbt.
func NewRedBlackTree[K types.Comparable[K], V any]() RedBlackTree[K, V] {
	Nil := redBlackNode[K, V]{parent: nil, left: nil, right: nil, color: Black}
	return &redBlackTree[K, V]{root: &Nil, Nil: &Nil}
}

// Update replaces the old node value with key k with the new value v. Returns the old value associated with the key and true if it exists otherwise
// zero value and false.
func (t *redBlackTree[K, V]) Update(k K, v V) (V, bool) {
	n := t.search(k)
	if n != t.Nil {
		temp := n.value
		n.value = v
		return temp, true
	}
	var e V
	return e, false
}

// insert adds the node z to the tree t.
func (t *redBlackTree[K, V]) insert(z *redBlackNode[K, V]) {
	var y *redBlackNode[K, V] = t.Nil
	x := t.root
	for x != t.Nil {
		y = x
		if z.key.Less(x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == t.Nil {
		t.root = z
	} else if z.key.Less(y.key) {
		y.left = z
	} else {
		y.right = z
	}
	z.color = Red
}

// insertFix fixes the tree t after an insertion.
func (t *redBlackTree[K, V]) insertFix(z *redBlackNode[K, V]) {
	var y *redBlackNode[K, V]
	for z.parent.color == Red {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}

				z.parent.color = Black
				z.parent.parent.color = Red
				t.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}

				z.parent.color = Black
				z.parent.parent.color = Red
				t.leftRotate(z.parent.parent)

			}
		}
	}
	t.root.color = Black
}

// Insert adds the key k with value v to the tree t.
func (t *redBlackTree[K, V]) Insert(k K, v V) bool {
	x := newRedBlackNode(k, v, t.Nil)
	t.insert(x)
	t.insertFix(x)
	t.len++
	return true
}

// leftRotate performs a left rotation on node x of t.
func (t *redBlackTree[K, V]) leftRotate(x *redBlackNode[K, V]) {
	y := x.right
	x.right = y.left

	if y.left != t.Nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == t.Nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

// rightRotate performs a right rotation on the node y of t.
func (t *redBlackTree[K, V]) rightRotate(x *redBlackNode[K, V]) {
	y := x.left
	x.left = y.right

	if y.right != t.Nil {
		y.right.parent = x
	}

	y.parent = x.parent
	if x.parent == t.Nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.right = x
	x.parent = y
}

// transplant performs transplant operation on t.
func (t *redBlackTree[K, V]) transplant(u *redBlackNode[K, V], v *redBlackNode[K, V]) {
	if u.parent == t.Nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

// minimum retrieves node with smallest key value in t.
func (t *redBlackTree[K, V]) minimum(r *redBlackNode[K, V]) *redBlackNode[K, V] {
	if r.left == t.Nil {
		return r
	} else {
		return t.minimum(r.left)
	}
}

// search finds the key k in t if its present otherwise gives nil.
func (t *redBlackTree[K, V]) search(k K) *redBlackNode[K, V] {
	x := t.root
	for x != t.Nil {
		if x.key.Equals(k) {
			return x
		} else if x.key.Less(k) {
			x = x.right
		} else {
			x = x.left
		}
	}
	return x
}

// Search checks if t contains the key k.
func (t *redBlackTree[K, V]) Search(k K) bool {
	if t.search(k) != t.Nil {
		return true
	}
	return false
}

// Get returns the value in the node with key k. If there is such a value returns value,true otherwise zero value,false.
func (t *redBlackTree[K, V]) Get(k K) (V, bool) {
	n := t.search(k)
	if n != t.Nil {
		return n.value, true
	}
	var e V
	return e, false
}

// Delete deletes the key k from t.
func (t *redBlackTree[K, V]) Delete(k K) bool {
	x := t.search(k)
	if x != t.Nil {
		t.delete(x)
		t.len--
		x.left = nil
		x.right = nil
		x = nil
		return true
	}
	return false
}

// Delete deletes the node z from t.
func (t *redBlackTree[K, V]) delete(z *redBlackNode[K, V]) {
	var x, y *redBlackNode[K, V]
	y = z
	yOriginalColor := y.color
	if z.left == t.Nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == t.Nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = t.minimum(z.right)
		yOriginalColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if yOriginalColor == Black {
		t.deleteFix(x)
	}
}

// deleteFix fixes t after a delete operation.
func (t *redBlackTree[K, V]) deleteFix(x *redBlackNode[K, V]) {
	var s *redBlackNode[K, V]
	for x != t.root && x.color == Black {
		if x == x.parent.left {
			s = x.parent.right
			if s.color == Red {
				s.color = Black
				x.parent.color = Red
				t.leftRotate(x.parent)
				s = x.parent.right
			}

			if s.left.color == Black && s.right.color == Black {
				s.color = Red
				x = x.parent
			} else {
				if s.right.color == Black {
					s.left.color = Black
					s.color = Red
					t.rightRotate(s)
					s = x.parent.right
				}

				s.color = x.parent.color
				x.parent.color = Black
				s.right.color = Black
				t.leftRotate(x.parent)
				x = t.root
			}
		} else {
			s = x.parent.left
			if s.color == Red {
				s.color = Black
				x.parent.color = Red
				t.rightRotate(x.parent)
				s = x.parent.left
			}

			if s.left.color == Black && s.right.color == Black {
				s.color = Red
				x = x.parent
			} else {
				if s.left.color == Black {
					s.right.color = Black
					s.color = Red
					t.leftRotate(s)
					s = x.parent.left
				}

				s.color = x.parent.color
				x.parent.color = Black
				s.left.color = Black
				t.rightRotate(x.parent)
				x = t.root
			}
		}
	}
	x.color = Black
}

// inOrdeTraversal performs an in order traversal of t appending resuls to string.
func (t *redBlackTree[K, V]) inOrderTraversal(n *redBlackNode[K, V], data *[]K) {
	if n == t.Nil {
		return
	}
	if n.left != t.Nil {
		t.inOrderTraversal(n.left, data)
	}
	*data = append(*data, n.key)
	if n.right != t.Nil {
		t.inOrderTraversal(n.right, data)
	}
}

// InOrderTraversal an in order traversal of t with results (values) returned as a slice.
func (t *redBlackTree[K, V]) InOrderTraversal() []K {
	data := []K{}
	t.inOrderTraversal(t.root, &data)
	return data
}

// values collects values in t into a slice using an in order traversal.
func (t *redBlackTree[K, V]) values(n *redBlackNode[K, V], data *[]V) {
	if n == t.Nil {
		return
	}

	if n.left != t.Nil {
		t.values(n.left, data)
	}
	*data = append(*data, n.value)
	if n.right != t.Nil {
		t.values(n.right, data)
	}
}

// Values collects values in t into a slice using an in order traversal.
func (t *redBlackTree[K, V]) Values() []V {
	data := []V{}
	t.values(t.root, &data)
	return data
}

// Keys collects the keys in t to a slice.
func (t *redBlackTree[K, V]) Keys() []K {
	data := []K{}
	t.keys(t.root, &data)
	return data
}

// keys collects the keys in t into a slice using an in order traversal.
func (t *redBlackTree[K, V]) keys(n *redBlackNode[K, V], data *[]K) {
	if n == t.Nil {
		return
	}
	if n.left != t.Nil {
		t.keys(n.left, data)
	}
	*data = append(*data, n.key)
	if n.right != t.Nil {
		t.keys(n.right, data)
	}
}

// Len returns number of nodes in the tree t.
func (t *redBlackTree[K, V]) Len() int {
	return t.len
}

// Clear deletes all the nodes in the tree t.
func (t *redBlackTree[K, V]) Clear() {
	keys := t.Keys()
	for _, k := range keys {
		t.Delete(k)
	}
}

// Empty checks if t has no elements.
func (t *redBlackTree[L, V]) Empty() bool {
	return t.root == t.Nil
}

func (t *redBlackTree[K, V]) printInOrder(n *redBlackNode[K, V], sb *strings.Builder) {
	if n == t.Nil {
		return
	}
	if n.left != t.Nil {
		t.printInOrder(n.left, sb)
	}
	sb.WriteString(fmt.Sprint(n.key) + n.Color() + " ")
	if n.right != t.Nil {
		t.printInOrder(n.right, sb)
	}
}

// String for pretty printing the tree t.
func (t *redBlackTree[K, V]) String() string {
	var sb strings.Builder
	t.printInOrder(t.root, &sb)
	return strings.TrimSpace(sb.String())
}
