package collections_test

import (
	"testing"

	"github.com/qiuapeng921/collections"
)

// Collection 边界情况测试

func TestFirstEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if c.First() != 0 {
		t.Error("First empty failed")
	}
}

func TestLastEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if c.Last() != 0 {
		t.Error("Last empty failed")
	}
}

func TestGetNegative(t *testing.T) {
	c := collections.Make(1, 2, 3)
	if c.Get(-1) != 0 {
		t.Error("Get negative failed")
	}
}

func TestPopEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if c.Pop() != 0 {
		t.Error("Pop empty failed")
	}
}

func TestShiftEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if c.Shift() != 0 {
		t.Error("Shift empty failed")
	}
}

func TestSliceOffsetBeyondLength(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.Slice(10)
	if result.Count() != 0 {
		t.Error("Slice offset beyond length failed")
	}
}

func TestChunkZeroSize(t *testing.T) {
	c := collections.Make(1, 2, 3)
	chunks := c.ChunkInto(0)
	if len(chunks) != 0 {
		t.Error("ChunkInto zero size failed")
	}
}

func TestChunkNegativeSize(t *testing.T) {
	c := collections.Make(1, 2, 3)
	chunks := c.Chunk(-1)
	if len(chunks) != 0 {
		t.Error("Chunk negative size failed")
	}
}

func TestSplitEmpty(t *testing.T) {
	c := collections.Empty[int]()
	groups := c.Split(3)
	if len(groups) != 0 {
		t.Error("Split empty failed")
	}
}

func TestSplitZero(t *testing.T) {
	c := collections.Make(1, 2, 3)
	groups := c.Split(0)
	if len(groups) != 0 {
		t.Error("Split zero failed")
	}
}

func TestRandomEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if c.Random() != 0 {
		t.Error("Random empty failed")
	}
}

func TestRandomNEmpty(t *testing.T) {
	c := collections.Empty[int]()
	result := c.RandomN(3)
	if result.Count() != 0 {
		t.Error("RandomN empty failed")
	}
}

func TestRandomNZero(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.RandomN(0)
	if result.Count() != 0 {
		t.Error("RandomN zero failed")
	}
}

func TestRandomNMoreThanLength(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.RandomN(10)
	if result.Count() != 3 {
		t.Error("RandomN more than length failed")
	}
}

func TestNthZero(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.Nth(0)
	if result.Count() != 0 {
		t.Error("Nth zero failed")
	}
}

func TestPadNoChange(t *testing.T) {
	c := collections.Make(1, 2, 3, 4, 5)
	result := c.Pad(3, 0)
	if result.Count() != 5 {
		t.Error("Pad no change failed")
	}
}

func TestPullInvalidIndex(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.Pull(-1)
	if result != 0 {
		t.Error("Pull invalid index failed")
	}
	result2 := c.Pull(100)
	if result2 != 0 {
		t.Error("Pull out of range failed")
	}
}

func TestPutInvalidIndex(t *testing.T) {
	c := collections.Make(1, 2, 3)
	c.Put(-1, 100)
	c.Put(100, 100)
	// No change expected
}

func TestForgetInvalidIndex(t *testing.T) {
	c := collections.Make(1, 2, 3)
	c.Forget(-1)
	c.Forget(100)
	if c.Count() != 3 {
		t.Error("Forget invalid index should not change count")
	}
}

func TestEachSpreadBreak(t *testing.T) {
	c := collections.Make(1, 2, 3, 4, 5)
	count := 0
	c.EachSpread(func(n int, i int) bool {
		count++
		return count < 3
	})
	if count != 3 {
		t.Error("EachSpread break failed")
	}
}

func TestSoleWhereNone(t *testing.T) {
	c := collections.Make(1, 2, 3)
	_, err := c.SoleWhere(func(n int) bool { return n > 10 })
	if err == nil {
		t.Error("SoleWhere none should return error")
	}
}

func TestSoleWhereMultiple(t *testing.T) {
	c := collections.Make(1, 2, 3)
	_, err := c.SoleWhere(func(n int) bool { return n > 0 })
	if err == nil {
		t.Error("SoleWhere multiple should return error")
	}
}

func TestFlip(t *testing.T) {
	c := collections.Make("a", "b", "c")
	flipped := collections.Flip(c)
	if flipped.Get("a") != 0 || flipped.Get("b") != 1 || flipped.Get("c") != 2 {
		t.Error("Flip failed")
	}
}

func TestToJSONError(t *testing.T) {
	// Test with valid data
	c := collections.Make(1, 2, 3)
	_, err := c.ToJSON()
	if err != nil {
		t.Error("ToJSON failed")
	}
}

func TestDump(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := c.Dump() // Just verify it doesn't panic
	if result == nil {
		t.Error("Dump failed")
	}
}

func TestMapDump(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	result := m.Dump()
	if result == nil {
		t.Error("Map Dump failed")
	}
}

func TestForPageZero(t *testing.T) {
	c := collections.Make(1, 2, 3, 4, 5)
	result := c.ForPage(0, 2)
	if result.First() != 1 {
		t.Error("ForPage zero should treat as page 1")
	}
}

func TestLastWhereNone(t *testing.T) {
	c := collections.Make(1, 2, 3)
	_, found := c.LastWhere(func(n int) bool { return n > 10 })
	if found {
		t.Error("LastWhere none should not find")
	}
}

func TestEveryEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if !c.Every(func(n int) bool { return n > 0 }) {
		t.Error("Every empty should return true")
	}
}

func TestPipe(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := collections.Pipe(c, func(col *collections.Collection[int]) int {
		return col.Count()
	})
	if result != 3 {
		t.Error("Pipe failed")
	}
}
