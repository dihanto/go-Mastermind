package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CartController interface {
	FindProductById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	AddToCart(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	GetCart(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	UpdateCart(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	DeleteCart(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
