package repository

import (
	"context"
	"database/sql"
	"transaction_database/model"
)

type TransactionDetailsRepository interface {
	// AddTrxDetails(ctx context.Context, trxDetails model.TransactionDetails) (model.TransactionDetails, error)
	// AddTrxDetails(ctx context.Context, tx *sql.Tx, trxDetails []model.TransactionDetails, trxId int) ([]model.TransactionDetails, error)
	// FindById(ctx context.Context, trxDetailsId int) (model.TransactionDetails, error)
	// SearchById(ctx context.Context, id int) (model.TransactionDetails, error)

	// buat Get TrxDetail by TrxId 
	GetTrxDetailsByTrxId(ctx context.Context, trxId int) ([]model.TransactionDetails, error)
	// AddTrxDetails(ctx context.Context, tx *sql.Tx, trxDetails []model.TransactionDetails, trxId int) ([]model.TransactionDetails, error)
	// eksperimen
	// AddTrxDetails(ctx context.Context, trxDetails model.TransactionDetails) (model.TransactionDetails, error)
	// DeleteByTransactionId(ctx context.Context, tx *sql.Tx, trxId int) error

	// ini buat Add trx detail yang masukin id trx, sehabis AddTrx
	AddTrxDetails(ctx context.Context, tx *sql.Tx, trxDetails model.TransactionDetails, trxId int) (model.TransactionDetails, error)
}
