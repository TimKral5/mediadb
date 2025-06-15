package routers

import (
	"encoding/json"
	"mediadb/db"
	"mediadb/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MovieRouter struct {
	BaseRoute string
	Mongo *db.MongoConnection
	Log   utils.Logger
}

func (self *MovieRouter) createMovie(w http.ResponseWriter, r *http.Request) {

}

func (self *MovieRouter) getMovie(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if !query.Has("id") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := bson.ObjectIDFromHex(query.Get("id"))

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	movie, err := self.Mongo.GetMovie(id)

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(movie)
	w.Write(json)
}

func (self *MovieRouter) updateMovie(w http.ResponseWriter, r *http.Request) {

}

func (self *MovieRouter) deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func (self *MovieRouter) Middleware(next http.Handler) http.Handler {
	ctx := http.NewServeMux()
	ctx.HandleFunc("GET " + self.BaseRoute, self.getMovie)
	return ctx
}
