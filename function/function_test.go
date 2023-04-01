package function

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/sets/hashset"
	"github.com/phantom820/collections/sets/linkedhashset"
	"github.com/phantom820/collections/types/optional"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {

	type mapTest struct {
		input    iterable.Iterable[int]
		expected []string
	}

	mapTests := []mapTest{
		{
			input:    iterable.Of[int](),
			expected: []string{},
		},
		{
			input:    iterable.Of(1),
			expected: []string{"2"},
		},
		{
			input:    iterable.Of(1, 2, 3, 4, 5),
			expected: []string{"2", "4", "6", "8", "10"},
		},
	}

	for _, test := range mapTests {
		assert.Equal(t, test.expected, (Map(test.input, func(i int) string { return fmt.Sprint(i * 2) })).ToSlice())
	}

}

func TestMapMethod(t *testing.T) {

	view := Filter(iterable.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), func(i int) bool { return i%2 == 0 }).
		Map(func(i int) int { return i * 2 })

	sum := view.Reduce(func(x, y int) int {
		return x + y
	})

	assert.Equal(t, 60, sum.Value())
}

func TestToSlice(t *testing.T) {

	assert.Equal(t, []int{}, Identity(iterable.Of[int]()).ToSlice())
	assert.Equal(t, []int{1, 2, 3}, Identity(iterable.Of(1, 2, 3)).ToSlice())

}

func TestToHashSet(t *testing.T) {

	assert.True(t, Identity(iterable.Of[int]()).ToHashSet().Equals(hashset.New[int]()))
	assert.True(t, Identity(iterable.Of(1, 2, 3)).ToHashSet().Equals(hashset.New(1, 2, 3)))

}

func TestToLinkedHashSet(t *testing.T) {

	assert.True(t, Identity(iterable.Of[int]()).ToLinkedHashSet().Equals(linkedhashset.New[int]()))
	assert.True(t, Identity(iterable.Of(1, 2, 3)).ToLinkedHashSet().Equals(linkedhashset.New(1, 2, 3)))

}

func TestToTreeSet(t *testing.T) {

	lessThan := func(i1, i2 int) bool { return i1 > i2 }
	assert.True(t, Identity(iterable.Of[int]()).ToTreeSet(lessThan).Equals(linkedhashset.New[int]()))
	assert.True(t, Identity(iterable.Of(1, 2, 3)).ToTreeSet(lessThan).Equals(linkedhashset.New(1, 2, 3)))

}

func TestToLinkedList(t *testing.T) {

	assert.True(t, Identity(iterable.Of[int]()).ToLinkedList().Equals(linkedlist.New[int]()))
	assert.True(t, Identity(iterable.Of(1, 2, 3)).ToLinkedList().Equals(linkedlist.New(1, 2, 3)))

}

func TestToForwardList(t *testing.T) {

	assert.True(t, Identity(iterable.Of[int]()).ToForwardList().Equals(forwardlist.New[int]()))
	assert.True(t, Identity(iterable.Of(1, 2, 3)).ToForwardList().Equals(forwardlist.New(1, 2, 3)))

}

func TestToVector(t *testing.T) {

	assert.True(t, Identity(iterable.Of[int]()).ToVector().Equals(vector.New[int]()))
	assert.True(t, Identity(iterable.Of(1, 2, 3)).ToVector().Equals(vector.New(1, 2, 3)))

}

func TestFilter(t *testing.T) {

	type filterTest struct {
		input    iterable.Iterable[int]
		expected []int
	}

	filterTests := []filterTest{
		{
			input:    iterable.Of[int](),
			expected: []int{},
		},
		{
			input:    iterable.Of(1, 3, 5),
			expected: []int{},
		},
		{
			input:    iterable.Of(1, 2, 3, 4, 5),
			expected: []int{2, 4},
		},
	}

	for _, test := range filterTests {
		assert.Equal(t, test.expected, (Filter(test.input, func(i int) bool { return i%2 == 0 })).ToSlice())
	}

}

func TestFilterMethod(t *testing.T) {

	slice := Filter(iterable.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
		func(i int) bool { return i%2 == 0 }).
		Filter(
			func(i int) bool { return i == 2 },
		).ToSlice()

	assert.Equal(t, []int{2}, slice)
}

func TestReduce(t *testing.T) {

	type reduceTest struct {
		input    iterable.Iterable[int]
		expected optional.Optional[int]
	}

	reduceTests := []reduceTest{
		{
			input:    iterable.Of[int](),
			expected: optional.Empty[int](),
		},
		{
			input:    iterable.Of(1),
			expected: optional.Of(1),
		},
		{
			input:    iterable.Of(1, 2),
			expected: optional.Of(3),
		},

		{
			input:    iterable.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			expected: optional.Of(55),
		},
	}

	for _, test := range reduceTests {
		assert.Equal(t, test.expected, Reduce(test.input, func(i1, i2 int) int { return i1 + i2 }))
	}

}

func TestReduceMethod(t *testing.T) {

	view := Filter(iterable.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), func(i int) bool { return i%2 == 0 })
	sum := view.Reduce(func(x, y int) int {
		return x + y
	})

	assert.Equal(t, 30, sum.Value())
}

func TestForEach(t *testing.T) {

	view := Filter(iterable.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), func(i int) bool { return i%2 == 0 })

	sum := 0
	view.ForEach(func(i int) {
		sum = sum + i
	})

	assert.Equal(t, 30, sum)
}

func TestGroupBy(t *testing.T) {

	type groupByTest struct {
		input    iterable.Iterable[int]
		expected hashmap.HashMap[string, []int]
	}

	groupByTests := []groupByTest{
		{
			input:    iterable.Of[int](),
			expected: map[string][]int{},
		},
		{
			input:    iterable.Of(1),
			expected: map[string][]int{"odd": {1}},
		},
		{
			input:    iterable.Of(1, 2),
			expected: map[string][]int{"odd": {1}, "even": {2}},
		},
		{
			input:    iterable.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			expected: map[string][]int{"odd": {1, 3, 5, 7, 9}, "even": {2, 4, 6, 8, 10}},
		},
	}

	discriminator := func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	}

	for _, test := range groupByTests {
		assert.Equal(t, test.expected, GroupBy(test.input, discriminator))
	}

}
