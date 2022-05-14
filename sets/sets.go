package sets

import (
	"github.com/phantom820/collections"
	"github.com/phantom820/collections/types"
)

type Set[T types.Hashable[T]] interface {
	collections.Collection[T]
	// Union(b Set[T]) Set[T]
	// Intersection(b Set[T]) Set[T]
}

// func Union[T types.Hashable[T]](a Set[T], b Set[T]) Set[T] {
// 	c := New()
// }
