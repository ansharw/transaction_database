package template

import (
	"database/sql"
	"fmt"
	"transaction_database/controller"
	"transaction_database/helper"
	"transaction_database/model"
)

type transactionTemplate struct {
	db                 *sql.DB
	transactionHandler controller.TransactionHandler
}

func NewTransactionTemplate(db *sql.DB, transactionHandler controller.TransactionHandler) *transactionTemplate {
	return &transactionTemplate{db, transactionHandler}
}

func (template *transactionTemplate) ListTransactions(trx []model.Transaction) {
	fmt.Println("List Transaksi")
	for i, v := range trx {
		fmt.Printf("%d %s\n", i+1, *v.GetNumber())
	}
}

func (template *transactionTemplate) EachTransaction() {
	helper.ClearScreen()
	transactions, err := template.transactionHandler.GetTransactions()
	if err != nil {
		panic(err)
	}

	if err == sql.ErrNoRows {
		fmt.Println("tidak ada transaksi")
	} else {
		template.ListTransactions(transactions)
		fmt.Print("Masukkan Nomer Transaksi :")
		var nomorTransaksi string
		fmt.Scanln(&nomorTransaksi)
		fmt.Println("Nomor Transaksi: ", nomorTransaksi)
		transaction, err := template.transactionHandler.GetTransactionByNumber(nomorTransaksi)
		if err == sql.ErrNoRows {
			fmt.Println("Data tidak ditemukan")
		}

		fmt.Println("Id Transaksi: ", *transaction.GetId())
		date := transaction.GetDate().Format("2006-01-02")
		fmt.Printf("Tanggal Transaksi\t %v\n", date)
		email := transaction.GetEmail()
		fmt.Println("Email :", *email)
		phone := transaction.GetPhone()
		fmt.Println("Phone: ", *phone)
		for _, v := range *transaction.GetTransactionDetails() {
			fmt.Println(*v.GetProdName(), *v.GetQty(), *v.GetTotal())
		}
	}
}
