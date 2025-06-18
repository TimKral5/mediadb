package auth_test

import (
	"mediadb/auth"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestJwtToken(t *testing.T) {
	secret := uuid.New().String()

	token, err := auth.GenerateJWT(secret, auth.JwtContext{
		RegisteredClaims: jwt.RegisteredClaims{},
		Username: "demo",
	})

	if err != nil {
		t.Error(err)
		return
	}

	valid, claims, err := auth.ParseJWT(secret, token)

	if err != nil {
		t.Error(err)
		return
	}

	if !valid {
		t.Error("The token is invalid")
		return
	}

	if claims.Username != "demo" {
		t.Error("Usernames do not match")
		return
	}
}
