package routers

import (
	"encoding/json"
	"mediadb/auth"
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

	w.Write([]byte(token))
}

func (self *AuthRouter) GetHandler() http.Handler {
	ctx := http.NewServeMux()
	ctx.HandleFunc("POST /", self.generateToken)
	//ctx.HandleFunc("GET " + self.BaseRoute, )
	return ctx
}
