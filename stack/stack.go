// Package stack provides list based and slice based implementations of a stack.
package stack

import (
	"errors"
)

// Errors for operations that may be inapplicable on a queue.
var (
	NoTopElementError = errors.New("stack has no top element.")
)

// Stack interface specifying a list of methods a stack implementation is expected to provide.
// type Stack[T types.Equitable[T]] interface {
// 	interfaces.Collection[T]
// 	Peek() T
// 	Pop() T
// }
