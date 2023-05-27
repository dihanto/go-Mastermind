package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type CartMiddleware struct {
	Handler http.Handler
}

func NewCartMiddleware(handler http.Handler) *CartMiddleware {
	return &CartMiddleware{
		Handler: handler,
	}
}

func (middleware *CartMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	logger := logrus.New()
	logger.Infoln(request.Method)
	logger.Infoln(request.RequestURI)

	middleware.Handler.ServeHTTP(writer, request)

}
