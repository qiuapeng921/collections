package collections

import (
	"encoding/json"
	"fmt"
	"slices"
	"sort"
)

// MapCollection represents a key-value collection (similar to PHP associative arrays).
type MapCollection[K comparable, V any] struct {
	items map[K]V
	keys  []K // Maintains insertion order
}

// NewMap creates a new MapCollection from a map.
func NewMap[K comparable, V any](items map[K]V) *MapCollection[K, V] {
	if items == nil {
		items = make(map[K]V)
	}
	keys := make([]K, 0, len(items))
	for k := range items {
		keys = append(keys, k)
	}
	return &MapCollection[K, V]{items: items, keys: keys}
}

// NewMapOrdered creates a new MapCollection with ordered keys.
func NewMapOrdered[K comparable, V any](items map[K]V, keys []K) *MapCollection[K, V] {
	if items == nil {
		items = make(map[K]V)
	}
	return &MapCollection[K, V]{items: items, keys: slices.Clone(keys)}
}

// All returns all items as a map.
func (m *MapCollection[K, V]) All() map[K]V {
	return m.items
}

// Keys returns all keys.
func (m *MapCollection[K, V]) Keys() *Collection[K] {
	return New(slices.Clone(m.keys))
}

// Values returns all values.
func (m *MapCollection[K, V]) Values() *Collection[V] {
	values := make([]V, len(m.keys))
	for i, k := range m.keys {
		values[i] = m.items[k]
	}
	return New(values)
}

// Get returns the value for the given key.
func (m *MapCollection[K, V]) Get(key K) V {
	return m.items[key]
}

// GetOr returns the value for the key or a default value.
func (m *MapCollection[K, V]) GetOr(key K, defaultValue V) V {
	if v, ok := m.items[key]; ok {
		return v
	}
	return defaultValue
}

// Has determines if a key exists.
func (m *MapCollection[K, V]) Has(keys ...K) bool {
	for _, key := range keys {
		if _, ok := m.items[key]; !ok {
			return false
		}
	}
	return true
}

// HasAny determines if any of the keys exist.
func (m *MapCollection[K, V]) HasAny(keys ...K) bool {
	for _, key := range keys {
		if _, ok := m.items[key]; ok {
			return true
		}
	}
	return false
}

// Put sets a key-value pair.
func (m *MapCollection[K, V]) Put(key K, value V) *MapCollection[K, V] {
	if _, exists := m.items[key]; !exists {
		m.keys = append(m.keys, key)
	}
	m.items[key] = value
	return m
}

// Pull gets and removes an item.
func (m *MapCollection[K, V]) Pull(key K) V {
	value := m.items[key]
	m.Forget(key)
	return value
}

// Forget removes one or more keys.
func (m *MapCollection[K, V]) Forget(keys ...K) *MapCollection[K, V] {
	for _, key := range keys {
		delete(m.items, key)
		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
	return m
}

// Count returns the number of items.
func (m *MapCollection[K, V]) Count() int {
	return len(m.items)
}

// IsEmpty determines if the collection is empty.
func (m *MapCollection[K, V]) IsEmpty() bool {
	return len(m.items) == 0
}

// IsNotEmpty determines if the collection is not empty.
func (m *MapCollection[K, V]) IsNotEmpty() bool {
	return !m.IsEmpty()
}

// Each iterates over each item.
func (m *MapCollection[K, V]) Each(callback func(K, V)) *MapCollection[K, V] {
	for _, k := range m.keys {
		callback(k, m.items[k])
	}
	return m
}

// EachBreak iterates and allows breaking.
func (m *MapCollection[K, V]) EachBreak(callback func(K, V) bool) *MapCollection[K, V] {
	for _, k := range m.keys {
		if !callback(k, m.items[k]) {
			break
		}
	}
	return m
}

// MapValues applies a callback to each value.
func MapValues[K comparable, V any, U any](m *MapCollection[K, V], callback func(V, K) U) *MapCollection[K, U] {
	result := make(map[K]U)
	for _, k := range m.keys {
		result[k] = callback(m.items[k], k)
	}
	return NewMapOrdered(result, m.keys)
}

// Filter returns items that pass the predicate.
func (m *MapCollection[K, V]) Filter(predicate func(V, K) bool) *MapCollection[K, V] {
	result := make(map[K]V)
	keys := make([]K, 0)
	for _, k := range m.keys {
		if predicate(m.items[k], k) {
			result[k] = m.items[k]
			keys = append(keys, k)
		}
	}
	return NewMapOrdered(result, keys)
}

// Reject returns items that don't pass the predicate.
func (m *MapCollection[K, V]) Reject(predicate func(V, K) bool) *MapCollection[K, V] {
	return m.Filter(func(v V, k K) bool {
		return !predicate(v, k)
	})
}

// Only returns items with the specified keys.
func (m *MapCollection[K, V]) Only(keys ...K) *MapCollection[K, V] {
	keySet := make(map[K]bool)
	for _, k := range keys {
		keySet[k] = true
	}

	result := make(map[K]V)
	resultKeys := make([]K, 0)
	for _, k := range m.keys {
		if keySet[k] {
			result[k] = m.items[k]
			resultKeys = append(resultKeys, k)
		}
	}
	return NewMapOrdered(result, resultKeys)
}

// Except returns items without the specified keys.
func (m *MapCollection[K, V]) Except(keys ...K) *MapCollection[K, V] {
	keySet := make(map[K]bool)
	for _, k := range keys {
		keySet[k] = true
	}

	result := make(map[K]V)
	resultKeys := make([]K, 0)
	for _, k := range m.keys {
		if !keySet[k] {
			result[k] = m.items[k]
			resultKeys = append(resultKeys, k)
		}
	}
	return NewMapOrdered(result, resultKeys)
}

// Merge merges other MapCollections.
func (m *MapCollection[K, V]) Merge(others ...*MapCollection[K, V]) *MapCollection[K, V] {
	result := make(map[K]V)
	resultKeys := make([]K, 0)

	// Copy current items
	for _, k := range m.keys {
		result[k] = m.items[k]
		resultKeys = append(resultKeys, k)
	}

	// Merge others
	for _, other := range others {
		for _, k := range other.keys {
			if _, exists := result[k]; !exists {
				resultKeys = append(resultKeys, k)
			}
			result[k] = other.items[k]
		}
	}

	return NewMapOrdered(result, resultKeys)
}

// Union combines collections, keeping existing values.
func (m *MapCollection[K, V]) Union(other *MapCollection[K, V]) *MapCollection[K, V] {
	result := make(map[K]V)
	resultKeys := make([]K, 0)

	// Copy current items
	for _, k := range m.keys {
		result[k] = m.items[k]
		resultKeys = append(resultKeys, k)
	}

	// Add from other only if not exists
	for _, k := range other.keys {
		if _, exists := result[k]; !exists {
			result[k] = other.items[k]
			resultKeys = append(resultKeys, k)
		}
	}

	return NewMapOrdered(result, resultKeys)
}

// DiffKeys returns items whose keys are not in the other collection.
func (m *MapCollection[K, V]) DiffKeys(other *MapCollection[K, V]) *MapCollection[K, V] {
	result := make(map[K]V)
	resultKeys := make([]K, 0)

	for _, k := range m.keys {
		if _, exists := other.items[k]; !exists {
			result[k] = m.items[k]
			resultKeys = append(resultKeys, k)
		}
	}

	return NewMapOrdered(result, resultKeys)
}

// IntersectByKeys returns items whose keys are in both collections.
func (m *MapCollection[K, V]) IntersectByKeys(other *MapCollection[K, V]) *MapCollection[K, V] {
	result := make(map[K]V)
	resultKeys := make([]K, 0)

	for _, k := range m.keys {
		if _, exists := other.items[k]; exists {
			result[k] = m.items[k]
			resultKeys = append(resultKeys, k)
		}
	}

	return NewMapOrdered(result, resultKeys)
}

// First returns the first value.
func (m *MapCollection[K, V]) First() V {
	if m.IsEmpty() {
		var zero V
		return zero
	}
	return m.items[m.keys[0]]
}

// Last returns the last value.
func (m *MapCollection[K, V]) Last() V {
	if m.IsEmpty() {
		var zero V
		return zero
	}
	return m.items[m.keys[len(m.keys)-1]]
}

// FirstKey returns the first key.
func (m *MapCollection[K, V]) FirstKey() K {
	if m.IsEmpty() {
		var zero K
		return zero
	}
	return m.keys[0]
}

// LastKey returns the last key.
func (m *MapCollection[K, V]) LastKey() K {
	if m.IsEmpty() {
		var zero K
		return zero
	}
	return m.keys[len(m.keys)-1]
}

// Clone returns a copy of the MapCollection.
func (m *MapCollection[K, V]) Clone() *MapCollection[K, V] {
	result := make(map[K]V)
	for k, v := range m.items {
		result[k] = v
	}
	return NewMapOrdered(result, slices.Clone(m.keys))
}

// ToJSON converts to JSON.
func (m *MapCollection[K, V]) ToJSON() ([]byte, error) {
	return json.Marshal(m.items)
}

// ToJSONString converts to JSON string.
func (m *MapCollection[K, V]) ToJSONString() string {
	data, err := m.ToJSON()
	if err != nil {
		return "{}"
	}
	return string(data)
}

// String returns a string representation.
func (m *MapCollection[K, V]) String() string {
	return m.ToJSONString()
}

// Dump prints the collection.
func (m *MapCollection[K, V]) Dump() *MapCollection[K, V] {
	fmt.Printf("%+v\n", m.items)
	return m
}

// Contains checks if any value passes the predicate.
func (m *MapCollection[K, V]) Contains(predicate func(V, K) bool) bool {
	for _, k := range m.keys {
		if predicate(m.items[k], k) {
			return true
		}
	}
	return false
}

// Every checks if all values pass the predicate.
func (m *MapCollection[K, V]) Every(predicate func(V, K) bool) bool {
	for _, k := range m.keys {
		if !predicate(m.items[k], k) {
			return false
		}
	}
	return true
}

// ReduceMap reduces the MapCollection to a single value.
func ReduceMap[K comparable, V any, R any](m *MapCollection[K, V], callback func(R, V, K) R, initial R) R {
	result := initial
	for _, k := range m.keys {
		result = callback(result, m.items[k], k)
	}
	return result
}

// SortKeys sorts the collection by keys.
func SortMapKeys[K interface {
	comparable
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}, V any](m *MapCollection[K, V]) *MapCollection[K, V] {
	result := m.Clone()
	sort.Slice(result.keys, func(i, j int) bool {
		return result.keys[i] < result.keys[j]
	})
	return result
}

// SortKeysDesc sorts the collection by keys in descending order.
func SortMapKeysDesc[K interface {
	comparable
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}, V any](m *MapCollection[K, V]) *MapCollection[K, V] {
	result := m.Clone()
	sort.Slice(result.keys, func(i, j int) bool {
		return result.keys[i] > result.keys[j]
	})
	return result
}

// GetOrPut gets a value or puts a default if not exists.
func (m *MapCollection[K, V]) GetOrPut(key K, defaultValue V) V {
	if v, exists := m.items[key]; exists {
		return v
	}
	m.Put(key, defaultValue)
	return defaultValue
}

// Tap passes the collection to a callback.
func (m *MapCollection[K, V]) Tap(callback func(*MapCollection[K, V])) *MapCollection[K, V] {
	callback(m)
	return m
}

// When applies the callback if condition is true.
func (m *MapCollection[K, V]) When(condition bool, callback func(*MapCollection[K, V]) *MapCollection[K, V]) *MapCollection[K, V] {
	if condition {
		return callback(m)
	}
	return m
}

// Flip swaps keys and values (for string-string maps).
func FlipMap(m *MapCollection[string, string]) *MapCollection[string, string] {
	result := make(map[string]string)
	for k, v := range m.items {
		result[v] = k
	}
	return NewMap(result)
}

// ToSlice converts to a slice of key-value pairs.
type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

func (m *MapCollection[K, V]) ToSlice() *Collection[KeyValue[K, V]] {
	result := make([]KeyValue[K, V], len(m.keys))
	for i, k := range m.keys {
		result[i] = KeyValue[K, V]{Key: k, Value: m.items[k]}
	}
	return New(result)
}

// FromSlice creates a MapCollection from a slice of KeyValue pairs.
func FromSlice[K comparable, V any](items []KeyValue[K, V]) *MapCollection[K, V] {
	result := make(map[K]V)
	keys := make([]K, len(items))
	for i, item := range items {
		result[item.Key] = item.Value
		keys[i] = item.Key
	}
	return NewMapOrdered(result, keys)
}
