package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dihanto/go-mastermind/helper"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func MindMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		logger := logrus.New()
		logger.Infoln(request.Method)
		logger.Infoln(request.RequestURI)

		authHeader := request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := helper.ParseJWTString(tokenString)
		if err != nil || !token.Valid {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		}

		role, err := helper.GenerateRoleFromToken(token)
		if err != nil {
			log.Println(err)
		}

		log.Println(role)

		next(writer, request, params)
	}
}
