package linkedhashset

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

	assert.Equal(t, []string{}, Of[string]().linkedHashMap.Keys())
	assert.Equal(t, []string{"A"}, Of("A").linkedHashMap.Keys())

}

func TestAdd(t *testing.T) {

	type addTest struct {
		input    []string
		expected LinkedHashSet[string]
	}

	addTests := []addTest{
		{
			input:    []string{},
			expected: Of[string](),
		},
		{
			input:    []string{"A", "A", "B"},
			expected: Of("A", "B"),
		},
	}

	f := func(values []string) LinkedHashSet[string] {
		set := New[string]()
		for _, value := range values {
			set.Add(value)
		}
		return *set
	}

	for _, test := range addTests {
		assert.Equal(t, test.expected.linkedHashMap.Keys(), f(test.input).linkedHashMap.Keys())
	}
}

func TestAddSlice(t *testing.T) {

	type addSliceTest struct {
		input    []string
		expected LinkedHashSet[string]
	}

	addSliceTests := []addSliceTest{
		{
			input:    []string{},
			expected: Of[string](),
		},
		{
			input:    []string{"A", "A", "B"},
			expected: Of("A", "B"),
		},
	}

	for _, test := range addSliceTests {
		set := New[string]()
		set.AddSlice(test.input)
		assert.Equal(t, test.expected.linkedHashMap.Keys(), set.linkedHashMap.Keys())
	}

}

func TestRemove(t *testing.T) {

	type removeTest struct {
		input           string
		expectedSet     LinkedHashSet[string]
		expectedBoolean bool
	}

	removeTests := []removeTest{
		{
			input:           "",
			expectedSet:     Of("A", "B", "C"),
			expectedBoolean: false,
		},
		{
			input:           "A",
			expectedSet:     Of("B", "C"),
			expectedBoolean: true,
		},
	}

	for _, test := range removeTests {
		set := Of("A", "B", "C")
		assert.Equal(t, test.expectedBoolean, set.Remove(test.input))
		assert.Equal(t, test.expectedSet.linkedHashMap.Keys(), set.linkedHashMap.Keys())
	}

}

func TestRemoveIf(t *testing.T) {

	type removeIfTest struct {
		input           LinkedHashSet[int]
		expectedBoolean bool
		expectedSet     LinkedHashSet[int]
	}

	removeIfTests := []removeIfTest{
		{
			input:           Of[int](),
			expectedBoolean: false,
			expectedSet:     Of[int](),
		},
		{
			input:           Of(2),
			expectedBoolean: false,
			expectedSet:     Of(2),
		},
		{
			input:           Of(1, 2, 3, 4, 5),
			expectedBoolean: true,
			expectedSet:     Of(2, 4),
		},
	}

	f := func(x int) bool {
		return x%2 != 0
	}

	for _, test := range removeIfTests {
		test.input.RemoveIf(f)
		assert.Equal(t, test.expectedSet.linkedHashMap.Keys(), test.input.linkedHashMap.Keys())
	}
}

func TestRemoveSlice(t *testing.T) {

	type removeSliceTest struct {
		input           LinkedHashSet[int]
		slice           []int
		expectedBoolean bool
		expectedSet     LinkedHashSet[int]
	}

	removeSliceTests := []removeSliceTest{
		{
			input:           Of[int](),
			slice:           []int{},
			expectedBoolean: false,
			expectedSet:     Of[int](),
		},
		{
			input:           Of(2),
			slice:           []int{3},
			expectedBoolean: false,
			expectedSet:     Of(2),
		},
		{
			input:           Of(1, 2, 3, 4, 5),
			slice:           []int{2, 3, 1, 4},
			expectedBoolean: true,
			expectedSet:     Of(5),
		},
	}

	for _, test := range removeSliceTests[2:] {
		test.input.RemoveSlice(test.slice)
		assert.Equal(t, test.expectedSet.linkedHashMap.Keys(), test.input.linkedHashMap.Keys())
	}
}

func TestClear(t *testing.T) {

	set := Of(1, 2, 3, 4, 5)
	set.Clear()

	assert.NotNil(t, set)
	assert.True(t, set.Empty())

}

func TestContains(t *testing.T) {

	type containsTest struct {
		input    int
		expected bool
	}

	set := Of(0, 4, 5)
	containsTests := []containsTest{
		{
			input:    1,
			expected: false,
		},
		{
			input:    2,
			expected: false,
		},
		{
			input:    4,
			expected: true,
		},
	}

	for _, test := range containsTests {
		assert.Equal(t, test.expected, set.Contains(test.input))
	}
}

func TestAddAll(t *testing.T) {

	type addAllTest struct {
		setA     LinkedHashSet[int]
		setB     LinkedHashSet[int]
		expected LinkedHashSet[int]
	}

	addAllTests := []addAllTest{
		{
			setA:     Of[int](),
			setB:     Of(1, 2, 3, 4, 5),
			expected: Of(1, 2, 3, 4, 5),
		},
		{
			setA:     Of(1, 2),
			setB:     Of(9, 11, 12),
			expected: Of(1, 2, 9, 11, 12),
		},
	}

	for _, test := range addAllTests {
		test.setA.AddAll(&test.setB)
		assert.Equal(t, test.expected.linkedHashMap.Keys(), test.setA.linkedHashMap.Keys())
	}

}

func TestRemoveAll(t *testing.T) {

	type removeAllTest struct {
		setA     LinkedHashSet[int]
		setB     LinkedHashSet[int]
		expected LinkedHashSet[int]
	}

	removeAllTests := []removeAllTest{
		{
			setA:     Of(1, 2, 3, 4, 5),
			setB:     Of[int](),
			expected: Of(1, 2, 3, 4, 5),
		},
		{
			setA:     Of(1, 2, 3, 4, 5),
			setB:     Of(9, 1, 2),
			expected: Of(3, 4, 5),
		},
		{
			setA:     Of(1, 2, 3, 4, 5),
			setB:     Of(9, 1, 2, 3, 4, 5),
			expected: Of[int](),
		},
	}

	for _, test := range removeAllTests {
		test.setA.RemoveAll(&test.setB)
		assert.Equal(t, test.expected.linkedHashMap.Keys(), test.setA.linkedHashMap.Keys())
	}

}

func TestRetainAll(t *testing.T) {

	type retainAllTest struct {
		setA     LinkedHashSet[int]
		setB     LinkedHashSet[int]
		expected LinkedHashSet[int]
	}

	retainAllTests := []retainAllTest{
		{
			setA:     Of(1, 2, 3, 4, 5),
			setB:     Of[int](),
			expected: Of[int](),
		},
		{
			setA:     Of(1, 2, 3, 4, 5),
			setB:     Of(9, 1, 2),
			expected: Of(1, 2),
		},
		{
			setA:     Of(1, 2, 3, 4, 5),
			setB:     Of(9, 1, 2, 3, 4, 5),
			expected: Of(1, 2, 3, 4, 5),
		},
	}

	for _, test := range retainAllTests {
		test.setA.RetainAll(&test.setB)
		assert.Equal(t, test.expected.linkedHashMap.Keys(), test.setA.linkedHashMap.Keys())
	}

}

func TestForEach(t *testing.T) {

	set := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0

	set.ForEach(func(i int) { sum = sum + i })

	assert.Equal(t, 55, sum)
}

func TestEquals(t *testing.T) {

	type equalsTest struct {
		a        LinkedHashSet[int]
		b        LinkedHashSet[int]
		expected bool
	}

	equalsTests := []equalsTest{
		{
			a:        Of[int](),
			b:        Of[int](),
			expected: true,
		},
		{
			a:        Of(1, 2),
			b:        Of[int](),
			expected: false,
		},
		{
			a:        Of(1, 2),
			b:        Of(2, 1),
			expected: true,
		},
		{
			a:        Of(1, 2, 3),
			b:        Of(10, 12, 14),
			expected: false,
		},
	}

	for _, test := range equalsTests {
		assert.Equal(t, test.expected, test.a.Equals(&test.b))
		assert.Equal(t, test.expected, test.b.Equals(&test.a))

	}

	identity := Of[int]()
	assert.True(t, identity.Equals(&identity))

}

func TestString(t *testing.T) {

	assert.Equal(t, "{}", Of[int]().String())
	assert.Equal(t, "{1, 2, 3}", Of(1, 2, 3).String())
}
