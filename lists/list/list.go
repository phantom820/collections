// Package list provides implementations of a doubly linked list.
package list

import (
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists"
	"github.com/phantom820/collections/types"

	"fmt"
	"strings"
)

// List an  implementation of a doubly linked list.
type List[T types.Equitable[T]] struct {
	head *node[T]
	tail *node[T]
	len  int
}

// New creates a list with the specified elements. If no elements are specified an empty list is created.
func New[T types.Equitable[T]](elements ...T) *List[T] {
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

// newNode creates a new node with the specified value.
func newNode[T types.Equitable[T]](value T) *node[T] {
	return &node[T]{value: value, prev: nil, next: nil}
}

// listIterator type to implement an iterator for a list.
type listIterator[T types.Equitable[T]] struct {
	n     *node[T] // Used for Next() and HasNext().
	start *node[T] // Used to cycle an iterator.
}

// HasNext checks if the iterator has a next element to yield.
func (it *listIterator[T]) HasNext() bool {
	return it.n != nil
}

// Next returns the next element in the iterator. Will panic if iterator has been exhausted.
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

// Iterator returns an iterator for the list.
func (l *List[T]) Iterator() iterator.Iterator[T] {
	return &listIterator[T]{n: l.head, start: l.head}
}

// Front returns the front element of the list. Will panic if l has no front element.
func (l *List[T]) Front() T {
	if l.head != nil {
		return l.head.value
	}
	panic(lists.ErrEmptyList)
}

// Back returns the back element of the list. Will panic if l has no back element.
func (l *List[T]) Back() T {
	if l.tail != nil {
		return l.tail.value
	}
	panic(lists.ErrEmptyList)
}

// Swap swaps the element at index i and the element at index j. This is done using links. Will panics if one or both of the specified indices
// out of bounds.
func (l *List[T]) Swap(i, j int) {
	if i < 0 || i >= l.len || j < 0 || j >= l.len {
		panic(lists.ErrOutOfBounds)
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

// nodeAt retrieves the node at index i in list. This is for internal use for supporting operations like Swap.
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

// At retrieves the element at index i in list. Will panic If i is out of bounds.
func (l *List[T]) At(i int) T {
	if i < 0 || i >= l.len {
		panic(lists.ErrOutOfBounds)
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

// AddFront adds element to the front of the list.
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

// AddBack adds element to the back of the list.
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

// AddAt adds element to list at specified index i. Will panic if i is out of bounds.
func (l *List[T]) AddAt(i int, e T) {
	if i < 0 || i >= l.len {
		panic(lists.ErrOutOfBounds)
	} else if i == 0 {
		l.AddFront(e)
		return
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

// Add adds elements to the back of the list.
func (l *List[T]) Add(elements ...T) bool {
	for _, e := range elements {
		l.AddBack(e)
	}
	return true
}

// Set replaces the element at index i in the list with the new element. Returns the old element that was at index i.
func (l *List[T]) Set(i int, e T) T {
	if i < 0 || i >= l.len {
		panic(lists.ErrOutOfBounds)
	}
	n := l.nodeAt(i)
	temp := n.value
	n.value = e
	return temp
}

// AddAll adds all elements from some iterable elements to the list.
func (l *List[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		l.Add(it.Next())
	}
}

// AddSlice adds element from a slice s into the list.
func (l *List[T]) AddSlice(s []T) {
	for _, e := range s {
		l.Add(e)
	}
}

// Len returns the size of the list.
func (l *List[T]) Len() int {
	return l.len
}

// search traverses the list looking for element. For internal use to support operations such as Contains, AddAt and  so on.
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

// Contains checks if element is in the list.
func (l *List[T]) Contains(e T) bool {
	return l.search(e) != nil
}

// RemoveFront removes and returns the front element of the list. Will panic if l has no front element.
func (l *List[T]) RemoveFront() T {
	if l.len == 0 {
		panic(lists.ErrEmptyList)
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

// RemoveBack removes and returns the back element of the list. Will panic if l has no back element.
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
		panic(lists.ErrEmptyList)
	} else if i < 0 || i >= l.len {
		panic(lists.ErrOutOfBounds)
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

// Remove removes element from the list if its present. This removes the first occurence of e.
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

// RemoveAll removes all the elements from the list that appear in some iterable elements.
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
	r := New[T]()
	h := l.head
	for h != nil {
		r.AddFront(h.value)
		h = h.next
	}
	return r
}

// Clear removes all elements from the list.
func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
	// for l.head != nil {
	// 	l.RemoveFront()
	// }
}

// Equals checks if list and list other are equal. If they are the same reference/ have same size and elements then they are equal.
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

// Empty checks if the list is empty.
func (l *List[T]) Empty() bool {
	return l.len == 0
}

// Collect collects all elements of the list into a slice.
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

// String string format for list.
func (l *List[T]) String() string {
	return l.traversal()
}

// Map transforms each element of the list using a function f and returns a new list with transformed elements.
func (l *List[T]) Map(f func(e T) T) *List[T] {
	newList := New[T]()
	for e := l.head; e != nil; e = e.next {
		newE := f(e.value)
		newList.Add(newE)
	}
	return newList
}

// Filter filters the elements of the list using a predicate function f and returns new list with elements satisfying predicate.
func (l *List[T]) Filter(f func(e T) bool) *List[T] {
	newList := New[T]()
	for e := l.head; e != nil; e = e.next {
		if f(e.value) {
			newList.Add(e.value)
		}
	}
	return newList
}
