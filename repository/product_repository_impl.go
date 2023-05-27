package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dihanto/go-mastermind/helper"
	"github.com/dihanto/go-mastermind/model/domain"
)

type CartRepositoryImpl struct {
}

func NewCartRepository() CartRepository {
	return &CartRepositoryImpl{}
}

func (repository *CartRepositoryImpl) FindProductById(ctx context.Context, tx *sql.Tx, productId int) (product domain.Product, err error) {
	script := "Select name, price, quantity, category_id from products where id = ?"
	rows, err := tx.QueryContext(ctx, script, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&product.Name, &product.Price, &product.Quantity, &product.CategoryId)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}

}

func (repository *CartRepositoryImpl) AddToCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	script := "insert into carts (product_id, quantity) values (?,?)"
	result, err := tx.ExecContext(ctx, script, cart.ProductId, cart.Quantity)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	cart.Id = int(id)
	return cart
}

func (repository *CartRepositoryImpl) GetCart(ctx context.Context, tx *sql.Tx) (cart []domain.Cart, err error) {
	script := "select id, product_id, quantity from carts"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	carts := []domain.Cart{}
	for rows.Next() {
		cart := domain.Cart{}
		err := rows.Scan(&cart.Id, &cart.ProductId, &cart.Quantity)
		helper.PanicIfError(err)
		carts = append(carts, cart)
	}
	return carts, err

}

func (repository *CartRepositoryImpl) UpdateCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) {
	script := "update carts set product_id=?,quantity=? where id=?"
	_, err := tx.ExecContext(ctx, script, cart.ProductId, cart.Quantity, cart.Id)
	helper.PanicIfError(err)
}
