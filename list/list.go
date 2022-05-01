package list

import (
	"collections/interfaces"
	"errors"
	"fmt"
	"strings"
)

// node for the list.
type node[T interfaces.Equitable[T]] struct {
	prev  *node[T]
	next  *node[T]
	value T
}

// newNode creates a new node/link with value v.
func newNode[T interfaces.Equitable[T]](v T) *node[T] {
	return &node[T]{value: v, prev: nil, next: nil}
}

var (
	EmptyListError   = errors.New("Cannot remove from an empty list.")
	OutOfBoundsError = errors.New("Tried accessing outside indexable range.")
)

// list actual concrete type for doubly linked list.
type list[T interfaces.Equitable[T]] struct {
	head *node[T]
	tail *node[T]
	len  int
}

// List interface to implement a doubly linked list.
type List[T interfaces.Equitable[T]] interface {
	interfaces.Collection[T]
	interfaces.Functional[T, List[T]]
	AddFront(e T)
	Front() T
	AddBack(e T)
	Back() T
	AddAt(i int, e T) error
	At(i int) (T, error)
	Swap(i, j int) error
	RemoveFront() (T, error)
	RemoveBack() (T, error)
}

// NewList creates a new empty list.
func NewList[T interfaces.Equitable[T]]() List[T] {
	l := list[T]{head: nil, len: 0}
	return &l
}

// listIterator a way of iterating over the set.
type listIterator[T interfaces.Equitable[T]] struct {
	n         *node[T]
	start     *node[T]
	exhausted bool
}

// Cycle resets this iterator if its been exhausted.
func (it *listIterator[T]) Cycle() {
	if it.exhausted {
		it.exhausted = false
		it.n = it.start
	}
}

// HasNext checks if there is a next value for iterator.
func (it *listIterator[T]) HasNext() bool {
	if it.n == nil {
		it.exhausted = true
		return false
	}
	return true
}

// Next returns thev next element in iteration. if the iterator has been exhausted just returns zero value.
func (it *listIterator[T]) Next() T {
	if it.exhausted {
		var e T
		return e
	}
	n := it.n
	it.n = it.n.next
	return n.value
}

// Front returns the reference to front element of the list l. if there no front then zero value returned.
func (l *list[T]) Front() T {
	if l.head != nil {
		return l.head.value
	}
	var e T
	return e
}

// Back returns the reference to the tail element of the list l. if there is no back returns zero value.
func (l *list[T]) Back() T {
	if l.tail != nil {
		return l.tail.value
	}
	var e T
	return e
}

// Swap swaps the node at index i and the node at index j. This is done using links as using value might be expensive when values are large {structs}.
func (l *list[T]) Swap(i, j int) error {
	if i < 0 || i > l.len || j < 0 || j > l.len {
		return OutOfBoundsError
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

		// Swapping Node1 and Node2
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
	return nil
}

// nodeAt retrives the node at index i in list l. Otherwise returns nil.
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

// At retrives the value at index i in list l. Otherwise returns the zero value and an error.
func (l *list[T]) At(i int) (T, error) {
	if i < 0 || i >= l.len {
		var v T
		return v, OutOfBoundsError
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
	return v, nil
}

func (l *list[T]) Iterator() interfaces.Iterator[T] {
	return &listIterator[T]{n: l.head, start: l.head}
}

// AddFront append value to front of the list.
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

// AddAt adds an element to list l at specified index i. If is is outside indexable range element will not be added.
func (l *list[T]) AddAt(i int, e T) error {
	if i < 0 || i >= l.len {
		return OutOfBoundsError
	} else if i == 0 {
		l.AddFront(e)
		return nil
	} else if i == l.len-1 {
		l.AddBack(e)
		return nil
	} else {
		j := 0
		n := newNode(e)
		for x := l.head; x != nil; x = x.next {
			if j == i-1 {
				n.prev = x
				n.next = x.next
				x.next = n
				return nil
			}
			j = j + 1
		}
		return nil
	}
}

// Add appends element e to the back of the list l.
func (l *list[T]) Add(e T) bool {
	l.AddBack(e)
	return true
}

// AddAll adds all elements from some iterable i  to the list l.
func (l *list[T]) AddAll(i interfaces.Iterable[T]) {
	for _, e := range i.Collect() {
		l.Add(e)
	}
}

// Len gets the size of the list l.
func (l *list[T]) Len() int {
	return l.len
}

// search traverses the list l looking for element e.
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

// RemoveFront removes and returns the front element of the list l. returns nil if absent.
func (l *list[T]) RemoveFront() (T, error) {
	if l.len == 1 {
		n := l.head
		l.head = n.next
		l.tail = nil
		n.next = nil
		n.prev = nil
		l.len -= 1
		return n.value, nil
	} else if l.head != nil {
		n := l.head
		l.head = n.next
		n.next = nil
		n.prev = nil
		l.len -= 1
		return n.value, nil
	}
	var e T
	return e, EmptyListError
}

// RemoveBack removes and returns the last element of the list l. returns nil if absent.
func (l *list[T]) RemoveBack() (T, error) {
	if l.len <= 1 {
		return l.RemoveFront()
	} else {
		n := l.tail
		l.tail = l.tail.prev
		l.tail.next = nil
		l.len -= 1
		return n.value, nil
	}
}

// Remove removes element e from the list l if its present. This removes the first occurence of e.
func (l *list[T]) Remove(e T) bool {
	curr := l.search(e)
	if curr == l.head {
		l.RemoveFront()
		return true
	} else if curr == l.tail {
		l.RemoveBack()
		return true
	} else if curr != nil {
		n := curr.prev
		n.next = curr.next
		n.next.prev = n
		curr.next = nil
		curr.prev = nil
		curr = nil
		l.len -= 1
		return true
	}
	return false
}

// RemoveAll removes all the elements from some iterable.
func (l *list[T]) RemoveAll(elements interfaces.Iterable[T]) {
	for _, e := range elements.Collect() {
		l.Remove(e)
	}
}

// Clear removes all elements from the list.
func (l *list[T]) Clear() {
	for l.head != nil {
		l.RemoveFront()
	}
}

// Empty checks if the list is empty.
func (l *list[T]) Empty() bool {
	return l.len == 0
}

// Collect collects all elements of the list into a slice.
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

// Filter filters the elements of the list l using some predicate func f and returns new list with elements satisfying filter.
func (l *list[T]) Filter(f func(e T) bool) List[T] {
	newList := NewList[T]()
	for e := l.head; e != nil; e = e.next {
		if f(e.value) {
			newList.Add(e.value)
		}
	}
	return newList
}
