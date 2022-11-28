package repository

import (
	"context"
	"database/sql"
	"transaction_database/model"
)

type TransactionRepository interface {
	AddTrx(ctx context.Context, tx *sql.Tx, trxD []model.TransactionDetails, trx model.Transaction) ([]model.TransactionDetails, error)
	FindById(ctx context.Context, trxId int) (model.Transaction, error)
	FindAll(ctx context.Context) ([]model.Transaction, error)
}
