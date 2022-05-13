// Package stacks provides common utils that stack implementations use.
package stacks

import (
	"errors"
)

// Errors for operations that may be inapplicable on a queue.
var (
	ErrNoTopElement = errors.New("stack has no top element")
)
