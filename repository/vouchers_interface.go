package repository

import (
	"context"
	"transaction_database/model"
)

type VouchersRepository interface {
	FindAll(ctx context.Context) ([]model.Vouchers, error)
	FindVoucher(ctx context.Context, code string) (model.Vouchers, error)
}
