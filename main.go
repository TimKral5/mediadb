package main

import (
	"mediadb/internals"
	"mediadb/routers"
	"net/http"
)

func main() {
	log := internals.NewLogger()

	helloRouter := routers.GetHelloRouter()

	handler := helloRouter

	server := http.Server{
		Addr: ":3000",
		Handler: EnableLogging(log, handler),
	}


	log.Info("Listening on port 3000...")
	server.ListenAndServe()
}
