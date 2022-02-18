package mongorely

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoModel interface {
	GetMongoCollName() string
}

type DbConfig struct {
	DbName   string
	UserName string
	Password string
	Host     string
	Port     string
	// IsReplica  bool
	// ReplicaSet string
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

func (cfg *DbConfig) ToConnectionString() string {
	link := fmt.Sprintf("%s:%s/?w=majority&retryWrites=false", cfg.Host, cfg.Port)

	// if cfg.IsReplica {
	// 	link = fmt.Sprintf("%s", cfg.ReplicaSet)
	// }

	var uri string
	if cfg.UserName == "" && cfg.Password == "" {
		uri = fmt.Sprintf("mongodb://%s", link)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s@%s", cfg.UserName, cfg.Password, link)
	}

	return uri
}

// TransactionConfig, you can just define the func, the transaction options are not really need with simple application
type TransactionConfig struct {
	Options *options.TransactionOptions
	Func    func(ctx mongo.SessionContext) (interface{}, error)
}

type CustomTransactionConfig struct {
	Options *options.TransactionOptions
	Client  *mongo.Client
	Func    func(ctx mongo.SessionContext) (interface{}, error)
}
