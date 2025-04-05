package main

import (
	"log"
	"net/http"
)

func main() {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	server := http.Server{
		Addr: ":3000",
		Handler: handler,
	}

	log.Println("Listening on port 3000")
	server.ListenAndServe()
}
