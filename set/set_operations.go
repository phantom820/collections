// // Package set provides an implementation of a set motivated by java collections version.
// // The operations here do not modify the input set(s).
package set

// import (
// 	"collections/interfaces"
// )

// // Union performs a union of the sets a and b and returns a new set c. If a & b are the same type then
// // resulting set is of same type otherwise this results in a HashSet.
// func Union[T interfaces.Hashable](a Set[T], b Set[T]) Set[T] {
// 	if a.Rank() == b.Rank() {
// 		c := newSet(a)
// 		c.AddAll(b)
// 		return c
// 	}
// 	c := NewHashSet[T]()
// 	c.AddAll(a)
// 	c.AddAll(b)
// 	return c

// }

// // Difference computes the set difference of a and b, returns a new set c.if a & b are the same type then
// // resulting set is of same type otherwise this results in a HashSet.
// func Difference[T interfaces.Hashable](a Set[T], b Set[T]) Set[T] {
// 	if a.Rank() == b.Rank() {
// 		c := newSet(a)
// 		c.RemoveAll(b)
// 		return c
// 	}
// 	c := NewHashSet[T]()
// 	c.AddAll(a)
// 	c.RemoveAll(b)
// 	return c
// }
