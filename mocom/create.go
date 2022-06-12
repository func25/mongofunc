package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create[T MongoModel](ctx context.Context, model *T, opts ...*options.InsertOneOptions) (interface{}, error) {
	col := collWrite((*model).CollName())
	if result, err := col.InsertOne(ctx, model, opts...); err != nil {
		return nil, err
	} else {
		return result.InsertedID, nil
	}
}
