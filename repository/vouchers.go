package repository

import "database/sql"

type VouchersRepository interface {
	
}

type vouchersRepository struct {
	db *sql.DB
}

func NewVouchersRepository(db *sql.DB) *vouchersRepository {
	return &vouchersRepository{db}
}
