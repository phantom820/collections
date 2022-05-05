package list

import (
	"collections/wrapper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFront(t *testing.T) {
	l := NewList[wrapper.Integer]()

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, EmptyListError, r.(error))
			}
		}()

		l.Front()
	})

	assert.Equal(t, 0, l.Len())
	l.AddFront(22)
	assert.Equal(t, wrapper.Integer(22), l.Front())
	assert.Equal(t, 1, l.Len())
	l.AddFront(23)
	assert.Equal(t, wrapper.Integer(23), l.Front())
	assert.Equal(t, 2, l.Len())

}

func TestAddBack(t *testing.T) {
	l := NewList[wrapper.Integer]()

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, EmptyListError, r.(error))
			}
		}()

		l.Back()
	})

	assert.Equal(t, 0, l.Len())
	l.AddBack(22)
	assert.Equal(t, wrapper.Integer(22), l.Back())
	assert.Equal(t, 1, l.Len())
	l.AddBack(23)
	assert.Equal(t, wrapper.Integer(23), l.Back())
	assert.Equal(t, 2, l.Len())

}

func TestAt(t *testing.T) {
	l := NewList[wrapper.Integer]()

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, OutOfBoundsError, r.(error))
			}
		}()

		l.At(0)
	})

	l.Add(1)
	l.Add(2)
	l.Add(3)
	assert.Equal(t, wrapper.Integer(1), l.At(0))
	assert.Equal(t, wrapper.Integer(2), l.At(1))
	assert.Equal(t, wrapper.Integer(3), l.At(2))

}

func TestEquals(t *testing.T) {

	l := NewList[wrapper.Integer]()
	other := NewList[wrapper.Integer]()

	assert.Equal(t, true, l.Equals(other)) // two empty list are equal.

	for i := 1; i < 6; i++ {
		other.Add(wrapper.Integer(i))
	}

	assert.Equal(t, true, l.Equals(l))
	assert.Equal(t, false, l.Equals(other))

	for i := 1; i < 6; i++ {
		l.Add(wrapper.Integer(i))
	}

	assert.Equal(t, true, l.Equals(other))
	assert.Equal(t, true, other.Equals(l))

	l.Clear()
	for i := 1; i < 6; i++ {
		l.Add(wrapper.Integer(i + 1))
	}

	assert.Equal(t, false, other.Equals(l))

}

func TestAddAll(t *testing.T) {
	l := NewList[wrapper.Integer]()
	other := NewList[wrapper.Integer]()

	for i := 1; i < 6; i++ {
		other.Add(wrapper.Integer(i))
	}

	l.AddAll(other)
	assert.Equal(t, true, l.Equals(other))

}

func TestAddAt(t *testing.T) {
	l := NewList[wrapper.Integer]()

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, OutOfBoundsError, r.(error))
			}
		}()

		l.AddAt(0, 0)
	})

	l.Add(10)
	l.Add(20)
	l.Add(30)

	l.AddAt(0, 22)
	assert.Equal(t, wrapper.Integer(22), l.At(0))
	l.AddAt(l.Len()-1, 25)
	assert.Equal(t, wrapper.Integer(25), l.At(l.Len()-1))
	l.AddAt(2, -5)
	assert.Equal(t, wrapper.Integer(-5), l.At(2))

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
	assert.Equal(t, wrapper.Integer(5), l.At(2))
	l.Swap(0, l.Len()-1)
	assert.Equal(t, wrapper.Integer(10), l.At(0))
	l.Swap(l.Len()-1, 0)
	assert.Equal(t, wrapper.Integer(10), l.At(l.Len()-1))

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, OutOfBoundsError, r.(error))
			}
		}()
		l.Swap(-1, 0)
	})

}

func TestRemoveFront(t *testing.T) {
	l := NewList[wrapper.Integer]()

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, EmptyListError, r.(error))
			}
		}()
		l.RemoveFront()
	})
	l.Add(22)
	l.Add(23)
	l.Add(234)
	assert.Equal(t, wrapper.Integer(22), l.RemoveFront())
	assert.Equal(t, wrapper.Integer(23), l.RemoveFront())
	assert.Equal(t, wrapper.Integer(234), l.RemoveFront())

}

func TestSet(t *testing.T) {

	l := NewList[wrapper.Integer]()

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, OutOfBoundsError, r.(error))
			}
		}()
		l.Set(0, 0)
	})

	l.Add(1)
	l.Add(2)
	l.Add(3)

	l.Set(0, 45)
	assert.Equal(t, wrapper.Integer(45), l.Front())
	l.Set(2, -33)
	assert.Equal(t, wrapper.Integer(-33), l.Back())

}

func TestRemoveBack(t *testing.T) {
	l := NewList[wrapper.Integer]()

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, EmptyListError, r.(error))
			}
		}()
		l.RemoveBack()
	})

	l.Add(22)
	l.Add(23)
	l.Add(234)
	l.Add(-2)

	assert.Equal(t, wrapper.Integer(-2), l.RemoveBack())
	assert.Equal(t, l.Len(), 3)
	assert.Equal(t, wrapper.Integer(234), l.RemoveBack())
	assert.Equal(t, l.Len(), 2)
	assert.Equal(t, wrapper.Integer(23), l.RemoveBack())
	assert.Equal(t, l.Len(), 1)
	assert.Equal(t, wrapper.Integer(22), l.RemoveBack())

}

func TestRemoveAt(t *testing.T) {
	l := NewList[wrapper.Integer]()

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, EmptyListError, r.(error))
			}
		}()
		l.RemoveAt(0)
	})

	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(6)
	l.Add(9)
	l.Add(80)

	l.RemoveAt(0)
	assert.Equal(t, wrapper.Integer(2), l.Front())
	assert.Equal(t, 5, l.Len())
	l.RemoveAt(2)
	assert.Equal(t, wrapper.Integer(9), l.At(2))
	assert.Equal(t, wrapper.Integer(80), l.RemoveAt(l.Len()-1))
	assert.Equal(t, 3, l.Len())

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, OutOfBoundsError, r.(error))
			}
		}()
		l.RemoveAt(-1)
	})

}
func TestRemove(t *testing.T) {
	l := NewList[wrapper.Integer]()
	other := NewList[wrapper.Integer]()

	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, EmptyListError, r.(error))
			}
		}()
		l.Remove(2)
	})

	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(6)
	l.Add(9)

	other.Add(1)
	other.Add(2)
	other.Add(3)
	other.Add(4)
	other.Add(6)
	other.Add(9)

	assert.Equal(t, true, l.Contains(1))
	l.Remove(1)
	assert.Equal(t, false, l.Contains(1))
	assert.Equal(t, true, l.Contains(9))
	l.Remove(9)
	assert.Equal(t, false, l.Contains(9))
	l.Remove(3)
	assert.Equal(t, false, l.Contains(3))
	l.RemoveAll(other)
	assert.Equal(t, true, l.Empty())
}

func TestIterator(t *testing.T) {
	l := NewList[wrapper.Integer]()
	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, NoNextElementError, r.(error))
			}
		}()
		it := l.Iterator()
		it.Next()
	})
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

func TestString(t *testing.T) {

	l := NewList[wrapper.Integer]()

	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(10)

	assert.Equal(t, "{2, 3, 4, 5, 10}", fmt.Sprint(l))

}
