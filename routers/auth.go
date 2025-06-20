package routers

import (
	"encoding/json"
	"mediadb/auth"
	"mediadb/media"
	"mediadb/utils"
	"net/http"
)

type AuthRouter struct {
	SessionUUID string
	Ldap *auth.LDAPConnection
	Log   utils.Logger
}

func (self *AuthRouter) generateToken(w http.ResponseWriter, r *http.Request) {
	creds := auth.Credentials{}
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	valid, err := self.Ldap.ValidateLogin(creds)

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
	}

	ctx := auth.JwtContext{
		Username: creds.Username,
	}

	token, err := auth.GenerateJWT(self.SessionUUID, ctx)

	if err != nil {
		self.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenWrapper := media.Token{
		Token: token,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tokenWrapper)
}

func (self *AuthRouter) GetHandler() http.Handler {
	ctx := http.NewServeMux()
	ctx.HandleFunc("POST /jwt", self.generateToken)
	//ctx.HandleFunc("GET " + self.BaseRoute, )
	return ctx
}
