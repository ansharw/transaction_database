package repository

import (
	"context"
	"database/sql"
	"transaction_database/model"
)

type transactionDetailsRepository struct {
	db *sql.DB
}

func NewTransactionDetailsRepository(db *sql.DB) *transactionDetailsRepository {
	return &transactionDetailsRepository{db}
}

func (repo *transactionDetailsRepository) GetTrxDetailsByTrxId(ctx context.Context, trxId int) ([]model.TransactionDetails, error) {
	var trxDetails []model.TransactionDetails
	var query string = "SELECT id, price, quantity, total FROM transaction_details WHERE transaction_id=?"

	rows, err := repo.db.QueryContext(ctx, query, trxId)
	if err != nil {
		return trxDetails, err
	}
	for rows.Next() {
		var trxDetail model.TransactionDetails
		rows.Scan(trxDetail.GetId(), trxDetail.GetPrice(), trxDetail.GetQty(), trxDetail.GetTotal())
		trxDetails = append(trxDetails, trxDetail)
	}
	return trxDetails, err
}


func (repo *transactionDetailsRepository) AddTrxDetails(ctx context.Context, trxDetails model.TransactionDetails, trxId int) (model.TransactionDetails, error) {
	var query string = "INSERT INTO transaction_details(transaction_id, product_id, product_name, price, quantity, total) VALUES(?,?,?,?,?,?)"

	// masukin query nya sesuai dengan yang di minta
	res, err := repo.db.ExecContext(ctx, query, trxId, *trxDetails.GetProdId(), *trxDetails.GetProdName(), *trxDetails.GetPrice(), *trxDetails.GetQty(), *trxDetails.GetTotal())

	if err != nil {
		return trxDetails, err
	}

	// ini insert lastId nya buat trx details
	lastInsertId, _ := res.LastInsertId()
	id := int(lastInsertId)
	trxDetails.SetTrxId(&trxId)
	trxDetails.SetId(&id)

	return trxDetails, nil
}

func (repo *transactionDetailsRepository) DeleteByTransactionId(ctx context.Context, tx *sql.Tx, trxId int) error {
	var query string = "DELETE FROM transaction_details WHERE transaction_id=?"

	_, err := tx.ExecContext(ctx, query, trxId)
	if err != nil {
		return err
	}

	return nil
}
