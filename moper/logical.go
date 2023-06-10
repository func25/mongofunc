package moper

import "go.mongodb.org/mongo-driver/bson"

func (d D) Or(filters ...D) D {
	return append(d, bson.E{Key: "$or", Value: filters})
}
