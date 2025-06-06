package db_test

import (
	"context"
	"mediadb/db"
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

