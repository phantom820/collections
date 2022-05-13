package vector

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

// TestAdd covers tests for Add, Empty, Contains.
func TestAdd(t *testing.T) {

	v := New[types.Integer]()

	// v Starts out as empty.
	assert.Equal(t, true, v.Empty())

	// Case 1 : Add individual elements.
	v.Add(1)
	fmt.Println(v)
	assert.Equal(t, false, v.Empty())
	assert.Equal(t, 1, v.Len())
	assert.Equal(t, true, v.Contains(1))
	v.Add(2)
	assert.Equal(t, true, v.Contains(2))

	l := forwardlist.New[types.Integer]()

	for i := 3; i <= 10; i++ {
		l.Add(types.Integer(i))
	}

	// Case 2 : Add a number of elements at once.
	v.AddAll(l)
	assert.Equal(t, 10, v.Len())

	// Case 3 : Adding a slice should work accordingly
	v.Clear()
	s := []types.Integer{1, 2, 3, 4}
	v.AddSlice(s)

	assert.ElementsMatch(t, s, v.Collect())

}

// Covers tests for Iterator.
func TestIterator(t *testing.T) {
	v := New[types.Integer]()

	// Case 1 : Next on empty vector should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, iterator.NoNextElementError, r.(error))
			}
		}()
		it := v.Iterator()
		it.Next()
	})

	// Case 2 : Iterator should work accordingly on populated queue.
	for i := 1; i < 6; i++ {
		v.Add(types.Integer(i))
	}
	a := v.Collect()
	b := make([]types.Integer, 0)
	it := v.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, types.Integer(1), it.Next())

}

// TestRemove covers tests for Remove and RemoveAll.
func TestRemove(t *testing.T) {

	v := New[types.Integer]()

	// Case 1 : Removing from empty.
	assert.Equal(t, false, v.Remove(22))

	// Case 2 : Removing from poplated.
	v.Add(1)
	v.Add(2)
	v.Add(4)
	v.Add(5)

	assert.Equal(t, true, v.Remove(5))
	assert.Equal(t, false, v.Contains(5))

	s := forwardlist.New[types.Integer]()
	s.Add(1)
	s.Add(2)

	// Case 3 : Removing multiple elements at once.
	v.RemoveAll(s)
	assert.Equal(t, 1, v.Len())

}

// TestAddAt
func TestAddAt(t *testing.T) {

	v := New[types.Integer]()

	// Case 1 : Next on empty vector should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, ErrOutOfBounds, r.(error))
			}
		}()
		v.AddAt(0, 0)
	})

	// Case 2 : Should work accordinglt in valid indices.
	v.Add(1, 2, 3)
	v.AddAt(0, -1)
	assert.Equal(t, types.Integer(-1), v.data[0])
	v.AddAt(v.Len()-1, 22)
	assert.Equal(t, types.Integer(22), v.data[v.Len()-2])

}

// TestRemoveAt covers tests for remove at
func TestRemoveAt(t *testing.T) {

	v := New[types.Integer]()

	// Case 1 : RemoveAt on empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, ErrOutOfBounds, r.(error))
			}
		}()
		v.RemoveAt(0)
	})

	// Case 2 : Should work accordinglt in valid indices.
	v.Add(1, 2, 3, 5, 6)
	assert.Equal(t, types.Integer(1), v.RemoveAt(0))
	assert.Equal(t, types.Integer(2), v.data[0])

}

// TestEquals for Equals method of vectors.
func TestEquals(t *testing.T) {

	v := New[types.Integer]()
	other := New[types.Integer]()

	// Case 1 : A list is equal to its self.
	assert.Equal(t, true, v.Equals(v))

	// Case 2 : 2 empty vectors are equal.
	assert.Equal(t, true, v.Equals(other))

	// Case 3 : vectors of unequal sizes should not be equal.
	for i := 1; i < 6; i++ {
		other.Add(types.Integer(i))
	}

	assert.Equal(t, false, v.Equals(other))

	// Case 4 : vectors of equal sizes but different elements should not be equal.
	for i := 1; i < 6; i++ {
		v.Add(types.Integer(i + 1))
	}

	assert.Equal(t, false, v.Equals(other))
	v.Clear()

	// Case 5 : vectors with same size and elements should be equal.

	for i := 1; i < 6; i++ {
		v.Add(types.Integer(i))
	}

	assert.Equal(t, true, other.Equals(v))

}

// TestAt covers tests for At.
func TestAt(t *testing.T) {

	v := New[types.Integer]()

	// Case 1 : At on empty vector should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, ErrOutOfBounds, r.(error))
			}
		}()
		v.At(0)
	})

	v.Add(1, 2, 3, 4)

	assert.Equal(t, types.Integer(2), v.At(1))
	assert.Equal(t, types.Integer(3), v.At(2))

}

// TestSet covers tests for set.
func TestSet(t *testing.T) {

	v := New[types.Integer]()

	// Case 1 : Set on empty vector should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, ErrOutOfBounds, r.(error))
			}
		}()
		v.Set(0, 0)
	})

	v.Add(1, 2, 3)
	v.Set(1, -4)
	assert.Equal(t, types.Integer(-4), v.At(1))

}

// TestMapFilter covers tests for Map and Filter.
func TestMapFilter(t *testing.T) {
	v := New[types.Integer]()

	for i := 0; i < 6; i++ {
		v.Add(types.Integer(i))
	}

	// Case 1 : Map to a new vector.
	other := v.Map(func(e types.Integer) types.Integer { return e + 10 })

	a := []types.Integer{10, 11, 12, 13, 14, 15}
	b := other.Collect()

	assert.ElementsMatch(t, a, b)

	// Case 2 : Filter to create new vector.
	c := []types.Integer{0, 2, 4}
	other = v.Filter(func(e types.Integer) bool { return e%2 == 0 })
	d := other.Collect()
	assert.ElementsMatch(t, c, d)
}

// TestString covers tests for String.
func TestString(t *testing.T) {
	v := New[types.Integer]()

	v.Add(1)
	v.Add(2)
	v.Add(3)

	assert.Equal(t, "[1 2 3]", fmt.Sprint(v))
}
