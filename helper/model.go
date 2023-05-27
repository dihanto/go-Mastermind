package helper

import (
	"github.com/dihanto/go-mastermind/model/domain"
	"github.com/dihanto/go-mastermind/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:         product.Id,
		Name:       product.Name,
		Price:      product.Price,
		Quantity:   product.Quantity,
		CategoryId: product.CategoryId,
	}
}

func ToCartResponse(cart domain.Cart) web.CartResponse {
	return web.CartResponse{
		Id:        cart.Id,
		ProductId: cart.ProductId,
		Quantity:  cart.Quantity,
	}
}
func ToCartResponses(carts []domain.Cart) []web.CartResponse {
	var cartResponse []web.CartResponse
	for _, cart := range carts {
		cartResponse = append(cartResponse, ToCartResponse(cart))
	}
	return cartResponse
}
