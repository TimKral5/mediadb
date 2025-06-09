package db

import (
	"context"
	"mediadb/internals"
	"structs"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConnection struct {
	Addr   string
	client *mongo.Client
	ctx    context.Context
}

func NewMongoConnection(addr string, ctx context.Context) (*MongoConnection, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(addr))

	conn := MongoConnection{
		Addr:   addr,
		client: client,
		ctx:    ctx,
	}

	return &conn, err
}

func (self *MongoConnection) Ping() error {
	err := self.client.Ping(self.ctx, nil)
	return err
}

func (self *MongoConnection) CreateMovie(movie internals.Movie) (bool, any) {
	res, _ := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		InsertOne(self.ctx, movie)
	return res.Acknowledged, res.InsertedID
}

func (self *MongoConnection) UpdateMovie(id any, movie internals.Movie) (bool, any) {
	filter := bson.D{{"_id", id}}

	res, err := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		ReplaceOne(self.ctx, filter, movie)

	if err != nil {
		return false, err
	}

	return res.Acknowledged, res.UpsertedID
}

