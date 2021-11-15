package mongoquery

import "go.mongodb.org/mongo-driver/bson"

func Or(filters ...interface{}) bson.E {
	return bson.E{Key: "$or", Value: filters}
}
