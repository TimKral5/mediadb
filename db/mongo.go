package db

import (
	"context"
	"mediadb/media"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConfig struct {
	Addr string
	Context context.Context
	CancelContext context.CancelFunc
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
	self.Config.CancelContext()
}

func (self *MongoConnection) Ping() error {
	err := self.client.Ping(self.Config.Context, nil)
	return err
}

func (self *MongoConnection) CreateMovie(movie media.Movie) (bool, any) {
	res, _ := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		InsertOne(self.Config.Context, movie)
	return res.Acknowledged, res.InsertedID
}

func (self *MongoConnection) UpdateMovie(id any, movie media.Movie) (bool, any) {
	filter := bson.D{bson.E{Key: "_id", Value: id}}

	res, err := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		ReplaceOne(self.Config.Context, filter, movie)

	if err != nil {
		return false, err
	}

	return res.Acknowledged, res.UpsertedID
}

func (self *MongoConnection) GetMovie(id any) (*media.Movie, error) {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	movie := &media.Movie{}

	res := self.client.
		Database("mediadb").
		Collection("mediadb_movies").
		FindOne(self.Config.Context, filter)

	if res == nil {
		return nil, nil
	}

	err := res.Decode(movie)

	if err != nil {
		return nil, err
	}

	return movie, nil
}
