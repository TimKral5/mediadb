package main

import (
	_ "mediadb/internals"
	"mediadb/routers"
	"net/http"
)

func main() {
	log := NewLogger()
	auth := NewAuthenticator()

	stack := CreateStack(
		log.Middleware,
		auth.Middleware,
	)

	helloRouter := routers.GetHelloRouter()

	handler := helloRouter

	server := http.Server{
		Addr: ":3000",
		Handler: stack(handler),
	}


	log.Info("Listening on port 3000...")
	server.ListenAndServe()
}
