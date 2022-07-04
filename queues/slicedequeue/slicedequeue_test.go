package slicedequeue

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	q := New[types.Int]()

	// Case 1 : Add with no alements
	assert.Equal(t, true, q.Empty())
	assert.Equal(t, false, q.Add())

	// Case 2 : Add individual elements.
	assert.Equal(t, false, q.Contains(1))
	q.Add(1)
	assert.Equal(t, false, q.Empty())
	assert.Equal(t, 1, q.Len())
	assert.Equal(t, true, q.Contains(1))
	q.Add(2)
	assert.Equal(t, true, q.Contains(2))

	l := forwardlist.New[types.Int](3, 4, 5, 6, 7, 8, 9, 10)

	// Case 3 : Add an iterable.
	q.AddAll(l)
	assert.Equal(t, 10, q.Len())

	// Case 4 : Adding a slice.
	q.Clear()
	s := []types.Int{1, 2, 3, 4}
	q.Add(s...)
	assert.ElementsMatch(t, s, q.Collect())

}

func TestFront(t *testing.T) {

	q := New[types.Int]()

	// Case 1 : Front on an empty queue should panic
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		q.Front()
	})

	// Case 2 : Front on a queue with elements.
	q.Add(-1, 2, 3, 4)
	assert.Equal(t, types.Int(-1), q.Front())
	assert.Equal(t, 4, q.Len())

}

func TestAddFront(t *testing.T) {

	q := New[types.Int]()

	// Case 1: Add front for an empty queue.
	assert.Equal(t, false, q.AddFront())
	q.AddFront(23)
	assert.Equal(t, types.Int(23), q.Front())

	// Case 2 : Add front to a queue with elements.
	q.AddFront(1)
	assert.Equal(t, types.Int(1), q.Front())
	q.AddFront(2, 3, 4, 5)
	assert.Equal(t, types.Int(5), q.Front())
	assert.Equal(t, 6, q.Len())
	q.AddFront(6, 7, 8, 9, 10, 11)
	assert.Equal(t, 12, q.Len())

}

func TestRemoveFront(t *testing.T) {

	q := New[types.Int]()

	// Case 1 : RemoveFront on an empty queue should panic
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		q.RemoveFront()
	})

	// Case 2 : RemoveFront on a queue with elements. Shrinking should occur if we remove a number of elemnts
	// and have a lot of free memory.
	q.AddFront(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	assert.Equal(t, types.Int(16), q.RemoveFront())
	assert.Equal(t, types.Int(15), q.RemoveFront())
	assert.Equal(t, types.Int(14), q.RemoveFront())
	assert.Equal(t, types.Int(13), q.RemoveFront())
	assert.Equal(t, types.Int(12), q.RemoveFront())
	assert.Equal(t, types.Int(11), q.RemoveFront())
	assert.Equal(t, types.Int(10), q.RemoveFront())
	assert.Equal(t, types.Int(9), q.RemoveFront())
	assert.Equal(t, types.Int(8), q.RemoveFront())
	assert.Equal(t, 16, len(q.data)) // memory should have shrunk.
	assert.Equal(t, types.Int(7), q.RemoveFront())
	assert.Equal(t, types.Int(6), q.RemoveFront())
	assert.Equal(t, types.Int(5), q.RemoveFront())
	assert.Equal(t, types.Int(4), q.RemoveFront())
	assert.Equal(t, 8, len(q.data)) // memory should have shrunk.
	assert.Equal(t, types.Int(3), q.RemoveFront())
	assert.Equal(t, types.Int(2), q.RemoveFront())
	assert.Equal(t, types.Int(1), q.RemoveFront())
	assert.Equal(t, true, q.Empty())
	q.AddFront(22)
	assert.Equal(t, types.Int(22), q.Front())
	q.AddFront(23, 27)
	assert.Equal(t, types.Int(27), q.Front())
	assert.Equal(t, types.Int(22), q.Back())

}

func TestBack(t *testing.T) {

	q := New[types.Int]()

	// Case 1 : Back on an empty queue should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		q.Back()
	})

	// Case 2 : Back on a queue with elements.
	q.Add(1)
	assert.Equal(t, types.Int(1), q.Back())
	q.Add(23)
	assert.Equal(t, types.Int(23), q.Back())

}

func TestRemoveBack(t *testing.T) {

	q := New[types.Int]()

	// Case 1 : RemoveBack on an empty queue should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		q.RemoveBack()
	})

	// Case 2 : RemoveBack on a queue with elements. Shrinking should occur if we remove a number of elemnts
	// and have a lot of free memory.
	q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	assert.Equal(t, types.Int(16), q.RemoveBack())
	assert.Equal(t, types.Int(15), q.RemoveBack())
	assert.Equal(t, types.Int(14), q.RemoveBack())
	assert.Equal(t, types.Int(13), q.RemoveBack())
	assert.Equal(t, types.Int(12), q.RemoveBack())
	assert.Equal(t, types.Int(11), q.RemoveBack())
	assert.Equal(t, types.Int(10), q.RemoveBack())
	assert.Equal(t, types.Int(9), q.RemoveBack())
	assert.Equal(t, types.Int(8), q.RemoveBack())
	assert.Equal(t, 16, cap(q.data)) // memory should have shrunk.
	assert.Equal(t, types.Int(7), q.RemoveBack())
	assert.Equal(t, types.Int(6), q.RemoveBack())
	assert.Equal(t, types.Int(5), q.RemoveBack())
	assert.Equal(t, types.Int(4), q.RemoveBack())
	assert.Equal(t, 8, cap(q.data)) // memory should have shrunk.
	assert.Equal(t, types.Int(3), q.RemoveBack())
	assert.Equal(t, types.Int(2), q.RemoveBack())
	assert.Equal(t, types.Int(1), q.RemoveBack())
	assert.Equal(t, true, q.Empty())
	q.Add(22)
	assert.Equal(t, types.Int(22), q.Back())
	q.Add(23, 27)
	assert.Equal(t, types.Int(22), q.Front())
	assert.Equal(t, types.Int(27), q.Back())

}

func TestIterator(t *testing.T) {

	q := New[types.Int]()

	// Case 1 : Next on empty queue should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoNextElement, r.(errors.Error).Code())
			}
		}()
		it := q.Iterator()
		it.Next()
	})

	// Case 2 : Iterator should work accordingly on a queue with elements.
	q.Add(1, 2, 3, 4, 5)

	a := q.Collect()
	b := make([]types.Int, 0)
	it := q.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)

}

func TestIteratorConcurrentModification(t *testing.T) {

	q := New[types.String]()
	for i := 1; i <= 20; i++ {
		q.Add(types.String(fmt.Sprint(i)))
	}

	// Recovery for concurrent modifications.
	recovery := func() {
		if r := recover(); r != nil {
			assert.Equal(t, errors.ConcurrentModification, r.(*errors.Error).Code())
		}
	}
	// Case 1 : Add.
	it := q.Iterator()
	t.Run("Add while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			q.Add(types.String("D"))
			it.Next()
		}
	})
	// Case 2 : AddFront.
	it = q.Iterator()
	t.Run("AddFront while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			q.AddFront(types.String("D"))
			it.Next()
		}
	})
	// Case 3 : RemoveFront.
	it = q.Iterator()
	t.Run("RemoveFront while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			q.RemoveFront()
			it.Next()
		}
	})
	// Case 4 : RemoveBack.
	it = q.Iterator()
	t.Run("RemoveBack while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			q.RemoveBack()
			it.Next()
		}
	})
	// Case 5 : Remove.
	it = q.Iterator()
	t.Run("Remove while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			q.Remove()
			it.Next()
		}
	})
	// Case 6 : Clear.
	it = q.Iterator()
	t.Run("Clear while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			q.Clear()
			it.Next()
		}
	})
}

func TestRemove(t *testing.T) {

	q := New[types.Int]()

	// Case 1 : Removing from an empty queue.
	assert.Equal(t, false, q.Remove(22))

	// Case 2 : Removing from poplated.
	q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	assert.Equal(t, false, q.Remove(0))
	assert.Equal(t, true, q.Remove(11))
	assert.Equal(t, types.Int(10), q.Back())
	assert.Equal(t, true, q.Remove(1))
	assert.Equal(t, types.Int(2), q.Front())
	assert.Equal(t, 9, q.Len())
	assert.Equal(t, types.Int(2), q.RemoveFront())
	assert.Equal(t, types.Int(10), q.RemoveBack())
	assert.Equal(t, 7, q.Len())

	l := forwardlist.New[types.Int](4, 5, 6)

	// Case 3 : Removing multiple elements at once.
	q.RemoveAll(l)
	assert.Equal(t, 4, q.Len())
	assert.Equal(t, types.Int(3), q.Front())
	assert.Equal(t, types.Int(9), q.Back())
	q.Remove(3, 9)
	assert.Equal(t, types.Int(7), q.Front())
	assert.Equal(t, types.Int(8), q.Back())
	assert.Equal(t, 2, q.Len())
	q.Remove(7, 8)
	assert.Equal(t, true, q.Empty())
	assert.Equal(t, -1, q.front)
	assert.Equal(t, -1, q.back)

}

func TestString(t *testing.T) {

	q := New[types.Int](1, 2, 3)
	assert.Equal(t, "[1 2 3]", fmt.Sprint(q))

}
