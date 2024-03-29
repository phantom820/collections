// package linkedhashmap defines a map implementation in which entries are iterated in the order they were inserted.
package linkedhashmap

import (
	"fmt"
	"strings"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/errors"
	"github.com/phantom820/collections/iterator"
	"github.com/phantom820/collections/maps/hashmap"
	"github.com/phantom820/collections/types/optional"
	"github.com/phantom820/collections/types/pair"
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

// LinkedHashMap implementation of a map with a predictable order of iteration.
type LinkedHashMap[K comparable, V any] struct {
	head    *node[K, V]
	hashMap hashmap.HashMap[K, *node[K, V]]
	tail    *node[K, V]
}

// New creates a map with the given key, value pairs.
func New[K comparable, V any](pairs ...pair.Pair[K, V]) *LinkedHashMap[K, V] {
	linkedHashMap := LinkedHashMap[K, V]{hashMap: hashmap.New[K, *node[K, V]]()}
	for _, pair := range pairs {
		linkedHashMap.Put(pair.Key(), pair.Value())
	}
	return &linkedHashMap
}

// Put adds a new key/value pair to the map and optionally returns previously bound value.
func (linkedHashMap *LinkedHashMap[K, V]) Put(key K, value V) optional.Optional[V] {
	if linkedHashMap.Empty() {
		node := newNode(key, value)
		linkedHashMap.head = node
		linkedHashMap.tail = node
		linkedHashMap.hashMap.Put(key, node)
		return optional.Empty[V]()
	} else if storedNode, ok := linkedHashMap.hashMap[key]; ok {
		// The key is already mapped and we swap out the value.
		storedValue := storedNode.value
		storedNode.value = value
		return optional.Of(storedValue)
	}
	// Effectively inserting at the back of a linked list.
	node := newNode(key, value)
	node.prev = linkedHashMap.tail
	node.prev.next = node
	node.next = nil
	linkedHashMap.tail = node
	linkedHashMap.hashMap.Put(key, node)
	return optional.Empty[V]()
}

// PutIfAbsent  adds a new key/value pair to the map if the key is not already bounded and optionally returns bound value.
func (linkedHashMap *LinkedHashMap[K, V]) PutIfAbsent(key K, value V) optional.Optional[V] {
	if storedValue, ok := linkedHashMap.hashMap[key]; ok {
		return optional.Of(storedValue.value)
	}
	linkedHashMap.Put(key, value)
	return optional.Empty[V]()
}

// Get optionally returns the value associated with a key.
func (linkedHashMap *LinkedHashMap[K, V]) Get(key K) optional.Optional[V] {
	node := linkedHashMap.hashMap.Get(key)
	if node.Empty() {
		return optional.Empty[V]()
	}
	return optional.Of(node.Value().value)
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

// Remove removes a key from the map, returning the value associated previously with that key as an option.
func (linkedHashMap *LinkedHashMap[K, V]) Remove(key K) optional.Optional[V] {
	node, ok := linkedHashMap.hashMap[key]
	delete(linkedHashMap.hashMap, key)
	if !ok {
		return optional.Empty[V]()
	} else if node == linkedHashMap.head {
		linkedHashMap.head = node.next
		node.next = nil
		node.prev = nil
		storedValue := node.value
		node = nil
		return optional.Of(storedValue)
	} else if node == linkedHashMap.tail {
		linkedHashMap.tail = linkedHashMap.tail.prev
		linkedHashMap.tail.next = nil
		node.next = nil
		node.prev = nil
		storedValue := node.value
		node = nil
		return optional.Of(storedValue)
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
	storedValue := node.value
	node = nil
	return optional.Of(storedValue)
}

// RemoveIf removes all the key, value mapping in which the key matches the given predicate.
func (linkedHashMap *LinkedHashMap[K, V]) RemoveIf(f func(K) bool) bool {
	n := linkedHashMap.Len()
	keysToRemove := make([]K, 0)
	for curr := linkedHashMap.head; curr != nil; curr = curr.next {
		if f(curr.key) {
			keysToRemove = append(keysToRemove, curr.key)
		}
	}

	for _, key := range keysToRemove {
		linkedHashMap.Remove(key)
	}
	return n != linkedHashMap.Len()
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

// Len returns the size of the map.
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
func (linkedHashMap LinkedHashMap[K, V]) Iterator() iterator.Iterator[pair.Pair[K, V]] {
	return &mapIterator[K, V]{initialized: false, initialize: func() *node[K, V] { return linkedHashMap.head }}
}

// mapIterator implementation of an mapIterator for [LinkedHashMap].
type mapIterator[K comparable, V any] struct {
	initialized bool
	initialize  func() *node[K, V]
	head        *node[K, V]
}

// HasNext returns true if the iterator has more elements.
func (it *mapIterator[K, V]) HasNext() bool {
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
func (it *mapIterator[K, V]) Next() pair.Pair[K, V] {
	if !it.HasNext() {
		panic(errors.NoSuchElement())
	}
	entry := pair.Of(it.head.key, it.head.value)
	it.head = it.head.next
	return entry
}

// String returns the string representation of the map.
func (linkedHashMap *LinkedHashMap[K, V]) String() string {
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

// Equals return true if the map is is equal to the given map. Two maps are equal if they contain the same
// key, value pairs.
func (linkedHashMap *LinkedHashMap[K, V]) Equals(other collections.Map[K, V], equals func(V, V) bool) bool {
	if linkedHashMap.Len() != other.Len() {
		return false
	}
	it := other.Iterator()
	for it.HasNext() {
		pair := it.Next()
		result := linkedHashMap.Get(pair.Key())
		if result.Empty() {
			return false
		} else if !equals(pair.Value(), result.Value()) {
			return false
		}
	}
	return true
}
