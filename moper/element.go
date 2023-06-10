package moper

import "go.mongodb.org/mongo-driver/bson"

// func ExistOrDefault(fieldName string, exist bool, value interface{}) bson.E {
// 	return Or(Equal(fieldName, value), Exist(fieldName, exist))
// }

func (d D) Exists(fieldName string, exist bool) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$exists": exist}})
}
