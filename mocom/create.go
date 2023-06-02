package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Ptr[T any] interface {
// 	*T
// }

func Create[T Model](ctx context.Context, model T, opts ...*options.InsertOneOptions) (interface{}, error) {
	col := CollWrite(model.CollName())

	result, err := col.InsertOne(ctx, model, opts...)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func CreateWithID[T IDModel](ctx context.Context, model T, opts ...*options.InsertOneOptions) error {
	col := CollWrite(model.CollName())

	result, err := col.InsertOne(ctx, model, opts...)
	if err != nil {
		return err
	}

	model.SetID(result.InsertedID)
	return nil
}

func CreateMany[T Model](ctx context.Context, models []T, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	if len(models) == 0 {
		return &mongo.InsertManyResult{}, nil
	}

	col := CollWrite(models[0].CollName())

	m := make([]interface{}, len(models))
	for i := range models {
		m[i] = models[i]
	}

	return col.InsertMany(ctx, m, opts...)
}
