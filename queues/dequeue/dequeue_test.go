package dequeue

import (
	"testing"

	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
	"github.com/phantom820/collections/types/optional"
	"github.com/stretchr/testify/assert"
)

func TestNewListDequeue(t *testing.T) {

	deq := NewListDequeue[int]()

	assert.NotNil(t, deq)
	assert.NotNil(t, deq.list)
	_, ok := deq.list.(*linkedlist.LinkedList[int])
	assert.True(t, ok)
	assert.True(t, deq.Empty())
}

func TestNewVectorDequeue(t *testing.T) {

	deq := NewVectorDequeue[int]()

	assert.NotNil(t, deq)
	assert.NotNil(t, deq.list)
	_, ok := deq.list.(*vector.Vector[int])
	assert.True(t, ok)
	assert.True(t, deq.Empty())

}

func TestOf(t *testing.T) {

	dequeTests := []struct {
		deq      Dequeue[int]
		expected []int
	}{
		{
			deq:      VectorDequeueOf[int](),
			expected: []int{},
		},
		{
			deq:      ListDequeueOf[int](),
			expected: []int{},
		},
		{
			deq:      VectorDequeueOf[int](1, 2),
			expected: []int{1, 2},
		},
		{
			deq:      ListDequeueOf[int](1, 2),
			expected: []int{1, 2},
		},
	}

	for _, test := range dequeTests {
		assert.Equal(t, test.expected, test.deq.ToSlice())
	}

}

func TestPeekFirst(t *testing.T) {

	pekkFirstTests := []struct {
		input    Dequeue[int]
		expected optional.Optional[int]
	}{
		{
			input:    ListDequeueOf[int](),
			expected: optional.Empty[int](),
		},
		{
			input:    VectorDequeueOf[int](),
			expected: optional.Empty[int](),
		},
		{
			input:    ListDequeueOf[int](1),
			expected: optional.Of(1),
		},
		{
			input:    VectorDequeueOf[int](1),
			expected: optional.Of(1),
		},
		{
			input:    ListDequeueOf[int](1, 2, 3),
			expected: optional.Of(1),
		},
		{
			input:    VectorDequeueOf[int](1, 2, 3),
			expected: optional.Of(1),
		},
	}

	for _, test := range pekkFirstTests {
		assert.Equal(t, test.expected, test.input.PeekFirst())
	}
}

func TestPeekLast(t *testing.T) {

	peekFirstTests := []struct {
		input    Dequeue[int]
		expected optional.Optional[int]
	}{
		{
			input:    ListDequeueOf[int](),
			expected: optional.Empty[int](),
		},
		{
			input:    VectorDequeueOf[int](),
			expected: optional.Empty[int](),
		},
		{
			input:    ListDequeueOf[int](1),
			expected: optional.Of(1),
		},
		{
			input:    VectorDequeueOf[int](1),
			expected: optional.Of(1),
		},
		{
			input:    ListDequeueOf[int](1, 2, 3),
			expected: optional.Of(3),
		},
		{
			input:    VectorDequeueOf[int](1, 2, 3),
			expected: optional.Of(3),
		},
	}

	for _, test := range peekFirstTests {
		assert.Equal(t, test.expected, test.input.PeekLast())
	}
}

func TestRemoveFirst(t *testing.T) {

	removeFirstTests := []struct {
		input    Dequeue[int]
		expected []int
	}{
		{
			input:    ListDequeueOf[int](),
			expected: []int{},
		},
		{
			input:    VectorDequeueOf[int](),
			expected: []int{},
		},
		{
			input:    ListDequeueOf[int](1),
			expected: []int{1},
		},
		{
			input:    VectorDequeueOf[int](1),
			expected: []int{1},
		},
		{
			input:    ListDequeueOf[int](1, 2, 3),
			expected: []int{1, 2, 3},
		},
		{
			input:    VectorDequeueOf[int](1, 2, 3),
			expected: []int{1, 2, 3},
		},
	}

	for _, test := range removeFirstTests {
		for _, value := range test.expected {
			front := test.input.RemoveFirst()
			assert.Equal(t, value, front.Value())
		}
	}
}

func TestRemoveLast(t *testing.T) {

	removeLastTests := []struct {
		input    Dequeue[int]
		expected []int
	}{
		{
			input:    ListDequeueOf[int](),
			expected: []int{},
		},
		{
			input:    VectorDequeueOf[int](),
			expected: []int{},
		},
		{
			input:    ListDequeueOf[int](1),
			expected: []int{1},
		},
		{
			input:    VectorDequeueOf[int](1),
			expected: []int{1},
		},
		{
			input:    ListDequeueOf[int](1, 2, 3),
			expected: []int{3, 2, 1},
		},
		{
			input:    VectorDequeueOf[int](1, 2, 3),
			expected: []int{3, 2, 1},
		},
	}

	for _, test := range removeLastTests {
		for _, value := range test.expected {
			front := test.input.RemoveLast()
			assert.Equal(t, value, front.Value())
		}
	}
}

func TestAddFirst(t *testing.T) {

	addFirstTests := []struct {
		input    Dequeue[int]
		elements []int
		expected []int
	}{
		{
			input:    ListDequeueOf[int](),
			elements: []int{},
			expected: []int{},
		},
		{
			input:    VectorDequeueOf[int](),
			elements: []int{},
			expected: []int{},
		},
		{
			input:    ListDequeueOf[int](),
			elements: []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			input:    VectorDequeueOf[int](),
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
		input    Dequeue[int]
		elements []int
		expected []int
	}{
		{
			input:    ListDequeueOf[int](),
			elements: []int{},
			expected: []int{},
		},
		{
			input:    VectorDequeueOf[int](),
			elements: []int{},
			expected: []int{},
		},
		{
			input:    ListDequeueOf[int](),
			elements: []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    VectorDequeueOf[int](),
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