package collections

import (
	"strings"
)

// ArrHelpers provides array helper functions similar to Laravel's Arr class.
type ArrHelpers struct{}

var Arr = ArrHelpers{}

// Get retrieves a value from a nested map using dot notation.
func (ArrHelpers) Get(data map[string]any, key string, defaultValue ...any) any {
	if key == "" {
		return data
	}

	// Check for direct key
	if val, ok := data[key]; ok {
		return val
	}

	// Check for dot notation
	if !strings.Contains(key, ".") {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}

	keys := strings.Split(key, ".")
	current := any(data)

	for _, k := range keys {
		switch v := current.(type) {
		case map[string]any:
			if val, ok := v[k]; ok {
				current = val
			} else {
				if len(defaultValue) > 0 {
					return defaultValue[0]
				}
				return nil
			}
		default:
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
			return nil
		}
	}

	return current
}

// Set sets a value in a nested map using dot notation.
func (ArrHelpers) Set(data map[string]any, key string, value any) map[string]any {
	if key == "" {
		return data
	}

	keys := strings.Split(key, ".")
	current := data

	for i, k := range keys {
		if i == len(keys)-1 {
			current[k] = value
		} else {
			if _, ok := current[k]; !ok {
				current[k] = make(map[string]any)
			}
			if next, ok := current[k].(map[string]any); ok {
				current = next
			} else {
				current[k] = make(map[string]any)
				current = current[k].(map[string]any)
			}
		}
	}

	return data
}

// Has checks if a key exists using dot notation.
func (ArrHelpers) Has(data map[string]any, keys ...string) bool {
	for _, key := range keys {
		if Arr.Get(data, key) == nil {
			return false
		}
	}
	return true
}

// HasAny checks if any key exists.
func (ArrHelpers) HasAny(data map[string]any, keys ...string) bool {
	for _, key := range keys {
		if Arr.Get(data, key) != nil {
			return true
		}
	}
	return false
}

// Forget removes keys from a map using dot notation.
func (ArrHelpers) Forget(data map[string]any, keys ...string) map[string]any {
	for _, key := range keys {
		if !strings.Contains(key, ".") {
			delete(data, key)
			continue
		}

		parts := strings.Split(key, ".")
		current := data

		for i, part := range parts {
			if i == len(parts)-1 {
				delete(current, part)
			} else {
				if next, ok := current[part].(map[string]any); ok {
					current = next
				} else {
					break
				}
			}
		}
	}
	return data
}

// Dot flattens a multi-dimensional map into dot notation.
func (ArrHelpers) Dot(data map[string]any, prepend ...string) map[string]any {
	result := make(map[string]any)
	prefix := ""
	if len(prepend) > 0 {
		prefix = prepend[0]
	}

	for key, value := range data {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		if nested, ok := value.(map[string]any); ok && len(nested) > 0 {
			for k, v := range Arr.Dot(nested, fullKey) {
				result[k] = v
			}
		} else {
			result[fullKey] = value
		}
	}

	return result
}

// Undot expands dot notation keys into a nested map.
func (ArrHelpers) Undot(data map[string]any) map[string]any {
	result := make(map[string]any)

	for key, value := range data {
		Arr.Set(result, key, value)
	}

	return result
}

// Only returns only the specified keys.
func (ArrHelpers) Only(data map[string]any, keys ...string) map[string]any {
	result := make(map[string]any)
	for _, key := range keys {
		if val, ok := data[key]; ok {
			result[key] = val
		}
	}
	return result
}

// Except returns all keys except the specified ones.
func (ArrHelpers) Except(data map[string]any, keys ...string) map[string]any {
	result := make(map[string]any)
	keySet := make(map[string]bool)
	for _, k := range keys {
		keySet[k] = true
	}

	for key, value := range data {
		if !keySet[key] {
			result[key] = value
		}
	}
	return result
}

// Add adds a key-value if the key doesn't exist.
func (ArrHelpers) Add(data map[string]any, key string, value any) map[string]any {
	if Arr.Get(data, key) == nil {
		Arr.Set(data, key, value)
	}
	return data
}

// Pull gets and removes a key.
func (ArrHelpers) Pull(data map[string]any, key string, defaultValue ...any) any {
	value := Arr.Get(data, key, defaultValue...)
	Arr.Forget(data, key)
	return value
}

// Exists checks if a key exists at the top level.
func (ArrHelpers) Exists(data map[string]any, key string) bool {
	_, ok := data[key]
	return ok
}

// Wrap wraps a value in a slice if it's not already one.
func (ArrHelpers) Wrap(value any) []any {
	if value == nil {
		return []any{}
	}
	if slice, ok := value.([]any); ok {
		return slice
	}
	return []any{value}
}

// First returns the first element.
func (ArrHelpers) First(items []any, predicate ...func(any) bool) any {
	if len(items) == 0 {
		return nil
	}

	if len(predicate) == 0 || predicate[0] == nil {
		return items[0]
	}

	for _, item := range items {
		if predicate[0](item) {
			return item
		}
	}
	return nil
}

// Last returns the last element.
func (ArrHelpers) Last(items []any, predicate ...func(any) bool) any {
	if len(items) == 0 {
		return nil
	}

	if len(predicate) == 0 || predicate[0] == nil {
		return items[len(items)-1]
	}

	for i := len(items) - 1; i >= 0; i-- {
		if predicate[0](items[i]) {
			return items[i]
		}
	}
	return nil
}

// Where filters items by a predicate.
func (ArrHelpers) Where(items []any, predicate func(any, any) bool) []any {
	result := make([]any, 0)
	for i, item := range items {
		if predicate(item, i) {
			result = append(result, item)
		}
	}
	return result
}

// WhereNotNull filters out nil values.
func (ArrHelpers) WhereNotNull(items []any) []any {
	return Arr.Where(items, func(item any, _ any) bool {
		return item != nil
	})
}

// Shuffle randomly shuffles items.
func (ah ArrHelpers) Shuffle(items []any) []any {
	c := New(items)
	return c.Shuffle().All()
}

// Random returns random items.
func (ah ArrHelpers) Random(items []any, n ...int) any {
	c := New(items)
	if len(n) == 0 {
		return c.Random()
	}
	return c.RandomN(n[0]).All()
}

// Divide splits a map into keys and values.
func (ArrHelpers) Divide(data map[string]any) ([]string, []any) {
	keys := make([]string, 0, len(data))
	values := make([]any, 0, len(data))
	for k, v := range data {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

// IsAssoc checks if a slice has string keys (always false for Go slices).
func (ArrHelpers) IsAssoc(items []any) bool {
	return false // Go slices are always numerically indexed
}

// IsList checks if a slice is a list.
func (ArrHelpers) IsList(items []any) bool {
	return true // Go slices are always lists
}

// Query builds a query string from a map.
func (ArrHelpers) Query(data map[string]string) string {
	parts := make([]string, 0, len(data))
	for k, v := range data {
		parts = append(parts, k+"="+v)
	}
	return strings.Join(parts, "&")
}

// Accessible returns true (Go slices are always accessible).
func (ArrHelpers) Accessible(value any) bool {
	return true
}

// Collapse flattens an array of arrays.
func (ArrHelpers) Collapse(items [][]any) []any {
	result := make([]any, 0)
	for _, arr := range items {
		result = append(result, arr...)
	}
	return result
}

// Prepend adds an item to the beginning.
func (ArrHelpers) Prepend(items []any, value any) []any {
	return append([]any{value}, items...)
}

// CrossJoin creates a cross product.
func (ArrHelpers) CrossJoin(arrays ...[]any) [][]any {
	if len(arrays) == 0 {
		return [][]any{}
	}

	result := [][]any{{}}
	for _, arr := range arrays {
		newResult := make([][]any, 0)
		for _, existing := range result {
			for _, item := range arr {
				newRow := make([]any, len(existing)+1)
				copy(newRow, existing)
				newRow[len(existing)] = item
				newResult = append(newResult, newRow)
			}
		}
		result = newResult
	}
	return result
}
