package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteOne[T Model](ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	var t T
	return collWrite(t.CollName()).DeleteOne(ctx, filter, opts...)
}

func Delete[T Model](ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	var t T
	return collWrite(t.CollName()).DeleteMany(ctx, filter, opts...)
}
