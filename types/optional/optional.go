// package optional defines an optional type that can be used to represent value that may or may not be present.
package optional

// Optional represent values that may or may not be present.
type Optional[T any] interface {
	Value() T    // Value returns the value contained by the optional.
	Empty() bool //  Returns true if the optional does not contain any value.
}

// empty defines an empty optional.
type empty[T any] struct{}

func (e *empty[T]) Empty() bool {
	return true
}

func (e *empty[T]) Value() T {
	var zero T
	return zero
}

// optional defines an optional with a value.
type optional[T any] struct {
	value T
}

func (o *optional[T]) Empty() bool {
	return false
}

func (o *optional[T]) Value() T {
	return o.value
}

// Of creates an optional with the given value
func Of[T any](value T) Optional[T] {
	return &optional[T]{value: value}
}

// Empty creates an empty optional.
func Empty[T any]() Optional[T] {
	return &empty[T]{}
}
