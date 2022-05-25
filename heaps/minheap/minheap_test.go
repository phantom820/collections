package minheap

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/heaps"
	"github.com/phantom820/collections/types"
	"github.com/stretchr/testify/assert"
)

// TestInsert also covers tests for Top, Contains, Clear , Empty.
func TestInsert(t *testing.T) {

	heap := New[types.Int]()

	assert.Equal(t, true, heap.Empty())

	// Case 1 : Insert to empty heap.
	heap.Insert(2)
	assert.Equal(t, 1, heap.Len())
	assert.Equal(t, types.Int(2), heap.Top())

	// Case 2 : Insert to heap with an element.
	assert.Equal(t, false, heap.Search(3))
	heap.Insert(3)
	assert.Equal(t, types.Int(2), heap.Top())
	assert.Equal(t, true, heap.Search(2))

	heap.Insert(-1)
	assert.Equal(t, types.Int(-1), heap.Top())

	heap.Insert(5)
	assert.Equal(t, types.Int(-1), heap.Top())

	heap.Clear()

	// Case 1 : Empty heap has no top.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, heaps.ErrEmptyHeap, r.(error))
			}
		}()
		heap.Top()
	})
	assert.Equal(t, true, heap.Empty())

}

func TestDeleteTop(t *testing.T) {

	heap := New[types.Int]()

	// Case 1 : Empty heap has no top.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, heaps.ErrEmptyHeap, r.(error))
			}
		}()
		heap.DeleteTop()
	})

	heap.Insert(11)
	heap.Insert(12)
	heap.Insert(0)
	heap.Insert(23)

	assert.Equal(t, types.Int(0), heap.DeleteTop())
	assert.Equal(t, types.Int(11), heap.DeleteTop())
	assert.Equal(t, types.Int(12), heap.DeleteTop())
	assert.Equal(t, types.Int(23), heap.DeleteTop())

	assert.Equal(t, true, heap.Empty())

}

func TestUpdate(t *testing.T) {

	heap := New[types.Int]()

	heap.Insert(11)
	heap.Insert(12)
	heap.Insert(0)
	heap.Insert(23)

	// Case 1: Update an element we can find.
	assert.Equal(t, types.Int(0), heap.Top())
	heap.Update(0, 92)
	assert.Equal(t, types.Int(11), heap.Top())

	// Case 2 : Update to an element not in the queue.
	assert.Equal(t, false, heap.Update(-2, 10))

}

func TestStrin(t *testing.T) {

	heap := New[types.Int]()

	heap.Insert(1)

	assert.Equal(t, "[1]", fmt.Sprint(heap))

}
