package db

import (
	"context"
	"mediadb/media"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConfig struct {
	Addr string
}

type MongoConnection struct {
	Config *MongoConfig
	client *mongo.Client
}

func NewMongoConnection(conf *MongoConfig) (*MongoConnection, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(conf.Addr))

	conn := MongoConnection{
		Config: conf,
		client: client,
	}

	return &conn, err
}

func (self *MongoConnection) Disconnect() {
}

func (self *MongoConnection) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := self.client.Ping(ctx, nil)
	return err
}

func (self *MongoConnection) CreateMovie(movie media.Movie) (any, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		InsertOne(ctx, movie)

	if err != nil {
		return nil, false, err
	}

	return res.InsertedID, res.Acknowledged, nil
}

func (self *MongoConnection) UpdateMovie(id any, movie media.Movie) (bool, any) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{bson.E{Key: "_id", Value: id}}

	res, err := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		ReplaceOne(ctx, filter, movie)

	if err != nil {
		return false, err
	}

	return res.Acknowledged, res.UpsertedID
}

func (self *MongoConnection) GetMovie(id any) (*media.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	movie := &media.Movie{}

	res := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		FindOne(ctx, filter)
	
	if err := res.Err(); err != nil {
		return nil, err
	}

	if err := res.Decode(movie); err != nil {
		return nil, err
	}

	return movie, nil
}
