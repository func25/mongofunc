package mongoquery_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/func25/mongofunc/mongoquery"
	"github.com/func25/mongofunc/mongorely"
	"github.com/func25/mongofunc/mongoseed"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

const (
	N               = 10
	DOCUMENT_COUNT  = N * (N + 1) / 2
	COLLECTION_NAME = "Heroes"
)

func init() {
	ctx := context.Background()
	var err error
	client, err = mongorely.Connect(ctx, mongorely.DefaultLocalDb())
	if err != nil {
		log.Fatal(err)
		return
	}

	value, err := mongorely.Count(ctx, COLLECTION_NAME, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if value != DOCUMENT_COUNT {
		mongorely.Flush(ctx, COLLECTION_NAME)
		mongoseed.Seed(ctx, N)
	}
}

func TestMIn(t *testing.T) {
	ctx := context.Background()

	if value, err := mongorely.Count(ctx, COLLECTION_NAME, mongoquery.MIn("damage", 5, 3)); value != 5+3 || err != nil {
		t.Error("something wrong", value, err)
	} else {
		fmt.Println(value)
	}
}

func TestMNotIn(t *testing.T) {
	ctx := context.Background()

	if value, err := mongorely.Count(ctx, COLLECTION_NAME, mongoquery.MNotIn("damage", 5, 6)); value != DOCUMENT_COUNT-5-6 || err != nil {
		t.Error("something wrong", value, err)
	} else {
		fmt.Println(value)
	}

}
