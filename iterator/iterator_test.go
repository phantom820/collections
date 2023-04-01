package iterator

import (
	"testing"

	"github.com/phantom820/collections/types/optional"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {

	type mapTest struct {
		input    Iterator[int]
		expected []int
	}

	mapTests := []mapTest{
		{
			input:    Of[int](),
			expected: []int{},
		},
		{
			input:    Of(1, 2, 3, 4, 5),
			expected: []int{2, 4, 6, 8, 10},
		},
	}

	for _, test := range mapTests {

		assert.Equal(t, test.expected, ToSlice(Map(test.input, func(i int) int { return i * 2 })))
	}

}

func TestFilter(t *testing.T) {

	type filterTest struct {
		input    Iterator[int]
		expected []int
	}

	filterTests := []filterTest{
		{
			input:    Of[int](),
			expected: []int{},
		},
		{
			input:    Of(1, 2, 3, 4, 5),
			expected: []int{2, 4},
		},
	}

	for _, test := range filterTests {

		assert.Equal(t, test.expected, ToSlice(Filter(test.input, func(i int) bool { return i%2 == 0 })))
	}

}

func TestReduce(t *testing.T) {

	type reduceTest struct {
		input    Iterator[int]
		expected optional.Optional[int]
	}

	reduceTests := []reduceTest{
		{
			input:    Of[int](),
			expected: optional.Empty[int](),
		},
		{
			input:    Of(1),
			expected: optional.Of(1),
		},
		{
			input:    Of(1, 2),
			expected: optional.Of(3),
		},

		{
			input:    Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			expected: optional.Of(55),
		},
	}

	for _, test := range reduceTests {
		assert.Equal(t, test.expected, Reduce(test.input, func(i1, i2 int) int { return i1 + i2 }))
	}

}
