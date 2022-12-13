package mocom

import (
	"context"
	"errors"
	"time"

	"github.com/func25/mongofunc/moper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func Count[T Model](ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	var t T
	return CollRead(t.CollName()).CountDocuments(ctx, filter, opts...)
}

func EstimatedCount[T Model](ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	var t T
	return CollRead(t.CollName()).EstimatedDocumentCount(ctx, opts...)
}

func Aggregate[T Model](ctx context.Context, req *AggregationRequest[T]) (res []bson.M, err error) {
	var t T
	col := db.Collection(t.CollName())

	cursor, err := col.Aggregate(ctx, req.Pipeline, req.Options...)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &res)

	return res, err
}

// Flush clears all records of collection and return number of deleted records
func Flush[T Model](ctx context.Context) (int64, error) {
	var t T
	result, err := db.Collection(t.CollName()).DeleteMany(ctx, moper.NewD())
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, err
}

// Tx -> transaction
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

// TxOptimal will do the transaction with majority write-concern
// snapshot read-concern, primary read preference
//
// This should be used when transaction does not contain any read
func TxOptimal(ctx context.Context, f func(ctx mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	if client == nil {
		return nil, errors.New("client is nil, please using mocom to create connection to mongo server or using your own client connection")
	}

	wc := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(5*time.Second))
	opts := options.Transaction().SetReadConcern(readconcern.Snapshot()).SetWriteConcern(wc).SetReadPreference(readpref.Primary())

	session, err := client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	return session.WithTransaction(ctx, f, opts)
}
