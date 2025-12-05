// Package collections provides a fluent, convenient wrapper for working with slices of data.
// It is inspired by Laravel's Collection class and provides similar functionality using Go generics.
package collections

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"slices"
	"sort"
)

// Collection represents a wrapper around a slice providing fluent operations.
// T is the type of elements in the collection.
type Collection[T any] struct {
	items []T
}

// New creates a new Collection from a slice.
func New[T any](items []T) *Collection[T] {
	if items == nil {
		items = make([]T, 0)
	}
	return &Collection[T]{items: items}
}

// Make creates a new Collection from variadic arguments.
func Make[T any](items ...T) *Collection[T] {
	return New(items)
}

// Range creates a new Collection with a range of integers.
func Range(from, to int) *Collection[int] {
	if from > to {
		items := make([]int, from-to+1)
		for i := from; i >= to; i-- {
			items[from-i] = i
		}
		return New(items)
	}
	items := make([]int, to-from+1)
	for i := from; i <= to; i++ {
		items[i-from] = i
	}
	return New(items)
}

// Times creates a new Collection by invoking a callback a given number of times.
func Times[T any](n int, callback func(int) T) *Collection[T] {
	items := make([]T, n)
	for i := 0; i < n; i++ {
		items[i] = callback(i + 1)
	}
	return New(items)
}

// Empty creates an empty Collection.
func Empty[T any]() *Collection[T] {
	return New([]T{})
}

// All returns all items in the collection.
func (c *Collection[T]) All() []T {
	return c.items
}

// Count returns the number of items in the collection.
func (c *Collection[T]) Count() int {
	return len(c.items)
}

// IsEmpty determines if the collection is empty.
func (c *Collection[T]) IsEmpty() bool {
	return len(c.items) == 0
}

// IsNotEmpty determines if the collection is not empty.
func (c *Collection[T]) IsNotEmpty() bool {
	return !c.IsEmpty()
}

// First returns the first item in the collection.
// Returns zero value if collection is empty.
func (c *Collection[T]) First() T {
	if c.IsEmpty() {
		var zero T
		return zero
	}
	return c.items[0]
}

// FirstOr returns the first item or the default value if empty.
func (c *Collection[T]) FirstOr(defaultValue T) T {
	if c.IsEmpty() {
		return defaultValue
	}
	return c.items[0]
}

// FirstWhere returns the first item matching the predicate.
func (c *Collection[T]) FirstWhere(predicate func(T) bool) (T, bool) {
	for _, item := range c.items {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// Last returns the last item in the collection.
func (c *Collection[T]) Last() T {
	if c.IsEmpty() {
		var zero T
		return zero
	}
	return c.items[len(c.items)-1]
}

// LastOr returns the last item or the default value if empty.
func (c *Collection[T]) LastOr(defaultValue T) T {
	if c.IsEmpty() {
		return defaultValue
	}
	return c.items[len(c.items)-1]
}

// LastWhere returns the last item matching the predicate.
func (c *Collection[T]) LastWhere(predicate func(T) bool) (T, bool) {
	for i := len(c.items) - 1; i >= 0; i-- {
		if predicate(c.items[i]) {
			return c.items[i], true
		}
	}
	var zero T
	return zero, false
}

// Get returns an item at the given index.
func (c *Collection[T]) Get(index int) T {
	if index < 0 || index >= len(c.items) {
		var zero T
		return zero
	}
	return c.items[index]
}

// GetOr returns an item at the given index or the default value.
func (c *Collection[T]) GetOr(index int, defaultValue T) T {
	if index < 0 || index >= len(c.items) {
		return defaultValue
	}
	return c.items[index]
}

// Each iterates over each item in the collection.
func (c *Collection[T]) Each(callback func(T, int)) *Collection[T] {
	for i, item := range c.items {
		callback(item, i)
	}
	return c
}

// EachSpread iterates over each item, spreading slice items to the callback.
func (c *Collection[T]) EachSpread(callback func(T, int) bool) *Collection[T] {
	for i, item := range c.items {
		if !callback(item, i) {
			break
		}
	}
	return c
}

// Map applies a callback to each item and returns a new collection.
func Map[T any, U any](c *Collection[T], callback func(T, int) U) *Collection[U] {
	result := make([]U, len(c.items))
	for i, item := range c.items {
		result[i] = callback(item, i)
	}
	return New(result)
}

// Filter returns a new collection with items that pass the predicate.
func (c *Collection[T]) Filter(predicate func(T) bool) *Collection[T] {
	result := make([]T, 0)
	for _, item := range c.items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return New(result)
}

// Reject returns a new collection without items that pass the predicate.
func (c *Collection[T]) Reject(predicate func(T) bool) *Collection[T] {
	return c.Filter(func(item T) bool {
		return !predicate(item)
	})
}

// Reduce reduces the collection to a single value.
func Reduce[T any, U any](c *Collection[T], callback func(U, T, int) U, initial U) U {
	result := initial
	for i, item := range c.items {
		result = callback(result, item, i)
	}
	return result
}

// Contains determines if an item exists in the collection.
func (c *Collection[T]) Contains(predicate func(T) bool) bool {
	for _, item := range c.items {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Every determines if all items pass the given truth test.
func (c *Collection[T]) Every(predicate func(T) bool) bool {
	for _, item := range c.items {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Some is an alias for Contains.
func (c *Collection[T]) Some(predicate func(T) bool) bool {
	return c.Contains(predicate)
}

// Push adds one or more items to the end of the collection.
func (c *Collection[T]) Push(items ...T) *Collection[T] {
	c.items = append(c.items, items...)
	return c
}

// Pop removes and returns the last item from the collection.
func (c *Collection[T]) Pop() T {
	if c.IsEmpty() {
		var zero T
		return zero
	}
	item := c.items[len(c.items)-1]
	c.items = c.items[:len(c.items)-1]
	return item
}

// Prepend adds one or more items to the beginning of the collection.
func (c *Collection[T]) Prepend(items ...T) *Collection[T] {
	c.items = append(items, c.items...)
	return c
}

// Shift removes and returns the first item from the collection.
func (c *Collection[T]) Shift() T {
	if c.IsEmpty() {
		var zero T
		return zero
	}
	item := c.items[0]
	c.items = c.items[1:]
	return item
}

// Slice returns a slice of the collection.
func (c *Collection[T]) Slice(offset int, length ...int) *Collection[T] {
	if offset < 0 {
		offset = max(0, len(c.items)+offset)
	}
	if offset >= len(c.items) {
		return Empty[T]()
	}

	end := len(c.items)
	if len(length) > 0 {
		end = min(offset+length[0], len(c.items))
	}

	return New(slices.Clone(c.items[offset:end]))
}

// Take returns a new collection with the first n items.
func (c *Collection[T]) Take(n int) *Collection[T] {
	if n < 0 {
		return c.Slice(n)
	}
	return c.Slice(0, n)
}

// Skip returns a new collection without the first n items.
func (c *Collection[T]) Skip(n int) *Collection[T] {
	return c.Slice(n)
}

// Chunk splits the collection into chunks of the given size.
func (c *Collection[T]) Chunk(size int) [][]*Collection[T] {
	if size <= 0 {
		return [][]*Collection[T]{}
	}

	chunks := make([]*Collection[T], 0)
	for i := 0; i < len(c.items); i += size {
		end := min(i+size, len(c.items))
		chunks = append(chunks, New(slices.Clone(c.items[i:end])))
	}
	return [][]*Collection[T]{chunks}
}

// ChunkInto splits the collection into chunks of the given size.
func (c *Collection[T]) ChunkInto(size int) []*Collection[T] {
	if size <= 0 {
		return []*Collection[T]{}
	}

	chunks := make([]*Collection[T], 0)
	for i := 0; i < len(c.items); i += size {
		end := min(i+size, len(c.items))
		chunks = append(chunks, New(slices.Clone(c.items[i:end])))
	}
	return chunks
}

// Split splits the collection into a given number of groups.
func (c *Collection[T]) Split(numberOfGroups int) []*Collection[T] {
	if c.IsEmpty() || numberOfGroups <= 0 {
		return []*Collection[T]{}
	}

	groups := make([]*Collection[T], 0)
	groupSize := len(c.items) / numberOfGroups
	remainder := len(c.items) % numberOfGroups

	start := 0
	for i := 0; i < numberOfGroups; i++ {
		size := groupSize
		if i < remainder {
			size++
		}
		if size > 0 {
			groups = append(groups, New(slices.Clone(c.items[start:start+size])))
			start += size
		}
	}
	return groups
}

// Reverse returns a new collection with items in reverse order.
func (c *Collection[T]) Reverse() *Collection[T] {
	result := make([]T, len(c.items))
	for i, item := range c.items {
		result[len(c.items)-1-i] = item
	}
	return New(result)
}

// Shuffle randomly shuffles the items in the collection.
func (c *Collection[T]) Shuffle() *Collection[T] {
	result := slices.Clone(c.items)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return New(result)
}

// Random returns a random item from the collection.
func (c *Collection[T]) Random() T {
	if c.IsEmpty() {
		var zero T
		return zero
	}
	return c.items[rand.Intn(len(c.items))]
}

// RandomN returns n random items from the collection.
func (c *Collection[T]) RandomN(n int) *Collection[T] {
	if n <= 0 || c.IsEmpty() {
		return Empty[T]()
	}
	if n >= len(c.items) {
		return c.Shuffle()
	}

	indices := rand.Perm(len(c.items))[:n]
	result := make([]T, n)
	for i, idx := range indices {
		result[i] = c.items[idx]
	}
	return New(result)
}

// Unique returns unique items from the collection using a key function.
func (c *Collection[T]) Unique(keyFn func(T) string) *Collection[T] {
	seen := make(map[string]bool)
	result := make([]T, 0)
	for _, item := range c.items {
		key := keyFn(item)
		if !seen[key] {
			seen[key] = true
			result = append(result, item)
		}
	}
	return New(result)
}

// Values resets the keys/indices and returns a new collection.
func (c *Collection[T]) Values() *Collection[T] {
	return New(slices.Clone(c.items))
}

// Merge merges other collections into this collection.
func (c *Collection[T]) Merge(others ...*Collection[T]) *Collection[T] {
	result := slices.Clone(c.items)
	for _, other := range others {
		result = append(result, other.items...)
	}
	return New(result)
}

// Concat adds items to the collection.
func (c *Collection[T]) Concat(items []T) *Collection[T] {
	result := slices.Clone(c.items)
	result = append(result, items...)
	return New(result)
}

// Clone returns a copy of the collection.
func (c *Collection[T]) Clone() *Collection[T] {
	return New(slices.Clone(c.items))
}

// Nth creates a new collection from every n-th item.
func (c *Collection[T]) Nth(step int, offset ...int) *Collection[T] {
	if step <= 0 {
		return Empty[T]()
	}

	start := 0
	if len(offset) > 0 {
		start = offset[0]
	}

	result := make([]T, 0)
	for i := start; i < len(c.items); i += step {
		result = append(result, c.items[i])
	}
	return New(result)
}

// ForPage returns a slice of items for the given page.
func (c *Collection[T]) ForPage(page, perPage int) *Collection[T] {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * perPage
	return c.Slice(offset, perPage)
}

// Pad pads the collection to the specified length with a value.
func (c *Collection[T]) Pad(size int, value T) *Collection[T] {
	currentLen := len(c.items)
	if currentLen >= abs(size) {
		return c.Clone()
	}

	newLen := abs(size)
	result := make([]T, newLen)

	if size > 0 {
		copy(result, c.items)
		for i := currentLen; i < newLen; i++ {
			result[i] = value
		}
	} else {
		padCount := newLen - currentLen
		for i := 0; i < padCount; i++ {
			result[i] = value
		}
		copy(result[padCount:], c.items)
	}
	return New(result)
}

// Tap passes the collection to a callback and returns the collection.
func (c *Collection[T]) Tap(callback func(*Collection[T])) *Collection[T] {
	callback(c)
	return c
}

// Pipe passes the collection to a callback and returns the result.
func Pipe[T any, U any](c *Collection[T], callback func(*Collection[T]) U) U {
	return callback(c)
}

// When applies the callback if the condition is true.
func (c *Collection[T]) When(condition bool, callback func(*Collection[T]) *Collection[T]) *Collection[T] {
	if condition {
		return callback(c)
	}
	return c
}

// WhenEmpty applies the callback if the collection is empty.
func (c *Collection[T]) WhenEmpty(callback func(*Collection[T]) *Collection[T]) *Collection[T] {
	return c.When(c.IsEmpty(), callback)
}

// WhenNotEmpty applies the callback if the collection is not empty.
func (c *Collection[T]) WhenNotEmpty(callback func(*Collection[T]) *Collection[T]) *Collection[T] {
	return c.When(c.IsNotEmpty(), callback)
}

// Unless applies the callback if the condition is false.
func (c *Collection[T]) Unless(condition bool, callback func(*Collection[T]) *Collection[T]) *Collection[T] {
	return c.When(!condition, callback)
}

// ContainsOneItem determines if the collection contains exactly one item.
func (c *Collection[T]) ContainsOneItem() bool {
	return c.Count() == 1
}

// TakeUntil takes items until the condition is met.
func (c *Collection[T]) TakeUntil(predicate func(T) bool) *Collection[T] {
	result := make([]T, 0)
	for _, item := range c.items {
		if predicate(item) {
			break
		}
		result = append(result, item)
	}
	return New(result)
}

// TakeWhile takes items while the condition is true.
func (c *Collection[T]) TakeWhile(predicate func(T) bool) *Collection[T] {
	result := make([]T, 0)
	for _, item := range c.items {
		if !predicate(item) {
			break
		}
		result = append(result, item)
	}
	return New(result)
}

// SkipUntil skips items until the condition is met.
func (c *Collection[T]) SkipUntil(predicate func(T) bool) *Collection[T] {
	result := make([]T, 0)
	skipping := true
	for _, item := range c.items {
		if skipping && predicate(item) {
			skipping = false
		}
		if !skipping {
			result = append(result, item)
		}
	}
	return New(result)
}

// SkipWhile skips items while the condition is true.
func (c *Collection[T]) SkipWhile(predicate func(T) bool) *Collection[T] {
	result := make([]T, 0)
	skipping := true
	for _, item := range c.items {
		if skipping && !predicate(item) {
			skipping = false
		}
		if !skipping {
			result = append(result, item)
		}
	}
	return New(result)
}

// Partition splits the collection into two based on a predicate.
func (c *Collection[T]) Partition(predicate func(T) bool) (*Collection[T], *Collection[T]) {
	pass := make([]T, 0)
	fail := make([]T, 0)
	for _, item := range c.items {
		if predicate(item) {
			pass = append(pass, item)
		} else {
			fail = append(fail, item)
		}
	}
	return New(pass), New(fail)
}

// Sole returns the only item in the collection, or error if not exactly one item.
func (c *Collection[T]) Sole() (T, error) {
	if c.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("item not found")
	}
	if c.Count() > 1 {
		var zero T
		return zero, fmt.Errorf("multiple items found")
	}
	return c.items[0], nil
}

// SoleWhere returns the only item matching the predicate.
func (c *Collection[T]) SoleWhere(predicate func(T) bool) (T, error) {
	filtered := c.Filter(predicate)
	return filtered.Sole()
}

// ToJSON converts the collection to JSON.
func (c *Collection[T]) ToJSON() ([]byte, error) {
	return json.Marshal(c.items)
}

// ToJSONString converts the collection to a JSON string.
func (c *Collection[T]) ToJSONString() string {
	data, err := c.ToJSON()
	if err != nil {
		return "[]"
	}
	return string(data)
}

// String returns a string representation of the collection.
func (c *Collection[T]) String() string {
	return c.ToJSONString()
}

// Dump prints the collection for debugging.
func (c *Collection[T]) Dump() *Collection[T] {
	fmt.Printf("%+v\n", c.items)
	return c
}

// DD dumps and dies (prints and panics).
func (c *Collection[T]) DD() {
	c.Dump()
	panic("DD called")
}

// Search finds the index of a value using a predicate.
func (c *Collection[T]) Search(predicate func(T) bool) int {
	for i, item := range c.items {
		if predicate(item) {
			return i
		}
	}
	return -1
}

// Transform applies a callback to each item, modifying in place.
func (c *Collection[T]) Transform(callback func(T, int) T) *Collection[T] {
	for i, item := range c.items {
		c.items[i] = callback(item, i)
	}
	return c
}

// Put sets the item at the given index.
func (c *Collection[T]) Put(index int, value T) *Collection[T] {
	if index >= 0 && index < len(c.items) {
		c.items[index] = value
	}
	return c
}

// Forget removes an item by index.
func (c *Collection[T]) Forget(indices ...int) *Collection[T] {
	// Sort indices in descending order to remove from end first
	sort.Sort(sort.Reverse(sort.IntSlice(indices)))
	for _, idx := range indices {
		if idx >= 0 && idx < len(c.items) {
			c.items = append(c.items[:idx], c.items[idx+1:]...)
		}
	}
	return c
}

// Pull gets and removes an item by index.
func (c *Collection[T]) Pull(index int) T {
	if index < 0 || index >= len(c.items) {
		var zero T
		return zero
	}
	item := c.items[index]
	c.Forget(index)
	return item
}

// Flip swaps keys and values (for string collections).
func Flip(c *Collection[string]) *MapCollection[string, int] {
	result := make(map[string]int)
	for i, item := range c.items {
		result[item] = i
	}
	return NewMap(result)
}

// Helper functions
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
