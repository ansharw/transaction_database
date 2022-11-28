package repository

import (
	"context"
	"database/sql"
	"transaction_database/model"
)

type vouchersRepository struct {
	db *sql.DB
}

func NewVouchersRepository(db *sql.DB) *vouchersRepository {
	return &vouchersRepository{db}
}

func (repo *vouchersRepository) FindAll(ctx context.Context) ([]model.Vouchers, error) {
	var query string = "SELECT id, code, value FROM vouchers"
	var vouchers []model.Vouchers

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return vouchers, err
	}
	for rows.Next() {
		var voucher model.Vouchers
		rows.Scan(voucher.GetId(), voucher.GetCode(), voucher.GetValue())
		vouchers = append(vouchers, voucher)
	}
	return vouchers, nil
}
