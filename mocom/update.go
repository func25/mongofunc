package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateOne[T MongoModel](ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	var t T
	return db.Collection(t.CollName()).UpdateOne(ctx, filter, update, opts...)
}

func UpdateAndReturn[T MongoModel](ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) (ptrT *T, err error) {
	ptrT = new(T)
	res := db.Collection((*ptrT).CollName()).FindOneAndUpdate(ctx, filter, update, opts...)

	err = res.Decode(ptrT)
	if err != nil {
		return ptrT, err
	}

	return
}

func UpdateMany[T MongoModel](ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	var t T
	return db.Collection(t.CollName()).UpdateMany(ctx, filter, update)
}
