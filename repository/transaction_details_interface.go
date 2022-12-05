package repository

import (
	"context"
	"transaction_database/model"
)

type TransactionDetailsRepository interface {
	// buat Get TrxDetail by TrxId
	GetTrxDetailsByTrxId(ctx context.Context, trxId int) ([]model.TransactionDetails, error)
	// ini buat Add trx detail yang masukin id trx, sehabis AddTrx
	AddTrxDetails(ctx context.Context, trxDetails model.TransactionDetails, trxId int) (model.TransactionDetails, error)
}
