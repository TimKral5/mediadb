package main

import (
	"mediadb/internals"
	"mediadb/middleware"
	"mediadb/routers"
	"net/http"
	"os"
)

type Env struct {
	mongoUrl string
	lokiUrl  string
	port     string
}

type Db struct {
	mongodb internals.MongoDB
}

func getenv() Env {
	return Env{
		mongoUrl: os.Getenv("MEDIADB_MONGODB_URL"),
		lokiUrl:  os.Getenv("MEDIADB_LOKI_URL"),
		port:     os.Getenv("MEDIADB_PORT"),
	}
}

func setupDbConnections(env Env) Db {
	var err error

	mongodb, err := internals.Connect(env.mongoUrl)
	if err != nil {
		panic(err)
	}

	db := Db{
		mongodb: mongodb,
	}

	return db
}

func main() {
	var err error

	log := middleware.NewLogger()
	auth := middleware.NewAuthenticator()

	env := getenv()
	_ = setupDbConnections(env)

	stack := middleware.CreateStack(
		log.Middleware,
		auth.Middleware,
	)

	helloRouter := routers.GetHelloRouter()

	handler := helloRouter

	server := http.Server{
		Addr:    ":3000",
		Handler: stack(handler),
	}

	log.Info("Listening on port 3000...")
	err = server.ListenAndServe()
	if err != nil {
		log.Error(err)
	}
}
