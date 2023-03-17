package treemap

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types/optional"
	"github.com/phantom820/collections/types/pair"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	treeMap := New[string, int](func(k1, k2 string) bool { return k1 < k2 })
	assert.NotNil(t, treeMap)
	assert.True(t, treeMap.Empty())
	assert.Equal(t, 0, treeMap.Len())

}

func TestPut(t *testing.T) {

	type putTest struct {
		input          *TreeMap[string, int]
		action         func(*TreeMap[string, int])
		expectedKeys   []string
		expectedValues []int
	}

	lessThan := func(k1, k2 string) bool { return k1 < k2 }
	putTests := []putTest{
		{input: New[string, int](lessThan),
			action: func(hm *TreeMap[string, int]) {
				hm.Put("A", 1)
			},
			expectedKeys:   []string{"A"},
			expectedValues: []int{1},
		},
		{input: New[string, int](lessThan),
			action: func(hm *TreeMap[string, int]) {
				hm.Put("A", 1)
				hm.Put("A", 2)
			},
			expectedKeys:   []string{"A"},
			expectedValues: []int{2},
		},
		{input: New[string, int](lessThan),
			action: func(hm *TreeMap[string, int]) {
				hm.Put("B", 2)
				hm.Put("C", 3)
				hm.Put("C", 4)
				hm.Put("A", 1)
			},
			expectedKeys:   []string{"A", "B", "C"},
			expectedValues: []int{1, 2, 4},
		},
		{input: New[string, int](func(k1, k2 string) bool { return k1 > k2 }),
			action: func(hm *TreeMap[string, int]) {
				hm.Put("B", 2)
				hm.Put("C", 3)
				hm.Put("C", 4)
				hm.Put("A", 1)
			},
			expectedKeys:   []string{"C", "B", "A"},
			expectedValues: []int{4, 2, 1},
		},
	}

	for _, test := range putTests {
		test.action(test.input)
		keys, values := test.input.Keys(), test.input.Values()
		assert.Equal(t, test.expectedKeys, keys)
		assert.Equal(t, test.expectedValues, values)

	}
}

func TestPutIfAbsent(t *testing.T) {

	type putIfAbsentTest struct {
		input          *TreeMap[string, int]
		action         func(*TreeMap[string, int])
		expectedKeys   []string
		expectedValues []int
	}

	lessThan := func(k1, k2 string) bool { return k1 < k2 }

	putIfAbsentTests := []putIfAbsentTest{
		{input: New[string, int](lessThan),
			action: func(hm *TreeMap[string, int]) {
				hm.PutIfAbsent("A", 1)
			},
			expectedKeys:   []string{"A"},
			expectedValues: []int{1},
		},
		{input: New[string, int](lessThan),
			action: func(hm *TreeMap[string, int]) {
				hm.PutIfAbsent("B", 3)
				hm.PutIfAbsent("A", 1)
				hm.PutIfAbsent("A", 2)

			},
			expectedKeys:   []string{"A", "B"},
			expectedValues: []int{1, 3},
		},
	}

	for _, test := range putIfAbsentTests {
		test.action(test.input)
		keys, values := test.input.Keys(), test.input.Values()
		assert.Equal(t, test.expectedKeys, keys)
		assert.Equal(t, test.expectedValues, values)

	}
}

func TestGet(t *testing.T) {

	type getTest struct {
		input    *TreeMap[string, int]
		key      string
		action   func(*TreeMap[string, int])
		expected optional.Optional[int]
	}

	lessThan := func(k1, k2 string) bool { return k1 < k2 }

	getTests := []getTest{
		{input: New[string, int](lessThan),
			key:      "A",
			action:   func(tm *TreeMap[string, int]) {},
			expected: optional.Empty[int](),
		},
		{input: New[string, int](lessThan),
			key: "B",
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
				tm.Put("B", 2)
			},
			expected: optional.Of(2),
		},
		{input: New[string, int](lessThan),
			key: "C",
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
				tm.Put("B", 2)
			},
			expected: optional.Empty[int](),
		},
	}

	for _, test := range getTests {
		test.action(test.input)
		assert.Equal(t, test.expected, test.input.Get(test.key))
	}
}

func TestGetIf(t *testing.T) {

	type getIfTest struct {
		input    *TreeMap[string, int]
		f        func(string) bool
		expected []int
	}

	lessThan := func(k1, k2 string) bool { return k1 < k2 }

	tm := New[string, int](lessThan)
	tm.Put("B", 2)
	tm.Put("C", 3)
	tm.Put("D", 4)
	tm.Put("A", 1)

	getIfTests := []getIfTest{
		{input: tm,
			f:        func(s string) bool { return s == "" },
			expected: []int{},
		},
		{input: tm,
			f:        func(s string) bool { return s == "A" || s == "B" },
			expected: []int{1, 2},
		},
		{input: tm,
			f:        func(s string) bool { return false },
			expected: []int{},
		},
	}

	for _, test := range getIfTests {
		assert.Equal(t, test.expected, test.input.GetIf(test.f))
	}
}

func TestRemove(t *testing.T) {

	type removeTest struct {
		input          *TreeMap[string, int]
		action         func(*TreeMap[string, int])
		key            string
		expectedKeys   []string
		expectedValues []int
	}

	lessThan := func(k1, k2 string) bool { return k1 < k2 }

	removeTests := []removeTest{
		{
			input:          New[string, int](lessThan),
			key:            "A",
			action:         func(tm *TreeMap[string, int]) {},
			expectedKeys:   []string{},
			expectedValues: []int{},
		},
		{
			input: New[string, int](lessThan),
			key:   "A",
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
			},
			expectedKeys:   []string{},
			expectedValues: []int{},
		},
		{
			input: New[string, int](lessThan),
			key:   "A",
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
				tm.Put("B", 2)
				tm.Put("C", 3)
				tm.Put("D", 4)
				tm.Put("E", 5)
			},
			expectedKeys:   []string{"B", "C", "D", "E"},
			expectedValues: []int{2, 3, 4, 5},
		},
		{
			input: New[string, int](lessThan),
			key:   "B",
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
				tm.Put("B", 2)
				tm.Put("C", 3)
				tm.Put("D", 4)
			},
			expectedKeys:   []string{"A", "C", "D"},
			expectedValues: []int{1, 3, 4},
		},
		{
			input: New[string, int](lessThan),
			key:   "D",
			action: func(tm *TreeMap[string, int]) {
				tm.Put("B", 2)
				tm.Put("C", 3)
				tm.Put("D", 4)
				tm.Put("A", 1)
			},
			expectedKeys:   []string{"A", "B", "C"},
			expectedValues: []int{1, 2, 3},
		},
	}

	for _, test := range removeTests {
		test.action(test.input)
		test.input.Remove(test.key)
		keys, values := test.input.Keys(), test.input.Values()
		assert.Equal(t, test.expectedKeys, keys)
		assert.Equal(t, test.expectedValues, values)
	}
}

func TestRemoveIf(t *testing.T) {

	type removeIfTest struct {
		input          *TreeMap[string, int]
		action         func(*TreeMap[string, int])
		f              func(string) bool
		expectedKeys   []string
		expectedValues []int
	}

	lessThan := func(k1, k2 string) bool { return k1 < k2 }
	removeIfTests := []removeIfTest{
		{
			input:          New[string, int](lessThan),
			f:              func(s string) bool { return s == "" },
			action:         func(tm *TreeMap[string, int]) {},
			expectedKeys:   []string{},
			expectedValues: []int{},
		},
		{
			input: New[string, int](lessThan),
			f:     func(s string) bool { return s == "A" },
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
			},
			expectedKeys:   []string{},
			expectedValues: []int{},
		},
		{
			input: New[string, int](lessThan),
			f:     func(s string) bool { return s == "A" || s == "C" },
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
				tm.Put("B", 2)
				tm.Put("C", 3)
				tm.Put("D", 4)
			},
			expectedKeys:   []string{"B", "D"},
			expectedValues: []int{2, 4},
		},
		{
			input: New[string, int](lessThan),
			f:     func(s string) bool { return true },
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
				tm.Put("B", 2)
				tm.Put("C", 3)
				tm.Put("D", 4)
				tm.Put("E", 4)
			},
			expectedKeys:   []string{},
			expectedValues: []int{},
		},
		{
			input: New[string, int](lessThan),
			f:     func(s string) bool { return s == "A" || s == "B" || s == "C" || s == "D" },
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
				tm.Put("B", 2)
				tm.Put("C", 3)
				tm.Put("D", 4)
				tm.Put("E", 5)
			},
			expectedKeys:   []string{"E"},
			expectedValues: []int{5},
		},
		{
			input: New[string, int](lessThan),
			f:     func(s string) bool { return s == "B" || s == "C" || s == "D" },
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
				tm.Put("B", 2)
				tm.Put("C", 3)
				tm.Put("D", 4)
			},
			expectedKeys:   []string{"A"},
			expectedValues: []int{1},
		},
	}

	for _, test := range removeIfTests {
		test.action(test.input)
		test.input.RemoveIf(test.f)
		keys, values := test.input.Keys(), test.input.Values()
		assert.Equal(t, test.expectedKeys, keys)
		assert.Equal(t, test.expectedValues, values)
	}
}

func TestContainsKey(t *testing.T) {

	type containsKeyTest struct {
		input    *TreeMap[string, int]
		action   func(*TreeMap[string, int])
		key      string
		expected bool
	}

	lessThan := func(k1, k2 string) bool { return k1 < k2 }
	containsKeyTests := []containsKeyTest{
		{
			input:    New[string, int](lessThan),
			key:      "A",
			action:   func(tm *TreeMap[string, int]) {},
			expected: false,
		},
		{
			input: New[string, int](lessThan),
			key:   "A",
			action: func(tm *TreeMap[string, int]) {
				tm.Put("C", 1)
				tm.Put("D", 1)
			},
			expected: false,
		},
		{
			input: New[string, int](lessThan),
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 1)
			},
			key:      "A",
			expected: true,
		},
	}

	for _, test := range containsKeyTests {
		test.action(test.input)
		assert.Equal(t, test.expected, test.input.ContainsKey(test.key))
	}
}

func TestContainsValue(t *testing.T) {

	type containsValueTest struct {
		input    *TreeMap[string, int]
		action   func(*TreeMap[string, int])
		value    int
		expected bool
	}

	lessThan := func(k1, k2 string) bool { return k1 < k2 }
	containsValueTests := []containsValueTest{
		{
			input:    New[string, int](lessThan),
			action:   func(tm *TreeMap[string, int]) {},
			value:    0,
			expected: false,
		},
		{
			input: New[string, int](lessThan),
			action: func(tm *TreeMap[string, int]) {
				tm.Put("A", 22)
				tm.Put("B", 45)
				tm.Put("C", 33)
			},
			value:    2,
			expected: false,
		},
		{
			input: New[string, int](lessThan),
			action: func(tm *TreeMap[string, int]) {
				tm.Put("B", 45)
				tm.Put("C", 33)
				tm.Put("A", 1)
			},
			value:    1,
			expected: true,
		},
	}
	equals := func(v1, v2 int) bool { return v1 == v2 }
	for _, test := range containsValueTests {
		test.action(test.input)
		assert.Equal(t, test.expected, test.input.ContainsValue(test.value, equals))
	}
}

func TestClear(t *testing.T) {

	tm := New[string, int](func(k1, k2 string) bool { return k1 < k2 })
	tm.Put("A", 1)
	tm.Put("B", 2)

	tm.Clear()
	assert.True(t, tm.Empty())

}

func TestKeys(t *testing.T) {

	tm := New[string, int](func(k1, k2 string) bool { return k1 < k2 })
	tm.Put("A", 1)
	tm.Put("B", 2)
	tm.Put("C", 3)
	tm.Put("D", 4)

	assert.Equal(t, []string{"A", "B", "C", "D"}, tm.Keys())
}

func TestValues(t *testing.T) {

	tm := New[string, int](func(k1, k2 string) bool { return k1 < k2 })
	tm.Put("A", 1)
	tm.Put("B", 2)
	tm.Put("C", 3)
	tm.Put("D", 4)

	assert.Equal(t, []int{1, 2, 3, 4}, tm.Values())
}

func TestForEach(t *testing.T) {

	m := make(map[string]int)
	tm := New[string, int](func(k1, k2 string) bool { return k1 < k2 })
	tm.Put("A", 1)
	tm.Put("B", 2)
	tm.Put("C", 3)
	tm.Put("D", 4)

	tm.ForEach(func(s string, i int) {
		m[s] = i
	})

	assert.Equal(t, map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}, m)
}

func TestIterator(t *testing.T) {

	iterate := func(it iterator.Iterator[pair.Pair[string, int]]) []pair.Pair[string, int] {
		entries := make([]pair.Pair[string, int], 0)
		for it.HasNext() {
			entries = append(entries, it.Next())
		}
		return entries
	}

	h := New[string, int](func(k1, k2 string) bool { return k1 < k2 })
	assert.Equal(t, []pair.Pair[string, int]{}, iterate(h.Iterator()))

	h = New[string, int](func(k1, k2 string) bool { return k1 < k2 })
	h.Put("A", 1)
	h.Put("B", 2)
	h.Put("C", 3)

	assert.Equal(t, []pair.Pair[string, int]{
		pair.New("A", 1), pair.New("B", 2), pair.New("C", 3)}, iterate(h.Iterator()))

	h = New[string, int](func(k1, k2 string) bool { return k1 < k2 })
	h.Put("A", 1)
	h.Put("B", 2)
	h.Put("C", 3)
	it := h.Iterator()
	h.Put("F", 23)

	assert.Equal(t, []pair.Pair[string, int]{
		pair.New("A", 1), pair.New("B", 2), pair.New("C", 3), pair.New("F", 23)}, iterate(it))

}

func TestString(t *testing.T) {

	assert.Equal(t, "{}", fmt.Sprint(New[string, string](func(k1, k2 string) bool { return k1 < k2 })))
	m := New[string, int](func(k1, k2 string) bool { return k1 < k2 })
	m.Put("A", 1)
	m.Put("B", 2)
	m.Put("C", 3)
	assert.Equal(t, "{A=1, B=2, C=3}", m.String())

}
