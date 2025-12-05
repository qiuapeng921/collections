package collections_test

import (
	"testing"

	"github.com/qiuapeng921/collections"
)

func TestMapCollectionAll(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	if len(m.All()) != 1 {
		t.Error("All failed")
	}
}

func TestMapCollectionPull(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2})
	v := m.Pull("a")
	if v != 1 || m.Has("a") {
		t.Error("Pull failed")
	}
}

func TestMapCollectionEach(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	count := 0
	m.Each(func(k string, v int) { count++ })
	if count != 1 {
		t.Error("Each failed")
	}
}

func TestMapCollectionEachBreak(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2, "c": 3})
	count := 0
	m.EachBreak(func(k string, v int) bool {
		count++
		return count < 2
	})
	if count != 2 {
		t.Error("EachBreak failed")
	}
}

func TestMapValues(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	doubled := collections.MapValues(m, func(v int, k string) int { return v * 2 })
	if doubled.Get("a") != 2 {
		t.Error("MapValues failed")
	}
}

func TestMapCollectionReject(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2})
	rejected := m.Reject(func(v int, k string) bool { return v > 1 })
	if rejected.Count() != 1 {
		t.Error("Reject failed")
	}
}

func TestMapCollectionUnion(t *testing.T) {
	m1 := collections.NewMap(map[string]int{"a": 1})
	m2 := collections.NewMap(map[string]int{"a": 2, "b": 3})
	union := m1.Union(m2)
	if union.Get("a") != 1 || union.Get("b") != 3 {
		t.Error("Union failed")
	}
}

func TestMapCollectionDiffKeys(t *testing.T) {
	m1 := collections.NewMap(map[string]int{"a": 1, "b": 2})
	m2 := collections.NewMap(map[string]int{"a": 1})
	diff := m1.DiffKeys(m2)
	if diff.Count() != 1 || !diff.Has("b") {
		t.Error("DiffKeys failed")
	}
}

func TestMapCollectionIntersectByKeys(t *testing.T) {
	m1 := collections.NewMap(map[string]int{"a": 1, "b": 2})
	m2 := collections.NewMap(map[string]int{"a": 1, "c": 3})
	inter := m1.IntersectByKeys(m2)
	if inter.Count() != 1 || !inter.Has("a") {
		t.Error("IntersectByKeys failed")
	}
}

func TestMapCollectionLastKey(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	key := m.LastKey()
	if key == "" {
		t.Error("LastKey failed")
	}
}

func TestMapCollectionLastKeyEmpty(t *testing.T) {
	m := collections.NewMap(map[string]int{})
	key := m.LastKey()
	if key != "" {
		t.Error("LastKey on empty should return zero")
	}
}

func TestMapCollectionFirstKeyEmpty(t *testing.T) {
	m := collections.NewMap(map[string]int{})
	key := m.FirstKey()
	if key != "" {
		t.Error("FirstKey on empty should return zero")
	}
}

func TestMapCollectionFirstEmpty(t *testing.T) {
	m := collections.NewMap(map[string]int{})
	v := m.First()
	if v != 0 {
		t.Error("First on empty should return zero")
	}
}

func TestMapCollectionLastEmpty(t *testing.T) {
	m := collections.NewMap(map[string]int{})
	v := m.Last()
	if v != 0 {
		t.Error("Last on empty should return zero")
	}
}

func TestMapCollectionToJSON(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	_, err := m.ToJSON()
	if err != nil {
		t.Error("ToJSON failed")
	}
}

func TestMapCollectionToJSONString(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	s := m.ToJSONString()
	if s == "" || s == "{}" {
		t.Error("ToJSONString failed")
	}
}

func TestMapCollectionString(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	s := m.String()
	if s == "" {
		t.Error("String failed")
	}
}

func TestMapCollectionReduceMap(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2})
	sum := collections.ReduceMap(m, func(acc int, v int, k string) int { return acc + v }, 0)
	if sum != 3 {
		t.Error("ReduceMap failed")
	}
}

func TestSortMapKeys(t *testing.T) {
	m := collections.NewMap(map[string]int{"c": 3, "a": 1, "b": 2})
	sorted := collections.SortMapKeys(m)
	keys := sorted.Keys()
	if keys.Get(0) != "a" {
		t.Error("SortMapKeys failed")
	}
}

func TestSortMapKeysDesc(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2, "c": 3})
	sorted := collections.SortMapKeysDesc(m)
	keys := sorted.Keys()
	if keys.Get(0) != "c" {
		t.Error("SortMapKeysDesc failed")
	}
}

func TestMapCollectionGetOrPut(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	v := m.GetOrPut("b", 2)
	if v != 2 || m.Get("b") != 2 {
		t.Error("GetOrPut new key failed")
	}
	v2 := m.GetOrPut("a", 10)
	if v2 != 1 {
		t.Error("GetOrPut existing key failed")
	}
}

func TestMapCollectionTap(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	called := false
	m.Tap(func(mc *collections.MapCollection[string, int]) { called = true })
	if !called {
		t.Error("Tap failed")
	}
}

func TestMapCollectionWhen(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	result := m.When(true, func(mc *collections.MapCollection[string, int]) *collections.MapCollection[string, int] {
		mc.Put("b", 2)
		return mc
	})
	if !result.Has("b") {
		t.Error("When true failed")
	}
	result2 := m.When(false, func(mc *collections.MapCollection[string, int]) *collections.MapCollection[string, int] {
		mc.Put("c", 3)
		return mc
	})
	if result2.Has("c") {
		t.Error("When false failed")
	}
}

func TestFlipMap(t *testing.T) {
	m := collections.NewMap(map[string]string{"a": "x", "b": "y"})
	flipped := collections.FlipMap(m)
	if flipped.Get("x") != "a" {
		t.Error("FlipMap failed")
	}
}

func TestKeyValue(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2})
	slice := m.ToSlice()
	if slice.Count() != 2 {
		t.Error("ToSlice failed")
	}
}

func TestFromSlice(t *testing.T) {
	slice := []collections.KeyValue[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
	}
	m := collections.FromSlice(slice)
	if m.Get("a") != 1 {
		t.Error("FromSlice failed")
	}
}

// 额外覆盖率测试

func TestNewMapNil(t *testing.T) {
	m := collections.NewMap[string, int](nil)
	if m.Count() != 0 {
		t.Error("NewMap nil should create empty map")
	}
}

func TestNewMapEmpty(t *testing.T) {
	emptyMap := collections.NewMap(map[string]int{})
	if emptyMap.Count() != 0 {
		t.Error("Empty map should have 0 count")
	}
}

func TestNewMapOrderedNilKeys(t *testing.T) {
	m := collections.NewMapOrdered(map[string]int{"a": 1}, nil)
	if m.Count() != 1 {
		t.Error("NewMapOrdered should still work with nil keys")
	}
}

func TestNewMapOrderedEmptyKeys(t *testing.T) {
	orderedEmpty := collections.NewMapOrdered(map[string]int{"a": 1}, []string{})
	if orderedEmpty.Count() != 1 {
		t.Error("Should still have items")
	}
}

func TestNewMapOrderedWithKeys(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := []string{"c", "a", "b"}
	m := collections.NewMapOrdered(data, keys)
	orderedKeys := m.Keys().All()
	if len(orderedKeys) != 3 || orderedKeys[0] != "c" {
		t.Error("NewMapOrdered should respect key order")
	}
}

func TestMapCollectionLastNonEmpty(t *testing.T) {
	m := collections.NewMapOrdered(map[string]int{"a": 1, "b": 2}, []string{"a", "b"})
	last := m.Last()
	if last != 2 {
		t.Error("Last should return last value")
	}
}

func TestMapCollectionFirstKeyNonEmpty(t *testing.T) {
	m := collections.NewMapOrdered(map[string]int{"a": 1, "b": 2}, []string{"a", "b"})
	firstKey := m.FirstKey()
	if firstKey != "a" {
		t.Error("FirstKey should return first key")
	}
}

func TestMapCollectionContainsTrue(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2})
	if !m.Contains(func(v int, k string) bool { return v == 1 }) {
		t.Error("Contains should return true")
	}
}

func TestMapCollectionContainsFalse(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2})
	if m.Contains(func(v int, k string) bool { return v == 100 }) {
		t.Error("Contains should return false")
	}
}

func TestMapCollectionEveryTrue(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2})
	if !m.Every(func(v int, k string) bool { return v > 0 }) {
		t.Error("Every should return true")
	}
}

func TestMapCollectionEveryFalse(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2})
	if m.Every(func(v int, k string) bool { return v > 1 }) {
		t.Error("Every should return false")
	}
}
