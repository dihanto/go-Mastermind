package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dihanto/go-mastermind/helper"
	"github.com/dihanto/go-mastermind/model/web"
	"github.com/dihanto/go-mastermind/service"
	"github.com/julienschmidt/httprouter"
)

type CartControllerImpl struct {
	CartService service.CartService
}

func NewCartContnroller(cartService service.CartService) CartController {
	return &CartControllerImpl{
		CartService: cartService,
	}
}

func (controller *CartControllerImpl) FindProductById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	productId := param.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := controller.CartService.FindProductById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	helper.PanicIfError(err)

}
func (controller *CartControllerImpl) AddToCart(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	cartRequest := web.CartAddRequest{}
	err := json.NewDecoder(request.Body).Decode(&cartRequest)
	helper.PanicIfError(err)

	cartResponse := controller.CartService.AddToCart(request.Context(), cartRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CartControllerImpl) GetCart(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	cartResponse := controller.CartService.GetCart(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}
	log.Println(cartResponse)

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(webResponse)
	helper.PanicIfError(err)
}
