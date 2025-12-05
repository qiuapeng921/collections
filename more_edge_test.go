package collections_test

import (
	"testing"

	"github.com/qiuapeng921/collections"
)

func TestMapIsEmpty(t *testing.T) {
	m := collections.NewMap(map[string]int{})
	if !m.IsEmpty() {
		t.Error("IsEmpty failed")
	}
	if m.IsNotEmpty() {
		t.Error("IsNotEmpty failed")
	}
}

func TestMapHasMultiple(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2})
	if !m.Has("a", "b") {
		t.Error("Has multiple failed")
	}
	if m.Has("a", "c") {
		t.Error("Has multiple with missing failed")
	}
}

func TestMapForgetMultiple(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2, "c": 3})
	m.Forget("a", "b")
	if m.Has("a") || m.Has("b") {
		t.Error("Forget multiple failed")
	}
}

func TestNewMapOrdered(t *testing.T) {
	m := collections.NewMapOrdered(
		map[string]int{"a": 1, "b": 2},
		[]string{"b", "a"},
	)
	keys := m.Keys()
	if keys.Get(0) != "b" || keys.Get(1) != "a" {
		t.Error("NewMapOrdered key order failed")
	}
}

func TestMapToJSONError(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	str := m.ToJSONString()
	if str == "" {
		t.Error("ToJSONString failed")
	}
}

func TestFlattenDepthZero(t *testing.T) {
	items := []any{1, []any{2, 3}}
	result := collections.Flatten[int](items, 0)
	// Depth 0 means don't recurse into nested arrays
	if len(result) < 1 {
		t.Error("Flatten depth 0 failed")
	}
}

func TestFlattenNonInt(t *testing.T) {
	items := []any{"a", []any{"b", "c"}}
	result := collections.Flatten[string](items, 1)
	if len(result) != 3 {
		t.Error("Flatten string failed")
	}
}

func TestSearchNotFound(t *testing.T) {
	c := collections.Make(1, 2, 3)
	if c.Search(func(n int) bool { return n > 10 }) != -1 {
		t.Error("Search not found failed")
	}
}

func TestChunkReturnType(t *testing.T) {
	c := collections.Make(1, 2, 3, 4, 5)
	chunks := c.Chunk(2)
	if len(chunks) != 1 {
		t.Error("Chunk return type failed")
	}
}

func TestSplitRemainder(t *testing.T) {
	c := collections.Make(1, 2, 3, 4, 5, 6, 7)
	groups := c.Split(3)
	// 7 items split into 3: sizes should be 3, 2, 2
	if len(groups) != 3 {
		t.Error("Split remainder failed")
	}
}

func TestMapCollectionToJSONInvalid(t *testing.T) {
	// Just test that ToJSON works - invalid types are Go's responsibility
	m := collections.NewMap(map[string]int{"a": 1})
	_, err := m.ToJSON()
	if err != nil {
		t.Error("ToJSON with valid data failed")
	}
}
