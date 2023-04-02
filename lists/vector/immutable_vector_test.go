package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {

	assert.True(t, Of[int]().Empty())
	assert.False(t, Of(1).Empty())
	assert.Equal(t, 0, Of[int]().Len())

}

func TestImmutableContains(t *testing.T) {

	assert.False(t, Of[int]().Contains(1))
	assert.True(t, Of(1).Contains(1))
	assert.False(t, Of(1, 2).Contains(3))

}

func TestImmutableAt(t *testing.T) {

	assert.Equal(t, 1, Of(1, 2, 3).At(0))
	assert.Equal(t, 2, Of(1, 2, 3).At(1))

}

func TestImmutableEquals(t *testing.T) {

	assert.True(t, Of[int]().Equals(Of[int]()))
	assert.True(t, Of(1).Equals(Of(1)))
	assert.True(t, Of(1, 2).Equals(Of(1, 2)))
	assert.False(t, Of(2, 1).Equals(Of(2)))
	assert.False(t, Of[int]().Equals(Of(2)))

}

func TestImmutableSubList(t *testing.T) {

	assert.Equal(t, []int{}, Of(1, 2, 3, 4).SubList(1, 1).ToSlice())
	assert.Equal(t, []int{2}, Of(1, 2, 3, 4).SubList(1, 2).ToSlice())
	assert.Equal(t, 1, Of(1, 2, 3, 4).SubList(1, 2).Len())
	assert.Equal(t, []int{3, 4, 5, 6}, Of(1, 2, 3, 4, 5, 6, 7).SubList(2, 6).ToSlice())
	assert.Equal(t, 4, Of(1, 2, 3, 4, 5, 6, 7).SubList(2, 6).Len())

}

func TestImmutableForEach(t *testing.T) {

	list := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0
	list.ForEach(func(i int) { sum = sum + i })
	assert.Equal(t, 55, sum)
}

func TestImmutableIterator(t *testing.T) {

	list := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0
	it := list.Iterator()
	for it.HasNext() {
		sum = sum + it.Next()
	}
	assert.Equal(t, 55, sum)
}

func TestImmutableToSlice(t *testing.T) {

	type toSliceTest struct {
		input    ImmutableVector[int]
		expected []int
	}

	toSliceTests := []toSliceTest{
		{
			input:    Of[int](),
			expected: []int{},
		},
		{
			input:    Of(1, 2, 3, 4),
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    Of(1, 2, 3, 4, 5),
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range toSliceTests {
		assert.Equal(t, test.expected, test.input.ToSlice())
	}
}

func TestImmutableString(t *testing.T) {

	assert.Equal(t, "[]", Of[int]().String())
	assert.Equal(t, "[1]", Of(1).String())
	assert.Equal(t, "[1 2]", Of(1, 2).String())
}
