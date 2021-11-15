package mongoquery

import "go.mongodb.org/mongo-driver/bson"

const (
	iNC  = "$inc"
	sET  = "$set"
	pUSH = "$push"
)

// INC
func DIncInt(pairs ...PairSetterInt) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}
	return bson.D{{Key: iNC, Value: updated}}
}

func DIncInt64(pairs ...PairSetterInt64) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}
	return bson.D{{Key: iNC, Value: updated}}
}

// SET
func DSet(pairs ...PairSetter) bson.D {
	return bson.D{{Key: sET, Value: dPair(pairs)}}
}

// PUSH
func DPush(pairs ...PairSetter) bson.D {
	return bson.D{{Key: pUSH, Value: dPair(pairs)}}
}

func dPair(pairs []PairSetter) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}

	return updated
}
