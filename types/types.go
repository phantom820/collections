// Package types contains the type constraints that will be applied on types that can be used with the data structures
// that will be implemented .
package types

// Equitable specifies method that a type must implement in order to be able to check two members of the
// type for equality. The method equals must be an equivalence relation.
type Equitable[T any] interface {
	Equals(other T) bool
}

// Hashable specifies methods that a type must implement in order to be able to generate a hash code for a member of the type.
// If x and y are of type T and satisy x.Equals(y) then they must also sastify x.HashCode() == y.HashCode().
type Hashable[T any] interface {
	HashCode() int
	Equitable[T]
}

// Comparable specifies methods that a type must implement in order to allow ordering between 2 members of that type.
// The Less method must be transitive.
type Comparable[T any] interface {
	Equitable[T]
	Less(other T) bool
}
