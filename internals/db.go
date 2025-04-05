package mediadb

import (
	"context"
	"log"

	//"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func Connect(uri string) *Database {
	log.Println("Connecting to database...")

	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(opts)

	if err != nil {
		log.Fatalln(err)
	}

	db := &Database{
		Client: client,
	}

	return db
}

func (db *Database) Disconnect() {
	db.Client.Disconnect(context.TODO())
}
