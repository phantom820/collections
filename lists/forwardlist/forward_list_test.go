package forwardlist

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/testutils"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

func TestAddFront(t *testing.T) {

	l := New[types.Int]()
	assert.Equal(t, true, l.Empty())

	// Case 1 : AddFront on an empty list.
	l.AddFront(1)
	assert.Equal(t, 1, l.Len())

	// Case 2 : AddFront on a list with elements.
	l.AddFront(2)
	assert.Equal(t, 2, l.Len())
	l.addFront(3)
	assert.Equal(t, 3, l.Len())

}

func TestFront(t *testing.T) {

	l := New[types.Int]()
	assert.Equal(t, true, l.Empty())

	// Case 1 : Front of an empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		l.Front()
	})

	// Case 2 : Front on a list with items.
	l.AddFront(1)
	assert.Equal(t, types.Int(1), l.Front())
	l.AddFront(2)
	assert.Equal(t, types.Int(2), l.Front())
	l.AddFront(3)
	assert.Equal(t, types.Int(3), l.Front())

}

func TestAdd(t *testing.T) {

	l := New[types.Int]()
	assert.Equal(t, true, l.Empty())

	// Case 1 : Add on an empty list.
	l.Add(1)
	assert.Equal(t, 1, l.Len())

	// Case 2 : Add on a list with elements.
	l.Add(2)
	assert.Equal(t, 2, l.Len())
	l.Add(3)
	assert.Equal(t, 3, l.Len())

}

func TestAddAt(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : AddAt with out of bounds index.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.IndexOutOfBounds, r.(errors.Error).Code())
			}
		}()
		l.AddAt(0, 0)
	})

	// Case 2 : AddAt with valid index.
	l.Add(10, 20, 30)
	l.AddAt(0, 22)
	assert.Equal(t, types.Int(22), l.At(0))
	l.AddAt(l.Len()-1, 25)
	assert.Equal(t, types.Int(25), l.At(l.Len()-2))
	l.AddAt(2, -5)
	assert.Equal(t, types.Int(-5), l.At(2))

}

func TestAddAll(t *testing.T) {

	l1 := New[types.Int]()
	l2 := New[types.Int](1, 2, 3, 4, 5, 6)

	// Case 1: AddAll on an empty list.
	assert.Equal(t, true, l1.Empty())
	l1.AddAll(l2)
	assert.Equal(t, 6, l1.Len())

	// Case 2: AddAll on a list with items.
	l1.AddAll(l2)
	assert.Equal(t, 12, l1.Len())

}

func TestAt(t *testing.T) {
	l := New[types.Int]()

	// Case 1 : At on an empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.IndexOutOfBounds, r.(errors.Error).Code())
			}
		}()
		l.At(0)
	})

	// Case 2 : At on a populated list.
	l.Add(1, 2, 3)

	assert.Equal(t, types.Int(1), l.At(0))
	assert.Equal(t, types.Int(2), l.At(1))
	assert.Equal(t, types.Int(3), l.At(2))

}

func TestBack(t *testing.T) {

	l := New[types.Int]()
	assert.Equal(t, true, l.Empty())

	// Case 1 : Back of an empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		l.Back()
	})

	// Case 2 : Back on a list with items.
	l.Add(1)
	assert.Equal(t, types.Int(1), l.Back())
	l.Add(2)
	assert.Equal(t, types.Int(2), l.Back())
	l.Add(3)
	assert.Equal(t, types.Int(3), l.Back())

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
	l.Add(22)
	assert.Equal(t, types.Int(22), l.Back())

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

func TestSwap(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Swapping out of indices out of bounds should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.IndexOutOfBounds, r.(errors.Error).Code())
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

func TestSet(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Set on empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.IndexOutOfBounds, r.(errors.Error).Code())
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
	l.Set(1, 90)
	assert.Equal(t, types.Int(90), l.At(1))

}

func TestRemoveFront(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Removing front from empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		l.RemoveFront()
	})

	// Case 2 : Removing front from list with elements.
	l.Add(22, 23, 234)

	assert.Equal(t, 3, l.Len())
	assert.Equal(t, types.Int(22), l.RemoveFront())
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, types.Int(23), l.RemoveFront())
	assert.Equal(t, 1, l.Len())
	assert.Equal(t, types.Int(234), l.RemoveFront())
	assert.Equal(t, true, l.Empty())

}

func TestRemoveBack(t *testing.T) {
	l := New[types.Int]()

	// Case 1 : Removing back from empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
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
				assert.Equal(t, errors.IndexOutOfBounds, r.(errors.Error).Code())
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

func TestRemoveAll(t *testing.T) {

	l1 := New[types.Int](1, 2, 3, 4, 5, 6)
	l2 := New[types.Int](2, 4, 6)

	l1.RemoveAll(l2)
	assert.Equal(t, 3, l1.Len())
	assert.Equal(t, false, l1.Contains(2))
	assert.Equal(t, false, l1.Contains(2))

}

func TestIterator(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Iterator on empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoNextElement, r.(errors.Error).Code())
			}
		}()
		it := l.Iterator()
		it.Next()
	})

	// Case 2 : Iterator on list with elements.
	l.Add(1, 2, 3, 4, 5)

	a := l.Collect()
	b := make([]types.Int, 0)
	it := l.Iterator()

	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.Equal(t, true, testutils.EqualSlices(a, b))

}

func TestIteratorConcurrentModification(t *testing.T) {

	l := New[types.String]()
	for i := 1; i <= 20; i++ {
		l.Add(types.String(fmt.Sprint(i)))
	}

	// Recovery for concurrent modifications.
	recovery := func() {
		if r := recover(); r != nil {
			assert.Equal(t, errors.ConcurrentModification, r.(*errors.Error).Code())
		}
	}
	// Case 1 : Add.
	it := l.Iterator()
	t.Run("Add while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			l.Add(types.String("D"))
			it.Next()
		}
	})
	// Case 2 : AddFront.
	it = l.Iterator()
	t.Run("AddFront while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			l.AddFront(types.String("D"))
			it.Next()
		}
	})
	// Case 3 : RemoveFront.
	it = l.Iterator()
	t.Run("RemoveFront while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			l.RemoveFront()
			it.Next()
		}
	})
	// Case 4 : RemoveBack.
	it = l.Iterator()
	t.Run("RemoveBack while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			l.RemoveBack()
			it.Next()
		}
	})
	// Case 5 : Remove.
	it = l.Iterator()
	t.Run("Remove while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			l.Remove()
			it.Next()
		}
	})
	// Case 6 : RemoveAt.
	it = l.Iterator()
	t.Run("RemoveAt while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			l.RemoveAt(0)
			it.Next()
		}
	})
	// Case 7 : Swap.
	it = l.Iterator()
	t.Run("Swap while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			l.Swap(0, 1)
			it.Next()
		}
	})
	// Case 8 : Reverse.
	it = l.Iterator()
	t.Run("Reverse while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			l.Reverse()
			it.Next()
		}
	})
	// Case 9 : Clear.
	it = l.Iterator()
	t.Run("Clear while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			l.Clear()
			it.Next()
		}
	})
	// Case 10 : Sort.
	it = l.Iterator()
	t.Run("Sort while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			Sort(l)
			it.Next()
		}
	})
	// Case 11 : SortBy.
	it = l.Iterator()
	t.Run("SortBy while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			SortBy(l, func(a, b types.String) bool { return a < b })
			it.Next()
		}
	})

}

func TestMap(t *testing.T) {

	l := New[types.Int](0, 1, 2, 3, 4, 5)

	// Case 1 : Map to a new list.
	other := l.Map(func(e types.Int) types.Int { return e + 10 })

	a := []types.Int{10, 11, 12, 13, 14, 15}
	b := other.Collect()

	assert.ElementsMatch(t, a, b)

}

func TestFilter(t *testing.T) {

	l := New[types.Int](0, 1, 2, 3, 4, 5)

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
	assert.Equal(t, true, testutils.EqualSlices(sorted, l.Collect()))

	// Try adding to sorted list to see if nothing broke.
	l.Add(100)
	assert.ElementsMatch(t, append(sorted, 100), l.Collect())
	l.AddFront(200)
	assert.ElementsMatch(t, append([]types.Int{200}, append(sorted, 100)...), l.Collect())

	// Case 2 : Sorting an already sorted list.
	l.Clear()
	l.Add(-10, 0, 1, 2, 3, 4, 5, 20)
	Sort(l)
	assert.Equal(t, true, testutils.EqualSlices(sorted, l.Collect()))

}

func TestSortBy(t *testing.T) {

	l := New[types.Int]()

	// Case 1 : Sorting an empty list does nothing.
	SortBy(l, func(a, b types.Int) bool { return a < b })
	assert.Equal(t, true, l.Empty())

	// Case 2 : Sorting a list with elements.
	l.Add(-10, 20, 0, 5, 4, 3, 2, 1)
	reverseSorted := []types.Int{20, 5, 4, 3, 2, 1, 0, -10}
	SortBy(l, func(a, b types.Int) bool { return a > b })
	assert.ElementsMatch(t, reverseSorted, l.Collect())

	// Try adding to sorted list to see if nothing broke.
	l.Add(100)
	assert.ElementsMatch(t, append(reverseSorted, 100), l.Collect())
	l.AddFront(200)
	assert.ElementsMatch(t, append([]types.Int{200}, append(reverseSorted, 100)...), l.Collect())

	// Case 3 : Sorting an already sorted list.
	l.Clear()
	sorted := []types.Int{-10, 0, 1, 2, 3, 4, 5, 20}
	l.Add(-10, 0, 1, 2, 3, 4, 5, 20)
	SortBy(l, func(a, b types.Int) bool { return a < b })
	assert.ElementsMatch(t, sorted, l.Collect())

}
