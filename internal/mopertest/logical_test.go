package mopertest

import (
	"context"
	"testing"

	"github.com/func25/mongofunc/v2/mocom"
	"github.com/func25/mongofunc/v2/moper"
)

func TestOr(t *testing.T) {
	ctx := context.Background()

	for i := 1; i < ROUND; i++ {
		for j := i + 1; j <= ROUND; j++ {
			filter := moper.Query().Or(
				moper.Query().Equal("damage", i),
				moper.Query().Equal("damage", j),
			)
			if count, err := mocom.Count[Hero](ctx, filter); err != nil {
				t.Error("[TestOr]", err)
			} else if count != int64(i+j) {
				t.Error("[TestOr]", count, "!=", int64(i+j))
			}
		}
	}
}
