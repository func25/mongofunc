package mongoquery

import "go.mongodb.org/mongo-driver/bson"

const (
	iNC   = "$inc"
	sET   = "$set"
	pUSH  = "$push"
	uNSET = "$unset"
)

// SET
func Set(pairs ...PairSetter) bson.E {
	return bson.E{Key: sET, Value: toPair(pairs)}
}

// UNSET
func Unset(keys ...string) bson.E {
	res := bson.D{}
	for i := range keys {
		res = append(res, bson.E{Key: keys[i], Value: ""})
	}
	return bson.E{Key: uNSET, Value: res}
}

// INC
func IncInt(pairs ...PairSetterInt) bson.E {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}
	return bson.E{Key: iNC, Value: updated}
}

func IncInt64(pairs ...PairSetterInt64) bson.E {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}
	return bson.E{Key: iNC, Value: updated}
}

// PUSH
func Push(pairs ...PairSetter) bson.E {
	return bson.E{Key: pUSH, Value: toPair(pairs)}
}

func toPair(pairs []PairSetter) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}

	return updated
}
