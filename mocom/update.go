package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateOne(ctx context.Context, collName string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return CollWrite(collName).UpdateOne(ctx, filter, update, opts...)
}

// UpdateOne updates one document from collection
func UpdateOneT[T Modeler](ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	var t T
	return UpdateOne(ctx, t.CollName(), filter, update, opts...)
}

// UpdateAndReturn updates one document from collection and return the updated document
// Pass options to return document AFTER or BEFORE the update
func UpdateAndReturn[T Modeler](ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) (*T, error) {
	var t T
	res := CollWrite(t.CollName()).FindOneAndUpdate(ctx, filter, update, opts...)

	err := res.Decode(&t)
	return &t, err
}

func UpdateMany(ctx context.Context, collName string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return CollWrite(collName).UpdateMany(ctx, filter, update, opts...)
}

// UpdateMany updates many documents from collection
// model should be implement `CollName() string`
func UpdateManyT[T Modeler](ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	var t T
	return UpdateMany(ctx, t.CollName(), filter, update, opts...)
}
