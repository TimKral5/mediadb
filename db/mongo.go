package db

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConnection struct {
	Addr string
	client *mongo.Client
	ctx context.Context
}

func NewMongoConnection(addr string, ctx context.Context) (*MongoConnection, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(addr))

	conn := MongoConnection{
		Addr: addr,
		client: client,
		ctx: ctx,
	}

	return &conn, err
}

func (self *MongoConnection) Ping() error {
	err := self.client.Ping(self.ctx, nil)
	return err
}

