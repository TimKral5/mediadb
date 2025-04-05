package main

import (
	mediadb "mediadb/internals"
	"net/http"
)

func main() {
	log := mediadb.NewLogger()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	server := http.Server{
		Addr: ":3000",
		Handler: EnableLogging(log, handler),
	}

	log.Info("Listening on port 3000...")
	server.ListenAndServe()
}
