// Package queues provides common utils that queue implementations must use.
package queues

import (
	"errors"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/types"
)

// Errors for operations that may be inapplicable on a queue.
var (
	ErrNoFrontElement = errors.New("queue has no front element")
)

type Queue[T types.Equitable[T]] interface {
	collections.Collection[T]
	Front() T       //  Returns the front element of the queue. Will panic if no front element.
	RemoveFront() T // Returns and removes the front element of the queue. Will panic if no front element.
}
