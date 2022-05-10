package stack

import (
	"testing"

	"github.com/phantom820/collections/list"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

// TestPeek covers tests for Peek and Add.
func TestPeek(t *testing.T) {

	s := NewListStack[types.Integer]()

	// Case 1 : Peek on an empty stack should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, NoTopElementError, r.(error))
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

// TestPop covers tests for Pop.
func TestPop(t *testing.T) {

	s := NewListStack[types.Integer]()

	// Case 1 : Pop on an empty stack should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, NoTopElementError, r.(error))
			}
		}()
		s.Pop()
	})

}

// TestAdd covers tests for AddAll and AddSlice.
func TestAdd(t *testing.T) {

	s := NewListStack[types.Integer]()

	// Case 1 : AddAll add items from an iterable should work accordingly.
	l := list.NewForwardList[types.Integer]()
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

// TestRemove covers tests for Remove, RemoveAll and Contains.
func TestRemove(t *testing.T) {

	s := NewListStack[types.Integer]()

	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)

	// Case 1 : Remove an individual item.
	assert.Equal(t, true, s.Remove(4))
	assert.Equal(t, false, s.Contains(4))
	assert.Equal(t, types.Integer(3), s.Peek())

	// Case 2 : RemoveAll remove a number of items from an iterable.
	l := list.NewForwardList[types.Integer]()

	l.Add(1)
	l.Add(2)
	l.Add(3)

	s.RemoveAll(l)
	assert.Equal(t, 0, s.Len())

}

// TestCollect covers tests for collect.
func TestCollect(t *testing.T) {

	s := NewListStack[types.Integer]()

	s.Add(1)
	s.Add(2)
	sl := []types.Integer{2, 1}

	assert.ElementsMatch(t, sl, s.Collect())

}

func TestIterator(t *testing.T) {

	s := NewListStack[types.Integer]()

	// No extensive tests here since this just return an iterator.
	assert.NotNil(t, s.Iterator())

}
