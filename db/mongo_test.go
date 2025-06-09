package db_test

import (
	"context"
	"mediadb/db"
	"mediadb/internals"
	"testing"
	"time"
)

func TestNewMongoConnection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := db.NewMongoConnection("mongodb://root:root@127.0.0.1/", ctx)

	if err != nil {
		t.Error(err)
		return
	}

	err = conn.Ping()

	if err != nil {
		t.Error(err)
		return
	}
}

func TestCreateMovie(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, _ := db.NewMongoConnection("mongodb://root:root@127.0.0.1/", ctx)
	_ = conn.Ping()

	succeeded, id := conn.CreateMovie(internals.Movie{
		Title: "Test Movie",
		Description: "Test Description",
	})

	if !succeeded {
		t.Errorf("Create operation failed")
		t.Error(id)
		return
	}
}

func TestUpdateMovie(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, _ := db.NewMongoConnection("mongodb://root:root@127.0.0.1/", ctx)
	_ = conn.Ping()

	succeeded, id := conn.CreateMovie(internals.Movie{
		Title: "Test Movie 2",
		Description: "Test Description 2",
	})

	if !succeeded {
		t.Errorf("Create operation failed")
		t.Error(id)
		return
	}

	succeeded, id = conn.UpdateMovie(id, internals.Movie{
		Title: "Updated Test Movie",
		Description: "Updated Test Description",
	})

	if !succeeded {
		t.Errorf("Update operation failed")
		t.Error(id)
		return
	}
}

