package list

import (
	"collections/interfaces"
	"errors"
	"fmt"
	"strings"
)

// Errors for operations that may be inapplicable on a list.
var (
	EmptyListError     = errors.New("cannot remove from an empty list.")
	OutOfBoundsError   = errors.New("index out of bounds.")
	NoNextElementError = errors.New("iterator has no next element.")
)

// List interface to abstract away underlying concrete implementation. Provides functions to operate on
// the underlying list.
type List[T interfaces.Equitable[T]] interface {
	interfaces.Collection[T]
	interfaces.Functional[T, List[T]]
	AddFront(e T)
	Front() T
	AddBack(e T)
	Back() T
	AddAt(i int, e T)
	RemoveAt(i int) T
	At(i int) T
	Set(i int, e T) T
	Swap(i, j int)
	RemoveFront() T
	RemoveBack() T
	Equals(other List[T]) bool
}

// node a node/link for a linked list. Stores a value of some type T.
type node[T interfaces.Equitable[T]] struct {
	prev  *node[T]
	next  *node[T]
	value T
}

// newNode creates a new node with the value v.
func newNode[T interfaces.Equitable[T]](v T) *node[T] {
	return &node[T]{value: v, prev: nil, next: nil}
}

// list actual concrete type for a list with unexported fields.
// head -> head node of the list , tail -> tail node of the list , len -> size of the list.
type list[T interfaces.Equitable[T]] struct {
	head *node[T]
	tail *node[T]
	len  int
}

// NewList creates a new empty list that can store values of type T.
func NewList[T interfaces.Equitable[T]]() List[T] {
	l := list[T]{head: nil, len: 0}
	return &l
}

// listIterator an iterator for the list.
type listIterator[T interfaces.Equitable[T]] struct {
	n     *node[T]
	start *node[T]
}

// HasNext checks if iterator it has a next element to produce.
func (it *listIterator[T]) HasNext() bool {
	if it.n == nil {
		return false
	}
	return true
}

// Next returns the next element in it. Panics if the iterator has been exhausted.
func (it *listIterator[T]) Next() T {
	if !it.HasNext() {
		panic(NoNextElementError)
	}
	n := it.n
	it.n = it.n.next
	return n.value
}

// Cycle resets the iterator.
func (it *listIterator[T]) Cycle() {
	it.n = it.start
}

// Front returns the element at the front of the list l. Panics if called on an empty list which has no front.
func (l *list[T]) Front() T {
	if l.head != nil {
		return l.head.value
	}
	panic(EmptyListError)
}

// Back returns the element at the back of the list l.  Panics if called on an empty list which has no back.
func (l *list[T]) Back() T {
	if l.tail != nil {
		return l.tail.value
	}
	panic(EmptyListError)
}

// Swap swaps the element at index i and the element at index j. This is done using links. Panics if one of the specified indices
// is out of bounds.
func (l *list[T]) Swap(i, j int) {
	if i < 0 || i > l.len || j < 0 || j > l.len {
		panic(OutOfBoundsError)
	}
	if i >= 0 && i < l.len && j >= 0 && j < l.len {
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

		// Swapping x and y
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

// nodeAt retrives the node at index i in list l. This is for internal use for supporting operations like swap.
func (l *list[T]) nodeAt(i int) *node[T] {
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

// At retrieves the element at index i in list l. If i is out of bounds will panic.
func (l *list[T]) At(i int) T {
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

// Iterator returns a listIterator for the list l.
func (l *list[T]) Iterator() interfaces.Iterator[T] {
	return &listIterator[T]{n: l.head, start: l.head}
}

// AddFront adds an element e to the front of the list l.
func (l *list[T]) AddFront(e T) {
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
func (l *list[T]) AddBack(e T) {
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

// AddAt adds an element to list l at specified index i. If i is out of bounds will panic.
func (l *list[T]) AddAt(i int, e T) {
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
func (l *list[T]) Add(e T) bool {
	l.AddBack(e)
	return true
}

// Set replace the element at index i in the list l with the new element e. Returns the old element at index i.
func (l list[T]) Set(i int, e T) T {
	if i < 0 || i >= l.len {
		panic(OutOfBoundsError)
	}
	n := l.nodeAt(i)
	temp := n.value
	n.value = e
	return temp
}

// AddAll adds all elements from some iterable i  to the list l.
func (l *list[T]) AddAll(i interfaces.Iterable[T]) {
	it := i.Iterator()
	for it.HasNext() {
		l.Add(it.Next())
	}
}

// Len gets the size of the list l.
func (l *list[T]) Len() int {
	return l.len
}

// search traverses the list l looking for element e. For internal use to support operations such as Contains, AddAt and  son on.
func (l *list[T]) search(e T) *node[T] {
	curr := l.head
	for curr != nil {
		if curr.value.Equals(e) {
			return curr
		}
		curr = curr.next
	}
	return nil
}

// Contains checks if element e is part of the list l.
func (l *list[T]) Contains(e T) bool {
	return l.search(e) != nil
}

// RemoveFront removes and returns the front element of the list l. Panics if l is an empty list with no front.
func (l *list[T]) RemoveFront() T {
	if l.len == 0 {
		panic(EmptyListError)
	} else if l.len == 1 {
		n := l.head
		l.head = n.next // subsequent operations are to avoid memory leaks.
		l.tail = nil
		n.next = nil
		n.prev = nil
		l.len -= 1
		return n.value
	} else {
		n := l.head
		l.head = n.next
		n.next = nil
		n.prev = nil
		l.len -= 1
		return n.value
	}
}

// RemoveBack removes and returns the back element of the list l. Panics if l is an empty list with no back.
func (l *list[T]) RemoveBack() T {
	if l.len <= 1 {
		return l.RemoveFront()
	} else {
		n := l.tail
		l.tail = l.tail.prev
		l.tail.next = nil
		l.len -= 1
		return n.value
	}
}

// RemoveAt removes the element at the specified index i. Panics if index is out of bounds.
func (l *list[T]) RemoveAt(i int) T {
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

// removeNode removes the specified node , for internal use for functions such as remove at.
func (l *list[T]) removeNode(curr *node[T]) T {
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
func (l *list[T]) Remove(e T) bool {
	if l.Empty() {
		panic(EmptyListError)
	}
	curr := l.search(e)
	if curr == l.head {
		l.RemoveFront()
		return true
	} else if curr == l.tail {
		l.RemoveBack()
		return true
	} else if curr != nil {
		l.removeNode(curr)
		return true
	}
	return false
}

// RemoveAll removes all the elements from some iterable.
func (l *list[T]) RemoveAll(elements interfaces.Iterable[T]) {
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

// Clear removes all elements from the list.
func (l *list[T]) Clear() {
	for l.head != nil {
		l.RemoveFront()
	}
}

// Equals checks if list l and other list are equal. If they are the same reference/ have same size and elements then they are equal.
// Otherwise they are not equal.
func (l *list[T]) Equals(other List[T]) bool {
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

// Empty checks if the list is empty.
func (l *list[T]) Empty() bool {
	return l.len == 0
}

// Collect collects all elements of the list l into a slice.
func (l *list[T]) Collect() []T {
	data := make([]T, l.len)
	i := 0
	for e := l.head; e != nil; e = e.next {
		data[i] = e.value
		i = i + 1
	}
	return data
}

// traversal for pretty printing purposes.
func (l *list[T]) traversal() string {
	sb := make([]string, 0, 0)
	for e := l.head; e != nil; e = e.next {
		sb = append(sb, fmt.Sprint(e.value))
	}
	return "{" + strings.Join(sb, ", ") + "}"
}

// String string formats for a list l.
func (l *list[T]) String() string {
	return l.traversal()
}

// Map transforms each element of the list l using some function f and returns a new list with transformed elements.
func (l *list[T]) Map(f func(e T) T) List[T] {
	newList := NewList[T]()
	for e := l.head; e != nil; e = e.next {
		newE := f(e.value)
		newList.Add(newE)
	}
	return newList
}

// Filter filters the elements of the list l using some predicate function f and returns new list with elements satisfying filter.
func (l *list[T]) Filter(f func(e T) bool) List[T] {
	newList := NewList[T]()
	for e := l.head; e != nil; e = e.next {
		if f(e.value) {
			newList.Add(e.value)
		}
	}
	return newList
}
