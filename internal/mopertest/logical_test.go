package mopertest

import (
	"context"
	"testing"

	"github.com/func25/mongofunc/mocom"
	"github.com/func25/mongofunc/moper"
)

func TestOr(t *testing.T) {
	ctx := context.Background()

	for i := 1; i < ROUND; i++ {
		for j := i + 1; j <= ROUND; j++ {
			filter := moper.NewD().Or(
				*moper.NewD().Equal("damage", i),
				*moper.NewD().Equal("damage", j),
			)
			if count, err := mocom.Count[Hero](ctx, filter); err != nil {
				t.Error("[TestOr]", err)
			} else if count != int64(i+j) {
				t.Error("[TestOr]", count, "!=", int64(i+j))
			}
		}
	}
}
