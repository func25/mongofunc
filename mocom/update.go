package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateOne[T Model](ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	var t T
	return CollWrite(t.CollName()).UpdateOne(ctx, filter, update, opts...)
}

func UpdateAndReturn[T Model](ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) (ptrT *T, err error) {
	var t T
	res := CollWrite(t.CollName()).FindOneAndUpdate(ctx, filter, update, opts...)

	err = res.Decode(ptrT)
	if err != nil {
		return ptrT, err
	}

	return
}

func UpdateMany[T Model](ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	var t T
	return CollWrite(t.CollName()).UpdateMany(ctx, filter, update, opts...)
}
