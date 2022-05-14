// Package types implements type constraints for types that can be used with collections. It also contains
// wrappers for string and int,int8,int16,.... primitives to allow them to be used with collections.
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
	Less(other T) bool // Defines a relation on the memebers of type T (what it means for member a < b). This results in what sequence members appear in once sorted.
}

// Int wrapper around int to make it compatible with collections.
type Int int

// HashCode gives the hash code for an Int which is its value.
func (i Int) HashCode() int {
	return int(i)
}

// Equals checks if x equals y.
func (x Int) Equals(y Int) bool {
	return x == y
}

// Less checks if x < y.
func (x Int) Less(y Int) bool {
	return x < y
}

// Int8 wrapper around int8 to make it compatible with collections.
type Int8 int8

// HashCode gives the hash code for an Int8 as its value.
func (i Int8) HashCode() int {
	return int(i)
}

// Equals checks if x equals y.
func (x Int8) Equals(y Int8) bool {
	return x == y
}

// Less checks if x < y.
func (x Int8) Less(y Int8) bool {
	return x < y
}

// Int16 wrapper around int16 to make it compatible with collections.
type Int16 int16

// HashCode gives the hash code for an Int16 as its value.
func (i Int16) HashCode() int {
	return int(i)
}

// Equals checks if x equals y.
func (x Int16) Equals(y Int16) bool {
	return x == y
}

// Less checks if x < y.
func (x Int16) Less(y Int16) bool {
	return x < y
}

// Int32 wrapper around int32 to make it compatible with collections.
type Int32 int32

// HashCode gives the hash code for an Int32 as its value.
func (i Int32) HashCode() int {
	return int(i)
}

// Equals checks if x equals y.
func (x Int32) Equals(y Int32) bool {
	return x == y
}

// Less checks if x < y.
func (x Int32) Less(y Int32) bool {
	return x < y
}

// Int64 wrapper around int64 to make it compatible with collections.
type Int64 int64

// HashCode gives the hash code for an Int64 as its value.
func (i Int64) HashCode() int {
	return int(i)
}

// Equals checks if x equals y.
func (x Int64) Equals(y Int64) bool {
	return x == y
}

// Less checks if x < y.
func (x Int64) Less(y Int64) bool {
	return x < y
}

// Uint wrapper around uint to make it compatible with collections.
type Uint uint

// HashCode gives the hash code for an Uint which is its value.
func (i Uint) HashCode() int {
	return int(i)
}

// Equals checks if x equals y.
func (x Uint) Equals(y Uint) bool {
	return x == y
}

// Less checks if x < y.
func (x Uint) Less(y Uint) bool {
	return x < y
}

// Uint8 wrapper around uint8 to make it compatible with collections.
type Uint8 uint8

// HashCode gives the hash code for an Uint8 as its value.
func (i Uint8) HashCode() int {
	return int(i)
}

// Equals checks if x equals y.
func (x Uint8) Equals(y Uint8) bool {
	return x == y
}

// Less checks if x < y.
func (x Uint8) Less(y Uint8) bool {
	return x < y
}

// Uint16 wrapper around uint16 to make it compatible with collections.
type Uint16 uint16

// HashCode gives the hash code for an Uint16 as its value.
func (i Uint16) HashCode() int {
	return int(i)
}

// Equals checks if x equals y.
func (x Uint16) Equals(y Uint16) bool {
	return x == y
}

// Less checks if x < y.
func (x Uint16) Less(y Uint16) bool {
	return x < y
}

// Uint32 wrapper around uint32 to make it compatible with collections.
type Uint32 uint32

// HashCode gives the hash code for an Uint32 as its value.
func (i Uint32) HashCode() int {
	return int(i)
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
