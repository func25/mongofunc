package mongofunctest

import (
	"context"
	"testing"

	"github.com/func25/mongofunc/mongoquery"
	"github.com/func25/mongofunc/mongorely"
)

func TestSearchHeroes_Exists(t *testing.T) {
	ctx := context.Background()

	filter := mongoquery.Init(
		mongoquery.Exists("omit", true),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery Or failed", err)
	} else if count != int64(ROUND) {
		t.Error("mongoquery Or wrong", count, "!=", int64(ROUND))
	}
}
