package moper

import "go.mongodb.org/mongo-driver/bson"

func (d *D) Or(filters ...D) *D {
	*d = append(*d, bson.E{Key: "$or", Value: filters})
	return d
}
