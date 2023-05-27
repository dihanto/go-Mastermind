package app

import (
	"github.com/dihanto/go-mastermind/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.CartController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/products/:productId", controller.FindProductById)
	router.POST("/api/carts", controller.AddToCart)
	router.GET("/api/carts", controller.GetCart)
	router.PATCH("/api/carts/:cartItemId", controller.UpdateCart)

	return router
}
