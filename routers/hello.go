package routers

import "net/http"

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func GetHelloRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /hello", getHello)

	return router
}
