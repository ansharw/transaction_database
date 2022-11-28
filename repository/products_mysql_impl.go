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
