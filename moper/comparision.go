package moper

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InEll - InEllipsis
func (d D) InEll(fieldName string, value ...interface{}) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$in": value}})
}

func (d D) InArray(fieldName string, value interface{}) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$in": value}})
}

// NotInEll is NotInEllipsis
func (d D) NotInEll(fieldName string, value ...interface{}) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$nin": value}})
}

func (d D) NotInArray(fieldName string, value interface{}) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$nin": value}})
}

// LESS OR EQUAL
func (d D) EqualLess(fieldName string, value interface{}) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$lte": value}})
}

func (d D) EqualLessTime(fieldName string, value time.Time) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$lte": primitive.NewDateTimeFromTime(value)}})
}

func (d D) Less(fieldName string, value interface{}) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$lt": value}})
}

func (d D) LessTime(fieldName string, value time.Time) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$lt": primitive.NewDateTimeFromTime(value)}})
}

// GREATER OR EQUAL
func (d D) EqualGreater(fieldName string, value interface{}) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$gte": value}})
}

func (d D) EqualGreaterTime(fieldName string, value time.Time) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$gte": primitive.NewDateTimeFromTime(value)}})
}

func (d D) Greater(fieldName string, value interface{}) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$gt": value}})
}

func (d D) GreaterTime(fieldName string, value time.Time) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$gt": primitive.NewDateTimeFromTime(value)}})
}

// EQUAL
func (d D) NotEqual(fieldName string, value interface{}) D {
	return append(d, bson.E{Key: fieldName, Value: bson.M{"$ne": value}})
}

func (d D) Equal(fieldName string, value interface{}) D {
	return append(d, bson.E{Key: fieldName, Value: value})
}
