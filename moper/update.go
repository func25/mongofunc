package moper

import "go.mongodb.org/mongo-driver/bson"

const (
	_inc   = "$inc"
	_set   = "$set"
	_push  = "$push"
	_unset = "$unset"
)

// SET
func (d D) Set(pairs ...P) D {
	return append(d, bson.E{Key: _set, Value: toPair(pairs)})
}

func (d D) SetD(pairs D) D {
	return append(d, bson.E{Key: _set, Value: pairs})
}

// UNSET
func (d D) Unset(keys ...string) D {
	res := D{}
	for i := range keys {
		res = append(res, bson.E{Key: keys[i], Value: ""})
	}
	return append(d, bson.E{Key: _unset, Value: res})
}

// INC
func (d D) Inc(pairs ...P) D {
	pairLen := len(pairs)
	updated := D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].K, Value: pairs[i].V})
	}
	return append(d, bson.E{Key: _inc, Value: updated})
}

func (d D) IncD(pairs D) D {
	return append(d, bson.E{Key: _inc, Value: pairs})
}

// PUSH
func (d D) Push(pairs ...P) D {
	return append(d, bson.E{Key: _push, Value: toPair(pairs)})
}

// PUSH
func (d D) PushD(pairs D) D {
	return append(d, bson.E{Key: _push, Value: pairs})
}

func toPair(pairs []P) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].K, Value: pairs[i].V})
	}

	return updated
}
