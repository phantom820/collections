package set

import (
	"collections/list"
	"collections/wrapper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	s := NewHashSet[wrapper.Integer]()

	// adding of individual elements, and try adding alements already in the set.
	assert.Equal(t, true, s.Empty())
	s.Add(1)
	assert.Equal(t, 1, s.Len())
	assert.Equal(t, false, s.Add(1))
	assert.Equal(t, 1, s.Len())
	s.Add(2)
	assert.Equal(t, 2, s.Len())
	assert.Equal(t, true, s.Contains(1))
	assert.Equal(t, true, s.Contains(2))

	// adding of elements coming from another iterable
	s = NewHashSet[wrapper.Integer]()
	l := list.NewList[wrapper.Integer]()
	for i := 0; i < 10; i++ {
		l.Add(wrapper.Integer(i))
	}

	s.AddAll(l)
	assert.Equal(t, l.Len(), s.Len())

	// should contain all the added elements
	it := l.Iterator()
	for it.HasNext() {
		assert.Equal(t, true, s.Contains(it.Next()))
	}
}

func TestIterator(t *testing.T) {

	s := NewHashSet[wrapper.Integer]()
	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, NoNextElementError, r.(error))
			}
		}()
		it := s.Iterator()
		it.Next()
	})

	for i := 1; i < 6; i++ {
		s.Add(wrapper.Integer(i))
	}
	a := s.Collect()
	b := make([]wrapper.Integer, 0)
	it := s.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	it.Cycle()
	assert.Equal(t, wrapper.Integer(1), it.Next())

	s.Clear()
	assert.Equal(t, true, s.Empty())

}

func TestRemove(t *testing.T) {

	s := NewHashSet[wrapper.Integer]()

	assert.Equal(t, false, s.Remove(1))
	s.Add(1)
	assert.Equal(t, true, s.Remove(1))
	assert.Equal(t, 0, s.Len())

	l := list.NewList[wrapper.Integer]()
	for i := 1; i <= 10; i++ {
		s.Add(wrapper.Integer(i))
		l.Add(wrapper.Integer(i))
	}

	assert.Equal(t, 10, s.Len())
	s.RemoveAll(l)
	assert.Equal(t, 0, s.Len())

}

func TestMapFilter(t *testing.T) {

	s := NewHashSet[wrapper.Integer]()
	for i := 0; i < 6; i++ {
		s.Add(wrapper.Integer(i))
	}
	other := s.Map(func(e wrapper.Integer) wrapper.Integer { return e + 10 })

	a := []wrapper.Integer{10, 11, 12, 13, 14, 15}
	b := other.Collect()

	assert.ElementsMatch(t, a, b)

	c := []wrapper.Integer{0, 2, 4}
	other = s.Filter(func(e wrapper.Integer) bool { return e%2 == 0 })
	d := other.Collect()

	assert.ElementsMatch(t, c, d)

}

func TestEquals(t *testing.T) {

	s := NewHashSet[wrapper.Integer]()
	other := NewHashSet[wrapper.Integer]()
	assert.Equal(t, true, s.Equals(s))
	assert.Equal(t, true, s.Equals(other)) // Two empty sets are equal.

	s.Add(1)
	assert.Equal(t, false, s.Equals(other))
	other.Add(1)
	assert.Equal(t, true, s.Equals(other))
	s.Add(2)
	other.Add(3)
	assert.Equal(t, false, s.Equals(other))

}

func TestString(t *testing.T) {

	s := NewHashSet[wrapper.Integer]()

	assert.Equal(t, "{}", fmt.Sprint(s))
	s.Add(1)
	assert.Equal(t, "{1}", fmt.Sprint(s))

}
