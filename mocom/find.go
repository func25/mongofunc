package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindT finds documents from from model's collection
// model should be implement `CollName() string`
func FindT[T Modeler](ctx context.Context, filter interface{}, opts ...*options.FindOptions) (res []T, err error) {
	var t T
	cur, err := CollRead(t.CollName()).Find(ctx, filter, opts...)
	if err != nil {
		return res, err
	}

	err = cur.All(ctx, &res)
	return res, err
}

// FindOneT finds one document from model's collection
// model should be implement `CollName() string`
func FindOneT[T Modeler](ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (res *T, err error) {
	res = new(T)
	var t T
	cur := CollRead(t.CollName()).FindOne(ctx, filter, opts...)
	err = cur.Decode(&res)
	return res, err
}
