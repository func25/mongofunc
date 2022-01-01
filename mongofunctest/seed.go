package mongofunctest

import (
	"context"
	"log"

	"github.com/func25/mongofunc/mongorely"
)

type Hero struct {
	mongorely.ObjectId `bson:",inline"`
	Name               string       `bson:"name"`
	Damage             int          `bson:"damage"`
	SkillIds           []int        `bson:"skillIds"`
	Omit               int          `bson:"omit,omitempty"`
	NotOmit            int          `bson:"notomit"`
	Subhero            map[int]Hero `bson:"subhero"`
}

const COLLECTION_NAME = "Heroes"

var (
	ROUND = 10
	TOTAL = -1
)

func (*Hero) GetMongoCollName() string {
	return "Heroes"
}

func init() {
	ctx := context.Background()
	_, err := mongorely.Connect(ctx, mongorely.DefaultLocalDb())
	if err != nil {
		log.Fatal("cannot connect mongo")
	}

	if err := Clear(ctx); err != nil {
		log.Fatal("cannot flush database")
	}

	if err := Seed(ctx, ROUND); err != nil {
		log.Fatal("cannot connect mongo")
	}

	TOTAL = ROUND * (ROUND + 1) / 2
}

func Seed(ctx context.Context, n int) error {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			x := 0
			if i == n-1 {
				x = 1
			}
			err := mongorely.Create(ctx, &Hero{
				Name:     "Mongorely",
				Damage:   i + 1,
				SkillIds: []int{1, 2, 3, 4, 5},
				Omit:     x,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Clear(ctx context.Context) error {
	return mongorely.Flush(ctx, COLLECTION_NAME)
}
