// Package stack provides an interface that stack implementation must satisfy.
package stack

import (
	"collections/interfaces"
	"collections/types"
	"errors"
)

// Errors for operations that may be inapplicable on a queue.
var (
	NoTopElementError = errors.New("stack has no top element.")
)

type Stack[T types.Equitable[T]] interface {
	interfaces.Collection[T]
	Peek() T
	Pop() T
}
