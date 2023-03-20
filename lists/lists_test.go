package lists

import (
	"math/rand"
	"testing"
	"time"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/linkedlist"
	"github.com/phantom820/collections/lists/vector"
	"github.com/stretchr/testify/assert"
)

func data(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rand.Intn(n)
	}
	return data
}

func shuffle(data []int) {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
}

func TestAdd(t *testing.T) {

	lists := []collections.List[int]{
		vector.New[int](),
		forwardlist.New[int](),
		linkedlist.New[int](),
	}

	data := data(100)
	for _, list := range lists {
		for _, e := range data {
			list.Add(e)
		}
		assert.Equal(t, data, list.ToSlice())
	}

}

func TestRemove(t *testing.T) {

	data := data(100)

	lists := []collections.List[int]{
		vector.New(data...),
		forwardlist.New(data...),
		linkedlist.New(data...),
	}

	shuffle(data)

	for _, list := range lists {
		for _, e := range data {
			list.Remove(e)
		}
		assert.Equal(t, []int{}, list.ToSlice())
	}

}
