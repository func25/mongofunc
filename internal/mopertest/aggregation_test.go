package mopertest

import (
	"context"
	"fmt"
	"testing"

	"github.com/func25/mongofunc/mocom"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAggregationTest(t *testing.T) {
	matchStage := bson.D{
		{
			Key:   "$match",
			Value: bson.D{
				// moper.In("damage", 1, 2, 3, 4, 5, 6, 7, 8, 9),
			},
		},
	}

	groupStage := bson.D{
		{
			Key: "$group",
			Value: bson.D{
				{
					Key:   "_id",
					Value: nil,
				},
				{
					Key: "total",
					Value: bson.D{
						{
							Key:   "$sum",
							Value: "$damage",
						},
					},
				},
			},
		},
	}

	req := &mocom.AggregationRequest{
		CollectionName: COLLECTION_NAME,
		Pipeline:       mongo.Pipeline{matchStage, groupStage},
		Options:        []*options.AggregateOptions{},
	}
	if err := mocom.Aggregate(context.Background(), req); err != nil {
		t.Error(err)
		return
	}

	fmt.Println(req.Result)
}
