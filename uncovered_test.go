package collections_test

import (
	"testing"

	"github.com/qiuapeng921/collections"
)

// 测试未覆盖的路径

func TestContainsEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if c.Contains(func(n int) bool { return n == 1 }) {
		t.Error("Contains empty should return false")
	}
}

func TestSomeEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if c.Some(func(n int) bool { return n == 1 }) {
		t.Error("Some empty should return false")
	}
}

func TestMapPutExisting(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	m.Put("a", 2)
	if m.Get("a") != 2 {
		t.Error("Put existing failed")
	}
}

func TestMapMergeEmpty(t *testing.T) {
	m1 := collections.NewMap(map[string]int{"a": 1})
	merged := m1.Merge()
	if merged.Count() != 1 {
		t.Error("Merge empty failed")
	}
}

func TestMapUnionEmpty(t *testing.T) {
	m1 := collections.NewMap(map[string]int{"a": 1})
	m2 := collections.NewMap(map[string]int{})
	union := m1.Union(m2)
	if union.Count() != 1 {
		t.Error("Union with empty failed")
	}
}

func TestCrossJoinSingle(t *testing.T) {
	a := collections.Make(1, 2)
	result := collections.CrossJoin(a)
	if result.Count() != 2 {
		t.Error("CrossJoin single failed")
	}
}

func TestReplaceEmpty(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.Replace([]int{})
	if result.Count() != 3 {
		t.Error("Replace empty failed")
	}
}

func TestReplaceLonger(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.Replace([]int{10, 20, 30, 40, 50})
	// Replace replaces indices 0,1,2 but doesn't extend
	if result.Count() < 3 {
		t.Error("Replace longer failed")
	}
}

func TestImplodeEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := collections.Implode(c, func(n int) string { return "" }, ", ")
	if result != "" {
		t.Error("Implode empty failed")
	}
}

func TestImplodeSingle(t *testing.T) {
	c := collections.Make(1)
	result := collections.Implode(c, func(n int) string { return "x" }, ", ")
	if result != "x" {
		t.Error("Implode single failed")
	}
}

func TestJoinStringsWithFinalGlue(t *testing.T) {
	c := collections.Make("a", "b", "c")
	result := collections.JoinStrings(c, ", ", " and ")
	if result != "a, b and c" {
		t.Error("JoinStrings with final glue failed")
	}
}

func TestSkipWhileNoMatch(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.SkipWhile(func(n int) bool { return n < 10 })
	if result.Count() != 0 {
		t.Error("SkipWhile no match failed")
	}
}

func TestTakeWhileAll(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.TakeWhile(func(n int) bool { return n < 10 })
	if result.Count() != 3 {
		t.Error("TakeWhile all failed")
	}
}

func TestKeyByDuplicate(t *testing.T) {
	type Item struct {
		Group int
		Name  string
	}
	items := collections.Make(Item{1, "a"}, Item{1, "b"})
	m := collections.KeyBy(items, func(i Item) int { return i.Group })
	// Last one wins
	if m.Get(1).Name != "b" {
		t.Error("KeyBy duplicate failed")
	}
}

func TestGroupByEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := collections.GroupBy(c, func(n int) int { return n })
	if result.Count() != 0 {
		t.Error("GroupBy empty failed")
	}
}

func TestCountByEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := collections.CountBy(c, func(n int) int { return n })
	if result.Count() != 0 {
		t.Error("CountBy empty failed")
	}
}

func TestPluckMapEmpty(t *testing.T) {
	type Item struct{ ID int }
	c := collections.Empty[Item]()
	result := collections.PluckMap(c, func(i Item) int { return i.ID }, func(i Item) int { return i.ID })
	if result.Count() != 0 {
		t.Error("PluckMap empty failed")
	}
}

func TestMapWithKeysEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := collections.MapWithKeys(c, func(n int, i int) (int, int) { return i, n })
	if result.Count() != 0 {
		t.Error("MapWithKeys empty failed")
	}
}

func TestDiffEmpty(t *testing.T) {
	c := collections.Empty[int]()
	other := collections.Make(1, 2, 3)
	result := collections.Diff(c, other)
	if result.Count() != 0 {
		t.Error("Diff empty failed")
	}
}

func TestIntersectEmpty(t *testing.T) {
	c := collections.Empty[int]()
	other := collections.Make(1, 2, 3)
	result := collections.Intersect(c, other)
	if result.Count() != 0 {
		t.Error("Intersect empty failed")
	}
}

func TestDuplicatesEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := collections.Duplicates(c)
	if result.Count() != 0 {
		t.Error("Duplicates empty failed")
	}
}

func TestDuplicatesNone(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := collections.Duplicates(c)
	if result.Count() != 0 {
		t.Error("Duplicates none failed")
	}
}

func TestUniqueComparableEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := collections.UniqueComparable(c)
	if result.Count() != 0 {
		t.Error("UniqueComparable empty failed")
	}
}

func TestFlipEmpty(t *testing.T) {
	c := collections.Empty[string]()
	result := collections.Flip(c)
	if result.Count() != 0 {
		t.Error("Flip empty failed")
	}
}
