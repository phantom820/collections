package hashset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImmutableOf(t *testing.T) {

	assert.True(t, ImmutableOf[int]().Empty())
	assert.False(t, ImmutableOf(1).Empty())
	assert.Equal(t, 0, ImmutableOf[int]().Len())

}

func TestImmutableContains(t *testing.T) {

	assert.False(t, ImmutableOf[int]().Contains(1))
	assert.True(t, ImmutableOf(1).Contains(1))

}

func TestImmutableEquals(t *testing.T) {

	assert.True(t, ImmutableOf[int]().Equals(ImmutableOf[int]()))
	assert.True(t, ImmutableOf(1).Equals(ImmutableOf(1)))
	assert.True(t, ImmutableOf(2, 1).Equals(ImmutableOf(1, 2)))
	assert.False(t, ImmutableOf(2, 1).Equals(ImmutableOf(2)))
	assert.False(t, ImmutableOf[int]().Equals(ImmutableOf(2)))

}

func TestImmutableForEach(t *testing.T) {

	set := ImmutableOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0
	set.ForEach(func(i int) { sum = sum + i })
	assert.Equal(t, 55, sum)
}

func TestImmutableIterator(t *testing.T) {

	set := ImmutableOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0
	it := set.Iterator()
	for it.HasNext() {
		sum = sum + it.Next()
	}
	assert.Equal(t, 55, sum)
}

func TestImmutableToSlice(t *testing.T) {

	type toSliceTest struct {
		input    ImmutableHashSet[int]
		expected []int
	}

	toSliceTests := []toSliceTest{
		{
			input:    ImmutableOf[int](),
			expected: []int{},
		},
		{
			input:    ImmutableOf(1, 2, 3, 4),
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    ImmutableOf(1, 2, 3, 4, 5),
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range toSliceTests {
		assert.ElementsMatch(t, test.expected, test.input.ToSlice())
	}
}

func TestImmutableString(t *testing.T) {

	assert.Equal(t, "{}", ImmutableOf[int]().String())
	assert.Equal(t, "{1}", ImmutableOf(1).String())
}
