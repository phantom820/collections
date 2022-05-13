package forwardlist

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

// TestAddFront covers tests for Front and  AddFront
func TestAddFront(t *testing.T) {

	l := New[types.Integer]()

	// Case 1 : Empty list jhas no front and should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrEmptyList, r.(error))
			}
		}()
		l.Front()
	})
	assert.Equal(t, 0, l.Len())

	// Case 2 : Elements added to front should be reflected by Front().
	l.AddFront(22)
	assert.Equal(t, types.Integer(22), l.Front())
	assert.Equal(t, 1, l.Len())
	l.AddFront(23)
	assert.Equal(t, types.Integer(23), l.Front())
	assert.Equal(t, 2, l.Len())

}

// TestAddBack covers tests Back and AddBack.
func TestAddBack(t *testing.T) {

	l := New[types.Integer]()

	// Case 1 : Empty list has no back element should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrEmptyList, r.(error))
			}
		}()
		l.Back()
	})

	assert.Equal(t, 0, l.Len())

	// Case 2 : Elements added to the back should be reflected by Back().
	l.AddBack(22)
	assert.Equal(t, types.Integer(22), l.Back())
	assert.Equal(t, 1, l.Len())
	l.AddBack(23)
	assert.Equal(t, types.Integer(23), l.Back())
	assert.Equal(t, 2, l.Len()) // len of the ;list should change accordingly.

}

// TestReverse covers tests for Reverse.
func TestReverse(t *testing.T) {

	l := New[types.Integer]()

	assert.Equal(t, true, l.Equals(l.Reverse()))

	l.Add(1)
	l.Add(2)
	l.Add(3)

	r := New[types.Integer]()
	r.Add(3)
	r.Add(2)
	r.Add(1)

	assert.Equal(t, true, l.Reverse().Equals(r))
	assert.Equal(t, true, l.Equals(r.Reverse()))
}

// TestAt covers tests for At
func TestAt(t *testing.T) {
	l := New[types.Integer]()

	// Case 1 : Empty list should not be indexable.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrOutOfBounds, r.(error))
			}
		}()
		l.At(0)
	})

	// Case 2 : Elements added should appear at appropiate index.
	l.AddBack(1)
	l.AddBack(2)
	l.AddBack(3)
	assert.Equal(t, types.Integer(1), l.At(0))
	assert.Equal(t, types.Integer(2), l.At(1))
	assert.Equal(t, types.Integer(3), l.At(2))

}

// TestEquals for Equals method of lists.
func TestEquals(t *testing.T) {

	l := New[types.Integer]()
	other := New[types.Integer]()

	// Case 1 : A list is equal to its self.
	assert.Equal(t, true, l.Equals(l))

	// Case 2 : 2 empty lists are equal.
	assert.Equal(t, true, l.Equals(other))

	// Case 3 : lists of unequal sizes should not be equal.
	for i := 1; i < 6; i++ {
		other.Add(types.Integer(i))
	}

	assert.Equal(t, false, l.Equals(other))

	// Case 4 : list of equal sizes but different elements should not be equal.
	for i := 1; i < 6; i++ {
		l.Add(types.Integer(i + 1))
	}

	assert.Equal(t, false, l.Equals(other))
	l.Clear()

	// Case 5 : lists with same size and elements should be equal.

	for i := 1; i < 6; i++ {
		l.Add(types.Integer(i))
	}

	assert.Equal(t, true, other.Equals(l))

}

// TestAdd covers tests for Add, AddAll, AddSlice.
func TestAdd(t *testing.T) {

	// Case 1 : Just add should add at back.

	l := New[types.Integer](1, 2)
	other := New[types.Integer]()

	assert.Equal(t, types.Integer(2), l.Back())
	l.Clear()

	// Case 2 : AddAll should add all the elements from another iterable.
	for i := 1; i < 6; i++ {
		other.Add(types.Integer(i))
	}

	l.AddAll(other)
	assert.Equal(t, true, l.Equals(other))

	l.Clear()

	// Case 4 : AddSlice should behave accordingly.
	s := []types.Integer{1, 2, 3, 4}
	l.AddSlice(s)

	assert.ElementsMatch(t, s, l.Collect())

}

// TestAddAt covers tests for AddAt adding at specified index.
func TestAddAt(t *testing.T) {
	l := New[types.Integer]()

	// Case 1 : Adding out of bounds.
	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrOutOfBounds, r.(error))
			}
		}()
		l.AddAt(0, 0)
	})

	// Case 2 : Adding at allowed index.
	l.Add(10)
	l.Add(20)
	l.Add(30)

	l.AddAt(0, 22)
	assert.Equal(t, types.Integer(22), l.At(0))
	l.AddAt(l.Len()-1, 25)
	assert.Equal(t, types.Integer(25), l.At(l.Len()-1))
	l.AddAt(2, -5)
	assert.Equal(t, types.Integer(-5), l.At(2))

}

// TestSwap covers tests for Swap swapping elements at specified indices.
func TestSwap(t *testing.T) {
	l := New[types.Integer]()

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

	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(10)

	l.Swap(0, 1)
	assert.Equal(t, types.Integer(3), l.Front())
	l.Swap(2, 3)
	l.Swap(1, 0)
	assert.Equal(t, types.Integer(5), l.At(2))
	l.Swap(0, l.Len()-1)
	assert.Equal(t, types.Integer(10), l.At(0))
	l.Swap(l.Len()-1, 0)
	assert.Equal(t, types.Integer(10), l.At(l.Len()-1))

}

// TestRemoveFront covers test for RemoveFront.
func TestRemoveFront(t *testing.T) {

	l := New[types.Integer]()

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
	l.Add(22)
	l.Add(23)
	l.Add(234)
	assert.Equal(t, types.Integer(22), l.RemoveFront())
	assert.Equal(t, types.Integer(23), l.RemoveFront())
	assert.Equal(t, types.Integer(234), l.RemoveFront())

}

// TestSet covers tests for Set overriding value at specified index.
func TestSet(t *testing.T) {

	l := New[types.Integer]()

	// Case 1 : set on empty list should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, lists.ErrOutOfBounds, r.(error))
			}
		}()
		l.Set(0, 0)
	})

	// Case 2 : Set at legal indices.
	l.Add(1)
	l.Add(2)
	l.Add(3)

	l.Set(0, 45)
	assert.Equal(t, types.Integer(45), l.Front())
	l.Set(2, -33)
	assert.Equal(t, types.Integer(-33), l.Back())

}

func TestRemoveBack(t *testing.T) {
	l := New[types.Integer]()

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
	l.Add(22)
	l.Add(23)
	l.Add(234)
	l.Add(-2)

	assert.Equal(t, types.Integer(-2), l.RemoveBack())
	assert.Equal(t, l.Len(), 3)
	assert.Equal(t, types.Integer(234), l.RemoveBack())
	assert.Equal(t, l.Len(), 2)
	assert.Equal(t, types.Integer(23), l.RemoveBack())
	assert.Equal(t, l.Len(), 1)
	assert.Equal(t, types.Integer(22), l.RemoveBack())

}

// TestRemoveAt covers tests for RemoveAt.
func TestRemoveAt(t *testing.T) {

	l := New[types.Integer]()

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
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(6)
	l.Add(9)
	l.Add(80)

	l.RemoveAt(0)
	assert.Equal(t, types.Integer(2), l.Front())
	assert.Equal(t, 5, l.Len())
	l.RemoveAt(2)
	assert.Equal(t, types.Integer(9), l.At(2))
	assert.Equal(t, types.Integer(80), l.RemoveAt(l.Len()-1))
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

// TestRemove covers tests for Remove.
func TestRemove(t *testing.T) {
	l := New[types.Integer]()
	other := New[types.Integer]()

	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(6)
	l.Add(9)

	other.Add(1)
	other.Add(2)
	other.Add(3)
	other.Add(4)
	other.Add(6)
	other.Add(9)

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

// TestIterator covers tests for Iterator.
func TestIterator(t *testing.T) {

	l := New[types.Integer]()

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
		l.Add(types.Integer(i))
	}
	a := l.Collect()
	b := make([]types.Integer, 0)
	it := l.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, types.Integer(1), it.Next())

}

// TestMapFilter covers tests for Map and Filter.
func TestMapFilter(t *testing.T) {
	l := New[types.Integer]()

	for i := 0; i < 6; i++ {
		l.Add(types.Integer(i))
	}

	// Case 1 : Map to a new list.
	other := l.Map(func(e types.Integer) types.Integer { return e + 10 })

	a := []types.Integer{10, 11, 12, 13, 14, 15}
	b := other.Collect()

	assert.ElementsMatch(t, a, b)

	// Case 2 : Filter to create new list.
	c := []types.Integer{0, 2, 4}
	other = l.Filter(func(e types.Integer) bool { return e%2 == 0 })
	d := other.Collect()
	assert.ElementsMatch(t, c, d)

}

// TestClear covers tests for Clear.
func TestClear(t *testing.T) {
	l := New[types.Integer]()

	for i := 0; i < 20; i++ {
		l.Add(types.Integer(i))
	}

	assert.Equal(t, 20, l.Len())
	l.Clear()
	assert.Equal(t, true, l.Empty())

}

// TestString covers tests for String used in printing.
func TestString(t *testing.T) {

	l := New[types.Integer]()

	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(10)

	assert.Equal(t, "[2 3 4 5 10]", fmt.Sprint(l))

}
