package mongofunctest

import (
	"context"
	"testing"

	"github.com/func25/mongofunc/mongoquery"
	"github.com/func25/mongofunc/mongorely"
)

func TestSearchHeroes_Equal(t *testing.T) {
	ctx := context.Background()
	num := 4
	filter := mongoquery.Init(
		mongoquery.Equal("damage", num),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery IN failed", err)
	} else if count != int64(num) {
		t.Error("mongoquery IN wrong", count, "!=", num)
	}
}

func TestSearchHeroes_NotEqual(t *testing.T) {
	ctx := context.Background()
	num := 4
	filter := mongoquery.Init(
		mongoquery.NotEqual("damage", num),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery IN failed", err)
	} else if count != int64(TOTAL-num) {
		t.Error("mongoquery IN wrong", count, "!=", TOTAL-num)
	}
}

func TestSearchHeroes_In(t *testing.T) {
	ctx := context.Background()
	exactCount1 := 2
	exactCount2 := 3
	filter := mongoquery.Init(
		mongoquery.InEllipsis("damage", exactCount1, exactCount2),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery IN failed", err)
	} else if count != int64(exactCount1+exactCount2) {
		t.Error("mongoquery IN wrong", count, "!=", exactCount1+exactCount2)
	}
}

func TestSearchHeroes_NotIn(t *testing.T) {
	ctx := context.Background()
	exactCount1 := 2
	exactCount2 := 3
	filter := mongoquery.Init(
		mongoquery.NotInEllipsis("damage", exactCount1, exactCount2),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery IN failed", err)
	} else if count != int64(TOTAL-exactCount1-exactCount2) {
		t.Error("mongoquery IN wrong", count, "!=", TOTAL-exactCount1-exactCount2)
	}
}

func TestSearchHeroes_EqualLess(t *testing.T) {
	ctx := context.Background()
	compNum := 3
	num := compNum * (compNum + 1) / 2

	filter := mongoquery.Init(
		mongoquery.EqualLess("damage", compNum),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery EqualLess failed", err)
	} else if count != int64(num) {
		t.Error("mongoquery EqualLess wrong", count, "!=", num)
	}
}

func TestSearchHeroes_GreaterInt(t *testing.T) {
	ctx := context.Background()
	compNum := 3
	num := TOTAL - compNum*(compNum+1)/2

	filter := mongoquery.Init(
		mongoquery.Greater("damage", compNum),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery GreaterInt failed", err)
	} else if count != int64(num) {
		t.Error("mongoquery GreaterInt wrong", count, "!=", num)
	}
}

func TestSearchHeroes_Less(t *testing.T) {
	ctx := context.Background()
	compNum := 3
	num := compNum * (compNum - 1) / 2

	filter := mongoquery.Init(
		mongoquery.Less("damage", compNum),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery Less failed", err)
	} else if count != int64(num) {
		t.Error("mongoquery Less wrong", count, "!=", num)
	}
}

func TestSearchHeroes_EqualGreaterInt(t *testing.T) {
	ctx := context.Background()
	compNum := 3
	num := TOTAL - compNum*(compNum-1)/2

	filter := mongoquery.Init(
		mongoquery.EqualGreater("damage", compNum),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery EqualGreaterInt failed", err)
	} else if count != int64(num) {
		t.Error("mongoquery EqualGreaterInt wrong", count, "!=", num)
	}
}
