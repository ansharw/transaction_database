package controller

import (
	"database/sql"
	"transaction_database/repository"
)

// kumpulan fungsi yang nanti dipakai di menu / buat end user
type TransactionHandler interface {
	
}

// nyimpen
type transactionHandler struct {
	db                           *sql.DB
	productsRepository           repository.ProductsRepository
	transactionRepository        repository.TransactionRepository
	transactionDetailsRepository repository.TransactionDetailsRepository
	vouchersRepository           repository.VouchersRepository
}

func NewTransactionHandler(db *sql.DB, productsRepository repository.ProductsRepository, transactionRepository repository.TransactionRepository, transactionDetailsRepository repository.TransactionDetailsRepository, vouchersRepository repository.VouchersRepository) *transactionHandler {
	return &transactionHandler{db, productsRepository, transactionRepository, transactionDetailsRepository, vouchersRepository}
}
