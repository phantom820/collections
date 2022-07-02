// Package forwadlist provides a singly linked list with a tail pointer.
package forwardlist

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types"
)

// ForwardList singly linked list with a tail pointer.
type ForwardList[T types.Equitable[T]] struct {
	head *node[T]
	tail *node[T]
	len  int
}

// New creates a list with the specified elements. If no elements are specified an empty list is created.
func New[T types.Equitable[T]](elements ...T) *ForwardList[T] {
	list := ForwardList[T]{head: nil, len: 0}
	list.Add(elements...)
	return &list
}

// node a link for a singly linked list. Stores a value of some type T along with next pointer. This type is for internal use
// in the implementation of a singly linked list.
type node[T types.Equitable[T]] struct {
	next  *node[T]
	value T
}

// newNode creates a new node with the specified value.
func newNode[T types.Equitable[T]](value T) *node[T] {
	return &node[T]{value: value, next: nil}
}

// forwadListIterator a type implement an iterator for the list.
type forwardListIterator[T types.Equitable[T]] struct {
	n     *node[T] // Used for Next() and HasNext().
	start *node[T] // Used to cycle an iterator.
}

// HasNext checks if the iterator has a next element to yield.
func (it *forwardListIterator[T]) HasNext() bool {
	return it.n != nil
}

// Next yields the next element in the iterator. Will panic if the iterator has no next element.
func (it *forwardListIterator[T]) Next() T {
	if !it.HasNext() {
		panic(errors.ErrNoNextElement())
	}
	n := it.n
	it.n = it.n.next
	return n.value
}

// Cycle resets the iterator.
func (it *forwardListIterator[T]) Cycle() {
	it.n = it.start
}

// Iterator returns an iterator for the list.
func (list *ForwardList[T]) Iterator() iterator.Iterator[T] {
	return &forwardListIterator[T]{n: list.head, start: list.head}
}

// Front returns the front of the list. Will panic if list has no front element.
func (list *ForwardList[T]) Front() T {
	if list.head == nil {
		panic(errors.ErrNoSuchElement(list.len))
	}
	return list.head.value
}

// Back returns the back element of the list.  Will panic if list has no back element.
func (list *ForwardList[T]) Back() T {
	if list.tail == nil {
		panic(errors.ErrNoSuchElement(list.len))
	}
	return list.tail.value
}

// Swap swaps the element at index i and the element at index j. This is done using links. Will panic if one/both of the specified indices is
//  out of bounds.
func (list *ForwardList[T]) Swap(i, j int) {
	if i < 0 || i >= list.len || j < 0 || j >= list.len {
		if i < 0 || i >= list.len {
			panic(errors.ErrIndexOutOfBounds(i, list.len))
		}
		panic(errors.ErrIndexOutOfBounds(j, list.len))
	} else {
		prevX, currX := list.nodePair(i)
		prevY, currY := list.nodePair(j)

		// If x is not head of linked list
		if prevX != nil {
			prevX.next = currY
		} else { // Else make y as new head
			list.head = currY
		}
		// If y is not head of linked list
		if prevY != nil {
			prevY.next = currX
		} else { // Else make x as new head
			list.head = currX
		}
		// Swap next pointers
		temp := currY.next
		currY.next = currX.next
		currX.next = temp
	}
}

// nodePair retrieves a node and the one before it. For internal use to support operation like Swap.
func (list *ForwardList[T]) nodePair(i int) (*node[T], *node[T]) {
	j := 0
	var p *node[T] = nil
	var e = list.head
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
func (list *ForwardList[T]) nodeAt(i int) *node[T] {
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

// At retrieves the element at index i in the list. Will panic If index is out of bounds.
func (list *ForwardList[T]) At(i int) T {
	if i < 0 || i >= list.len {
		panic(errors.ErrIndexOutOfBounds(i, list.len))
	}
	iter := list.Iterator()
	j := 0
	var v T
	for iter.HasNext() {
		e := iter.Next()
		if j == i {
			v = e
			break
		}
		j++
	}
	return v
}

// addFront adds element to the front of the list. For internal use to support AddFront.
func (list *ForwardList[T]) addFront(element T) {
	n := newNode(element)
	if list.head != nil {
		n.next = list.head
		list.head = n
		list.len++
		return
	}
	list.head = n
	list.tail = n
	list.len++
}

// AddFront adds elements to the front of the list.
func (list *ForwardList[T]) AddFront(elements ...T) {
	for _, element := range elements {
		list.addFront(element)
	}
}

// addBack adds element to the back of the list.
func (list *ForwardList[T]) addBack(element T) {
	if list.head == nil {
		list.AddFront(element)
		return
	}
	n := newNode(element)
	list.tail.next = n
	list.tail = n
	list.len++
}

// AddAt adds an element to the list at specified index, all subsequent elements will be shifted right. Will panic if index is out of bounds.
func (list *ForwardList[T]) AddAt(i int, e T) {
	if i < 0 || i >= list.len {
		panic(errors.ErrIndexOutOfBounds(i, list.len))
	} else if i == 0 {
		list.AddFront(e)
		return
	}
	j := 0
	n := newNode(e)
	for x := list.head; x != nil; x = x.next {
		if j == i-1 {
			n.next = x.next
			x.next = n
			list.len++
			break
		}
		j = j + 1
	}
	return

}

// Add adds elements to the back of the list.
func (list *ForwardList[T]) Add(elements ...T) bool {
	if len(elements) == 0 {
		return false
	}
	for _, element := range elements {
		list.addBack(element)
	}
	return true
}

// Set replaces the element at the specified index in the list with the new element. Returns the old element that was at the index. Will panic
// if index is out of bounds.
func (list *ForwardList[T]) Set(i int, element T) T {
	if i < 0 || i >= list.len {
		panic(errors.ErrIndexOutOfBounds(i, list.len))
	}
	n := list.nodeAt(i)
	temp := n.value
	n.value = element
	return temp
}

// AddAll adds all elements from iterable elements to the list.
func (list *ForwardList[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		list.Add(it.Next())
	}
}

// Len returns the size of the list.
func (list *ForwardList[T]) Len() int {
	return list.len
}

// search traverses the list looking for element. For internal use to support operations such as Contains, AddAt and  so on.
func (list *ForwardList[T]) search(e T) *node[T] {
	curr := list.head
	for curr != nil {
		if curr.value.Equals(e) {
			return curr
		}
		curr = curr.next
	}
	return nil
}

// Contains checks if the element is in the list.
func (list *ForwardList[T]) Contains(element T) bool {
	return list.search(element) != nil
}

// RemoveFront removes and returns the front element of the list. Will panic if list has no front element.
func (list *ForwardList[T]) RemoveFront() T {
	if list.len == 0 {
		panic(errors.ErrNoSuchElement(list.len))
	} else if list.len == 1 {
		n := list.head
		list.head = n.next
		list.tail = nil
		list.len -= 1
		v := n.value
		n.next = nil
		n = nil
		return v
	} else {
		n := list.head
		list.head = n.next
		list.len -= 1
		v := n.value
		n.next = nil
		n = nil
		return v
	}
}

// RemoveBack removes and returns the back element of the list. Will panic if the list has no back element.
func (list *ForwardList[T]) RemoveBack() T {
	if list.len <= 1 {
		return list.RemoveFront()
	} else {
		prevN, n := list.nodePair(list.len - 1)
		prevN.next = nil
		list.tail = prevN
		list.len -= 1
		v := n.value
		n.next = nil
		n = nil
		return v
	}
}

// RemoveAt removes the element at the specified index in the list. Will panic if index is out of bounds.
func (list *ForwardList[T]) RemoveAt(i int) T {
	if i < 0 || i >= list.len {
		panic(errors.ErrIndexOutOfBounds(i, list.len))
	} else if i == 0 {
		return list.RemoveFront()
	} else if i == list.len-1 {
		return list.RemoveBack()
	} else {
		prevN, n := list.nodePair(i)
		return list.removeNode(prevN, n)
	}
}

// removeNode removes the specified node curr where prev is the node before it. For internal use to support operations such as RemoveAt.
func (list *ForwardList[T]) removeNode(prev *node[T], curr *node[T]) T {
	prev.next = curr.next
	list.len--
	v := curr.value
	curr.next = nil
	curr = nil
	return v
}

// Remove removes elements from the list. Only the first occurence of an element is removed.
func (list *ForwardList[T]) Remove(elements ...T) bool {
	n := list.len
	for _, element := range elements {
		list.remove(element)
		if list.Empty() {
			break
		}
	}
	return (n != list.len)
}

// remove removes element from the list if its present. This removes the first occurence of the element.
func (list *ForwardList[T]) remove(e T) bool {
	curr := list.search(e)
	if curr == nil {
		return false
	} else if curr == list.head {
		list.RemoveFront()
		return true
	} else if curr == list.tail {
		list.RemoveBack()
		return true
	} else {
		var p *node[T]
		var n = list.head
		for n != nil {
			if n.value.Equals(e) {
				break
			}
			p = n
			n = n.next
		}
		list.removeNode(p, n)
		return true
	}
}

// RemoveAll removes all the elements in the list that appear in the iterable.
func (list *ForwardList[T]) RemoveAll(iterable iterator.Iterable[T]) {
	defer func() {
		if r := recover(); r != nil {
			// do nothing just fail safe if l ends up empty from the removals.
		}
	}()
	it := iterable.Iterator()
	for it.HasNext() {
		list.Remove(it.Next())
	}
}

// Clear removes all elements from the list.
func (list *ForwardList[T]) Clear() {
	for list.head != nil {
		list.RemoveFront()
	}
}

// Equals checks if the list is equal to another list. Two lists are equal if they are the same reference or have the same size and their elements match.
func (list *ForwardList[T]) Equals(other *ForwardList[T]) bool {
	if list == other {
		return true
	} else if list.len != other.Len() {
		return false
	} else {
		it := list.Iterator()
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
func (list *ForwardList[T]) Empty() bool {
	return list.len == 0
}

// Reverse returns a new list that is the reverse of l. This uses extra memory since we inserting into a new list.
func (list *ForwardList[T]) Reverse() {
	// Initialize current, previous and next pointers
	tail := list.head
	current := list.head
	var prev *node[T]
	var next *node[T]

	for current != nil {
		// Store next
		next = current.next
		// Reverse current node's pointer
		current.next = prev
		// Move pointers one position ahead.
		prev = current
		current = next
	}
	list.head = prev
	list.tail = tail
}

// Collect returns a slice containing all the elements in the list.
func (list *ForwardList[T]) Collect() []T {
	data := make([]T, list.len)
	i := 0
	for e := list.head; e != nil; e = e.next {
		data[i] = e.value
		i = i + 1
	}
	return data
}

// traversal for pretty printing purposes.
func (list *ForwardList[T]) traversal() string {
	sb := make([]string, 0)
	for e := list.head; e != nil; e = e.next {
		sb = append(sb, fmt.Sprint(e.value))
	}
	return "[" + strings.Join(sb, " ") + "]"
}

// String string formats for a list.
func (list *ForwardList[T]) String() string {
	return list.traversal()
}

// Map applies a transformation on each element of the list, using the function f and returns a new list with the
// transformed elements.
func (list *ForwardList[T]) Map(f func(element T) T) *ForwardList[T] {
	newList := New[T]()
	for e := list.head; e != nil; e = e.next {
		newE := f(e.value)
		newList.Add(newE)
	}
	return newList
}

// Filter filters the list using the predicate function  f and returns a new list containing only elements that satisfy the predicate.
func (list *ForwardList[T]) Filter(f func(element T) bool) *ForwardList[T] {
	newList := New[T]()
	for e := list.head; e != nil; e = e.next {
		if f(e.value) {
			newList.Add(e.value)
		}
	}
	return newList
}

// locateMid finds the mid of a list using The Tortoise and The Hare approach.  For internal use to support sorting.
func locateMid[T types.Equitable[T]](head *node[T]) *node[T] {
	slow := head
	fast := head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// merge combines 2 list that have been sorted by natural ordering of elements. For internal use to support Sort.
func merge[T types.Comparable[T]](leftHead *node[T], rightHead *node[T]) (*node[T], *node[T]) {

	falseHead := &node[T]{}
	sentinel := falseHead

	// merge by comparing front of each list and traversing.
	for leftHead != nil && rightHead != nil {
		if leftHead.value.Less(rightHead.value) {
			sentinel.next = leftHead
			leftHead = leftHead.next
		} else {
			sentinel.next = rightHead
			rightHead = rightHead.next
		}
		sentinel = sentinel.next
	}

	// at the end one of the 2 list must have been exhauted.
	for leftHead != nil {
		sentinel.next = leftHead
		leftHead = leftHead.next
		sentinel = sentinel.next
	}

	for rightHead != nil {
		sentinel.next = rightHead
		rightHead = rightHead.next
		sentinel = sentinel.next
	}

	return falseHead.next, sentinel
}

// sort sorts the list using natural ordering of elements. The sorting algorithm is merge sort.
func sort[T types.Comparable[T]](head *node[T]) (*node[T], *node[T]) {
	if head.next == nil {
		return head, nil
	}
	mid := locateMid(head)
	rightHead := mid.next
	mid.next = nil
	leftHeadSorted, _ := sort(head)
	rightHeadSorted, _ := sort(rightHead)
	finalHead, finalTail := merge(leftHeadSorted, rightHeadSorted)
	return finalHead, finalTail
}

// Sort sorts the list using natural ordering of elements.
func Sort[T types.Comparable[T]](list *ForwardList[T]) {
	if list.Empty() || list.len == 1 {
		return
	}
	head, tail := sort(list.head)
	list.head = head
	list.tail = tail
}

// mergeBy combines 2 list that have been sorted by a specified function f.
func mergeBy[T types.Equitable[T]](leftHead *node[T], rightHead *node[T], less func(a, b T) bool) (*node[T], *node[T]) {

	falseHead := &node[T]{}
	sentinel := falseHead

	// merge by comparing front of each list and traversing.
	for leftHead != nil && rightHead != nil {
		if less(leftHead.value, rightHead.value) {
			sentinel.next = leftHead
			leftHead = leftHead.next
		} else {
			sentinel.next = rightHead
			rightHead = rightHead.next
		}
		sentinel = sentinel.next
	}

	// at the end one of the 2 list must have been exhauted.
	for leftHead != nil {
		sentinel.next = leftHead
		leftHead = leftHead.next
		sentinel = sentinel.next
	}

	for rightHead != nil {
		sentinel.next = rightHead
		rightHead = rightHead.next
		sentinel = sentinel.next
	}

	return falseHead.next, sentinel
}

// sortBy sorts the list using the specified function less for ordering. The underlying sorting algorithms is merge sort.
func sortBy[T types.Equitable[T]](head *node[T], less func(a, b T) bool) (*node[T], *node[T]) {
	if head.next == nil {
		return head, nil
	}
	mid := locateMid(head)
	rightHead := mid.next
	mid.next = nil
	leftHeadSorted, _ := sortBy(head, less)
	rightHeadSorted, _ := sortBy(rightHead, less)
	finalHead, finalTail := mergeBy(leftHeadSorted, rightHeadSorted, less)
	return finalHead, finalTail
}

// SortBy sorts the list using the function less for comparison of two element . If less(a,b) = true then a comes before b in the sorted list.
func SortBy[T types.Equitable[T]](list *ForwardList[T], less func(a, b T) bool) {
	if list.Empty() || list.len == 1 {
		return
	}
	head, tail := sortBy(list.head, less)
	list.head = head
	list.tail = tail
}
