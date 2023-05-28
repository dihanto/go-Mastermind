package service

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-mastermind/exception"
	"github.com/dihanto/go-mastermind/helper"
	"github.com/dihanto/go-mastermind/model/domain"
	"github.com/dihanto/go-mastermind/model/web"
	"github.com/dihanto/go-mastermind/repository"
	"github.com/go-playground/validator/v10"
)

type CartServiceImpl struct {
	CartRepository repository.CartRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewCartService(cartRepository repository.CartRepository, db *sql.DB, validate *validator.Validate) CartService {
	return &CartServiceImpl{
		CartRepository: cartRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *CartServiceImpl) FindProductById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.CartRepository.FindProductById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}
func (service *CartServiceImpl) AddToCart(ctx context.Context, request web.CartAddRequest) web.CartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := domain.CartItem{
		CartId:    request.CartId,
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

func (service *CartServiceImpl) UpdateCart(ctx context.Context, request web.CartUpdateRequest) web.CartUpdateResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// cartItem := domain.CartItem{
	// 	CartItemId: request.CartItemId,
	// 	CartId:     request.CartId,
	// 	ProductId:  request.ProductId,
	// 	Quantity:   request.Quantity,
	// }

	cartItem, err := service.CartRepository.UpdateCart(ctx, tx, domain.CartItem(request))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	cartResponse := web.CartUpdateResponse{
		CartId:    cartItem.CartId,
		ProductId: cartItem.ProductId,
		Quantity:  cartItem.Quantity,
	}

	return cartResponse
}

func (service *CartServiceImpl) DeleteCart(ctx context.Context, cartItemId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.CartRepository.DeleteCart(ctx, tx, cartItemId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

}
