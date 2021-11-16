package mongofunctest

import (
	"context"
	"testing"

	"github.com/func25/mongofunc/mongoquery"
	"github.com/func25/mongofunc/mongorely"
)

func TestUpdateHero_Set(t *testing.T) {
	ctx := context.Background()

	filter := mongoquery.Init(
		mongoquery.Equal("damage", ROUND),
	)
	update := mongoquery.Set(mongoquery.PairSetter{
		FieldName: "damage",
		Value:     -ROUND,
	})

	if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update); err != nil {
		t.Error("mongoquery Set failed 1", err)
		return
	}

	filter2 := mongoquery.Init(
		mongoquery.Equal("damage", -ROUND),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter2); err != nil {
		t.Error("mongoquery Set failed 2", err)
		return
	} else if count != int64(ROUND) {
		t.Error("mongoquery Set wrong 3", count, "!=", ROUND)
		return
	}

	update2 := mongoquery.Set(mongoquery.PairSetter{
		FieldName: "damage",
		Value:     ROUND,
	})
	if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter2, update2); err != nil {
		t.Error("mongoquery Set failed 4", err)
		return
	}

	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery Set failed 5", err)
		return
	} else if count != int64(ROUND) {
		t.Error("mongoquery Set wrong 6", count, "!=", ROUND)
		return
	}
}

func TestUpdateHero_IncInt(t *testing.T) {
	ctx := context.Background()
	extra := 2

	filter := mongoquery.Init(
		mongoquery.Equal("damage", ROUND),
	)
	update := mongoquery.IncInt(mongoquery.PairSetterInt{
		FieldName: "damage",
		Value:     extra,
	})

	if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update); err != nil {
		t.Error("mongoquery Set failed 1", err)
		return
	}

	filter2 := mongoquery.Init(
		mongoquery.Equal("damage", ROUND+extra),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter2); err != nil {
		t.Error("mongoquery Set failed 2", err)
		return
	} else if count != int64(ROUND) {
		t.Error("mongoquery Set wrong 3", count, "!=", ROUND)
		return
	}

	update2 := mongoquery.IncInt(mongoquery.PairSetterInt{
		FieldName: "damage",
		Value:     -extra,
	})
	if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter2, update2); err != nil {
		t.Error("mongoquery Set failed 4", err)
		return
	}

	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
		t.Error("mongoquery Set failed 5", err)
		return
	} else if count != int64(ROUND) {
		t.Error("mongoquery Set wrong 6", count, "!=", ROUND)
		return
	}
}

func TestUpdateHero_Push(t *testing.T) {
	ctx := context.Background()

	filter := mongoquery.Init(
		mongoquery.Equal("damage", ROUND),
	)
	update := mongoquery.Push(mongoquery.PairSetter{FieldName: "skillIds", Value: 6})

	if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update); err != nil {
		t.Error("mongoquery Set failed 1", err)
		return
	}

	filter2 := mongoquery.Init(
		mongoquery.Equal("skillIds", []int{1, 2, 3, 4, 5, 6}),
	)
	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter2); err != nil {
		t.Error("mongoquery Set failed 2", err)
		return
	} else if count != int64(ROUND) {
		t.Error("mongoquery Set wrong 3", count, "!=", ROUND)
		return
	}
}
