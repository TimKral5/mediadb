package main

import (
	"mediadb/utils"
	"net/http"
)

func main() {
	log := utils.NewLogger()
	ctx := http.NewServeMux()
	log.Info("Launching MediaDB v0.1.0-alpha...")

	stack := utils.CreateStack(
		log.Middleware,
	)

	ctx.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Content-Type", "application/json")

		w.WriteHeader(500)
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	server := http.Server{
		Addr: ":3000",
		Handler: stack(ctx),
	}

	log.Info("Running HTTP server...")
	server.ListenAndServe()
}
