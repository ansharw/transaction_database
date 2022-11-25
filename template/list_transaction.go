package template

import (
	"database/sql"
	"transaction_database/controller"
)

type transactionTemplate struct {
	db                 *sql.DB
	transactionHandler controller.TransactionHandler
}

func NewTransactionTemplate(db *sql.DB, transactionHandler controller.TransactionHandler) *transactionTemplate {
	return &transactionTemplate{db, transactionHandler}
}

func (template *transactionTemplate) ListTransaction() {

}
