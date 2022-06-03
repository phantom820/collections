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
	list.Add(elements...)
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

// Next yields the next element in the iterator. Will panic if the iterator has no next element.
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

// Front returns the front of the list. Will panic if list has no front element.
func (list *List[T]) Front() T {
	if list.head == nil {
		panic(lists.ErrEmptyList)
	}
	return list.head.value
}

// Back returns the back element of the list.  Will panic if list has no back element.
func (list *List[T]) Back() T {
	if list.tail != nil {
		return list.tail.value
	}
	panic(lists.ErrEmptyList)
}

// Swap swaps the element at index i and the element at index j. This is done using links. Will panic if one/both of the specified indices is
//  out of bounds.
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

// At retrieves the element at index i in the list. Will panic If index is out of bounds.
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

// addFront adds element to the front of the list. For internal use to support AddFront.
func (list *List[T]) addFront(element T) {
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

// AddFront adds elements to the front of the list.
func (list *List[T]) AddFront(elements ...T) {
	for _, element := range elements {
		list.addFront(element)
	}
}

// addBack add an element to the back of the list. For internal use to support AddBack.
func (list *List[T]) addBack(element T) {
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

// AddBack adds elements to the back of the list.
func (list *List[T]) AddBack(elements ...T) {
	for _, element := range elements {
		list.addBack(element)
	}
}

// AddAt adds an element to the list at specified index, all subsequent elements will be shifted right. Will panic if index is out of bounds.
func (list *List[T]) AddAt(i int, element T) {
	if i < 0 || i >= list.len {
		panic(lists.ErrOutOfBounds)
	} else if i == 0 {
		list.AddFront(element)
		return
	}
	j := 0
	n := newNode(element)
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

// Add adds elements to the back of the list.
func (list *List[T]) Add(elements ...T) bool {
	if len(elements) == 0 {
		return false
	}
	for _, e := range elements {
		list.AddBack(e)
	}
	return true
}

// Set replaces the element at the specified index in the list with the new element. Returns the old element that was at the index. Will panic
// if index is out of bounds.
func (list *List[T]) Set(i int, element T) T {
	if i < 0 || i >= list.len {
		panic(lists.ErrOutOfBounds)
	}
	n := list.nodeAt(i)
	temp := n.value
	n.value = element
	return temp
}

// AddAll adds all elements from iterable elements to the list.
func (list *List[T]) AddAll(elements iterator.Iterable[T]) {
	it := elements.Iterator()
	for it.HasNext() {
		list.Add(it.Next())
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

// Contains checks if the element is in the list.
func (list *List[T]) Contains(element T) bool {
	return list.search(element) != nil
}

// RemoveFront removes and returns the front element of the list. Will panic if list has no front element.
func (list *List[T]) RemoveFront() T {
	if list.len == 0 {
		panic(lists.ErrEmptyList)
	} else if list.len == 1 {
		n := list.head
		list.head = n.next
		list.tail = nil
		v := n.value
		n.next = nil
		n.prev = nil
		n = nil
		list.len -= 1
		return v
	}
	n := list.head
	list.head = n.next
	v := n.value
	n.next = nil
	n.prev = nil
	n = nil
	list.len -= 1
	return v
}

// RemoveBack removes and returns the back element of the list. Will panic if the list has no back element.
func (list *List[T]) RemoveBack() T {
	if list.len <= 1 {
		return list.RemoveFront()
	}
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

// RemoveAt removes the element at the specified index in the list. Will panic if index is out of bounds.
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

// removeNode removes the specified node. For internal use to support RemoveAt.
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

// Remove removes elements from the list. Only the first occurence of an element is removed.
func (list *List[T]) Remove(elements ...T) bool {
	n := list.Len()
	for _, element := range elements {
		list.remove(element)
		if list.Empty() {
			break
		}
	}
	return (n != list.Len())
}

// remove removes element from the list if its present. For internal use to support Remove.
func (list *List[T]) remove(element T) bool {
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

// RemoveAll removes all the elements in the list that appear in the iterable.
func (list *List[T]) RemoveAll(iterable iterator.Iterable[T]) {
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

// Reverse reverses the list in place.
func (list *List[T]) Reverse() {

	var temp *node[T] = nil
	current := list.head

	/* swap next and prev for all nodes of
	doubly linked list */
	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}

	/* Before changing the head, check for the cases like empty
	   list and list with only one node */
	if temp != nil {
		prevHead := list.head
		list.head = temp.prev
		list.head.prev = nil
		list.tail = prevHead
	}

}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.len = 0
}

// Equals checks if the list is equal to another list. Two lists are equal if they are the same reference or have the same size and their elements match.
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

// Collect returns a slice containing all the elements in the list.
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

// Map applies a transformation on each element of the list, using the function f and returns a new list with the
// transformed elements.
func (list *List[T]) Map(f func(element T) T) *List[T] {
	newList := New[T]()
	for e := list.head; e != nil; e = e.next {
		newE := f(e.value)
		newList.Add(newE)
	}
	return newList
}

// Filter filters the list using the predicate function  f and returns a new list containing only elements that satisfy the predicate.
func (list *List[T]) Filter(f func(element T) bool) *List[T] {
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
		temp := sentinel
		sentinel = sentinel.next
		sentinel.prev = temp

	}

	// at the end one of the 2 list must have been exhauted.
	for leftHead != nil {
		sentinel.next = leftHead
		leftHead = leftHead.next
		temp := sentinel
		sentinel = sentinel.next
		sentinel.prev = temp
	}

	for rightHead != nil {
		sentinel.next = rightHead
		rightHead = rightHead.next
		temp := sentinel
		sentinel = sentinel.next
		sentinel.prev = temp
	}
	falseHead.next.prev = nil
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
func Sort[T types.Comparable[T]](list *List[T]) {
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
func SortBy[T types.Equitable[T]](list *List[T], less func(a, b T) bool) {
	if list.Empty() || list.len == 1 {
		return
	}
	head, tail := sortBy(list.head, less)
	list.head = head
	list.tail = tail
}
