package mocom

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoModel interface {
	GetCollName() string
}

type ObjectId struct {
	ID interface{} `json:"id" bson:"_id,omitempty"`
}

type AggregationRequest struct {
	CollectionName string
	Result         []bson.M
	Pipeline       mongo.Pipeline
	Options        []*options.AggregateOptions
}

// TransactionConfig, you can just define the func, the transaction options are not really need with simple application
type TransactionConfig struct {
	Options    *options.TransactionOptions
	SessConfig *options.SessionOptions
	Func       func(ctx mongo.SessionContext) (interface{}, error)
}
