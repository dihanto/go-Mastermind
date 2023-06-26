package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/go-mastermind/model/entity"
)

type ProductRepositoryImpl struct {
}

func (repository *ProductRepositoryImpl) AddProduct(ctx context.Context, tx *sql.Tx, request entity.Product) (product entity.Product, err error) {
	query := "INSERT INTO products (id, id_seller, name, price, created_at) VALUES($1, $2, $3, $4, $5)"
	_, err = tx.ExecContext(ctx, query, request.Id, request.IdSeller, request.Name, request.Price, request.CreatedAt)
	if err != nil {
		return
	}
	product = entity.Product{
		Id:        request.Id,
		IdSeller:  request.IdSeller,
		Name:      request.Name,
		Price:     request.Price,
		CreatedAt: request.CreatedAt,
	}
	return
}

func (repository *ProductRepositoryImpl) GetProduct(ctx context.Context, tx *sql.Tx) (products []entity.Product, err error) {
	query := "SELECT id, name, price FROM products"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.Id, product.Name, product.Price)
		if err != nil {
			return
		}
		products = append(products, product)
	}
	return
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (product entity.Product, err error) {
	query := "Select id_seller, name, price, created_at, updated_at FROM products WHERE id=$1"
	err = tx.QueryRowContext(ctx, query, id).Scan(&product.IdSeller, product.Name, product.Price, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		return
	}

	product.Id = id

	return
}

func (repository *ProductRepositoryImpl) UpdateProduct(ctx context.Context, tx *sql.Tx, request entity.Product) (product entity.Product, err error) {
	query := "UPDATE products SET name=$1, price=$2 WHERE id=$3"
	_, err = tx.ExecContext(ctx, query, request.Name, request.Price, request.Id)
	if err != nil {
		return
	}
	queryProduct := "SELECT name, price, created_at, updated_at FROM products WHERE id=$1"
	err = tx.QueryRowContext(ctx, queryProduct, request.Id).Scan(&product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return
	}
	product.Id = request.Id

	return
}

func (repository *ProductRepositoryImpl) DeleteProduct(ctx context.Context, tx *sql.Tx, deleteTime int32, id int) (err error) {
	query := "UPDATE products SET deleted_at=$1 WHERE id=$2"
	_, err = tx.ExecContext(ctx, query, deleteTime, id)
	if err != nil {
		return
	}
	return
}
