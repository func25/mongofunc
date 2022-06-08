package mocom

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
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
	// fmt.Println(cursor)

	err = cursor.All(ctx, &req.Result)

	return err
}

//Flush clears all records of collection and return number of deleted records, use it carefully
func Flush(ctx context.Context, collName string) (int64, error) {
	result, err := db.Collection(collName).DeleteMany(ctx, bson.D{})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, err
}

//Tx -> transaction
func Tx(ctx context.Context, cfg TransactionConfig) (interface{}, error) {
	if client == nil {
		return nil, errors.New("client is nil, please using mocom to create connection to mongo server or using your own client connection")
	}

	session, err := client.StartSession(cfg.SessConfig)
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	return session.WithTransaction(ctx, cfg.Func, cfg.Options)
}

// TxOptimal will do the transaction with majority write-concern and local read-concern, client default read pref
func TxOptimal(ctx context.Context, f func(ctx mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	if client == nil {
		return nil, errors.New("client is nil, please using mocom to create connection to mongo server or using your own client connection")
	}

	wc := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(5*time.Second))
	opts := options.Transaction().SetReadConcern(readconcern.Local()).SetWriteConcern(wc)

	session, err := client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	return session.WithTransaction(ctx, f, opts)
}
