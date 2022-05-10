// Package queue provides an interface that a queue implementation should satisfy. Different implementations of a queue
// are present in  this package , ListQueue (based on ForwardList ).
package queue

import (
	"errors"
)

// Errors for operations that may be inapplicable on a queue.
var (
	NoFrontElementError = errors.New("queue has no front element.")
)

// Queue interface specifying a list of methods a queue implementation is expected to provide.
// type Queue[T types.Equitable[T]] interface {
// 	interfaces.Collection[T]
// 	Front() T
// 	RemoveFront() T
// }
