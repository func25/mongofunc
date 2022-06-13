package mocom

import (
	"github.com/func25/mongofunc/moper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model interface {
	CollName() string
}

type IDModel interface {
	Model
	SetID(t interface{})
}

type ID struct {
	ID interface{} `json:"id" bson:"_id,omitempty"`
}

func (id *ID) SetID(t interface{}) {
	id.ID = t
}

type AggregationRequest[T Model] struct {
	Pipeline []moper.D
	Options  []*options.AggregateOptions
}

// TransactionConfig, you can just define the func, the transaction options are not really need with simple application
type TransactionConfig struct {
	Options    *options.TransactionOptions
	SessConfig *options.SessionOptions
	Func       func(ctx mongo.SessionContext) (interface{}, error)
}
