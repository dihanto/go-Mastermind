package service

import (
	"context"

	"github.com/dihanto/go-mastermind/model/web"
)

type CartService interface {
	FindProductById(ctx context.Context, productId int) web.ProductResponse
	AddToCart(ctx context.Context, request web.CartAddRequest) web.CartResponse
	GetCart(ctx context.Context) []web.CartResponse
	UpdateCart(ctx context.Context, request web.CartUpdateRequest) web.CartUpdateResponse
	DeleteCart(ctx context.Context, cartItemId int)
}
