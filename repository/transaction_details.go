package repository

import "database/sql"

type TransactionDetailsRepository interface {
	
}

type transactionDetailsRepository struct {
	db *sql.DB
}

func NewTransactionDetailsRepository(db *sql.DB) *transactionDetailsRepository {
	return &transactionDetailsRepository{db}
}
