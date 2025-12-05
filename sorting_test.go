package collections_test

import (
	"testing"

	"github.com/qiuapeng921/collections"
)

func TestSortFunc(t *testing.T) {
	c := collections.Make(3, 1, 2)
	sorted := c.SortFunc(func(a, b int) int { return a - b })
	if sorted.Get(0) != 1 {
		t.Error("SortFunc failed")
	}
}

func TestSortStableFunc(t *testing.T) {
	c := collections.Make(3, 1, 2)
	sorted := c.SortStableFunc(func(a, b int) int { return a - b })
	if sorted.Get(0) != 1 {
		t.Error("SortStableFunc failed")
	}
}

func TestMinEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if collections.Min(c) != 0 {
		t.Error("Min on empty should return zero")
	}
}

func TestMaxEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if collections.Max(c) != 0 {
		t.Error("Max on empty should return zero")
	}
}

func TestMinByEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := collections.MinBy(c, func(n int) int { return n })
	if result != 0 {
		t.Error("MinBy on empty should return zero")
	}
}

func TestMaxByEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := collections.MaxBy(c, func(n int) int { return n })
	if result != 0 {
		t.Error("MaxBy on empty should return zero")
	}
}

func TestAvgEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if collections.Avg(c) != 0 {
		t.Error("Avg on empty should return 0")
	}
}

func TestAvgByEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if collections.AvgBy(c, func(n int) int { return n }) != 0 {
		t.Error("AvgBy on empty should return 0")
	}
}

func TestMedianEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if collections.Median(c) != 0 {
		t.Error("Median on empty should return 0")
	}
}

func TestModeEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if collections.Mode(c) != nil {
		t.Error("Mode on empty should return nil")
	}
}

func TestSortByKeys(t *testing.T) {
	type Item struct {
		Name string
		Age  int
	}
	items := collections.Make(
		Item{"B", 30},
		Item{"A", 25},
		Item{"A", 20},
	)
	sorted := collections.SortByKeys(items, []collections.SortKey[Item]{
		{KeyFn: func(i Item) any { return i.Name }, Descending: false},
		{KeyFn: func(i Item) any { return i.Age }, Descending: true},
	})
	if sorted.Get(0).Age != 25 {
		t.Error("SortByKeys failed")
	}
}

func TestSortByKeysAllTypes(t *testing.T) {
	// Test int8
	c1 := collections.Make(int8(3), int8(1), int8(2))
	collections.SortByKeys(c1, []collections.SortKey[int8]{{KeyFn: func(i int8) any { return i }}})

	// Test int16
	c2 := collections.Make(int16(3), int16(1))
	collections.SortByKeys(c2, []collections.SortKey[int16]{{KeyFn: func(i int16) any { return i }}})

	// Test int32
	c3 := collections.Make(int32(3), int32(1))
	collections.SortByKeys(c3, []collections.SortKey[int32]{{KeyFn: func(i int32) any { return i }}})

	// Test int64
	c4 := collections.Make(int64(3), int64(1))
	collections.SortByKeys(c4, []collections.SortKey[int64]{{KeyFn: func(i int64) any { return i }}})

	// Test uint
	c5 := collections.Make(uint(3), uint(1))
	collections.SortByKeys(c5, []collections.SortKey[uint]{{KeyFn: func(i uint) any { return i }}})

	// Test uint8
	c6 := collections.Make(uint8(3), uint8(1))
	collections.SortByKeys(c6, []collections.SortKey[uint8]{{KeyFn: func(i uint8) any { return i }}})

	// Test uint16
	c7 := collections.Make(uint16(3), uint16(1))
	collections.SortByKeys(c7, []collections.SortKey[uint16]{{KeyFn: func(i uint16) any { return i }}})

	// Test uint32
	c8 := collections.Make(uint32(3), uint32(1))
	collections.SortByKeys(c8, []collections.SortKey[uint32]{{KeyFn: func(i uint32) any { return i }}})

	// Test uint64
	c9 := collections.Make(uint64(3), uint64(1))
	collections.SortByKeys(c9, []collections.SortKey[uint64]{{KeyFn: func(i uint64) any { return i }}})

	// Test float32
	c10 := collections.Make(float32(3.0), float32(1.0))
	collections.SortByKeys(c10, []collections.SortKey[float32]{{KeyFn: func(i float32) any { return i }}})

	// Test float64
	c11 := collections.Make(3.0, 1.0)
	collections.SortByKeys(c11, []collections.SortKey[float64]{{KeyFn: func(i float64) any { return i }}})

	// Test string
	c12 := collections.Make("b", "a")
	collections.SortByKeys(c12, []collections.SortKey[string]{{KeyFn: func(i string) any { return i }}})

	// Test default case
	type Custom struct{ v int }
	c13 := collections.Make(Custom{1}, Custom{2})
	collections.SortByKeys(c13, []collections.SortKey[Custom]{{KeyFn: func(i Custom) any { return i }}})
}

func TestSortByKeysDescending(t *testing.T) {
	c := collections.Make(1, 3, 2)
	sorted := collections.SortByKeys(c, []collections.SortKey[int]{
		{KeyFn: func(i int) any { return i }, Descending: true},
	})
	if sorted.Get(0) != 3 {
		t.Error("SortByKeys descending failed")
	}
}
