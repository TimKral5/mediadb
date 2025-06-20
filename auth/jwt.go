package auth

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JwtContext struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func GenerateJWT(secret string, ctx JwtContext) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ctx)

	jwtStr, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return jwtStr, nil
}

func ParseJWT(secret string, jwtStr string) (bool, *JwtContext, error) {
	token, err := jwt.ParseWithClaims(jwtStr, &JwtContext{}, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return false, nil, err
	}

	claims, ok := token.Claims.(*JwtContext)

	if !ok {
		return false, nil, nil
	}

	return token.Valid, claims, nil
}

// Remove prefix "Bearer " and parse JWT
func ParseJWTFromHeader(secret string, header string) (bool, *JwtContext, error) {
	if !strings.Contains(header, "Bearer ") {
		return false, nil, nil
	}

	jwtStr := strings.Split(header, "Bearer ")[1]
	return ParseJWT(secret, jwtStr)
}

