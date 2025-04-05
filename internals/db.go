package mediadb

import (
	"log"

	//"go.mongodb.org/mongo-driver/v2/bson"
	//"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Database struct {
	Client *mongo.Client
}

func Connect() *Database {
	log.Println("Connecting to database...")

	client, err := mongo.Connect()

	if err != nil {
		log.Fatalln(err)
	}

	db := &Database{
		Client: client,
	}

	return db
}
