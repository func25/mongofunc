package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddIndex adds an index to collection
func AddIndex(ctx context.Context, collName string, index mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return CollWrite(collName).Indexes().CreateOne(ctx, index, opts...)
}

// AddIndex adds an index to collection
// model should be implement `CollName() string`
func AddIndexT[T Modeler](ctx context.Context, index mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	var t T
	return AddIndex(ctx, t.CollName(), index, opts...)
}

// AddIndexes adds indexes to collection
func AddIndexes(ctx context.Context, collName string, indexes []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	return CollWrite(collName).Indexes().CreateMany(ctx, indexes, opts...)
}

// AddIndexes adds indexes to collection
// model should be implement `CollName() string`
func AddIndexesT[T Modeler](ctx context.Context, indexes []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	var t T
	return AddIndexes(ctx, t.CollName(), indexes, opts...)
}
