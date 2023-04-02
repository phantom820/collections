package hashset

import (
	"testing"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/queues/vectordequeue"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	set := New[string]()
	assert.NotNil(t, set)
	assert.True(t, set.Empty())
	assert.Equal(t, 0, set.Len())

}

func TestAdd(t *testing.T) {

	type addTest struct {
		input    *HashSet[int]
		elements []int
		expected []int
	}

	addTests := []addTest{
		{
			input:    New[int](),
			elements: []int{1},
			expected: []int{1},
		},
		{
			input:    New[int](),
			elements: []int{1, 2, 3, 1, 2},
			expected: []int{1, 2, 3},
		},
	}

	for _, test := range addTests {
		for _, element := range test.elements {
			test.input.Add(element)
		}
		assert.ElementsMatch(t, test.expected, test.input.ToSlice())
	}
}

func TestAddSlice(t *testing.T) {

	type addSliceTest struct {
		input    *HashSet[int]
		slice    []int
		expected []int
	}

	addSliceTests := []addSliceTest{
		{
			input:    New[int](),
			slice:    []int{1},
			expected: []int{1},
		},
		{
			input:    New[int](),
			slice:    []int{1, 2, 3, 1, 2},
			expected: []int{1, 2, 3},
		},
	}

	for _, test := range addSliceTests {
		test.input.AddSlice(test.slice)
		assert.ElementsMatch(t, test.expected, test.input.ToSlice())
	}

}

func TestRemove(t *testing.T) {

	type output struct {
		elements []int
		len      int
		modified bool
	}

	type removeTest struct {
		input    *HashSet[int]
		element  int
		expected output
	}

	removeTests := []removeTest{
		{
			input:   New[int](),
			element: 0,
			expected: output{
				elements: []int{},
				len:      0,
				modified: false,
			},
		},
		{
			input:   New(1, 2, 3),
			element: 1,
			expected: output{
				elements: []int{3, 2},
				len:      2,
				modified: true,
			},
		},
		{
			input:   New(1, 2, 3),
			element: 3,
			expected: output{
				elements: []int{1, 2},
				len:      2,
				modified: true,
			},
		},
		{
			input:   New(3),
			element: 3,
			expected: output{
				elements: []int{},
				len:      0,
				modified: true,
			},
		},
	}

	for _, test := range removeTests {
		output := output{
			modified: test.input.Remove(test.element),
			elements: test.input.ToSlice(),
			len:      test.input.Len(),
		}
		assert.ElementsMatch(t, test.expected.elements, output.elements)
		assert.Equal(t, test.expected.modified, output.modified)
		assert.Equal(t, test.expected.len, output.len)

	}

}

func TestRemoveIf(t *testing.T) {

	type output struct {
		elements []int
		len      int
		modified bool
	}

	type removeIfTest struct {
		input     *HashSet[int]
		predicate func(int) bool
		expected  output
	}

	removeIfTests := []removeIfTest{
		{
			input:     New[int](),
			predicate: func(i int) bool { return i%2 == 0 },
			expected: output{
				elements: []int{},
				len:      0,
				modified: false,
			},
		},
		{
			input:     New(1, 2, 3),
			predicate: func(i int) bool { return i%2 != 0 },
			expected: output{
				elements: []int{2},
				len:      1,
				modified: true,
			},
		},
		{
			input:     New(1, 2, 3),
			predicate: func(i int) bool { return false },
			expected: output{
				elements: []int{1, 2, 3},
				len:      3,
				modified: false,
			},
		},
		{
			input:     New(3),
			predicate: func(i int) bool { return i == 3 },
			expected: output{
				elements: []int{},
				len:      0,
				modified: true,
			},
		},
	}

	for _, test := range removeIfTests {
		output := output{
			modified: test.input.RemoveIf(test.predicate),
			elements: test.input.ToSlice(),
			len:      test.input.Len(),
		}
		assert.ElementsMatch(t, test.expected.elements, output.elements)
		assert.Equal(t, test.expected.modified, output.modified)
		assert.Equal(t, test.expected.len, output.len)

	}
}

func TestRemoveSlice(t *testing.T) {

	type output struct {
		elements []int
		len      int
		modified bool
	}

	type removeSliceTest struct {
		input    *HashSet[int]
		slice    []int
		expected output
	}

	removeSliceTests := []removeSliceTest{
		{
			input: New[int](),
			slice: []int{},
			expected: output{
				elements: []int{},
				len:      0,
				modified: false,
			},
		},
		{
			input: New(2),
			slice: []int{3},
			expected: output{
				elements: []int{2},
				len:      1,
				modified: false,
			},
		},
		{
			input: New(1, 2, 3, 4, 5),
			slice: []int{2, 3, 1},
			expected: output{
				elements: []int{5, 4},
				len:      2,
				modified: true,
			},
		},
		{
			input: New(1, 2, 3, 4, 5),
			slice: []int{2, 3, 1, 4, 5},
			expected: output{
				elements: []int{},
				len:      0,
				modified: true,
			},
		},
	}

	for _, test := range removeSliceTests {
		output := output{
			modified: test.input.RemoveSlice(test.slice),
			elements: test.input.ToSlice(),
			len:      test.input.Len(),
		}
		test.input.RemoveSlice(test.slice)
		assert.ElementsMatch(t, test.expected.elements, output.elements)
		assert.Equal(t, test.expected.modified, output.modified)
		assert.Equal(t, test.expected.len, output.len)
	}
}

func TestClear(t *testing.T) {

	set := New(1, 2, 3, 4, 5)
	set.Clear()

	assert.NotNil(t, set)
	assert.True(t, set.Empty())

}

func TestContains(t *testing.T) {

	type containsTest struct {
		input    *HashSet[int]
		element  int
		expected bool
	}

	containsTests := []containsTest{
		{
			input:    New(0, 4, 5),
			element:  1,
			expected: false,
		},
		{
			input:    New(0, 4, 5),
			element:  2,
			expected: false,
		},
		{
			input:    New(0, 4, 5),
			element:  4,
			expected: true,
		},
	}

	for _, test := range containsTests {
		assert.Equal(t, test.expected, test.input.Contains(test.element))
	}
}

func TestContainsAll(t *testing.T) {

	type containsAllTest struct {
		input    *HashSet[int]
		elements []int
		expected bool
	}

	containsAllTests := []containsAllTest{
		{
			input:    New(0, 4, 5),
			elements: []int{},
			expected: true,
		},
		{
			input:    New(0, 4, 5),
			elements: []int{1},
			expected: false,
		},
		{
			input:    New(0, 4, 5),
			elements: []int{4, 5},
			expected: true,
		},
		{
			input:    New(0, 4, 5),
			elements: []int{0, 4, 5},
			expected: true,
		},
		{
			input:    New(0, 4, 5),
			elements: []int{3},
			expected: false,
		},
		{
			input:    New(0, 4, 5),
			elements: []int{0, 4, 5, 8},
			expected: false,
		},
	}

	for _, test := range containsAllTests {
		iterable := New(test.elements...)
		assert.Equal(t, test.expected, test.input.ContainsAll(iterable))
	}
}

func TestAddAll(t *testing.T) {

	type addAllTest struct {
		a        *HashSet[int]
		b        *HashSet[int]
		expected []int
	}

	addAllTests := []addAllTest{
		{
			a:        New[int](),
			b:        New(1, 2, 3, 4, 5),
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			a:        New(1, 2),
			b:        New(9, 11, 12),
			expected: []int{1, 2, 9, 11, 12},
		},
	}

	for _, test := range addAllTests {
		test.a.AddAll(test.b)
		assert.ElementsMatch(t, test.expected, test.a.ToSlice())
	}

}

func TestRemoveAll(t *testing.T) {

	type removeAllTest struct {
		a        *HashSet[int]
		b        *HashSet[int]
		expected []int
	}

	removeAllTests := []removeAllTest{
		{
			a:        New[int](),
			b:        New[int](),
			expected: []int{},
		},
		{
			a:        New(1, 2, 3, 4, 5),
			b:        New[int](),
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			a:        New(1, 2, 3, 4, 5),
			b:        New(9, 1, 2),
			expected: []int{3, 4, 5},
		},
		{
			a:        New(1, 2, 3, 4, 5),
			b:        New(9, 1, 2, 3, 4, 5),
			expected: []int{},
		},
	}

	for _, test := range removeAllTests {
		test.a.RemoveAll(test.b)
		assert.ElementsMatch(t, test.expected, test.a.ToSlice())
	}

}

func TestRetainAll(t *testing.T) {

	type retainAllTest struct {
		a        *HashSet[int]
		b        collections.Collection[int]
		expected []int
	}

	retainAllTests := []retainAllTest{
		{
			a:        New(1, 2, 3, 4, 5),
			b:        New[int](),
			expected: []int{},
		},
		{
			a:        New(1, 2, 3, 4, 5),
			b:        New(9, 1, 2),
			expected: []int{1, 2},
		},
		{
			a:        New(1, 2, 3, 4, 5),
			b:        New(9, 1, 2, 3, 4, 5),
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			a:        New(1, 2, 3, 4, 5),
			b:        vectordequeue.New(9, 1, 2, 3, 4, 5),
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range retainAllTests {
		test.a.RetainAll(test.b)
		assert.ElementsMatch(t, test.expected, test.a.ToSlice())
	}

}

func TestForEach(t *testing.T) {

	set := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0

	set.ForEach(func(i int) { sum = sum + i })

	assert.Equal(t, 55, sum)
}

func TestEquals(t *testing.T) {

	type equalsTest struct {
		a        *HashSet[int]
		b        *HashSet[int]
		expected bool
	}

	equalsTests := []equalsTest{
		{
			a:        New[int](),
			b:        New[int](),
			expected: true,
		},
		{
			a:        New(1, 2),
			b:        New[int](),
			expected: false,
		},
		{
			a:        New(1, 2),
			b:        New(2, 1),
			expected: true,
		},
		{
			a:        New(1, 2, 3),
			b:        New(10, 12, 14),
			expected: false,
		},
	}

	for _, test := range equalsTests {
		assert.Equal(t, test.expected, test.a.Equals(test.b))
		assert.Equal(t, test.expected, test.b.Equals(test.a))

	}

	identity := New[int]()
	assert.True(t, identity.Equals(identity))

}

func TestToSlice(t *testing.T) {

	type toSliceTest struct {
		input    *HashSet[int]
		expected []int
	}

	toSliceTests := []toSliceTest{
		{
			input:    New[int](),
			expected: []int{},
		},
		{
			input:    New(1, 2, 3, 4),
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    New(1, 2, 3, 4, 5),
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range toSliceTests {
		assert.ElementsMatch(t, test.expected, test.input.ToSlice())
	}
}

func TestString(t *testing.T) {

	assert.Equal(t, "{}", New[int]().String())
	assert.Equal(t, "{1}", New(1).String())
}
