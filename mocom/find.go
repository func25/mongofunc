package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find[T Model](ctx context.Context, filter interface{}, opts ...*options.FindOptions) (res []T, err error) {
	var t T
	cur, err := CollRead(t.CollName()).Find(ctx, filter, opts...)
	if err != nil {
		return res, err
	}

	err = cur.All(ctx, &res)
	return res, err
}

func FindOne[T Model](ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (res *T, err error) {
	res = new(T)
	var t T
	cur := CollRead(t.CollName()).FindOne(ctx, filter, opts...)
	err = cur.Decode(&res)
	return res, err
}
