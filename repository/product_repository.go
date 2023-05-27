package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-mastermind/model/domain"
)

type CartRepository interface {
	FindProductById(ctx context.Context, tx *sql.Tx, productId int) (product domain.Product, err error)
	AddToCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	GetCart(ctx context.Context, tx *sql.Tx) (cart []domain.Cart, err error)
	UpdateCart(ctx context.Context, tx *sql.Tx, cart domain.Cart)
}
