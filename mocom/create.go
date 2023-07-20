package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(ctx context.Context, collName string, model any, opts ...*options.InsertOneOptions) (interface{}, error) {
	col := CollWrite(collName)

	result, err := col.InsertOne(ctx, model, opts...)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func CreateT[T Modeler](ctx context.Context, model T, opts ...*options.InsertOneOptions) (interface{}, error) {
	return Create(ctx, model.CollName(), model, opts...)
}

func CreateWithID[T IDModeler](ctx context.Context, model T, opts ...*options.InsertOneOptions) error {
	col := CollWrite(model.CollName())

	result, err := col.InsertOne(ctx, model, opts...)
	if err != nil {
		return err
	}

	model.SetID(result.InsertedID)
	return nil
}

func CreateMany(ctx context.Context, collName string, models []any, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	if len(models) == 0 {
		return &mongo.InsertManyResult{}, nil
	}

	col := CollWrite(collName)

	return col.InsertMany(ctx, models, opts...)
}

func CreateManyT[T Modeler](ctx context.Context, models []T, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	if len(models) == 0 {
		return &mongo.InsertManyResult{}, nil
	}

	m := make([]interface{}, len(models))
	for i := range models {
		m[i] = models[i]
	}

	return CreateMany(ctx, models[0].CollName(), m, opts...)
}
