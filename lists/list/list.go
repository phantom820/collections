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
	list := List[T]{head: nil, len: 0}
	list.AddSlice(elements)
	return &list
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
func (iterator *listIterator[T]) HasNext() bool {
	return iterator.n != nil
}

// Next returns the next element in the iterator it. Will panic if iterator has no next element.
func (iter *listIterator[T]) Next() T {
	if !iter.HasNext() {
		panic(iterator.NoNextElementError)
	}
	n := iter.n
	iter.n = iter.n.next
	return n.value
}

// Cycle resets the iterator.
func (it *listIterator[T]) Cycle() {
	it.n = it.start
}

// Iterator returns an iterator for the list.
func (list *List[T]) Iterator() iterator.Iterator[T] {
	return &listIterator[T]{n: list.head, start: list.head}
}

// Front returns the front element of the list. Will panic if l has no front element.
func (list *List[T]) Front() T {
	if list.head == nil {
		panic(lists.ErrEmptyList)
	}
	return list.head.value
}

// Back returns the back element of the list. Will panic if l has no back element.
func (list *List[T]) Back() T {
	if list.tail != nil {
		return list.tail.value
	}
	panic(lists.ErrEmptyList)
}

// Swap swaps the element at index i and the element at index j. This is done using links. Will panics if one or both of the specified indices
// out of bounds.
func (list *List[T]) Swap(i, j int) {
	if i < 0 || i >= list.len || j < 0 || j >= list.len {
		panic(lists.ErrOutOfBounds)
	} else {
		x := list.nodeAt(i)
		y := list.nodeAt(j)
		if x == list.head {
			list.head = y
		} else if y == list.head {
			list.head = x

		}
		if x == list.tail {
			list.tail = y

		} else if y == list.tail {
			list.tail = x

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
func (list *List[T]) nodeAt(i int) *node[T] {
	j := 0
	var n *node[T]
	for e := list.head; e != nil; e = e.next {
		if j == i {
			n = e
		}
		j++
	}
	return n
}

// At retrieves the element at index i in list. Will panic If i is out of bounds.
func (list *List[T]) At(i int) T {
	if i < 0 || i >= list.len {
		panic(lists.ErrOutOfBounds)
	}
	iterator := list.Iterator()
	j := 0
	var v T
	for iterator.HasNext() {
		e := iterator.Next()
		if j == i {
			v = e
			break
		}
		j++
	}
	return v
}

// AddFront adds element to the front of the list.
func (list *List[T]) AddFront(element T) {
	n := newNode(element)
	if list.head != nil {
		n.next = list.head
		list.head.prev = n
		list.head = n
		list.len++
		return
	}
	list.head = n
	list.tail = n
	list.len++
}

// AddBack adds element to the back of the list.
func (list *List[T]) AddBack(element T) {
	if list.head == nil {
		list.AddFront(element)
		return
	}
	n := newNode(element)
	list.tail.next = n
	n.prev = list.tail
	list.tail = n
	list.len++
}

// AddAt adds element to list at specified index i. Will panic if i is out of bounds.
func (list *List[T]) AddAt(i int, elements T) {
	if i < 0 || i >= list.len {
		panic(lists.ErrOutOfBounds)
	} else if i == 0 {
		list.AddFront(elements)
		return
	} else {
		j := 0
		n := newNode(elements)
		for x := list.head; x != nil; x = x.next {
			if j == i-1 {
				n.prev = x
				n.next = x.next
				x.next = n
				list.len++
				break
			}
			j = j + 1
		}
		return
	}
}

// Add adds elements to the back of the list.
func (list *List[T]) Add(elements ...T) bool {
	for _, e := range elements {
		list.AddBack(e)
	}
	return true
}

// Set replaces the element at index i in the list with the new element. Returns the old element that was at index i.
func (list *List[T]) Set(i int, element T) T {
	if i < 0 || i >= list.len {
		panic(lists.ErrOutOfBounds)
	}
	n := list.nodeAt(i)
	temp := n.value
	n.value = element
	return temp
}

// AddAll adds all elements from some iterable elements to the list.
func (list *List[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		list.Add(it.Next())
	}
}

// AddSlice adds element from a slice to the list.
func (list *List[T]) AddSlice(slice []T) {
	for _, element := range slice {
		list.Add(element)
	}
}

// Len returns the size of the list.
func (list *List[T]) Len() int {
	return list.len
}

// search traverses the list looking for element. For internal use to support operations such as Contains, AddAt and  so on.
func (list *List[T]) search(element T) *node[T] {
	curr := list.head
	for curr != nil {
		if curr.value.Equals(element) {
			return curr
		}
		curr = curr.next
	}
	return nil
}

// Contains checks if element is in the list.
func (list *List[T]) Contains(element T) bool {
	return list.search(element) != nil
}

// RemoveFront removes and returns the front element of the list. Will panic if the list has no front element.
func (list *List[T]) RemoveFront() T {
	if list.len == 0 {
		panic(lists.ErrEmptyList)
	} else if list.len == 1 {
		n := list.head
		list.head = n.next // subsequent operations are to avoid memory leaks.
		list.tail = nil
		v := n.value
		n.next = nil
		n.prev = nil
		n = nil
		list.len -= 1
		return v
	} else {
		n := list.head
		list.head = n.next
		v := n.value
		n.next = nil
		n.prev = nil
		n = nil
		list.len -= 1
		return v
	}
}

// RemoveBack removes and returns the back element of the list. Will panic if l has no back element.
func (list *List[T]) RemoveBack() T {
	if list.len <= 1 {
		return list.RemoveFront()
	} else {
		n := list.tail
		list.tail = list.tail.prev
		list.tail.next = nil
		list.len -= 1
		v := n.value
		n.prev = nil
		n.next = nil
		n = nil
		return v
	}
}

// RemoveAt removes the element at the specified index i. Will panic if index i is out of bounds.
func (list *List[T]) RemoveAt(i int) T {
	if list.Empty() {
		panic(lists.ErrEmptyList)
	} else if i < 0 || i >= list.len {
		panic(lists.ErrOutOfBounds)
	} else if i == 0 {
		return list.RemoveFront()
	} else if i == list.len-1 {
		return list.RemoveBack()
	} else {
		n := list.nodeAt(i)
		return list.removeNode(n)
	}
}

// removeNode removes the specified node. For internal use for functions such as RemoveAt.
func (list *List[T]) removeNode(curr *node[T]) T {
	n := curr.prev
	n.next = curr.next
	n.next.prev = n
	curr.next = nil
	curr.prev = nil
	curr = nil
	list.len -= 1
	return n.value
}

// Remove removes element from the list if its present. This removes the first occurence of the element.
func (list *List[T]) Remove(element T) bool {
	curr := list.search(element)
	if curr == nil {
		return false
	} else if curr == list.head {
		list.RemoveFront()
		return true
	} else if curr == list.tail {
		list.RemoveBack()
		return true
	} else {
		list.removeNode(curr)
		return true
	}
}

// RemoveAll removes all the elements from the list that appear in some iterable elements.
func (list *List[T]) RemoveAll(elements iterator.Iterable[T]) {
	defer func() {
		if r := recover(); r != nil {
			// do nothing just fail safe if l ends up empty from the removals.
		}
	}()
	it := elements.Iterator()
	for it.HasNext() {
		list.Remove(it.Next())
	}
}

// Reverse returns a new list that is the reverse of l. This uses extra memory since we inserting into a new list.
func (list *List[T]) Reverse() *List[T] {
	r := New[T]()
	h := list.head
	for h != nil {
		r.AddFront(h.value)
		h = h.next
	}
	return r
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.len = 0
	// for list.head != nil {
	// 	list.RemoveFront()
	// }
}

// Equals checks if list and list other are equal. If they are the same reference/ have same size and elements then they are equal.
// Otherwise they are not equal.
func (list *List[T]) Equals(other *List[T]) bool {
	if list == other {
		return true
	} else if list.len != other.Len() {
		return false
	} else {
		iter := list.Iterator()
		otherIter := other.Iterator()
		for iter.HasNext() {
			a := iter.Next()
			b := otherIter.Next()
			if !a.Equals(b) {
				return false
			}
		}
		return true
	}
}

// Empty checks if the list is empty.
func (list *List[T]) Empty() bool {
	return list.len == 0
}

// Collect collects all elements of the list into a slice.
func (list *List[T]) Collect() []T {
	data := make([]T, list.len)
	i := 0
	for e := list.head; e != nil; e = e.next {
		data[i] = e.value
		i = i + 1
	}
	return data
}

// traversal for pretty printing purposes.
func (list *List[T]) traversal() string {
	sb := make([]string, 0)
	for e := list.head; e != nil; e = e.next {
		sb = append(sb, fmt.Sprint(e.value))
	}
	return "[" + strings.Join(sb, " ") + "]"
}

// String string format for list.
func (list *List[T]) String() string {
	return list.traversal()
}

// Map transforms each element of the list using a function f and returns a new list with transformed elements.
func (list *List[T]) Map(f func(e T) T) *List[T] {
	newList := New[T]()
	for e := list.head; e != nil; e = e.next {
		newE := f(e.value)
		newList.Add(newE)
	}
	return newList
}

// Filter filters the elements of the list using a predicate function f and returns new list with elements satisfying predicate.
func (list *List[T]) Filter(f func(e T) bool) *List[T] {
	newList := New[T]()
	for e := list.head; e != nil; e = e.next {
		if f(e.value) {
			newList.Add(e.value)
		}
	}
	return newList
}
