package vector

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/lists"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/testutils"
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

func TestAdd(t *testing.T) {

	v := New[types.Int]()

	// Case 1 : Add individual elements.
	assert.Equal(t, true, v.Empty())
	v.Add(1)
	assert.Equal(t, false, v.Empty())
	assert.Equal(t, 1, v.Len())
	assert.Equal(t, true, v.Contains(1))
	v.Add(2)
	assert.Equal(t, true, v.Contains(2))

	l := forwardlist.New[types.Int](3, 4, 5, 6, 7, 8, 9, 10)

	// Case 2 : Add a number of elements at once.
	v.AddAll(l)
	assert.Equal(t, 10, v.Len())

	// Case 3 : Adding a slice should work accordingly
	v.Clear()
	s := []types.Int{1, 2, 3, 4}
	v.Add(s...)

	assert.ElementsMatch(t, s, v.Collect())

}

func TestIterator(t *testing.T) {

	v := New[types.Int]()

	// Case 1 : Next on empty vector should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoNextElement, r.(errors.Error).Code())
			}
		}()
		it := v.Iterator()
		it.Next()
	})

	// Case 2 : Iterator should work accordingly on populated queue.
	v.Add(1, 2, 3, 4, 5)

	a := v.Collect()
	b := make([]types.Int, 0)

	it := v.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)

}

func TestIteratorConcurrentModification(t *testing.T) {

	v := New[types.String]()
	for i := 1; i <= 20; i++ {
		v.Add(types.String(fmt.Sprint(i)))
	}

	// Recovery for concurrent modifications.
	recovery := func() {
		if r := recover(); r != nil {
			assert.Equal(t, errors.ConcurrentModification, r.(*errors.Error).Code())
		}
	}
	// Case 1 : Add.
	it := v.Iterator()
	t.Run("Add while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			v.Add(types.String("D"))
			it.Next()
		}
	})
	// Case 2 : AddFront.
	it = v.Iterator()
	t.Run("AddFront while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			v.AddFront(types.String("D"))
			it.Next()
		}
	})
	// Case 3 : RemoveFront.
	it = v.Iterator()
	t.Run("RemoveFront while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			v.RemoveFront()
			it.Next()
		}
	})
	// Case 4 : RemoveBack.
	it = v.Iterator()
	t.Run("RemoveBack while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			v.RemoveBack()
			it.Next()
		}
	})
	// Case 5 : Remove.
	it = v.Iterator()
	t.Run("Remove while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			v.Remove()
			it.Next()
		}
	})
	// Case 6 : RemoveAt.
	it = v.Iterator()
	t.Run("RemoveAt while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			v.RemoveAt(0)
			it.Next()
		}
	})
	// Case 7 : Swap.
	it = v.Iterator()
	t.Run("Swap while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			v.Swap(0, 1)
			it.Next()
		}
	})
	// Case 8 : Reverse.
	it = v.Iterator()
	t.Run("Reverse while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			v.Reverse()
			it.Next()
		}
	})
	// Case 9 : Clear.
	it = v.Iterator()
	t.Run("Clear while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			v.Clear()
			it.Next()
		}
	})
	// Case 10 : Sort.
	it = v.Iterator()
	t.Run("Sort while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			Sort(v)
			it.Next()
		}
	})
	// Case 11 : SortBy.
	it = v.Iterator()
	t.Run("SortBy while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			SortBy(v, func(a, b types.String) bool { return a < b })
			it.Next()
		}
	})
}

func TestRemove(t *testing.T) {

	v := New[types.Int]()

	// Case 1 : Removing from empty.
	assert.Equal(t, false, v.Remove(22))

	// Case 2 : Removing from poplated.
	v.Add(1, 2, 4, 5)

	assert.Equal(t, true, v.Remove(5))
	assert.Equal(t, false, v.Contains(5))

	// Case 3 : Removing multiple elements at once.
	l := forwardlist.New[types.Int](1, 2)
	v.RemoveAll(l)
	assert.Equal(t, 1, v.Len())

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

func TestReverse(t *testing.T) {

	l := New[types.Int](1, 2, 3)

	// Case 1 : Reverse a list with odd number of elements.
	l.Reverse()
	assert.Equal(t, true, testutils.EqualSlices([]types.Int{3, 2, 1}, l.Collect()))
	assert.Equal(t, types.Int(3), l.Front())
	assert.Equal(t, types.Int(1), l.Back())

	// Case 2 : Reverse a list with an even number of elements.
	l.Add(4, 5, 6)
	l.Reverse()
	assert.Equal(t, true, testutils.EqualSlices([]types.Int{6, 5, 4, 1, 2, 3}, l.Collect()))
	assert.Equal(t, types.Int(6), l.Front())
	assert.Equal(t, types.Int(3), l.Back())

}

func TestAddAt(t *testing.T) {

	v := New[types.Int]()

	// Case 1 : Next on empty vector should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.IndexOutOfBounds, r.(errors.Error).Code())
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

func TestRemoveAt(t *testing.T) {

	v := New[types.Int]()

	// Case 1 : RemoveAt on empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.IndexOutOfBounds, r.(errors.Error).Code())
			}
		}()
		v.RemoveAt(0)
	})

	// Case 2 : Should work accordinglt in valid indices.
	v.Add(1, 2, 3, 5, 6)
	assert.Equal(t, types.Int(1), v.RemoveAt(0))
	assert.Equal(t, types.Int(2), v.data[0])

}

func TestEquals(t *testing.T) {

	v := New[types.Int]()
	other := New[types.Int]()

	// Case 1 : A list is equal to its self.
	assert.Equal(t, true, v.Equals(v))

	// Case 2 : 2 empty vectors are equal.
	assert.Equal(t, true, v.Equals(other))

	// Case 3 : vectors of unequal sizes should not be equal.
	other.Add(1, 2, 3, 4, 5)
	assert.Equal(t, false, v.Equals(other))

	// Case 4 : vectors of equal sizes but different elements should not be equal.
	v.Add(2, 3, 4, 5, 6)
	assert.Equal(t, false, v.Equals(other))
	v.Clear()

	// Case 5 : vectors with same size and elements should be equal.
	v.Add(1, 2, 3, 4, 5)
	assert.Equal(t, true, other.Equals(v))

}

func TestAt(t *testing.T) {

	v := New[types.Int]()

	// Case 1 : At on empty vector should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.IndexOutOfBounds, r.(errors.Error).Code())
			}
		}()
		v.At(0)
	})

	v.Add(1, 2, 3, 4)

	assert.Equal(t, types.Int(2), v.At(1))
	assert.Equal(t, types.Int(3), v.At(2))

}

func TestSet(t *testing.T) {

	v := New[types.Int]()

	// Case 1 : Set on empty vector should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.IndexOutOfBounds, r.(errors.Error).Code())
			}
		}()
		v.Set(0, 0)
	})

	v.Add(1, 2, 3)
	v.Set(1, -4)
	assert.Equal(t, types.Int(-4), v.At(1))
	v.Set(v.Len()-1, -5)
	assert.Equal(t, types.Int(-5), v.At(v.Len()-1))

}

func TestMap(t *testing.T) {

	v := New[types.Int](0, 1, 2, 3, 4, 5)

	// Case 1 : Map to a new vector.
	other := v.Map(func(e types.Int) types.Int { return e + 10 })

	a := []types.Int{10, 11, 12, 13, 14, 15}
	b := other.Collect()
	assert.ElementsMatch(t, a, b)
}

func TestFilter(t *testing.T) {

	v := New[types.Int](0, 1, 2, 3, 4, 5)

	// Case 2 : Filter to create new vector.
	c := []types.Int{0, 2, 4}
	other := v.Filter(func(e types.Int) bool { return e%2 == 0 })
	d := other.Collect()
	assert.ElementsMatch(t, c, d)

}

func TestString(t *testing.T) {

	v := New[types.Int](1, 2, 3)
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
