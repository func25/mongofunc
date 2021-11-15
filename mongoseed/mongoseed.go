package mongoseed

import (
	"context"

	"github.com/func25/mongofunc/monmongofunc/mongorely"
)

type Hero struct {
	mongorely.ObjectId `bson:",inline"`
	Name               string `bson:"name"`
	Damage             int    `bson:"damage"`
	SkillIds           []int  `bson:"skillIds"`
}

func (*Hero) GetMongoCollName() string {
	return "Heroes"
}

func Seed(ctx context.Context, n int) error {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			err := mongorely.Create(ctx, &Hero{
				Name:     "Mongorely",
				Damage:   i + 1,
				SkillIds: []int{1, 2, 3, 4, 5},
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
