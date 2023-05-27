package service

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-mastermind/helper"
	"github.com/dihanto/go-mastermind/model/domain"
	"github.com/dihanto/go-mastermind/model/web"
	"github.com/dihanto/go-mastermind/repository"
)

type CartServiceImpl struct {
	CartRepository repository.CartRepository
	DB             *sql.DB
}

func NewCartService(cartRepository repository.CartRepository, db *sql.DB) CartService {
	return &CartServiceImpl{
		CartRepository: cartRepository,
		DB:             db,
	}
}

func (service *CartServiceImpl) FindProductById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.CartRepository.FindProductById(ctx, tx, productId)
	helper.PanicIfError(err)

	return helper.ToProductResponse(product)
}
func (service *CartServiceImpl) AddToCart(ctx context.Context, request web.CartAddOrUpdateRequest) web.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := domain.Cart{
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
	}

	cart = service.CartRepository.AddToCart(ctx, tx, cart)

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImpl) GetCart(ctx context.Context) []web.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	carts, err := service.CartRepository.GetCart(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToCartResponses(carts)

}

func (service *CartServiceImpl) UpdateCart(ctx context.Context, request web.CartUpdateRequest) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := domain.Cart{
		Id:        request.Id,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
	}

	service.CartRepository.UpdateCart(ctx, tx, cart)
}
