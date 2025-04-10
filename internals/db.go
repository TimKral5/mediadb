package internals

import (
	"context"
	//"mediadb/middleware"

	//"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

func Connect(uri string) (MongoDB, error) {
	db := MongoDB{}

	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(opts)

	db.Client = client
	return db, err
}

func (db *MongoDB) Disconnect() {
	db.Client.Disconnect(context.TODO())
}
