package monmongofunctest

import (
	"context"
	"errors"
	"testing"

	"github.com/func25/mongofunc/monmongofunc/mongoquery"
	"github.com/func25/mongofunc/monmongofunc/mongorely"
	"github.com/func25/mongofunc/monmongofunc/mongoseed"

	"go.mongodb.org/mongo-driver/bson"
)

func TestSeed(t *testing.T) {
	// Connect to mongo
	ctx := context.Background()
	client, err := mongorely.Connect(ctx, mongorely.DefaultLocalDb())
	if err != nil {
		t.Error(errors.New("cannot connect to mongodb"))
		return
	}
	defer client.Disconnect(ctx)

	// err = mongoseed.Seed(ctx, 1000)
	// if err != nil {
	// 	t.Error(err)
	// }

	SearchHeroesWithBIn(ctx)
}

func BenchmarkInt(b *testing.B) {
	// Connect to mongo
	ctx := context.Background()
	client, err := mongorely.Connect(ctx, mongorely.DefaultLocalDb())
	if err != nil {
		return
	}
	defer client.Disconnect(ctx)

	for i := 0; i < b.N; i++ {
		SearchHeroesWithBInt(ctx)
	}
}

func BenchmarkInterface(b *testing.B) {
	// Connect to mongo
	ctx := context.Background()
	client, err := mongorely.Connect(ctx, mongorely.DefaultLocalDb())
	if err != nil {
		return
	}
	defer client.Disconnect(ctx)

	for i := 0; i < b.N; i++ {
		SearchHeroesWithBIn(ctx)
	}
}

func SearchHeroesWithBIn(ctx context.Context) {
	filter := mongoquery.MIn("damage", 0, 2)
	heroes := []mongoseed.Hero{}
	mongorely.Find(ctx, "Heroes", &heroes, filter)
	// fmt.Println(len(heroes))
}

func SearchHeroesWithBInt(ctx context.Context) {
	filter := MInt("damage", 0, 2)
	heroes := []mongoseed.Hero{}
	mongorely.Find(ctx, "Heroes", &heroes, filter)
	// fmt.Println(len(heroes))
}

func Int(fieldName string, value ...int) bson.M {
	filter := bson.M{fieldName: bson.M{"$in": value}}
	return filter
}

func NotInt(fieldName string, value ...int) bson.M {
	filter := bson.M{fieldName: bson.M{"$in": value}}
	return filter
}
