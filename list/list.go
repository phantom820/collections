// Package list provides implementation of a homogeneous linked list (All elements must be same type.)
// forward list (singly linked list with tail pointer) , list (doubly linked list see double_list.go)
package list

import (
	"collections/interfaces"
	"collections/types"
	"errors"
)

// Errors for operations that may be inapplicable on a list.
var (
	EmptyListError   = errors.New("cannot remove from an empty list.")
	OutOfBoundsError = errors.New("index out of bounds.")
)

// _List interface specifying methods that an implementation of a linked list must provide.
type _List[T types.Equitable[T]] interface {
	interfaces.Collection[T]
	AddFront(e T)
	Front() T
	AddBack(e T)
	Back() T
	AddAt(i int, e T)
	RemoveAt(i int) T
	At(i int) T
	Set(i int, e T) T
	Swap(i, j int)
	RemoveFront() T
	RemoveBack() T
}
