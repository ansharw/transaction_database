package repository

import (
	"context"
	"transaction_database/model"
)

type TransactionRepository interface {
	// AddTrx(ctx context.Context, tx *sql.Tx, trxD []model.TransactionDetails, trx model.Transaction) ([]model.TransactionDetails, error)
	AddTrx(ctx context.Context, trx model.Transaction) (model.Transaction, error) 
	FindById(ctx context.Context, trxId int) (model.Transaction, error)
	FindAll(ctx context.Context) ([]model.Transaction, error)
}
