package forwadlist

import (
	"fmt"
	"testing"

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
		input    []string
		expected []string
	}

	addTests := []addTest{
		{
			input:    []string{},
			expected: []string{},
		},
		{
			input:    []string{"A", "A", "B"},
			expected: []string{"A", "A", "B"},
		},
		{
			input:    []string{"1", "2", "3"},
			expected: []string{"1", "2", "3"},
		},
	}

	f := func(values []string) []string {
		list := New[string]()
		for _, value := range values {
			list.Add(value)
		}
		return list.ToSlice()
	}

	for _, test := range addTests {
		assert.Equal(t, test.expected, f(test.input))
	}
}

func TestAddSlice(t *testing.T) {

	type addSliceTest struct {
		input    []string
		expected ForwardList[string]
	}

	addSliceTests := []addSliceTest{
		{
			input:    []string{},
			expected: Of[string](),
		},
		{
			input:    []string{"A", "A", "B"},
			expected: Of("A", "A", "B"),
		},
	}

	for _, test := range addSliceTests {
		list := New[string]()
		list.AddSlice(test.input)
		assert.Equal(t, test.expected, *list)
	}

}

func TestAddAt(t *testing.T) {

	type addAtTest struct {
		input    ForwardList[int]
		index    int
		value    int
		expected []int
	}

	addAtTests := []addAtTest{
		{
			input:    Of(1),
			index:    0,
			value:    -1,
			expected: []int{-1, 1},
		},
		{
			input:    Of(1, 2, 3),
			index:    1,
			value:    -2,
			expected: []int{1, -2, 2, 3},
		},
		{
			input:    Of(1, 2, 3),
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

	type removeTest struct {
		input           ForwardList[string]
		element         string
		expectedList    []string
		expectedBoolean bool
		expectedLen     int
	}

	removeTests := []removeTest{
		{
			input:           Of("A", "B", "C", "D"),
			element:         "F",
			expectedList:    []string{"A", "B", "C", "D"},
			expectedBoolean: false,
			expectedLen:     4,
		},
		{
			input:           Of("A", "B", "C", "D"),
			element:         "A",
			expectedList:    []string{"B", "C", "D"},
			expectedBoolean: true,
			expectedLen:     3,
		},
		{
			input:           Of("A", "B", "C", "D"),
			element:         "B",
			expectedList:    []string{"A", "C", "D"},
			expectedBoolean: true,
			expectedLen:     3,
		},
		{
			input:           Of("A", "B", "C", "D"),
			element:         "D",
			expectedList:    []string{"A", "B", "C"},
			expectedBoolean: true,
			expectedLen:     3,
		},
		{
			input:           Of("A"),
			element:         "A",
			expectedList:    []string{},
			expectedBoolean: true,
			expectedLen:     0,
		},
	}

	for _, test := range removeTests {
		assert.Equal(t, test.expectedBoolean, test.input.Remove(test.element))
		assert.Equal(t, test.expectedList, test.input.ToSlice())
		assert.Equal(t, test.expectedLen, test.input.Len())
	}

}

func TestRemoveAt(t *testing.T) {

	type removeAtTest struct {
		input           ForwardList[int]
		index           int
		expectedElement int
		expectedList    []int
	}

	removeAtTests := []removeAtTest{
		{
			input:           Of(1, 2, 3, 4),
			index:           0,
			expectedElement: 1,
			expectedList:    []int{2, 3, 4},
		},
		{
			input:           Of(1, 2, 3, 4),
			index:           3,
			expectedElement: 4,
			expectedList:    []int{1, 2, 3},
		},
		{
			input:           Of(1, 2, 3, 4),
			index:           2,
			expectedElement: 3,
			expectedList:    []int{1, 2, 4},
		},
	}

	for _, test := range removeAtTests {
		assert.Equal(t, test.expectedElement, test.input.RemoveAt(test.index))
		assert.Equal(t, test.expectedList, test.input.ToSlice())
	}

}

func TestRemoveIf(t *testing.T) {

	type removeIfTest struct {
		input           ForwardList[int]
		expectedBoolean bool
		expectedList    []int
	}

	removeIfTests := []removeIfTest{
		{
			input:           Of[int](),
			expectedBoolean: false,
			expectedList:    []int{},
		},
		{
			input:           Of(2),
			expectedBoolean: false,
			expectedList:    []int{2},
		},
		{
			input:           Of(1, 2, 3, 4, 5),
			expectedBoolean: true,
			expectedList:    []int{2, 4},
		},
		{
			input:           Of(1, 3, 5, 7),
			expectedBoolean: true,
			expectedList:    []int{},
		},
	}

	f := func(x int) bool {
		return x%2 != 0
	}

	for _, test := range removeIfTests[2:3] {
		test.input.RemoveIf(f)
		assert.Equal(t, test.expectedList, test.input.ToSlice())
	}
}

func TestRemoveSlice(t *testing.T) {

	type removeSliceTest struct {
		input           ForwardList[int]
		slice           []int
		expectedBoolean bool
		expectedList    []int
	}

	removeSliceTests := []removeSliceTest{
		{
			input:           Of[int](),
			slice:           []int{},
			expectedBoolean: false,
			expectedList:    []int{},
		},
		{
			input:           Of(2),
			slice:           []int{3},
			expectedBoolean: false,
			expectedList:    []int{2},
		},
		{
			input:           Of(1, 2, 3, 4, 5),
			slice:           []int{2, 3, 1, 4},
			expectedBoolean: true,
			expectedList:    []int{5},
		},
		{
			input:           Of(1, 2, 3, 4, 5),
			slice:           []int{2, 3, 1, 4, 5},
			expectedBoolean: true,
			expectedList:    []int{},
		},
	}

	for _, test := range removeSliceTests {
		test.input.RemoveSlice(test.slice)
		assert.Equal(t, test.expectedList, test.input.ToSlice())
	}
}

func TestClear(t *testing.T) {

	list := Of(1, 2, 3, 4, 5)
	list.Clear()

	assert.NotNil(t, list)
	assert.True(t, list.Empty())
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)

}

func TestSet(t *testing.T) {

	type setTest struct {
		input    ForwardList[int]
		index    int
		value    int
		expected []int
	}

	setTests := []setTest{
		{
			input:    Of(1),
			index:    0,
			value:    -1,
			expected: []int{-1},
		},
		{
			input:    Of(1, 2, 3),
			index:    1,
			value:    -2,
			expected: []int{1, -2, 3},
		},
		{
			input:    Of(1, 2, 3),
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
		input    ForwardList[int]
		element  int
		expected bool
	}

	containsTests := []containsTest{
		{
			input:    Of(0, 4, 5),
			element:  1,
			expected: false,
		},
		{
			input:    Of(0, 4, 5),
			element:  2,
			expected: false,
		},
		{
			input:    Of(0, 4, 5),
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
		input    ForwardList[int]
		index    int
		expected int
	}

	atTests := []atTest{
		{
			input:    Of(1, 2, 3, 4),
			index:    0,
			expected: 1,
		},
		{
			input:    Of(1, 2, 3, 4),
			index:    3,
			expected: 4,
		},
		{
			input:    Of(1, 2, 3, 4),
			index:    1,
			expected: 2,
		},
	}

	for _, test := range atTests {
		fmt.Println(test.input.ToSlice())
		assert.Equal(t, test.expected, test.input.At(test.index))
	}
}
