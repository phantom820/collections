// Package queues provides common utils that queue implementations must use.
package queues

import (
	"github.com/phantom820/collections"
	"github.com/phantom820/collections/types"
)

// Queue an interface that a queue implementation should satisfy.
type Queue[T types.Equitable[T]] interface {
	collections.Collection[T]
	Front() T       //  Returns the front element of the queue. Will panic if no front element.
	RemoveFront() T // Returns and removes the front element of the queue. Will panic if no front element.
}

// Dequeue an interface that a dequeue implementation should satisfy.
type Dequeue[T types.Equitable[T]] interface {
	collections.Collection[T]
	AddFront(elements ...T) bool // Adds elements to the front of the queue.
	Front() T                    //  Returns the front element of the queue. Will panic if no front element.
	RemoveFront() T              // Returns and removes the front element of the queue. Will panic if no front element.
	Back() T                     //  Returns the back element of the queue. Will panic if no back element.
	RemoveBack() T               //  Returns and removes the item at the back of the queue
}
