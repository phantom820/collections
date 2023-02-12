package hashset

import (
	"testing"

	"github.com/phantom820/collections/maps/hashmap"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	set := New[string]()
	assert.NotNil(t, set)
	assert.True(t, set.Empty())
	assert.Equal(t, 0, set.Len())
}

func TestOf(t *testing.T) {

	assert.Equal(t, hashmap.New[string, struct{}](), Of[string]().hashmap)
	assert.Equal(t, hashmap.HashMap[string, struct{}](map[string]struct{}{"A": {}}), Of("A").hashmap)

}

func TestAdd(t *testing.T) {

	type addTest struct {
		input    []string
		expected HashSet[string]
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

	f := func(values []string) HashSet[string] {
		set := New[string]()
		for _, value := range values {
			set.Add(value)
		}
		return *set
	}

	for _, test := range addTests {
		assert.Equal(t, test.expected, f(test.input))
	}
}

func TestAddSlice(t *testing.T) {

	type addSliceTest struct {
		input    []string
		expected HashSet[string]
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
		assert.Equal(t, test.expected, *set)
	}

}

func TestRemove(t *testing.T) {

	type removeTest struct {
		input           string
		expectedSet     HashSet[string]
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
		assert.Equal(t, test.expectedSet, set)
	}

}

func TestRemoveIf(t *testing.T) {

	type removeIfTest struct {
		input           HashSet[int]
		expectedBoolean bool
		expectedSet     HashSet[int]
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
		assert.Equal(t, test.expectedSet, test.input)
	}
}

func TestRemoveSlice(t *testing.T) {

	type removeSliceTest struct {
		input           HashSet[int]
		slice           []int
		expectedBoolean bool
		expectedSet     HashSet[int]
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

	for _, test := range removeSliceTests {
		test.input.RemoveSlice(test.slice)
		assert.Equal(t, test.expectedSet, test.input)
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

	set := Of(0, 4, 5)
	for _, test := range containsTests {
		assert.Equal(t, test.expected, set.Contains(test.input))
	}
}

func TestAddAll(t *testing.T) {

	type addAllTest struct {
		setA     HashSet[int]
		setB     HashSet[int]
		expected HashSet[int]
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
		assert.Equal(t, test.expected, test.setA)
	}

}

func TestRemoveAll(t *testing.T) {

	type removeAllTest struct {
		setA     HashSet[int]
		setB     HashSet[int]
		expected HashSet[int]
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
		assert.Equal(t, test.expected, test.setA)
	}

}

func TestRetainAll(t *testing.T) {

	type retainAllTest struct {
		setA     HashSet[int]
		setB     HashSet[int]
		expected HashSet[int]
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
		assert.Equal(t, test.expected, test.setA)
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
		a        HashSet[int]
		b        HashSet[int]
		expected bool
	}

	equalsTests := []equalsTest{
		{
			a:        Of[int](),
			b:        Of[int](),
			expected: true,
		},
		{
			a:        Of[int](1, 2),
			b:        Of[int](),
			expected: false,
		},
		{
			a:        Of[int](1, 2),
			b:        Of[int](2, 1),
			expected: true,
		},
		{
			a:        Of[int](1, 2, 3),
			b:        Of[int](10, 12, 14),
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
	assert.Equal(t, "{1}", Of(1).String())
}
