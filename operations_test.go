package collections_test

import (
	"errors"
	"testing"

	"github.com/qiuapeng921/collections"
)

func TestPluckMap(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}
	users := collections.Make(User{1, "A"}, User{2, "B"})
	m := collections.PluckMap(users, func(u User) string { return u.Name }, func(u User) int { return u.ID })
	if m.Get(1) != "A" {
		t.Error("PluckMap failed")
	}
}

func TestCollapse(t *testing.T) {
	c1 := collections.Make(1, 2)
	c2 := collections.Make(3, 4)
	outer := collections.Make(c1, c2)
	result := collections.Collapse(outer)
	if result.Count() != 4 {
		t.Error("Collapse failed")
	}
}

func TestCollapseSlices(t *testing.T) {
	c := collections.Make([]int{1, 2}, []int{3, 4})
	result := collections.CollapseSlices(c)
	if result.Count() != 4 {
		t.Error("CollapseSlices failed")
	}
}

func TestFlatten(t *testing.T) {
	items := []any{1, []any{2, 3}, []any{4, []any{5}}}
	result := collections.Flatten[int](items, 2)
	if len(result) != 5 {
		t.Error("Flatten failed")
	}
}

func TestMapWithKeys(t *testing.T) {
	c := collections.Make("a", "b")
	m := collections.MapWithKeys(c, func(s string, i int) (int, string) { return i, s })
	if m.Get(0) != "a" {
		t.Error("MapWithKeys failed")
	}
}

func TestMapToDictionary(t *testing.T) {
	c := collections.Make(1, 2, 3, 4)
	m := collections.MapToDictionary(c, func(n int, i int) (string, int) {
		if n%2 == 0 {
			return "even", n
		}
		return "odd", n
	})
	if len(m.Get("even")) != 2 {
		t.Error("MapToDictionary failed")
	}
}

func TestMapToGroups(t *testing.T) {
	c := collections.Make(1, 2)
	m := collections.MapToGroups(c, func(n int, i int) (string, int) { return "all", n })
	if len(m.Get("all")) != 2 {
		t.Error("MapToGroups failed")
	}
}

func TestZipEmpty(t *testing.T) {
	result := collections.Zip[int]()
	if result.Count() != 0 {
		t.Error("Zip empty failed")
	}
}

func TestZipDifferentLengths(t *testing.T) {
	a := collections.Make(1, 2, 3)
	b := collections.Make(4, 5)
	result := collections.Zip(a, b)
	if result.Count() != 2 {
		t.Error("Zip different lengths failed")
	}
}

func TestCrossJoinEmpty(t *testing.T) {
	result := collections.CrossJoin[int]()
	if result.Count() != 0 {
		t.Error("CrossJoin empty failed")
	}
}

func TestSpliceNegativeOffset(t *testing.T) {
	c := collections.Make(1, 2, 3, 4, 5)
	_, removed := c.Splice(-2, 2)
	if removed.Count() != 2 {
		t.Error("Splice negative offset failed")
	}
}

func TestSpliceReplace(t *testing.T) {
	c := collections.Make(1, 2, 3, 4, 5)
	_, removed := c.SpliceReplace(1, 2, []int{10, 20, 30})
	if removed.Count() != 2 {
		t.Error("SpliceReplace failed")
	}
}

func TestSpliceReplaceNegative(t *testing.T) {
	c := collections.Make(1, 2, 3, 4, 5)
	c.SpliceReplace(-2, 1, []int{100})
}

func TestJoinStringsEmpty(t *testing.T) {
	c := collections.Empty[string]()
	if collections.JoinStrings(c, ", ") != "" {
		t.Error("JoinStrings empty failed")
	}
}

func TestJoinStringsSingle(t *testing.T) {
	c := collections.Make("hello")
	if collections.JoinStrings(c, ", ") != "hello" {
		t.Error("JoinStrings single failed")
	}
}

func TestJoinEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if collections.Join(c, func(n int) string { return "" }, ", ") != "" {
		t.Error("Join empty failed")
	}
}

func TestJoinSingle(t *testing.T) {
	c := collections.Make(1)
	result := collections.Join(c, func(n int) string { return "x" }, ", ")
	if result != "x" {
		t.Error("Join single failed")
	}
}

func TestJoinWithFinalGlue(t *testing.T) {
	c := collections.Make(1, 2, 3)
	result := collections.Join(c, func(n int) string { return "x" }, ", ", " and ")
	if result != "x, x and x" {
		t.Error("Join with final glue failed")
	}
}

func TestSlidingEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if len(c.Sliding(2)) != 0 {
		t.Error("Sliding empty failed")
	}
}

func TestSlidingInvalidSize(t *testing.T) {
	c := collections.Make(1, 2, 3)
	if len(c.Sliding(0)) != 0 {
		t.Error("Sliding invalid size failed")
	}
}

func TestWhereType(t *testing.T) {
	c := collections.Make[any](1, "hello", 2, "world", 3)
	result := collections.WhereType[any, int](c)
	if result.Count() != 3 {
		t.Error("WhereType failed")
	}
}

func TestToCollection(t *testing.T) {
	items := []int{1, 2, 3}
	c := collections.ToCollection(items)
	if c.Count() != 3 {
		t.Error("ToCollection failed")
	}
}

// Helper tests
func TestValue(t *testing.T) {
	if collections.Value(42) != 42 {
		t.Error("Value failed")
	}
}

func TestValueFn(t *testing.T) {
	fn := func() int { return 42 }
	if collections.ValueFn[int](fn) != 42 {
		t.Error("ValueFn with function failed")
	}
	if collections.ValueFn[int](42) != 42 {
		t.Error("ValueFn with value failed")
	}
	if collections.ValueFn[int]("invalid") != 0 {
		t.Error("ValueFn with invalid type failed")
	}
}

func TestDataGetNil(t *testing.T) {
	if collections.DataGet(nil, "key") != nil {
		t.Error("DataGet nil failed")
	}
	if collections.DataGet(nil, "key", "default") != "default" {
		t.Error("DataGet nil with default failed")
	}
}

func TestDataGetEmptyKey(t *testing.T) {
	data := map[string]any{"a": 1}
	result := collections.DataGet(data, "")
	if result == nil {
		t.Error("DataGet empty key failed")
	}
}

func TestDataGetNonMap(t *testing.T) {
	if collections.DataGet("string", "key", "default") != "default" {
		t.Error("DataGet non-map failed")
	}
}

func TestDataSetNonMap(t *testing.T) {
	result := collections.DataSet("string", "key", "value")
	if result != "string" {
		t.Error("DataSet non-map failed")
	}
}

func TestDataForgetNonMap(t *testing.T) {
	result := collections.DataForget("string", "key")
	if result != "string" {
		t.Error("DataForget non-map failed")
	}
}

func TestCollectSlice(t *testing.T) {
	c := collections.CollectSlice([]int{1, 2, 3})
	if c.Count() != 3 {
		t.Error("CollectSlice failed")
	}
}

func TestCollectMap(t *testing.T) {
	m := collections.CollectMap(map[string]int{"a": 1})
	if m.Get("a") != 1 {
		t.Error("CollectMap failed")
	}
}

func TestHeadEmpty(t *testing.T) {
	if collections.Head([]int{}) != 0 {
		t.Error("Head empty failed")
	}
}

func TestTailEmpty(t *testing.T) {
	if len(collections.Tail([]int{})) != 0 {
		t.Error("Tail empty failed")
	}
}

func TestTailSingle(t *testing.T) {
	if len(collections.Tail([]int{1})) != 0 {
		t.Error("Tail single failed")
	}
}

func TestInitEmpty(t *testing.T) {
	if len(collections.Init([]int{})) != 0 {
		t.Error("Init empty failed")
	}
}

func TestInitSingle(t *testing.T) {
	if len(collections.Init([]int{1})) != 0 {
		t.Error("Init single failed")
	}
}

func TestLastItemEmpty(t *testing.T) {
	if collections.LastItem([]int{}) != 0 {
		t.Error("LastItem empty failed")
	}
}

func TestTransformBlank(t *testing.T) {
	result := collections.Transform("", func(s string) int { return len(s) })
	if result != 0 {
		t.Error("Transform blank failed")
	}
}

func TestWith(t *testing.T) {
	if collections.With(42) != 42 {
		t.Error("With failed")
	}
}

func TestOptionalGetPanic(t *testing.T) {
	defer func() { recover() }()
	none := collections.None[int]()
	none.Get()
	t.Error("Expected panic")
}

func TestOptionalMap(t *testing.T) {
	some := collections.Some(2)
	mapped := some.Map(func(n int) int { return n * 2 })
	if mapped.Get() != 4 {
		t.Error("Optional Map failed")
	}
	none := collections.None[int]()
	mapped2 := none.Map(func(n int) int { return n * 2 })
	if mapped2.HasValue() {
		t.Error("Optional Map on None failed")
	}
}

func TestOptionalFilter(t *testing.T) {
	some := collections.Some(2)
	filtered := some.Filter(func(n int) bool { return n > 1 })
	if !filtered.HasValue() {
		t.Error("Optional Filter pass failed")
	}
	filtered2 := some.Filter(func(n int) bool { return n > 10 })
	if filtered2.HasValue() {
		t.Error("Optional Filter fail failed")
	}
}

func TestRetry(t *testing.T) {
	count := 0
	result, err := collections.Retry(3, func() (int, error) {
		count++
		if count < 3 {
			return 0, errors.New("fail")
		}
		return 42, nil
	})
	if err != nil || result != 42 {
		t.Error("Retry failed")
	}
}

func TestRetryAllFail(t *testing.T) {
	_, err := collections.Retry(2, func() (int, error) {
		return 0, errors.New("fail")
	})
	if err == nil {
		t.Error("Retry all fail should return error")
	}
}

func TestOnce(t *testing.T) {
	count := 0
	fn := collections.Once(func() int {
		count++
		return count
	})
	first := fn()
	second := fn()
	if first != 1 || second != 1 || count != 1 {
		t.Error("Once failed")
	}
}

func TestRescue(t *testing.T) {
	result := collections.Rescue(func() int { return 42 }, 0)
	if result != 42 {
		t.Error("Rescue normal failed")
	}
}

func TestThrowIf(t *testing.T) {
	defer func() { recover() }()
	collections.ThrowIf(true, "error")
	t.Error("Expected panic")
}

func TestThrowIfFalse(t *testing.T) {
	collections.ThrowIf(false, "error")
}

func TestThrowUnless(t *testing.T) {
	defer func() { recover() }()
	collections.ThrowUnless(false, "error")
	t.Error("Expected panic")
}

func TestThrowUnlessTrue(t *testing.T) {
	collections.ThrowUnless(true, "error")
}
