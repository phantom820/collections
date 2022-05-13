package lists

import (
	"errors"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/types"
)

// Errors for  operations that may be invalid on a list.
var (
	ErrEmptyList   = errors.New("cannot remove from an empty list")
	ErrOutOfBounds = errors.New("index out of bounds")
)

// List interface which implementations of a linked list must satisfy.
type List[T types.Equitable[T]] interface {
	collections.Collection[T]
	Front() T         // Returns the front element in the list. Will panic if there is no front element.
	RemoveFront() T   // Returns and removes the front element in the list.
	Back() T          // Returns the element at the back of the list. Will panic if no back element.
	RemoveBack() T    // Returns and removes the element at the back of the list. Will panic if no back element.
	Set(i int, e T) T // Replaces the element at the specified index with the new element and returns old element. Will panic if index out of bounds.
	Swap(i, j int)    // Swaps the element at index i with the element at index j. Will panic if one or both indices out of bounds.
}
