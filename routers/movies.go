package routers

import (
	"encoding/json"
	"mediadb/db"
	"mediadb/media"
	"mediadb/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MovieRouter struct {
	BaseRoute string
	Mongo *db.MongoConnection
	Log   utils.Logger
}

func (self *MovieRouter) createMovie(w http.ResponseWriter, r *http.Request) {
	movie := media.Movie{}

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	succeeded, _, err := self.Mongo.CreateMovie(movie)

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !succeeded {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Write([]byte("Hello"))
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

	if err == mongo.ErrNoDocuments {
		w.WriteHeader(204)
		return
	}

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
	ctx.HandleFunc("POST " + self.BaseRoute, self.createMovie)
	ctx.HandleFunc("GET " + self.BaseRoute, self.getMovie)
	return ctx
}
