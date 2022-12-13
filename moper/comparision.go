package moper

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InEll is InEllipsis
func (d *D) InEll(fieldName string, value ...interface{}) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$in": value}})
	return d
}

func (d *D) InArray(fieldName string, value interface{}) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$in": value}})
	return d
}

// NotInEll is NotInEllipsis
func (d *D) NotInEll(fieldName string, value ...interface{}) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$nin": value}})
	return d
}

func (d *D) NotInArray(fieldName string, value interface{}) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$nin": value}})
	return d
}

// LESS OR EQUAL
func (d *D) EqualLess(fieldName string, value interface{}) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$lte": value}})
	return d
}

func (d *D) EqualLessTime(fieldName string, value time.Time) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$lte": primitive.NewDateTimeFromTime(value)}})
	return d
}

func (d *D) Less(fieldName string, value interface{}) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$lt": value}})
	return d
}

func (d *D) LessTime(fieldName string, value time.Time) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$lt": primitive.NewDateTimeFromTime(value)}})
	return d
}

// GREATER OR EQUAL
func (d *D) EqualGreater(fieldName string, value interface{}) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$gte": value}})
	return d
}

func (d *D) EqualGreaterTime(fieldName string, value time.Time) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$gte": primitive.NewDateTimeFromTime(value)}})
	return d
}

func (d *D) Greater(fieldName string, value interface{}) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$gt": value}})
	return d
}

func (d *D) GreaterTime(fieldName string, value time.Time) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$gt": primitive.NewDateTimeFromTime(value)}})
	return d
}

// EQUAL

func (d *D) NotEqual(fieldName string, value interface{}) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: bson.M{"$ne": value}})
	return d
}

func (d *D) Equal(fieldName string, value interface{}) *D {
	*d = append(*d, bson.E{Key: fieldName, Value: value})
	return d
}
