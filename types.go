package collections

// Numeric is a constraint for numeric types.
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Ordered is a constraint for ordered types.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// MapFunc is a function that maps T to U.
type MapFunc[T any, U any] func(T, int) U

// FilterFunc is a function that filters items.
type FilterFunc[T any] func(T) bool

// ReduceFunc is a function that reduces items.
type ReduceFunc[T any, U any] func(U, T, int) U

// CompareFunc is a function that compares two items.
type CompareFunc[T any] func(a, b T) int

// KeyFunc extracts a key from an item.
type KeyFunc[T any, K any] func(T) K
