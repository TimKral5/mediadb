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

func (self *MongoConnection) CreateMovie(movie media.Movie) (bool, bson.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		InsertOne(ctx, movie)

	if err != nil {
		return false, bson.NewObjectID(), err
	}

	return res.Acknowledged, res.InsertedID.(bson.ObjectID), nil
}

func (self *MongoConnection) UpdateMovie(id bson.ObjectID, movie media.Movie) (bool, bson.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{bson.E{Key: "_id", Value: id}}

	res, err := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		ReplaceOne(ctx, filter, movie)

	if err != nil {
		return false, bson.NewObjectID(), err
	}

	if !res.Acknowledged {
		return false, bson.NewObjectID(), nil
	}

	if res.UpsertedID == nil {
		return true, id, nil
	}

	return true, res.UpsertedID.(bson.ObjectID), nil
}

func (self *MongoConnection) GetMovie(id bson.ObjectID) (*media.Movie, error) {
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
