package treeset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	lessThan    = func(k1, k2 string) bool { return k1 < k2 }
	lessThanInt = func(k1, k2 int) bool { return k1 < k2 }
)

func TestNew(t *testing.T) {

	set := New[string](lessThan)
	assert.NotNil(t, set)
	assert.True(t, set.Empty())
	assert.Equal(t, 0, set.Len())
}

func TestOf(t *testing.T) {

	assert.Equal(t, []string{}, Of[string](lessThan).treeMap.Keys())
	assert.Equal(t, []string{"A"}, Of(lessThan, "A").treeMap.Keys())

}

func TestAdd(t *testing.T) {

	type addTest struct {
		input    []string
		expected TreeSet[string]
	}

	addTests := []addTest{
		{
			input:    []string{},
			expected: Of[string](lessThan),
		},
		{
			input:    []string{"A", "A", "B"},
			expected: Of(lessThan, "A", "B"),
		},
	}

	f := func(values []string) TreeSet[string] {
		set := New[string](lessThan)
		for _, value := range values {
			set.Add(value)
		}
		return *set
	}

	for _, test := range addTests {
		assert.Equal(t, test.expected.treeMap.Keys(), f(test.input).treeMap.Keys())
	}
}

func TestAddSlice(t *testing.T) {

	type addSliceTest struct {
		input    []string
		expected TreeSet[string]
	}

	addSliceTests := []addSliceTest{
		{
			input:    []string{},
			expected: Of[string](lessThan),
		},
		{
			input:    []string{"A", "A", "B"},
			expected: Of(lessThan, "A", "B"),
		},
	}

	for _, test := range addSliceTests {
		set := New[string](lessThan)
		set.AddSlice(test.input)
		assert.Equal(t, test.expected.treeMap.Keys(), set.treeMap.Keys())
	}

}

func TestRemove(t *testing.T) {

	type removeTest struct {
		input           string
		expectedSet     TreeSet[string]
		expectedBoolean bool
	}

	removeTests := []removeTest{
		{
			input:           "",
			expectedSet:     Of(lessThan, "A", "B", "C"),
			expectedBoolean: false,
		},
		{
			input:           "A",
			expectedSet:     Of(lessThan, "B", "C"),
			expectedBoolean: true,
		},
	}

	for _, test := range removeTests {
		set := Of(lessThan, "A", "B", "C")
		assert.Equal(t, test.expectedBoolean, set.Remove(test.input))
		assert.Equal(t, test.expectedSet.treeMap.Keys(), set.treeMap.Keys())
	}

}

func TestRemoveIf(t *testing.T) {

	type removeIfTest struct {
		input           TreeSet[int]
		expectedBoolean bool
		expectedSet     TreeSet[int]
	}

	removeIfTests := []removeIfTest{
		{
			input:           Of[int](lessThanInt),
			expectedBoolean: false,
			expectedSet:     Of[int](lessThanInt),
		},
		{
			input:           Of(lessThanInt, 2),
			expectedBoolean: false,
			expectedSet:     Of(lessThanInt, 2),
		},
		{
			input:           Of(lessThanInt, 1, 2, 3, 4, 5),
			expectedBoolean: true,
			expectedSet:     Of(lessThanInt, 2, 4),
		},
	}

	f := func(x int) bool {
		return x%2 != 0
	}

	for _, test := range removeIfTests {
		test.input.RemoveIf(f)
		assert.Equal(t, test.expectedSet.treeMap.Keys(), test.input.treeMap.Keys())
	}
}

func TestRemoveSlice(t *testing.T) {

	type removeSliceTest struct {
		input           TreeSet[int]
		slice           []int
		expectedBoolean bool
		expectedSet     TreeSet[int]
	}

	removeSliceTests := []removeSliceTest{
		{
			input:           Of[int](lessThanInt),
			slice:           []int{},
			expectedBoolean: false,
			expectedSet:     Of[int](lessThanInt),
		},
		{
			input:           Of(lessThanInt, 2),
			slice:           []int{3},
			expectedBoolean: false,
			expectedSet:     Of(lessThanInt, 2),
		},
		{
			input:           Of(lessThanInt, 1, 2, 3, 4, 5),
			slice:           []int{2, 3, 1, 4},
			expectedBoolean: true,
			expectedSet:     Of(lessThanInt, 5),
		},
	}

	for _, test := range removeSliceTests {
		test.input.RemoveSlice(test.slice)
		assert.Equal(t, test.expectedSet.treeMap.Keys(), test.input.treeMap.Keys())
	}
}

func TestClear(t *testing.T) {

	set := Of(lessThanInt, 1, 2, 3, 4, 5)
	set.Clear()

	assert.NotNil(t, set)
	assert.True(t, set.Empty())

}

func TestContains(t *testing.T) {

	type containsTest struct {
		input    TreeSet[int]
		element  int
		expected bool
	}

	containsTests := []containsTest{
		{
			input:    Of(lessThanInt, 0, 4, 5),
			element:  1,
			expected: false,
		},
		{
			input:    Of(lessThanInt, 0, 4, 5),
			element:  2,
			expected: false,
		},
		{
			input:    Of(lessThanInt, 0, 4, 5),
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
		a        TreeSet[int]
		b        TreeSet[int]
		expected TreeSet[int]
	}

	addAllTests := []addAllTest{
		{
			a:        Of[int](lessThanInt),
			b:        Of(lessThanInt, 1, 2, 3, 4, 5),
			expected: Of(lessThanInt, 1, 2, 3, 4, 5),
		},
		{
			a:        Of(lessThanInt, 1, 2),
			b:        Of(lessThanInt, 9, 11, 12),
			expected: Of(lessThanInt, 1, 2, 9, 11, 12),
		},
	}

	for _, test := range addAllTests {
		test.a.AddAll(&test.b)
		assert.Equal(t, test.expected.treeMap.Keys(), test.a.treeMap.Keys())
	}

}

func TestRemoveAll(t *testing.T) {

	type removeAllTest struct {
		a        TreeSet[int]
		b        TreeSet[int]
		expected TreeSet[int]
	}

	removeAllTests := []removeAllTest{
		{
			a:        Of[int](lessThanInt),
			b:        Of[int](lessThanInt),
			expected: Of[int](lessThanInt),
		},
		{
			a:        Of(lessThanInt, 1, 2, 3, 4, 5),
			b:        Of[int](lessThanInt),
			expected: Of(lessThanInt, 1, 2, 3, 4, 5),
		},
		{
			a:        Of(lessThanInt, 1, 2, 3, 4, 5),
			b:        Of(lessThanInt, 9, 1, 2),
			expected: Of(lessThanInt, 3, 4, 5),
		},
		{
			a:        Of(lessThanInt, 1, 2, 3, 4, 5),
			b:        Of(lessThanInt, 9, 1, 2, 3, 4, 5),
			expected: Of[int](lessThanInt),
		},
	}

	for _, test := range removeAllTests {
		test.a.RemoveAll(&test.b)
		assert.Equal(t, test.expected.treeMap.Keys(), test.a.treeMap.Keys())
	}

}

func TestRetainAll(t *testing.T) {

	type retainAllTest struct {
		a        TreeSet[int]
		b        TreeSet[int]
		expected TreeSet[int]
	}

	retainAllTests := []retainAllTest{
		{
			a:        Of(lessThanInt, 1, 2, 3, 4, 5),
			b:        Of[int](lessThanInt),
			expected: Of[int](lessThanInt),
		},
		{
			a:        Of(lessThanInt, 1, 2, 3, 4, 5),
			b:        Of(lessThanInt, 9, 1, 2),
			expected: Of(lessThanInt, 1, 2),
		},
		{
			a:        Of(lessThanInt, 1, 2, 3, 4, 5),
			b:        Of(lessThanInt, 9, 1, 2, 3, 4, 5),
			expected: Of(lessThanInt, 1, 2, 3, 4, 5),
		},
	}

	for _, test := range retainAllTests {
		test.a.RetainAll(&test.b)
		assert.Equal(t, test.expected.treeMap.Keys(), test.a.treeMap.Keys())
	}

}

func TestForEach(t *testing.T) {

	set := Of(lessThanInt, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0

	set.ForEach(func(i int) { sum = sum + i })

	assert.Equal(t, 55, sum)
}

func TestEquals(t *testing.T) {

	type equalsTest struct {
		a        TreeSet[int]
		b        TreeSet[int]
		expected bool
	}

	equalsTests := []equalsTest{
		{
			a:        Of[int](lessThanInt),
			b:        Of[int](lessThanInt),
			expected: true,
		},
		{
			a:        Of(lessThanInt, 1, 2),
			b:        Of[int](lessThanInt),
			expected: false,
		},
		{
			a:        Of(lessThanInt, 1, 2),
			b:        Of(lessThanInt, 2, 1),
			expected: true,
		},
		{
			a:        Of(lessThanInt, 1, 2, 3),
			b:        Of(lessThanInt, 10, 12, 14),
			expected: false,
		},
	}

	for _, test := range equalsTests {
		assert.Equal(t, test.expected, test.a.Equals(&test.b))
		assert.Equal(t, test.expected, test.b.Equals(&test.a))

	}

	identity := Of[int](lessThanInt)
	assert.True(t, identity.Equals(&identity))

}

func TestToSlice(t *testing.T) {

	type toSliceTest struct {
		input    TreeSet[int]
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
		assert.Equal(t, test.expected, test.input.ToSlice())
	}
}

func TestString(t *testing.T) {

	assert.Equal(t, "{}", Of(lessThanInt).String())
	assert.Equal(t, "{1, 2, 3}", Of(lessThanInt, 1, 2, 3).String())
}
