package treeset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImmutableOf(t *testing.T) {

	assert.True(t, ImmutableOf[int](lessThanInt).Empty())
	assert.False(t, ImmutableOf(lessThanInt, 1).Empty())
	assert.Equal(t, 0, ImmutableOf[int](lessThanInt).Len())

}

func TestImmutableContains(t *testing.T) {

	assert.False(t, ImmutableOf[int](lessThanInt).Contains(1))
	assert.True(t, ImmutableOf(lessThanInt, 1).Contains(1))

}

func TestImmutableContainsAll(t *testing.T) {

	type containsAllTest struct {
		input    ImmutableTreeSet[int]
		elements []int
		expected bool
	}

	containsAllTests := []containsAllTest{
		{
			input:    ImmutableOf(lessThanInt, 0, 4, 5),
			elements: []int{},
			expected: true,
		},
		{
			input:    ImmutableOf(lessThanInt, 0, 4, 5),
			elements: []int{1},
			expected: false,
		},
		{
			input:    ImmutableOf(lessThanInt, 0, 4, 5),
			elements: []int{4, 5},
			expected: true,
		},
		{
			input:    ImmutableOf(lessThanInt, 0, 4, 5),
			elements: []int{0, 4, 5},
			expected: true,
		},
		{
			input:    ImmutableOf(lessThanInt, 0, 4, 5),
			elements: []int{3},
			expected: false,
		},
		{
			input:    ImmutableOf(lessThanInt, 0, 4, 5),
			elements: []int{0, 4, 5, 8},
			expected: false,
		},
	}

	for _, test := range containsAllTests {
		iterable := Of(lessThanInt, test.elements...)
		assert.Equal(t, test.expected, test.input.ContainsAll(&iterable))
	}
}

func TestImmutableEquals(t *testing.T) {

	assert.True(t, ImmutableOf[int](lessThanInt).Equals(ImmutableOf[int](lessThanInt)))
	assert.True(t, ImmutableOf(lessThanInt, 1).Equals(ImmutableOf(lessThanInt, 1)))
	assert.True(t, ImmutableOf(lessThanInt, 2, 1).Equals(ImmutableOf(lessThanInt, 1, 2)))
	assert.False(t, ImmutableOf(lessThanInt, 2, 1).Equals(ImmutableOf(lessThanInt, 2)))
	assert.False(t, ImmutableOf[int](lessThanInt).Equals(ImmutableOf(lessThanInt, 2)))

}

func TestImmutableForEach(t *testing.T) {

	set := ImmutableOf(lessThanInt, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0
	set.ForEach(func(i int) { sum = sum + i })
	assert.Equal(t, 55, sum)
}

func TestImmutableIterator(t *testing.T) {

	set := ImmutableOf(lessThanInt, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	numbers := []int{}
	it := set.Iterator()
	for it.HasNext() {
		numbers = append(numbers, it.Next())
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, numbers)
}

func TestImmutableToSlice(t *testing.T) {

	type toSliceTest struct {
		input    ImmutableTreeSet[int]
		expected []int
	}

	toSliceTests := []toSliceTest{
		{
			input:    ImmutableOf[int](lessThanInt),
			expected: []int{},
		},
		{
			input:    ImmutableOf(lessThanInt, 1, 2, 3, 4),
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    ImmutableOf(lessThanInt, 1, 2, 3, 4, 5),
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range toSliceTests {
		assert.ElementsMatch(t, test.expected, test.input.ToSlice())
	}
}

func TestImmutableString(t *testing.T) {

	assert.Equal(t, "{}", ImmutableOf(lessThanInt).String())
	assert.Equal(t, "{1, 2, 3}", ImmutableOf(lessThanInt, 1, 2, 3).String())
}
