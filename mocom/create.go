package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(ctx context.Context, model MongoModel, opts ...*options.InsertOneOptions) (interface{}, error) {
	col := db.Collection(model.GetCollName())
	if result, err := col.InsertOne(ctx, model, opts...); err != nil {
		return nil, err
	} else {
		return result.InsertedID, nil
	}
}

func CreateCustom(ctx context.Context, collectionName string, model interface{}, opts ...*options.InsertOneOptions) (interface{}, error) {
	col := db.Collection(collectionName)
	if result, err := col.InsertOne(ctx, model, opts...); err != nil {
		return nil, err
	} else {
		return result.InsertedID, err
	}
}
