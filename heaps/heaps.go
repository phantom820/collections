// Package heaps specifies an interface that a heap implementation will fulfill.
package heaps

import (
	"errors"

	"github.com/phantom820/collections/types"
)

var (
	ErrEmptyHeap = errors.New("operation is not allowed on an empty heap")
)

// Heap an interface a heap implementation need to fulfill.
type Heap[T types.Comparable[T]] interface {
	Top() T                   // Returns the element at the top of the heap. Will panic if heap has no top element.
	DeleteTop() T             // Returns and removes the top element in the heap. Will panic if heap has no top element.
	Insert(element T)         // Inserts the element into the heap.
	Update(old T, new T) bool // Updates the old element with the new element in order to change its priority.
	Search(element T) bool    // Searches for the element in the heap.
	Delete(element T) bool    // Removes the element from the heap if its present. Will only remove first element found.
	Len() int                 // Returns the size of the heap.
	Clear()                   // Removes all elements from the heap.
	Collect() []T             // Return elements of the heap as a slice in particular order.
	Empty() bool              // Checks if the heap is empty.

}
