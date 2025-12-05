package collections_test

import (
	"testing"

	"github.com/qiuapeng921/collections"
)

func TestCollectionString(t *testing.T) {
	c := collections.Make(1, 2, 3)
	s := c.String()
	if s == "" {
		t.Error("String failed")
	}
}

func TestCollectionWhenEmpty(t *testing.T) {
	empty := collections.Empty[int]()
	result := empty.WhenEmpty(func(c *collections.Collection[int]) *collections.Collection[int] {
		return collections.Make(1, 2, 3)
	})
	if result.Count() != 3 {
		t.Error("WhenEmpty failed for empty")
	}

	notEmpty := collections.Make(1)
	result2 := notEmpty.WhenEmpty(func(c *collections.Collection[int]) *collections.Collection[int] {
		return collections.Make(4, 5, 6)
	})
	if result2.Count() != 1 {
		t.Error("WhenEmpty failed for non-empty")
	}
}

func TestCollectionWhenNotEmpty(t *testing.T) {
	notEmpty := collections.Make(1, 2, 3)
	result := notEmpty.WhenNotEmpty(func(c *collections.Collection[int]) *collections.Collection[int] {
		return c.Take(1)
	})
	if result.Count() != 1 {
		t.Error("WhenNotEmpty failed for non-empty")
	}

	empty := collections.Empty[int]()
	result2 := empty.WhenNotEmpty(func(c *collections.Collection[int]) *collections.Collection[int] {
		return collections.Make(4, 5, 6)
	})
	if result2.Count() != 0 {
		t.Error("WhenNotEmpty failed for empty")
	}
}

func TestTakeUntilEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := c.TakeUntil(func(n int) bool { return n > 1 })
	if result.Count() != 0 {
		t.Error("TakeUntil empty failed")
	}
}

func TestSkipUntilEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := c.SkipUntil(func(n int) bool { return n > 1 })
	if result.Count() != 0 {
		t.Error("SkipUntil empty failed")
	}
}

func TestSkipUntilNoMatch(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.SkipUntil(func(n int) bool { return n > 10 })
	if result.Count() != 0 {
		t.Error("SkipUntil no match failed")
	}
}

func TestCollectionSole(t *testing.T) {
	single := collections.Make(42)
	v, err := single.Sole()
	if err != nil || v != 42 {
		t.Error("Sole single failed")
	}

	multi := collections.Make(1, 2)
	_, err = multi.Sole()
	if err == nil {
		t.Error("Sole multi should fail")
	}

	empty := collections.Empty[int]()
	_, err = empty.Sole()
	if err == nil {
		t.Error("Sole empty should fail")
	}
}

func TestCollectionSoleWhere(t *testing.T) {
	c := collections.Make(1, 2, 3, 4)
	v, err := c.SoleWhere(func(n int) bool { return n == 3 })
	if err != nil || v != 3 {
		t.Error("SoleWhere found failed")
	}
}

func TestCombineUnequalLengths(t *testing.T) {
	keys := collections.Make("a", "b", "c")
	values := collections.Make(1, 2) // One less
	m := collections.Combine(keys, values)
	// Should handle gracefully
	if m == nil {
		t.Error("Combine unequal failed")
	}
}

func TestReplacePartial(t *testing.T) {
	c := collections.Make(1, 2, 3, 4, 5)
	result := c.Replace([]int{10, 20})
	if result.Get(0) != 10 || result.Get(2) != 3 {
		t.Error("Replace partial failed")
	}
}

func TestSpliceNoLength(t *testing.T) {
	c := collections.Make(1, 2, 3, 4, 5)
	remaining, removed := c.Splice(2)
	if remaining.Count() != 2 {
		t.Error("Splice no length remaining failed")
	}
	if removed.Count() != 3 {
		t.Error("Splice no length removed failed")
	}
}

func TestNewMapOrderedExtraKeys(t *testing.T) {
	m := collections.NewMapOrdered(
		map[string]int{"a": 1},
		[]string{"b", "c", "a"}, // Extra keys not in map
	)
	if m.Count() != 1 {
		t.Error("NewMapOrdered extra keys failed")
	}
}
