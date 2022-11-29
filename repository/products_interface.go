package repository

import (
	"context"
	"transaction_database/model"
)

type ProductsRepository interface {
	FindAll(ctx context.Context) ([]model.Products, error)
	FindById(ctx context.Context, product_id int) (model.Products, error)
}
