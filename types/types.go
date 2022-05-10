// Package types implements type constraints for types that can be used with collections. It also contains
// wrappers for string and int primitives to allow them to be used with collections.
package types

// Equitable specifies a method that a type must implement in order to be able to check two members of the
// type for equality. The method Equals must be an equivalence relation.
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

// Integer wrapper around int to make it compatible with collections.
type Integer int

// HashCode gives the hash code for an integer as its value.
func (i Integer) HashCode() int {
	return int(i)
}

// Equals checks if Integer x equals Integer y.
func (x Integer) Equals(y Integer) bool {
	return x == y
}

// Less checks if Integer x is less than Integer y.
func (x Integer) Less(y Integer) bool {
	return x < y
}

// String wrapper around string to make it compatible with collections.
type String string

// HashCode generates a hash code for a given String.
func (s String) HashCode() int {
	const (
		m = 1e9 + 9 // To be used for modulation.
		p = 53      // To be used as prime.
	)
	p_pow := 1
	runes := []rune(s)
	code := 0
	for _, r := range runes {
		code = (code + p_pow*(int(r)+1)) % m
		p_pow = (p_pow * p) % m
	}
	return code
}

// Equals check if String s1 and String s2 are equal.
func (s String) Equals(other String) bool {
	return s == other
}

// Less checks if String s is less than String other (lexographical comparison).
func (s String) Less(other String) bool {
	return s < other
}
