package mongorely

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DefaultLocalDb() DbConfig {
	return DbConfig{
		DbName:   "defaultdb",
		UserName: "",
		Password: "",
		Host:     "localhost",
		Port:     "27017",
	}
}

var client *mongo.Client
var db *mongo.Database

func Connect(ctx context.Context, cfg DbConfig) (*mongo.Client, error) {
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.ToConnectionString()))
	db = client.Database(cfg.DbName)
	return client, err
}

func Setup(database *mongo.Database) {
	db = database
	client = db.Client()
}

func RetrieveClient() *mongo.Client {
	return client
}
