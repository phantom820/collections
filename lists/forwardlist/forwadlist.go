// package forwardlist defines singly linked list with a tail pointer.
package forwardlist

import (
	"fmt"
	"math"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/sets"
	"github.com/phantom820/collections/types/optional"
)

type node[T comparable] struct {
	next    *node[T]
	element T
}

// ForwardList represents a mutable disperse ordered collection of elements of type T.
type ForwardList[T comparable] struct {
	head *node[T]
	len  int
	tail *node[T]
}

// New creates a mutable list with the given elements.
func New[T comparable](elements ...T) *ForwardList[T] {
	list := ForwardList[T]{}
	for _, e := range elements {
		list.addBack(e)
	}
	return &list
}

// Of creates an immutable list with the given elements.
func Of[T comparable](elements ...T) ImmutableForwadList[T] {
	list := ForwardList[T]{}
	for _, e := range elements {
		list.addBack(e)
	}
	return ImmutableForwadList[T]{list: list}
}

// AddSlice adds all the elements in the slice to the list.
func (list *ForwardList[T]) AddSlice(s []T) bool {
	for _, e := range s {
		list.addBack(e)
	}
	return true
}

// AddAll adds all of the elements in the iterable to the list.
func (list *ForwardList[T]) AddAll(iterable iterable.Iterable[T]) bool {
	it := iterable.Iterator()
	for it.HasNext() {
		list.Add(it.Next())
	}
	return true
}

// Empty returns true if the list contains no elements.
func (list *ForwardList[T]) Empty() bool {
	return list.len == 0
}

// addFront adds the element to the front of the list.
func (list *ForwardList[T]) addFront(e T) {
	if list.head == nil {
		list.head = &node[T]{element: e}
		list.tail = list.head
		list.len++
		return
	}
	temp := list.head
	list.head = &node[T]{element: e}
	list.head.next = temp
	list.len++
}

// addBack adds the given element to the end of the list.
func (list *ForwardList[T]) addBack(e T) {
	if list.Empty() {
		list.addFront(e)
		return
	}
	temp := &node[T]{element: e}
	list.tail.next = temp
	list.tail = temp
	list.len++
}

// Add appends the specified element to the end of the list.
func (list *ForwardList[T]) Add(e T) bool {
	if list.Empty() {
		list.addFront(e)
		return true
	}
	list.addBack(e)
	return true
}

// chaseIndex chase the given index using a current pointer and previous pointer.
func chaseIndex[T comparable](start *node[T], i int) (*node[T], *node[T]) {
	var prev *node[T]
	curr := start
	j := 0
	for curr != nil {
		if j == i {
			return prev, curr
		}
		prev = curr
		curr = curr.next
		j++
	}
	return nil, nil
}

// AddAt inserts the specified element at the specified index in the list.
func (list *ForwardList[T]) AddAt(i int, e T) {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	} else if i == 0 {
		list.addFront(e)
		return
	} else if i == list.Len()-1 {
		list.Add(e)
		return
	}
	node := node[T]{element: e}
	prev, curr := chaseIndex(list.head, i)
	prev.next = &node
	node.next = curr
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
func (list *ForwardList[T]) At(i int) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	} else if i == 0 {
		return list.head.element
	} else if i == list.len-1 {
		return list.tail.element
	}
	node := at(i, list.head)
	return node.element
}

// Set replaces the element at the specified index in the list with the specified element.
func (list *ForwardList[T]) Set(i int, e T) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	} else if i == 0 {
		temp := list.head.element
		list.head.element = e
		return temp
	} else if i == list.len-1 {
		temp := list.tail.element
		list.tail.element = e
		return temp
	}
	node := at(i, list.head)
	temp := node.element
	node.element = e
	return temp
}

// Len returns the number of elements in the list.
func (list *ForwardList[T]) Len() int {
	return list.len
}

// Clear removes all of the elements from the list.
func (list *ForwardList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.len = 0
}

// Contains returns true if the list contains the specified element.
func (list *ForwardList[T]) Contains(e T) bool {
	for curr := list.head; curr != nil; curr = curr.next {
		if curr.element == e {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of the specified element in the list.
func (list *ForwardList[T]) IndexOf(e T) optional.Optional[int] {
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

// chaseValue chase the given value using a current pointer and previous pointer.
func chaseValue[T comparable](start *node[T], e T) (*node[T], *node[T]) {
	var prev *node[T]
	curr := start
	for curr != nil {
		if curr.element == e {
			return prev, curr
		}
		prev = curr
		curr = curr.next
	}
	return nil, nil
}

// removeFront removes the front node from the list.
func (list *ForwardList[T]) removeFront() T {
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

// removeBack removes the back node from the list, the passed node is the second last.
func (list *ForwardList[T]) removeBack(prev *node[T]) T {
	e := list.tail.element
	list.tail = nil
	prev.next = nil
	list.tail = prev
	list.len = int(math.Max(0, float64(list.len-1)))
	return e
}

// remove removes the current node which is preceded by the prev node.
func (list *ForwardList[T]) remove(prev *node[T], curr *node[T]) {
	if curr == list.head {
		list.removeFront()
		return
	} else if curr == list.tail {
		list.removeBack(prev)
		return
	}
	prev.next = curr.next
	curr.next = nil
	curr = nil
	list.len = int(math.Max(0, float64(list.len-1)))
}

// Remove removes the first occurrence of the specified element from this list.
func (list *ForwardList[T]) Remove(e T) bool {
	prev, curr := chaseValue(list.head, e)
	if curr == nil {
		return false
	}
	list.remove(prev, curr)
	return true
}

// RemoveAt removes the element at the specified index in the list.
func (list *ForwardList[T]) RemoveAt(i int) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	} else if i == 0 {
		return list.removeFront()
	} else if i == list.Len()-1 {
		prev, _ := chaseIndex(list.head, i)
		return list.removeBack(prev)
	}
	prev, curr := chaseIndex(list.head, i)
	e := curr.element
	list.remove(prev, curr)
	return e
}

// RemoveIf removes all of the elements of the list that satisfy the given predicate.
func (list *ForwardList[T]) RemoveIf(f func(e T) bool) bool {
	n := list.len
	var prev *node[T] = nil
	curr := list.head

	// Chase curr and prev pointers and perform normal remove when predicate is satisfied.
	for curr != nil {
		if f(curr.element) {
			next := curr.next
			list.remove(prev, curr)
			curr = next
			continue
		}

		prev = curr
		curr = curr.next
	}

	return n != list.len
}

// RemoveAll removes from the list elements that are contained in the specified iterable.
func (list *ForwardList[T]) RemoveAll(iterable iterable.Iterable[T]) bool {
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

// RetainAll retains only the elements in the list that are contained in the specified collection.
func (list *ForwardList[T]) RetainAll(c collections.Collection[T]) bool {
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

// RemoveSlice removes all of the list elements that are contained in the specified slice.
func (list *ForwardList[T]) RemoveSlice(s []T) bool {
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
func (list *ForwardList[T]) ToSlice() []T {
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
func (list *ForwardList[T]) ForEach(f func(T)) {
	it := list.Iterator()
	for it.HasNext() {
		f(it.Next())
	}
}

// copy copies the values stored in the nodes from start to end  into a new list.
func (list *ForwardList[T]) copy(start, end *node[T]) *ForwardList[T] {
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
func (list *ForwardList[T]) SubList(start int, end int) *ForwardList[T] {
	if start < 0 || start >= list.Len() {
		panic(errors.IndexOutOfBounds(start, list.Len()))
	} else if end < 0 || end > list.Len() {
		panic(errors.IndexOutOfBounds(end, list.Len()))
	} else if start > end {
		panic(errors.IndexBoundsOutOfRange(start, end))
	} else if start == end {
		return New[T]()
	}
	_, startNode := chaseIndex(list.head, start)
	endNode, _ := chaseIndex(list.head, end)
	return list.copy(startNode, endNode)
}

// ImmutableCopy returns an immutable copy of the list.
func (list *ForwardList[T]) ImmutableCopy() ImmutableForwadList[T] {
	copy := ForwardList[T]{}
	list.ForEach(func(e T) {
		copy.Add(e)
	})
	return ImmutableForwadList[T]{copy}
}

// Copy returns a copy of the list.
func (list *ForwardList[T]) Copy() *ForwardList[T] {
	copy := New[T]()
	list.ForEach(func(e T) {
		copy.Add(e)
	})
	return copy
}

// Equals returns true if the list is equivalent to the given list. Two lists are equal if they have the same size
// and contain the same elements in the same order.
func (list *ForwardList[T]) Equals(other collections.List[T]) bool {
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

// Iterator returns an iterator over the elements in the list.
func (list *ForwardList[T]) Iterator() iterator.Iterator[T] {
	return &listIterator[T]{initialized: false, initialize: func() (*node[T], int) { return list.head, list.len }}
}

// listIterator iterator implememantation for [ForwardList].
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
func (list ForwardList[T]) String() string {
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

// locateMid finds the mid of a list using the Tortoise and Hare approach.  For internal use to support sorting.
func locateMid[T comparable](head *node[T]) *node[T] {
	slow := head
	fast := head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// merge combines 2 list that have been sorted by natural ordering of elements. For internal use to support sorting.
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
		sentinel = sentinel.next
	}

	// at the end one of the 2 lists must have been exhauted.
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

// Sort sorts the list using the given less function. if less(a,b) = true then a would be before b in a sorted list.
func (list *ForwardList[T]) Sort(less func(a, b T) bool) {
	if list.Empty() || list.len == 1 {
		return
	}
	head, tail := sort(list.head, less)
	list.head = head
	list.tail = tail
}
