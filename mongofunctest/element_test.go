package mongofunctest

import (
	"context"
	"testing"

	"github.com/func25/mongofunc/mongoquery"
	"github.com/func25/mongofunc/mongorely"
)

func TestSearchHeroes_Or(t *testing.T) {
	ctx := context.Background()

	num1 := 2
	num2 := 3

	filter := mongoquery.Init(
		mongoquery.Or(
			mongoquery.Init(mongoquery.Equal("damage", num1)),
			mongoquery.Init(mongoquery.Equal("damage", num2)),
		),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery Or failed", err)
	} else if count != int64(num1+num2) {
		t.Error("mongoquery Or wrong", count, "!=", int64(num1+num2))
	}
}
