package queue

import (
	"collections/list"
	"collections/wrapper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceAdd(t *testing.T) {

	q := NewSliceQueue[wrapper.Integer]()

	assert.Equal(t, true, q.Empty())
	q.Add(1)
	assert.Equal(t, false, q.Empty())
	assert.Equal(t, 1, q.Len())
	assert.Equal(t, true, q.Contains(1))
	q.Add(2)
	assert.Equal(t, true, q.Contains(2))

	l := list.NewList[wrapper.Integer]()

	for i := 3; i <= 10; i++ {
		l.Add(wrapper.Integer(i))
	}

	q.AddAll(l)
	assert.Equal(t, 10, q.Len())

}

func TestSliceFront(t *testing.T) {

	q := NewSliceQueue[wrapper.Integer]()
	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, NoFrontElementError, r.(error))
			}
		}()

		q.Front()
	})

	for i := 1; i <= 10; i++ {
		q.Add(wrapper.Integer(i))
	}

	assert.Equal(t, wrapper.Integer(1), q.Front())
	assert.Equal(t, wrapper.Integer(1), q.RemoveFront())
	assert.Equal(t, wrapper.Integer(2), q.RemoveFront())
	assert.Equal(t, wrapper.Integer(3), q.RemoveFront())

	q.Clear()
	assert.Equal(t, true, q.Empty())

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, NoFrontElementError, r.(error))
			}
		}()

		q.RemoveFront()
	})

}

func TestSliceIterator(t *testing.T) {
	q := NewSliceQueue[wrapper.Integer]()

	for i := 1; i < 6; i++ {
		q.Add(wrapper.Integer(i))
	}
	a := q.Collect()
	b := make([]wrapper.Integer, 0)
	it := q.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, wrapper.Integer(1), it.Next())

}

func TestSliceRemove(t *testing.T) {

	q := NewSliceQueue[wrapper.Integer]()

	assert.Equal(t, false, q.Remove(22))

	q.Add(1)
	q.Add(2)
	q.Add(4)
	q.Add(5)

	assert.Equal(t, true, q.Remove(5))

	s := list.NewList[wrapper.Integer]()
	s.Add(1)
	s.Add(2)

	q.RemoveAll(s)
	assert.Equal(t, 1, q.Len())

}
