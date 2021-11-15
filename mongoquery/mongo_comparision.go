package mongoquery

import "go.mongodb.org/mongo-driver/bson"

func Init(e ...bson.E) bson.D {
	return e
}

func In(fieldName string, value ...interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$in": value}}
}

func NotIn(fieldName string, value ...interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$nin": value}}
}

// LESS OR EQUAL
func EqualLess(fieldName string, value interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$lte": value}}
}

func Less(fieldName string, value interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$lt": value}}
}

// GREATER OR EQUAL
func EqualGreaterInt(fieldName string, value int) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$gte": value}}
}

func EqualGreaterInt64(fieldName string, value int64) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$gte": value}}
}

func GreaterInt(fieldName string, value int) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$gt": value}}
}

func GreaterInt64(fieldName string, value int64) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$gt": value}}
}

func NotEqual(fieldName string, value interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$ne": value}}
}

func Equal(fieldName string, value interface{}) bson.E {
	return bson.E{Key: fieldName, Value: value}
}
