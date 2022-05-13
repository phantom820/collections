// Package queues provides common utils that queue implementations must use.
package queues

import (
	"errors"
)

// Errors for operations that may be inapplicable on a queue.
var (
	ErrNoFrontElement = errors.New("queue has no front element")
)
