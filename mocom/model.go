package mocom

import (
	"github.com/func25/mongofunc/moper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// -------- MODEL --------
type Model interface {
	CollName() string
}

// -------- MODEL + ID --------
type IDModel interface {
	Model
	SetID(t interface{})
}

type ID = IDT[any]

type ObjectID = IDT[primitive.ObjectID]

type IntID = IDT[int]

type StringID = IDT[string]

type IDT[T any] struct {
	ID T `json:"id" bson:"_id,omitempty"`
}

func (id *IDT[T]) SetID(t T) {
	id.ID = t
}

// -------- AGGREGATION --------
type AggregationRequest[T Model] struct {
	Pipeline []moper.D
	Options  []*options.AggregateOptions
}

// -------- TRANSACTION --------
// TransactionConfig, you only need to define the func; the transaction options are not necessary for a simple application
type TransactionConfig struct {
	Options    *options.TransactionOptions
	SessConfig *options.SessionOptions
	Func       func(ctx mongo.SessionContext) (interface{}, error)
}
