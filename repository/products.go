package repository

import "database/sql"

type ProductsRepository interface {
	
}

type productsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *productsRepository {
	return &productsRepository{db}
}
