package hashmap

import (
	"testing"

	"github.com/phantom820/collections/maps"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	hashMap := New[string, int]()
	assert.NotNil(t, hashMap)
	assert.True(t, hashMap.Empty())
	assert.Equal(t, 0, hashMap.Len())
}

func TestPut(t *testing.T) {

	type putTest struct {
		input    HashMap[string, int]
		action   func(HashMap[string, int])
		expected HashMap[string, int]
	}

	putTests := []putTest{
		{input: New[string, int](),
			action: func(hm HashMap[string, int]) {
				hm.Put("A", 1)
			},
			expected: map[string]int{"A": 1},
		},
		{input: New[string, int](),
			action: func(hm HashMap[string, int]) {
				hm.Put("A", 1)
				hm.Put("A", 2)
			},
			expected: map[string]int{"A": 2},
		},
		{input: New[string, int](),
			action: func(hm HashMap[string, int]) {
				hm.Put("A", 1)
				hm.Put("B", 2)
				hm.Put("C", 3)
			},
			expected: map[string]int{"A": 1, "B": 2, "C": 3},
		},
	}

	for _, test := range putTests {
		test.action(test.input)
		assert.Equal(t, test.expected, test.input)
	}
}

func TestPutIfAbsent(t *testing.T) {

	type putIfAbsentTest struct {
		input    HashMap[string, int]
		action   func(HashMap[string, int])
		expected HashMap[string, int]
	}

	putIfAbsentTests := []putIfAbsentTest{
		{input: New[string, int](),
			action: func(hm HashMap[string, int]) {
				hm.PutIfAbsent("A", 1)
			},
			expected: map[string]int{"A": 1},
		},
		{input: New[string, int](),
			action: func(hm HashMap[string, int]) {
				hm.PutIfAbsent("A", 1)
				hm.PutIfAbsent("A", 2)
			},
			expected: map[string]int{"A": 1},
		},
	}

	for _, test := range putIfAbsentTests {
		test.action(test.input)
		assert.Equal(t, test.expected, test.input)
	}
}

func TestGet(t *testing.T) {

	type getTest struct {
		input    HashMap[string, int]
		key      string
		expected int
	}

	getTests := []getTest{
		{input: New[string, int](),
			key:      "A",
			expected: 0,
		},
		{input: map[string]int{"A": 1},
			key:      "A",
			expected: 1,
		},
		{input: map[string]int{"A": 1},
			key:      "B",
			expected: 0,
		},
	}

	for _, test := range getTests {
		assert.Equal(t, test.expected, test.input.Get(test.key))
	}
}

func TestGetIf(t *testing.T) {

	type getIfTest struct {
		input    HashMap[string, int]
		f        func(string) bool
		expected []int
	}

	getIfTests := []getIfTest{
		{input: New[string, int](),
			f:        func(s string) bool { return s == "" },
			expected: []int{},
		},
		{input: map[string]int{"A": 1, "B": 2, "C": 3},
			f:        func(s string) bool { return s == "A" || s == "B" },
			expected: []int{1, 2},
		},
		{input: map[string]int{"A": 1},
			f:        func(s string) bool { return false },
			expected: []int{},
		},
	}

	for _, test := range getIfTests {
		assert.ElementsMatch(t, test.expected, test.input.GetIf(test.f))
	}
}

func TestRemove(t *testing.T) {

	type removeTest struct {
		input    HashMap[string, int]
		key      string
		expected int
	}

	removeTests := []removeTest{
		{input: New[string, int](),
			key:      "A",
			expected: 0,
		},
		{input: map[string]int{"A": 1},
			key:      "A",
			expected: 1,
		},
		{input: map[string]int{"A": 1},
			key:      "B",
			expected: 0,
		},
	}

	for _, test := range removeTests {
		assert.Equal(t, test.expected, test.input.Remove(test.key))
	}
}

func TestRemoveIf(t *testing.T) {

	type removeIfTest struct {
		input    HashMap[string, int]
		f        func(string) bool
		expected HashMap[string, int]
	}

	removeIfTests := []removeIfTest{
		{input: New[string, int](),
			f:        func(s string) bool { return s == "" },
			expected: make(HashMap[string, int]),
		},
		{input: map[string]int{"A": 1, "B": 2, "C": 3},
			f:        func(s string) bool { return s == "A" || s == "B" },
			expected: map[string]int{"C": 3},
		},
		{input: map[string]int{"A": 1},
			f:        func(s string) bool { return false },
			expected: map[string]int{"A": 1},
		},
	}

	for _, test := range removeIfTests {
		test.input.RemoveIf(test.f)
		assert.Equal(t, test.expected, test.input)
	}
}

func TestContainsKey(t *testing.T) {

	type containsKeyTest struct {
		input    HashMap[string, int]
		key      string
		expected bool
	}

	containsKeyTests := []containsKeyTest{
		{
			input:    New[string, int](),
			key:      "A",
			expected: false,
		},
		{
			input:    map[string]int{"C": 1, "D": 1},
			key:      "A",
			expected: false,
		},
		{
			input:    map[string]int{"A": 1},
			key:      "A",
			expected: true,
		},
	}

	for _, test := range containsKeyTests {
		assert.Equal(t, test.expected, test.input.ContainsKey(test.key))
	}
}

func TestContainsValue(t *testing.T) {

	type containsValueTest struct {
		input    HashMap[string, int]
		value    int
		expected bool
	}

	containsValueTests := []containsValueTest{
		{
			input:    New[string, int](),
			value:    0,
			expected: false,
		},
		{
			input:    map[string]int{"C": 1, "D": 1},
			value:    2,
			expected: false,
		},
		{
			input:    map[string]int{"A": 1},
			value:    1,
			expected: true,
		},
	}
	equals := func(v1, v2 int) bool { return v1 == v2 }
	for _, test := range containsValueTests {
		assert.Equal(t, test.expected, test.input.ContainsValue(test.value, equals))
	}
}

func TestClear(t *testing.T) {

	type clearTest struct {
		input    HashMap[string, int]
		expected HashMap[string, int]
	}

	clearTests := []clearTest{
		{
			input:    New[string, int](),
			expected: make(HashMap[string, int]),
		},
		{
			input:    map[string]int{"A": 1, "B": 2},
			expected: make(HashMap[string, int]),
		},
	}

	for _, test := range clearTests {
		test.input.Clear()
		assert.Equal(t, test.expected, test.input)
	}
}

func TestKeys(t *testing.T) {

	type keyTest struct {
		input    HashMap[string, int]
		expected []string
	}

	keyTests := []keyTest{
		{
			input:    New[string, int](),
			expected: []string{},
		},
		{
			input:    map[string]int{"A": 1, "B": 2},
			expected: []string{"A", "B"},
		},
	}

	for _, test := range keyTests {
		assert.ElementsMatch(t, test.expected, test.input.Keys())
	}
}

func TestValues(t *testing.T) {

	type valuesTest struct {
		input    HashMap[string, int]
		expected []int
	}

	valuesTests := []valuesTest{
		{
			input:    New[string, int](),
			expected: []int{},
		},
		{
			input:    map[string]int{"A": 1, "B": 2},
			expected: []int{1, 2},
		},
	}

	for _, test := range valuesTests {
		assert.ElementsMatch(t, test.expected, test.input.Values())
	}
}

func TestForEach(t *testing.T) {

	type forEachTest struct {
		input    HashMap[string, int]
		expected map[string]int
	}

	forEachTests := []forEachTest{
		{
			input:    New[string, int](),
			expected: map[string]int{},
		},
		{
			input:    map[string]int{"A": 1, "B": 2},
			expected: map[string]int{"A": 1, "B": 2},
		},
	}

	for _, test := range forEachTests {
		m := make(map[string]int)
		f := func(s string, i int) {
			m[s] = i
		}
		test.input.ForEach(f)
		assert.Equal(t, test.expected, m)
	}
}

func TestIterator(t *testing.T) {

	iterate := func(it maps.MapIterator[string, int]) []maps.Entry[string, int] {
		entries := make([]maps.Entry[string, int], 0)
		for it.HasNext() {
			entries = append(entries, it.Next())
		}
		return entries
	}

	h := HashMap[string, int](map[string]int{})
	assert.ElementsMatch(t, []maps.Entry[string, int]{}, iterate(h.Iterator()))

	h = HashMap[string, int](map[string]int{"A": 1, "B": 2, "C": 3})
	assert.ElementsMatch(t, []maps.Entry[string, int]{
		maps.NewEntry("A", 1), maps.NewEntry("B", 2), maps.NewEntry("C", 3)}, iterate(h.Iterator()))

	h = HashMap[string, int](map[string]int{"A": 1, "B": 2, "C": 3})
	it := h.Iterator()
	h.Put("F", 23)

	assert.ElementsMatch(t, []maps.Entry[string, int]{
		maps.NewEntry("A", 1), maps.NewEntry("B", 2), maps.NewEntry("C", 3), maps.NewEntry("F", 23)}, iterate(it))

}

func TestString(t *testing.T) {

	assert.Equal(t, "{}", (New[string, string]().String()))
	assert.Equal(t, "{A=1}", (HashMap[string, int](map[string]int{"A": 1}).String()))
}
