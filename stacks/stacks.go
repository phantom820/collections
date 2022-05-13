// Package stacks provides common utils that stack implementations use.
package stacks

import (
	"errors"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/types"
)

// Errors for operations that may be inapplicable on a queue.
var (
	ErrNoTopElement = errors.New("stack has no top element")
)

// Stack an interface that stack implementation must satisfy.
type Stack[T types.Equitable[T]] interface {
	collections.Collection[T]
	Peek() T // Returns the top element in the stack. Will panic if no top element.
	Pop() T  // Returns and  removes the top element in the stack. Will panic if no top element.
}
