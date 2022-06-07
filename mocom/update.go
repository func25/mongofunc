package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateOne(ctx context.Context, collName string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return db.Collection(collName).UpdateOne(ctx, filter, update, opts...)
}

func UpdateMany(ctx context.Context, collName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return db.Collection(collName).UpdateMany(ctx, filter, update)
}
