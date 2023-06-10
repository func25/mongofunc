package mopertest

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/func25/mongofunc/v2/mocom"
	"github.com/func25/mongofunc/v2/moper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestTransactionSuccess(t *testing.T) {
	ctx := context.Background()
	x := 2
	y := 3
	z := 4

	_, err := mocom.Tx(ctx, mocom.TransactionConfig{
		Options: &options.TransactionOptions{},
		Func: func(ctx mongo.SessionContext) (interface{}, error) {
			filter := moper.Query().Equal("damage", x)
			update := moper.Query().Set(moper.P{K: "damage", V: y})

			_, err := mocom.UpdateMany[Hero](ctx, filter, update)
			if err != nil {
				return nil, err
			}

			// update damage y to z
			filter2 := moper.Query().Equal("damage", y)
			update2 := moper.Query().Set(moper.P{K: "damage", V: z})

			result2, err := mocom.UpdateMany[Hero](ctx, filter2, update2)
			if err != nil {
				return nil, err
			}

			return result2, nil
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	// get all hero has damage x or y
	filter := moper.Query().InEll("damage", x, y)
	count, err := mocom.Count[Hero](ctx, filter)
	if err != nil {
		t.Error(err)
	} else if count != 0 {
		t.Error(errors.New("Transaction was operated wrong with count 1:" + strconv.Itoa(int(count))))
	}

	// get all hero has damage z
	filter = moper.Query().Equal("damage", z)
	count, err = mocom.Count[Hero](ctx, filter)
	if err != nil {
		t.Error(err)
	} else if int(count) != x+y+z {
		t.Error(errors.New("Transaction was operated wrong with count 2:" + strconv.Itoa(int(count))))
	}

	Clear(ctx)
	Seed(ctx, ROUND)
}

func TestNestedTransactionSuccess(t *testing.T) {
	ctx := context.Background()
	x := 2
	y := 3
	z := 4

	_, err := mocom.Tx(ctx, mocom.TransactionConfig{
		Options: &options.TransactionOptions{},
		Func: func(ctx mongo.SessionContext) (interface{}, error) {
			return NestedTransaction(ctx, x, y, z)
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	// get all hero has damage x or y
	filter := moper.Query().InEll("damage", x, y)
	count, err := mocom.Count[Hero](ctx, filter)
	if err != nil {
		t.Error(err)
	} else if count != 0 {
		t.Error(errors.New("Transaction was operated wrong with count 1:" + strconv.Itoa(int(count))))
	}

	// get all hero has damage z
	filter = moper.Query().Equal("damage", z)
	count, err = mocom.Count[Hero](ctx, filter)
	if err != nil {
		t.Error(err)
	} else if int(count) != x+y+z {
		t.Error(errors.New("Transaction was operated wrong with count 2:" + strconv.Itoa(int(count))))
	}

	Clear(ctx)
	Seed(ctx, ROUND)
}

func NestedTransaction(ctx context.Context, x, y, z int) (any, error) {
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	return mocom.TxOptimal(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		filter := moper.Query().Equal("damage", x)
		update := moper.Query().Set(moper.P{K: "damage", V: y})

		_, err := mocom.UpdateMany[Hero](ctx, filter, update)
		if err != nil {
			return nil, err
		}

		// update damage y to z
		filter2 := moper.Query().Equal("damage", y)
		update2 := moper.Query().Set(moper.P{K: "damage", V: z})

		result2, err := mocom.UpdateMany[Hero](ctx, filter2, update2)
		if err != nil {
			return nil, err
		}

		return result2, nil
	})
}

func TestTransactionFailed(t *testing.T) {
	x := 2
	y := 3
	z := 4

	ctx := context.Background()

	_, err := mocom.Tx(ctx, mocom.TransactionConfig{
		Options: &options.TransactionOptions{},
		Func: func(ctx mongo.SessionContext) (interface{}, error) {
			// update damage x to y
			filter := moper.Query().Equal("damage", x)
			update := moper.Query().Set(moper.P{K: "damage", V: y})

			_, err := mocom.UpdateMany[Hero](ctx, filter, update)
			if err != nil {
				t.Error(err)
				return nil, err
			}

			// update damage y to z
			filter2 := moper.Query().Equal("damage", y)
			update2 := moper.Query().Set(moper.P{K: "damage", V: z})

			_, err = mocom.UpdateMany[Hero](ctx, filter2, update2)
			if err != nil {
				return nil, err
			}

			return nil, errors.New("fake error")
		},
	})

	if err == nil {
		t.Error(errors.New("transaction expected to failed but not"))
		return
	}

	// get all hero has damage x or y
	filter := moper.Query().InEll("damage", x, y)
	count, err := mocom.Count[Hero](ctx, filter)
	if err != nil {
		t.Error(err)
	} else if int(count) != x+y {
		t.Error(errors.New("Transaction was operated wrong with count 1:" + strconv.Itoa(int(count))))
	}

	// get all hero has damage z
	filter = moper.Query().Equal("damage", z)
	count, err = mocom.Count[Hero](ctx, filter)
	if err != nil {
		t.Error(err)
	} else if int(count) != z {
		t.Error(errors.New("Transaction was operated wrong with count 2:" + strconv.Itoa(int(count))))
	}
}
