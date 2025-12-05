package collections

import (
	"slices"
	"strings"
)

// Pluck extracts values by a key function.
func Pluck[T any, V any](c *Collection[T], keyFn func(T) V) *Collection[V] {
	return Map(c, func(item T, _ int) V {
		return keyFn(item)
	})
}

// PluckMap extracts values by a key function into a map.
func PluckMap[T any, K comparable, V any](c *Collection[T], valueFn func(T) V, keyFn func(T) K) *MapCollection[K, V] {
	result := make(map[K]V)
	keys := make([]K, 0)
	for _, item := range c.items {
		k := keyFn(item)
		result[k] = valueFn(item)
		keys = append(keys, k)
	}
	return NewMapOrdered(result, keys)
}

// KeyBy keys the collection by a function.
func KeyBy[T any, K comparable](c *Collection[T], keyFn func(T) K) *MapCollection[K, T] {
	result := make(map[K]T)
	keys := make([]K, 0)
	for _, item := range c.items {
		k := keyFn(item)
		if _, exists := result[k]; !exists {
			keys = append(keys, k)
		}
		result[k] = item
	}
	return NewMapOrdered(result, keys)
}

// GroupBy groups items by a key function.
func GroupBy[T any, K comparable](c *Collection[T], keyFn func(T) K) *MapCollection[K, *Collection[T]] {
	result := make(map[K]*Collection[T])
	keys := make([]K, 0)
	for _, item := range c.items {
		k := keyFn(item)
		if _, exists := result[k]; !exists {
			result[k] = Empty[T]()
			keys = append(keys, k)
		}
		result[k].Push(item)
	}
	return NewMapOrdered(result, keys)
}

// CountBy counts items by a key function.
func CountBy[T any, K comparable](c *Collection[T], keyFn func(T) K) *MapCollection[K, int] {
	result := make(map[K]int)
	keys := make([]K, 0)
	for _, item := range c.items {
		k := keyFn(item)
		if _, exists := result[k]; !exists {
			keys = append(keys, k)
		}
		result[k]++
	}
	return NewMapOrdered(result, keys)
}

// Collapse flattens a collection of collections.
func Collapse[T any](c *Collection[*Collection[T]]) *Collection[T] {
	result := make([]T, 0)
	for _, item := range c.items {
		result = append(result, item.items...)
	}
	return New(result)
}

// CollapseSlices flattens a collection of slices.
func CollapseSlices[T any](c *Collection[[]T]) *Collection[T] {
	result := make([]T, 0)
	for _, item := range c.items {
		result = append(result, item...)
	}
	return New(result)
}

// Flatten flattens nested slices.
func Flatten[T any](items []any, depth int) []T {
	result := make([]T, 0)
	flattenRecursive(items, depth, &result)
	return result
}

func flattenRecursive[T any](items []any, depth int, result *[]T) {
	for _, item := range items {
		switch v := item.(type) {
		case []any:
			if depth > 0 {
				flattenRecursive[T](v, depth-1, result)
			} else {
				for _, inner := range v {
					if t, ok := inner.(T); ok {
						*result = append(*result, t)
					}
				}
			}
		case T:
			*result = append(*result, v)
		}
	}
}

// FlatMap maps and flattens in one step.
func FlatMap[T any, U any](c *Collection[T], callback func(T, int) []U) *Collection[U] {
	result := make([]U, 0)
	for i, item := range c.items {
		result = append(result, callback(item, i)...)
	}
	return New(result)
}

// MapWithKeys creates a MapCollection from a collection.
func MapWithKeys[T any, K comparable, V any](c *Collection[T], callback func(T, int) (K, V)) *MapCollection[K, V] {
	result := make(map[K]V)
	keys := make([]K, 0)
	for i, item := range c.items {
		k, v := callback(item, i)
		if _, exists := result[k]; !exists {
			keys = append(keys, k)
		}
		result[k] = v
	}
	return NewMapOrdered(result, keys)
}

// MapToDictionary maps each item to a dictionary.
func MapToDictionary[T any, K comparable, V any](c *Collection[T], callback func(T, int) (K, V)) *MapCollection[K, []V] {
	result := make(map[K][]V)
	keys := make([]K, 0)
	for i, item := range c.items {
		k, v := callback(item, i)
		if _, exists := result[k]; !exists {
			result[k] = make([]V, 0)
			keys = append(keys, k)
		}
		result[k] = append(result[k], v)
	}
	return NewMapOrdered(result, keys)
}

// MapToGroups is an alias for MapToDictionary.
func MapToGroups[T any, K comparable, V any](c *Collection[T], callback func(T, int) (K, V)) *MapCollection[K, []V] {
	return MapToDictionary(c, callback)
}

// Zip zips collections together.
func Zip[T any](collections ...*Collection[T]) *Collection[[]T] {
	if len(collections) == 0 {
		return Empty[[]T]()
	}

	// Find minimum length
	minLen := collections[0].Count()
	for _, c := range collections[1:] {
		if c.Count() < minLen {
			minLen = c.Count()
		}
	}

	result := make([][]T, minLen)
	for i := 0; i < minLen; i++ {
		row := make([]T, len(collections))
		for j, c := range collections {
			row[j] = c.items[i]
		}
		result[i] = row
	}
	return New(result)
}

// CrossJoin creates a cross product of collections.
func CrossJoin[T any](collections ...*Collection[T]) *Collection[[]T] {
	if len(collections) == 0 {
		return Empty[[]T]()
	}

	result := [][]T{{}}
	for _, c := range collections {
		newResult := make([][]T, 0)
		for _, existing := range result {
			for _, item := range c.items {
				newRow := slices.Clone(existing)
				newRow = append(newRow, item)
				newResult = append(newResult, newRow)
			}
		}
		result = newResult
	}
	return New(result)
}

// Combine creates a map using one collection as keys and another as values.
func Combine[K comparable, V any](keys *Collection[K], values *Collection[V]) *MapCollection[K, V] {
	result := make(map[K]V)
	resultKeys := make([]K, 0)
	minLen := min(keys.Count(), values.Count())
	for i := 0; i < minLen; i++ {
		k := keys.items[i]
		result[k] = values.items[i]
		resultKeys = append(resultKeys, k)
	}
	return NewMapOrdered(result, resultKeys)
}

// Replace replaces items in the collection.
func (c *Collection[T]) Replace(items []T) *Collection[T] {
	result := slices.Clone(c.items)
	for i := 0; i < min(len(result), len(items)); i++ {
		result[i] = items[i]
	}
	return New(result)
}

// Splice removes and replaces items.
func (c *Collection[T]) Splice(offset int, length ...int) (*Collection[T], *Collection[T]) {
	if offset < 0 {
		offset = max(0, len(c.items)+offset)
	}

	removeLen := len(c.items) - offset
	if len(length) > 0 {
		removeLen = min(length[0], len(c.items)-offset)
	}

	removed := slices.Clone(c.items[offset : offset+removeLen])
	c.items = append(c.items[:offset], c.items[offset+removeLen:]...)
	return c, New(removed)
}

// SpliceReplace removes and replaces items with new items.
func (c *Collection[T]) SpliceReplace(offset, length int, replacement []T) (*Collection[T], *Collection[T]) {
	if offset < 0 {
		offset = max(0, len(c.items)+offset)
	}

	removeLen := min(length, len(c.items)-offset)
	removed := slices.Clone(c.items[offset : offset+removeLen])

	newItems := make([]T, 0, len(c.items)-removeLen+len(replacement))
	newItems = append(newItems, c.items[:offset]...)
	newItems = append(newItems, replacement...)
	newItems = append(newItems, c.items[offset+removeLen:]...)
	c.items = newItems

	return c, New(removed)
}

// Implode joins strings with a separator.
func ImplodeStrings(c *Collection[string], separator string) string {
	return strings.Join(c.items, separator)
}

// Implode joins items with a separator using a string function.
func Implode[T any](c *Collection[T], toStringFn func(T) string, separator string) string {
	parts := make([]string, len(c.items))
	for i, item := range c.items {
		parts[i] = toStringFn(item)
	}
	return strings.Join(parts, separator)
}

// Join joins items with glue, optionally using a final separator.
func JoinStrings(c *Collection[string], glue string, finalGlue ...string) string {
	if c.IsEmpty() {
		return ""
	}
	if c.Count() == 1 {
		return c.First()
	}

	if len(finalGlue) == 0 || finalGlue[0] == "" {
		return ImplodeStrings(c, glue)
	}

	lastItem := c.items[len(c.items)-1]
	rest := c.Slice(0, c.Count()-1)
	return ImplodeStrings(rest, glue) + finalGlue[0] + lastItem
}

// Join with custom string function.
func Join[T any](c *Collection[T], toStringFn func(T) string, glue string, finalGlue ...string) string {
	if c.IsEmpty() {
		return ""
	}
	if c.Count() == 1 {
		return toStringFn(c.First())
	}

	if len(finalGlue) == 0 || finalGlue[0] == "" {
		return Implode(c, toStringFn, glue)
	}

	lastItem := toStringFn(c.items[len(c.items)-1])
	rest := c.Slice(0, c.Count()-1)
	return Implode(rest, toStringFn, glue) + finalGlue[0] + lastItem
}

// Sliding creates a sliding window view.
func (c *Collection[T]) Sliding(size int, step ...int) []*Collection[T] {
	stepSize := 1
	if len(step) > 0 && step[0] > 0 {
		stepSize = step[0]
	}

	if size <= 0 || c.IsEmpty() {
		return []*Collection[T]{}
	}

	result := make([]*Collection[T], 0)
	for i := 0; i+size <= len(c.items); i += stepSize {
		result = append(result, New(slices.Clone(c.items[i:i+size])))
	}
	return result
}

// WhereType filters items that are of a specific type.
// Note: This requires type assertion at runtime.
func WhereType[T any, U any](c *Collection[T]) *Collection[U] {
	result := make([]U, 0)
	for _, item := range c.items {
		if v, ok := any(item).(U); ok {
			result = append(result, v)
		}
	}
	return New(result)
}

// First converts a collection item to another collection type.
func ToCollection[T any](items []T) *Collection[T] {
	return New(items)
}
