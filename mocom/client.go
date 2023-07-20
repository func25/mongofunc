package mocom

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// mocom only support one client, one database for now
var client *mongo.Client // default client
var db Database          // default database

type Database struct {
	*mongo.Database
}

// Connect mongodb://localhost:27017/?w=majority&retryWrites=false
func Connect(ctx context.Context, uri string, dbName string) error {
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	db.Database = client.Database(dbName)

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	return err
}

func SetDatabase(database *mongo.Database) {
	db.Database = database
	client = db.Client()
}

func GetClient() *mongo.Client {
	return client
}

func CollRead(collName string) *mongo.Collection {
	return db.Collection(collName, options.Collection().SetReadPreference(readpref.Nearest()))
}

func CollWrite(collName string) *mongo.Collection {
	return db.Collection(collName, options.Collection().SetReadPreference(readpref.Primary()))
}
