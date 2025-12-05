package collections_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/qiuapeng921/collections"
)

// ============================================
// Collection 创建测试
// ============================================

func TestNew(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	if c.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", c.Count())
	}
}

func TestNewWithNil(t *testing.T) {
	c := collections.New[int](nil)
	if c.Count() != 0 {
		t.Errorf("Expected 0 items, got %d", c.Count())
	}
}

func TestMake(t *testing.T) {
	c := collections.Make(1, 2, 3)
	if c.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", c.Count())
	}
}

func TestRange(t *testing.T) {
	c := collections.Range(1, 5)
	expected := []int{1, 2, 3, 4, 5}
	if c.Count() != 5 {
		t.Errorf("Expected 5 items, got %d", c.Count())
	}
	for i, v := range c.All() {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestRangeReverse(t *testing.T) {
	c := collections.Range(5, 1)
	expected := []int{5, 4, 3, 2, 1}
	if c.Count() != 5 {
		t.Errorf("Expected 5 items, got %d", c.Count())
	}
	for i, v := range c.All() {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestTimes(t *testing.T) {
	c := collections.Times(5, func(i int) int {
		return i * 2
	})
	expected := []int{2, 4, 6, 8, 10}
	for i, v := range c.All() {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestEmpty(t *testing.T) {
	c := collections.Empty[int]()
	if !c.IsEmpty() {
		t.Error("Expected collection to be empty")
	}
	if c.Count() != 0 {
		t.Errorf("Expected 0 items, got %d", c.Count())
	}
}

// ============================================
// 访问测试
// ============================================

func TestIsEmpty(t *testing.T) {
	empty := collections.Empty[int]()
	if !empty.IsEmpty() {
		t.Error("Expected collection to be empty")
	}
	if empty.IsNotEmpty() {
		t.Error("Expected IsNotEmpty to return false")
	}

	notEmpty := collections.New([]int{1})
	if notEmpty.IsEmpty() {
		t.Error("Expected collection to not be empty")
	}
	if !notEmpty.IsNotEmpty() {
		t.Error("Expected IsNotEmpty to return true")
	}
}

func TestFirstLast(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	if c.First() != 1 {
		t.Errorf("Expected First() to be 1, got %d", c.First())
	}
	if c.Last() != 3 {
		t.Errorf("Expected Last() to be 3, got %d", c.Last())
	}
}

func TestFirstOrLastOr(t *testing.T) {
	empty := collections.Empty[int]()
	if empty.FirstOr(100) != 100 {
		t.Error("Expected FirstOr to return default value")
	}
	if empty.LastOr(200) != 200 {
		t.Error("Expected LastOr to return default value")
	}
}

func TestFirstWhere(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	value, found := c.FirstWhere(func(n int) bool { return n > 3 })
	if !found {
		t.Error("Expected to find item")
	}
	if value != 4 {
		t.Errorf("Expected 4, got %d", value)
	}

	_, notFound := c.FirstWhere(func(n int) bool { return n > 10 })
	if notFound {
		t.Error("Expected not to find item")
	}
}

func TestLastWhere(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	value, found := c.LastWhere(func(n int) bool { return n < 4 })
	if !found {
		t.Error("Expected to find item")
	}
	if value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}
}

func TestGet(t *testing.T) {
	c := collections.New([]int{10, 20, 30})
	if c.Get(1) != 20 {
		t.Errorf("Expected Get(1) to be 20, got %d", c.Get(1))
	}
	if c.Get(10) != 0 {
		t.Error("Expected Get with invalid index to return zero value")
	}
}

func TestGetOr(t *testing.T) {
	c := collections.New([]int{10, 20, 30})
	if c.GetOr(1, 100) != 20 {
		t.Error("Expected GetOr to return existing value")
	}
	if c.GetOr(10, 100) != 100 {
		t.Error("Expected GetOr to return default value for invalid index")
	}
}

func TestContainsOneItem(t *testing.T) {
	single := collections.Make(1)
	if !single.ContainsOneItem() {
		t.Error("Expected ContainsOneItem to be true")
	}
	multi := collections.Make(1, 2)
	if multi.ContainsOneItem() {
		t.Error("Expected ContainsOneItem to be false")
	}
}

// ============================================
// 过滤测试
// ============================================

func TestFilter(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5, 6})
	evens := c.Filter(func(n int) bool { return n%2 == 0 })
	if evens.Count() != 3 {
		t.Errorf("Expected 3 even numbers, got %d", evens.Count())
	}
}

func TestReject(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5, 6})
	notEvens := c.Reject(func(n int) bool { return n%2 == 0 })
	if notEvens.Count() != 3 {
		t.Errorf("Expected 3 odd numbers, got %d", notEvens.Count())
	}
}

func TestSlice(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	sliced := c.Slice(1, 3)
	expected := []int{2, 3, 4}
	if sliced.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", sliced.Count())
	}
	for i, v := range sliced.All() {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestSliceNegativeOffset(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	sliced := c.Slice(-3)
	if sliced.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", sliced.Count())
	}
	if sliced.First() != 3 {
		t.Errorf("Expected first item to be 3, got %d", sliced.First())
	}
}

func TestTakeSkip(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	taken := c.Take(2)
	if taken.Count() != 2 || taken.First() != 1 {
		t.Error("Take(2) failed")
	}

	skipped := c.Skip(3)
	if skipped.Count() != 2 || skipped.First() != 4 {
		t.Error("Skip(3) failed")
	}
}

func TestTakeNegative(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	taken := c.Take(-2)
	if taken.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", taken.Count())
	}
	if taken.First() != 4 {
		t.Errorf("Expected first item to be 4, got %d", taken.First())
	}
}

func TestTakeWhile(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5, 4, 3, 2, 1})
	result := c.TakeWhile(func(n int) bool { return n < 4 })
	if result.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", result.Count())
	}
}

func TestSkipWhile(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	result := c.SkipWhile(func(n int) bool { return n < 3 })
	if result.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", result.Count())
	}
	if result.First() != 3 {
		t.Errorf("Expected first item to be 3, got %d", result.First())
	}
}

func TestTakeUntil(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	result := c.TakeUntil(func(n int) bool { return n > 3 })
	if result.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", result.Count())
	}
}

func TestSkipUntil(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	result := c.SkipUntil(func(n int) bool { return n > 3 })
	if result.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", result.Count())
	}
	if result.First() != 4 {
		t.Errorf("Expected first item to be 4, got %d", result.First())
	}
}

// ============================================
// 转换测试
// ============================================

func TestMap(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	doubled := collections.Map(c, func(n int, i int) int { return n * 2 })
	expected := []int{2, 4, 6}
	for i, v := range doubled.All() {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestMapToString(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	strings := collections.Map(c, func(n int, i int) string {
		return fmt.Sprintf("num_%d", n)
	})
	if strings.First() != "num_1" {
		t.Errorf("Expected 'num_1', got '%s'", strings.First())
	}
}

func TestReduce(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	sum := collections.Reduce(c, func(acc int, n int, i int) int { return acc + n }, 0)
	if sum != 15 {
		t.Errorf("Expected sum to be 15, got %d", sum)
	}

	product := collections.Reduce(c, func(acc int, n int, i int) int { return acc * n }, 1)
	if product != 120 {
		t.Errorf("Expected product to be 120, got %d", product)
	}
}

func TestFlatMap(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	expanded := collections.FlatMap(c, func(n int, i int) []int {
		return []int{n, n * 10}
	})
	if expanded.Count() != 6 {
		t.Errorf("Expected 6 items, got %d", expanded.Count())
	}
}

func TestEach(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	sum := 0
	c.Each(func(n int, i int) {
		sum += n
	})
	if sum != 6 {
		t.Errorf("Expected sum to be 6, got %d", sum)
	}
}

// ============================================
// 条件测试
// ============================================

func TestContains(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	if !c.Contains(func(n int) bool { return n == 2 }) {
		t.Error("Expected collection to contain 2")
	}
	if c.Contains(func(n int) bool { return n == 10 }) {
		t.Error("Expected collection to not contain 10")
	}
}

func TestEvery(t *testing.T) {
	c := collections.New([]int{2, 4, 6})
	if !c.Every(func(n int) bool { return n%2 == 0 }) {
		t.Error("Expected all items to be even")
	}

	c2 := collections.New([]int{2, 3, 4})
	if c2.Every(func(n int) bool { return n%2 == 0 }) {
		t.Error("Expected not all items to be even")
	}
}

func TestSome(t *testing.T) {
	c := collections.New([]int{1, 3, 5, 6})
	if !c.Some(func(n int) bool { return n%2 == 0 }) {
		t.Error("Expected some items to be even")
	}
}

// ============================================
// 修改测试
// ============================================

func TestPushPop(t *testing.T) {
	c := collections.New([]int{1, 2})
	c.Push(3)
	if c.Last() != 3 {
		t.Errorf("Expected 3 to be pushed, got %d", c.Last())
	}
	popped := c.Pop()
	if popped != 3 {
		t.Errorf("Expected popped value to be 3, got %d", popped)
	}
}

func TestPrependShift(t *testing.T) {
	c := collections.New([]int{2, 3})
	c.Prepend(1)
	if c.First() != 1 {
		t.Errorf("Expected 1 to be prepended, got %d", c.First())
	}
	shifted := c.Shift()
	if shifted != 1 {
		t.Errorf("Expected shifted value to be 1, got %d", shifted)
	}
}

func TestPut(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	c.Put(1, 100)
	if c.Get(1) != 100 {
		t.Errorf("Expected Put to set value, got %d", c.Get(1))
	}
}

func TestForget(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	c.Forget(2)
	if c.Count() != 4 {
		t.Errorf("Expected 4 items after forget, got %d", c.Count())
	}
	if c.Get(2) != 4 {
		t.Errorf("Expected index 2 to be 4, got %d", c.Get(2))
	}
}

func TestPull(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	pulled := c.Pull(2)
	if pulled != 3 {
		t.Errorf("Expected pulled value to be 3, got %d", pulled)
	}
	if c.Count() != 4 {
		t.Errorf("Expected 4 items after pull, got %d", c.Count())
	}
}

func TestTransform(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	c.Transform(func(n int, i int) int { return n * 10 })
	if c.Get(0) != 10 || c.Get(1) != 20 || c.Get(2) != 30 {
		t.Error("Transform failed")
	}
}

// ============================================
// 分组测试
// ============================================

func TestChunk(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	chunks := c.ChunkInto(2)
	if len(chunks) != 3 {
		t.Errorf("Expected 3 chunks, got %d", len(chunks))
	}
	if chunks[0].Count() != 2 {
		t.Errorf("Expected first chunk to have 2 items, got %d", chunks[0].Count())
	}
	if chunks[2].Count() != 1 {
		t.Errorf("Expected last chunk to have 1 item, got %d", chunks[2].Count())
	}
}

func TestSplit(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	groups := c.Split(3)
	if len(groups) != 3 {
		t.Errorf("Expected 3 groups, got %d", len(groups))
	}
}

func TestPartition(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5, 6})
	evens, odds := c.Partition(func(n int) bool { return n%2 == 0 })
	if evens.Count() != 3 || odds.Count() != 3 {
		t.Errorf("Partition failed: evens=%d, odds=%d", evens.Count(), odds.Count())
	}
}

func TestSliding(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	windows := c.Sliding(3)
	if len(windows) != 3 {
		t.Errorf("Expected 3 windows, got %d", len(windows))
	}
}

func TestSlidingWithStep(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5, 6})
	windows := c.Sliding(2, 2)
	if len(windows) != 3 {
		t.Errorf("Expected 3 windows, got %d", len(windows))
	}
}

// ============================================
// 排序测试
// ============================================

func TestReverse(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	reversed := c.Reverse()
	expected := []int{3, 2, 1}
	for i, v := range reversed.All() {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestSort(t *testing.T) {
	c := collections.New([]int{3, 1, 4, 1, 5, 9, 2, 6})
	sorted := collections.Sort(c)
	prev := sorted.First()
	sorted.Each(func(n int, i int) {
		if n < prev {
			t.Errorf("Collection not sorted: %d came after %d", n, prev)
		}
		prev = n
	})
}

func TestSortDesc(t *testing.T) {
	c := collections.New([]int{3, 1, 4, 1, 5})
	sorted := collections.SortDesc(c)
	expected := []int{5, 4, 3, 1, 1}
	for i, v := range sorted.All() {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

type TestUser struct {
	ID   int
	Name string
	Age  int
}

func TestSortBy(t *testing.T) {
	users := collections.New([]TestUser{
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 3, Name: "Charlie", Age: 20},
	})

	sorted := collections.SortBy(users, func(u TestUser) int { return u.Age })
	if sorted.First().Name != "Charlie" {
		t.Errorf("Expected Charlie first, got %s", sorted.First().Name)
	}
}

func TestSortByDesc(t *testing.T) {
	users := collections.New([]TestUser{
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 1, Name: "Alice", Age: 30},
	})

	sorted := collections.SortByDesc(users, func(u TestUser) int { return u.Age })
	if sorted.First().Name != "Alice" {
		t.Errorf("Expected Alice first, got %s", sorted.First().Name)
	}
}

// ============================================
// 聚合测试
// ============================================

func TestSumAvg(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	if collections.Sum(c) != 15 {
		t.Errorf("Expected sum to be 15, got %d", collections.Sum(c))
	}
	if collections.Avg(c) != 3.0 {
		t.Errorf("Expected avg to be 3.0, got %f", collections.Avg(c))
	}
}

func TestMinMax(t *testing.T) {
	c := collections.New([]int{3, 1, 4, 1, 5, 9, 2, 6})
	if collections.Min(c) != 1 {
		t.Errorf("Expected min to be 1, got %d", collections.Min(c))
	}
	if collections.Max(c) != 9 {
		t.Errorf("Expected max to be 9, got %d", collections.Max(c))
	}
}

func TestMedian(t *testing.T) {
	odd := collections.New([]int{1, 2, 3, 4, 5})
	if collections.Median(odd) != 3.0 {
		t.Errorf("Expected median to be 3.0, got %f", collections.Median(odd))
	}

	even := collections.New([]int{1, 2, 3, 4})
	if collections.Median(even) != 2.5 {
		t.Errorf("Expected median to be 2.5, got %f", collections.Median(even))
	}
}

func TestMode(t *testing.T) {
	c := collections.New([]int{1, 2, 2, 3, 3, 3, 4})
	mode := collections.Mode(c)
	if len(mode) != 1 || mode[0] != 3 {
		t.Errorf("Expected mode to be [3], got %v", mode)
	}
}

func TestSumBy(t *testing.T) {
	users := collections.New([]TestUser{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
	})
	sum := collections.SumBy(users, func(u TestUser) int { return u.Age })
	if sum != 55 {
		t.Errorf("Expected sum to be 55, got %d", sum)
	}
}

func TestMinMaxBy(t *testing.T) {
	users := collections.New([]TestUser{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	})

	youngest := collections.MinBy(users, func(u TestUser) int { return u.Age })
	if youngest.Name != "Bob" {
		t.Errorf("Expected Bob, got %s", youngest.Name)
	}

	oldest := collections.MaxBy(users, func(u TestUser) int { return u.Age })
	if oldest.Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", oldest.Name)
	}
}

// ============================================
// 集合操作测试
// ============================================

func TestUnique(t *testing.T) {
	c := collections.New([]int{1, 2, 2, 3, 3, 3})
	unique := collections.UniqueComparable(c)
	if unique.Count() != 3 {
		t.Errorf("Expected 3 unique items, got %d", unique.Count())
	}
}

func TestUniqueWithKeyFn(t *testing.T) {
	c := collections.New([]string{"apple", "apricot", "banana", "berry"})
	unique := c.Unique(func(s string) string {
		return string(s[0])
	})
	if unique.Count() != 2 {
		t.Errorf("Expected 2 unique items, got %d", unique.Count())
	}
}

func TestDuplicates(t *testing.T) {
	c := collections.New([]int{1, 2, 2, 3, 3, 3, 4})
	duplicates := collections.Duplicates(c)
	if duplicates.Count() != 2 {
		t.Errorf("Expected 2 duplicates, got %d", duplicates.Count())
	}
}

func TestDiffIntersect(t *testing.T) {
	c1 := collections.New([]int{1, 2, 3, 4, 5})
	c2 := collections.New([]int{3, 4, 5, 6, 7})

	diff := collections.Diff(c1, c2)
	if diff.Count() != 2 {
		t.Errorf("Expected 2 diff items, got %d", diff.Count())
	}

	intersect := collections.Intersect(c1, c2)
	if intersect.Count() != 3 {
		t.Errorf("Expected 3 intersect items, got %d", intersect.Count())
	}
}

func TestMerge(t *testing.T) {
	c1 := collections.New([]int{1, 2, 3})
	c2 := collections.New([]int{4, 5, 6})
	merged := c1.Merge(c2)
	if merged.Count() != 6 {
		t.Errorf("Expected 6 items, got %d", merged.Count())
	}
}

func TestConcat(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	result := c.Concat([]int{4, 5})
	if result.Count() != 5 {
		t.Errorf("Expected 5 items, got %d", result.Count())
	}
}

// ============================================
// 其他操作测试
// ============================================

func TestNth(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5, 6})
	nth := c.Nth(2)
	expected := []int{1, 3, 5}
	if nth.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", nth.Count())
	}
	for i, v := range nth.All() {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestNthWithOffset(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5, 6})
	nth := c.Nth(2, 1)
	expected := []int{2, 4, 6}
	for i, v := range nth.All() {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestForPage(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	page := c.ForPage(2, 3)
	if page.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", page.Count())
	}
	if page.First() != 4 {
		t.Errorf("Expected first item to be 4, got %d", page.First())
	}
}

func TestPad(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	padded := c.Pad(5, 0)
	if padded.Count() != 5 {
		t.Errorf("Expected 5 items, got %d", padded.Count())
	}
	if padded.Last() != 0 {
		t.Errorf("Expected last item to be 0, got %d", padded.Last())
	}
}

func TestPadNegative(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	padded := c.Pad(-5, 0)
	if padded.Count() != 5 {
		t.Errorf("Expected 5 items, got %d", padded.Count())
	}
	if padded.First() != 0 {
		t.Errorf("Expected first item to be 0, got %d", padded.First())
	}
}

func TestClone(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	cloned := c.Clone()
	cloned.Push(4)
	if c.Count() != 3 {
		t.Error("Original should not be affected by clone modification")
	}
	if cloned.Count() != 4 {
		t.Error("Clone should have 4 items")
	}
}

func TestSearch(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	index := c.Search(func(n int) bool { return n == 3 })
	if index != 2 {
		t.Errorf("Expected index 2, got %d", index)
	}

	notFound := c.Search(func(n int) bool { return n == 10 })
	if notFound != -1 {
		t.Errorf("Expected -1 for not found, got %d", notFound)
	}
}

func TestSole(t *testing.T) {
	single := collections.Make(42)
	value, err := single.Sole()
	if err != nil || value != 42 {
		t.Error("Sole failed for single item collection")
	}

	multi := collections.Make(1, 2)
	_, err = multi.Sole()
	if err == nil {
		t.Error("Expected error for multiple items")
	}

	empty := collections.Empty[int]()
	_, err = empty.Sole()
	if err == nil {
		t.Error("Expected error for empty collection")
	}
}

func TestToJSON(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	jsonStr := c.ToJSONString()
	if jsonStr != "[1,2,3]" {
		t.Errorf("Expected '[1,2,3]', got '%s'", jsonStr)
	}
}

func TestWhen(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	result := c.When(true, func(col *collections.Collection[int]) *collections.Collection[int] {
		return col.Filter(func(n int) bool { return n > 1 })
	})
	if result.Count() != 2 {
		t.Errorf("Expected 2 items when condition is true, got %d", result.Count())
	}

	result2 := c.When(false, func(col *collections.Collection[int]) *collections.Collection[int] {
		return col.Filter(func(n int) bool { return n > 1 })
	})
	if result2.Count() != 3 {
		t.Errorf("Expected 3 items when condition is false, got %d", result2.Count())
	}
}

func TestUnless(t *testing.T) {
	c := collections.New([]int{1, 2, 3})
	result := c.Unless(false, func(col *collections.Collection[int]) *collections.Collection[int] {
		return col.Take(1)
	})
	if result.Count() != 1 {
		t.Errorf("Expected 1 item when condition is false, got %d", result.Count())
	}
}

// ============================================
// MapCollection 测试
// ============================================

func TestMapCollection(t *testing.T) {
	m := collections.NewMap(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})

	if m.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", m.Count())
	}

	if m.Get("b") != 2 {
		t.Errorf("Expected Get('b') to be 2, got %d", m.Get("b"))
	}

	if !m.Has("a") {
		t.Error("Expected Has('a') to be true")
	}

	if m.Has("z") {
		t.Error("Expected Has('z') to be false")
	}
}

func TestMapKeysValues(t *testing.T) {
	m := collections.NewMap(map[string]int{
		"a": 1,
		"b": 2,
	})

	keys := m.Keys()
	if keys.Count() != 2 {
		t.Errorf("Expected 2 keys, got %d", keys.Count())
	}

	values := m.Values()
	if values.Count() != 2 {
		t.Errorf("Expected 2 values, got %d", values.Count())
	}
}

func TestMapPutForget(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	m.Put("b", 2)
	if m.Get("b") != 2 {
		t.Error("Put failed")
	}

	m.Forget("a")
	if m.Has("a") {
		t.Error("Forget failed")
	}
}

func TestMapFilter(t *testing.T) {
	m := collections.NewMap(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})

	filtered := m.Filter(func(v int, k string) bool { return v > 1 })
	if filtered.Count() != 2 {
		t.Errorf("Expected 2 filtered items, got %d", filtered.Count())
	}
}

func TestMapOnlyExcept(t *testing.T) {
	m := collections.NewMap(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})

	only := m.Only("a", "b")
	if only.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", only.Count())
	}

	except := m.Except("a")
	if except.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", except.Count())
	}
}

func TestMapMerge(t *testing.T) {
	m1 := collections.NewMap(map[string]int{"a": 1})
	m2 := collections.NewMap(map[string]int{"b": 2})
	merged := m1.Merge(m2)
	if merged.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", merged.Count())
	}
}

// ============================================
// Pluck 和 GroupBy 测试
// ============================================

func TestPluck(t *testing.T) {
	users := collections.New([]TestUser{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
	})

	names := collections.Pluck(users, func(u TestUser) string { return u.Name })
	if names.Count() != 2 || names.First() != "Alice" {
		t.Error("Pluck failed")
	}
}

func TestKeyBy(t *testing.T) {
	users := collections.New([]TestUser{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
	})

	keyed := collections.KeyBy(users, func(u TestUser) int { return u.ID })
	if keyed.Get(1).Name != "Alice" {
		t.Error("KeyBy failed")
	}
}

func TestGroupBy(t *testing.T) {
	users := collections.New([]TestUser{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 30},
	})

	grouped := collections.GroupBy(users, func(u TestUser) int { return u.Age })
	if grouped.Get(30).Count() != 2 {
		t.Error("GroupBy failed")
	}
}

func TestCountBy(t *testing.T) {
	users := collections.New([]TestUser{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 30},
	})

	counted := collections.CountBy(users, func(u TestUser) int { return u.Age })
	if counted.Get(30) != 2 {
		t.Errorf("Expected 2, got %d", counted.Get(30))
	}
}

// ============================================
// Zip 和 CrossJoin 测试
// ============================================

func TestZip(t *testing.T) {
	a := collections.New([]int{1, 2, 3})
	b := collections.New([]int{4, 5, 6})
	zipped := collections.Zip(a, b)
	if zipped.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", zipped.Count())
	}
}

func TestCrossJoin(t *testing.T) {
	a := collections.New([]string{"a", "b"})
	b := collections.New([]string{"1", "2"})
	product := collections.CrossJoin(a, b)
	if product.Count() != 4 {
		t.Errorf("Expected 4 items, got %d", product.Count())
	}
}

func TestCombine(t *testing.T) {
	keys := collections.New([]string{"name", "age"})
	values := collections.New([]string{"Alice", "30"})
	combined := collections.Combine(keys, values)
	if combined.Get("name") != "Alice" {
		t.Error("Combine failed")
	}
}

// ============================================
// 字符串操作测试
// ============================================

func TestImplodeStrings(t *testing.T) {
	names := collections.New([]string{"Alice", "Bob", "Charlie"})
	result := collections.ImplodeStrings(names, ", ")
	if result != "Alice, Bob, Charlie" {
		t.Errorf("Expected 'Alice, Bob, Charlie', got '%s'", result)
	}
}

func TestJoinStrings(t *testing.T) {
	names := collections.New([]string{"Alice", "Bob", "Charlie"})
	result := collections.JoinStrings(names, ", ", " and ")
	if result != "Alice, Bob and Charlie" {
		t.Errorf("Expected 'Alice, Bob and Charlie', got '%s'", result)
	}
}

func TestImplode(t *testing.T) {
	numbers := collections.New([]int{1, 2, 3})
	result := collections.Implode(numbers, func(n int) string {
		return fmt.Sprintf("%d", n)
	}, "-")
	if result != "1-2-3" {
		t.Errorf("Expected '1-2-3', got '%s'", result)
	}
}

// ============================================
// IndexOf 测试
// ============================================

func TestIndexOf(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	if collections.IndexOf(c, 3) != 2 {
		t.Error("IndexOf failed")
	}
	if collections.IndexOf(c, 10) != -1 {
		t.Error("IndexOf should return -1 for not found")
	}
}

func TestLastIndexOf(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 2, 1})
	if collections.LastIndexOf(c, 2) != 3 {
		t.Error("LastIndexOf failed")
	}
}

// ============================================
// ContainsComparable 测试
// ============================================

func TestContainsComparable(t *testing.T) {
	c := collections.New([]string{"apple", "banana", "cherry"})
	if !collections.ContainsComparable(c, "banana") {
		t.Error("Expected to contain 'banana'")
	}
	if collections.ContainsComparable(c, "grape") {
		t.Error("Expected not to contain 'grape'")
	}
}

// ============================================
// Tap 测试
// ============================================

func TestTap(t *testing.T) {
	var tapResult int
	c := collections.New([]int{1, 2, 3})
	result := c.Tap(func(col *collections.Collection[int]) {
		tapResult = col.Count()
	})
	if tapResult != 3 {
		t.Error("Tap callback not called correctly")
	}
	if result.Count() != 3 {
		t.Error("Tap should return the same collection")
	}
}

// ============================================
// Replace 和 Splice 测试
// ============================================

func TestReplace(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	replaced := c.Replace([]int{10, 20})
	if replaced.Get(0) != 10 || replaced.Get(1) != 20 || replaced.Get(2) != 3 {
		t.Error("Replace failed")
	}
}

func TestSplice(t *testing.T) {
	c := collections.New([]int{1, 2, 3, 4, 5})
	_, removed := c.Splice(2, 2)
	if removed.Count() != 2 {
		t.Errorf("Expected 2 removed items, got %d", removed.Count())
	}
}

// ============================================
// WhenEmpty 和 WhenNotEmpty 测试
// ============================================

func TestWhenEmpty(t *testing.T) {
	empty := collections.Empty[int]()
	result := empty.WhenEmpty(func(c *collections.Collection[int]) *collections.Collection[int] {
		return collections.Make(1, 2, 3)
	})
	if result.Count() != 3 {
		t.Error("WhenEmpty should execute callback for empty collection")
	}

	notEmpty := collections.Make(1)
	result2 := notEmpty.WhenEmpty(func(c *collections.Collection[int]) *collections.Collection[int] {
		return collections.Make(4, 5, 6)
	})
	if result2.Count() != 1 {
		t.Error("WhenEmpty should not execute callback for non-empty collection")
	}
}

func TestWhenNotEmpty(t *testing.T) {
	notEmpty := collections.Make(1, 2, 3)
	result := notEmpty.WhenNotEmpty(func(c *collections.Collection[int]) *collections.Collection[int] {
		return c.Take(1)
	})
	if result.Count() != 1 {
		t.Error("WhenNotEmpty should execute callback for non-empty collection")
	}
}

// ============================================
// Error 类型测试
// ============================================

func TestFirstOrFail(t *testing.T) {
	c := collections.Make(1, 2, 3)
	value, err := c.FirstOrFail()
	if err != nil || value != 1 {
		t.Error("FirstOrFail should return first item")
	}

	empty := collections.Empty[int]()
	_, err = empty.FirstOrFail()
	if err == nil {
		t.Error("FirstOrFail should return error for empty collection")
	}
}

func TestLastOrFail(t *testing.T) {
	c := collections.Make(1, 2, 3)
	value, err := c.LastOrFail()
	if err != nil || value != 3 {
		t.Error("LastOrFail should return last item")
	}
}

func TestGetOrFail(t *testing.T) {
	c := collections.Make(1, 2, 3)
	value, err := c.GetOrFail(1)
	if err != nil || value != 2 {
		t.Error("GetOrFail should return item at index")
	}

	_, err = c.GetOrFail(10)
	if err == nil {
		t.Error("GetOrFail should return error for invalid index")
	}
}

// ============================================
// MapCollection 额外测试
// ============================================

func TestMapGetOr(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	if m.GetOr("a", 100) != 1 {
		t.Error("GetOr should return existing value")
	}
	if m.GetOr("z", 100) != 100 {
		t.Error("GetOr should return default for missing key")
	}
}

func TestMapHasAny(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2})
	if !m.HasAny("a", "z") {
		t.Error("HasAny should return true if any key exists")
	}
	if m.HasAny("x", "y", "z") {
		t.Error("HasAny should return false if no key exists")
	}
}

func TestMapFirstLast(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2, "c": 3})
	// Note: map order is not guaranteed, just check they return valid values
	first := m.First()
	if first < 1 || first > 3 {
		t.Error("First should return a valid value")
	}
}

func TestMapContains(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1, "b": 2, "c": 3})
	if !m.Contains(func(v int, k string) bool { return v == 2 }) {
		t.Error("Contains should find value")
	}
}

func TestMapEvery(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 2, "b": 4, "c": 6})
	if !m.Every(func(v int, k string) bool { return v%2 == 0 }) {
		t.Error("Every should return true for all even values")
	}
}

func TestMapClone(t *testing.T) {
	m := collections.NewMap(map[string]int{"a": 1})
	cloned := m.Clone()
	cloned.Put("b", 2)
	if m.Has("b") {
		t.Error("Original should not be affected by clone")
	}
}

// ============================================
// Arr 辅助函数测试
// ============================================

func TestArrGet(t *testing.T) {
	data := map[string]any{
		"user": map[string]any{
			"name": "Alice",
			"profile": map[string]any{
				"age": 30,
			},
		},
	}

	name := collections.Arr.Get(data, "user.name")
	if name != "Alice" {
		t.Errorf("Expected 'Alice', got '%v'", name)
	}

	age := collections.Arr.Get(data, "user.profile.age")
	if age != 30 {
		t.Errorf("Expected 30, got '%v'", age)
	}

	missing := collections.Arr.Get(data, "user.missing", "default")
	if missing != "default" {
		t.Errorf("Expected 'default', got '%v'", missing)
	}
}

func TestArrSet(t *testing.T) {
	data := map[string]any{}
	collections.Arr.Set(data, "user.name", "Bob")

	name := collections.Arr.Get(data, "user.name")
	if name != "Bob" {
		t.Errorf("Expected 'Bob', got '%v'", name)
	}
}

func TestArrHas(t *testing.T) {
	data := map[string]any{
		"user": map[string]any{
			"name": "Alice",
		},
	}

	if !collections.Arr.Has(data, "user.name") {
		t.Error("Expected Has to return true")
	}
	if collections.Arr.Has(data, "user.missing") {
		t.Error("Expected Has to return false")
	}
}

func TestArrForget(t *testing.T) {
	data := map[string]any{
		"a": 1,
		"b": 2,
	}
	collections.Arr.Forget(data, "a")
	if collections.Arr.Exists(data, "a") {
		t.Error("Expected key to be forgotten")
	}
}

func TestArrDot(t *testing.T) {
	data := map[string]any{
		"user": map[string]any{
			"name": "Alice",
			"age":  30,
		},
	}
	dotted := collections.Arr.Dot(data)
	if dotted["user.name"] != "Alice" {
		t.Error("Dot notation failed")
	}
}

func TestArrUndot(t *testing.T) {
	data := map[string]any{
		"user.name": "Alice",
		"user.age":  30,
	}
	undotted := collections.Arr.Undot(data)
	user := undotted["user"].(map[string]any)
	if user["name"] != "Alice" {
		t.Error("Undot failed")
	}
}

func TestArrOnlyExcept(t *testing.T) {
	data := map[string]any{"a": 1, "b": 2, "c": 3}

	only := collections.Arr.Only(data, "a", "b")
	if len(only) != 2 {
		t.Error("Only failed")
	}

	except := collections.Arr.Except(data, "a")
	if len(except) != 2 || collections.Arr.Exists(except, "a") {
		t.Error("Except failed")
	}
}

func TestArrWrap(t *testing.T) {
	// Wrap single value
	wrapped := collections.Arr.Wrap("hello")
	if len(wrapped) != 1 {
		t.Error("Wrap single value failed")
	}

	// Wrap slice
	slice := []any{1, 2, 3}
	wrapped2 := collections.Arr.Wrap(slice)
	if len(wrapped2) != 3 {
		t.Error("Wrap slice failed")
	}

	// Wrap nil
	wrapped3 := collections.Arr.Wrap(nil)
	if len(wrapped3) != 0 {
		t.Error("Wrap nil failed")
	}
}

// ============================================
// Helper 函数测试
// ============================================

func TestCollect(t *testing.T) {
	c := collections.Collect(1, 2, 3)
	if c.Count() != 3 {
		t.Error("Collect failed")
	}
}

func TestHeadTail(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	if collections.Head(items) != 1 {
		t.Error("Head failed")
	}

	tail := collections.Tail(items)
	if len(tail) != 4 || tail[0] != 2 {
		t.Error("Tail failed")
	}
}

func TestInit(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	init := collections.Init(items)
	if len(init) != 4 || init[3] != 4 {
		t.Error("Init failed")
	}
}

func TestLastItem(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	if collections.LastItem(items) != 5 {
		t.Error("LastItem failed")
	}
}

func TestBlankFilled(t *testing.T) {
	if !collections.Blank("") {
		t.Error("Blank should return true for empty string")
	}
	if collections.Blank("hello") {
		t.Error("Blank should return false for non-empty string")
	}
	if !collections.Filled("hello") {
		t.Error("Filled should return true for non-empty string")
	}
}

func TestTransformHelper(t *testing.T) {
	result := collections.Transform("hello", func(s string) string {
		return strings.ToUpper(s)
	})
	if result != "HELLO" {
		t.Error("Transform helper failed")
	}

	result2 := collections.Transform("", func(s string) string {
		return "transformed"
	}, "default")
	if result2 != "default" {
		t.Error("Transform should return default for blank value")
	}
}

func TestOptional(t *testing.T) {
	some := collections.Some(42)
	if !some.HasValue() {
		t.Error("Some should have value")
	}
	if some.Get() != 42 {
		t.Error("Some.Get failed")
	}

	none := collections.None[int]()
	if none.HasValue() {
		t.Error("None should not have value")
	}
	if none.GetOr(100) != 100 {
		t.Error("None.GetOr failed")
	}
}

func TestTapHelper(t *testing.T) {
	var called bool
	result := collections.Tap(42, func(v int) {
		called = true
	})
	if !called || result != 42 {
		t.Error("Tap helper failed")
	}
}

func TestIdentity(t *testing.T) {
	if collections.Identity(42) != 42 {
		t.Error("Identity failed")
	}
}
