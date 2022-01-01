package mongoquery

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Init(e ...bson.E) bson.D {
	return e
}

func InEllipsis(fieldName string, value ...interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$in": value}}
}

func InArray(fieldName string, value interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$in": value}}
}

func NotInEllipsis(fieldName string, value ...interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$nin": value}}
}

func NotInArray(fieldName string, value interface{}) bson.E {
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
func EqualGreater(fieldName string, value interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$gte": value}}
}

func EqualGreaterTime(fieldName string, value time.Time) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$gte": primitive.NewDateTimeFromTime(value)}}
}

func Greater(fieldName string, value interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$gt": value}}
}

func NotEqual(fieldName string, value interface{}) bson.E {
	return bson.E{Key: fieldName, Value: bson.M{"$ne": value}}
}

func Equal(fieldName string, value interface{}) bson.E {
	return bson.E{Key: fieldName, Value: value}
}
