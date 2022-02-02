package mongorely

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(ctx context.Context, model MongoModel, opts ...*options.InsertOneOptions) (interface{}, error) {
	col := db.Collection(model.GetMongoCollName())
	result, err := col.InsertOne(ctx, model, opts...)
	return result.InsertedID, err
}

func CreateCustom(ctx context.Context, collectionName string, model interface{}, opts ...*options.InsertOneOptions) (interface{}, error) {
	col := db.Collection(collectionName)
	result, err := col.InsertOne(ctx, model, opts...)
	return result.InsertedID, err
}
