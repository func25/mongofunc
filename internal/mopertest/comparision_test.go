package mopertest

import (
	"context"
	"testing"

	"github.com/func25/mongofunc/v2/mocom"
	"github.com/func25/mongofunc/v2/moper"
)

func TestEquals(t *testing.T) {
	ctx := context.Background()
	for i := 0; i < ROUND; i++ {
		filter := moper.Query().Equal("damage", i+1)

		if count, err := mocom.CountT[Hero](ctx, filter); err != nil {
			t.Error("[TestEquals]:", err)
		} else if count != int64(i+1) {
			t.Error("[TestEquals]:", count, "!=", i+1)
		}
	}
}

func TestNotEquals(t *testing.T) {
	ctx := context.Background()

	for i := 0; i < ROUND; i++ {
		num := i + 1

		filter := moper.Query().NotEqual("damage", num)

		if count, err := mocom.CountT[Hero](ctx, filter); err != nil {
			t.Error("[TestNotEquals]:", err)
		} else if count != int64(TOTAL-num) {
			t.Error("[TestNotEquals]:", count, "!=", TOTAL-num)
		}
	}

}

func TestIn(t *testing.T) {
	ctx := context.Background()

	for i := 0; i < ROUND; i++ {
		dmg1 := i + 1

		for j := 0; j < i; j++ {
			dmg2 := j + 1

			filter := moper.Query().InEll("damage", dmg1, dmg2)

			if count, err := mocom.CountT[Hero](ctx, filter); err != nil {
				t.Error("[TestIn]", err)
			} else if count != int64(dmg1+dmg2) {
				t.Error("[TestNotIn]", count, "!=", dmg1+dmg2)
			}
		}
	}
}

func TestNotIn(t *testing.T) {
	ctx := context.Background()

	for i := 0; i < ROUND; i++ {
		dmg1 := i + 1

		for j := 0; j < i; j++ {
			dmg2 := j + 1

			filter := moper.Query().NotInEll("damage", dmg1, dmg2)

			if count, err := mocom.CountT[Hero](ctx, filter); err != nil {
				t.Error("[TestNotIn]:", err)
			} else if count != int64(TOTAL-dmg1-dmg2) {
				t.Error("[TestNotIn]:", count, "!=", TOTAL-dmg1-dmg2)
			}
		}
	}
}

func TestLess(t *testing.T) {
	ctx := context.Background()

	for i := 0; i <= ROUND; i++ {
		num := i * (i - 1) / 2

		filter := moper.Query().Less("damage", i)

		if count, err := mocom.CountT[Hero](ctx, filter); err != nil {
			t.Error("[TestLess]:", err)
		} else if count != int64(num) {
			t.Error("[TestLess]", count, "!=", num)
		}
	}
}

func TestEqualLess(t *testing.T) {
	ctx := context.Background()

	for i := 0; i <= ROUND; i++ {
		num := i * (i + 1) / 2

		filter := moper.Query().EqualLess("damage", i)

		if count, err := mocom.CountT[Hero](ctx, filter); err != nil {
			t.Error("[TestEqualLess]:", err)
		} else if count != int64(num) {
			t.Error("[TestEqualLess]:", count, "!=", num)
		}
	}
}

func TestGreater(t *testing.T) {
	ctx := context.Background()

	for i := 0; i <= ROUND; i++ {
		num := TOTAL - i*(i+1)/2

		filter := moper.Query().Greater("damage", i)

		if count, err := mocom.CountT[Hero](ctx, filter); err != nil {
			t.Error("[TestGreater]:", err)
		} else if count != int64(num) {
			t.Error("[TestGreater]:", count, "!=", num)
		}
	}

}

func TestEqualGreater(t *testing.T) {
	ctx := context.Background()

	for i := 0; i <= ROUND; i++ {
		num := TOTAL - i*(i-1)/2

		filter := moper.Query().EqualGreater("damage", i)

		if count, err := mocom.CountT[Hero](ctx, filter); err != nil {
			t.Error("[TestEqualGreater]:", err)
		} else if count != int64(num) {
			t.Error("[TestEqualGreater]:", count, "!=", num)
		}
	}

}
