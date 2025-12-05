package collections_test

import (
	"testing"

	"github.com/qiuapeng921/collections"
)

func TestItemNotFoundExceptionError(t *testing.T) {
	e := &collections.ItemNotFoundException{}
	if e.Error() != "item not found" {
		t.Error("Default message failed")
	}
	e2 := &collections.ItemNotFoundException{Message: "custom"}
	if e2.Error() != "custom" {
		t.Error("Custom message failed")
	}
}

func TestMultipleItemsFoundExceptionError(t *testing.T) {
	e := &collections.MultipleItemsFoundException{}
	if e.Error() != "multiple items found" {
		t.Error("Default message failed")
	}
	e2 := &collections.MultipleItemsFoundException{Message: "custom"}
	if e2.Error() != "custom" {
		t.Error("Custom message failed")
	}
}

func TestInvalidArgumentExceptionError(t *testing.T) {
	e := &collections.InvalidArgumentException{}
	if e.Error() != "invalid argument" {
		t.Error("Default message failed")
	}
	e2 := &collections.InvalidArgumentException{Message: "custom"}
	if e2.Error() != "custom" {
		t.Error("Custom message failed")
	}
}

func TestFirstWhereOrFail(t *testing.T) {
	c := collections.Make(1, 2, 3)
	v, err := c.FirstWhereOrFail(func(n int) bool { return n > 1 })
	if err != nil || v != 2 {
		t.Error("FirstWhereOrFail found failed")
	}
	_, err = c.FirstWhereOrFail(func(n int) bool { return n > 10 })
	if err == nil {
		t.Error("FirstWhereOrFail not found failed")
	}
}

func TestRandomOrFail(t *testing.T) {
	c := collections.Make(1, 2, 3)
	_, err := c.RandomOrFail()
	if err != nil {
		t.Error("RandomOrFail with items failed")
	}
	empty := collections.Empty[int]()
	_, err = empty.RandomOrFail()
	if err == nil {
		t.Error("RandomOrFail empty failed")
	}
}

func TestPopOrFail(t *testing.T) {
	c := collections.Make(1, 2, 3)
	v, err := c.PopOrFail()
	if err != nil || v != 3 {
		t.Error("PopOrFail with items failed")
	}
	empty := collections.Empty[int]()
	_, err = empty.PopOrFail()
	if err == nil {
		t.Error("PopOrFail empty failed")
	}
}

func TestShiftOrFail(t *testing.T) {
	c := collections.Make(1, 2, 3)
	v, err := c.ShiftOrFail()
	if err != nil || v != 1 {
		t.Error("ShiftOrFail with items failed")
	}
	empty := collections.Empty[int]()
	_, err = empty.ShiftOrFail()
	if err == nil {
		t.Error("ShiftOrFail empty failed")
	}
}
