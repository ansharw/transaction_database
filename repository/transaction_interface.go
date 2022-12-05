package repository

import (
	"context"
	"transaction_database/model"
)

type TransactionRepository interface {
	AddTrx(ctx context.Context, trx model.Transaction) (model.Transaction, error)
	FindByNumber(ctx context.Context, trxNumber string) (model.Transaction, error)
	FindAll(ctx context.Context) ([]model.Transaction, error)
}
