package optional

import "fmt"

// Optional represent values that may or may not be present.
type Optional[T any] struct {
	value T
	empty bool
}

// Empty returns true if the optional does not contain any value.
func (optional *Optional[T]) Empty() bool {
	return optional.empty
}

// Value returns the value contained by the optional.
func (optional *Optional[T]) Value() T {
	return optional.value
}

// Of creates an optional with the given value.
func Of[T any](value T) Optional[T] {
	return Optional[T]{value: value, empty: false}
}

// Empty creates an empty optional.
func Empty[T any]() Optional[T] {
	return Optional[T]{empty: true}
}

// String returns string representation of optional.
func (optional Optional[T]) String() string {
	if optional.Empty() {
		return "{}"
	}

	return fmt.Sprintf("{%v}", optional.value)
}
