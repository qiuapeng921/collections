// Package collections provides helper functions for collections.
package collections

// Value returns the value if it's not a function, or calls it if it is.
func Value[T any](value T) T {
	return value
}

// ValueFn returns the result of a function or the value itself.
func ValueFn[T any](value any) T {
	if fn, ok := value.(func() T); ok {
		return fn()
	}
	if v, ok := value.(T); ok {
		return v
	}
	var zero T
	return zero
}

// DataGet retrieves a value from nested data structures.
func DataGet(target any, key string, defaultValue ...any) any {
	if target == nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}

	if key == "" {
		return target
	}

	// Handle map[string]any
	if data, ok := target.(map[string]any); ok {
		return Arr.Get(data, key, defaultValue...)
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return nil
}

// DataSet sets a value in nested data structures.
func DataSet(target any, key string, value any) any {
	if data, ok := target.(map[string]any); ok {
		return Arr.Set(data, key, value)
	}
	return target
}

// DataForget removes a key from nested data structures.
func DataForget(target any, keys ...string) any {
	if data, ok := target.(map[string]any); ok {
		return Arr.Forget(data, keys...)
	}
	return target
}

// Collect creates a new collection from items.
func Collect[T any](items ...T) *Collection[T] {
	return New(items)
}

// CollectSlice creates a new collection from a slice.
func CollectSlice[T any](items []T) *Collection[T] {
	return New(items)
}

// CollectMap creates a new map collection from a map.
func CollectMap[K comparable, V any](items map[K]V) *MapCollection[K, V] {
	return NewMap(items)
}

// Head returns the first item from a slice.
func Head[T any](items []T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	return items[0]
}

// Tail returns all items except the first.
func Tail[T any](items []T) []T {
	if len(items) <= 1 {
		return []T{}
	}
	return items[1:]
}

// Init returns all items except the last.
func Init[T any](items []T) []T {
	if len(items) <= 1 {
		return []T{}
	}
	return items[:len(items)-1]
}

// LastItem returns the last item from a slice.
func LastItem[T any](items []T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	return items[len(items)-1]
}

// Blank returns true if the value is empty.
func Blank[T comparable](value T) bool {
	var zero T
	return value == zero
}

// Filled returns true if the value is not empty.
func Filled[T comparable](value T) bool {
	return !Blank(value)
}

// Transform applies a callback to a value if it's truthy.
func Transform[T comparable, U any](value T, callback func(T) U, defaultValue ...U) U {
	if Blank(value) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		var zero U
		return zero
	}
	return callback(value)
}

// With returns the value.
func With[T any](value T) T {
	return value
}

// Optional returns a wrapper that provides safe access to nullable values.
type Optional[T any] struct {
	value    T
	hasValue bool
}

// Some creates an Optional with a value.
func Some[T any](value T) Optional[T] {
	return Optional[T]{value: value, hasValue: true}
}

// None creates an empty Optional.
func None[T any]() Optional[T] {
	return Optional[T]{hasValue: false}
}

// HasValue returns true if the Optional has a value.
func (o Optional[T]) HasValue() bool {
	return o.hasValue
}

// Get returns the value or panics if empty.
func (o Optional[T]) Get() T {
	if !o.hasValue {
		panic("optional has no value")
	}
	return o.value
}

// GetOr returns the value or a default.
func (o Optional[T]) GetOr(defaultValue T) T {
	if !o.hasValue {
		return defaultValue
	}
	return o.value
}

// Map applies a function to the value if present.
func (o Optional[T]) Map(fn func(T) T) Optional[T] {
	if !o.hasValue {
		return o
	}
	return Some(fn(o.value))
}

// Filter returns None if the predicate fails.
func (o Optional[T]) Filter(predicate func(T) bool) Optional[T] {
	if !o.hasValue || !predicate(o.value) {
		return None[T]()
	}
	return o
}

// Retry retries a function n times until it succeeds.
func Retry[T any](times int, callback func() (T, error)) (T, error) {
	var lastErr error
	for i := 0; i < times; i++ {
		result, err := callback()
		if err == nil {
			return result, nil
		}
		lastErr = err
	}
	var zero T
	return zero, lastErr
}

// Once creates a function that only executes once.
func Once[T any](fn func() T) func() T {
	var result T
	var called bool
	return func() T {
		if !called {
			result = fn()
			called = true
		}
		return result
	}
}

// Identity returns the input unchanged.
func Identity[T any](value T) T {
	return value
}

// Tap executes a callback and returns the original value.
func Tap[T any](value T, callback func(T)) T {
	callback(value)
	return value
}

// Rescue executes a callback and returns a default on panic.
func Rescue[T any](callback func() T, defaultValue T) T {
	defer func() {
		recover()
	}()
	return callback()
}

// ThrowIf panics with an error if the condition is true.
func ThrowIf(condition bool, message string) {
	if condition {
		panic(message)
	}
}

// ThrowUnless panics with an error if the condition is false.
func ThrowUnless(condition bool, message string) {
	if !condition {
		panic(message)
	}
}
