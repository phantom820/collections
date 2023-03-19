// package collections defines the interfaces for common container data structures. Each container type is an iterable and an iterator can be obtained to access
// its elements. See below for some brief descriptions.
//
//  1. Maps[K comparable, V any] : A Map is an Iterable consisting of pairs of keys and values (also named mappings or associations).
//     1.1 HashMap[K, V] : This is a wrapper around a standard map[K]V i.e has map[K]V as its base type and can be ranged over.
//     1.2 LinkedHashMap[K, V] : This is similar to a HashMap[K, V] however elements are iterated over following their insertion order.
//     1.2 TreeMap[K, V] : A sorted map that stored elements in a sorted order, this backed by a Red Black Tree.
//
// 2.Collection[T comparable] : This is an interface satisfied by
//
//		2.1 List[T] : Linear ordered data structures that support index based operations, this interface is satisfied by the following concrete types.
//		   a. Vector[T] : This is backed by a standard slice.
//		   b. ImmutableVector[T] : An immutable version of a [Vector[T]].
//		   c. ForwardList[T] : A singly linked list with a tail pointer.
//		   d. ImmutableForwardList[T] : An immutable version of a [ForwardList[T]].
//		   e. LinkedList[T] : A doubly linkedList.
//
//	 2.2 Dequeue[T] : This is a double ended queue and can either be backed by a Vector[T] or a LinkedList[T].
//	 2.3 Set[T] : Non-linear data structure that stores unique elements and has quick lookups., this interface is satisfied by the following concrete types.
//		   a. HashSet[T] : A set implementation backed by a [HashMap] with no particular ordering for element iteration.
//		   b. LinkedHashSet[T] : A set implementation backed by a [LinkedHashMap] in which elements are iterated on following their insertion order.
//		   c.TreeSet[T] : A set implementation backed by a [TreeMap] in which elements are iterated on following particular ordering.
package collections

import (
	"reflect"

	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/types/optional"
	"github.com/phantom820/collections/types/pair"
)

type Map[K comparable, V any] interface {
	iterable.Iterable[pair.Pair[K, V]]
	ContainsKey(k K) bool
	ContainsValue(v V, f func(V, V) bool) bool
	ForEach(f func(K, V))
	Clear()
	Get(k K) optional.Optional[V]
	GetIf(f func(K) bool) []V
	Put(k K, v V) optional.Optional[V]
	PutIfAbsent(k K, v V) optional.Optional[V]
	Len() int
	Remove(k K) optional.Optional[V]
	RemoveIf(f func(K) bool) bool
	Keys() []K
	Values() []V
	Empty() bool
}

type Collection[T comparable] interface {
	iterable.Iterable[T]
	Add(e T) bool
	AddAll(iterable iterable.Iterable[T]) bool
	AddSlice(s []T) bool
	Clear()
	Contains(e T) bool
	Empty() bool
	Remove(e T) bool
	RemoveIf(func(T) bool) bool
	RemoveAll(iterable iterable.Iterable[T]) bool
	RemoveSlice(s []T) bool
	RetainAll(c Collection[T]) bool
	ForEach(func(T))
	Len() int
	ToSlice() []T
}

type List[T comparable] interface {
	Collection[T]
	AddAt(i int, e T)
	At(i int) T
	Set(i int, e T) T
	RemoveAt(i int) T
	Equals(list List[T]) bool
	Sort(less func(a, b T) bool)
}

type Set[T comparable] interface {
	Collection[T]
	ContainsAll(iterable iterable.Iterable[T]) bool
}

func IsNil[T comparable](c Collection[T]) bool {
	return c == nil || reflect.ValueOf(c).IsNil()
}
