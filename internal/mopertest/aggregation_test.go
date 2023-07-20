package mopertest

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/func25/mongofunc/v2/mocom"
	"github.com/func25/mongofunc/v2/moper"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAggregation(t *testing.T) {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	matchStage := moper.Query().MatchD(moper.Query().InArray("damage", intArr))
	groupStage := moper.Query().Group(
		moper.P{K: "_id", V: nil},
		moper.P{K: "total", V: moper.Query().Sum("damage")},
	)

	req := &mocom.AggregationRequest{
		Pipeline: []moper.D{matchStage, groupStage},
		Options:  []*options.AggregateOptions{},
	}
	result, err := mocom.AggregateT[Hero](context.Background(), req)
	if err != nil {
		t.Error(err)
		return
	}

	expect := 0
	for _, v := range intArr {
		expect += v * v
	}

	if int(result[0]["total"].(int32)) != expect {
		t.Error("wrong result", result[0]["total"], expect)
		return
	}
}

func TestLookup(t *testing.T) {
	intArr := []int{1}
	matchStage := moper.Query().MatchD(moper.Query().InArray("damage", intArr))

	lookupStage := moper.Query().LookUp().
		From(Weapon{}.CollName()).
		LocalField("damage").
		ForeignField("damage").
		As("weapon")

	unwindStage := moper.Query().Equal("$unwind", moper.Query().Equal("path", "$weapon").Equal("preserveNullAndEmptyArrays", false))

	req := &mocom.AggregationRequest{
		Pipeline: []moper.D{matchStage, lookupStage.D(), unwindStage},
		Options:  []*options.AggregateOptions{},
	}
	result, err := mocom.AggregateT[Hero](context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	_, err = json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}

	// expect := 0
	// for _, v := range intArr {
	// 	expect += v * v
	// }

	// if result[0]["total"] != expect {
	// 	t.Error("wrong result", result[0]["total"], expect)
	// }
}
