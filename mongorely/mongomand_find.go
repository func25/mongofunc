package mongorely

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find(ctx context.Context, collName string, models interface{}, filter interface{}, opts ...*options.FindOptions) error {
	cur, err := db.Collection(collName).Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	return cur.All(ctx, models)
}

func FindOne(ctx context.Context, collName string, model interface{}, filter interface{}, opts ...*options.FindOneOptions) error {
	cur := db.Collection(collName).FindOne(ctx, filter, opts...)
	return cur.Decode(model)
}
