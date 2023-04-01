package listdequeue

import (
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/types/optional"
	"github.com/stretchr/testify/assert"
)

func data(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = i + 1
	}
	return data
}

func TestNew(t *testing.T) {

	deq := New[int]()

	assert.NotNil(t, deq)
	assert.NotNil(t, deq.list)
	_, ok := deq.list.(*linkedlist.LinkedList[int])
	assert.True(t, ok)
	assert.True(t, deq.Empty())
}

func TestPeekFirst(t *testing.T) {

	pekkFirstTests := []struct {
		input    *ListDequeue[int]
		expected optional.Optional[int]
	}{
		{
			input:    New[int](),
			expected: optional.Empty[int](),
		},
		{
			input:    New(1),
			expected: optional.Of(1),
		},
		{
			input:    New(1, 2, 3),
			expected: optional.Of(1),
		},
	}

	for _, test := range pekkFirstTests {
		assert.Equal(t, test.expected, test.input.PeekFirst())
	}
}

func TestPeekLast(t *testing.T) {

	peekFirstTests := []struct {
		input    *ListDequeue[int]
		expected optional.Optional[int]
	}{
		{
			input:    New[int](),
			expected: optional.Empty[int](),
		},
		{
			input:    New(1),
			expected: optional.Of(1),
		},
		{
			input:    New(1, 2, 3),
			expected: optional.Of(3),
		},
	}

	for _, test := range peekFirstTests {
		assert.Equal(t, test.expected, test.input.PeekLast())
	}
}

func TestAddFirst(t *testing.T) {

	addFirstTests := []struct {
		input    *ListDequeue[int]
		elements []int
		expected []int
	}{
		{
			input:    New[int](),
			elements: []int{},
			expected: []int{},
		},
		{
			input:    New[int](),
			elements: []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, test := range addFirstTests {
		for _, e := range test.elements {
			test.input.AddFirst(e)
		}
		assert.Equal(t, test.expected, test.input.ToSlice())
	}

}

func TestAddLast(t *testing.T) {

	addLastTests := []struct {
		input    *ListDequeue[int]
		elements []int
		expected []int
	}{
		{
			input:    New[int](),
			elements: []int{},
			expected: []int{},
		},
		{
			input:    New[int](),
			elements: []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range addLastTests {
		for _, e := range test.elements {
			test.input.AddLast(e)
		}
		assert.Equal(t, test.expected, test.input.ToSlice())
	}

}

func TestRemoveFirst(t *testing.T) {

	queue := New[int]()
	data := data(1000)

	for _, e := range data {
		queue.AddLast(e)
	}

	for _, e := range data {
		assert.Equal(t, e, queue.RemoveFirst().Value())
	}

	assert.True(t, queue.Empty())

}

func TestRemoveLast(t *testing.T) {

	queue := New[int]()
	data := data(1000)

	for _, e := range data {
		queue.AddFirst(e)
	}

	for _, e := range data {
		assert.Equal(t, e, queue.RemoveLast().Value())
	}

	assert.True(t, queue.Empty())
	assert.Equal(t, optional.Empty[int](), queue.RemoveFirst())

}

func TestContains(t *testing.T) {

	type containsTest struct {
		input    *ListDequeue[int]
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
		{
			input:    New[int](),
			element:  4,
			expected: false,
		},
	}

	for _, test := range containsTests {
		assert.Equal(t, test.expected, test.input.Contains(test.element))
	}
}

func TestIterator(t *testing.T) {

	type iteratorTest struct {
		input    *ListDequeue[int]
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

func TestAddSlice(t *testing.T) {

	type addSliceTest struct {
		input    *ListDequeue[int]
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

func TestAddAll(t *testing.T) {

	type addAllTest struct {
		a        *ListDequeue[int]
		b        *ListDequeue[int]
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

func TestClear(t *testing.T) {

	queue := New(1, 2, 3, 4, 5)
	queue.Clear()

	assert.NotNil(t, queue)
	assert.True(t, queue.Empty())

}

func TestForEach(t *testing.T) {

	// Empty dequeue.
	queue := New[int]()
	sum := 0
	queue.ForEach(func(i int) { sum = sum + i })
	assert.Equal(t, 0, sum)

	// Dequeue with elements.
	queue = New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum = 0
	queue.ForEach(func(i int) { sum = sum + i })
	assert.Equal(t, 55, sum)

}

func TestString(t *testing.T) {

	assert.Equal(t, "[]", New[int]().String())
	assert.Equal(t, "[1]", New(1).String())
	assert.Equal(t, "[1 2]", New(1, 2).String())
}

func TestEquals(t *testing.T) {

	type equalsTest struct {
		a        *ListDequeue[int]
		b        *ListDequeue[int]
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
