package set

// import (
// 	"collections/interfaces"
// 	"collections/tree"
// 	"fmt"
// 	"strings"
// 	"sync"
// )

// // set the actual underlying set.
// type hashSet[T interfaces.Hashable] struct {
// 	*sync.Mutex
// 	data            map[int](tree.RedBlackTree[T])
// 	len             int
// 	rank            int8
// 	loadFactorLimit float32
// 	capacity        int
// }

// // HashSet implements Set interface with some type T and all operations that result in a set
// // will return a HashSet.
// type HashSet[T interfaces.Hashable] interface {
// 	Set[T, HashSet[T]]
// }

// // NewHashSet creates a new empty set with default initial capacity and load factor limit.
// func NewHashSet[T interfaces.Hashable]() HashSet[T] {
// 	m := make(map[int]tree.RedBlackTree[T], Capacity)
// 	s := hashSet[T]{data: m, len: 0, rank: HashSetRank, loadFactorLimit: LoadFactorLimit,
// 		capacity: Capacity}
// 	s.Mutex = new(sync.Mutex)
// 	return &s
// }

// // NewHashSetWith creates a new empty  HashSet with the given initial capacity and load factor limit.
// func NewHashSetWith[T interfaces.Hashable](capacity int, loadFactorLimit float32) HashSet[T] {
// 	m := make(map[int]tree.RedBlackTree[T], capacity)
// 	s := hashSet[T]{data: m, len: 0, rank: HashSetRank, loadFactorLimit: loadFactorLimit,
// 		capacity: capacity}
// 	s.Mutex = new(sync.Mutex)
// 	return &s
// }

// //
// type hashSetIterator[T interfaces.Equitable] struct {
// 	keys     []int
// 	n        int
// 	called   bool
// 	callback func(mapKeys []int) // lazy evaluation only initialize keys and the iterator if it is actually use.
// }

// func (it *hashSetIterator[T]) HasNext() bool {
// 	if it.n < len(it.keys) {
// 		return true
// 	}
// 	return false
// }

// func (it *hashSetIterator[T]) Next() *T {
// 	if
// }

// // String formats the set for printing.
// func (s hashSet[T]) String() string {
// 	sb := make([]string, 0, s.len)
// 	for k := range s.data {
// 		for _, e := range s.data[k].Collect() {
// 			sb = append(sb, fmt.Sprint(e))
// 		}
// 	}
// 	return "{" + strings.Join(sb, ", ") + "}"
// }

// // Len returns the size of the set.
// func (s *hashSet[T]) Len() int {
// 	return s.len
// }

// // Rank returns the rank of the hashSet.
// func (s *hashSet[T]) Rank() int8 {
// 	return s.rank
// }

// // Contains checks if an element is in the set.
// func (s *hashSet[T]) Contains(e T) bool {
// 	key := e.HashCode()
// 	if _, ok := s.data[key]; ok {
// 		return s.data[key].Search(e)
// 	}
// 	return false
// }

// // LoadFactor computes the load factor of the set.
// func (a *hashSet[T]) LoadFactor() float32 {
// 	return float32(a.len) / float32(a.capacity)
// }

// // resize resizes the set once it crosses load factor limit.
// func (s *hashSet[T]) resize() {
// 	newSet := NewHashSetWith[T](s.capacity*2, s.loadFactorLimit)
// 	newSet.AddAll(s)
// 	newSetPtr, ok := newSet.(*hashSet[T])
// 	if ok {
// 		*s = *newSetPtr
// 	}
// }

// // Add adds the element  if its not already in the set.
// func (s *hashSet[T]) Add(e T) bool {
// 	s.Lock()
// 	defer s.Unlock()
// 	if float32(s.len) > s.loadFactorLimit*float32(s.capacity) {
// 		s.resize()
// 	}
// 	key := e.HashCode() % s.capacity
// 	if s.data[key] == nil {
// 		s.data[key] = tree.NewRedBlackTree[T]()
// 		s.data[key].Insert(e)
// 		s.len = s.len + 1
// 		return true
// 	} else {
// 		if !s.Contains(e) {
// 			s.data[key].Insert(e)
// 			s.len = s.len + 1
// 			return true
// 		} else {
// 			ok := s.data[key].Insert(e)
// 			if ok {
// 				s.len = s.len + 1
// 			}
// 			return ok
// 		}
// 	}
// }

// // AddAll adds all elements from an iterable to the set.
// func (a *hashSet[T]) AddAll(b interfaces.Iterable[T]) {
// 	for _, e := range b.Collect() {
// 		a.Add(e)
// 	}
// }

// // Remove remove the element from the set if it is present.
// func (a *hashSet[T]) Remove(e T) bool {
// 	a.Lock()
// 	defer a.Unlock()
// 	key := e.HashCode()
// 	if a.Contains(e) {
// 		c := a.data[key]
// 		c.Delete(e)
// 		if c.Empty() {
// 			delete(a.data, key)
// 		}
// 		a.len = a.len - 1
// 		return true
// 	}
// 	return false
// }

// // RemoveAll deletes all entries of set b from set a.
// func (a *hashSet[T]) RemoveAll(b interfaces.Iterable[T]) {
// 	for _, e := range b.Collect() {
// 		a.Remove(e)
// 	}
// }

// // Clear clears the set and returns to its initial state.
// func (s *hashSet[T]) Clear() {
// 	for k := range s.data {
// 		delete(s.data, k)
// 	}
// 	newSet := NewHashSetWith[T](s.capacity, s.loadFactorLimit)
// 	newSetPtr, ok := newSet.(*hashSet[T])
// 	if ok {
// 		*s = *newSetPtr
// 	}
// }

// // Empty checks if the set is empty.
// func (s *hashSet[T]) Empty() bool {
// 	return s.len == 0
// }

// // Collect collects all elements of the set into a slice.
// func (s *hashSet[T]) Collect() []T {
// 	data := make([]T, s.len)
// 	i := 0
// 	for _, c := range s.data {
// 		for _, e := range c.Collect() {
// 			data[i] = e
// 			i += 1
// 		}
// 	}
// 	return data
// }

// // Map applies some transformation on elements of the set to produce a new set.
// func (a *hashSet[T]) Map(f func(e T) T) HashSet[T] {
// 	b := NewHashSet[T]()
// 	for _, c := range a.data {
// 		for _, e := range c.Collect() {
// 			b.Add(f(e))
// 		}
// 	}
// 	return b
// }

// // Filter filters this set and produces a new set.
// func (a *hashSet[T]) Filter(f func(e T) bool) HashSet[T] {
// 	b := NewHashSet[T]()
// 	for _, c := range a.data {
// 		for _, e := range c.Collect() {
// 			if f(e) {
// 				b.Add(e)
// 			}
// 		}
// 	}
// 	return b
// }
