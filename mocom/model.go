package mocom

import (
	"fmt"

	"github.com/func25/mongofunc/v2/moper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// -------- MODEL --------
type Modeler interface {
	CollName() string
}

// -------- MODEL + ID --------
type IDer interface {
	SetID(t any)
}

type IDModeler interface {
	Modeler
	IDer
}

type ID = IDT[any]

type ObjectID = IDT[primitive.ObjectID]

type IntID = IDT[int]

type StringID = IDT[string]

type IDT[T any] struct {
	ID T `json:"id" bson:"_id,omitempty"`
}

func (id *IDT[T]) SetID(t any) {
	var ok bool

	id.ID, ok = t.(T)
	if !ok {
		fmt.Println("SetID: type assertion failed")
	}
}

// -------- AGGREGATION --------
type AggregationRequest struct {
	Pipeline []moper.D
	Options  []*options.AggregateOptions
}

// -------- TRANSACTION --------
// TransactionConfig, you only need to define the func; the transaction options are not necessary to be defined for simple use cases
type TransactionConfig struct {
	Options    *options.TransactionOptions
	SessConfig *options.SessionOptions
	Func       func(ctx mongo.SessionContext) (interface{}, error)
}
