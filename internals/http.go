package internals

import (
	"mediadb/utils"
	"net/http"
)

func LaunchHttpServer(log utils.Logger, ctx *http.ServeMux) {
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

	server.ListenAndServe()
}
