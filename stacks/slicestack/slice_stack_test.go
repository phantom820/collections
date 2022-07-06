package slicestack

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

func TestPeek(t *testing.T) {

	s := New[types.Int]()

	// Case 1 : Peek on an empty stack should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		s.Peek()
	})

	// Case 2 : Peek on stack with elements.
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
				assert.Equal(t, errors.NoSuchElement, r.(errors.Error).Code())
			}
		}()
		s.Pop()
	})

	// Case 2 : Pop on a stack with elements.
	s.Add(1, 2, 3)
	assert.Equal(t, types.Int(3), s.Pop())
	assert.Equal(t, types.Int(2), s.Pop())
	assert.Equal(t, types.Int(1), s.Pop())

}

func TestAdd(t *testing.T) {

	s := New[types.Int]()

	// Case 1 : Add with no elements.
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
	assert.Equal(t, true, s.Contains(4))
	assert.Equal(t, true, s.Remove(4))
	assert.Equal(t, false, s.Remove(41))
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

	assert.ElementsMatch(t, sl, s.Collect())

}

func TestIterator(t *testing.T) {

	s := New[types.Int]()

	// Case 1 : Next on empty stack should panic.
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, errors.NoNextElement, r.(errors.Error).Code())
			}
		}()
		it := s.Iterator()
		it.Next()
	})

	// Case 2 : Iterator on a populated stack.
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

}

func TestIteratorConcurrentModification(t *testing.T) {

	s := New[types.String]()
	for i := 1; i <= 20; i++ {
		s.Add(types.String(fmt.Sprint(i)))
	}

	// Recovery for concurrent modifications.
	recovery := func() {
		if r := recover(); r != nil {
			assert.Equal(t, errors.ConcurrentModification, r.(*errors.Error).Code())
		}
	}
	// Case 1 : Add.
	it := s.Iterator()
	t.Run("Add while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			s.Add(types.String("D"))
			it.Next()
		}
	})
	// Case 2 : RemoveFront.
	it = s.Iterator()
	t.Run("Remove while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			s.Remove()
			it.Next()
		}
	})
	// Case 3 : Pop.
	it = s.Iterator()
	t.Run("Pop while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			s.Pop()
			it.Next()
		}
	})
	// Case 4 : Clear.
	it = s.Iterator()
	t.Run("Clear while iterating", func(t *testing.T) {
		defer recovery()
		for it.HasNext() {
			s.Clear()
			it.Next()
		}
	})
}

func TestString(t *testing.T) {
	s := New[types.Int](1, 2, 3, 4)
	assert.Equal(t, "[1 2 3 4]", fmt.Sprint(s))
}
