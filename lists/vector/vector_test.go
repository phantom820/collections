package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	set := New[string]()
	assert.NotNil(t, set)
	assert.True(t, set.Empty())
	assert.Equal(t, 0, set.Len())
}

func TestOf(t *testing.T) {

	assert.Equal(t, []string{}, Of[string]().data)
	assert.Equal(t, []string{"A"}, Of("A").data)

}

func TestAdd(t *testing.T) {

	type addTest struct {
		input    []string
		expected Vector[string]
	}

	addTests := []addTest{
		{
			input:    []string{},
			expected: Of[string](),
		},
		{
			input:    []string{"A", "A", "B"},
			expected: Of("A", "A", "B"),
		},
	}

	f := func(values []string) Vector[string] {
		list := New[string]()
		for _, value := range values {
			list.Add(value)
		}
		return *list
	}

	for _, test := range addTests {
		assert.Equal(t, test.expected, f(test.input))
	}
}

func TestAddSlice(t *testing.T) {

	type addSliceTest struct {
		input    []string
		expected Vector[string]
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
		set := New[string]()
		set.AddSlice(test.input)
		assert.Equal(t, test.expected, *set)
	}

}

func TestRemove(t *testing.T) {

	type removeTest struct {
		input           Vector[string]
		element         string
		expectedList    Vector[string]
		expectedBoolean bool
		expectedLen     int
	}

	removeTests := []removeTest{
		{
			input:           Of("A", "B", "C", "D"),
			element:         "",
			expectedList:    Of("A", "B", "C", "D"),
			expectedBoolean: false,
			expectedLen:     4,
		},
		{
			input:           Of("A", "B", "C", "D"),
			element:         "A",
			expectedList:    Of("B", "C", "D"),
			expectedBoolean: true,
			expectedLen:     3,
		},
		{
			input:           Of("A", "B", "C", "D"),
			element:         "B",
			expectedList:    Of("A", "C", "D"),
			expectedBoolean: true,
			expectedLen:     3,
		},
		{
			input:           Of("A", "B", "C", "D"),
			element:         "D",
			expectedList:    Of("A", "B", "C"),
			expectedBoolean: true,
			expectedLen:     3,
		},
		{
			input:           Of("A"),
			element:         "A",
			expectedList:    Of[string](),
			expectedBoolean: true,
			expectedLen:     0,
		},
	}

	for _, test := range removeTests {
		assert.Equal(t, test.expectedBoolean, test.input.Remove(test.element))
		assert.Equal(t, test.expectedList, test.input)
		assert.Equal(t, test.expectedLen, test.input.Len())
	}

}

func TestRemoveIf(t *testing.T) {

	type removeIfTest struct {
		input           Vector[int]
		expectedBoolean bool
		expectedList    Vector[int]
	}

	removeIfTests := []removeIfTest{
		{
			input:           Of[int](),
			expectedBoolean: false,
			expectedList:    Of[int](),
		},
		{
			input:           Of(2),
			expectedBoolean: false,
			expectedList:    Of(2),
		},
		{
			input:           Of(1, 2, 3, 4, 5),
			expectedBoolean: true,
			expectedList:    Of(2, 4),
		},
		{
			input:           Of(1, 3, 5, 7),
			expectedBoolean: true,
			expectedList:    Of[int](),
		},
	}

	f := func(x int) bool {
		return x%2 != 0
	}

	for _, test := range removeIfTests {
		test.input.RemoveIf(f)
		assert.Equal(t, test.expectedList, test.input)
	}
}

func TestRemoveSlice(t *testing.T) {

	type removeSliceTest struct {
		input           Vector[int]
		slice           []int
		expectedBoolean bool
		expectedList    Vector[int]
	}

	removeSliceTests := []removeSliceTest{
		{
			input:           Of[int](),
			slice:           []int{},
			expectedBoolean: false,
			expectedList:    Of[int](),
		},
		{
			input:           Of(2),
			slice:           []int{3},
			expectedBoolean: false,
			expectedList:    Of(2),
		},
		{
			input:           Of(1, 2, 3, 4, 5),
			slice:           []int{2, 3, 1, 4},
			expectedBoolean: true,
			expectedList:    Of(5),
		},
	}

	for _, test := range removeSliceTests {
		test.input.RemoveSlice(test.slice)
		assert.Equal(t, test.expectedList, test.input)
	}
}

func TestClear(t *testing.T) {

	list := Of(1, 2, 3, 4, 5)
	list.Clear()

	assert.NotNil(t, list)
	assert.True(t, list.Empty())

}

func TestContains(t *testing.T) {

	type containsTest struct {
		input    Vector[int]
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

func TestAddAll(t *testing.T) {

	type addAllTest struct {
		a        Vector[int]
		b        Vector[int]
		expected Vector[int]
	}

	addAllTests := []addAllTest{
		{
			a:        Of[int](),
			b:        Of(1, 2, 3, 4, 5),
			expected: Of(1, 2, 3, 4, 5),
		},
		{
			a:        Of(1, 2),
			b:        Of(9, 11, 12),
			expected: Of(1, 2, 9, 11, 12),
		},
	}

	for _, test := range addAllTests {
		test.a.AddAll(&test.b)
		assert.Equal(t, test.expected, test.a)
	}

}

func TestRetainAll(t *testing.T) {

	type retainAllTest struct {
		a        Vector[int]
		b        Vector[int]
		expected Vector[int]
	}

	retainAllTests := []retainAllTest{
		{
			a:        Of(1, 2, 3, 4, 5),
			b:        Of[int](),
			expected: Of[int](),
		},
		{
			a:        Of(1, 2, 3, 4, 5),
			b:        Of(9, 1, 2),
			expected: Of(1, 2),
		},
		{
			a:        Of(1, 2, 3, 4, 5),
			b:        Of(9, 1, 2, 3, 4, 5),
			expected: Of(1, 2, 3, 4, 5),
		},
	}

	for _, test := range retainAllTests {
		test.a.RetainAll(&test.b)
		assert.Equal(t, test.expected, test.a)
	}

}

func TestRemoveAll(t *testing.T) {

	type removeAllTest struct {
		a        Vector[int]
		b        Vector[int]
		expected Vector[int]
	}

	removeAllTests := []removeAllTest{
		{
			a:        Of[int](),
			b:        Of[int](),
			expected: Of[int](),
		},
		{
			a:        Of(1, 2, 3, 4, 5),
			b:        Of[int](),
			expected: Of(1, 2, 3, 4, 5),
		},
		{
			a:        Of(1, 2, 3, 4, 5),
			b:        Of(9, 1, 2),
			expected: Of(3, 4, 5),
		},
		{
			a:        Of(1, 2, 3, 4, 5),
			b:        Of(9, 1, 2, 3, 4, 5),
			expected: Of[int](),
		},
	}

	for _, test := range removeAllTests {
		test.a.RemoveAll(&test.b)
		assert.Equal(t, test.expected, test.a)
	}

}

func TestAt(t *testing.T) {

	type atTest struct {
		input    Vector[int]
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
		assert.Equal(t, test.expected, test.input.At(test.index))
	}
}

func TestForEach(t *testing.T) {

	list := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0

	list.ForEach(func(i int) { sum = sum + i })

	assert.Equal(t, 55, sum)
}
