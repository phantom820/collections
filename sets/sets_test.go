package sets

import (
	"testing"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/sets/hashset"
	"github.com/phantom820/collections/sets/linkedhashset"
	"github.com/phantom820/collections/sets/treeset"
	"github.com/stretchr/testify/assert"
)

func TestIsSet(t *testing.T) {

	type isSetTest struct {
		input    collections.Collection[int]
		expected bool
	}

	a := hashset.Of[int]()
	b := linkedhashset.Of[int]()
	c := treeset.Of(func(e1, e2 int) bool { return e1 < e2 })
	isSetTests := []isSetTest{
		{
			input:    nil,
			expected: false,
		},
		{
			input:    hashset.New[int](),
			expected: true,
		},
		{
			input:    &a,
			expected: true,
		},
		{
			input:    a,
			expected: true,
		},
		{
			input:    linkedhashset.New[int](),
			expected: true,
		},
		{
			input:    &b,
			expected: true,
		},
		{
			input:    &b,
			expected: true,
		},
		{
			input:    treeset.New(func(e1, e2 int) bool { return e1 < e2 }),
			expected: true,
		},
		{
			input:    &c,
			expected: true,
		},
		{
			input:    &c,
			expected: true,
		},
	}

	for _, test := range isSetTests {
		assert.Equal(t, test.expected, IsSet[int](test.input))
	}
}

func TestUnion(t *testing.T) {

	type unionTest struct {
		inputs           func() (collections.Set[int], collections.Set[int])
		element          int
		expectedSlice    []int
		expectedLen      int
		expectedContains bool
	}

	unionTests := []unionTest{
		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of[int](), linkedhashset.Of[int]()
				return &a, &b
			},
			element:          0,
			expectedSlice:    []int{},
			expectedLen:      0,
			expectedContains: false,
		},
		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of(1, 2, 3), linkedhashset.Of(4, 5, 6)
				return &a, &b
			},
			element:          4,
			expectedSlice:    []int{1, 2, 3, 4, 5, 6},
			expectedLen:      6,
			expectedContains: true,
		},
		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of(1, 2, 3), treeset.Of(func(e1, e2 int) bool { return e1 > e2 }, 4, 5, 6)
				return &a, &b
			},
			element:          4,
			expectedSlice:    []int{1, 2, 3, 4, 5, 6},
			expectedLen:      6,
			expectedContains: true,
		},
		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of(1, 2, 3, 4), linkedhashset.Of(4, 5, 6)
				return &a, &b
			},
			element:          10,
			expectedSlice:    []int{1, 2, 3, 4, 5, 6},
			expectedLen:      6,
			expectedContains: false,
		},
	}

	for _, test := range unionTests {
		a, b := test.inputs()
		c := Union(a, b)
		assert.ElementsMatch(t, test.expectedSlice, c.ToSlice())
		assert.Equal(t, test.expectedLen, c.Len())
		assert.Equal(t, test.expectedContains, c.Contains(test.element))
	}
}

func TestDifference(t *testing.T) {

	type differenceTest struct {
		inputs           func() (collections.Set[int], collections.Set[int])
		element          int
		expectedSlice    []int
		expectedLen      int
		expectedContains bool
	}

	differenceTests := []differenceTest{
		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of[int](), linkedhashset.Of[int]()
				return &a, &b
			},
			element:          0,
			expectedSlice:    []int{},
			expectedLen:      0,
			expectedContains: false,
		},

		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of(1, 2, 3), linkedhashset.Of(4, 5, 6)
				return &a, &b
			},
			element:          2,
			expectedSlice:    []int{1, 2, 3},
			expectedLen:      3,
			expectedContains: true,
		},
		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of(1, 2, 3, 4), linkedhashset.Of(1, 2, 3)
				return &a, &b
			},
			element:          3,
			expectedSlice:    []int{4},
			expectedLen:      1,
			expectedContains: false,
		},
		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of(1, 2, 3), linkedhashset.Of(1, 2, 3)
				return &a, &b
			},
			element:          3,
			expectedSlice:    []int{},
			expectedLen:      0,
			expectedContains: false,
		},
	}

	for _, test := range differenceTests {
		a, b := test.inputs()
		c := Difference(a, b)
		assert.ElementsMatch(t, test.expectedSlice, c.ToSlice())
		assert.Equal(t, test.expectedLen, c.Len())
		assert.Equal(t, test.expectedContains, c.Contains(test.element))
	}
}

func TestIntersection(t *testing.T) {

	type intersectionTest struct {
		inputs           func() (collections.Set[int], collections.Set[int])
		element          int
		expectedSlice    []int
		expectedLen      int
		expectedContains bool
	}

	intersectionTests := []intersectionTest{
		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of[int](), linkedhashset.Of[int]()
				return &a, &b
			},
			element:          0,
			expectedSlice:    []int{},
			expectedLen:      0,
			expectedContains: false,
		},

		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of(1, 2, 3), linkedhashset.Of(4, 5, 6)
				return &a, &b
			},
			element:          2,
			expectedSlice:    []int{},
			expectedLen:      0,
			expectedContains: false,
		},
		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of(1, 2, 3, 4), linkedhashset.Of(1, 2, 5)
				return &a, &b
			},
			element:          2,
			expectedSlice:    []int{1, 2},
			expectedLen:      2,
			expectedContains: true,
		},
		{
			inputs: func() (collections.Set[int], collections.Set[int]) {
				a, b := hashset.Of(1, 2, 3), linkedhashset.Of(8, 7, 6, 9, 11, 1, 3)
				return &a, &b
			},
			element:          3,
			expectedSlice:    []int{1, 3},
			expectedLen:      2,
			expectedContains: true,
		},
	}

	for _, test := range intersectionTests {
		a, b := test.inputs()
		c := Intersection(a, b)
		assert.ElementsMatch(t, test.expectedSlice, c.ToSlice())
		assert.Equal(t, test.expectedLen, c.Len())
		assert.Equal(t, test.expectedContains, c.Contains(test.element))
	}
}

func TestToHashSet(t *testing.T) {

	a := hashset.Of(1, 2, 3, 4)
	b := linkedhashset.Of(5, 6, 7, 8)
	c := Union[int](&a, &b).ToHashSet()

	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, c.ToSlice())

}

func TestToLinkedHashSet(t *testing.T) {

	a := hashset.Of(1, 2, 3, 4)
	b := linkedhashset.Of(5, 6, 7, 8)
	c := Union[int](&a, &b).ToLinkedHashSet()

	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, c.ToSlice())

}

func TestToTreeSet(t *testing.T) {

	a := hashset.Of(1, 2, 3, 4)
	b := linkedhashset.Of(5, 6, 7, 8)
	c := Union[int](&a, &b).ToTreeSet(func(e1, e2 int) bool { return e1 >= e2 })

	assert.Equal(t, []int{8, 7, 6, 5, 4, 3, 2, 1}, c.ToSlice())

}
