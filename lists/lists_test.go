package lists

import (
	"testing"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	newTests := []struct {
		input collections.List[int]
	}{
		{
			input: forwardlist.New[int](),
		},
		{
			input: linkedlist.New[int](),
		},
		{
			input: vector.New[int](),
		},
	}

	for _, test := range newTests {
		assert.NotNil(t, test.input)
		assert.True(t, test.input.Empty())
	}

}

func TestAdd(t *testing.T) {

	addTests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range addTests {
		lists := []collections.List[int]{forwardlist.New[int](), linkedlist.New[int](), vector.New[int]()}
		for _, list := range lists {
			for _, element := range test.input {
				list.Add(element)
			}
			assert.Equal(t, test.expected, list.ToSlice())
		}
	}

}
