package hashset

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

// TestAdd covers tests for Add,AddAll and AddSlice.
func TestAdd(t *testing.T) {

	s := New[types.Integer]()
	assert.Equal(t, true, s.Empty())

	// Case 1 : Add to an empty set.
	s.Add(1)
	assert.Equal(t, 1, s.Len())
	assert.Equal(t, false, s.Add(1))
	assert.Equal(t, 1, s.Len())
	s.Add(2)
	assert.Equal(t, 2, s.Len())
	assert.Equal(t, true, s.Contains(1))
	assert.Equal(t, true, s.Contains(2))

	// Case 2 : Add multiple elements from another iterable.
	s = New[types.Integer]()
	l := list.New[types.Integer]()
	for i := 0; i < 10; i++ {
		l.Add(types.Integer(i))
	}

	s.AddAll(l)
	assert.Equal(t, l.Len(), s.Len())

	// should contain all the added elements
	it := l.Iterator()
	for it.HasNext() {
		assert.Equal(t, true, s.Contains(it.Next()))
	}

	// Case 3 Adding a slice should work accordingly.
	s.Clear()

	sl := []types.Integer{1, 1, 2, 3, 4, 5}
	s.AddSlice(sl)

	sm := []types.Integer{1, 2, 3, 4, 5}
	assert.ElementsMatch(t, sm, s.Collect())

}

func TestIterator(t *testing.T) {

	s := New[types.Integer]()
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
		s.Add(types.Integer(i))
	}
	a := s.Collect()
	b := make([]types.Integer, 0)
	it := s.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, types.Integer(1), it.Next())

	s.Clear()
	assert.Equal(t, true, s.Empty())

}

func TestRemove(t *testing.T) {

	s := New[types.Integer]()

	assert.Equal(t, false, s.Remove(1))
	s.Add(1)
	assert.Equal(t, true, s.Remove(1))
	assert.Equal(t, 0, s.Len())

	l := list.New[types.Integer]()
	for i := 1; i <= 10; i++ {
		s.Add(types.Integer(i))
		l.Add(types.Integer(i))
	}

	assert.Equal(t, 10, s.Len())
	s.RemoveAll(l)
	assert.Equal(t, 0, s.Len())

}

// TestMapFilter covers tests for Map and Filter
func TestMapFilter(t *testing.T) {

	s := New[types.Integer]()
	for i := 0; i < 6; i++ {
		s.Add(types.Integer(i))
	}
	other := s.Map(func(e types.Integer) types.Integer { return e + 10 })

	a := []types.Integer{10, 11, 12, 13, 14, 15}
	b := other.Collect()

	assert.ElementsMatch(t, a, b)

	c := []types.Integer{0, 2, 4}
	other = s.Filter(func(e types.Integer) bool { return e%2 == 0 })
	d := other.Collect()

	assert.ElementsMatch(t, c, d)

}

// TestEquals covers tests for Equals.
func TestEquals(t *testing.T) {

	s := New[types.Integer]()
	other := New[types.Integer]()
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

	s := New[types.Integer]()

	assert.Equal(t, "{}", fmt.Sprint(s))
	s.Add(1)
	assert.Equal(t, "{1}", fmt.Sprint(s))

}