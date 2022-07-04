package treeset

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/testutils"
	"github.com/phantom820/collections/types"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	s := New[types.Int]()
	assert.Equal(t, true, s.Empty())

	// Case 1 : Add with no args.
	assert.Equal(t, false, s.Add())

	// Case 2 : Add to an empty set.
	s.Add(1)
	assert.Equal(t, 1, s.Len())
	assert.Equal(t, false, s.Add(1))
	assert.Equal(t, 1, s.Len())
	s.Add(2)
	assert.Equal(t, 2, s.Len())
	assert.Equal(t, true, s.Contains(1))
	assert.Equal(t, true, s.Contains(2))

	// Case 3 : Add multiple elements from another iterable.
	s = New[types.Int]()
	l := list.New[types.Int]()
	for i := 0; i < 10; i++ {
		l.Add(types.Int(i))
	}

	s.AddAll(l)
	assert.Equal(t, l.Len(), s.Len())

	// should contain all the added elements
	it := l.Iterator()
	for it.HasNext() {
		assert.Equal(t, true, s.Contains(it.Next()))
	}

	// Case 4 : Adding a slice should work accordingly.
	s.Clear()

	sl := []types.Int{1, 1, 2, 3, 4, 5}
	s.Add(sl...)

	sm := []types.Int{1, 2, 3, 4, 5}
	assert.Equal(t, true, testutils.EqualSlices(sm, s.Collect()))

	// Case 5 : Clear and add a vast colelction of values
	s.Clear()
	slice := []types.Int{}
	for i := 50; i >= 0; i-- {
		s.Add(types.Int(i))
		slice = append(slice, types.Int(50-i))
	}
	assert.Equal(t, true, testutils.EqualSlices(slice, s.Collect()))

}

func TestIterator(t *testing.T) {

	s := New[types.Int]()

	// Case 1 : Next on iterator for an empty set should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoNextElement, r.(errors.Error).Code())
			}
		}()
		it := s.Iterator()
		it.Next()
	})
	s.Add(1, 2, 3, 4, 5)

	a := s.Collect()
	b := make([]types.Int, 0)
	it := s.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}

	assert.Equal(t, true, testutils.EqualSlices(a, b))

	s.Clear()
	assert.Equal(t, true, s.Empty())

}

func TestIteratorConcurrentModification(t *testing.T) {

	s := New[types.String]()
	for i := 1; i <= 20; i++ {
		s.Add(types.String(strconv.Itoa(i)))
	}

	// Recovery for concurrent modifications.
	recovery := func() {
		if r := recover(); r != nil {
			assert.Equal(t, errors.ConcurrentModification, r.(*errors.Error).Code())
		}
	}
	// Case 1 : Put.
	it := s.Iterator()
	t.Run("Add while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			s.Add(types.String("D"))
			it.Next()
		}
	})
	// Case 2 : Remove.
	it = s.Iterator()
	t.Run("Remove while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			s.Remove(types.String("D"))
			it.Next()
		}
	})
	// Case 3 : RemoveIf.
	it = s.Iterator()
	t.Run("RemoveIf while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			s.RemoveIf(func(element types.String) bool { return element == "" })
			it.Next()
		}
	})
	// Case 3 : RetainAll.
	it = s.Iterator()
	t.Run("RetainAll while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			s.RetainAll(New[types.String]())
			it.Next()
		}
	})
	// Case 4 : Clear.
	it = s.Iterator()
	t.Run("Clear while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			s.Clear()
			it.Next()
		}
	})

}

func TestRemove(t *testing.T) {

	s := New[types.Int]()

	assert.Equal(t, false, s.Remove(1))
	s.Add(1)
	assert.Equal(t, true, s.Remove(1))
	assert.Equal(t, 0, s.Len())

	l := list.New[types.Int]()
	for i := 1; i <= 10; i++ {
		s.Add(types.Int(i))
		l.Add(types.Int(i))
	}

	assert.Equal(t, 10, s.Len())
	s.RemoveAll(l)
	assert.Equal(t, 0, s.Len())

}

func TestRemoveIf(t *testing.T) {

	s := New[types.Int]()

	// Case 1 : RemoveIf on an empty set.
	assert.Equal(t, false, s.RemoveIf(func(element types.Int) bool { return element%2 == 0 }))

	// Case 2 : RemoveIf on a set with elements but none satisfy predicates.
	for i := 1; i <= 2000; i++ {
		s.Add(types.Int(i))
	}
	assert.Equal(t, false, s.RemoveIf(func(element types.Int) bool { return element > 2000 }))

	// Case 3 : RemoveIf on a set with elements and some satisfy predicate.
	assert.Equal(t, true, s.RemoveIf(func(element types.Int) bool { return element%2 == 0 }))
	assert.Equal(t, 1000, s.Len())

}

// TestUnion covers tests for Union operation.
func TestUnion(t *testing.T) {

	a := New[types.Int]()
	b := New[types.Int]()

	// Case 1 : Union of empty sets should return empty.
	c := a.Union(b)
	assert.Equal(t, true, c.Empty())

	// Case 2 : Union of populatet set and empty set should return popuilated set.
	a.Add(1, 2, 3)
	c = a.Union(b)
	assert.Equal(t, true, c.Equals(a))

	// Case 3 : Union of populated sets.
	b.Add(1, 2, 4, 5, 6)
	d := New[types.Int](1, 2, 3, 4, 5, 6)

	assert.Equal(t, true, d.Equals(a.Union(b)))
}

// TestIntersection covers tests for Intersection operation.
func TestIntersection(t *testing.T) {

	a := New[types.Int]()
	b := New[types.Int]()

	// Case 1 : Intersection of empty sets should return empty.
	c := a.Intersection(b)
	assert.Equal(t, true, c.Empty())

	// Case 2 : Intersection of populated set and empty set should return empty set.
	a.Add(1, 2, 3)
	c = a.Intersection(b)
	assert.Equal(t, true, c.Empty())

	// Case 3 : Intersection of populated sets.
	b.Add(1, 2, 4, 5, 6)
	d := New[types.Int](1, 2)

	assert.Equal(t, true, d.Equals(a.Intersection(b)))
	assert.Equal(t, true, d.Equals(b.Intersection(a)))

}

func TestMap(t *testing.T) {

	s := New[types.Int]()
	for i := 0; i < 6; i++ {
		s.Add(types.Int(i))
	}
	other := s.Map(func(e types.Int) types.Int { return e + 10 })

	a := []types.Int{10, 11, 12, 13, 14, 15}
	b := other.Collect()

	assert.Equal(t, true, testutils.EqualSlices(a, b))

}

func TestFilter(t *testing.T) {

	s := New[types.Int]()

	for i := 0; i < 6; i++ {
		s.Add(types.Int(i))
	}

	c := []types.Int{0, 2, 4}
	other := s.Filter(func(e types.Int) bool { return e%2 == 0 })
	d := other.Collect()

	assert.Equal(t, true, testutils.EqualSlices(c, d))

}

// TestEquals covers tests for Equals.
func TestEquals(t *testing.T) {

	s := New[types.Int]()
	other := New[types.Int]()
	assert.Equal(t, true, s.Equals(s))
	assert.Equal(t, true, s.Equals(other)) // Two empty sets are equal.

	s.Add(1)
	assert.Equal(t, false, s.Equals(other))
	other.Add(1)
	assert.Equal(t, true, s.Equals(other))
	s.Add(2)
	other.Add(3)
	assert.Equal(t, false, s.Equals(other))

}

// TestString covers tests for String.
func TestString(t *testing.T) {

	s := New[types.Int]()

	assert.Equal(t, "{}", fmt.Sprint(s))
	s.Add(1)
	assert.Equal(t, "{1}", fmt.Sprint(s))

}

func TestRetainAll(t *testing.T) {

	a := New[types.Int](1, 2, 3, 4)
	b := New[types.Int](2, 4, 7, 8, 9)

	a.RetainAll(b)
	assert.Equal(t, 2, a.Len())

}

func TestLeftSubset(t *testing.T) {

	s := New[types.Int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Case 1 : LeftSubset non inclusive.
	subset := s.LeftSubset(4, false)
	assert.Equal(t, true, testutils.EqualSlices([]types.Int{1, 2, 3}, subset.Collect()))

	// Case 2 : LeftSubset inclusive.
	subset = s.LeftSubset(4, true)
	assert.Equal(t, true, testutils.EqualSlices([]types.Int{1, 2, 3, 4}, subset.Collect()))

}

func TestRightSubset(t *testing.T) {

	s := New[types.Int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Case 1 : RightSubset non inclusive.
	subset := s.RightSubset(4, false)
	assert.Equal(t, true, testutils.EqualSlices([]types.Int{5, 6, 7, 8, 9, 10}, subset.Collect()))

	// Case 2 : RightSubset inclusive.
	subset = s.RightSubset(4, true)
	assert.Equal(t, true, testutils.EqualSlices([]types.Int{4, 5, 6, 7, 8, 9, 10}, subset.Collect()))

}

func TestSubset(t *testing.T) {

	s := New[types.Int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Case 1 : Subset with open ends.
	subset := s.Subset(2, false, 6, false)
	assert.Equal(t, true, testutils.EqualSlices([]types.Int{3, 4, 5}, subset.Collect()))

	// Case 2 : Subset with left closed end and right open.
	subset = s.Subset(2, true, 6, false)
	assert.Equal(t, true, testutils.EqualSlices([]types.Int{2, 3, 4, 5}, subset.Collect()))

	// Case 3 : Subset with left open end and right closed.
	subset = s.Subset(2, false, 6, true)
	assert.Equal(t, true, testutils.EqualSlices([]types.Int{3, 4, 5, 6}, subset.Collect()))

}
