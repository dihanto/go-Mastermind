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
	script := "Select name, price, category_id from products where product_id = ?"
	rows, err := tx.QueryContext(ctx, script, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&product.Name, &product.Price, &product.CategoryId)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}

}

func (repository *CartRepositoryImpl) AddToCart(ctx context.Context, tx *sql.Tx, cartItem domain.CartItem) domain.CartItem {

	script := "insert into cart_items (cart_id, product_id, quantity) values (?,?,?)"
	result, err := tx.ExecContext(ctx, script, cartItem.CartId, cartItem.ProductId, cartItem.Quantity)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	cartItem.CartItemId = int(id)
	cartItemResult := cartItem
	return cartItemResult
}

func (repository *CartRepositoryImpl) GetCart(ctx context.Context, tx *sql.Tx) (cartItems []domain.CartItem, err error) {
	script := "select cart_item_id, cart_id, product_id, quantity from cart_items"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	cartItems = []domain.CartItem{}
	for rows.Next() {
		cartItem := domain.CartItem{}
		err := rows.Scan(&cartItem.CartItemId, &cartItem.CartId, &cartItem.ProductId, &cartItem.Quantity)
		helper.PanicIfError(err)
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, err

}

func (repository *CartRepositoryImpl) UpdateCart(ctx context.Context, tx *sql.Tx, cartItem domain.CartItem) (domain.CartItem, error) {
	checkScript := "select count(*) from cart_items where cart_item_id=?"
	var count int
	err := tx.QueryRowContext(ctx, checkScript, cartItem.CartItemId).Scan(&count)
	helper.PanicIfError(err)
	if count == 0 {
		return cartItem, errors.New("cart item not found")
	}

	if cartItem.ProductId != 0 && cartItem.Quantity == 0 {
		script := "update cart_items set product_id=? where cart_item_id=?"
		_, err := tx.ExecContext(ctx, script, cartItem.ProductId, cartItem.CartItemId)
		helper.PanicIfError(err)

		return cartItem, err

	} else if cartItem.ProductId == 0 && cartItem.Quantity != 0 {
		script := "update cart_items set quantity=? where cart_item_id=?"
		_, err := tx.ExecContext(ctx, script, cartItem.Quantity, cartItem.CartItemId)
		helper.PanicIfError(err)

		return cartItem, err

	} else {
		script := "update cart_items set product_id=?,quantity=? where cart_item_id=?"
		_, err := tx.ExecContext(ctx, script, cartItem.ProductId, cartItem.Quantity, cartItem.CartItemId)
		helper.PanicIfError(err)

		return cartItem, err
	}
}

func (repository *CartRepositoryImpl) DeleteCart(ctx context.Context, tx *sql.Tx, cartItemId int) error {
	checkScript := "select count(*) from cart_items where cart_item_id=?"
	var count int
	err := tx.QueryRowContext(ctx, checkScript, cartItemId).Scan(&count)
	helper.PanicIfError(err)
	if count == 0 {
		return errors.New("cart item not found")
	}

	script := "delete from cart_items where cart_item_id =?"
	_, err = tx.ExecContext(ctx, script, cartItemId)
	helper.PanicIfError(err)

	return err
}
