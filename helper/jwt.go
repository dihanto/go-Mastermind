package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	JWTSecret     = "mastermind"
	TokenDuration = time.Hour * 5
)

func GenerateJWTToken(id uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(TokenDuration).Unix()
	claims["role"] = "customer"

	return token.SignedString([]byte(JWTSecret))
}

func GenerateIdFromToken(tokenString string) (idString string, err error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("mastermind"), nil
	})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if idString, ok = claims["id"].(string); !ok {
			return
		}
		return idString, nil
	}

	return
}
