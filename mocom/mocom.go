package mocom

import (
	"context"
	"errors"
	"time"

	"github.com/func25/mongofunc/v2/moper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// Count counts documents from collection
func Count(ctx context.Context, collName string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return CollRead(collName).CountDocuments(ctx, filter, opts...)
}

// CountT counts documents from collection
// model should be implement `CollName() string`
func CountT[T Modeler](ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	var t T
	return Count(ctx, t.CollName(), filter, opts...)
}

// EstimatedCount counts documents from collection but it is not accurate, still faster than Count
func EstimatedCount(ctx context.Context, collName string, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return CollRead(collName).EstimatedDocumentCount(ctx, opts...)
}

// EstimatedCountT counts documents from collection but it is not accurate, still faster than Count
// model should be implement `CollName() string`
func EstimatedCountT[T Modeler](ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	var t T
	return EstimatedCount(ctx, t.CollName(), opts...)
}

// Aggregate aggregates documents based on Pipeline and Options from request
func Aggregate(ctx context.Context, collName string, req *AggregationRequest) (res []bson.M, err error) {
	col := db.Collection(collName)

	cursor, err := col.Aggregate(ctx, req.Pipeline, req.Options...)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &res)

	return res, err
}

// Aggregate aggregates documents based on Pipeline and Options from request
// model should be implement `CollName() string`
func AggregateT[T Modeler](ctx context.Context, req *AggregationRequest) (res []bson.M, err error) {
	var t T
	return Aggregate(ctx, t.CollName(), req)
}

// Flush clears all records of collection and return number of deleted records
func Flush[T Modeler](ctx context.Context) (int64, error) {
	var t T
	result, err := db.Collection(t.CollName()).DeleteMany(ctx, moper.Query())
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, err
}

type SessionKey struct{}

// Tx executes a MongoDB transaction based on provided configuration (cfg).
// ctx is the context for the transaction, while cfg includes session configuration,
// transaction options, and the transaction function.
// Returns the result of transaction function execution or an error if the client is nil or
// any issues occurred during the session creation or transaction.
// If a nested transaction is detected, then this transaction will be executed with the passed-in context.
func Tx(ctx context.Context, cfg TransactionConfig) (interface{}, error) {
	if preSession := ctx.Value(SessionKey{}); preSession != nil {
		if mongoSession, ok := preSession.(mongo.Session); ok {
			return cfg.Func(mongo.NewSessionContext(ctx, mongoSession))
		}
	}

	if client == nil {
		return nil, errors.New("client is nil, please using mocom to create connection to mongo server or using your own client connection")
	}

	session, err := client.StartSession(cfg.SessConfig)
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	ctx = context.WithValue(ctx, SessionKey{}, session)
	return session.WithTransaction(ctx, cfg.Func, cfg.Options)
}

// TxOptimal will do the transaction with majority write-concern
// snapshot read-concern, primary read preference
//
// This should be used when transaction does not contain any read
// If a nested transaction is detected, then this transaction will be executed with the passed-in context.
func TxOptimal(ctx context.Context, f func(ctx mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	if preSession := ctx.Value(SessionKey{}); preSession != nil {
		if mongoSession, ok := preSession.(mongo.Session); ok {
			return f(mongo.NewSessionContext(ctx, mongoSession))
		}
	}

	if client == nil {
		return nil, errors.New("client is nil, please using mocom.Connect to create connection to mongo server or using your own client connection")
	}

	wc := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(5*time.Second))
	opts := options.Transaction().SetReadConcern(readconcern.Snapshot()).SetWriteConcern(wc).SetReadPreference(readpref.Primary())

	session, err := client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	ctx = context.WithValue(ctx, SessionKey{}, session)
	return session.WithTransaction(ctx, f, opts)
}
