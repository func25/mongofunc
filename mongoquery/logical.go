package mongoquery

import "go.mongodb.org/mongo-driver/bson"

func Or(filters ...bson.D) bson.E {
	return bson.E{Key: "$or", Value: filters}
}
