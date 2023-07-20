package moper

import "go.mongodb.org/mongo-driver/bson"

func (d D) ExistOrDefault(fieldName string, exist bool, value interface{}) D {
	return d.Or(d.Equal(fieldName, value), d.Exists(fieldName, exist))
}

func (d D) Exists(fieldName string, exist bool) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$exists": exist}})
}
