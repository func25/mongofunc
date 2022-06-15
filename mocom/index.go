package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddIndex[T Model](ctx context.Context, index mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	var t T
	col := collWrite(t.CollName())
	return col.Indexes().CreateOne(ctx, index, opts...)
}

func AddIndexes[T Model](ctx context.Context, indexes []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	var t T
	col := collWrite(t.CollName())
	return col.Indexes().CreateMany(ctx, indexes, opts...)
}
