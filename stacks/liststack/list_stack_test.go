package liststack

import (
	"testing"

	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/stacks"
	"github.com/phantom820/collections/testutils"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

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
	assert.Equal(t, true, s.Empty())
	s.Add(2, 3)
	assert.Equal(t, types.Int(3), s.Peek())

}

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

	// Case 2 : Pop on a stack with elements.
	s.Add(11, 12)
	assert.Equal(t, types.Int(12), s.Pop())
	assert.Equal(t, types.Int(11), s.Pop())

}

func TestAdd(t *testing.T) {

	s := New[types.Int]()

	// Case 1  Add with empty elements.
	assert.Equal(t, false, s.Add())

	// Case 2 : AddAll add items from an iterable.
	l := forwardlist.New[types.Int]()
	l.Add(1, 2, 3)

	s.AddAll(l)

	assert.Equal(t, types.Int(3), s.Pop())
	assert.Equal(t, types.Int(2), s.Pop())
	assert.Equal(t, types.Int(1), s.Pop())

	// Case 3 : Adding a slice.
	s.Clear()
	slice := []types.Int{2, 4, 6}
	s.Add(slice...)

	assert.Equal(t, types.Int(6), s.Pop())
	assert.Equal(t, types.Int(4), s.Pop())
	assert.Equal(t, types.Int(2), s.Pop())

}

func TestRemove(t *testing.T) {

	s := New[types.Int]()

	s.Add(1, 2, 3, 4)

	// Case 1 : Remove an individual item.
	assert.Equal(t, true, s.Remove(4))
	assert.Equal(t, false, s.Contains(4))
	assert.Equal(t, types.Int(3), s.Peek())

	// Case 2 : RemoveAll remove a number of items from an iterable.
	l := forwardlist.New[types.Int]()
	l.Add(1, 2, 3)

	s.RemoveAll(l)
	assert.Equal(t, 0, s.Len())

}

func TestCollect(t *testing.T) {
	s := New[types.Int]()

	s.Add(1, 2)

	sl := []types.Int{2, 1}
	assert.Equal(t, true, testutils.EqualSlices(sl, s.Collect()))

}

func TestIterator(t *testing.T) {

	s := New[types.Int]()

	// No extensive tests here since this just return an iterator for the list which  has been tested.
	assert.NotNil(t, s.Iterator())

}
