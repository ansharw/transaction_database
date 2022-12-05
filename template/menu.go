package template

import (
	"database/sql"
	"fmt"
	"os"
	"transaction_database/controller"
	"transaction_database/helper"
	"transaction_database/repository"
)

func Menu(db *sql.DB) {
	productsRepository := repository.NewProductsRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	transactionDetailsRepository := repository.NewTransactionDetailsRepository(db)
	vouchersRepository := repository.NewVouchersRepository(db)

	// ngehandler repository buat di panggil di end user
	transactionHandler := controller.NewTransactionHandler(db, productsRepository, transactionRepository, transactionDetailsRepository, vouchersRepository)

	// ini yang manggil handler untuk templating di end user
	transactionTemplate := NewTransactionTemplate(db, transactionHandler)

	helper.ClearScreen()
	fmt.Println("===================================")
	fmt.Println("SELAMAT DATANG DI PROGRAM PENJUALAN")
	fmt.Println("===================================")
	fmt.Println("1. Tambah Penjualan               =")
	fmt.Println("2. Daftar List Penjualan          =")
	fmt.Println("3. Tampilkan Daftar Produk        =")
	fmt.Println("4. Tampilkan Daftar Voucher       =")
	fmt.Println("5. Exit                           =")
	fmt.Println("===================================")

	var input int
	fmt.Print("Pilih menu : ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		transactionTemplate.AddTransactionTemplate()
	case 2:
		transactionTemplate.EachTransaction()
	case 3:
		transactionTemplate.ListProduct()
	case 4:
		transactionTemplate.ListVoucher()
	case 5:
		os.Exit(3)
	}
}
