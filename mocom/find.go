package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindMany[T Model](ctx context.Context, filter interface{}, opts ...*options.FindOptions) (res []T, err error) {
	var t T
	cur, err := collRead(t.CollName()).Find(ctx, filter, opts...)
	if err != nil {
		return res, err
	}

	err = cur.All(ctx, &res)
	return res, err
}

func Find[T Model](ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (res *T, err error) {
	res = new(T)
	cur := collRead((*res).CollName()).FindOne(ctx, filter, opts...)
	err = cur.Decode(&res)
	return res, err
}
