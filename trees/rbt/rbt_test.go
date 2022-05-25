package rbt

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/trees"
	"github.com/phantom820/collections/types"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {

	tree := New[types.Int, int]()

	// Note these tests check the state of the tree after an insertion using an in order traversal. The state of the tree should
	// correspond to the results obtained by doing the operations by hand.

	assert.Equal(t, true, tree.Empty())
	assert.Equal(t, "", fmt.Sprint(tree))
	tree.Insert(20, 1)
	assert.Equal(t, "20(B)", fmt.Sprint(tree))
	tree.Insert(30, 1)
	assert.Equal(t, "20(B) 30(R)", fmt.Sprint(tree))
	tree.Insert(40, 1)
	assert.Equal(t, "20(R) 30(B) 40(R)", fmt.Sprint(tree))
	tree.Insert(10, 1)
	assert.Equal(t, "10(R) 20(B) 30(B) 40(B)", fmt.Sprint(tree))
	tree.Insert(15, 1)
	assert.Equal(t, "10(R) 15(B) 20(R) 30(B) 40(B)", fmt.Sprint(tree))
	tree.Insert(25, 2)
	assert.Equal(t, "10(B) 15(R) 20(B) 25(R) 30(B) 40(B)", fmt.Sprint(tree))
	tree.Insert(24, 11)
	assert.Equal(t, "10(B) 15(R) 20(R) 24(B) 25(R) 30(B) 40(B)", fmt.Sprint(tree))
	assert.Equal(t, 7, tree.Len())

}

func TestDelete(t *testing.T) {

	tree := New[types.Int, int]()

	// Note these tests check the state of the tree after a deletion using an in order traversal. The state of the tree should
	// correspond to the results obtained by doing the operations by hand.

	tree.Insert(20, 10)
	tree.Insert(30, 20)
	tree.Insert(40, 0)
	tree.Insert(10, 30)
	tree.Insert(15, 1)
	tree.Insert(25, 2)
	tree.Insert(24, 11)
	tree.Insert(21, 11)
	tree.Insert(17, 11)
	tree.Insert(41, 11)
	tree.Insert(39, 11)

	assert.Equal(t, "10(B) 15(R) 17(R) 20(B) 21(R) 24(B) 25(B) 30(R) 39(R) 40(B) 41(R)", fmt.Sprint(tree))
	tree.Delete(10)
	assert.Equal(t, "15(B) 17(R) 20(R) 21(B) 24(B) 25(B) 30(R) 39(R) 40(B) 41(R)", fmt.Sprint(tree))
	tree.Delete(15)
	assert.Equal(t, "17(B) 20(R) 21(B) 24(B) 25(B) 30(R) 39(R) 40(B) 41(R)", fmt.Sprint(tree))
	tree.Delete(30)
	assert.Equal(t, "17(B) 20(R) 21(B) 24(B) 25(B) 39(R) 40(B) 41(R)", fmt.Sprint(tree))
	tree.Delete(24)
	assert.Equal(t, "17(B) 20(R) 21(B) 25(B) 39(B) 40(R) 41(B)", fmt.Sprint(tree))
	tree.Delete(25)
	assert.Equal(t, "17(B) 20(R) 21(B) 39(B) 40(B) 41(R)", fmt.Sprint(tree))
	tree.Delete(39)
	assert.Equal(t, "17(B) 20(R) 21(B) 40(B) 41(B)", fmt.Sprint(tree))
	tree.Delete(41)
	assert.Equal(t, "17(B) 20(B) 21(R) 40(B)", fmt.Sprint(tree))
	tree.Delete(40)
	assert.Equal(t, "17(B) 20(B) 21(B)", fmt.Sprint(tree))
	tree.Insert(14, 11)
	tree.Delete(21)
	assert.Equal(t, "14(B) 17(B) 20(B)", fmt.Sprint(tree))
	tree.Insert(18, 11)
	tree.Insert(23, 11)
	tree.Insert(21, 11)
	tree.Delete(17)
	assert.Equal(t, "14(B) 18(B) 20(B) 21(R) 23(B)", fmt.Sprint(tree))

	tree.Clear()
	assert.Equal(t, true, tree.Empty())
	_, ok := tree.Delete(1)
	assert.Equal(t, false, ok)

	tree.Insert(50, 1)
	tree.Insert(80, 1)
	tree.Insert(90, 1)
	tree.Insert(100, 1)
	tree.Insert(120, 1)
	tree.Insert(140, 1)
	tree.Insert(150, 1)
	tree.Insert(110, 1)
	tree.Insert(122, 1)
	tree.Delete(110)
	tree.Delete(150)
	assert.Equal(t, "50(B) 80(R) 90(B) 100(B) 120(B) 122(R) 140(B)", fmt.Sprint(tree))

}

func TestUpdate(t *testing.T) {

	tree := New[types.Int, int]()

	tree.Insert(10, 10)
	tree.Insert(20, 20)

	_, ok := tree.Update(-1, 20)
	assert.Equal(t, false, ok)
	value, _ := tree.Update(10, -20)
	assert.Equal(t, 10, value)
	value, _ = tree.Get(10)
	assert.Equal(t, -20, value)

}

func TestSearch(t *testing.T) {

	tree := New[types.Int, int]()

	for i := -10; i < 12; i++ {
		tree.Insert(types.Int(i), i)
	}

	assert.Equal(t, true, tree.Search(0))
	assert.Equal(t, true, tree.Search(11))
	assert.Equal(t, false, tree.Search(-11))
	_, b := tree.Get(-22)
	assert.Equal(t, false, b)

	for i := -10; i < 12; i++ {
		tree.Delete(types.Int(i))
	}

}

func TestInOrderTraversal(t *testing.T) {

	tree := New[types.Int, int]()

	keys := make([]types.Int, 0)
	assert.Equal(t, keys, tree.InOrderTraversal())

	tree.Insert(22, 1)
	tree.Insert(13, 2)
	tree.Insert(3, 2)

	keys = []types.Int{3, 13, 22}
	assert.Equal(t, keys, tree.InOrderTraversal())

}

func TestKeys(t *testing.T) {

	tree := New[types.Int, int]()

	keys := make([]types.Int, 0)
	assert.ElementsMatch(t, keys, tree.Keys())
	tree.Insert(1, 1)
	tree.Insert(2, 2)
	tree.Insert(3, 3)

	keys = []types.Int{1, 2, 3}
	assert.ElementsMatch(t, keys, tree.Keys())

}

func TestNodes(t *testing.T) {

	tree := New[types.Int, int]()

	nodes := []trees.Node[types.Int, int]{}
	assert.ElementsMatch(t, nodes, tree.Nodes())

	tree.Insert(1, 2)
	tree.Insert(-10, 1)
	tree.Insert(11, 0)

	nodes = []trees.Node[types.Int, int]{{-10, 1}, {1, 2}, {11, 0}}
	assert.ElementsMatch(t, nodes, tree.Nodes())

}
func TestValues(t *testing.T) {

	tree := New[types.Int, int]()

	values := make([]int, 0)
	assert.ElementsMatch(t, values, tree.Values())
	tree.Insert(1, 1)
	tree.Insert(2, 2)
	tree.Insert(3, 3)

	values = []int{1, 2, 3}
	assert.ElementsMatch(t, values, tree.Values())
}
