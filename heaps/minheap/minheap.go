// Package min heap provides an implementation of a min binary heap.
package minheap

import (
	"fmt"

	"github.com/phantom820/collections/heaps"
	"github.com/phantom820/collections/types"
)

// MinHeap an implementation of a max heap based on a slice.
type MinHeap[T types.Comparable[T]] struct {
	data []T
}

// New creates an empty min heap.
func New[T types.Comparable[T]]() *MinHeap[T] {
	data := make([]T, 0)
	heap := MinHeap[T]{data: data}
	return &heap
}

// heapify preserves heap property.
func (heap *MinHeap[T]) heapify(i int) {
	size := len(heap.data)
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < size && !heap.data[largest].Less(heap.data[l]) {
		largest = l
	}
	if r < size && !heap.data[largest].Less(heap.data[r]) {
		largest = r
	}

	if largest != i {
		temp := heap.data[largest]
		heap.data[largest] = heap.data[i]
		heap.data[largest] = heap.data[i]
		heap.data[i] = temp
		heap.heapify(largest)
	}
}

// Len returns the size of the heap.
func (heap *MinHeap[T]) Len() int {
	return len(heap.data)
}

// Insert insert the element into the heap.
func (heap *MinHeap[T]) Insert(element T) {
	size := heap.Len()
	if size == 0 {
		heap.data = append(heap.data, element)
	} else {
		heap.data = append(heap.data, element)
		size = heap.Len()
		for i := size/2 - 1; i >= 0; i-- {
			heap.heapify(i)
		}
	}
}

// Search checks if the element is in the heap.
func (heap *MinHeap[T]) Search(element T) bool {
	for _, e := range heap.data {
		if e.Equals(element) {
			return true
		}
	}
	return false
}

// Update updates the old element in the heap with the new element.
func (heap *MinHeap[T]) Update(old T, new T) bool {
	ok := heap.Delete(old)
	if ok {
		heap.Insert(new)
	}
	return ok
}

// Delete deletes the element from the heap.
func (heap *MinHeap[T]) Delete(element T) bool {
	size := heap.Len()
	var i int = -1
	for j := 0; j < size; j++ {
		if element.Equals(heap.data[j]) {
			i = j
			break
		}
	}

	if i == -1 {
		return false
	}

	temp := heap.data[i]
	heap.data[i] = heap.data[heap.Len()-1]
	heap.data[size-1] = temp

	heap.data = heap.data[0 : size-1]

	for j := size/2 - 1; j >= 0; j-- {
		heap.heapify(j)
	}

	return true
}

// DeleteTop deletes and returns the element at the top of the heap. Will panic if heap has no top element.
func (heap *MinHeap[T]) DeleteTop() T {
	size := heap.Len()
	i := 0

	if i >= size {
		panic(heaps.ErrEmptyHeap)
	}

	temp := heap.data[i]
	heap.data[i] = heap.data[heap.Len()-1]
	heap.data[size-1] = temp

	heap.data = heap.data[0 : size-1]

	for j := size/2 - 1; j >= 0; j-- {
		heap.heapify(j)
	}

	return temp
}

// Top returns the top element of the heap. Will panic if no top element.
func (heap *MinHeap[T]) Top() T {
	if heap.Empty() {
		panic(heaps.ErrEmptyHeap)
	}
	return heap.data[0]
}

// Clear clears the heap.
func (heap *MinHeap[T]) Clear() {
	heap.data = nil
	heap.data = make([]T, 0)
}

// Empty checks if the heap is empty.
func (heap *MinHeap[T]) Empty() bool {
	return heap.Len() == 0
}

// Collect returns the heap is a slice with no particular ordering.
func (heap *MinHeap[T]) Collect() []T {
	return heap.data
}

// String for printing the heap.
func (heap *MinHeap[T]) String() string {
	return fmt.Sprint(heap.data)
}
