package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReplaceOne replaces one document from collection
// upsert = true if you want to upsert or just use mocom.AddOrUpdate
func ReplaceOne(ctx context.Context, collName string, filter interface{}, model any, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return CollWrite(collName).ReplaceOne(ctx, filter, model, opts...)
}

// ReplaceOneT replaces one document from collection
// upsert = true if you want to upsert
// model should be implement `CollName() string`
func ReplaceOneT[T Modeler](ctx context.Context, filter interface{}, model T, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return ReplaceOne(ctx, model.CollName(), filter, model, opts...)
}
