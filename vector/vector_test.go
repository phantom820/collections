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

	v := New[types.Int]()

	// v Starts out as empty.
	assert.Equal(t, true, v.Empty())

	// Case 1 : Add individual elements.
	v.Add(1)
	assert.Equal(t, false, v.Empty())
	assert.Equal(t, 1, v.Len())
	assert.Equal(t, true, v.Contains(1))
	v.Add(2)
	assert.Equal(t, true, v.Contains(2))

	l := forwardlist.New[types.Int]()

	for i := 3; i <= 10; i++ {
		l.Add(types.Int(i))
	}

	// Case 2 : Add a number of elements at once.
	v.AddAll(l)
	assert.Equal(t, 10, v.Len())

	// Case 3 : Adding a slice should work accordingly
	v.Clear()
	s := []types.Int{1, 2, 3, 4}
	v.Add(s...)

	assert.ElementsMatch(t, s, v.Collect())

}

// Covers tests for Iterator.
func TestIterator(t *testing.T) {
	v := New[types.Int]()

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
		v.Add(types.Int(i))
	}
	a := v.Collect()
	b := make([]types.Int, 0)
	it := v.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, types.Int(1), it.Next())

}

// TestRemove covers tests for Remove and RemoveAll.
func TestRemove(t *testing.T) {

	v := New[types.Int]()

	// Case 1 : Removing from empty.
	assert.Equal(t, false, v.Remove(22))

	// Case 2 : Removing from poplated.
	v.Add(1)
	v.Add(2)
	v.Add(4)
	v.Add(5)

	assert.Equal(t, true, v.Remove(5))
	assert.Equal(t, false, v.Contains(5))

	s := forwardlist.New[types.Int]()
	s.Add(1)
	s.Add(2)

	// Case 3 : Removing multiple elements at once.
	v.RemoveAll(s)
	assert.Equal(t, 1, v.Len())

}

// TestAddAt
func TestAddAt(t *testing.T) {

	v := New[types.Int]()

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
	assert.Equal(t, types.Int(-1), v.data[0])
	v.AddAt(v.Len()-1, 22)
	assert.Equal(t, types.Int(22), v.data[v.Len()-2])
	v.AddAt(2, 23)
	assert.Equal(t, types.Int(23), v.At(2))

}

// TestRemoveAt covers tests for remove at
func TestRemoveAt(t *testing.T) {

	v := New[types.Int]()

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
	assert.Equal(t, types.Int(1), v.RemoveAt(0))
	assert.Equal(t, types.Int(2), v.data[0])

}

// TestEquals for Equals method of vectors.
func TestEquals(t *testing.T) {

	v := New[types.Int]()
	other := New[types.Int]()

	// Case 1 : A list is equal to its self.
	assert.Equal(t, true, v.Equals(v))

	// Case 2 : 2 empty vectors are equal.
	assert.Equal(t, true, v.Equals(other))

	// Case 3 : vectors of unequal sizes should not be equal.
	for i := 1; i < 6; i++ {
		other.Add(types.Int(i))
	}

	assert.Equal(t, false, v.Equals(other))

	// Case 4 : vectors of equal sizes but different elements should not be equal.
	for i := 1; i < 6; i++ {
		v.Add(types.Int(i + 1))
	}

	assert.Equal(t, false, v.Equals(other))
	v.Clear()

	// Case 5 : vectors with same size and elements should be equal.

	for i := 1; i < 6; i++ {
		v.Add(types.Int(i))
	}

	assert.Equal(t, true, other.Equals(v))

}

// TestAt covers tests for At.
func TestAt(t *testing.T) {

	v := New[types.Int]()

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

	assert.Equal(t, types.Int(2), v.At(1))
	assert.Equal(t, types.Int(3), v.At(2))

}

// TestSet covers tests for set.
func TestSet(t *testing.T) {

	v := New[types.Int]()

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
	assert.Equal(t, types.Int(-4), v.At(1))

}

// TestMapFilter covers tests for Map and Filter.
func TestMapFilter(t *testing.T) {
	v := New[types.Int]()

	for i := 0; i < 6; i++ {
		v.Add(types.Int(i))
	}

	// Case 1 : Map to a new vector.
	other := v.Map(func(e types.Int) types.Int { return e + 10 })

	a := []types.Int{10, 11, 12, 13, 14, 15}
	b := other.Collect()

	assert.ElementsMatch(t, a, b)

	// Case 2 : Filter to create new vector.
	c := []types.Int{0, 2, 4}
	other = v.Filter(func(e types.Int) bool { return e%2 == 0 })
	d := other.Collect()
	assert.ElementsMatch(t, c, d)
}

// TestString covers tests for String.
func TestString(t *testing.T) {
	v := New[types.Int]()

	v.Add(1)
	v.Add(2)
	v.Add(3)

	assert.Equal(t, "[1 2 3]", fmt.Sprint(v))
}

func TestSort(t *testing.T) {

	v := New[types.Int]()

	// Case 1 : Sorting an empty vector does nothing.
	Sort(v)
	assert.Equal(t, true, v.Empty())

	// Case 2 : Sorting a vector with elements.
	v.Add(-10, 20, 0, 5, 4, 3, 2, 1)
	sorted := []types.Int{-10, 0, 1, 2, 3, 4, 5, 20}
	Sort(v)
	assert.ElementsMatch(t, sorted, v.Collect())

	// Try adding to sorted vector to see if nothing broke.
	v.Add(100)
	assert.ElementsMatch(t, append(sorted, 100), v.Collect())
	v.AddAt(0, 200)
	assert.ElementsMatch(t, append([]types.Int{200}, append(sorted, 100)...), v.Collect())

	// Case 2 : Sorting an already sorted vector.
	v.Clear()
	v.Add(-10, 0, 1, 2, 3, 4, 5, 20)
	Sort(v)
	assert.ElementsMatch(t, sorted, v.Collect())

}

func TestSortBy(t *testing.T) {

	v := New[types.Int]()

	// Case 1 : Sorting an empty vector does nothing.
	SortBy(v, func(a, b types.Int) bool { return a < b })
	assert.Equal(t, true, v.Empty())

	// Case 2 : Sorting a vector with elements.
	v.Add(-10, 20, 0, 5, 4, 3, 2, 1)
	sorted := []types.Int{20, 5, 4, 3, 2, 1, 0, -10}
	SortBy(v, func(a, b types.Int) bool { return a > b })
	assert.ElementsMatch(t, sorted, v.Collect())

	// Try adding to sorted vector to see if nothing broke.
	v.Add(100)
	assert.ElementsMatch(t, append(sorted, 100), v.Collect())
	v.AddAt(0, 200)
	assert.ElementsMatch(t, append([]types.Int{200}, append(sorted, 100)...), v.Collect())

	// Case 2 : Sorting an already sorted vector.
	v.Clear()
	v.Add(-10, 0, 1, 2, 3, 4, 5, 20)
	SortBy(v, func(a, b types.Int) bool { return a < b })
	assert.ElementsMatch(t, sorted, v.Collect())

}
