package db

import (
	"context"
	"mediadb/internals"

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

func (self *MongoConnection) CreateMovie(movie internals.Movie) any {
	res, _ := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		InsertOne(self.ctx, movie)
	return res.InsertedID
}
