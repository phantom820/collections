package hashset

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

}

func TestImmutableContainsAll(t *testing.T) {

	type containsAllTest struct {
		input    ImmutableHashSet[int]
		elements []int
		expected bool
	}

	containsAllTests := []containsAllTest{
		{
			input:    Of(0, 4, 5),
			elements: []int{},
			expected: true,
		},
		{
			input:    Of(0, 4, 5),
			elements: []int{1},
			expected: false,
		},
		{
			input:    Of(0, 4, 5),
			elements: []int{4, 5},
			expected: true,
		},
		{
			input:    Of(0, 4, 5),
			elements: []int{0, 4, 5},
			expected: true,
		},
		{
			input:    Of(0, 4, 5),
			elements: []int{3},
			expected: false,
		},
		{
			input:    Of(0, 4, 5),
			elements: []int{0, 4, 5, 8},
			expected: false,
		},
	}

	for _, test := range containsAllTests {
		iterable := Of(test.elements...)
		assert.Equal(t, test.expected, test.input.ContainsAll(&iterable))
	}
}

func TestImmutableEquals(t *testing.T) {

	assert.True(t, Of[int]().Equals(Of[int]()))
	assert.True(t, Of(1).Equals(Of(1)))
	assert.True(t, Of(2, 1).Equals(Of(1, 2)))
	assert.False(t, Of(2, 1).Equals(Of(2)))
	assert.False(t, Of[int]().Equals(Of(2)))

}

func TestImmutableForEach(t *testing.T) {

	set := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0
	set.ForEach(func(i int) { sum = sum + i })
	assert.Equal(t, 55, sum)
}

func TestImmutableIterator(t *testing.T) {

	set := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
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
		assert.ElementsMatch(t, test.expected, test.input.ToSlice())
	}
}

func TestImmutableString(t *testing.T) {

	assert.Equal(t, "{}", Of[int]().String())
	assert.Equal(t, "{1}", Of(1).String())
}
