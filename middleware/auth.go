package middleware

import (
	"net/http"
)

//
// Authenticator
//

type Authenticator struct {
}

func NewAuthenticator() Authenticator {
	return Authenticator{}
}

func (auth *Authenticator) ValidateToken(token string) bool {
	if token != "1234" {
		return false
	}
	return true
}

//
// Authenticator Middleware
//

func (auth *Authenticator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Api-Key")
		if auth.ValidateToken(token) {
			next.ServeHTTP(w, r)
                        return
		}
		w.WriteHeader(401)
	})
}
