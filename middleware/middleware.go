package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type AuthMiddleware struct {
	Next   httprouter.Handle
	JWTKey []byte
}

func NewAuthMiddleware(next httprouter.Handle, jwtKey []byte) *AuthMiddleware {
	return &AuthMiddleware{
		Next:   next,
		JWTKey: jwtKey,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logger := logrus.New()
	logger.Infoln(request.Method)
	logger.Infoln(request.RequestURI)

	authHeader := request.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return middleware.JWTKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	middleware.Next(writer, request, params)
}
