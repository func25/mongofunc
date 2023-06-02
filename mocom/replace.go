package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// REPLACE
func ReplaceOne[T Model](ctx context.Context, filter interface{}, model interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	var t T
	return CollWrite(t.CollName()).ReplaceOne(ctx, filter, model, opts...)
}
