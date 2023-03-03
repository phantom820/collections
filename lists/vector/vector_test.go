package vector

import (
	"testing"

	"github.com/phantom820/collections"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	list := New[string]()
	assert.NotNil(t, list)
	assert.True(t, list.Empty())
	assert.Equal(t, 0, list.Len())
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
		list := New[string]()
		list.AddSlice(test.input)
		assert.Equal(t, test.expected, *list)
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
		{
			a:        Of[int](),
			b:        Of(9, 1, 2, 3, 4, 5),
			expected: Of[int](),
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

func TestIndexOf(t *testing.T) {

	type indexOfTest struct {
		input    Vector[int]
		expected int
	}

	indexOfTests := []indexOfTest{
		{
			input:    Of[int](),
			expected: -1,
		},
		{
			input:    Of[int](1, 2, 3, 4),
			expected: 0,
		},
		{
			input:    Of[int](0, 2, 1, 4),
			expected: 2,
		},
		{
			input:    Of[int](0, 1, 1, 4),
			expected: 1,
		},
	}

	for _, test := range indexOfTests {
		assert.Equal(t, test.expected, test.input.IndexOf(1))
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

func TestRemoveAt(t *testing.T) {

	type removeAtTest struct {
		input           Vector[int]
		index           int
		expectedElement int
		expectedList    Vector[int]
	}

	removeAtTests := []removeAtTest{
		{
			input:           Of(1, 2, 3, 4),
			index:           0,
			expectedElement: 1,
			expectedList:    Of(2, 3, 4),
		},
		{
			input:           Of(1, 2, 3, 4),
			index:           3,
			expectedElement: 4,
			expectedList:    Of(1, 2, 3),
		},
		{
			input:           Of(1, 2, 3, 4),
			index:           2,
			expectedElement: 3,
			expectedList:    Of(1, 2, 4),
		},
	}

	for _, test := range removeAtTests {
		assert.Equal(t, test.expectedElement, test.input.RemoveAt(test.index))
		assert.Equal(t, test.expectedList, test.input)
	}

}

func TestAddAt(t *testing.T) {

	type addAtTest struct {
		input    Vector[int]
		index    int
		value    int
		expected Vector[int]
	}

	addAtTests := []addAtTest{
		{
			input:    Of(1),
			index:    0,
			value:    -1,
			expected: Of(-1, 1),
		},
		{
			input:    Of(1, 2, 3),
			index:    1,
			value:    -2,
			expected: Of(1, -2, 2, 3),
		},
		{
			input:    Of(1, 2, 3),
			index:    2,
			value:    4,
			expected: Of(1, 2, 3, 4),
		},
	}

	for _, test := range addAtTests {
		test.input.AddAt(test.index, test.value)
		assert.Equal(t, test.expected, test.input)
	}
}

func TestSet(t *testing.T) {

	type setTest struct {
		input    Vector[int]
		index    int
		value    int
		expected Vector[int]
	}

	setTests := []setTest{
		{
			input:    Of(1),
			index:    0,
			value:    -1,
			expected: Of(-1),
		},
		{
			input:    Of(1, 2, 3),
			index:    1,
			value:    -2,
			expected: Of(1, -2, 3),
		},
		{
			input:    Of(1, 2, 3),
			index:    2,
			value:    4,
			expected: Of(1, 2, 4),
		},
	}

	for _, test := range setTests {
		test.input.Set(test.index, test.value)
		assert.Equal(t, test.expected, test.input)
	}
}

func TestSubList(t *testing.T) {

	type subListTest struct {
		input      Vector[int]
		start, end int
		expected   Vector[int]
	}

	subListTests := []subListTest{
		{
			input:    Of(1),
			start:    0,
			end:      0,
			expected: Of[int](),
		},
		{
			input:    Of(1, 2),
			start:    0,
			end:      1,
			expected: Of(1),
		},
		{
			input:    Of(1, 2, 3, 4, 5),
			start:    0,
			end:      4,
			expected: Of(1, 2, 3, 4),
		},
		{
			input:    Of(1, 2, 3, 4, 5),
			start:    1,
			end:      4,
			expected: Of(2, 3, 4),
		},
		{
			input:    Of(1, 2, 3, 4, 5),
			start:    0,
			end:      5,
			expected: Of(1, 2, 3, 4, 5),
		},
		{
			input:    Of(1, 2, 3, 4, 5),
			start:    2,
			end:      5,
			expected: Of(3, 4, 5),
		},
	}

	for _, test := range subListTests {
		assert.Equal(t, test.expected.data, test.input.SubList(test.start, test.end).ToSlice())
	}
}

// func TestImmutableSubList(t *testing.T) {

// 	type immutableSubListTest struct {
// 		input      Vector[int]
// 		start, end int
// 		expected   ImmutableVector[int]
// 	}

// 	subListTests := []immutableSubListTest{
// 		{
// 			input:    Of(1),
// 			start:    0,
// 			end:      0,
// 			expected: ImmutableOf[int](),
// 		},
// 		{
// 			input:    Of(1, 2),
// 			start:    0,
// 			end:      1,
// 			expected: ImmutableOf(1),
// 		},
// 		{
// 			input:    Of(1, 2, 3, 4, 5),
// 			start:    0,
// 			end:      4,
// 			expected: ImmutableOf(1, 2, 3, 4),
// 		},
// 	}

//		for _, test := range subListTests {
//			assert.Equal(t, test.expected, test.input.ImmutableSubList(test.start, test.end))
//		}
//	}
func TestForEach(t *testing.T) {

	list := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := 0

	list.ForEach(func(i int) { sum = sum + i })

	assert.Equal(t, 55, sum)
}

func TestEquals(t *testing.T) {

	type equalsTest struct {
		a        Vector[int]
		b        Vector[int]
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
			b:        Of[int](1, 2),
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

func TestIterator(t *testing.T) {

	type iteratorTest struct {
		input    Vector[int]
		expected []int
	}

	iteratorTests := []iteratorTest{
		{
			input:    Of[int](),
			expected: []int{},
		},
		{
			input:    Of[int](1, 2, 3, 4),
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    Of[int](1),
			expected: []int{1},
		},
	}

	iterate := func(it collections.Iterator[int]) []int {
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

	assert.Equal(t, "[]", Of[int]().String())
	assert.Equal(t, "[1 2]", Of(1, 2).String())
}

func TestSort(t *testing.T) {

	type sortTest struct {
		input    Vector[int]
		less     func(int, int) bool
		expected Vector[int]
	}

	sortTests := []sortTest{
		{
			input:    Of[int](),
			less:     func(i1, i2 int) bool { return i1 < i2 },
			expected: Of[int](),
		},
		{
			input:    Of[int](2, 1, 4),
			less:     func(i1, i2 int) bool { return i1 < i2 },
			expected: Of[int](1, 2, 4),
		},
		{
			input:    Of[int](1, 2, 3, 5, 4),
			less:     func(i1, i2 int) bool { return i1 < i2 },
			expected: Of[int](1, 2, 3, 4, 5),
		},
		{
			input:    Of[int](5, 4, 3, 2, 1),
			less:     func(i1, i2 int) bool { return i1 <= i2 },
			expected: Of[int](1, 2, 3, 4, 5),
		},
		{
			input:    Of[int](1, 2, 3, 4, 5),
			less:     func(i1, i2 int) bool { return i1 >= i2 },
			expected: Of[int](5, 4, 3, 2, 1),
		},
	}

	for _, test := range sortTests {
		test.input.Sort(test.less)
		assert.Equal(t, test.expected, test.input)
	}

}

func TestCopy(t *testing.T) {

	type copyTest struct {
		input    Vector[int]
		expected []int
	}

	copyTests := []copyTest{
		{
			input:    Of[int](),
			expected: []int{},
		},
		{
			input:    Of[int](1, 2, 3, 4),
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, test := range copyTests {
		assert.Equal(t, test.expected, test.input.Copy().ToSlice())
		assert.Equal(t, test.expected, test.input.ImmutableCopy().ToSlice())
	}
}
