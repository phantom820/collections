package sets

import (
	"testing"

	"github.com/phantom820/collections/sets/hashset"
	"github.com/phantom820/collections/types"
	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {

	a := hashset.New[types.Int]()
	b := hashset.New[types.Int]()
	c := hashset.New[types.Int]()

	// Case 1 : Union of empty sets.
	_ = Union[types.Int](a, b, c)
	assert.Equal(t, true, c.Equals(hashset.New[types.Int]()))

	// Case 2 : Union of sets with elements.
	a.Add(1, 2, 3)
	b.Add(3, 4, 5, 6)
	_ = Union[types.Int](a, b, c)
	assert.Equal(t, true, c.Equals(hashset.New[types.Int](1, 2, 3, 4, 5, 6)))

	// Case 3 : Passing a non empty result set.
	err := Union[types.Int](a, b, c)
	assert.Equal(t, errInvalidDestinationSet, err)

}

func TestIntersection(t *testing.T) {

	a := hashset.New[types.Int]()
	b := hashset.New[types.Int]()
	c := hashset.New[types.Int]()

	// Case 1 : Intersection of empty sets.
	_ = Union[types.Int](a, b, c)
	assert.Equal(t, true, c.Equals(hashset.New[types.Int]()))

	// Case 2 : Intersection of sets with elements.
	a.Add(1, 2, 3)
	b.Add(3, 4, 5, 6)
	_ = Intersection[types.Int](a, b, c)
	assert.Equal(t, true, c.Equals(hashset.New[types.Int](3)))

	// make a the larger set
	a.Add(4, 11, 12, 13, 14)
	c.Clear()
	_ = Intersection[types.Int](a, b, c)
	assert.Equal(t, true, c.Equals(hashset.New[types.Int](3, 4)))

	// Case 3 : Passing a non empty result set,
	err := Intersection[types.Int](a, b, c)
	assert.Equal(t, errInvalidDestinationSet, err)

}

func TestDifference(t *testing.T) {

	a := hashset.New[types.Int]()
	b := hashset.New[types.Int]()
	c := hashset.New[types.Int]()

	// Case 1 : Difference of empty sets.
	_ = Difference[types.Int](a, b, c)
	assert.Equal(t, true, c.Equals(hashset.New[types.Int]()))

	// Case 2 : Difference of sets with elements.
	a.Add(1, 2, 3)
	b.Add(3, 4, 5, 6)
	_ = Difference[types.Int](a, b, c)
	assert.Equal(t, true, c.Equals(hashset.New[types.Int](1, 2)))

	// Case 3 : Passing a non empty result set,
	err := Difference[types.Int](a, b, c)
	assert.Equal(t, errInvalidDestinationSet, err)

}
