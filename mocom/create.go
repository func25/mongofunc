package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create[T Model](ctx context.Context, model T, opts ...*options.InsertOneOptions) (interface{}, error) {
	col := collWrite(model.CollName())
	if result, err := col.InsertOne(ctx, model, opts...); err != nil {
		return nil, err
	} else {
		return result.InsertedID, nil
	}
}

func CreateWithID[T IDModel](ctx context.Context, model T, opts ...*options.InsertOneOptions) error {
	col := collWrite(model.CollName())
	if result, err := col.InsertOne(ctx, model, opts...); err != nil {
		return err
	} else {
		model.SetID(result.InsertedID)
		return nil
	}
}

func CreateMany[T Model](ctx context.Context, models []T, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	var t T
	col := collWrite(t.CollName())

	m := make([]interface{}, len(models))
	for i := range models {
		m[i] = models[i]
	}

	return col.InsertMany(ctx, m, opts...)
}
