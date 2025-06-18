package middleware

import (
	"mediadb/auth"
	"mediadb/utils"
	"net/http"
)

type Authenticator struct {
	SessionUUID string
	Log         utils.Logger
}

func (self *Authenticator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			r.Header.Set("X-JWT-Status", "null")
			next.ServeHTTP(w, r)
			return
		}

		valid, jwtToken, err := auth.ParseJWTFromHeader(self.SessionUUID, authHeader)

		switch {
		case err != nil:
			self.Log.Error(err)
			r.Header.Set("X-JWT-Status", "error")
		case !valid:
			r.Header.Set("X-JWT-Status", "invalid")
		default:
			r.Header.Set("X-JWT-Username", jwtToken.Username)
			r.Header.Set("X-JWT-Status", "valid")
		}

		next.ServeHTTP(w, r)
	})
}
