package mongorely

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrCollNotFound = errors.New("cannot find the collection of model")
)

func Count(ctx context.Context, collName string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	col := db.Collection(collName)
	if col == nil {
		return -1, ErrCollNotFound
	}

	return col.CountDocuments(ctx, filter, opts...)
}

func Create(ctx context.Context, model MongoModel, opts ...*options.InsertOneOptions) error {
	col := db.Collection(model.GetMongoCollName())
	if col == nil {
		return ErrCollNotFound
	}

	_, err := col.InsertOne(ctx, model, opts...)

	return err
}

func UpdateMany(ctx context.Context, collName string, filter interface{}, update interface{}) error {
	col := db.Collection(collName)

	_, err := col.UpdateMany(ctx, filter, update)

	return err
}

func UpdateManyWithResult(ctx context.Context, collName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	col := db.Collection(collName)
	return col.UpdateMany(ctx, filter, update)
}

func Aggregate(ctx context.Context, req *AggregationRequest) error {
	col := db.Collection(req.CollectionName)

	cursor, err := col.Aggregate(ctx, req.Pipeline, req.Options...)
	if err != nil {
		return err
	}
	fmt.Println(cursor)

	err = cursor.All(ctx, &req.Result)

	return err
}

func Find(ctx context.Context, collName string, models interface{}, filter interface{}, opts ...*options.FindOptions) error {
	col := db.Collection(collName)

	cur, err := col.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	return cur.All(ctx, models)
}

//Flush, Clear all records of collection, use it carefully
func Flush(ctx context.Context, collName string) error {
	col := db.Collection(collName)

	_, err := col.DeleteMany(ctx, bson.D{})

	return err
}
