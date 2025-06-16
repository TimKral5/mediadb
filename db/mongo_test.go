package db_test

import (
	"mediadb/db"
	"mediadb/internals"
	"mediadb/media"
	"testing"
)


var initialized = false
var mongoConfig *db.MongoConfig

func getConfig(t *testing.T) {
	if initialized {
		return
	}

	env, err := internals.LoadEnvironment()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	mongoConfig = env.MongoConfig
}

func TestNewMongoConnection(t *testing.T) {
	getConfig(t)
	conn, err := db.NewMongoConnection(mongoConfig)

	if err != nil {
		t.Error(err)
		return
	}

	err = conn.Ping()

	if err != nil {
		t.Error(err)
		return
	}

	conn.Disconnect()
}

func TestCreateMovie(t *testing.T) {
	getConfig(t)
	conn, err := db.NewMongoConnection(mongoConfig)

	if err != nil {
		t.Error(err)
		return
	}

	err = conn.Ping()

	if err != nil {
		t.Error(err)
		return
	}

	succeeded, _, err := conn.CreateMovie(media.Movie{
		Title: "Test Movie",
		Description: "Test Description",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if !succeeded {
		t.Errorf("Create operation failed")
		return
	}

	conn.Disconnect()
}

func TestUpdateMovie(t *testing.T) {
	getConfig(t)
	conn, err := db.NewMongoConnection(mongoConfig)

	if err != nil {
		t.Error(err)
		return
	}

	err = conn.Ping()

	if err != nil {
		t.Error(err)
		return
	}

	succeeded, id, err := conn.CreateMovie(media.Movie{
		Title: "Test Movie",
		Description: "Test Description",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if !succeeded {
		t.Error("Create operation failed")
		return
	}

	succeeded, id, err = conn.UpdateMovie(id, media.Movie{
		Title: "Updated Test Movie",
		Description: "Updated Test Description",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if !succeeded {
		t.Error("Update operation failed")
		return
	}

	conn.Disconnect()
}

func TestGetMovie(t *testing.T) {
	getConfig(t)
	conn, err := db.NewMongoConnection(mongoConfig)

	if err != nil {
		t.Error(err)
		return
	}

	err = conn.Ping()
	
	if err != nil {
		t.Error(err)
		return
	}

	succeeded, id, err := conn.CreateMovie(media.Movie{
		Title: "Test Movie 2",
		Description: "Test Description 2",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if !succeeded {
		t.Errorf("Create operation failed")
		return
	}

	movie, err := conn.GetMovie(id)
	if err != nil {
		t.Error(err)
		return
	}

	if movie == nil {
		t.Errorf("Movie could not be fetched")
		return
	}

	if movie.Title != "Test Movie 2" {
		t.Errorf("Title does not match")
		return
	}

	conn.Disconnect()
}

