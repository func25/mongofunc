package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// REPLACE
func ReplaceOne[T Model](ctx context.Context, filter interface{}, model T, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return CollWrite(model.CollName()).ReplaceOne(ctx, filter, model, opts...)
}

func AddOrUpdate[T Model](ctx context.Context, filter interface{}, model T) (*mongo.UpdateResult, error) {
	return CollWrite(model.CollName()).ReplaceOne(ctx, filter, model, options.Replace().SetUpsert(true))
}
