package mopertest

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/func25/mongofunc/mocom"
	"github.com/func25/mongofunc/moper"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAggregation(t *testing.T) {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	matchStage := moper.D{}.MatchD(moper.D{}.InArray("damage", intArr))
	groupStage := moper.D{}.Group(
		moper.P{K: "_id", V: nil},
		moper.P{K: "total", V: moper.D{}.Sum("damage")},
	)

	req := &mocom.AggregationRequest[Hero]{
		Pipeline: []moper.D{matchStage, groupStage},
		Options:  []*options.AggregateOptions{},
	}
	result, err := mocom.Aggregate(context.Background(), req)
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
	matchStage := moper.D{}.MatchD(moper.D{}.InArray("damage", intArr))

	lookupStage := moper.D{}.LookUp().
		From(Weapon{}.CollName()).
		LocalField("damage").
		ForeignField("damage").
		As("weapon")

	unwindStage := moper.D{}.Equal("$unwind", moper.D{}.Equal("path", "$weapon").Equal("preserveNullAndEmptyArrays", false))

	req := &mocom.AggregationRequest[Hero]{
		Pipeline: []moper.D{matchStage, lookupStage.D(), unwindStage},
		Options:  []*options.AggregateOptions{},
	}
	result, err := mocom.Aggregate(context.Background(), req)
	if err != nil {
		t.Error(err)
		return
	}

	x, err := json.Marshal(result)
	fmt.Println(string(x))

	// expect := 0
	// for _, v := range intArr {
	// 	expect += v * v
	// }

	// if result[0]["total"] != expect {
	// 	t.Error("wrong result", result[0]["total"], expect)
	// }
}
