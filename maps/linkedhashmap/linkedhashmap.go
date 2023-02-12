package linkedhashmap

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections/maps"
	"github.com/phantom820/collections/maps/hashmap"
)

const (
	DEFAULT_CAPACITY = 16 // Initial capacity of the hash table.
)

type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

func newNode[K comparable, V any](key K, value V) *node[K, V] {
	return &node[K, V]{key: key, value: value}
}

// LinkedHashMap implementation of [HashMap] with a predicatble order of key,value pairs for iteration.
type LinkedHashMap[K comparable, V any] struct {
	head    *node[K, V]
	hashMap hashmap.HashMap[K, *node[K, V]]
	tail    *node[K, V]
}

// New creates an empty LinkedHashMap.
func New[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{hashMap: hashmap.New[K, *node[K, V]]()}
}

// Put associates the specified value with the specified key in the map.
func (linkedHashMap *LinkedHashMap[K, V]) Put(key K, value V) V {
	if linkedHashMap.Empty() {
		node := newNode(key, value)
		linkedHashMap.head = node
		linkedHashMap.tail = node
		linkedHashMap.hashMap.Put(key, node)
		var zeroValue V
		return zeroValue
	} else if storedNode, ok := linkedHashMap.hashMap[key]; ok {
		// The key is alredy mapped and we swap out the value.
		storedValue := storedNode.value
		storedNode.value = value
		return storedValue
	}
	// Effectively inserting at the back of a linked list.
	node := newNode(key, value)
	node.prev = linkedHashMap.tail
	node.prev.next = node
	node.next = nil
	linkedHashMap.tail = node
	linkedHashMap.hashMap.Put(key, node)
	var zeroValue V
	return zeroValue
}

// PutIfAbsent associates the specified key with the given value if the key is not already mapped. Will return the
// current value if present otherwise the zero value.
func (linkedHashMap *LinkedHashMap[K, V]) PutIfAbsent(key K, value V) bool {
	if _, ok := linkedHashMap.hashMap[key]; ok {
		return false
	}
	linkedHashMap.Put(key, value)
	return true
}

// Get returns the value to which the specified key is mapped, or the zero value if the key is not present.
func (linkedHashMap *LinkedHashMap[K, V]) Get(key K) V {
	node := linkedHashMap.hashMap.Get(key)
	if node == nil {
		var zeroValue V
		return zeroValue
	}
	return node.value
}

// GetIf returns the values mapped by keys that match the given predicate.
func (linkedHashMap LinkedHashMap[K, V]) GetIf(f func(K) bool) []V {
	values := make([]V, 0)
	for curr := linkedHashMap.head; curr != nil; curr = curr.next {
		if f(curr.key) {
			values = append(values, curr.value)
		}

	}
	return values
}

// Remove removes the mapping for the specified key from the map if present
func (linkedHashMap *LinkedHashMap[K, V]) Remove(key K) V {
	node, ok := linkedHashMap.hashMap[key]
	delete(linkedHashMap.hashMap, key)
	if !ok {
		var zero V
		return zero
	} else if node == linkedHashMap.head {
		linkedHashMap.head = node.next
		node.next = nil
		node.prev = nil
		storedValue := node.value
		node = nil
		return storedValue
	} else if node == linkedHashMap.tail {
		linkedHashMap.tail = linkedHashMap.tail.prev
		linkedHashMap.tail.next = nil
		node.next = nil
		node.prev = nil
		storedValue := node.value
		node = nil
		return storedValue
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
	storedValue := node.value
	node = nil
	return storedValue
}

// RemoveIf removes all the key, value mapping in which the key matches the given predicate.
func (linkedHashMap *LinkedHashMap[K, V]) RemoveIf(f func(K) bool) {
	keysToRemove := make([]K, 0)
	for curr := linkedHashMap.head; curr != nil; curr = curr.next {
		if f(curr.key) {
			keysToRemove = append(keysToRemove, curr.key)
		}
	}

	for _, key := range keysToRemove {
		linkedHashMap.Remove(key)
	}

}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (linkedHashMap *LinkedHashMap[K, V]) ContainsKey(key K) bool {
	_, ok := linkedHashMap.hashMap[key]
	return ok
}

// ContainsValue returns true if this map maps one or more keys to the specified value.
func (linkedHashMap *LinkedHashMap[K, V]) ContainsValue(value V, equals func(v1, v2 V) bool) bool {
	for _, storedValue := range linkedHashMap.hashMap {
		if equals(storedValue.value, value) {
			return true
		}
	}
	return false
}

// Clear removes all of the mappings from this map.
func (linkedHashMap *LinkedHashMap[K, V]) Clear() {
	linkedHashMap.hashMap.Clear()
	linkedHashMap.head = nil
	linkedHashMap.tail = nil
}

// Keys returns a slice containing the keys in the map.
func (linkedHashMap *LinkedHashMap[K, V]) Keys() []K {
	keys := make([]K, 0)
	// i := 0
	for curr := linkedHashMap.head; curr != nil; curr = curr.next {
		keys = append(keys, curr.key)
	}
	return keys
}

// Values returns a slice containing the values in the map.
func (linkedHashMap *LinkedHashMap[K, V]) Values() []V {
	values := make([]V, len(linkedHashMap.hashMap))
	i := 0
	for curr := linkedHashMap.head; curr != nil; curr = curr.next {
		values[i] = curr.value
		i++
	}
	return values
}

// Len returns the number of key, value mappings in the map.
func (linkedHashMap *LinkedHashMap[K, V]) Len() int {
	return len(linkedHashMap.hashMap)
}

// Empty returns true if the map has no elements.
func (linkedHashMap *LinkedHashMap[K, V]) Empty() bool {
	return len(linkedHashMap.hashMap) == 0
}

// ForEach performs the given action for each key, value mapping in the map.
func (linkedHashMap *LinkedHashMap[K, V]) ForEach(f func(K, V)) {
	for curr := linkedHashMap.head; curr != nil; curr = curr.next {
		f(curr.key, curr.value)
	}
}

// Iterator returns an iterator over the map. Elements are iterated over following their insertion order.
func (linkedHashMap LinkedHashMap[K, V]) Iterator() maps.MapIterator[K, V] {
	return &iterator[K, V]{initialized: false, initialize: func() *node[K, V] { return linkedHashMap.head }}
}

// iterator implementation of an iterator for [LinkedHashMap].
type iterator[K comparable, V any] struct {
	initialized bool
	initialize  func() *node[K, V]
	head        *node[K, V]
}

// HasNext returns true if the iterator has more elements.
func (it *iterator[K, V]) HasNext() bool {
	if !it.initialized {
		it.initialized = true
		it.head = it.initialize()
	}
	if it.head != nil {
		return true
	}
	return false
}

// Next returns the next element in the iterator.
func (it *iterator[K, V]) Next() maps.Entry[K, V] {
	if !it.HasNext() {
		panic("iterator things shoould panic here")
	}
	entry := maps.NewEntry(it.head.key, it.head.value)
	it.head = it.head.next
	return entry
}

// String returns the string representation of the map.
func (linkedHashMap LinkedHashMap[K, V]) String() string {
	var sb strings.Builder
	if linkedHashMap.Empty() {
		return "{}"
	}
	sb.WriteString("{")
	sb.WriteString(fmt.Sprintf("%v=%v", linkedHashMap.head.key, linkedHashMap.head.value))
	for curr := linkedHashMap.head.next; curr != nil; curr = curr.next {
		sb.WriteString(fmt.Sprintf(", %v=%v", curr.key, curr.value))
	}
	sb.WriteString("}")
	return sb.String()
}
