// package linkedlist defines doubly linked list.
package linkedlist

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"unsafe"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/sets"
	"github.com/phantom820/collections/types/optional"
)

type node[T comparable] struct {
	prev    *node[T]
	next    *node[T]
	element T
}

// LinkedList a doubly linked list.
type LinkedList[T comparable] struct {
	head *node[T]
	len  int
	tail *node[T]
}

// New creates a mutable list with the given elements.
func New[T comparable](elements ...T) *LinkedList[T] {
	list := LinkedList[T]{head: nil, len: 0}
	for _, e := range elements {
		list.Add(e)
	}
	return &list
}

// Of creates an immutable list with the given elements.
func Of[T comparable](elements ...T) forwardlist.ImmutableForwadList[T] {
	return forwardlist.Of(elements...)
}

// AddSlice adds all the elements in the slice to the list.
func (list *LinkedList[T]) AddSlice(s []T) bool {
	for _, e := range s {
		list.addBack(e)
	}
	return true
}

// AddAll adds all of the elements in the specified iterable to the set.
func (list *LinkedList[T]) AddAll(iterable iterable.Iterable[T]) bool {
	it := iterable.Iterator()
	for it.HasNext() {
		list.Add(it.Next())
	}
	return true
}

// Empty returns true if the list contains no elements.
func (list *LinkedList[T]) Empty() bool {
	return list.len == 0
}

// addFront add the element to the front of the list.
func (list *LinkedList[T]) addFront(e T) {
	if list.head == nil {
		list.head = &node[T]{element: e}
		list.tail = list.head
		list.len++
		return
	}
	temp := list.head
	list.head = &node[T]{element: e}
	list.head.next = temp
	temp.prev = list.head
	list.len++
}

// addBack appends the given element to the end of the list.
func (list *LinkedList[T]) addBack(e T) {
	if list.Empty() {
		list.addFront(e)
		return
	}
	newTail := &node[T]{element: e}
	list.tail.next = newTail
	newTail.prev = list.tail
	list.tail = newTail
	list.len++
}

// Add appends the specified element to the end of the list.
func (list *LinkedList[T]) Add(e T) bool {
	if list.Empty() {
		list.addFront(e)
		return true
	}
	list.addBack(e)
	return true
}

// chaseIndex chase the given index.
func chaseIndex[T comparable](start *node[T], i int) *node[T] {
	curr := start
	j := 0
	for curr != nil {
		if j == i {
			return curr
		}
		curr = curr.next
		j++
	}
	return curr
}

// AddAt inserts the specified element at the specified index in the list.
func (list *LinkedList[T]) AddAt(i int, e T) {
	if i < 0 || i >= list.Len() && !list.Empty() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	} else if i == 0 {
		list.addFront(e)
		return
	} else if i == list.Len()-1 {
		list.Add(e)
		return
	}
	node := node[T]{element: e}
	curr := at(i, list.head)
	prev := curr.prev
	prev.next = &node
	node.prev = prev
	node.next = curr
	curr.prev = &node
	list.len++
}

// at returns the node at the given index.
func at[T comparable](i int, start *node[T]) *node[T] {
	var curr *node[T]
	j := 0
	for curr := start; curr != nil; curr = curr.next {
		if j == i {
			return curr
		}
		j++
	}
	return curr
}

// At returns the element at the specified index in the list.
func (list *LinkedList[T]) At(i int) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	}
	node := at(i, list.head)
	return node.element
}

// Set replaces the element at the specified index in the list with the specified element.
func (list *LinkedList[T]) Set(i int, e T) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	}
	node := at(i, list.head)
	temp := node.element
	node.element = e
	return temp
}

// Len returns the number of elements in the list.
func (list *LinkedList[T]) Len() int {
	return list.len
}

// Clear removes all of the elements from the list.
func (list *LinkedList[T]) Clear() {
	list.head.next = nil
	list.head = nil
	list.tail.prev = nil
	list.tail = nil
	list.len = 0
}

// Contains returns true if the list contains the specified element.
func (list *LinkedList[T]) Contains(e T) bool {
	for curr := list.head; curr != nil; curr = curr.next {
		if curr.element == e {
			return true
		}
	}
	return false
}

// findValue returns the node with the given value.
func findValue[T comparable](start *node[T], e T) *node[T] {
	curr := start
	for curr != nil {
		if curr.element == e {
			return curr
		}
		curr = curr.next
	}
	return nil
}

// removeFront removes the front node from the list.
func (list *LinkedList[T]) removeFront() T {
	if list.head != list.tail {
		temp := list.head
		e := temp.element
		list.head = list.head.next
		temp.next = nil
		temp = nil
		list.len = int(math.Max(0, float64(list.len-1)))
		return e
	}
	e := list.head.element
	list.head.next = nil
	list.head = nil
	list.tail = nil
	list.len = int(math.Max(0, float64(list.len-1)))
	return e
}

// removeBack removes the back node from the list.
func (list *LinkedList[T]) removeBack() T {
	e := list.tail.element
	// temp := list.tail
	list.tail = list.tail.prev
	// temp = nil
	list.len = int(math.Max(0, float64(list.len-1)))
	return e
}

// remove removes the current node.
func (list *LinkedList[T]) remove(curr *node[T]) {
	if curr == list.head { // the value is in the head node.
		list.removeFront()
		return
	} else if curr == list.tail {
		list.removeBack()
		return
	}
	prev := curr.prev
	prev.next = curr.next
	curr.next = nil
	curr.prev = nil
	curr = nil
	list.len = int(math.Max(0, float64(list.len-1)))
}

// Remove removes the first occurrence of the specified element from this list, if it is present.
func (list *LinkedList[T]) Remove(e T) bool {
	curr := findValue(list.head, e)
	if curr == nil { // the value does not exist in the list.
		return false
	}
	list.remove(curr)
	return true
}

// RemoveAt removes the element at the specified index in the list.
func (list *LinkedList[T]) RemoveAt(i int) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	} else if i == 0 {
		return list.removeFront()
	} else if i == list.Len()-1 {
		return list.removeBack()
	}
	curr := chaseIndex(list.head, i)
	e := curr.element
	list.remove(curr)
	return e
}

// RemoveIf removes all of the elements of the list that satisfy the given predicate.
func (list *LinkedList[T]) RemoveIf(f func(T) bool) bool {
	n := list.len
	curr := list.head

	// chase curr and prev pointers and perform normal remove when predicate.
	for curr != nil {
		if f(curr.element) {
			next := curr.next
			list.remove(curr)
			curr = next
			continue
		}
		curr = curr.next
	}
	return n != list.len
}

// RemoveAll removes from the list all of its elements that are contained in the specified tion.
func (list *LinkedList[T]) RemoveAll(iterable iterable.Iterable[T]) bool {
	if list.Empty() {
		return false
	} else if sets.IsSet(iterable) {
		set := iterable.(collections.Set[T])
		return list.RemoveIf(func(e T) bool {
			return set.Contains(e)
		})
	}
	// Extra memory O(n) for map , time complexity to populate map O(n), worst case
	// we remove at each instance assuming fixed cost k.
	// k + k + k + k + ... + k = n*k  = O(n) (n is size of the list)
	set := make(map[T]struct{})
	it := iterable.Iterator()
	for it.HasNext() {
		set[it.Next()] = struct{}{}
	}
	return list.RemoveIf(func(e T) bool {
		_, ok := set[e]
		return ok
	})
}

// RetainAll retains only the elements in the list that are contained in the specified slice.
func (list *LinkedList[T]) RetainAll(c collections.Collection[T]) bool {
	if list.Empty() {
		return false
	} else if sets.IsSet[T](c) {
		return list.RemoveIf(func(t T) bool { return !c.Contains(t) })
	}
	set := make(map[T]struct{})
	it := c.Iterator()
	for it.HasNext() {
		set[it.Next()] = struct{}{}
	}
	return list.RemoveIf(func(e T) bool {
		_, ok := set[e]
		return !ok
	})
}

// RemoveSlice removes all of the list elements that are also contained in the specified slice.
func (list *LinkedList[T]) RemoveSlice(s []T) bool {
	if list.Empty() {
		return false
	}
	set := make(map[T]struct{})
	for _, e := range s {
		set[e] = struct{}{}
	}
	return list.RemoveIf(func(e T) bool {
		_, ok := set[e]
		return ok
	})
}

// ToSlice returns a slice containing the elements of the list.
func (list *LinkedList[T]) ToSlice() []T {
	data := make([]T, list.len)
	j := 0
	it := list.Iterator()
	for it.HasNext() {
		data[j] = it.Next()
		j++
	}
	return data
}

// ForEach performs the given action for each element of the list.
func (list *LinkedList[T]) ForEach(f func(T)) {
	it := list.Iterator()
	for it.HasNext() {
		f(it.Next())
	}
}

// copy copies the values stored in the nodes from start to end  into a new list.
func (list *LinkedList[T]) copy(start, end *node[T]) *LinkedList[T] {
	copy := New[T]()
	for curr := start; curr != nil; curr = curr.next {
		if curr == end {
			copy.Add(curr.element)
			break
		}
		copy.Add(curr.element)
	}
	return copy
}

// SubList returns a copy of the portion of the list between the specified start and end indices (exclusive).
func (list *LinkedList[T]) SubList(start int, end int) *LinkedList[T] {
	if start < 0 || start >= list.Len() {
		panic(errors.IndexOutOfBounds(start, list.Len()))
	} else if end < 0 || end > list.Len() {
		panic(errors.IndexOutOfBounds(end, list.Len()))
	} else if start > end {
		panic(errors.IndexBoundsOutOfRange(start, end))
	} else if start == end {
		return New[T]()
	}
	startNode := chaseIndex(list.head, start)
	endNode := chaseIndex(list.head, end-1)
	return list.copy(startNode, endNode)
}

// setUnexportedField set an unexported field.
func setUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}

// ImmutableCopy returns an immutable copy of the list.
func (list *LinkedList[T]) ImmutableCopy() forwardlist.ImmutableForwadList[T] {
	if list.Empty() {
		return forwardlist.Of[T]()
	}
	forwardList := forwardlist.New[T]()
	list.ForEach(func(e T) {
		forwardList.Add(e)
	})
	immutableList := forwardlist.Of[T]()
	field := reflect.ValueOf(&immutableList).Elem().FieldByName("list")
	setUnexportedField(field, *forwardList)
	return immutableList
}

// Copy returns a copy of the list.
func (list *LinkedList[T]) Copy() *LinkedList[T] {
	copy := New[T]()
	list.ForEach(func(e T) {
		copy.Add(e)
	})
	return copy
}

// Equals returns true if the list is equivalent to the given list. Two lists are equal if they have the same size
// and contain the same elements in the same order.
func (list *LinkedList[T]) Equals(other collections.List[T]) bool {
	if list == other {
		return true
	} else if list.Len() != other.Len() {
		return false
	}
	it1, it2 := list.Iterator(), other.Iterator()
	_, _ = it1.HasNext(), it2.HasNext() // initializes each iterator.
	for it1.HasNext() {
		if it1.Next() != it2.Next() {
			return false
		}
	}
	return true

}

// IndexOf returns the index of the first occurrence of the specified element in the list.
func (list *LinkedList[T]) IndexOf(e T) optional.Optional[int] {
	j := 0
	it := list.Iterator()
	for it.HasNext() {
		if it.Next() == e {
			return optional.Of(j)
		}
		j++
	}
	return optional.Empty[int]()
}

// Iterator returns an iterator over the elements in the list.
func (list *LinkedList[T]) Iterator() iterator.Iterator[T] {
	return &listIterator[T]{initialized: false, initialize: func() (*node[T], int) { return list.head, list.len }}
}

// listIterator implememantation for [LinkedList].
type listIterator[T comparable] struct {
	initialized bool
	initialize  func() (*node[T], int)
	node        *node[T]
	len         int
	index       int
}

// HasNext returns true if the iterator has more elements.
func (it *listIterator[T]) HasNext() bool {
	if !it.initialized {
		it.node, it.len = it.initialize()
		it.initialized = true
	} else if it.node == nil {
		return false
	}
	return it.node != nil && it.index < it.len
}

// Next returns the next element in the iterator.
func (it *listIterator[T]) Next() T {
	if !it.HasNext() {
		panic(errors.NoSuchElement())
	}
	e := it.node.element
	it.node = it.node.next
	it.index++
	return e
}

// String returns the string representation of the list.
func (list LinkedList[T]) String() string {
	var sb strings.Builder
	if list.Empty() {
		return "[]"
	}
	it := list.Iterator()
	sb.WriteString(fmt.Sprintf("[%v", it.Next()))
	for it.HasNext() {
		sb.WriteString(fmt.Sprintf(" %v", it.Next()))
	}
	sb.WriteString("]")
	return sb.String()
}

// locateMid finds the mid of a list using The Tortoise and The Hare approach.  For internal use to support sorting.
func locateMid[T comparable](head *node[T]) *node[T] {
	slow := head
	fast := head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// merge combines 2 list that have been sorted by using given less function. For internal use to support Sort.
func merge[T comparable](leftHead *node[T], rightHead *node[T], less func(a, b T) bool) (*node[T], *node[T]) {

	falseHead := &node[T]{}
	sentinel := falseHead

	// merge by comparing front of each list and traversing.
	for leftHead != nil && rightHead != nil {
		if less(leftHead.element, rightHead.element) {
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

// sort sorts the list using the given less function. The sorting algorithm is merge sort.
func sort[T comparable](head *node[T], less func(a, b T) bool) (*node[T], *node[T]) {
	if head.next == nil {
		return head, nil
	}
	mid := locateMid(head)
	rightHead := mid.next
	mid.next = nil
	leftHeadSorted, _ := sort(head, less)
	rightHeadSorted, _ := sort(rightHead, less)
	finalHead, finalTail := merge(leftHeadSorted, rightHeadSorted, less)
	return finalHead, finalTail
}

// Sort sorts the list using the given less function. if less(a,b) = true then a would be before b in a sortled list.
func (list *LinkedList[T]) Sort(less func(a, b T) bool) {
	if list.Empty() || list.len == 1 {
		return
	}
	head, tail := sort(list.head, less)
	list.head = head
	list.tail = tail
}
