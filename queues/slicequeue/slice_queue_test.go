package slicequeue

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/testutils"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	q := New[types.Int]()

	// Case 1 : Add with no elements.
	assert.Equal(t, true, q.Empty())
	assert.Equal(t, false, q.Add())

	// Case 2 : Add individual elements.
	q.Add(1)
	assert.Equal(t, false, q.Empty())
	assert.Equal(t, 1, q.Len())
	assert.Equal(t, true, q.Contains(1))
	q.Add(2)
	assert.Equal(t, true, q.Contains(2))

	l := forwardlist.New[types.Int](3, 4, 5, 6, 7, 8, 9, 10)

	// Case 3 : Add a number of elements at once.
	q.AddAll(l)
	assert.Equal(t, 10, q.Len())

	// Case 4 : Adding a slice.
	q.Clear()
	s := []types.Int{1, 2, 3, 4}
	q.Add(s...)

	assert.Equal(t, true, testutils.EqualSlices(s, q.Collect()))

}

func TestFront(t *testing.T) {

	q := New[types.Int]()

	// Case 1 : Front on an empty queue should paanic
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		q.Front()
	})

	// Case 2 : Front and RemoveFront should behave accordingly.
	q.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	assert.Equal(t, types.Int(1), q.Front())
	assert.Equal(t, types.Int(1), q.RemoveFront())
	assert.Equal(t, types.Int(2), q.RemoveFront())
	assert.Equal(t, types.Int(3), q.RemoveFront())

	q.Clear()
	assert.Equal(t, true, q.Empty())

	// Case 3 : RemoveFront should panic on an empty queue
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		q.RemoveFront()
	})

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

	assert.Equal(t, true, testutils.EqualSlices(a, b))

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
	// Case 2 : RemoveFront.
	it = q.Iterator()
	t.Run("RemoveFront while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			q.RemoveFront()
			it.Next()
		}
	})
	// Case 3 : Remove.
	it = q.Iterator()
	t.Run("Remove while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			q.Remove()
			it.Next()
		}
	})
	// Case 4 : Clear.
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

	// Case 1 : Removing from empty.
	assert.Equal(t, false, q.Remove(22))

	// Case 2 : Removing from poplated.
	q.Add(1, 2, 4, 5)

	assert.Equal(t, true, q.Remove(5))
	assert.Equal(t, false, q.Contains(5))

	s := forwardlist.New[types.Int](1, 2)

	// Case 3 : Removing multiple elements at once.
	q.RemoveAll(s)

	assert.Equal(t, 1, q.Len())
	assert.Equal(t, types.Int(4), q.Front())

	q.Add(4, 45, 90)
	q.Remove(90)
	assert.Equal(t, false, q.Contains(90))

}

func TestString(t *testing.T) {
	q := New[types.Int]()

	q.Add(1, 2, 3)
	assert.Equal(t, "[1 2 3]", fmt.Sprint(q))
}
