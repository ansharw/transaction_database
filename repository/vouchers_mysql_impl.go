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

func (repo *vouchersRepository) FindVoucher(ctx context.Context, code string) (model.Vouchers, error) {
	var vouch model.Vouchers
	query := "SELECT id, code, value FROM vouchers WHERE code LIKE ?"

	rows := repo.db.QueryRowContext(ctx, query, code)
	err := rows.Scan(vouch.GetId(), vouch.GetCode(), vouch.GetValue())
	if err != nil {
		return vouch, err
	}
	return vouch, nil
}