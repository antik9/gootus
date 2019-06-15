package list

import (
	"testing"
)

func TestPushBackFirst(t *testing.T) {
	list := NewList()
	list.PushBack("XY")

	if list.First().Value() != "XY" || list.Last().Value() != "XY" {
		t.Errorf("error with PushBack to empty List")
	}

	if list.First().Prev() != nil || list.First().Next() != nil {
		t.Errorf("first value in list unexpectedly has neighbours")
	}

	if list.Last().Prev() != nil || list.Last().Next() != nil {
		t.Errorf("last value in list unexpectedly has neighbours")
	}

	if list.Len() != 1 {
		t.Errorf("got list with length %d; wants list with length 1", list.Len())
	}
}

func TestPushFrontFirst(t *testing.T) {
	list := NewList()
	list.PushFront("XY")

	if list.First().Value() != "XY" || list.Last().Value() != "XY" {
		t.Errorf("error with PushBack to empty List")
	}

	if list.First().Prev() != nil || list.First().Next() != nil {
		t.Errorf("first value in list unexpectedly has neighbours")
	}

	if list.Last().Prev() != nil || list.Last().Next() != nil {
		t.Errorf("last value in list unexpectedly has neighbours")
	}

	if list.Len() != 1 {
		t.Errorf("got list with length %d; wants list with length 1", list.Len())
	}
}

func TestOrderWithPush(t *testing.T) {
	list := NewList()
	list.PushFront("XY")
	list.PushBack("ZW")

	if list.Last().Value() != "ZW" {
		t.Errorf(
			"got last element with value %s; wants element with value ZW",
			list.Last().Value())
	}

	if list.First().Value() != "XY" {
		t.Errorf(
			"got first element with value %s; wants element with value XY",
			list.First().Value())
	}

	list.PushFront("IK")

	if list.First().Value() != "IK" {
		t.Errorf(
			"got first element with value %s; wants element with value IK",
			list.First().Value())
	}

	if list.Len() != 3 {
		t.Errorf("got list with length %d; wants list with length 3", list.Len())
	}
}

func TestRemoveFromMiddleOfTheList(t *testing.T) {
	list := NewList()
	list.PushFront("XY")
	list.PushBack("ZW")
	list.PushBack("IK")

	list.Last().Prev().Remove()

	if list.First().Value() != "XY" || list.First().Next() != list.Last() {
		t.Error("incorrect handling with previous element on remove")
	}

	if list.Last().Value() != "IK" || list.Last().Prev() != list.First() {
		t.Error("incorrect handling with next element on remove")
	}

	if list.Len() != 2 {
		t.Errorf("got list with length %d; wants list with length 2", list.Len())
	}
}

func TestRemoveFromTheEndOfTheList(t *testing.T) {
	list := NewList()
	list.PushFront("XY")
	list.PushBack("ZW")

	list.Last().Remove()

	if list.First().Next() != nil || list.First() != list.Last() {
		t.Error("last element should be shifted on remove")
	}
}

func TestRemoveFromTheStartOfTheList(t *testing.T) {
	list := NewList()
	list.PushFront("XY")
	list.PushBack("ZW")

	list.First().Remove()

	if list.Last().Prev() != nil || list.First() != list.Last() {
		t.Error("first element should be shifted on remove")
	}
}

func TestRemoveLastElement(t *testing.T) {
	list := NewList()
	list.PushFront("XY")

	list.First().Remove()

	if list.First() != nil || list.Last() != nil {
		t.Error("the empty list should have nil as first and nil as last element")
	}

	if list.Len() != 0 {
		t.Errorf("got list with length %d; wants list with length 0", list.Len())
	}
}
