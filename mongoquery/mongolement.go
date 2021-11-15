package mongoquery

import "go.mongodb.org/mongo-driver/bson"

func ExistOrDefault(fieldName string, exist bool, value interface{}) bson.E {
	return Or(Equal(fieldName, value), Exist(fieldName, exist))
}

func Exist(fieldName string, exist bool) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$exist": exist}}
}
