package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DeleteOne deletes one document from collection
func DeleteOne(ctx context.Context, collName string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return CollWrite(collName).DeleteOne(ctx, filter, opts...)
}

// DeleteOne deletes one document from collection
func DeleteOneT[T Modeler](ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	var t T
	return DeleteOne(ctx, t.CollName(), filter, opts...)
}

// Delete deletes many documents from collection
func Delete(ctx context.Context, collName string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return CollWrite(collName).DeleteMany(ctx, filter, opts...)
}

// Delete deletes many documents from collection
func DeleteT[T Modeler](ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	var t T
	return Delete(ctx, t.CollName(), filter, opts...)
}
