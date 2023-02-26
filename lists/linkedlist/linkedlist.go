package linkedlist

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"unsafe"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/sets/hashset"
)

type node[T comparable] struct {
	prev  *node[T]
	next  *node[T]
	value T
}

type LinkedList[T comparable] struct {
	head *node[T]
	len  int
	tail *node[T]
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, len: 0}
}

// Of creates a list with the given elements.
func Of[T comparable](elements ...T) LinkedList[T] {
	list := LinkedList[T]{}
	for _, e := range elements {
		list.addBack(e)
	}
	return list
}

// AddSlice adds all the elements in the slice to the list.
func (list *LinkedList[T]) AddSlice(s []T) bool {
	for _, e := range s {
		list.addBack(e)
	}
	return true
}

// AddAll adds all of the elements in the specified iterable to the set.
func (list *LinkedList[T]) AddAll(iterable collections.Iterable[T]) bool {
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
		list.head = &node[T]{value: e}
		list.tail = list.head
		list.len++
		return
	}
	temp := list.head
	list.head = &node[T]{value: e}
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
	temp := &node[T]{value: e}
	list.tail.next = temp
	temp.prev = list.tail
	list.tail = temp
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
func (list *LinkedList[T]) AddAt(i int, e T) {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	} else if i == 0 {
		list.addFront(e)
		return
	} else if i == list.Len()-1 {
		list.Add(e)
		return
	}
	node := node[T]{value: e}
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
	return node.value
}

// Set replaces the element at the specified index in the list with the specified element.
func (list *LinkedList[T]) Set(i int, e T) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	}
	node := at(i, list.head)
	temp := node.value
	node.value = e
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
		if curr.value == e {
			return true
		}
	}
	return false
}

// findValue returns the node with the given value.
func findValue[T comparable](start *node[T], e T) *node[T] {
	curr := start
	for curr != nil {
		if curr.value == e {
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
		e := temp.value
		list.head = list.head.next
		temp.next = nil
		temp = nil
		list.len = int(math.Max(0, float64(list.len-1)))
		return e
	}
	e := list.head.value
	list.head.next = nil
	list.head = nil
	list.tail = nil
	list.len = int(math.Max(0, float64(list.len-1)))
	return e
}

// removeBack removes the back node from the list.
func (list *LinkedList[T]) removeBack(curr *node[T]) T {
	e := list.tail.value
	list.tail = curr.prev
	curr.prev = nil
	curr.next = nil
	curr = nil
	list.len = int(math.Max(0, float64(list.len-1)))
	return e
}

// remove removes the current node.
func (list *LinkedList[T]) remove(curr *node[T]) {
	if curr == list.head { // the value is in the head node.
		list.removeFront()
		return
	} else if curr == list.tail {
		list.removeBack(curr)
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
		prev, _ := chaseIndex(list.head, i)
		return list.removeBack(prev)
	}
	_, curr := chaseIndex(list.head, i)
	e := curr.value
	list.remove(curr)
	return e
}

// RemoveIf removes all of the elements of the list that satisfy the given predicate.
func (list *LinkedList[T]) RemoveIf(f func(T) bool) bool {
	n := list.len
	curr := list.head

	// chase curr and prev pointers and perform normal remove when predicate.
	for curr != nil {
		if f(curr.value) {
			next := curr.next
			list.remove(curr)
			curr = next
			continue
		}
		curr = curr.next
	}
	return n != list.len
}

// RemoveAll removes from the list all of its elements that are contained in the specified collection.
func (list *LinkedList[T]) RemoveAll(iterable collections.Iterable[T]) bool {
	if list.Empty() {
		return false
	}
	// introduce a set so we can ensure the lookups fast, we only want to do a single linear pass in removing elements
	// so the algorithm here is O(n) i.e 2 linear passes.
	set := hashset.New[T]()
	it := iterable.Iterator()
	for it.HasNext() {
		set.Add(it.Next())
	}
	return list.RemoveIf(func(t T) bool { return set.Contains(t) })
}

// RetainAll retains only the elements in the list that are contained in the specified collection.
func (list *LinkedList[T]) RetainAll(c collections.Collection[T]) bool {
	if list.Empty() {
		return false
	}
	// create a predicate that removes elements that are not in the passed collection.
	// performance here is mainly affected by how the given collection performs with contains.
	return list.RemoveIf(func(t T) bool { return !c.Contains(t) })
}

// RemoveSlice removes all of the list elements that are also contained in the specified slice.
func (list *LinkedList[T]) RemoveSlice(s []T) bool {
	if list.Empty() {
		return false
	}
	// introduce a set so we can make the lookups fast, also passing a collection here introduces
	// uncertainty about performance of contains so we just need an iterable and enforce the set.
	set := hashset.New[T]()
	set.AddSlice(s)
	return list.RemoveIf(func(t T) bool { return set.Contains(t) })
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
			copy.Add(curr.value)
			break
		}
		copy.Add(curr.value)
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
	_, startNode := chaseIndex(list.head, start)
	endNode, _ := chaseIndex(list.head, end)
	return list.copy(startNode, endNode)
}

// setUnexportedFieldToAvoid set an unexported field.
func setUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}

// ImmutableCopy returns an immutable copy of the list.
func (list *LinkedList[T]) ImmutableCopy() forwardlist.ImmutableForwadList[T] {
	forwardList := forwardlist.Of[T]()
	list.ForEach(func(e T) {
		forwardList.Add(e)
	})
	immutableList := forwardlist.ImmutableOf[T]()
	// using reflection here to avoid an unnecessary memory allocations.
	field := reflect.ValueOf(&immutableList).Elem().FieldByName("list")
	setUnexportedField(field, forwardList)
	return immutableList
}

// Copy returns a copy of the list.
func (list *LinkedList[T]) Copy() *LinkedList[T] {
	copy := Of[T]()
	list.ForEach(func(e T) {
		copy.Add(e)
	})
	return &copy
}

// Equals returns true if the list is equivalent to the given list. Two lists are equal if they have the same size
// and contain the same elements in the same order.
func (list *LinkedList[T]) Equals(other *LinkedList[T]) bool {
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
func (list *LinkedList[T]) Iterator() collections.Iterator[T] {
	return &iterator[T]{initialized: false, initialize: func() (*node[T], int) { return list.head, list.len }}
}

// iterator implememantation for [ForwardList].
type iterator[T comparable] struct {
	initialized bool
	initialize  func() (*node[T], int)
	node        *node[T]
	len         int
	index       int
}

// HasNext returns true if the iterator has more elements.
func (it *iterator[T]) HasNext() bool {
	if !it.initialized {
		it.node, it.len = it.initialize()
		it.initialized = true
	} else if it.node == nil {
		return false
	}
	return it.node != nil && it.index < it.len
}

// Next returns the next element in the iterator.
func (it *iterator[T]) Next() T {
	if !it.HasNext() {
		panic("iterator things shoould panic here")
	}
	e := it.node.value
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
