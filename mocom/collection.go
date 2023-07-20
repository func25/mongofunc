package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateCollection[T Modeler](ctx context.Context, opts ...*options.CreateCollectionOptions) error {
	var t T
	return db.CreateCollection(ctx, t.CollName(), opts...)
}
