package mopertest

import (
	"context"
	"log"
	"strconv"

	"github.com/func25/mongofunc/mocom"
)

// weapon
type Weapon struct {
	mocom.ID `bson:",inline"`
	Type     int `json:"type" bson:"type"`
	Damage   int `json:"damage" bson:"damage"`
}

func (Weapon) CollName() string {
	return "Weapons"
}

// hero
type Hero struct {
	mocom.ID `bson:",inline"`
	WeaponID interface{} `bson:"weaponId"`
	Name     string      `bson:"name"`
	Damage   int         `bson:"damage"`
	SkillIds []int       `bson:"skillIds"`
	Omit     bool        `bson:"omit,omitempty"`
}

var (
	ROUND = 10
	TOTAL = -1
)

func (Hero) CollName() string {
	return "Heroes"
}

// %s:%s/?w=majority&retryWrites=false
func init() {
	ctx := context.Background()
	_, err := mocom.Connect(ctx, "mongodb://localhost:27017/?w=majority&retryWrites=false", "defaultdb")
	if err != nil {
		log.Fatal("cannot connect mongo", err)
	}

	if err := Clear(ctx); err != nil {
		log.Fatalln("cannot flush database", err)
	}

	if err := Seed(ctx, ROUND); err != nil {
		log.Fatal("cannot connect mongo")
	}

	TOTAL = ROUND * (ROUND + 1) / 2
}

//Seed create 1 hero has 1 damage, 2 heroes have 2 damages,... until n (n == 10)
func Seed(ctx context.Context, n int) error {
	count := 0
	weapons := []*Weapon{}
	for i := 0; i < 3; i++ {
		x := &Weapon{
			Type:   i,
			Damage: 1,
		}
		err := mocom.CreateWithID(ctx, x)
		if err != nil {
			return err
		}
		weapons = append(weapons, x)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j, count = j+1, count+1 {
			_, err := mocom.Create(ctx, &Hero{
				Name:     "hero" + strconv.Itoa(count),
				Damage:   i + 1,
				SkillIds: []int{1, 2, 3, 4, 5},
				Omit:     j == i,
				WeaponID: weapons[j%3].ID.ID,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Clear(ctx context.Context) error {
	_, err := mocom.Flush[Hero](ctx)
	_, err = mocom.Flush[Weapon](ctx)
	return err
}
