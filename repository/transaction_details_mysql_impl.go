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

// backup
// func (repo *transactionDetailsRepository) GetTrxDetailsByTrxId(ctx context.Context, trxId int) ([]model.TransactionDetails, error) {
// 	var trxDetails []model.TransactionDetails
// 	var query string = "SELECT id, transaction_id, product_id, product_name, price, quantity, total, created_at FROM transaction_details WHERE transaction_id=?"

// 	rows, err := repo.db.QueryContext(ctx, query, trxId)
// 	if err != nil {
// 		return trxDetails, err
// 	}
// 	for rows.Next() {
// 		var trxDetail model.TransactionDetails
// 		rows.Scan(trxDetail.GetId(), trxDetail.GetPrice(), trxDetail.GetQty(), trxDetail.GetTotal(), trxDetail.GetCreatedAt())
// 		trxDetails = append(trxDetails, trxDetail)
// 	}
// 	return trxDetails, err
// }

// eksperimen
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

// backup
// func (repo *transactionDetailsRepository) AddTrxDetails(ctx context.Context, tx *sql.Tx, trxDetails []model.TransactionDetails, trxId int) ([]model.TransactionDetails, error) {
// 	var query string = "INSERT INTO transaction_details(transaction_id, product_id, product_name, price, quantity, total, created_at) VALUES(?,?,?,?,?,?,?)"

// 	stmt, err := tx.PrepareContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	for _, v := range trxDetails {
// 		res, err := stmt.ExecContext(ctx, trxId, v.GetProdId(), v.GetProdName(), v.GetPrice(), v.GetQty(), v.GetTotal(), v.GetCreatedAt())
// 		if err != nil {
// 			return nil, err
// 		}
// 		lastInsertId, _ := res.LastInsertId()
// 		id := int(lastInsertId)
// 		v.SetId(&id)
// 	}
// 	return trxDetails, nil
// }

// eksperimen
// func (repo *transactionDetailsRepository) AddTrxDetails(ctx context.Context, trxDetails model.TransactionDetails) (model.TransactionDetails, error) {
// 	var query string = "INSERT INTO transaction_details(transaction_id, product_id, product_name, price, quantity, total) VALUES(?,?,?,?,?,?,?)"

// 	res, err := repo.db.ExecContext(ctx, query, trxDetails.GetProdId(), trxDetails.GetProdName(), trxDetails.GetPrice(), trxDetails.GetQty(), trxDetails.GetTotal())

// 	if err != nil {
// 		return trxDetails, err
// 	}
// 	lastInsertId, _ := res.LastInsertId()
// 	id := int(lastInsertId)
// 	trxDetails.SetId(&id)

// 	return trxDetails, err
// }

// eksperimen 2
// add transaction detail
func (repo *transactionDetailsRepository) AddTrxDetails(ctx context.Context, tx *sql.Tx, trxDetails model.TransactionDetails, trxId int) (model.TransactionDetails, error) {
	var query string = "INSERT INTO transaction_details(transaction_id, product_id, product_name, price, quantity, total) VALUES(?,?,?,?,?,?)"

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return trxDetails, err
	}
	defer stmt.Close()

	// masukin query nya sesuai dengan yang di minta
	res, err := stmt.ExecContext(ctx, trxId, trxDetails.GetProdId(), trxDetails.GetProdName(), trxDetails.GetPrice(), trxDetails.GetQty(), trxDetails.GetTotal())

	if err != nil {
		return trxDetails, err
	}

	// ini insert lastId nya buat trx details
	lastInsertId, _ := res.LastInsertId()
	id := int(lastInsertId)
	trxDetails.SetId(&id)
	trxDetails.SetTrxId(&trxId)

	return trxDetails, err
}

// func (repo *transactionDetailsRepository) FindById(ctx context.Context, trxDetailsId int) (model.TransactionDetails, error) {
// 	var query string = "SELECT id, transaction_id, product_id, product_name, price, quantity, total, created_at FROM transaction_details WHERE id=?"
// 	var trxDetails model.TransactionDetails

// 	row := repo.db.QueryRowContext(ctx, query, trxDetailsId)
// 	err := row.Scan(trxDetails.GetId(), trxDetails.GetTrxId(), trxDetails.GetProdId(), trxDetails.GetProdName(), trxDetails.GetPrice(), trxDetails.GetQty(), trxDetails.GetTotal(), trxDetails.GetCreatedAt())
// 	if err != nil {
// 		return trxDetails, err
// 	}

// 	return trxDetails, nil
// }

func (repo *transactionDetailsRepository) DeleteByTransactionId(ctx context.Context, tx *sql.Tx, trxId int) error {
	var query string = "DELETE FROM transaction_details WHERE transaction_id=?"

	_, err := tx.ExecContext(ctx, query, trxId)
	if err != nil {
		return err
	}

	return nil
}
