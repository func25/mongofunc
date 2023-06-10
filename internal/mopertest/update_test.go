package mopertest

import (
	"context"
	"testing"

	"github.com/func25/mongofunc/v2/mocom"
	"github.com/func25/mongofunc/v2/moper"
)

func TestSet(t *testing.T) {
	ctx := context.Background()

	// change all damages to negative
	for i := 1; i <= ROUND; i++ {
		filter := moper.Query().Equal("damage", i)
		update := moper.Query().Set(moper.P{K: "damage", V: -i})

		result, err := mocom.UpdateMany[Hero](ctx, filter, update)
		if err != nil {
			t.Error("[TestSet]", err)
			return
		}

		if result.ModifiedCount != int64(i) {
			t.Error("[TestSet]:", result.ModifiedCount, "!=", i)
			return
		}
	}

	// change damages to positive
	for i := 1; i <= ROUND; i++ {
		filter := moper.Query().Equal("damage", -i)
		update := moper.Query().Set(moper.P{K: "damage", V: i})

		result, err := mocom.UpdateMany[Hero](ctx, filter, update)
		if err != nil {
			t.Error("[TestSet]", err)
			return
		}

		if result.ModifiedCount != int64(i) {
			t.Error("[TestSet]:", result.ModifiedCount, "!=", i)
			return
		}
	}
}

func TestInc(t *testing.T) {
	ctx := context.Background()
	for i := ROUND; i >= 0; i-- {
		filter := moper.Query().Equal("damage", i)

		update := moper.Query().Inc(moper.P{K: "damage", V: i})

		result, err := mocom.UpdateMany[Hero](ctx, filter, update)
		if err != nil {
			t.Error("[TestInc]", err)
			return
		}

		if result.ModifiedCount != int64(i) {
			t.Error("[TestInc]", result.ModifiedCount, "!=", i)
			return
		}
	}

	for i := 1; i <= ROUND; i++ {
		filter := moper.Query().Equal("damage", i*2)
		update := moper.Query().Inc(moper.P{
			K: "damage",
			V: -i,
		})

		result, err := mocom.UpdateMany[Hero](ctx, filter, update)
		if err != nil {
			t.Error("[TestInc]", err)
			return
		}

		if result.ModifiedCount != int64(i) {
			t.Error("[TestInc]", result.ModifiedCount, "!=", i)
			return
		}
	}
}

func TestPush(t *testing.T) {
	ctx := context.Background()

	filter := moper.Query().Equal("damage", ROUND)
	update := moper.Query().Push(moper.P{K: "skillIds", V: 6})

	result, err := mocom.UpdateMany[Hero](ctx, filter, update)
	if err != nil {
		t.Error("[TestPush]", err)
		return
	}

	filter2 := moper.Query().Equal("skillIds", []int{1, 2, 3, 4, 5, 6})

	if count, err := mocom.Count[Hero](ctx, filter2); err != nil {
		t.Error("[TestPush]", err)
		return
	} else if count != int64(ROUND) || result.ModifiedCount != int64(ROUND) {
		t.Error("[TestPush]", count, "!=", ROUND)
		return
	}
}
