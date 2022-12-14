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

// addTrx untuk add transaction ke database()
func (repo *transactionRepository) AddTrx(ctx context.Context, trx model.Transaction) (model.Transaction, error) {

	var query string = "INSERT INTO transactions(number, customer_name, email, phone, date, quantity, discount, total, pay) VALUES(?,?,?,?,?,?,?,?,?)"

	res, err := repo.db.ExecContext(ctx, query, *trx.GetNumber(), *trx.GetCustomerName(), *trx.GetEmail(), *trx.GetPhone(), *trx.GetDate(), *trx.GetQty(), *trx.GetDiscount(), *trx.GetTotal(), *trx.GetPay())

	if err != nil {
		return trx, nil
	}

	lastInsertId, _ := res.LastInsertId()
	id := int(lastInsertId)
	trx.SetId(&id)

	return trx, nil
}

// eksperimen
func (repo *transactionRepository) FindByNumber(ctx context.Context, trxNumber string) (model.Transaction, error) {
	var query string = "SELECT id, number, customer_name, email, phone, total, pay, date FROM transactions WHERE number = ?"
	var trx model.Transaction

	row := repo.db.QueryRowContext(ctx, query, trxNumber)
	err := row.Scan(trx.GetId(), trx.GetNumber(), trx.GetCustomerName(), trx.GetEmail(), trx.GetPhone(), trx.GetTotal(), trx.GetPay(), trx.GetDate())
	if err != nil {
		return trx, err
	}

	return trx, nil
}

// eksperimen
func (repo *transactionRepository) FindAll(ctx context.Context) ([]model.Transaction, error) {
	var query string = "SELECT id, number, customer_name, email, phone, total, pay FROM transactions"
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
