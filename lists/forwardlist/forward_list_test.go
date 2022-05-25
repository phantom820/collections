package forwardlist

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

func TestAddFront(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Front on an empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrEmptyList, r.(error))
			}
		}()
		l.Front()
	})

	// Case 2 : Add front to an empty list.
	assert.Equal(t, true, l.Empty())
	l.AddFront(1)
	assert.Equal(t, 1, l.Len())
	assert.Equal(t, types.Int(1), l.Front())

	// Case 3 : Add front to a populated list.
	l.AddFront(2)
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, types.Int(2), l.Front())

}

func TestAddBack(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Back of an empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrEmptyList, r.(error))
			}
		}()
		l.Back()
	})

	// Case 2 : Add back to an empty list.
	assert.Equal(t, true, l.Empty())
	l.AddBack(1)
	assert.Equal(t, 1, l.Len())
	assert.Equal(t, types.Int(1), l.Back())

	// Case 3 : Add back to a populated list.
	l.AddFront(2)
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, types.Int(2), l.Front())

}

func TestReverse(t *testing.T) {

	l := New[types.Int]()

	assert.Equal(t, true, l.Equals(l.Reverse()))

	l.Add(1)
	l.Add(2)
	l.Add(3)

	r := New[types.Int]()
	r.Add(3)
	r.Add(2)
	r.Add(1)

	assert.Equal(t, true, l.Reverse().Equals(r))
	assert.Equal(t, true, l.Equals(r.Reverse()))
}

func TestAt(t *testing.T) {
	l := New[types.Int]()

	// Case 1 : At on an empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrOutOfBounds, r.(error))
			}
		}()
		l.At(0)
	})

	// Case 2 : At on a populated list.
	l.AddBack(1)
	l.AddBack(2)
	l.AddBack(3)
	assert.Equal(t, types.Int(1), l.At(0))
	assert.Equal(t, types.Int(2), l.At(1))
	assert.Equal(t, types.Int(3), l.At(2))

}

func TestEquals(t *testing.T) {

	l := New[types.Int]()
	other := New[types.Int]()

	// Case 1 : Self equivalence and empty lists.
	assert.Equal(t, true, l.Equals(l))
	assert.Equal(t, true, l.Equals(other))

	// Case 2 : Lists of unequal sizes should not be equal.
	for i := 1; i < 6; i++ {
		other.Add(types.Int(i))
	}
	assert.Equal(t, false, l.Equals(other))

	// Case 3 : Lists of equal sizes but different elements should not be equal.
	for i := 1; i < 6; i++ {
		l.Add(types.Int(i + 1))
	}

	assert.Equal(t, false, l.Equals(other))
	l.Clear()

	// Case 4 : Lists with same size and elements should be equal.

	for i := 1; i < 6; i++ {
		l.Add(types.Int(i))
	}

	assert.Equal(t, true, other.Equals(l))

}

func TestAdd(t *testing.T) {

	l := New[types.Int]()
	other := New[types.Int]()

	// Case 1 : Add with no elements.
	assert.Equal(t, false, l.Add())

	// Case 2 : Just add should add to the back of the list.
	assert.Equal(t, true, l.Add(1, 2, 3))
	assert.Equal(t, types.Int(3), l.Back())
	l.Clear()

	// Case 3 : AddAll should add all the elements from another iterable.
	for i := 1; i < 6; i++ {
		other.Add(types.Int(i))
	}

	l.AddAll(other)
	assert.Equal(t, true, l.Equals(other))

}

func TestAddAt(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Adding out of bounds.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrOutOfBounds, r.(error))
			}
		}()
		l.AddAt(0, 0)
	})

	// Case 2 : Adding at allowed index.
	l.Add(10, 20, 30)
	l.AddAt(0, 22)
	assert.Equal(t, types.Int(22), l.At(0))
	l.AddAt(l.Len()-1, 25)
	assert.Equal(t, types.Int(25), l.At(l.Len()-1))
	l.AddAt(2, -5)
	assert.Equal(t, types.Int(-5), l.At(2))

}

func TestSwap(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Swapping out of bounds should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrOutOfBounds, r.(error))
			}
		}()
		l.Swap(-1, 0)
	})

	// Case 2 : Swapping at legal index.
	l.Add(2, 3, 4, 5, 10)

	l.Swap(0, 1)
	assert.Equal(t, types.Int(3), l.Front())
	l.Swap(2, 3)
	l.Swap(1, 0)
	assert.Equal(t, types.Int(5), l.At(2))
	l.Swap(0, l.Len()-1)
	assert.Equal(t, types.Int(10), l.At(0))
	l.Swap(l.Len()-1, 0)
	assert.Equal(t, types.Int(10), l.At(l.Len()-1))

}

func TestRemoveFront(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Removing front from empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrEmptyList, r.(error))
			}
		}()
		l.RemoveFront()
	})

	// Case 2 : Removing front from list with elements.
	l.Add(22, 23, 234)

	assert.Equal(t, types.Int(22), l.RemoveFront())
	assert.Equal(t, types.Int(23), l.RemoveFront())
	assert.Equal(t, types.Int(234), l.RemoveFront())

}

func TestSet(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Set on empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrOutOfBounds, r.(error))
			}
		}()
		l.Set(0, 0)
	})

	// Case 2 : Set at legal indices.
	l.Add(1, 2, 3)
	l.Set(0, 45)
	assert.Equal(t, types.Int(45), l.Front())
	l.Set(2, -33)
	assert.Equal(t, types.Int(-33), l.Back())

}

func TestRemoveBack(t *testing.T) {
	l := New[types.Int]()

	// Case 1 : Removing back from empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrEmptyList, r.(error))
			}
		}()
		l.RemoveBack()
	})

	// Case 2 : Remove back from list with elements.
	l.Add(22, 23, 234, -2)

	assert.Equal(t, types.Int(-2), l.RemoveBack())
	assert.Equal(t, l.Len(), 3)
	assert.Equal(t, types.Int(234), l.RemoveBack())
	assert.Equal(t, l.Len(), 2)
	assert.Equal(t, types.Int(23), l.RemoveBack())
	assert.Equal(t, l.Len(), 1)
	assert.Equal(t, types.Int(22), l.RemoveBack())

}

func TestRemoveAt(t *testing.T) {

	l := New[types.Int]()

	/// Case 1 : Remmoving at in empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrEmptyList, r.(error))
			}
		}()
		l.RemoveAt(0)
	})

	// Case 2 : Remove from list with elements.
	l.Add(1, 2, 3, 6, 9, 80)

	l.RemoveAt(0)
	assert.Equal(t, types.Int(2), l.Front())
	assert.Equal(t, 5, l.Len())
	l.RemoveAt(2)
	assert.Equal(t, types.Int(9), l.At(2))
	assert.Equal(t, types.Int(80), l.RemoveAt(l.Len()-1))
	assert.Equal(t, 3, l.Len())

	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrOutOfBounds, r.(error))
			}
		}()
		l.RemoveAt(-1)
	})

}

func TestRemove(t *testing.T) {
	l := New[types.Int]()
	other := New[types.Int]()

	l.Add(1, 2, 3, 6, 9)

	other.Add(1, 2, 3, 4, 6, 9)

	assert.Equal(t, true, l.Contains(1))
	l.Remove(1)
	assert.Equal(t, false, l.Contains(1))
	assert.Equal(t, true, l.Contains(9))
	l.Remove(9)
	assert.Equal(t, false, l.Contains(9))
	l.Remove(3)
	assert.Equal(t, false, l.Contains(3))
	l.RemoveAll(other)
	assert.Equal(t, true, l.Empty())

}

func TestIterator(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Iterator on empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, iterator.NoNextElementError, r.(error))
			}
		}()
		it := l.Iterator()
		it.Next()
	})

	// Case 2 : Iterator on list with elements.
	for i := 1; i < 6; i++ {
		l.Add(types.Int(i))
	}
	a := l.Collect()
	b := make([]types.Int, 0)
	it := l.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, types.Int(1), it.Next())

}

func TestMapFilter(t *testing.T) {
	l := New[types.Int]()

	for i := 0; i < 6; i++ {
		l.Add(types.Int(i))
	}

	// Case 1 : Map to a new list.
	other := l.Map(func(e types.Int) types.Int { return e + 10 })

	a := []types.Int{10, 11, 12, 13, 14, 15}
	b := other.Collect()

	assert.ElementsMatch(t, a, b)

	// Case 2 : Filter to create new list.
	c := []types.Int{0, 2, 4}
	other = l.Filter(func(e types.Int) bool { return e%2 == 0 })
	d := other.Collect()
	assert.ElementsMatch(t, c, d)

}

func TestFilter(t *testing.T) {
	l := New[types.Int]()

	for i := 0; i < 6; i++ {
		l.Add(types.Int(i))
	}

	// Case 2 : Filter to create new list.
	c := []types.Int{0, 2, 4}
	other := l.Filter(func(e types.Int) bool { return e%2 == 0 })
	d := other.Collect()
	assert.ElementsMatch(t, c, d)

}

func TestClear(t *testing.T) {
	l := New[types.Int]()

	for i := 0; i < 20; i++ {
		l.Add(types.Int(i))
	}

	assert.Equal(t, 20, l.Len())
	l.Clear()
	assert.Equal(t, true, l.Empty())

}

func TestString(t *testing.T) {

	l := New[types.Int]()

	l.Add(2, 3, 4, 5, 10)
	assert.Equal(t, "[2 3 4 5 10]", fmt.Sprint(l))

}

func TestSort(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Sorting an empty list does nothing.
	Sort(l)
	assert.Equal(t, true, l.Empty())

	// Case 2 : Sorting a list with elements.
	l.Add(-10, 20, 0, 5, 4, 3, 2, 1)
	sorted := []types.Int{-10, 0, 1, 2, 3, 4, 5, 20}
	Sort(l)
	assert.ElementsMatch(t, sorted, l.Collect())

	// Try adding to sorted list to see if nothing broke.
	l.Add(100)
	assert.ElementsMatch(t, append(sorted, 100), l.Collect())
	l.AddFront(200)
	assert.ElementsMatch(t, append([]types.Int{200}, append(sorted, 100)...), l.Collect())

	// Case 2 : Sorting an already sorted list.
	l.Clear()
	l.Add(-10, 0, 1, 2, 3, 4, 5, 20)
	Sort(l)
	assert.ElementsMatch(t, sorted, l.Collect())

}

func TestSortBy(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Sorting an empty list does nothing.
	SortBy(l, func(a, b types.Int) bool { return a < b })
	assert.Equal(t, true, l.Empty())

	// Case 2 : Sorting a list with elements.
	l.Add(-10, 20, 0, 5, 4, 3, 2, 1)
	sorted := []types.Int{20, 5, 4, 3, 2, 1, 0, -10}
	SortBy(l, func(a, b types.Int) bool { return a > b })
	assert.ElementsMatch(t, sorted, l.Collect())

	// Try adding to sorted list to see if nothing broke.
	l.Add(100)
	assert.ElementsMatch(t, append(sorted, 100), l.Collect())
	l.AddFront(200)
	assert.ElementsMatch(t, append([]types.Int{200}, append(sorted, 100)...), l.Collect())

	// Case 3 : Sorting an already sorted list.
	l.Clear()
	l.Add(-10, 0, 1, 2, 3, 4, 5, 20)
	SortBy(l, func(a, b types.Int) bool { return a < b })
	assert.ElementsMatch(t, sorted, l.Collect())

}
