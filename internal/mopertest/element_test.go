package mopertest

import (
	"context"
	"testing"

	"github.com/func25/mongofunc/mocom"
	"github.com/func25/mongofunc/moper"
)

func TestExists(t *testing.T) {
	ctx := context.Background()

	filter := moper.D{}.Exists("omit", true)
	if count, err := mocom.Count[Hero](ctx, filter); err != nil {
		t.Error("[TestExists]", err)
	} else if count != int64(ROUND) {
		t.Error("[TestExists]", count, "!=", int64(ROUND))
	}

	filter = moper.D{}.Exists("omit", false)
	if count, err := mocom.Count[Hero](ctx, filter); err != nil {
		t.Error("[TestExists]", err)
	} else if count != int64(ROUND*(ROUND-1)/2) {
		t.Error("[TestExists]", count, "!=", int64(ROUND))
	}
}
