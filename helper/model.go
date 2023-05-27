package helper

import (
	"github.com/dihanto/go-mastermind/model/domain"
	"github.com/dihanto/go-mastermind/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:         product.ProductId,
		Name:       product.Name,
		Price:      product.Price,
		CategoryId: product.CategoryId,
	}
}

func ToCartResponse(cartItem domain.CartItem) web.CartResponse {
	return web.CartResponse{
		CartItemId: cartItem.CartItemId,
		CartId:     cartItem.CartId,
		ProductId:  cartItem.ProductId,
		Quantity:   cartItem.Quantity,
	}
}
func ToCartResponses(cartItems []domain.CartItem) []web.CartResponse {
	var cartResponse []web.CartResponse
	for _, cart := range cartItems {
		cartResponse = append(cartResponse, ToCartResponse(cart))
	}
	return cartResponse
}
