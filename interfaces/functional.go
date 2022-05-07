package interfaces

// Functional interface for functional programming style on collections.
// Here T is an element of some container C.
// Map applies some transformation to the elements of some collection and returns a collection of the same as original with
// new transfromed elements.
// Filter uses some predicate function to filter a collection and returns a new collection that contains
// elements from old collection that satisfied the predicate.
type Functional[T any, C any] interface {
	Map(f func(e T) T) C       // f:e[T] -> e'[T]
	Filter(f func(e T) bool) C // f:e[T] -> [true,false]
}
