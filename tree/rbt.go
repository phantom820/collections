package tree

import (
	"collections/interfaces"
	"fmt"
	"strings"
)

const (
	Black bool = true
	Red   bool = false
)

// redBlackNode a struct representing how node of an rbt looks like.
type redBlackNode[K interfaces.Comparable[K], V any] struct {
	parent *redBlackNode[K, V]
	left   *redBlackNode[K, V]
	right  *redBlackNode[K, V]
	color  bool
	key    K
	value  V
}

// newRedBlackNode constructs a new rbt node.
func newRedBlackNode[K interfaces.Comparable[K], V any](k K, v V, Nil *redBlackNode[K, V]) *redBlackNode[K, V] {
	return &redBlackNode[K, V]{parent: Nil, left: Nil, right: Nil, key: k, value: v}
}

// RedBlackTree rbt datastructure with tree functions.
type RedBlackTree[K interfaces.Comparable[K], V any] interface {
	Tree[K, V]
}

// redBlackTree struct for the rbt.
type redBlackTree[K interfaces.Comparable[K], V any] struct {
	root *redBlackNode[K, V]
	Nil  *redBlackNode[K, V]
}

// NewRedBlackTree creates an empty rbt.
func NewRedBlackTree[K interfaces.Comparable[K], V any]() RedBlackTree[K, V] {
	Nil := redBlackNode[K, V]{parent: nil, left: nil, right: nil, color: Black}
	return &redBlackTree[K, V]{root: &Nil, Nil: &Nil}
}

// insert adds the node z to t.
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

// Update updates the old node value with key k with the new node value v.
func (t *redBlackTree[K, V]) Update(k K, v V) bool {
	n := t.search(k)
	if n != t.Nil {
		n.value = v
		return true
	}
	return false
}

// Insert adds the key k with value v to t.
func (t *redBlackTree[K, V]) Insert(k K, v V) bool {
	x := newRedBlackNode(k, v, t.Nil)
	t.insert(x)
	t.insertFix(x)
	return true
}

// Delete deletes the key k from t.
func (t *redBlackTree[K, V]) Delete(k K) bool {
	x := t.search(k)
	if x != t.Nil {
		t.delete(x)
		return true
	}
	return false
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
func (t *redBlackTree[K, V]) rightRotate(y *redBlackNode[K, V]) {
	x := y.left
	y.left = x.right

	if x.right != t.Nil {
		x.right.parent = y
	}

	x.parent = y.parent

	if y.parent == t.Nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	}

	x.right = y
	y.parent = x
}

// insertFix fixes t after an insertion.
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

// minimum retrieves node with smallest key value.
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

// deleteFix fixes t after a delete operation.
func (tree *redBlackTree[K, V]) deleteFix(x *redBlackNode[K, V]) {
	t := tree.root
	for x != tree.Nil && x != t && x.color == Black {
		if x == x.parent.left {
			w := x.parent.right

			if w.color == Red {
				w.color = Black
				x.parent.color = Red
				tree.leftRotate(x.parent)
				w = x.parent.right
			}

			if w.left.color == Black && w.right.color == Black {
				w.color = Red
				x = x.parent
			} else if w.right.color == Black {
				w.left.color = Black
				w.color = Red
				tree.rightRotate(w)
				w = x.parent.right
			}
			w.color = x.parent.color
			x.parent.color = Black
			w.right.color = Black
			tree.leftRotate(x.parent)
			x = t
		} else if x == x.parent.right {
			w := x.parent.left
			if w.color == false {
				w.color = Black
				w.parent.color = Red
				tree.rightRotate(x.parent)
				w = x.parent.right
			}

			if w.left.color == Black && w.right.color == Black {
				w.color = Red
				x = x.parent
			} else if w.left.color == Black {
				w.right.color = Black
				w.color = Red
				tree.leftRotate(w)
				w = x.parent.left
			}

			w.color = x.parent.color
			w.left.color = Black
			tree.leftRotate(x.parent)
			x = t
		}
	}
	x.color = Black
}

// Delete deletes the node z from t.
func (t *redBlackTree[K, V]) delete(z *redBlackNode[K, V]) {
	y := z
	var x *redBlackNode[K, V]
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

// inOrdeTraversal performs an in order traversal of t appending resuls to string.
func (t *redBlackTree[K, V]) inOrderTraversal(n *redBlackNode[K, V], sb *strings.Builder) {
	if n.left != t.Nil {
		t.inOrderTraversal(n.left, sb)
	}
	sb.WriteString(fmt.Sprint(n.value) + " ")
	if n.right != t.Nil {
		t.inOrderTraversal(n.right, sb)
	}
}

// InOrderTraversal an in order traversal of t with results returned as a string.
func (t *redBlackTree[K, V]) InOrderTraversal() string {
	var sb strings.Builder
	if t.root == t.Nil {
		return ""
	}
	t.inOrderTraversal(t.root, &sb)
	return strings.TrimSpace(sb.String())
}

// collect collects values in t into a slice using an in order traversal.
func (t *redBlackTree[K, V]) collect(n *redBlackNode[K, V], data *[]V) {
	if n.left != t.Nil {
		t.collect(n.left, data)
	}
	*data = append(*data, n.value)
	if n.right != t.Nil {
		t.collect(n.right, data)
	}
}

// collects values in t into a slice using an in order traversal.
func (t *redBlackTree[K, V]) Collect() []V {
	data := []V{}
	t.collect(t.root, &data)
	return data
}

// Keys collects the keys in t to a slice.
func (t *redBlackTree[K, V]) Keys() []K {
	data := []K{}
	t.keys(t.root, &data)
	return data
}

// lazy evaluation
func (t *redBlackNode[K, V]) tr() {

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

// Empty checks if t has no elements.
func (t *redBlackTree[L, V]) Empty() bool {
	return t.root == t.Nil
}
