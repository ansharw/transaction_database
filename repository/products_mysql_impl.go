package repository

import (
	"context"
	"database/sql"
	"transaction_database/model"
)

type productsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *productsRepository {
	return &productsRepository{db}
}

func (repo *productsRepository) FindAll(ctx context.Context) ([]model.Products, error) {
	var query string = "SELECT id, name, price FROM products"
	var products []model.Products

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return products, err
	}
	for rows.Next() {
		var product model.Products
		rows.Scan(product.GetId(), product.GetName(), product.GetPrice())
		products = append(products, product)
	}
	return products, nil
}

func (repo *productsRepository) FindProduct(ctx context.Context, id int) (model.Products, error) {
	var prod model.Products
	query := "SELECT id, name, price FROM products WHERE id=?"

	rows := repo.db.QueryRowContext(ctx, query, id)
	err := rows.Scan(prod.GetId(), prod.GetName(), prod.GetPrice())
	if err != nil {
		return prod, err
	}
	return prod, nil
}