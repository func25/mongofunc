package mongoquery

import "go.mongodb.org/mongo-driver/bson"

const (
	iNC  = "$inc"
	sET  = "$set"
	pUSH = "$push"
)

// SET
func Set(pairs ...PairSetter) bson.D {
	return bson.D{{Key: sET, Value: Pair(pairs)}}
}

// INC
func IncInt(pairs ...PairSetterInt) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}
	return bson.D{{Key: iNC, Value: updated}}
}

func IncInt64(pairs ...PairSetterInt64) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}
	return bson.D{{Key: iNC, Value: updated}}
}

// PUSH
func Push(pairs ...PairSetter) bson.D {
	return bson.D{{Key: pUSH, Value: Pair(pairs)}}
}

func Pair(pairs []PairSetter) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}

	return updated
}
