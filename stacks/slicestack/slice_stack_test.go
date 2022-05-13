package slicestack

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/stacks"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

// TestSlicePeek covers tests for Peek and Add.
func TestSlicePeek(t *testing.T) {

	s := New[types.Integer]()

	// Case 1 : Peek on an empty stack should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, stacks.ErrNoTopElement, r.(error))
			}
		}()
		s.Peek()
	})

	// Case 2 : Peek on stack with elements should work accordingly.
	assert.Equal(t, true, s.Empty()) // stack starts ou empty.
	s.Add(2)
	s.Add(3)

	assert.Equal(t, types.Integer(3), s.Peek())

}

// TestSlicePop covers tests for Pop.
func TestSlicePop(t *testing.T) {

	s := New[types.Integer]()

	// Case 1 : Pop on an empty stack should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, stacks.ErrNoTopElement, r.(error))
			}
		}()
		s.Pop()
	})

}

// TestSliceAdd covers tests for AddAll and AddSlice.
func TestSliceAdd(t *testing.T) {

	s := New[types.Integer]()

	// Case 1 : AddAll add items from an iterable should work accordingly.
	l := forwardlist.New[types.Integer]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	s.AddAll(l)

	assert.Equal(t, types.Integer(3), s.Pop())
	assert.Equal(t, types.Integer(2), s.Pop())
	assert.Equal(t, types.Integer(1), s.Pop())

	// Case 2 : AddSlice add items from a slice.
	s.Clear()
	slice := []types.Integer{2, 4, 6}
	s.AddSlice(slice)

	assert.Equal(t, types.Integer(6), s.Pop())
	assert.Equal(t, types.Integer(4), s.Pop())
	assert.Equal(t, types.Integer(2), s.Pop())

}

// TestSliceRemove covers tests for Remove, RemoveAll and Contains.
func TestSliceRemove(t *testing.T) {

	s := New[types.Integer]()

	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)

	// Case 1 : Remove an individual item.
	assert.Equal(t, true, s.Contains(4))
	assert.Equal(t, true, s.Remove(4))
	assert.Equal(t, false, s.Remove(41))
	assert.Equal(t, false, s.Contains(4))
	assert.Equal(t, types.Integer(3), s.Peek())

	// Case 2 : RemoveAll remove a number of items from an iterable.
	l := forwardlist.New[types.Integer]()

	l.Add(1)
	l.Add(2)
	l.Add(3)

	s.RemoveAll(l)
	assert.Equal(t, 0, s.Len())

}

// TestSliceCollect covers tests for collect.
func TestSliceCollect(t *testing.T) {

	s := New[types.Integer]()

	s.Add(1)
	s.Add(2)
	sl := []types.Integer{2, 1}

	assert.ElementsMatch(t, sl, s.Collect())

}

func TestSliceIterator(t *testing.T) {

	s := New[types.Integer]()

	// Case 1 : Next on empty stack should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, iterator.NoNextElementError, r.(error))
			}
		}()
		it := s.Iterator()
		it.Next()
	})

	// Case 2 : Iterator should work accordingly on populated stack.
	for i := 1; i < 6; i++ {
		s.Add(types.Integer(i))
	}
	a := s.Collect()
	b := make([]types.Integer, 0)
	it := s.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, types.Integer(5), it.Next())

}

func TestString(t *testing.T) {
	s := New[types.Integer](1, 2, 3, 4)
	assert.Equal(t, "[1 2 3 4]", fmt.Sprint(s))
}
