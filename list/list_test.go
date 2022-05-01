package list

import (
	"collections/wrapper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFront(t *testing.T) {
	l := NewList[wrapper.Integer]()

	assert.Equal(t, wrapper.Integer(0), l.Front())
	l.AddFront(22)
	assert.Equal(t, wrapper.Integer(22), l.Front())
	l.AddFront(23)
	assert.Equal(t, wrapper.Integer(23), l.Front())

}

func TestAddBack(t *testing.T) {
	l := NewList[wrapper.Integer]()

	assert.Equal(t, wrapper.Integer(0), l.Back())
	l.AddBack(22)
	assert.Equal(t, wrapper.Integer(22), l.Back())
	l.AddBack(23)
	assert.Equal(t, wrapper.Integer(23), l.Back())

}

func TestAddAt(t *testing.T) {
	l := NewList[wrapper.Integer]()

	err := l.AddAt(0, 1)
	assert.Equal(t, OutOfBoundsError, err)
	l.Add(1)
	l.Add(22)
	l.Add(23)
	l.Add(24)
	err = l.AddAt(2, 45)
	assert.Nil(t, err)
	assert.Equal(t, true, l.Contains(45))
}

func TestAt(t *testing.T) {
	l := NewList[wrapper.Integer]()

	_, err := l.At(0)
	assert.Equal(t, OutOfBoundsError, err)
	l.Add(22)
	v, _ := l.At(0)
	assert.Equal(t, wrapper.Integer(22), v)

}

func TestSwap(t *testing.T) {
	l := NewList[wrapper.Integer]()

	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(10)

	l.Swap(0, 1)
	assert.Equal(t, wrapper.Integer(3), l.Front())
	l.Swap(2, 3)
	l.Swap(1, 0)
	v, _ := l.At(2)

	assert.Equal(t, wrapper.Integer(5), v)
}

func TestRemoveFront(t *testing.T) {
	l := NewList[wrapper.Integer]()
	_, err := l.RemoveFront()
	assert.Equal(t, EmptyListError, err)
	l.Add(22)
	l.Add(23)
	v, _ := l.RemoveFront()
	assert.Equal(t, wrapper.Integer(22), v)
	v, _ = l.RemoveFront()
	assert.Equal(t, wrapper.Integer(23), v)

}

func TestRemoveBack(t *testing.T) {
	l := NewList[wrapper.Integer]()
	_, err := l.RemoveBack()
	assert.Equal(t, EmptyListError, err)
	l.Add(22)
	l.Add(33)
	l.Add(23)
	v, _ := l.RemoveBack()
	assert.Equal(t, wrapper.Integer(23), v)
	v, _ = l.RemoveBack()
	assert.Equal(t, wrapper.Integer(33), v)

}

func TestRemove(t *testing.T) {
	l := NewList[wrapper.Integer]()

	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(6)
	l.Add(9)

	assert.Equal(t, true, l.Contains(1))
	l.Remove(1)
	assert.Equal(t, false, l.Contains(1))
	assert.Equal(t, true, l.Contains(9))
	l.Remove(9)
	assert.Equal(t, false, l.Contains(9))
	l.Remove(3)
	assert.Equal(t, false, l.Contains(3))

}

func TestIterator(t *testing.T) {
	l := NewList[wrapper.Integer]()
	for i := 1; i < 6; i++ {
		l.Add(wrapper.Integer(i))
	}
	a := l.Collect()
	b := make([]wrapper.Integer, 0)
	it := l.Iterator()
	for it.HasNext() {
		b = append(b, it.Next())
	}
	assert.ElementsMatch(t, a, b)
	assert.Equal(t, wrapper.Integer(0), it.Next())
	it.Cycle()
	assert.Equal(t, wrapper.Integer(1), it.Next())

}

func TestMapFilter(t *testing.T) {
	l := NewList[wrapper.Integer]()
	for i := 0; i < 6; i++ {
		l.Add(wrapper.Integer(i))
	}
	other := l.Map(func(e wrapper.Integer) wrapper.Integer { return e + 10 })

	a := []wrapper.Integer{10, 11, 12, 13, 14, 15}
	b := other.Collect()

	assert.ElementsMatch(t, a, b)

	c := []wrapper.Integer{0, 2, 4}
	other = l.Filter(func(e wrapper.Integer) bool { return e%2 == 0 })
	d := other.Collect()
	assert.ElementsMatch(t, c, d)
}

func TestClear(t *testing.T) {
	l := NewList[wrapper.Integer]()

	for i := 0; i < 20; i++ {
		l.Add(wrapper.Integer(i))
	}
	assert.Equal(t, 20, l.Len())
	l.Clear()
	assert.Equal(t, true, l.Empty())

}
