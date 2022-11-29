package controller

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"transaction_database/model"
	"transaction_database/repository"
)

// kumpulan fungsi yang nanti dipakai di menu / buat end user
type TransactionHandler interface {
	GetProducts() ([]model.Products, error)
	GetVouchers() ([]model.Vouchers, error)

	GetTransactions() ([]model.Transaction, error)
	// AddTransaction(custName, email, phone string, total, pay float64, trxDetails []model.TransactionDetails) error
	// AddTransaction(prodsId int, quantity int, custName, custEmail, custPhone, discount string, pay float64) error
	// AddTransaction(prodsId int, quantity int, custName, custEmail, custPhone string, discount string, pay float64) error
	AddTransaction(prodsId int, quantity int, custName, custEmail, custPhone string, discount string, pay float64) (model.Transaction, []model.TransactionDetails, error)
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

func (handler *transactionHandler) GetProducts() ([]model.Products, error) {
	ctx := context.Background()

	products, err := handler.productsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (handler *transactionHandler) ShowProducts() ([]model.Products, error) {
	ctx := context.Background()

	fmt.Println("====================================================")
	fmt.Println("ID || Nama \t\t|| Price \t\t  ||")
	fmt.Println("====================================================")

	products, err := handler.productsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		fmt.Println("Data Kosong")
	} else {
		for _, v := range products {
			fmt.Printf("%v  || %v \t|| %v \t\t  ||\n", *v.GetId(), *v.GetName(), *v.GetPrice())
		}
	}
	fmt.Println("====================================================")

	return products, nil
}

func (handler *transactionHandler) GetVouchers() ([]model.Vouchers, error) {
	ctx := context.Background()

	vouchers, err := handler.vouchersRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return vouchers, nil
}

// eksperimen
func (handler *transactionHandler) GetTransactions() ([]model.Transaction, error) {
	ctx := context.Background()

	transactions, err := handler.transactionRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for i, v := range transactions {
		transactionDetails, err := handler.transactionDetailsRepository.GetTrxDetailsByTrxId(ctx, *v.GetId())
		if err != nil {
			fmt.Println("ini dari transaction controller ", v)
			fmt.Println(v.GetId())
			return nil, err
		}
		transactions[i].SetTransactionDetails(transactionDetails)
	}

	return transactions, nil
}

// eksperimen
// ini untuk manggil keseluruhan yaitu transaction detail abis itu transaction
// func (handler *transactionHandler) AddTransaction(prodsId int, quantity int, custName, custEmail, custPhone, discount string, pay float64) error {
// 	ctx := context.Background()

// 	var transaction model.Transaction
// 	var transactionDetail model.TransactionDetails
// 	var transactionDetails []model.TransactionDetails

// 	transactionDetail.SetProdId(&prodsId)
// 	transactionDetail.SetQty(&quantity)

// 	products, err := handler.GetProducts()
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, v := range products {
// 		if v.GetId() == &prodsId {
// 			transactionDetail.SetProdName(v.GetName())
// 			transactionDetail.SetPrice(v.GetPrice())
// 		}
// 	}

// 	var price float64 = *transactionDetail.GetPrice()
// 	var qty float64 = float64(quantity)
// 	total := qty * price
// 	transactionDetail.SetTotal(&total)

// 	trxDetail, err := handler.transactionDetailsRepository.AddTrxDetails(ctx, transactionDetail)
// 	if err != nil {
// 		return err
// 	}
// 	transactionDetails = append(transactionDetails, trxDetail)

// 	tx, err := handler.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	timeNow := time.Now()
// 	transaction.SetDate(&timeNow)

// 	trxNumber := GenerateTrxNumber()
// 	transaction.SetNumber(trxNumber)
// 	transaction.SetEmail(&custEmail)
// 	transaction.SetPhone(&custPhone)

// 	vouchers, err := handler.GetVouchers()
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, v := range vouchers {
// 		if *v.GetCode() == discount {
// 			total := transactionDetail.GetTotal()
// 			disc := *v.GetValue() / float64(100)
// 			discounting := *total * disc
// 			totalFinal := *total - discounting
// 			transaction.SetDiscount(&totalFinal)
// 		}
// 	}
// 	transaction.SetPay(&pay)

// 	trxFinal, err := handler.transactionRepository.AddTrx(ctx, tx, transactionDetails, transaction)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	tx.Commit()
// 	transaction.SetTransactionDetails(trxFinal)

// 	return nil
// }

// eksperimen 2
// ini untuk manggil keseluruhan yaitu transaction detail abis itu transaction
// func (handler *transactionHandler) AddTransaction(prodsId int, quantity int, custName, custEmail, custPhone string, discount string, pay float64) (model.Transaction, []model.TransactionDetails, error) {
// 	ctx := context.Background()

// 	var transaction model.Transaction
// 	var transactionDetail model.TransactionDetails
// 	var transactionDetails []model.TransactionDetails

// 	// masukkin data prodId dan qty
// 	transactionDetail.SetProdId(&prodsId)
// 	transactionDetail.SetQty(&quantity)

// 	// ini buat dapetin productnya
// 	// kalo id nya sama ya dimasukin aja
// 	products, err := handler.GetProducts()
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, v := range products {
// 		if *v.GetId() == prodsId {
// 			transactionDetail.SetProdName(v.GetName())
// 			transactionDetail.SetPrice(v.GetPrice())
// 		}
// 	}

// 	var price float64 = *transactionDetail.GetPrice()
// 	var qty float64 = float64(quantity)
// 	total := qty * price
// 	transactionDetail.SetTotal(&total)

// 	var now = time.Now()
// 	nows := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
// 	layout := "2006-01-02 15:04:05"
// 	date, _ := time.Parse(layout, nows)
// 	trxNumber := GenerateTrxNumber()

// 	vouchers, err := handler.GetVouchers()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// cek discount
// 	if total > 300000 {
// 		for _, v := range vouchers {
// 			if *v.GetCode() == discount {
// 				total := transactionDetail.GetTotal()
// 				transaction.SetTotal(transactionDetail.GetTotal())
// 				disc := *v.GetValue() / float64(100)
// 				discounting := *total * disc
// 				// totalFinal := *total - discounting
// 				// transaction.SetDiscount(&totalFinal)
// 				transaction.SetDiscount(&discounting)
// 			} else if discount == "" || *v.GetCode() != discount {
// 				var nol *float64
// 				transaction.SetDiscount(nol)
// 			}
// 		}
// 	}

// 	// masukin data ke transaction
// 	transaction.SetPay(&pay)
// 	transaction.SetDate(&date)
// 	transaction.SetNumber(&trxNumber)
// 	transaction.SetEmail(&custEmail)
// 	transaction.SetPhone(&custPhone)
// 	transaction.SetCustomerName(&custName)
// 	transaction.SetQty(transactionDetail.GetQty())

// 	// fmt.Println(date)
// 	// fmt.Println(transaction.GetDate())
// 	fmt.Println("ini transaction detail", transactionDetail)
// 	fmt.Println("ini transaction", transaction)

// 	// masukin semua transaction ke handler
// 	trx, err := handler.transactionRepository.AddTrx(ctx, transaction)
// 	// fmt.Println(trx)
// 	if err != nil {
// 		return transaction, transactionDetails, err
// 	}
// 	// fmt.Println(trx)

// 	tx, err := handler.db.BeginTx(ctx, nil)
// 	// fmt.Println(err)
// 	if err != nil {
// 		return transaction, transactionDetails, err
// 	}

// 	// masukin semua transaction detail ke handler beserta id trx nya
// 	trxD, err := handler.transactionDetailsRepository.AddTrxDetails(ctx, tx, transactionDetail, *trx.GetId())
// 	// fmt.Println(err)
// 	if err != nil {
// 		tx.Rollback()
// 		return transaction, transactionDetails, err
// 	}
// 	// fmt.Println(trxD)

// 	tx.Commit()
// 	transactionDetails = append(transactionDetails, trxD)

// 	transaction.SetTransactionDetails(transactionDetails)
// 	return transaction, transactionDetails, nil
// }

func GenerateTrxNumber() string {
	var trxNumber string = strconv.Itoa(rand.Int())
	return trxNumber[:5]
}

// Eksperimen 3
// ini untuk manggil keseluruhan yaitu transaction detail abis itu transaction
func (handler *transactionHandler) AddTransaction(prodsId int, quantity int, custName, custEmail, custPhone string, discount string, pay float64) (model.Transaction, []model.TransactionDetails, error) {
	ctx := context.Background()

	var transaction model.Transaction
	var transactionDetail model.TransactionDetails
	var transactionDetails []model.TransactionDetails

outer:
	for {
		handler.ShowProducts()
		// template.InputNameOfProduct(&nameProduct)
		template.InputIdOfProduct(&idProduct)
		// quantity of product
		template.InputQtyOfProduct(&qtyProduct)

		tempMap := map[string]int{
			"idProduct":  idProduct,
			"qtyProduct": qtyProduct,
		}
		tempSliceMap = append(tempSliceMap, tempMap)
		fmt.Println("Input data kembali? (y/n)")
		var option string
		fmt.Scanln(&option)
		switch option {
		case "n":
			break outer
		case "y":
			continue
		default:
			break outer
		}
	}

	// masukkin data prodId dan qty
	transactionDetail.SetProdId(&prodsId)
	transactionDetail.SetQty(&quantity)

	// ini buat dapetin productnya
	// kalo id nya sama ya dimasukin aja
	products, err := handler.GetProducts()
	if err != nil {
		panic(err)
	}

	for _, v := range products {
		if *v.GetId() == prodsId {
			transactionDetail.SetProdName(v.GetName())
			transactionDetail.SetPrice(v.GetPrice())
		}
	}

	var price float64 = *transactionDetail.GetPrice()
	var qty float64 = float64(quantity)
	total := qty * price
	transactionDetail.SetTotal(&total)

	var now = time.Now()
	nows := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	layout := "2006-01-02 15:04:05"
	date, _ := time.Parse(layout, nows)
	trxNumber := GenerateTrxNumber()

	vouchers, err := handler.GetVouchers()
	if err != nil {
		panic(err)
	}

	template.ShowVoucher()
	fmt.Print("Masukkan Code Voucher : ")
	fmt.Scanln(&discount)
	fmt.Print("Masukkan Uang Anda : ")
	fmt.Scanln(&pay)

	// cek discount
	if total > 300000 {
		for _, v := range vouchers {
			if *v.GetCode() == discount {
				total := transactionDetail.GetTotal()
				transaction.SetTotal(transactionDetail.GetTotal())
				disc := *v.GetValue() / float64(100)
				discounting := *total * disc
				// totalFinal := *total - discounting
				// transaction.SetDiscount(&totalFinal)
				transaction.SetDiscount(&discounting)
			} else if discount == "" || *v.GetCode() != discount {
				var nol *float64
				transaction.SetDiscount(nol)
			}
		}
	}

	// masukin data ke transaction
	transaction.SetPay(&pay)
	transaction.SetDate(&date)
	transaction.SetNumber(&trxNumber)
	transaction.SetEmail(&custEmail)
	transaction.SetPhone(&custPhone)
	transaction.SetCustomerName(&custName)
	transaction.SetQty(transactionDetail.GetQty())

	// fmt.Println(date)
	// fmt.Println(transaction.GetDate())
	fmt.Println("ini transaction detail", transactionDetail)
	fmt.Println("ini transaction", transaction)

	// masukin semua transaction ke handler
	trx, err := handler.transactionRepository.AddTrx(ctx, transaction)
	// fmt.Println(trx)
	if err != nil {
		return transaction, transactionDetails, err
	}
	// fmt.Println(trx)

	tx, err := handler.db.BeginTx(ctx, nil)
	// fmt.Println(err)
	if err != nil {
		return transaction, transactionDetails, err
	}

	// masukin semua transaction detail ke handler beserta id trx nya
	trxD, err := handler.transactionDetailsRepository.AddTrxDetails(ctx, tx, transactionDetail, *trx.GetId())
	// fmt.Println(err)
	if err != nil {
		tx.Rollback()
		return transaction, transactionDetails, err
	}
	// fmt.Println(trxD)

	tx.Commit()
	transactionDetails = append(transactionDetails, trxD)

	transaction.SetTransactionDetails(transactionDetails)
	return transaction, transactionDetails, nil
}

// eksperimen
// func (handler *transactionHandler) AddTransactionefef(custName, email, phone string, total, pay float64, trxDetails []model.TransactionDetails) error {
// 	// number, date, created_at
// 	ctx := context.Background()

// 	var transaction model.Transaction
// 	var transactionDetails []model.TransactionDetails

// 	transaction.SetCustomerName(&custName)
// 	transaction.SetEmail(&email)
// 	transaction.SetPhone(&phone)
// 	transaction.SetTotal(&total)
// 	transaction.SetPay(&pay)

// 	transaction, err := handler.transactionRepository.AddTrx(ctx, transaction)
// 	if err != nil {
// 		return err
// 	}

// 	tx, err := handler.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	transactionDetails, err = handler.transactionDetailsRepository.AddTrxDetails(ctx, tx, transactionDetails, *transaction.GetId())

// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	tx.Commit()
// 	transaction.SetTransactionDetails(transactionDetails)

// 	return nil
// }
