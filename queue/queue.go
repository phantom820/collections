package queue

import "collections/interfaces"

type Queue[T interfaces.Equitable[T]] interface {
	interfaces.Collection[T]
	Front() T
	RemoveFront() T
}
