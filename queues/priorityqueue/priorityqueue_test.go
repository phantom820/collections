package priorityqueue

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/queues"
	"github.com/phantom820/collections/types"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	q := New[types.Int](true)

	// Case 1 : Add with no alements
	assert.Equal(t, false, q.Add())
	assert.Equal(t, true, q.Empty())

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

	// Case 4 : Adding a slice should work accordingly
	q.Clear()

	assert.Equal(t, true, q.Empty())

}

func TestFront(t *testing.T) {

	// Min priority queue.
	minQ := New[types.Int](true)

	// Case 1 : Front on empty queue should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, queues.ErrNoFrontElement, r.(error))
			}
		}()
		minQ.Front()
	})

	// Case 2 : Front on populated queue.
	minQ.Add(20)
	assert.Equal(t, types.Int(20), minQ.Front())
	minQ.Add(2)
	assert.Equal(t, types.Int(2), minQ.Front())
	minQ.Add(5, 7, 8, 9, 10, -10)
	assert.Equal(t, types.Int(-10), minQ.Front())

	// Max priority queue.

	maxQ := New[types.Int](false)
	// Case 1 : Front on empty queue should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, queues.ErrNoFrontElement, r.(error))
			}
		}()
		maxQ.Front()
	})

	// Case 2 : Front on populated queue.
	maxQ.Add(20)
	assert.Equal(t, types.Int(20), maxQ.Front())
	maxQ.Add(2)
	assert.Equal(t, types.Int(20), maxQ.Front())
	maxQ.Add(5, 7, 8, 9, 10, -10, 100)
	assert.Equal(t, types.Int(100), maxQ.Front())

}

func TestRemoveFront(t *testing.T) {

	// Min priority queue.
	minQ := New[types.Int](true)

	// Case 1 : Front on empty queue should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, queues.ErrNoFrontElement, r.(error))
			}
		}()
		minQ.RemoveFront()
	})

	minQ.Add(1, 2, 3, 10, 8, 7)
	assert.Equal(t, types.Int(1), minQ.RemoveFront())
	assert.Equal(t, 5, minQ.Len())
	assert.Equal(t, types.Int(2), minQ.RemoveFront())
	assert.Equal(t, 4, minQ.Len())
	assert.Equal(t, types.Int(3), minQ.RemoveFront())
	assert.Equal(t, 3, minQ.Len())
	assert.Equal(t, types.Int(7), minQ.RemoveFront())
	assert.Equal(t, 2, minQ.Len())

}

func TestRemove(t *testing.T) {

	maxQ := New[types.Int](false)

	// Case 1 : Remove from an empty queue.
	assert.Equal(t, false, maxQ.Remove(11))

	// Case 2 : Remove from a queue with elements.
	maxQ.Add(2, 4, 6, 8, 10)
	assert.Equal(t, true, maxQ.Remove(10))
	assert.Equal(t, types.Int(8), maxQ.Front())
	assert.Equal(t, 4, maxQ.Len())

	list := forwardlist.New[types.Int](2, 4, 6, 8)

	maxQ.RemoveAll(list)
	assert.Equal(t, true, maxQ.Empty())

}

func TestIterator(t *testing.T) {

	minQ := New[types.Int](true)

	// Case 1 : Next on empty queue should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, iterator.NoNextElementError, r.(error))
			}
		}()
		it := minQ.Iterator()
		it.Next()
	})

	// Case 2 : Iterator on populated queue
	minQ.Add(1, 2, 3, 4, 5)
	it := minQ.Iterator()

	assert.Equal(t, types.Int(1), it.Next())
	for it.HasNext() {
		it.Next()
	}
	it.Cycle()
	assert.Equal(t, types.Int(1), it.Next())

}

func TestString(t *testing.T) {

	minQ := New[types.Int](true, 1)
	assert.Equal(t, "[1]", fmt.Sprint(minQ))

}
