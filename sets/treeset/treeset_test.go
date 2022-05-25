package treeset

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/list"
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
	assert.ElementsMatch(t, sm, s.Collect())

	// Case 5 : Clear and add a vast colelction of values
	s.Clear()
	slice := []types.Int{}
	for i := 50; i >= 0; i-- {
		s.Add(types.Int(i))
		slice = append(slice, types.Int(50-i))
	}
	assert.ElementsMatch(t, slice, s.Collect())

}

func TestIterator(t *testing.T) {

	s := New[types.Int]()
	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, iterator.NoNextElementError, r.(error))
			}
		}()
		it := s.Iterator()
		it.Next()
	})

	for i := 1; i < 6; i++ {
		s.Add(types.Int(i))
	}
	a := s.Collect()
	b := make([]types.Int, 0)
	it := s.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, types.Int(1), it.Next())

	s.Clear()
	assert.Equal(t, true, s.Empty())

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

// TestMapFilter covers tests for Map and Filter
func TestMapFilter(t *testing.T) {

	s := New[types.Int]()
	for i := 0; i < 6; i++ {
		s.Add(types.Int(i))
	}
	other := s.Map(func(e types.Int) types.Int { return e + 10 })

	a := []types.Int{10, 11, 12, 13, 14, 15}
	b := other.Collect()

	assert.ElementsMatch(t, a, b)

	c := []types.Int{0, 2, 4}
	other = s.Filter(func(e types.Int) bool { return e%2 == 0 })
	d := other.Collect()

	assert.ElementsMatch(t, c, d)

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
