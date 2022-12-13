package moper

import "go.mongodb.org/mongo-driver/bson"

const (
	_inc   = "$inc"
	_set   = "$set"
	_push  = "$push"
	_unset = "$unset"
)

// SET
func (d *D) Set(pairs ...P) *D {
	*d = append(*d, bson.E{Key: _set, Value: toPair(pairs)})
	return d
}

func (d *D) SetD(pairs D) *D {
	*d = append(*d, bson.E{Key: _set, Value: pairs})
	return d
}

// UNSET
func (d *D) Unset(keys ...string) *D {
	res := D{}
	for i := range keys {
		res = append(res, bson.E{Key: keys[i], Value: ""})
	}
	*d = append(*d, bson.E{Key: _unset, Value: res})
	return d
}

// INC
func (d *D) Inc(pairs ...P) *D {
	pairLen := len(pairs)
	updated := D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].K, Value: pairs[i].V})
	}
	*d = append(*d, bson.E{Key: _inc, Value: updated})
	return d
}

func (d *D) IncD(pairs D) *D {
	*d = append(*d, bson.E{Key: _inc, Value: pairs})
	return d
}

// PUSH
func (d *D) Push(pairs ...P) *D {
	*d = append(*d, bson.E{Key: _push, Value: toPair(pairs)})
	return d
}

// PUSH
func (d *D) PushD(pairs D) *D {
	*d = append(*d, bson.E{Key: _push, Value: pairs})
	return d
}

func toPair(pairs []P) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].K, Value: pairs[i].V})
	}

	return updated
}
