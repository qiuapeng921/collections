package collections

import "fmt"

// ItemNotFoundException is returned when an item is not found.
type ItemNotFoundException struct {
	Message string
}

func (e *ItemNotFoundException) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "item not found"
}

// MultipleItemsFoundException is returned when multiple items are found but only one expected.
type MultipleItemsFoundException struct {
	Message string
}

func (e *MultipleItemsFoundException) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "multiple items found"
}

// InvalidArgumentException is returned for invalid arguments.
type InvalidArgumentException struct {
	Message string
}

func (e *InvalidArgumentException) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "invalid argument"
}

// FirstOrFail returns the first item or returns an error if empty.
func (c *Collection[T]) FirstOrFail() (T, error) {
	if c.IsEmpty() {
		var zero T
		return zero, &ItemNotFoundException{}
	}
	return c.items[0], nil
}

// FirstWhereOrFail returns the first matching item or an error.
func (c *Collection[T]) FirstWhereOrFail(predicate func(T) bool) (T, error) {
	for _, item := range c.items {
		if predicate(item) {
			return item, nil
		}
	}
	var zero T
	return zero, &ItemNotFoundException{}
}

// LastOrFail returns the last item or an error if empty.
func (c *Collection[T]) LastOrFail() (T, error) {
	if c.IsEmpty() {
		var zero T
		return zero, &ItemNotFoundException{}
	}
	return c.items[len(c.items)-1], nil
}

// GetOrFail returns an item at the given index or an error.
func (c *Collection[T]) GetOrFail(index int) (T, error) {
	if index < 0 || index >= len(c.items) {
		var zero T
		return zero, &ItemNotFoundException{Message: fmt.Sprintf("index %d out of range", index)}
	}
	return c.items[index], nil
}

// RandomOrFail returns a random item or an error if empty.
func (c *Collection[T]) RandomOrFail() (T, error) {
	if c.IsEmpty() {
		var zero T
		return zero, &ItemNotFoundException{}
	}
	return c.Random(), nil
}

// PopOrFail removes and returns the last item or an error.
func (c *Collection[T]) PopOrFail() (T, error) {
	if c.IsEmpty() {
		var zero T
		return zero, &ItemNotFoundException{}
	}
	return c.Pop(), nil
}

// ShiftOrFail removes and returns the first item or an error.
func (c *Collection[T]) ShiftOrFail() (T, error) {
	if c.IsEmpty() {
		var zero T
		return zero, &ItemNotFoundException{}
	}
	return c.Shift(), nil
}
