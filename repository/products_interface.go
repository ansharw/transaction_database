package repository

import (
	"context"
	"transaction_database/model"
)

type ProductsRepository interface {
	FindAll(ctx context.Context) ([]model.Products, error)
	FindProduct(ctx context.Context, id int) (model.Products, error)
}
