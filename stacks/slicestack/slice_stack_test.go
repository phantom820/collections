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

// TestPeek covers tests for Peek and Add.
func TestPeek(t *testing.T) {

	s := New[types.Int]()

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

	assert.Equal(t, types.Int(3), s.Peek())

}

// TestPop covers tests for Pop.
func TestPop(t *testing.T) {

	s := New[types.Int]()

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

// TestAdd covers tests for AddAll and AddSlice.
func TestAdd(t *testing.T) {

	s := New[types.Int]()

	// Case 1 : Add with no elements.
	assert.Equal(t, false, s.Add())

	// Case 2 : AddAll add items from an iterable should work accordingly.
	l := forwardlist.New[types.Int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	s.AddAll(l)

	assert.Equal(t, types.Int(3), s.Pop())
	assert.Equal(t, types.Int(2), s.Pop())
	assert.Equal(t, types.Int(1), s.Pop())

	// Case 3 : AddSlice add items from a slice.
	s.Clear()
	slice := []types.Int{2, 4, 6}
	s.Add(slice...)

	assert.Equal(t, types.Int(6), s.Pop())
	assert.Equal(t, types.Int(4), s.Pop())
	assert.Equal(t, types.Int(2), s.Pop())

}

// TestRemove covers tests for Remove, RemoveAll and Contains.
func TestRemove(t *testing.T) {

	s := New[types.Int]()

	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)

	// Case 1 : Remove an individual item.
	assert.Equal(t, true, s.Contains(4))
	assert.Equal(t, true, s.Remove(4))
	assert.Equal(t, false, s.Remove(41))
	assert.Equal(t, false, s.Contains(4))
	assert.Equal(t, types.Int(3), s.Peek())

	// Case 2 : RemoveAll remove a number of items from an iterable.
	l := forwardlist.New[types.Int]()

	l.Add(1)
	l.Add(2)
	l.Add(3)

	s.RemoveAll(l)
	assert.Equal(t, 0, s.Len())

}

// TestCollect covers tests for collect.
func TestCollect(t *testing.T) {

	s := New[types.Int]()

	s.Add(1)
	s.Add(2)
	sl := []types.Int{2, 1}

	assert.ElementsMatch(t, sl, s.Collect())

}

func TestIterator(t *testing.T) {

	s := New[types.Int]()

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
		s.Add(types.Int(i))
	}
	a := s.Collect()
	b := make([]types.Int, 0)
	it := s.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, types.Int(5), it.Next())

}

func TestString(t *testing.T) {
	s := New[types.Int](1, 2, 3, 4)
	assert.Equal(t, "[1 2 3 4]", fmt.Sprint(s))
}
