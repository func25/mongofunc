package moper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type D primitive.D

// P is pair
type P struct {
	K string
	V interface{}
}

func Pr(k string, v interface{}) P {
	return P{k, v}
}

func Query() D {
	return D{}
}

func (d D) MarshalBSON() ([]byte, error) {
	return bson.Marshal(primitive.D(d))
}
