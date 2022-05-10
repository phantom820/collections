package queue

import (
	"collections/iterator"
	"collections/list"
	"collections/types"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAdd covers tests for Add, Empty, Contains.
func TestAdd(t *testing.T) {

	q := NewListQueue[types.Integer]()

	// q Starts out as empty.
	assert.Equal(t, true, q.Empty())

	// Case 1 : Add individual elements.
	q.Add(1)
	assert.Equal(t, false, q.Empty())
	assert.Equal(t, 1, q.Len())
	assert.Equal(t, true, q.Contains(1))
	q.Add(2)
	assert.Equal(t, true, q.Contains(2))

	l := list.NewForwardList[types.Integer]()

	for i := 3; i <= 10; i++ {
		l.Add(types.Integer(i))
	}

	// Case 2 : Add a number of elements at once.
	q.AddAll(l)
	assert.Equal(t, 10, q.Len())

	// Case 3 : Adding a slice should work accordingly
	q.Clear()

	s := []types.Integer{1, 2, 3, 4}
	q.AddSlice(s)

	assert.ElementsMatch(t, s, q.Collect())

}

// TestFront covers tests for Front and RemoveFront.
func TestFront(t *testing.T) {

	q := NewListQueue[types.Integer]()

	// Case 1 : Front on an empty queue should paanic
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, NoFrontElementError, r.(error))
			}
		}()
		q.Front()
	})

	// Case 2 : Front and RemoveFront should behave accordingly.
	for i := 1; i <= 10; i++ {
		q.Add(types.Integer(i))
	}

	assert.Equal(t, types.Integer(1), q.Front())
	assert.Equal(t, types.Integer(1), q.RemoveFront())
	assert.Equal(t, types.Integer(2), q.RemoveFront())
	assert.Equal(t, types.Integer(3), q.RemoveFront())

	q.Clear()
	assert.Equal(t, true, q.Empty())

	// Case 3 : RemoveFront should panic on an empty queue
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, NoFrontElementError, r.(error))
			}
		}()
		q.RemoveFront()
	})

}

// Covers tests for Iterator.
func TestIterator(t *testing.T) {
	q := NewListQueue[types.Integer]()

	// Case 1 : Next on empty queue should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, iterator.NoNextElementError, r.(error))
			}
		}()
		it := q.Iterator()
		it.Next()
	})

	// Case 2 : Iterator should work accordingly on populated queue.
	for i := 1; i < 6; i++ {
		q.Add(types.Integer(i))
	}
	a := q.Collect()
	b := make([]types.Integer, 0)
	it := q.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, types.Integer(1), it.Next())

}

// TestRemove covers tests for Remove and RemoveAll.
func TestRemove(t *testing.T) {

	q := NewListQueue[types.Integer]()

	// Case 1 : Removing from empty.
	assert.Equal(t, false, q.Remove(22))

	// Case 2 : Removing from poplated.
	q.Add(1)
	q.Add(2)
	q.Add(4)
	q.Add(5)

	assert.Equal(t, true, q.Remove(5))

	s := list.NewForwardList[types.Integer]()
	s.Add(1)
	s.Add(2)

	// Case 3 : Removing multiple elements at once.
	q.RemoveAll(s)
	assert.Equal(t, 1, q.Len())

}

// TestString covers tests for String.
func TestString(t *testing.T) {
	q := NewListQueue[types.Integer]()

	q.Add(1)
	q.Add(2)
	q.Add(3)

	assert.Equal(t, "[1 2 3]", fmt.Sprint(q))
}
