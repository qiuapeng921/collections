package collections_test

import (
	"testing"

	"github.com/qiuapeng921/collections"
)

func TestArrGetEmptyKey(t *testing.T) {
	data := map[string]any{"a": 1}
	result := collections.Arr.Get(data, "")
	if result == nil {
		t.Error("Get empty key failed")
	}
}

func TestArrGetDirectKey(t *testing.T) {
	data := map[string]any{"a": 1}
	if collections.Arr.Get(data, "a") != 1 {
		t.Error("Get direct key failed")
	}
}

func TestArrGetMissingNoDot(t *testing.T) {
	data := map[string]any{"a": 1}
	if collections.Arr.Get(data, "b") != nil {
		t.Error("Get missing no dot failed")
	}
}

func TestArrGetNestedMissing(t *testing.T) {
	data := map[string]any{"a": map[string]any{"b": 1}}
	if collections.Arr.Get(data, "a.c", "default") != "default" {
		t.Error("Get nested missing failed")
	}
}

func TestArrGetNonMapNested(t *testing.T) {
	data := map[string]any{"a": "string"}
	if collections.Arr.Get(data, "a.b", "default") != "default" {
		t.Error("Get non-map nested failed")
	}
}

func TestArrSetEmptyKey(t *testing.T) {
	data := map[string]any{"a": 1}
	result := collections.Arr.Set(data, "", "value")
	if result["a"] != 1 {
		t.Error("Set empty key failed")
	}
}

func TestArrSetOverwriteNonMap(t *testing.T) {
	data := map[string]any{"a": "string"}
	collections.Arr.Set(data, "a.b", "value")
	if collections.Arr.Get(data, "a.b") != "value" {
		t.Error("Set overwrite non-map failed")
	}
}

func TestArrHasMultiple(t *testing.T) {
	data := map[string]any{"a": 1, "b": 2}
	if !collections.Arr.Has(data, "a", "b") {
		t.Error("Has multiple failed")
	}
	if collections.Arr.Has(data, "a", "c") {
		t.Error("Has multiple with missing failed")
	}
}

func TestArrHasAnyNone(t *testing.T) {
	data := map[string]any{"a": 1}
	if collections.Arr.HasAny(data, "b", "c") {
		t.Error("HasAny none failed")
	}
}

func TestArrForgetNoDot(t *testing.T) {
	data := map[string]any{"a": 1, "b": 2}
	collections.Arr.Forget(data, "a")
	if collections.Arr.Exists(data, "a") {
		t.Error("Forget no dot failed")
	}
}

func TestArrForgetNestedMissing(t *testing.T) {
	data := map[string]any{"a": "string"}
	collections.Arr.Forget(data, "a.b.c")
}

func TestArrDotWithPrepend(t *testing.T) {
	data := map[string]any{"a": 1}
	result := collections.Arr.Dot(data, "prefix")
	if result["prefix.a"] != 1 {
		t.Error("Dot with prepend failed")
	}
}

func TestArrDotEmptyNested(t *testing.T) {
	data := map[string]any{"a": map[string]any{}}
	result := collections.Arr.Dot(data)
	if result["a"] == nil {
		t.Error("Dot empty nested failed")
	}
}

func TestArrAddExisting(t *testing.T) {
	data := map[string]any{"a": 1}
	collections.Arr.Add(data, "a", 2)
	if data["a"] != 1 {
		t.Error("Add existing failed")
	}
}

func TestArrPullWithDefault(t *testing.T) {
	data := map[string]any{"a": 1}
	result := collections.Arr.Pull(data, "b", "default")
	if result != "default" {
		t.Error("Pull with default failed")
	}
}

func TestArrFirstEmpty(t *testing.T) {
	if collections.Arr.First([]any{}) != nil {
		t.Error("First empty failed")
	}
}

func TestArrFirstWithPredicate(t *testing.T) {
	items := []any{1, 2, 3}
	result := collections.Arr.First(items, func(v any) bool { return v.(int) > 1 })
	if result != 2 {
		t.Error("First with predicate failed")
	}
}

func TestArrFirstPredicateNoMatch(t *testing.T) {
	items := []any{1, 2, 3}
	result := collections.Arr.First(items, func(v any) bool { return v.(int) > 10 })
	if result != nil {
		t.Error("First predicate no match failed")
	}
}

func TestArrLastEmpty(t *testing.T) {
	if collections.Arr.Last([]any{}) != nil {
		t.Error("Last empty failed")
	}
}

func TestArrLastWithPredicate(t *testing.T) {
	items := []any{1, 2, 3}
	result := collections.Arr.Last(items, func(v any) bool { return v.(int) < 3 })
	if result != 2 {
		t.Error("Last with predicate failed")
	}
}

func TestArrLastPredicateNoMatch(t *testing.T) {
	items := []any{1, 2, 3}
	result := collections.Arr.Last(items, func(v any) bool { return v.(int) > 10 })
	if result != nil {
		t.Error("Last predicate no match failed")
	}
}

func TestArrWhere(t *testing.T) {
	items := []any{1, 2, 3}
	result := collections.Arr.Where(items, func(v any, i any) bool { return v.(int) > 1 })
	if len(result) != 2 {
		t.Error("Where failed")
	}
}

func TestArrWhereNotNull(t *testing.T) {
	items := []any{1, nil, 2, nil, 3}
	result := collections.Arr.WhereNotNull(items)
	if len(result) != 3 {
		t.Error("WhereNotNull failed")
	}
}

func TestArrShuffle(t *testing.T) {
	items := []any{1, 2, 3}
	result := collections.Arr.Shuffle(items)
	if len(result) != 3 {
		t.Error("Shuffle failed")
	}
}

func TestArrRandomSingle(t *testing.T) {
	items := []any{1, 2, 3}
	result := collections.Arr.Random(items)
	if result == nil {
		t.Error("Random single failed")
	}
}

func TestArrRandomMultiple(t *testing.T) {
	items := []any{1, 2, 3, 4, 5}
	result := collections.Arr.Random(items, 2)
	if slice, ok := result.([]any); !ok || len(slice) != 2 {
		t.Error("Random multiple failed")
	}
}

func TestArrDivide(t *testing.T) {
	data := map[string]any{"a": 1, "b": 2}
	keys, values := collections.Arr.Divide(data)
	if len(keys) != 2 || len(values) != 2 {
		t.Error("Divide failed")
	}
}

func TestArrIsAssoc(t *testing.T) {
	if collections.Arr.IsAssoc([]any{1, 2}) {
		t.Error("IsAssoc should return false")
	}
}

func TestArrIsList(t *testing.T) {
	if !collections.Arr.IsList([]any{1, 2}) {
		t.Error("IsList should return true")
	}
}

func TestArrQuery(t *testing.T) {
	data := map[string]string{"a": "1", "b": "2"}
	result := collections.Arr.Query(data)
	if result == "" {
		t.Error("Query failed")
	}
}

func TestArrAccessible(t *testing.T) {
	if !collections.Arr.Accessible(nil) {
		t.Error("Accessible should return true")
	}
}

func TestArrCollapse(t *testing.T) {
	items := [][]any{{1, 2}, {3, 4}}
	result := collections.Arr.Collapse(items)
	if len(result) != 4 {
		t.Error("Collapse failed")
	}
}

func TestArrPrepend(t *testing.T) {
	items := []any{2, 3}
	result := collections.Arr.Prepend(items, 1)
	if len(result) != 3 || result[0] != 1 {
		t.Error("Prepend failed")
	}
}

func TestArrCrossJoinEmpty(t *testing.T) {
	result := collections.Arr.CrossJoin()
	if len(result) != 0 {
		t.Error("CrossJoin empty failed")
	}
}

func TestArrCrossJoin(t *testing.T) {
	result := collections.Arr.CrossJoin([]any{1, 2}, []any{"a", "b"})
	if len(result) != 4 {
		t.Error("CrossJoin failed")
	}
}
