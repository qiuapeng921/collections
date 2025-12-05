package collections

import (
	"cmp"
	"slices"
	"sort"
)

// Sort returns a sorted copy of the collection.
func Sort[T cmp.Ordered](c *Collection[T]) *Collection[T] {
	result := slices.Clone(c.items)
	slices.Sort(result)
	return New(result)
}

// SortDesc returns a sorted copy in descending order.
func SortDesc[T cmp.Ordered](c *Collection[T]) *Collection[T] {
	result := slices.Clone(c.items)
	slices.SortFunc(result, func(a, b T) int {
		return cmp.Compare(b, a)
	})
	return New(result)
}

// SortBy sorts the collection by a key function.
func SortBy[T any, K cmp.Ordered](c *Collection[T], keyFn func(T) K) *Collection[T] {
	result := slices.Clone(c.items)
	slices.SortStableFunc(result, func(a, b T) int {
		return cmp.Compare(keyFn(a), keyFn(b))
	})
	return New(result)
}

// SortByDesc sorts the collection by a key function in descending order.
func SortByDesc[T any, K cmp.Ordered](c *Collection[T], keyFn func(T) K) *Collection[T] {
	result := slices.Clone(c.items)
	slices.SortStableFunc(result, func(a, b T) int {
		return cmp.Compare(keyFn(b), keyFn(a))
	})
	return New(result)
}

// SortFunc sorts the collection using a custom comparison function.
func (c *Collection[T]) SortFunc(less func(a, b T) int) *Collection[T] {
	result := slices.Clone(c.items)
	slices.SortFunc(result, less)
	return New(result)
}

// SortStableFunc sorts stably using a custom comparison function.
func (c *Collection[T]) SortStableFunc(less func(a, b T) int) *Collection[T] {
	result := slices.Clone(c.items)
	slices.SortStableFunc(result, less)
	return New(result)
}

// Min returns the minimum value in the collection.
func Min[T cmp.Ordered](c *Collection[T]) T {
	if c.IsEmpty() {
		var zero T
		return zero
	}
	return slices.Min(c.items)
}

// Max returns the maximum value in the collection.
func Max[T cmp.Ordered](c *Collection[T]) T {
	if c.IsEmpty() {
		var zero T
		return zero
	}
	return slices.Max(c.items)
}

// MinBy returns the item with the minimum key value.
func MinBy[T any, K cmp.Ordered](c *Collection[T], keyFn func(T) K) T {
	if c.IsEmpty() {
		var zero T
		return zero
	}
	minItem := c.items[0]
	minKey := keyFn(minItem)
	for _, item := range c.items[1:] {
		key := keyFn(item)
		if key < minKey {
			minItem = item
			minKey = key
		}
	}
	return minItem
}

// MaxBy returns the item with the maximum key value.
func MaxBy[T any, K cmp.Ordered](c *Collection[T], keyFn func(T) K) T {
	if c.IsEmpty() {
		var zero T
		return zero
	}
	maxItem := c.items[0]
	maxKey := keyFn(maxItem)
	for _, item := range c.items[1:] {
		key := keyFn(item)
		if key > maxKey {
			maxItem = item
			maxKey = key
		}
	}
	return maxItem
}

// Sum returns the sum of all items.
func Sum[T Numeric](c *Collection[T]) T {
	var sum T
	for _, item := range c.items {
		sum += item
	}
	return sum
}

// SumBy returns the sum of values returned by the key function.
func SumBy[T any, N Numeric](c *Collection[T], keyFn func(T) N) N {
	var sum N
	for _, item := range c.items {
		sum += keyFn(item)
	}
	return sum
}

// Avg returns the average of all items.
func Avg[T Numeric](c *Collection[T]) float64 {
	if c.IsEmpty() {
		return 0
	}
	return float64(Sum(c)) / float64(c.Count())
}

// AvgBy returns the average of values returned by the key function.
func AvgBy[T any, N Numeric](c *Collection[T], keyFn func(T) N) float64 {
	if c.IsEmpty() {
		return 0
	}
	return float64(SumBy(c, keyFn)) / float64(c.Count())
}

// Median returns the median value of the collection.
func Median[T Numeric](c *Collection[T]) float64 {
	if c.IsEmpty() {
		return 0
	}

	sorted := Sort(c)
	count := sorted.Count()
	middle := count / 2

	if count%2 == 0 {
		return (float64(sorted.items[middle-1]) + float64(sorted.items[middle])) / 2
	}
	return float64(sorted.items[middle])
}

// Mode returns the most common value(s) in the collection.
func Mode[T comparable](c *Collection[T]) []T {
	if c.IsEmpty() {
		return nil
	}

	counts := make(map[T]int)
	maxCount := 0

	for _, item := range c.items {
		counts[item]++
		if counts[item] > maxCount {
			maxCount = counts[item]
		}
	}

	result := make([]T, 0)
	for item, count := range counts {
		if count == maxCount {
			result = append(result, item)
		}
	}
	return result
}

// Diff returns items not present in the other collection.
func Diff[T comparable](c *Collection[T], other *Collection[T]) *Collection[T] {
	set := make(map[T]bool)
	for _, item := range other.items {
		set[item] = true
	}

	result := make([]T, 0)
	for _, item := range c.items {
		if !set[item] {
			result = append(result, item)
		}
	}
	return New(result)
}

// Intersect returns items present in both collections.
func Intersect[T comparable](c *Collection[T], other *Collection[T]) *Collection[T] {
	set := make(map[T]bool)
	for _, item := range other.items {
		set[item] = true
	}

	result := make([]T, 0)
	for _, item := range c.items {
		if set[item] {
			result = append(result, item)
		}
	}
	return New(result)
}

// Duplicates returns duplicate items.
func Duplicates[T comparable](c *Collection[T]) *Collection[T] {
	seen := make(map[T]bool)
	duplicates := make(map[T]bool)
	result := make([]T, 0)

	for _, item := range c.items {
		if seen[item] && !duplicates[item] {
			duplicates[item] = true
			result = append(result, item)
		}
		seen[item] = true
	}
	return New(result)
}

// UniqueComparable returns unique items from a comparable collection.
func UniqueComparable[T comparable](c *Collection[T]) *Collection[T] {
	seen := make(map[T]bool)
	result := make([]T, 0)
	for _, item := range c.items {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return New(result)
}

// ContainsComparable checks if the collection contains the given item.
func ContainsComparable[T comparable](c *Collection[T], item T) bool {
	for _, v := range c.items {
		if v == item {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of an item.
func IndexOf[T comparable](c *Collection[T], item T) int {
	for i, v := range c.items {
		if v == item {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of an item.
func LastIndexOf[T comparable](c *Collection[T], item T) int {
	for i := len(c.items) - 1; i >= 0; i-- {
		if c.items[i] == item {
			return i
		}
	}
	return -1
}

// SortByMultiple sorts by multiple keys.
type SortKey[T any] struct {
	KeyFn      func(T) any
	Descending bool
}

// SortByKeys sorts by multiple keys in order.
func SortByKeys[T any](c *Collection[T], keys []SortKey[T]) *Collection[T] {
	result := slices.Clone(c.items)
	sort.SliceStable(result, func(i, j int) bool {
		for _, key := range keys {
			vi := key.KeyFn(result[i])
			vj := key.KeyFn(result[j])

			cmpResult := compareAny(vi, vj)
			if cmpResult != 0 {
				if key.Descending {
					return cmpResult > 0
				}
				return cmpResult < 0
			}
		}
		return false
	})
	return New(result)
}

// compareAny compares two values of the same underlying type.
func compareAny(a, b any) int {
	switch va := a.(type) {
	case int:
		return cmp.Compare(va, b.(int))
	case int8:
		return cmp.Compare(va, b.(int8))
	case int16:
		return cmp.Compare(va, b.(int16))
	case int32:
		return cmp.Compare(va, b.(int32))
	case int64:
		return cmp.Compare(va, b.(int64))
	case uint:
		return cmp.Compare(va, b.(uint))
	case uint8:
		return cmp.Compare(va, b.(uint8))
	case uint16:
		return cmp.Compare(va, b.(uint16))
	case uint32:
		return cmp.Compare(va, b.(uint32))
	case uint64:
		return cmp.Compare(va, b.(uint64))
	case float32:
		return cmp.Compare(va, b.(float32))
	case float64:
		return cmp.Compare(va, b.(float64))
	case string:
		return cmp.Compare(va, b.(string))
	default:
		return 0
	}
}
