// Package list provides implementations for a singly linked list (ForwardList) and a doubly linked list (List).
package list

import (
	"collections/iterator"
	"collections/types"
	"errors"

	"fmt"
	"strings"
)

// Errors for  operations that may be invalid on a list.
var (
	EmptyListError   = errors.New("cannot remove from an empty list.")
	OutOfBoundsError = errors.New("index out of bounds.")
)

// List an  implementation of a doubly linked list.
type List[T types.Equitable[T]] struct {
	head *node[T]
	tail *node[T]
	len  int
}

// NewList creates a list with the specified elements. If no elements are specified an empty list is created.
func NewList[T types.Equitable[T]](elements ...T) *List[T] {
	l := List[T]{head: nil, len: 0}
	l.AddSlice(elements)
	return &l
}

// node a link for a doubly linked list. Stores a value of some type T along with prev and next pointer. This type is for internal use.
type node[T types.Equitable[T]] struct {
	prev  *node[T]
	next  *node[T]
	value T
}

// newNode creates a new node with the value v.
func newNode[T types.Equitable[T]](v T) *node[T] {
	return &node[T]{value: v, prev: nil, next: nil}
}

// listIterator type to implement an iterator for a list.
type listIterator[T types.Equitable[T]] struct {
	n     *node[T] // Used for Next() and HasNext().
	start *node[T] // Used to cycle an iterator.
}

// HasNext checks if the iterator it has a next element to yield.
func (it *listIterator[T]) HasNext() bool {
	return it.n != nil
}

// Next returns the next element in the iterator it. Will panic if iterator has been exhausted.
func (it *listIterator[T]) Next() T {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	n := it.n
	it.n = it.n.next
	return n.value
}

// Cycle resets the iterator.
func (it *listIterator[T]) Cycle() {
	it.n = it.start
}

// Iterator returns an iterator for the list l.
func (l *List[T]) Iterator() iterator.Iterator[T] {
	return &listIterator[T]{n: l.head, start: l.head}
}

// Front returns the front element of the list l. Will panic if l has no front element.
func (l *List[T]) Front() T {
	if l.head != nil {
		return l.head.value
	}
	panic(EmptyListError)
}

// Back returns the back element of the list l. Will panic if l has no back element.
func (l *List[T]) Back() T {
	if l.tail != nil {
		return l.tail.value
	}
	panic(EmptyListError)
}

// Swap swaps the element at index i and the element at index j. This is done using links. Will panics if one or both of the specified indices
// out of bounds.
func (l *List[T]) Swap(i, j int) {
	if i < 0 || i >= l.len || j < 0 || j >= l.len {
		panic(OutOfBoundsError)
	} else {
		x := l.nodeAt(i)
		y := l.nodeAt(j)
		if x == l.head {
			l.head = y
		} else if y == l.head {
			l.head = x

		}
		if x == l.tail {
			l.tail = y

		} else if y == l.tail {
			l.tail = x

		}

		// Swapping x and y.
		var temp *node[T]
		temp = x.next
		x.next = y.next
		y.next = temp

		if x.next != nil {
			x.next.prev = x

		}
		if y.next != nil {
			y.next.prev = y

		}

		temp = x.prev
		x.prev = y.prev
		y.prev = temp

		if x.prev != nil {
			x.prev.next = x

		}
		if y.prev != nil {
			y.prev.next = y
		}

	}
}

// nodeAt retrieves the node at index i in list l. This is for internal use for supporting operations like Swap.
func (l *List[T]) nodeAt(i int) *node[T] {
	j := 0
	var n *node[T]
	for e := l.head; e != nil; e = e.next {
		if j == i {
			n = e
		}
		j++
	}
	return n
}

// At retrieves the element at index i in list l. Will panic If i is out of bounds.
func (l *List[T]) At(i int) T {
	if i < 0 || i >= l.len {
		panic(OutOfBoundsError)
	}
	it := l.Iterator()
	j := 0
	var v T
	for it.HasNext() {
		e := it.Next()
		if j == i {
			v = e
			break
		}
		j++
	}
	return v
}

// AddFront adds element e to the front of the list l.
func (l *List[T]) AddFront(e T) {
	n := newNode(e)
	if l.head != nil {
		n.next = l.head
		l.head.prev = n
		l.head = n
		l.len++
		return
	}
	l.head = n
	l.tail = n
	l.len++
}

// AddBack adds element e to the back of the list l.
func (l *List[T]) AddBack(e T) {
	if l.head == nil {
		l.AddFront(e)
		return
	}
	n := newNode(e)
	l.tail.next = n
	n.prev = l.tail
	l.tail = n
	l.len++
}

// AddAt adds element to list l at specified index i. Will panic if i is out of bounds.
func (l *List[T]) AddAt(i int, e T) {
	if i < 0 || i >= l.len {
		panic(OutOfBoundsError)
	} else if i == 0 {
		l.AddFront(e)
		return
	} else if i == l.len-1 {
		l.AddBack(e)
	} else {
		j := 0
		n := newNode(e)
		for x := l.head; x != nil; x = x.next {
			if j == i-1 {
				n.prev = x
				n.next = x.next
				x.next = n
				l.len++
				break
			}
			j = j + 1
		}
		return
	}
}

// Add adds element e to the back of the list l.
func (l *List[T]) Add(e T) bool {
	l.AddBack(e)
	return true
}

// Set replaces the element at index i in the list l with the new element e. Returns the old element that was at index i.
func (l *List[T]) Set(i int, e T) T {
	if i < 0 || i >= l.len {
		panic(OutOfBoundsError)
	}
	n := l.nodeAt(i)
	temp := n.value
	n.value = e
	return temp
}

// AddAll adds all elements from some iterable elements to the list l.
func (l *List[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		l.Add(it.Next())
	}
}

// AddSlice adds element from a slice s into the list l.
func (l *List[T]) AddSlice(s []T) {
	for _, e := range s {
		l.Add(e)
	}
}

// Len returns the size of the list l.
func (l *List[T]) Len() int {
	return l.len
}

// search traverses the list l looking for element e. For internal use to support operations such as Contains, AddAt and  so on.
func (l *List[T]) search(e T) *node[T] {
	curr := l.head
	for curr != nil {
		if curr.value.Equals(e) {
			return curr
		}
		curr = curr.next
	}
	return nil
}

// Contains checks if element e is in the list l.
func (l *List[T]) Contains(e T) bool {
	return l.search(e) != nil
}

// RemoveFront removes and returns the front element of the list l. Will panic if l has no front element.
func (l *List[T]) RemoveFront() T {
	if l.len == 0 {
		panic(EmptyListError)
	} else if l.len == 1 {
		n := l.head
		l.head = n.next // subsequent operations are to avoid memory leaks.
		l.tail = nil
		v := n.value
		n.next = nil
		n.prev = nil
		n = nil
		l.len -= 1
		return v
	} else {
		n := l.head
		l.head = n.next
		v := n.value
		n.next = nil
		n.prev = nil
		n = nil
		l.len -= 1
		return v
	}
}

// RemoveBack removes and returns the back element of the list l. Will panic if l has no back element.
func (l *List[T]) RemoveBack() T {
	if l.len <= 1 {
		return l.RemoveFront()
	} else {
		n := l.tail
		l.tail = l.tail.prev
		l.tail.next = nil
		l.len -= 1
		v := n.value
		n.prev = nil
		n.next = nil
		n = nil
		return v
	}
}

// RemoveAt removes the element at the specified index i. Will panic if index i is out of bounds.
func (l *List[T]) RemoveAt(i int) T {
	if l.Empty() {
		panic(EmptyListError)
	} else if i < 0 || i >= l.len {
		panic(OutOfBoundsError)
	} else if i == 0 {
		return l.RemoveFront()
	} else if i == l.len-1 {
		return l.RemoveBack()
	} else {
		n := l.nodeAt(i)
		return l.removeNode(n)
	}
}

// removeNode removes the specified node. For internal use for functions such as RemoveAt.
func (l *List[T]) removeNode(curr *node[T]) T {
	n := curr.prev
	n.next = curr.next
	n.next.prev = n
	curr.next = nil
	curr.prev = nil
	curr = nil
	l.len -= 1
	return n.value
}

// Remove removes element e from the list l if its present. This removes the first occurence of e.
func (l *List[T]) Remove(e T) bool {
	curr := l.search(e)
	if curr == nil {
		return false
	} else if curr == l.head {
		l.RemoveFront()
		return true
	} else if curr == l.tail {
		l.RemoveBack()
		return true
	} else {
		l.removeNode(curr)
		return true
	}
}

// RemoveAll removes all the elements from the list l that appear in some iterable elements.
func (l *List[T]) RemoveAll(elements iterator.Iterable[T]) {
	defer func() {
		if r := recover(); r != nil {
			// do nothing just fail safe if l ends up empty from the removals.
		}
	}()
	it := elements.Iterator()
	for it.HasNext() {
		l.Remove(it.Next())
	}
}

// Reverse returns a new list that is the reverse of l. This uses extra memory since we inserting into a new list.
func (l *List[T]) Reverse() *List[T] {
	r := NewList[T]()
	h := l.head
	for h != nil {
		r.AddFront(h.value)
		h = h.next
	}
	return r
}

// Clear removes all elements from the list l.
func (l *List[T]) Clear() {
	for l.head != nil {
		l.RemoveFront()
	}
}

// Equals checks if list l and list other are equal. If they are the same reference/ have same size and elements then they are equal.
// Otherwise they are not equal.
func (l *List[T]) Equals(other *List[T]) bool {
	if l == other {
		return true
	} else if l.len != other.Len() {
		return false
	} else {
		it := l.Iterator()
		otherIt := other.Iterator()
		for it.HasNext() {
			a := it.Next()
			b := otherIt.Next()
			if !a.Equals(b) {
				return false
			}
		}
		return true
	}
}

// Empty checks if the list l is empty.
func (l *List[T]) Empty() bool {
	return l.len == 0
}

// Collect collects all elements of the list l into a slice.
func (l *List[T]) Collect() []T {
	data := make([]T, l.len)
	i := 0
	for e := l.head; e != nil; e = e.next {
		data[i] = e.value
		i = i + 1
	}
	return data
}

// traversal for pretty printing purposes.
func (l *List[T]) traversal() string {
	sb := make([]string, 0)
	for e := l.head; e != nil; e = e.next {
		sb = append(sb, fmt.Sprint(e.value))
	}
	return "[" + strings.Join(sb, " ") + "]"
}

// String string format for list l.
func (l *List[T]) String() string {
	return l.traversal()
}

// Map transforms each element of the list l using a function f and returns a new list with transformed elements.
func (l *List[T]) Map(f func(e T) T) *List[T] {
	newList := NewList[T]()
	for e := l.head; e != nil; e = e.next {
		newE := f(e.value)
		newList.Add(newE)
	}
	return newList
}

// Filter filters the elements of the list l using a predicate function f and returns new list with elements satisfying predicate.
func (l *List[T]) Filter(f func(e T) bool) *List[T] {
	newList := NewList[T]()
	for e := l.head; e != nil; e = e.next {
		if f(e.value) {
			newList.Add(e.value)
		}
	}
	return newList
}
