package repository

import (
	"context"
	"database/sql"
	"transaction_database/model"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db}
}

// backup
// func (repo *transactionRepository) AddTrx(ctx context.Context, trx model.Transaction) (model.Transaction, error) {

// 	var query string = "INSERT INTO transaction(number, customer_name, email, phone, date, quantity, discount, total, pay, created_at) VALUES(?,?,?,?,?,?,?,?,?)"

// 	res, err := repo.db.ExecContext(ctx, query, trx.GetNumber(), trx.GetCustomerName(), trx.GetEmail(), trx.GetPhone(), trx.GetDate(), trx.GetQty(), trx.GetDiscount(), trx.GetTotal(), trx.GetPay(), trx.GetCreatedAt())
// 	if err != nil {
// 		return trx, nil
// 	}
// 	lastInsertId, _ := res.LastInsertId()
// 	id := int(lastInsertId)
// 	trx.SetId(&id)

// 	return trx, nil
// }

// eksperimen
func (repo *transactionRepository) AddTrx(ctx context.Context, tx *sql.Tx, trxD []model.TransactionDetails, trx model.Transaction) ([]model.TransactionDetails, error) {

	var trxDetail model.TransactionDetails
	var query string = "INSERT INTO transaction(number, customer_name, email, phone, date, quantity, discount, total, pay) VALUES(?,?,?,?,?,?,?,?,?)"

	res, err := repo.db.ExecContext(ctx, query, trx.GetNumber(), trx.GetEmail(), trx.GetPhone(), trx.GetDate(), trx.GetQty(), trx.GetDiscount(), trx.GetTotal(), trx.GetPay())
	if err != nil {
		return trxD, nil
	}

	lastInsertId, _ := res.LastInsertId()
	id := int(lastInsertId)
	trx.SetId(&id)
	trxDetail.SetTrxId(trx.GetId())

	return trxD, nil
}

// eksperimen 2

// backup
// func (repo *transactionRepository) FindById(ctx context.Context, trxId int) (model.Transaction, error) {
// 	var query string = "SELECT id, number, customer_name, email, phone, date, quantity, discount, total, pay, created_at FROM transaction WHERE id=?"
// 	var trx model.Transaction

// 	row := repo.db.QueryRowContext(ctx, query, trxId)
// 	err := row.Scan(trx.GetId(), trx.GetNumber(), trx.GetCustomerName(), trx.GetEmail(), trx.GetPhone(), trx.GetDate(), trx.GetQty(), trx.GetDiscount(), trx.GetTotal(), trx.GetPay(), trx.GetCreatedAt())
// 	if err != nil {
// 		return trx, err
// 	}

// 	return trx, nil
// }

// eksperimen
func (repo *transactionRepository) FindById(ctx context.Context, trxId int) (model.Transaction, error) {
	var query string = "SELECT id, number, customer_name, email, phone, total, pay FROM transaction WHERE id=?"
	var trx model.Transaction

	row := repo.db.QueryRowContext(ctx, query, trxId)
	err := row.Scan(trx.GetId(), trx.GetNumber(), trx.GetCustomerName(), trx.GetEmail(), trx.GetPhone(), trx.GetTotal(), trx.GetPay())
	if err != nil {
		return trx, err
	}

	return trx, nil
}

// backup
// func (repo *transactionRepository) FindAll(ctx context.Context) ([]model.Transaction, error) {
// 	var query string = "SELECT id, number, customer_name, email, phone, date, quantity, discount, total, pay, created_at FROM transaction"
// 	var transactions []model.Transaction

// 	rows, err := repo.db.QueryContext(ctx, query)
// 	if err != nil {
// 		return transactions, err
// 	}

// 	for rows.Next() {
// 		var transaction model.Transaction
// 		rows.Scan(transaction.GetId(), transaction.GetNumber(), transaction.GetCustomerName(), transaction.GetEmail(), transaction.GetPhone(), transaction.GetDate(), transaction.GetQty(), transaction.GetDiscount(), transaction.GetTotal(), transaction.GetPay(), transaction.GetCreatedAt())
// 		transactions = append(transactions, transaction)
// 	}
// 	return transactions, nil
// }

// eksperimen
func (repo *transactionRepository) FindAll(ctx context.Context) ([]model.Transaction, error) {
	var query string = "SELECT id, number, customer_name, email, phone, total, pay FROM transaction"
	var transactions []model.Transaction

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return transactions, err
	}

	for rows.Next() {
		var transaction model.Transaction
		rows.Scan(transaction.GetId(), transaction.GetNumber(), transaction.GetCustomerName(), transaction.GetEmail(), transaction.GetPhone(), transaction.GetTotal(), transaction.GetPay())
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
