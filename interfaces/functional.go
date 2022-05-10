package interfaces

// Functional interface for functional programming style on collections.
// Map applies a transformation to each element in a collection and returns a new collection with transformed elements.
// Filter uses a predicate function to filter a collection and returns a new collection that contains only the elements that
// satisfy the predicate.
type Functional[T any, C any] interface {
	Map(f func(e T) T) C       // f:e[T] -> e'[T]
	Filter(f func(e T) bool) C // f:e[T] -> [true,false]
}
