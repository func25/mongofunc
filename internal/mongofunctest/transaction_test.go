package mongofunctest

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/func25/mongofunc/mongoquery"
	"github.com/func25/mongofunc/mongorely"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestTransactionSuccess(t *testing.T) {
	return

	x := 2
	y := 3
	z := 4

	ctx := context.Background()

	_, err := mongorely.DoTransaction(ctx, mongorely.TransactionConfig{
		Options: &options.TransactionOptions{},
		Func: func(ctx mongo.SessionContext) (interface{}, error) {
			// update damage x to y
			filter := mongoquery.Init(mongoquery.Equal("damage", x))
			update := mongoquery.Init(mongoquery.Set(mongoquery.PairSetter{FieldName: "damage", Value: y}))

			_, err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update)
			if err != nil {
				t.Error(err)
				return nil, err
			}

			// get all hero has damage x
			heroes := []Hero{}
			err = mongorely.Find(ctx, COLLECTION_NAME, &heroes, filter)
			if err != nil {
				t.Error(err)
				return nil, err
			}

			// update damage y to z
			filter2 := mongoquery.Init(mongoquery.Equal("damage", y))
			update2 := mongoquery.Init(mongoquery.Set(mongoquery.PairSetter{FieldName: "damage", Value: z}))

			result2, err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter2, update2)
			if err != nil {
				t.Error(err)
				return nil, err
			}

			return result2, nil
		},
	})

	if err != nil {
		t.Error(err)
	}

	// get all hero has damage x or y
	filter := mongoquery.Init(mongoquery.InEllipsis("damage", x, y))
	count, err := mongorely.Count(ctx, COLLECTION_NAME, filter)
	if err != nil {
		t.Error(err)
	} else if count != 0 {
		t.Error(errors.New("Transaction was operated wrong with count 1:" + strconv.Itoa(int(count))))
	}

	// get all hero has damage z
	filter = mongoquery.Init(mongoquery.Equal("damage", z))
	count, err = mongorely.Count(ctx, COLLECTION_NAME, filter)
	if err != nil {
		t.Error(err)
	} else if int(count) != x+y+z {
		t.Error(errors.New("Transaction was operated wrong with count 2:" + strconv.Itoa(int(count))))
	}

	Clear(ctx)
	Seed(ctx, ROUND)
}

func TestTransactionFailed(t *testing.T) {
	return

	x := 2
	y := 3
	z := 4

	ctx := context.Background()

	_, err := mongorely.DoTransaction(ctx, mongorely.TransactionConfig{
		Options: &options.TransactionOptions{},
		Func: func(ctx mongo.SessionContext) (interface{}, error) {
			// update damage x to y
			filter := mongoquery.Init(mongoquery.Equal("damage", x))
			update := mongoquery.Init(mongoquery.Set(mongoquery.PairSetter{FieldName: "damage", Value: y}))

			_, err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update)
			if err != nil {
				t.Error(err)
				return nil, err
			}

			// get all hero has damage x
			heroes := []Hero{}
			err = mongorely.Find(ctx, COLLECTION_NAME, &heroes, filter)
			if err != nil {
				t.Error(err)
				return nil, err
			}

			// update damage y to z
			filter2 := mongoquery.Init(mongoquery.Equal("damage", y))
			update2 := mongoquery.Set(mongoquery.PairSetter{FieldName: "damage", Value: z})

			result2, err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter2, update2)
			if err != nil {
				return nil, err
			}

			return result2, nil
		},
	})

	if err == nil {
		t.Error(errors.New("transaction expected to failed but not"))
	}

	// get all hero has damage x or y
	filter := mongoquery.Init(mongoquery.InEllipsis("damage", x, y))
	count, err := mongorely.Count(ctx, COLLECTION_NAME, filter)
	if err != nil {
		t.Error(err)
	} else if int(count) != x+y {
		t.Error(errors.New("Transaction was operated wrong with count 1:" + strconv.Itoa(int(count))))
	}

	// get all hero has damage z
	filter = mongoquery.Init(mongoquery.Equal("damage", z))
	count, err = mongorely.Count(ctx, COLLECTION_NAME, filter)
	if err != nil {
		t.Error(err)
	} else if int(count) != z {
		t.Error(errors.New("Transaction was operated wrong with count 2:" + strconv.Itoa(int(count))))
	}
}
