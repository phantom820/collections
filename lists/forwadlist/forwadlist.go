package forwadlist

import (
	"math"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/sets/hashset"
)

type node[T comparable] struct {
	next  *node[T]
	value T
}

type ForwardList[T comparable] struct {
	head *node[T]
	len  int
	tail *node[T]
}

func New[T comparable]() *ForwardList[T] {
	return &ForwardList[T]{head: nil, len: 0}
}

// Of creates a list with the given elements.
func Of[T comparable](elements ...T) ForwardList[T] {
	list := ForwardList[T]{}
	for _, e := range elements {
		list.addBack(e)
	}
	return list
}

// AddSlice adds all the elements in the slice to the list.
func (list *ForwardList[T]) AddSlice(s []T) bool {
	for _, e := range s {
		list.addBack(e)
	}
	return true
}

// Empty returns true if the list contains no elements.
func (list *ForwardList[T]) Empty() bool {
	return list.len == 0
}

// addFront add the element to the front of the list.
func (list *ForwardList[T]) addFront(e T) {
	if list.head == nil {
		list.head = &node[T]{value: e}
		list.tail = list.head
		list.len++
		return
	}
	temp := list.head
	list.head = &node[T]{value: e}
	list.head.next = temp
	list.len++
}

// addBack appends the given element to the end of the list.
func (list *ForwardList[T]) addBack(e T) {
	if list.Empty() {
		list.addFront(e)
		return
	}
	temp := &node[T]{value: e}
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
	node := node[T]{value: e}
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
	}
	node := at(i, list.head)
	return node.value
}

// Set replaces the element at the specified index in the list with the specified element.
func (list *ForwardList[T]) Set(i int, e T) T {
	if i < 0 || i >= list.Len() {
		panic(errors.IndexOutOfBounds(i, list.Len()))
	}
	node := at(i, list.head)
	temp := node.value
	node.value = e
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
		if curr.value == e {
			return true
		}
	}
	return false
}

// chaseValue chase the given value using a current pointer and previous pointer.
func chaseValue[T comparable](start *node[T], e T) (*node[T], *node[T]) {
	var prev *node[T]
	curr := start
	for curr != nil {
		if curr.value == e {
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
func (list *ForwardList[T]) removeBack(prev *node[T]) T {
	if list.head == list.tail {
		return list.removeFront()
	}
	e := list.tail.value
	list.tail = nil
	prev.next = nil
	list.tail = prev
	list.len = int(math.Max(0, float64(list.len-1)))
	return e
}

// remove removes the current node which is prceded by the prev node.
func (list *ForwardList[T]) remove(prev *node[T], curr *node[T]) {
	if curr == list.head { // the value is in the head node.
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

// Remove removes the first occurrence of the specified element from this list, if it is present.
func (list *ForwardList[T]) Remove(e T) bool {
	prev, curr := chaseValue(list.head, e)
	if curr == nil { // the value does not exist in the list.
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
	e := curr.value
	list.remove(prev, curr)
	return e
}

// RemoveIf removes all of the elements of the list that satisfy the given predicate.
func (list *ForwardList[T]) RemoveIf(f func(T) bool) bool {
	n := list.len
	var prev *node[T] = nil
	curr := list.head
	removeFront := false

	// chase curr and prev pointers and perform normal remove when predicate.
	for curr != nil {
		if f(curr.value) {
			if prev == nil { // prev can only be nil for the front node.
				removeFront = true
			} else {
				next := curr.next
				list.remove(prev, curr)
				curr = next
				continue
			}
		}
		prev = curr
		curr = curr.next
	}

	// check if we should remove front.
	if removeFront {
		list.removeFront()
	}

	return n != list.len
}

// RemoveSlice removes all of the list elements that are also contained in the specified slice.
func (list *ForwardList[T]) RemoveSlice(s []T) bool {
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
func (list *ForwardList[T]) ToSlice() []T {
	data := make([]T, list.len)
	j := 0
	for curr := list.head; curr != nil; curr = curr.next {
		data[j] = curr.value
		j++
	}
	return data
}

// Equals returns true if the list is equivalent to the given list. Two lists are equal if they are the same reference or have the same size and contain
// the same elements in the same order.
func (list *ForwardList[T]) Equals(other *ForwardList[T]) bool {
	if list == other {
		return true
	} else if list.Len() != other.Len() {
		return false
	}
	head, otherHead := list.head, other.head
	for curr := head; curr != nil; curr = curr.next {
		if curr.value != otherHead.value {
			return false
		}
		otherHead = otherHead.next
	}
	return true

}
