package repository

import "database/sql"

type productsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *productsRepository {
	return &productsRepository{db}
}

