package forwardlist

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists"
	"github.com/phantom820/collections/types"
)

// ForwardList an implementation of a singly linked list.
type ForwardList[T types.Equitable[T]] struct {
	head *singleNode[T]
	tail *singleNode[T]
	len  int
}

// New creates a list with the specified elements. If no elements are specified an empty list is created.
func New[T types.Equitable[T]](elements ...T) *ForwardList[T] {
	l := ForwardList[T]{head: nil, len: 0}
	l.AddSlice(elements)
	return &l
}

// singleNode a link for a singly linked list. Stores a value of some type T along with next pointer. This type is for internal use
// in the implementation of a singly linked list.
type singleNode[T types.Equitable[T]] struct {
	next  *singleNode[T]
	value T
}

// newSingleNode creates a new singleNode with the value v.
func newSingleNode[T types.Equitable[T]](v T) *singleNode[T] {
	return &singleNode[T]{value: v, next: nil}
}

// forwadListIterator struct to implement an iterator for a ForwardList.
type forwardListIterator[T types.Equitable[T]] struct {
	n     *singleNode[T] // Used for Next() and HasNext().
	start *singleNode[T] // Used to cycle an iterator.
}

// HasNext checks if the iterator it has a next element to yield.
func (it *forwardListIterator[T]) HasNext() bool {
	return it.n != nil
}

// Next returns the next element in the iterator it. Will panic if iterator is exhausted.
func (it *forwardListIterator[T]) Next() T {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	n := it.n
	it.n = it.n.next
	return n.value
}

// Cycle resets the iterator it.
func (it *forwardListIterator[T]) Cycle() {
	it.n = it.start
}

// Iterator returns an iterator for the list.
func (l *ForwardList[T]) Iterator() iterator.Iterator[T] {
	return &forwardListIterator[T]{n: l.head, start: l.head}
}

// Front returns the front of the list. Will panic if list has no front element.
func (l *ForwardList[T]) Front() T {
	if l.head != nil {
		return l.head.value
	}
	panic(lists.ErrEmptyList)
}

// Back returns the back element of the list.  Will panic if list has no back element.
func (l *ForwardList[T]) Back() T {
	if l.tail != nil {
		return l.tail.value
	}
	panic(lists.ErrEmptyList)
}

// Swap swaps the element at index i and the element at index j. This is done using links. Will panic if one/both of the specified indices
//  out of bounds.
func (l *ForwardList[T]) Swap(i, j int) {
	if i < 0 || i >= l.len || j < 0 || j >= l.len {
		panic(lists.ErrOutOfBounds)
	} else {
		prevX, currX := l.nodePair(i)
		prevY, currY := l.nodePair(j)

		// If x is not head of linked list
		if prevX != nil {
			prevX.next = currY
		} else { // Else make y as new head
			l.head = currY
		}
		// If y is not head of linked list
		if prevY != nil {
			prevY.next = currX
		} else { // Else make x as new head
			l.head = currX
		}
		// Swap next pointers
		temp := currY.next
		currY.next = currX.next
		currX.next = temp
	}
}

// nodePair retrieves a node and the one before it. For internal use to support operation like Swap.
func (l *ForwardList[T]) nodePair(i int) (*singleNode[T], *singleNode[T]) {
	j := 0
	var p *singleNode[T] = nil
	var e = l.head
	for e != nil {
		if j == i {
			break
		}
		j++
		p = e
		e = e.next
	}
	return p, e
}

// nodeAt retrieves the node at index i in list. This is for internal use for supporting operations like Swap.
func (l *ForwardList[T]) nodeAt(i int) *singleNode[T] {
	j := 0
	var n *singleNode[T]
	for e := l.head; e != nil; e = e.next {
		if j == i {
			n = e
		}
		j++
	}
	return n
}

// At retrieves the element at index i in list. Will panic If i is out of bounds.
func (l *ForwardList[T]) At(i int) T {
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
func (l *ForwardList[T]) AddFront(e T) {
	n := newSingleNode(e)
	if l.head != nil {
		n.next = l.head
		l.head = n
		l.len++
		return
	}
	l.head = n
	l.tail = n
	l.len++
}

// AddBack adds element to the back of the list.
func (l *ForwardList[T]) AddBack(e T) {
	if l.head == nil {
		l.AddFront(e)
		return
	}
	n := newSingleNode(e)
	l.tail.next = n
	l.tail = n
	l.len++
}

// AddAt adds an element to list at specified index i. Will panic if i is out of bounds.
func (l *ForwardList[T]) AddAt(i int, e T) {
	if i < 0 || i >= l.len {
		panic(lists.ErrOutOfBounds)
	} else if i == 0 {
		l.AddFront(e)
		return
	} else if i == l.len-1 {
		l.AddBack(e)
	} else {
		j := 0
		n := newSingleNode(e)
		for x := l.head; x != nil; x = x.next {
			if j == i-1 {
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
func (l *ForwardList[T]) Add(elements ...T) bool {
	for _, e := range elements {
		l.AddBack(e)
	}
	return true
}

// Set replaces the element at index i in the list with the new element. Returns the old element that was at index i.
func (l *ForwardList[T]) Set(i int, e T) T {
	if i < 0 || i >= l.len {
		panic(lists.ErrOutOfBounds)
	}
	n := l.nodeAt(i)
	temp := n.value
	n.value = e
	return temp
}

// AddAll adds all elements from an iterable elements to the list.
func (l *ForwardList[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		l.Add(it.Next())
	}
}

// AddSlice adds element from a slice s into the list.
func (l *ForwardList[T]) AddSlice(s []T) {
	for _, e := range s {
		l.Add(e)
	}
}

// Len returns the size of the list.
func (l *ForwardList[T]) Len() int {
	return l.len
}

// search traverses the list looking for element. For internal use to support operations such as Contains, AddAt and  so on.
func (l *ForwardList[T]) search(e T) *singleNode[T] {
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
func (l *ForwardList[T]) Contains(e T) bool {
	return l.search(e) != nil
}

// RemoveFront removes and returns the front element of the list. Will panic if l has no front element.
func (l *ForwardList[T]) RemoveFront() T {
	if l.len == 0 {
		panic(lists.ErrEmptyList)
	} else if l.len == 1 {
		n := l.head
		l.head = n.next // subsequent operations are to avoid memory leaks.
		l.tail = nil
		l.len -= 1
		v := n.value
		n.next = nil
		n = nil
		return v
	} else {
		n := l.head
		l.head = n.next
		l.len -= 1
		v := n.value
		n.next = nil
		n = nil
		return v
	}
}

// RemoveBack removes and returns the back element of the list. Will panic if l has no back element.
func (l *ForwardList[T]) RemoveBack() T {
	if l.len <= 1 {
		return l.RemoveFront()
	} else {
		prevN, n := l.nodePair(l.len - 1)
		prevN.next = nil
		l.tail = prevN
		l.len -= 1
		v := n.value
		n.next = nil
		n = nil
		return v
	}
}

// RemoveAt removes the element at the specified index i in the list. Will panic if index i is out of bounds.
func (l *ForwardList[T]) RemoveAt(i int) T {
	if l.Empty() {
		panic(lists.ErrEmptyList)
	} else if i < 0 || i >= l.len {
		panic(lists.ErrOutOfBounds)
	} else if i == 0 {
		return l.RemoveFront()
	} else if i == l.len-1 {
		return l.RemoveBack()
	} else {
		prevN, n := l.nodePair(i)
		return l.removeNode(prevN, n)
	}
}

// removeNode removes the specified node curr where prev is the node before it. For internal use for functions such as remove at.
func (l *ForwardList[T]) removeNode(prev *singleNode[T], curr *singleNode[T]) T {
	prev.next = curr.next
	l.len -= 1
	v := curr.value
	curr.next = nil
	curr = nil
	return v
}

// Remove removes element from the list if its present. This removes the first occurence of e.
func (l *ForwardList[T]) Remove(e T) bool {
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
		var p *singleNode[T]
		var n = l.head
		for n != nil {
			if n.value.Equals(e) {
				break
			}
			p = n
			n = n.next
		}
		l.removeNode(p, n)
		return true
	}
}

// RemoveAll removes all the elements from list that appear in iterable elements.
func (l *ForwardList[T]) RemoveAll(elements iterator.Iterable[T]) {
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
func (l *ForwardList[T]) Clear() {
	for l.head != nil {
		l.RemoveFront()
	}
}

// Equals checks if list and other list are equal. If they are the same reference/ have same size and elements then they are equal.
// Otherwise they are not equal.
func (l *ForwardList[T]) Equals(other *ForwardList[T]) bool {
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
func (l *ForwardList[T]) Empty() bool {
	return l.len == 0
}

// Reverse returns a new list that is the reverse of l. This uses extra memory since we inserting into a new list.
func (l *ForwardList[T]) Reverse() *ForwardList[T] {
	r := New[T]()
	h := l.head
	for h != nil {
		r.AddFront(h.value)
		h = h.next
	}
	return r
}

// Collect collects all elements of the list into a slice.
func (l *ForwardList[T]) Collect() []T {
	data := make([]T, l.len)
	i := 0
	for e := l.head; e != nil; e = e.next {
		data[i] = e.value
		i = i + 1
	}
	return data
}

// traversal for pretty printing purposes.
func (l *ForwardList[T]) traversal() string {
	sb := make([]string, 0, 0)
	for e := l.head; e != nil; e = e.next {
		sb = append(sb, fmt.Sprint(e.value))
	}
	return "[" + strings.Join(sb, " ") + "]"
}

// String string formats for a list.
func (l *ForwardList[T]) String() string {
	return l.traversal()
}

// Map transforms each element of the list using a function f and returns a new list with transformed elements.
func (l *ForwardList[T]) Map(f func(e T) T) *ForwardList[T] {
	newList := New[T]()
	for e := l.head; e != nil; e = e.next {
		newE := f(e.value)
		newList.Add(newE)
	}
	return newList
}

// Filter filters the elements of the list using a predicate function f and returns new list with elements satisfying predicate.
func (l *ForwardList[T]) Filter(f func(e T) bool) *ForwardList[T] {
	newList := New[T]()
	for e := l.head; e != nil; e = e.next {
		if f(e.value) {
			newList.Add(e.value)
		}
	}
	return newList
}
