package treeset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {

	assert.Equal(t, []string{}, New(lessThan).treeMap.Keys())
	assert.Equal(t, []string{"A"}, New(lessThan, "A").treeMap.Keys())

}

func TestImmutableContains(t *testing.T) {

	assert.False(t, Of(lessThanInt).Contains(1))
	assert.True(t, Of(lessThanInt, 1).Contains(1))

}

func TestImmutableContainsAll(t *testing.T) {

	type containsAllTest struct {
		input    ImmutableTreeSet[int]
		elements []int
		expected bool
	}

	containsAllTests := []containsAllTest{
		{
			input:    Of(lessThanInt, 0, 4, 5),
			elements: []int{},
			expected: true,
		},
		{
			input:    Of(lessThanInt, 0, 4, 5),
			elements: []int{1},
			expected: false,
		},
		{
			input:    Of(lessThanInt, 0, 4, 5),
			elements: []int{4, 5},
			expected: true,
		},
		{
			input:    Of(lessThanInt, 0, 4, 5),
			elements: []int{0, 4, 5},
			expected: true,
		},
		{
			input:    Of(lessThanInt, 0, 4, 5),
			elements: []int{3},
			expected: false,
		},
		{
			input:    Of(lessThanInt, 0, 4, 5),
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

	assert.True(t, Of[int](lessThanInt).Equals(Of[int](lessThanInt)))
	assert.True(t, Of(lessThanInt, 1).Equals(Of(lessThanInt, 1)))
	assert.True(t, Of(lessThanInt, 2, 1).Equals(Of(lessThanInt, 1, 2)))
	assert.False(t, Of(lessThanInt, 2, 1).Equals(Of(lessThanInt, 2)))
	assert.False(t, Of[int](lessThanInt).Equals(Of(lessThanInt, 2)))

}

func TestImmutableForEach(t *testing.T) {

	set := Of(lessThanInt, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0
	set.ForEach(func(i int) { sum = sum + i })
	assert.Equal(t, 55, sum)
}

func TestImmutableIterator(t *testing.T) {

	set := Of(lessThanInt, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
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
			input:    Of[int](lessThanInt),
			expected: []int{},
		},
		{
			input:    Of(lessThanInt, 1, 2, 3, 4),
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    Of(lessThanInt, 1, 2, 3, 4, 5),
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range toSliceTests {
		assert.ElementsMatch(t, test.expected, test.input.ToSlice())
	}
}

func TestImmutableString(t *testing.T) {

	assert.Equal(t, "{}", Of(lessThanInt).String())
	assert.Equal(t, "{1, 2, 3}", Of(lessThanInt, 1, 2, 3).String())
}
