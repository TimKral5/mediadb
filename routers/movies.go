package routers

import (
	"encoding/json"
	"mediadb/auth"
	"mediadb/db"
	"mediadb/media"
	"mediadb/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MovieRouter struct {
	BaseRoute string
	SessionUUID string
	Mongo *db.MongoConnection
	Ldap *auth.LDAPConnection
	Log   utils.Logger
}

func (self *MovieRouter) createMovie(w http.ResponseWriter, r *http.Request) {
	status := r.Header.Get("X-JWT-Status")
	username := r.Header.Get("X-JWT-Username")

	if status != "valid" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	isMember, err := self.Ldap.IsUserGroupMember(username, "mediadb_create_movie")

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !isMember {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	movie := media.Movie{}

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	succeeded, id, err := self.Mongo.CreateMovie(movie)

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !succeeded {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id.Hex())
}

func (self *MovieRouter) getMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	status := r.Header.Get("X-JWT-Status")
	username := r.Header.Get("X-JWT-Username")

	if status != "valid" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	isMember, err := self.Ldap.IsUserGroupMember(username, "mediadb_get_movie")

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !isMember {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	id, err := bson.ObjectIDFromHex(idStr)

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
	status := r.Header.Get("X-JWT-Status")
	username := r.Header.Get("X-JWT-Username")
	idStr := r.PathValue("id")

	if status != "valid" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	isMember, err := self.Ldap.IsUserGroupMember(username, "mediadb_update_movie")

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !isMember {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	movie := media.Movie{}

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := bson.ObjectIDFromHex(idStr)

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	succeeded, id, err := self.Mongo.UpdateMovie(id, movie)

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !succeeded {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (self *MovieRouter) deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func (self *MovieRouter) GetHandler() http.Handler {
	ctx := http.NewServeMux()
	ctx.HandleFunc("POST " + utils.ConcatUrls(self.BaseRoute, "/", true), self.createMovie)
	ctx.HandleFunc("GET " + utils.ConcatUrls(self.BaseRoute, "/{id}", false), self.getMovie)
	ctx.HandleFunc("PUT " + utils.ConcatUrls(self.BaseRoute, "/{id}", false), self.updateMovie)
	return ctx
}
