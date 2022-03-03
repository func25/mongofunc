package mongorely

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Count(ctx context.Context, collName string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return db.Collection(collName).CountDocuments(ctx, filter, opts...)
}

func EstimatedCount(ctx context.Context, collName string, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return db.Collection(collName).EstimatedDocumentCount(ctx, opts...)
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

//Flush, Clear all records of collection and return number of deleted records, use it carefully
func Flush(ctx context.Context, collName string) (int64, error) {
	result, err := db.Collection(collName).DeleteMany(ctx, bson.D{})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, err
}

func DoTransaction(ctx context.Context, cfg TransactionConfig) (interface{}, error) {
	if cfg.Client == nil {
		cfg.Client = client
	}

	if client == nil {
		return nil, errors.New("client is nil, please using mongorely to create connection to mongo server or using your own client connection")
	}

	session, err := cfg.Client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	return session.WithTransaction(ctx, cfg.Func, cfg.Options)
}
