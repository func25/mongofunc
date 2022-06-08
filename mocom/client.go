package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database

//Connect mongodb://localhost:27017/?w=majority&retryWrites=false
func Connect(ctx context.Context, uri string, dbName string) (*mongo.Client, error) {
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	db = client.Database(dbName)
	return client, err
}

func Setup(database *mongo.Database) {
	db = database
	client = db.Client()
}

func RetrieveClient() *mongo.Client {
	return client
}
