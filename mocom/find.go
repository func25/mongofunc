package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find[T MongoModel](ctx context.Context, filter interface{}, opts ...*options.FindOptions) (res []T, err error) {
	var t T
	cur, err := db.Collection(t.CollName()).Find(ctx, filter, opts...)
	if err != nil {
		return res, err
	}

	err = cur.All(ctx, &res)
	return res, err
}

func FindOne[T MongoModel](ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (res *T, err error) {
	res = new(T)
	cur := db.Collection((*res).CollName()).FindOne(ctx, filter, opts...)
	err = cur.Decode(&res)
	return res, err
}
