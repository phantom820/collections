package list

import (
	"collections/interfaces"
	"collections/iterator"
	"collections/types"
	"fmt"
	"strings"
)

// ForwadList interface to abstract away underlying concrete data. Provides various methods to operate on
// the underlying singly linked list with a tail pointer.
type ForwardList[T types.Equitable[T]] interface {
	_List[T]
	interfaces.Functional[T, ForwardList[T]]
	Equals(other ForwardList[T]) bool
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

// list actual concrete type for a singly linked list.
// head -> head node of the list , tail -> tail node of the list , len -> size of the list.
type forwardList[T types.Equitable[T]] struct {
	head *singleNode[T]
	tail *singleNode[T]
	len  int
}

// NewForwardList creates an empty singy linked list that can store values of type T.
func NewForwardList[T types.Equitable[T]]() ForwardList[T] {
	l := forwardList[T]{head: nil, len: 0}
	return &l
}

// forwadListIterator struct to implement an iterator for a forwardList.
type forwardListIterator[T types.Equitable[T]] struct {
	n     *singleNode[T] // Used for Next() and HasNext().
	start *singleNode[T] // Used to cycle an iterator.
}

// HasNext checks if the iterator it has a next element to produce.
func (it *forwardListIterator[T]) HasNext() bool {
	if it.n == nil {
		return false
	}
	return true
}

// Next returns the next element in the iterator it. Panics if called on an iterator that has been exhausted.
func (it *forwardListIterator[T]) Next() T {
	if !it.HasNext() {
		panic(iterator.NoNextElementError)
	}
	n := it.n
	it.n = it.n.next
	return n.value
}

// Cycle resets the iterator.
func (it *forwardListIterator[T]) Cycle() {
	it.n = it.start
}

// Iterator returns a listIterator for the list l.
func (l *forwardList[T]) Iterator() iterator.Iterator[T] {
	return &forwardListIterator[T]{n: l.head, start: l.head}
}

// Front returns the element at the front of the list l. Will panic if called on an empty list which has no front.
func (l *forwardList[T]) Front() T {
	if l.head != nil {
		return l.head.value
	}
	panic(EmptyListError)
}

// Back returns the element at the back of the list l.  Will panics if called on an empty list which has no back.
func (l *forwardList[T]) Back() T {
	if l.tail != nil {
		return l.tail.value
	}
	panic(EmptyListError)
}

// Swap swaps the element at index i and the element at index j. This is done using links. Will panics if one/both of the specified indices
//  out of bounds.
func (l *forwardList[T]) Swap(i, j int) {
	if i < 0 || i >= l.len || j < 0 || j >= l.len {
		panic(OutOfBoundsError)
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
func (l *forwardList[T]) nodePair(i int) (*singleNode[T], *singleNode[T]) {
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

// nodeAt retrieves the node at index i in list l. This is for internal use for supporting operations like Swap.
func (l *forwardList[T]) nodeAt(i int) *singleNode[T] {
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

// At retrieves the element at index i in list l. Will panic If i is out of bounds.
func (l *forwardList[T]) At(i int) T {
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

// AddFront adds an element e to the front of the list l.
func (l *forwardList[T]) AddFront(e T) {
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

// AddBack adds element e to the back of the list l.
func (l *forwardList[T]) AddBack(e T) {
	if l.head == nil {
		l.AddFront(e)
		return
	}
	n := newSingleNode(e)
	l.tail.next = n
	l.tail = n
	l.len++
}

// AddAt adds an element to list l at specified index i. Will panic if i is out of bounds will panic.
func (l *forwardList[T]) AddAt(i int, e T) {
	if i < 0 || i >= l.len {
		panic(OutOfBoundsError)
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

// Add adds element e to the back of the list l.
func (l *forwardList[T]) Add(e T) bool {
	l.AddBack(e)
	return true
}

// Set replaces the element at index i in the list l with the new element e. Returns the old element at index i.
func (l *forwardList[T]) Set(i int, e T) T {
	if i < 0 || i >= l.len {
		panic(OutOfBoundsError)
	}
	n := l.nodeAt(i)
	temp := n.value
	n.value = e
	return temp
}

// AddAll adds all elements from some iterable elements to the list l.
func (l *forwardList[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		l.Add(it.Next())
	}
}

// AddSlice adds element from a slice s into the list l.
func (l *forwardList[T]) AddSlice(s []T) {
	for _, e := range s {
		l.Add(e)
	}
}

// Len gets the size of the list l.
func (l *forwardList[T]) Len() int {
	return l.len
}

// search traverses the list l looking for element e. For internal use to support operations such as Contains, AddAt and  so on.
func (l *forwardList[T]) search(e T) *singleNode[T] {
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
func (l *forwardList[T]) Contains(e T) bool {
	return l.search(e) != nil
}

// RemoveFront removes and returns the front element of the list l. Will panics if l is an empty list with no front.
func (l *forwardList[T]) RemoveFront() T {
	if l.len == 0 {
		panic(EmptyListError)
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

// RemoveBack removes and returns the back element of the list l. Will panic if l is an empty list with no back.
func (l *forwardList[T]) RemoveBack() T {
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

// RemoveAt removes the element at the specified index i. Will panic if index i is out of bounds.
func (l *forwardList[T]) RemoveAt(i int) T {
	if l.Empty() {
		panic(EmptyListError)
	} else if i < 0 || i >= l.len {
		panic(OutOfBoundsError)
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
func (l *forwardList[T]) removeNode(prev *singleNode[T], curr *singleNode[T]) T {
	prev.next = curr.next
	l.len -= 1
	v := curr.value
	curr.next = nil
	curr = nil
	return v
}

// Remove removes element e from the list l if its present. This removes the first occurence of e.
func (l *forwardList[T]) Remove(e T) bool {
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

// RemoveAll removes all the elements from some iterable.
func (l *forwardList[T]) RemoveAll(elements iterator.Iterable[T]) {
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
func (l *forwardList[T]) Clear() {
	for l.head != nil {
		l.RemoveFront()
	}
}

// Equals checks if list l and other list are equal. If they are the same reference/ have same size and elements then they are equal.
// Otherwise they are not equal.
func (l *forwardList[T]) Equals(other ForwardList[T]) bool {
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
func (l *forwardList[T]) Empty() bool {
	return l.len == 0
}

// Collect collects all elements of the list l into a slice.
func (l *forwardList[T]) Collect() []T {
	data := make([]T, l.len)
	i := 0
	for e := l.head; e != nil; e = e.next {
		data[i] = e.value
		i = i + 1
	}
	return data
}

// traversal for pretty printing purposes.
func (l *forwardList[T]) traversal() string {
	sb := make([]string, 0, 0)
	for e := l.head; e != nil; e = e.next {
		sb = append(sb, fmt.Sprint(e.value))
	}
	return "[" + strings.Join(sb, " ") + "]"
}

// String string formats for a list l.
func (l *forwardList[T]) String() string {
	return l.traversal()
}

// Map transforms each element of the list l using some function f and returns a new list with transformed elements.
func (l *forwardList[T]) Map(f func(e T) T) ForwardList[T] {
	newList := NewForwardList[T]()
	for e := l.head; e != nil; e = e.next {
		newE := f(e.value)
		newList.Add(newE)
	}
	return newList
}

// Filter filters the elements of the list l using some predicate function f and returns new list with elements satisfying predicate.
func (l *forwardList[T]) Filter(f func(e T) bool) ForwardList[T] {
	newList := NewForwardList[T]()
	for e := l.head; e != nil; e = e.next {
		if f(e.value) {
			newList.Add(e.value)
		}
	}
	return newList
}
