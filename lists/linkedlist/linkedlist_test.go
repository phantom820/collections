package linkedlist

import (
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types/optional"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	list := New[string]()
	assert.NotNil(t, list)
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
	assert.True(t, list.Empty())
	assert.Equal(t, 0, list.Len())

}
func TestAdd(t *testing.T) {

	type addTest struct {
		input    *LinkedList[int]
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
			elements: []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, test := range addTests {
		for _, element := range test.elements {
			test.input.Add(element)
		}
		assert.Equal(t, test.expected, test.input.ToSlice())
	}
}

func TestAddSlice(t *testing.T) {

	type addSliceTest struct {
		input    *LinkedList[int]
		elements []int
		expected []int
	}

	addSliceTests := []addSliceTest{
		{
			input:    New[int](),
			elements: []int{1},
			expected: []int{1},
		},
		{
			input:    New[int](),
			elements: []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, test := range addSliceTests {
		test.input.AddSlice(test.elements)
		assert.Equal(t, test.expected, test.input.ToSlice())
	}

}

func TestAddAt(t *testing.T) {

	type addAtTest struct {
		input    *LinkedList[int]
		index    int
		value    int
		expected []int
	}

	addAtTests := []addAtTest{
		{
			input:    New(1),
			index:    0,
			value:    -1,
			expected: []int{-1, 1},
		},
		{
			input:    New(1, 2, 3),
			index:    1,
			value:    -2,
			expected: []int{1, -2, 2, 3},
		},
		{
			input:    New(1, 2, 3),
			index:    2,
			value:    4,
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, test := range addAtTests {
		test.input.AddAt(test.index, test.value)
		assert.Equal(t, test.expected, test.input.ToSlice())
	}
}

func TestRemove(t *testing.T) {

	type output struct {
		elements []int
		len      int
		modified bool
	}

	type removeTest struct {
		input    *LinkedList[int]
		element  int
		expected output
	}

	removeTests := []removeTest{
		{
			input:   New[int](),
			element: 1,
			expected: output{
				elements: []int{},
				len:      0,
				modified: false,
			},
		},
		{
			input:   New(1, 2, 3, 4),
			element: 1,
			expected: output{
				elements: []int{2, 3, 4},
				len:      3,
				modified: true,
			},
		},
		{
			input:   New(1, 2, 3, 4),
			element: 2,
			expected: output{
				elements: []int{1, 3, 4},
				len:      3,
				modified: true,
			},
		},
		{
			input:   New(1, 2, 3, 4),
			element: 4,
			expected: output{
				elements: []int{1, 2, 3},
				len:      3,
				modified: true,
			},
		},
	}

	for _, test := range removeTests {
		assert.Equal(t, test.expected, output{
			modified: test.input.Remove(test.element),
			elements: test.input.ToSlice(),
			len:      test.input.len,
		})
	}

}

func TestRemoveAt(t *testing.T) {

	type output struct {
		elements []int
		len      int
		element  int
	}

	type removeAtTest struct {
		input    *LinkedList[int]
		index    int
		expected output
	}

	removeAtTests := []removeAtTest{
		{
			input: New(1, 2, 3, 4),
			index: 0,
			expected: output{
				elements: []int{2, 3, 4},
				len:      3,
				element:  1,
			},
		},
		{
			input: New(1, 2, 3, 4),
			index: 2,
			expected: output{
				elements: []int{1, 2, 4},
				len:      3,
				element:  3,
			},
		},
		{
			input: New(1, 2, 3, 4),
			index: 3,
			expected: output{
				elements: []int{1, 2, 3},
				len:      3,
				element:  4,
			},
		},
	}

	for _, test := range removeAtTests {
		assert.Equal(t, test.expected, output{
			element:  test.input.RemoveAt(test.index),
			elements: test.input.ToSlice(),
			len:      test.input.len,
		})
	}

}

func TestRemoveIf(t *testing.T) {

	type output struct {
		elements []int
		len      int
		modified bool
	}

	type removeIfTest struct {
		input     *LinkedList[int]
		predicate func(int) bool
		expected  output
	}

	removeIfTests := []removeIfTest{
		{
			input:     New[int](),
			predicate: func(i int) bool { return true },
			expected: output{
				elements: []int{},
				len:      0,
				modified: false,
			},
		},
		{
			input:     New(1, 2, 3, 4),
			predicate: func(i int) bool { return i%2 != 0 },
			expected: output{
				elements: []int{2, 4},
				len:      2,
				modified: true,
			},
		},
		{
			input:     New(1, 2, 3, 4),
			predicate: func(i int) bool { return i != 0 },
			expected: output{
				elements: []int{},
				len:      0,
				modified: true,
			},
		},
		{
			input:     New(1, 2, 3, 4),
			predicate: func(i int) bool { return i == 4 },
			expected: output{
				elements: []int{1, 2, 3},
				len:      3,
				modified: true,
			},
		},
	}

	for _, test := range removeIfTests {
		assert.Equal(t, test.expected,
			output{
				modified: test.input.RemoveIf(test.predicate),
				elements: test.input.ToSlice(),
				len:      test.input.len,
			},
		)
	}
}

func TestRemoveSlice(t *testing.T) {

	type output struct {
		elements []int
		len      int
		modified bool
	}

	type removeSliceTest struct {
		input    *LinkedList[int]
		slice    []int
		expected output
	}

	removeSliceTests := []removeSliceTest{
		{
			input: New[int](),
			slice: []int{1},
			expected: output{
				elements: []int{},
				len:      0,
				modified: false,
			},
		},
		{
			input: New(1, 2, 3, 4),
			slice: []int{1},
			expected: output{
				elements: []int{2, 3, 4},
				len:      3,
				modified: true,
			},
		},
		{
			input: New(1, 2, 3, 4),
			slice: []int{1, 2},
			expected: output{
				elements: []int{3, 4},
				len:      2,
				modified: true,
			},
		},
		{
			input: New(1, 2, 3, 4),
			slice: []int{1, 2, 3, 4},
			expected: output{
				elements: []int{},
				len:      0,
				modified: true,
			},
		},
	}

	for _, test := range removeSliceTests {
		assert.Equal(t, test.expected, output{
			modified: test.input.RemoveSlice(test.slice),
			elements: test.input.ToSlice(),
			len:      test.input.len,
		})
	}
}

func TestClear(t *testing.T) {

	list := New(1, 2, 3, 4, 5)
	list.Clear()

	assert.NotNil(t, list)
	assert.True(t, list.Empty())
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)

}

func TestIndexOf(t *testing.T) {

	type indexOfTest struct {
		input    *LinkedList[int]
		expected optional.Optional[int]
	}

	indexOfTests := []indexOfTest{
		{
			input:    New[int](),
			expected: optional.Empty[int](),
		},
		{
			input:    New(1, 2, 3, 4),
			expected: optional.Of(0),
		},
		{
			input:    New(0, 2, 1, 4),
			expected: optional.Of(2),
		},
		{
			input:    New(0, 1, 1, 4),
			expected: optional.Of(1),
		},
	}

	for _, test := range indexOfTests {
		assert.Equal(t, test.expected, test.input.IndexOf(1))
	}
}

func TestSet(t *testing.T) {

	type setTest struct {
		input    *LinkedList[int]
		index    int
		value    int
		expected []int
	}

	setTests := []setTest{
		{
			input:    New(1),
			index:    0,
			value:    -1,
			expected: []int{-1},
		},
		{
			input:    New(1, 2, 3),
			index:    1,
			value:    -2,
			expected: []int{1, -2, 3},
		},
		{
			input:    New(1, 2, 3),
			index:    2,
			value:    4,
			expected: []int{1, 2, 4},
		},
	}

	for _, test := range setTests {
		test.input.Set(test.index, test.value)
		assert.Equal(t, test.expected, test.input.ToSlice())
	}
}

func TestContains(t *testing.T) {

	type containsTest struct {
		input    *LinkedList[int]
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

func TestAt(t *testing.T) {

	type atTest struct {
		input    *LinkedList[int]
		index    int
		expected int
	}

	atTests := []atTest{
		{
			input:    New(1, 2, 3, 4),
			index:    0,
			expected: 1,
		},
		{
			input:    New(1, 2, 3, 4),
			index:    3,
			expected: 4,
		},
		{
			input:    New(1, 2, 3, 4),
			index:    1,
			expected: 2,
		},
	}

	for _, test := range atTests {
		assert.Equal(t, test.expected, test.input.At(test.index))
	}
}

func TestForEach(t *testing.T) {

	list := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0

	list.ForEach(func(i int) { sum = sum + i })

	assert.Equal(t, 55, sum)
}

func TestRetainAll(t *testing.T) {

	type retainAllTest struct {
		a        *LinkedList[int]
		b        *LinkedList[int]
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
			a:        New[int](),
			b:        New(9, 1, 2, 3, 4, 5),
			expected: []int{},
		},
	}

	for _, test := range retainAllTests {
		test.a.RetainAll(test.b)
		assert.Equal(t, test.expected, test.a.ToSlice())
	}

}

func TestAddAll(t *testing.T) {

	type addAllTest struct {
		a        *LinkedList[int]
		b        *LinkedList[int]
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
		assert.Equal(t, test.expected, test.a.ToSlice())
	}

}

func TestRemoveAll(t *testing.T) {

	type removeAllTest struct {
		a        *LinkedList[int]
		b        *LinkedList[int]
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
		assert.Equal(t, test.expected, test.a.ToSlice())
	}

}

func TestEquals(t *testing.T) {

	type equalsTest struct {
		a        *LinkedList[int]
		b        *LinkedList[int]
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
			b:        New(1, 2),
			expected: true,
		},
		{
			a:        New(1, 2, 3),
			b:        New(10, 12, 14),
			expected: false,
		},
	}

	for _, test := range equalsTests {
		assert.True(t, test.a.Equals(test.a))
		assert.Equal(t, test.expected, test.a.Equals(test.b))
		assert.Equal(t, test.expected, test.b.Equals(test.a))

	}

}

func TestSubList(t *testing.T) {

	type subListTest struct {
		input      *LinkedList[int]
		start, end int
		expected   []int
	}

	subListTests := []subListTest{
		{
			input:    New(1),
			start:    0,
			end:      0,
			expected: []int{},
		},
		{
			input:    New(1, 2),
			start:    0,
			end:      1,
			expected: []int{1},
		},
		{
			input:    New(1, 2, 3, 4, 5),
			start:    0,
			end:      4,
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    New(1, 2, 3, 4, 5),
			start:    1,
			end:      4,
			expected: []int{2, 3, 4},
		},
		{
			input:    New(1, 2, 3, 4, 5),
			start:    0,
			end:      5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    New(1, 2, 3, 4, 5),
			start:    2,
			end:      5,
			expected: []int{3, 4, 5},
		},
	}

	for _, test := range subListTests {
		assert.Equal(t, test.expected, test.input.SubList(test.start, test.end).ToSlice())
	}
}

func TestIterator(t *testing.T) {

	type iteratorTest struct {
		input    *LinkedList[int]
		expected []int
	}

	iteratorTests := []iteratorTest{
		{
			input:    New[int](),
			expected: []int{},
		},
		{
			input:    New(1, 2, 3, 4),
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    New(1),
			expected: []int{1},
		},
	}

	iterate := func(it iterator.Iterator[int]) []int {
		data := make([]int, 0)
		for it.HasNext() {
			data = append(data, it.Next())
		}
		return data
	}

	for _, test := range iteratorTests {
		assert.Equal(t, test.expected, iterate(test.input.Iterator()))
	}
}

func TestString(t *testing.T) {

	assert.Equal(t, "[]", New[int]().String())
	assert.Equal(t, "[1]", New(1).String())
	assert.Equal(t, "[1 2]", New(1, 2).String())
}

func TestSort(t *testing.T) {

	type sortTest struct {
		input    *LinkedList[int]
		less     func(int, int) bool
		expected []int
	}

	sortTests := []sortTest{
		{
			input:    New[int](),
			less:     func(i1, i2 int) bool { return i1 < i2 },
			expected: []int{},
		},
		{
			input:    New(2, 1, 4),
			less:     func(i1, i2 int) bool { return i1 < i2 },
			expected: []int{1, 2, 4},
		},
		{
			input:    New(1, 2, 3, 5, 4),
			less:     func(i1, i2 int) bool { return i1 < i2 },
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    New(5, 4, 3, 2, 1),
			less:     func(i1, i2 int) bool { return i1 <= i2 },
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    New(1, 2, 3, 4, 5),
			less:     func(i1, i2 int) bool { return i1 >= i2 },
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, test := range sortTests {
		test.input.Sort(test.less)
		assert.Equal(t, test.expected, test.input.ToSlice())
	}

}

func TestCopy(t *testing.T) {

	type copyTest struct {
		input    *LinkedList[int]
		expected []int
	}

	copyTests := []copyTest{
		{
			input:    New[int](),
			expected: []int{},
		},
		{
			input:    New(1, 2, 3, 4),
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, test := range copyTests {
		assert.Equal(t, test.expected, test.input.Copy().ToSlice())
		assert.Equal(t, test.expected, test.input.ImmutableCopy().ToSlice())
	}
}
