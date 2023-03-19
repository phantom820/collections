package linkedhashmap

import (
	"fmt"
	"testing"

	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/types/optional"
	"github.com/phantom820/collections/types/pair"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	linkedHashMap := New[string, int]()
	assert.NotNil(t, linkedHashMap)
	assert.True(t, linkedHashMap.Empty())
	assert.Equal(t, 0, linkedHashMap.Len())
	assert.Nil(t, linkedHashMap.head)
	assert.Nil(t, linkedHashMap.tail)

}

func TestPut(t *testing.T) {

	type putTest struct {
		input          *LinkedHashMap[string, int]
		action         func(*LinkedHashMap[string, int])
		expectedKeys   []string
		expectedValues []int
	}

	putTests := []putTest{
		{input: New[string, int](),
			action: func(hm *LinkedHashMap[string, int]) {
				hm.Put("A", 1)
			},
			expectedKeys:   []string{"A"},
			expectedValues: []int{1},
		},
		{input: New[string, int](),
			action: func(hm *LinkedHashMap[string, int]) {
				hm.Put("A", 1)
				hm.Put("A", 2)
			},
			expectedKeys:   []string{"A"},
			expectedValues: []int{2},
		},
		{input: New[string, int](),
			action: func(hm *LinkedHashMap[string, int]) {
				hm.Put("A", 1)
				hm.Put("B", 2)
				hm.Put("C", 3)
				hm.Put("C", 4)
			},
			expectedKeys:   []string{"A", "B", "C"},
			expectedValues: []int{1, 2, 4},
		},
	}

	for _, test := range putTests {
		test.action(test.input)
		keys, values := collect(test.input)
		assert.Equal(t, test.expectedKeys, keys)
		assert.Equal(t, test.expectedValues, values)

	}
}

func TestPutIfAbsent(t *testing.T) {

	type putIfAbsentTest struct {
		input          *LinkedHashMap[string, int]
		action         func(*LinkedHashMap[string, int])
		expectedKeys   []string
		expectedValues []int
	}

	putIfAbsentTests := []putIfAbsentTest{
		{input: New[string, int](),
			action: func(hm *LinkedHashMap[string, int]) {
				hm.PutIfAbsent("A", 1)
			},
			expectedKeys:   []string{"A"},
			expectedValues: []int{1},
		},
		{input: New[string, int](),
			action: func(hm *LinkedHashMap[string, int]) {
				hm.PutIfAbsent("A", 1)
				hm.PutIfAbsent("A", 2)
				hm.PutIfAbsent("B", 3)

			},
			expectedKeys:   []string{"A", "B"},
			expectedValues: []int{1, 3},
		},
	}

	for _, test := range putIfAbsentTests {
		test.action(test.input)
		keys, values := collect(test.input)
		assert.Equal(t, test.expectedKeys, keys)
		assert.Equal(t, test.expectedValues, values)

	}
}

func TestGet(t *testing.T) {

	type getTest struct {
		input    *LinkedHashMap[string, int]
		key      string
		action   func(*LinkedHashMap[string, int])
		expected optional.Optional[int]
	}

	getTests := []getTest{
		{input: New[string, int](),
			key:      "A",
			action:   func(lhm *LinkedHashMap[string, int]) {},
			expected: optional.Empty[int](),
		},
		{input: New[string, int](),
			key: "B",
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
				lhm.Put("B", 2)
			},
			expected: optional.Of(2),
		},
		{input: New[string, int](),
			key: "C",
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
				lhm.Put("B", 2)
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
		input    *LinkedHashMap[string, int]
		f        func(string) bool
		expected []int
	}

	lhm := New[string, int]()
	lhm.Put("A", 1)
	lhm.Put("B", 2)
	lhm.Put("C", 3)
	lhm.Put("D", 4)

	getIfTests := []getIfTest{
		{input: lhm,
			f:        func(s string) bool { return s == "" },
			expected: []int{},
		},
		{input: lhm,
			f:        func(s string) bool { return s == "A" || s == "B" },
			expected: []int{1, 2},
		},
		{input: lhm,
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
		input          *LinkedHashMap[string, int]
		action         func(*LinkedHashMap[string, int])
		key            string
		expectedKeys   []string
		expectedValues []int
	}

	removeTests := []removeTest{
		{
			input:          New[string, int](),
			key:            "A",
			action:         func(lhm *LinkedHashMap[string, int]) {},
			expectedKeys:   []string{},
			expectedValues: []int{},
		},
		{
			input: New[string, int](),
			key:   "A",
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
			},
			expectedKeys:   []string{},
			expectedValues: []int{},
		},
		{
			input: New[string, int](),
			key:   "A",
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
				lhm.Put("B", 2)
				lhm.Put("C", 3)
				lhm.Put("D", 4)
				lhm.Put("E", 5)
			},
			expectedKeys:   []string{"B", "C", "D", "E"},
			expectedValues: []int{2, 3, 4, 5},
		},
		{
			input: New[string, int](),
			key:   "B",
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
				lhm.Put("B", 2)
				lhm.Put("C", 3)
				lhm.Put("D", 4)
			},
			expectedKeys:   []string{"A", "C", "D"},
			expectedValues: []int{1, 3, 4},
		},
		{
			input: New[string, int](),
			key:   "D",
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
				lhm.Put("B", 2)
				lhm.Put("C", 3)
				lhm.Put("D", 4)
			},
			expectedKeys:   []string{"A", "B", "C"},
			expectedValues: []int{1, 2, 3},
		},
	}

	for _, test := range removeTests {
		test.action(test.input)
		test.input.Remove(test.key)
		keys, values := collect(test.input)
		assert.Equal(t, test.expectedKeys, keys)
		assert.Equal(t, test.expectedValues, values)
	}
}

func TestRemoveIf(t *testing.T) {

	type removeIfTest struct {
		input          *LinkedHashMap[string, int]
		action         func(*LinkedHashMap[string, int])
		f              func(string) bool
		expectedKeys   []string
		expectedValues []int
	}

	removeIfTests := []removeIfTest{
		{
			input:          New[string, int](),
			f:              func(s string) bool { return s == "" },
			action:         func(lhm *LinkedHashMap[string, int]) {},
			expectedKeys:   []string{},
			expectedValues: []int{},
		},
		{
			input: New[string, int](),
			f:     func(s string) bool { return s == "A" },
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
			},
			expectedKeys:   []string{},
			expectedValues: []int{},
		},
		{
			input: New[string, int](),
			f:     func(s string) bool { return s == "A" || s == "C" },
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
				lhm.Put("B", 2)
				lhm.Put("C", 3)
				lhm.Put("D", 4)
			},
			expectedKeys:   []string{"B", "D"},
			expectedValues: []int{2, 4},
		},
		{
			input: New[string, int](),
			f:     func(s string) bool { return true },
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
				lhm.Put("B", 2)
				lhm.Put("C", 3)
				lhm.Put("D", 4)
				lhm.Put("E", 4)
			},
			expectedKeys:   []string{},
			expectedValues: []int{},
		},
		{
			input: New[string, int](),
			f:     func(s string) bool { return s == "A" || s == "B" || s == "C" || s == "D" },
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
				lhm.Put("B", 2)
				lhm.Put("C", 3)
				lhm.Put("D", 4)
				lhm.Put("E", 5)
			},
			expectedKeys:   []string{"E"},
			expectedValues: []int{5},
		},
		{
			input: New[string, int](),
			f:     func(s string) bool { return s == "B" || s == "C" || s == "D" },
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
				lhm.Put("B", 2)
				lhm.Put("C", 3)
				lhm.Put("D", 4)
			},
			expectedKeys:   []string{"A"},
			expectedValues: []int{1},
		},
	}

	for _, test := range removeIfTests {
		test.action(test.input)
		test.input.RemoveIf(test.f)
		keys, values := collect(test.input)
		assert.Equal(t, test.expectedKeys, keys)
		assert.Equal(t, test.expectedValues, values)
	}
}

func TestContainsKey(t *testing.T) {

	type containsKeyTest struct {
		input    *LinkedHashMap[string, int]
		action   func(*LinkedHashMap[string, int])
		key      string
		expected bool
	}

	containsKeyTests := []containsKeyTest{
		{
			input:    New[string, int](),
			key:      "A",
			action:   func(lhm *LinkedHashMap[string, int]) {},
			expected: false,
		},
		{
			input: New[string, int](),
			key:   "A",
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("C", 1)
				lhm.Put("D", 1)
			},
			expected: false,
		},
		{
			input: New[string, int](),
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 1)
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
		input    *LinkedHashMap[string, int]
		action   func(*LinkedHashMap[string, int])
		value    int
		expected bool
	}

	containsValueTests := []containsValueTest{
		{
			input:    New[string, int](),
			action:   func(lhm *LinkedHashMap[string, int]) {},
			value:    0,
			expected: false,
		},
		{
			input: New[string, int](),
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("A", 22)
				lhm.Put("B", 45)
				lhm.Put("C", 33)
			},
			value:    2,
			expected: false,
		},
		{
			input: New[string, int](),
			action: func(lhm *LinkedHashMap[string, int]) {
				lhm.Put("B", 45)
				lhm.Put("C", 33)
				lhm.Put("A", 1)
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

	lhm := New[string, int]()
	lhm.Put("A", 1)
	lhm.Put("B", 2)

	lhm.Clear()
	assert.Nil(t, lhm.head)
	assert.Nil(t, lhm.tail)
	assert.True(t, lhm.Empty())

}

func TestKeys(t *testing.T) {

	lhm := New[string, int]()
	lhm.Put("A", 1)
	lhm.Put("B", 2)
	lhm.Put("C", 3)
	lhm.Put("D", 4)

	assert.Equal(t, []string{"A", "B", "C", "D"}, lhm.Keys())
}

func TestValues(t *testing.T) {

	lhm := New[string, int]()
	lhm.Put("A", 1)
	lhm.Put("B", 2)
	lhm.Put("C", 3)
	lhm.Put("D", 4)

	assert.Equal(t, []int{1, 2, 3, 4}, lhm.Values())
}

func TestForEach(t *testing.T) {

	m := make(map[string]int)
	lhm := New[string, int]()
	lhm.Put("A", 1)
	lhm.Put("B", 2)
	lhm.Put("C", 3)
	lhm.Put("D", 4)

	lhm.ForEach(func(s string, i int) {
		m[s] = i
	})

	assert.Equal(t, map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}, m)
}

func collect(linkedHashMap *LinkedHashMap[string, int]) ([]string, []int) {
	keys := make([]string, 0)
	values := make([]int, 0)
	for curr := linkedHashMap.head; curr != nil; curr = curr.next {
		keys = append(keys, curr.key)
		values = append(values, curr.value)
	}
	return keys, values
}

func TestIterator(t *testing.T) {

	iterate := func(it iterator.Iterator[pair.Pair[string, int]]) []pair.Pair[string, int] {
		entries := make([]pair.Pair[string, int], 0)
		for it.HasNext() {
			entries = append(entries, it.Next())
		}
		return entries
	}

	h := New[string, int]()
	assert.Equal(t, []pair.Pair[string, int]{}, iterate(h.Iterator()))

	h = New[string, int]()
	h.Put("A", 1)
	h.Put("B", 2)
	h.Put("C", 3)

	assert.Equal(t, []pair.Pair[string, int]{
		pair.Of("A", 1), pair.Of("B", 2), pair.Of("C", 3)}, iterate(h.Iterator()))

	h = New[string, int]()
	h.Put("A", 1)
	h.Put("B", 2)
	h.Put("C", 3)
	it := h.Iterator()
	h.Put("F", 23)

	assert.Equal(t, []pair.Pair[string, int]{
		pair.Of("A", 1), pair.Of("B", 2), pair.Of("C", 3), pair.Of("F", 23)}, iterate(it))

}

func TestString(t *testing.T) {

	assert.Equal(t, "{}", fmt.Sprint(New[string, string]()))
	m := New[string, int]()
	m.Put("A", 1)
	m.Put("B", 2)
	m.Put("C", 3)
	assert.Equal(t, "{A=1, B=2, C=3}", m.String())

}

func TestEquals(t *testing.T) {

	type equalsTest struct {
		a        *LinkedHashMap[int, int]
		b        *LinkedHashMap[int, int]
		expected bool
	}

	equalsTests := []equalsTest{
		{
			a:        New[int, int](),
			b:        New[int, int](),
			expected: true,
		},
		{
			a:        New[int, int](),
			b:        New(pair.Of(1, 1)),
			expected: false,
		},
		{
			a:        New(pair.Of(1, 2)),
			b:        New(pair.Of(1, 1)),
			expected: false,
		},
		{
			a:        New(pair.Of(2, 1)),
			b:        New(pair.Of(1, 1)),
			expected: false,
		},
		{
			a:        New(pair.Of(2, 2), pair.Of(1, 1)),
			b:        New(pair.Of(1, 1), pair.Of(2, 2)),
			expected: true,
		},
	}

	equals := func(i1, i2 int) bool { return i1 == i2 }
	for _, test := range equalsTests {
		assert.Equal(t, test.expected, test.a.Equals(test.b, equals))
		assert.Equal(t, test.expected, test.b.Equals(test.a, equals))

	}
}
