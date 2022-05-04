package interfaces

// Functional interface for functional programming style on collections.
// Here T is an element of some container C.
type Functional[T any, C any] interface {
	Map(f func(e T) T) C
	Filter(f func(e T) bool) C
}
