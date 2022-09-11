package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddIndex[T Model](ctx context.Context, index mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	var t T
	return CollWrite(t.CollName()).Indexes().CreateOne(ctx, index, opts...)
}

func AddIndexes[T Model](ctx context.Context, indexes []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	var t T
	return CollWrite(t.CollName()).Indexes().CreateMany(ctx, indexes, opts...)
}
